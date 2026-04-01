# TxPool Optimization

## Background

The transaction pool (`core/tx_pool.go`) processes incoming transactions from the network. Under distributed traffic — many unique sender addresses each sending one transaction — two bottlenecks emerged:

1. **Stateless checks run under the write lock.** Checks like size, transaction type, gas price floor, and fee cap validation do not touch shared pool state, but they execute sequentially inside `pool.mu.Lock()` along with every other validation step.

2. **Cold statedb reads hold the write lock during promotion.** `promoteExecutables` calls `GetNonce` and `GetBalance` for every queued account while `pool.mu` is held. For N unique senders with uncached state this means N sequential disk/trie reads inside the critical section, blocking all concurrent transaction submissions for the full duration.

---

## Changes

### Change A — Parallel stateless validation (`validateTxStateless`)

**File:** `core/tx_pool.go`

A new `validateTxStateless` function extracts all checks from `validateTx` that do not require `pool.currentState`:

- Transaction type support (EIP-2718 / EIP-1559 fork flags)
- Maximum transaction size
- Negative value guard
- Gas limit vs `pool.currentMaxGas`
- Fee cap / tip sanity (bit length, tip ≤ feecap)
- Gas price floor vs `pool.gasPrice`
- Sender recovery (already concurrently cached by `senderCacher`)

`addTxs` now runs `validateTxStateless` concurrently across all incoming transactions before acquiring `pool.mu`. Transactions that fail are rejected early. Only survivors reach the locked phase where stateful checks (nonce ordering, balance sufficiency) and insertion happen.

```
Before:
  lock → [stateless + stateful + insert] × N  (sequential)

After:
  [stateless] × N  (concurrent, no lock)
  → filter
  lock → [stateful + insert] × survivors  (sequential)
```

### Change B — Pre-read account state before promotion (`runReorg`)

**File:** `core/tx_pool.go`

`runReorg` now reads nonces and balances for all queued addresses **outside** `pool.mu` before calling `promoteExecutables`. The results are passed as a pre-fetched map, and `promoteExecutables` uses those values instead of calling statedb directly.

```
Before:
  pool.mu.Lock()
    for addr in N accounts:
      GetNonce(addr)    ← disk/trie, serialized under write lock
      GetBalance(addr)  ← disk/trie, serialized under write lock
      promote(addr)
  pool.mu.Unlock()

After:
  pool.mu.RLock()
  addrs := keys(pool.queue)
  pool.mu.RUnlock()

  for addr in addrs:             ← I/O outside write lock
    prefetch[addr] = {nonce, balance}

  pool.mu.Lock()
    for addr in accounts:
      use prefetch[addr]         ← map lookup, nanoseconds
      promote(addr)
  pool.mu.Unlock()
```

The total number of statedb reads is unchanged. The write lock hold time drops from `O(N × disk_latency)` to `O(N × map_lookup)`.

---

## What Was Not Changed

- `txNoncer` caching logic
- `txPricedList` heap operations
- Signature recovery (already batched and concurrent)
- Queue / pending data structures
- All consensus and validation semantics

---

## Benchmark

### Running the benchmarks

```bash
# Run all three scenarios
go test ./core/ -run "TestTxPoolBenchmark" -v -count=1

# Run a specific scenario
go test ./core/ -run "TestTxPoolBenchmark_ManyUniqueSenders" -v -count=1
```

### Scenarios

| Scenario | Senders | Tx per sender | Total txs | What it measures |
|---|---|---|---|---|
| `SingleSender` | 1 | 200 | 200 | Hot/warm baseline — state cached after first read |
| `ManyUniqueSenders` | 200 | 1 | 200 | Cold path — every sender is a fresh address |
| `MixedLoad` | 50 | 4 | 200 | Realistic blend of new and returning accounts |

### Comparing before and after

```bash
# On the benchmark commit (original code):
git checkout <benchmark-commit>
go test ./core/ -run "TestTxPoolBenchmark" -v -count=3 2>&1 | grep BENCH

# On the optimization commit:
git checkout <optimization-commit>
go test ./core/ -run "TestTxPoolBenchmark" -v -count=3 2>&1 | grep BENCH
```

The `ManyUniqueSenders` scenario shows the largest improvement because it maximises cold statedb reads — exactly the path targeted by both optimizations.
