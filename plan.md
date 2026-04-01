# TxPool Optimization Plan

## Overview

Two optimizations targeting throughput for distributed traffic (many unique senders):
1. **Parallel stateless validation** — decouple cheap pre-checks from the write lock
2. **Pre-read account state before promotion** — move cold statedb reads outside `pool.mu`

---

## Commit 1: Benchmark Test

### File: `core/tx_pool_bench_test.go`

Three timing scenarios. All use `t.Logf` to print results — they never call `t.Fatal` or `t.Error`, so they always pass. Pool setup reuses the existing `testBlockChain` + in-memory statedb pattern from `tx_pool_test.go`.

**Scenario A — Single Sender (hot path baseline)**
- 1 sender, 200 sequential nonces
- Tests throughput when state is already warm (best case)

**Scenario B — Many Unique Senders (distributed traffic)**
- 200 unique senders, 1 tx each
- Each sender is a new address → cold statedb reads on every nonce/balance check
- This is the worst case the optimization targets

**Scenario C — Mixed Load (realistic)**
- 50 unique senders, 4 txs each
- Blend of the above two scenarios

Each scenario:
1. Generates keys + pre-signs all transactions (excluded from timer)
2. Pre-funds each sender in statedb (excluded from timer)
3. Starts `time.Now()`
4. Calls `pool.AddRemotesSync(txs)`
5. Prints elapsed time and per-tx throughput via `t.Logf`

### File: `core/TXPOOL_OPTIMIZATION.md`

README summarizing:
- What was measured
- What was changed and why
- How to run benchmarks and compare before/after

---

## Commit 2: Optimizations

### Change A — Split `validateTx` into stateless + stateful phases

**Location:** `core/tx_pool.go`

**What:** Extract the checks inside `validateTx` that do NOT touch `pool.currentState` into a new `validateTxStateless` function.

Stateless checks (thread-safe, no statedb):
- Transaction type support (EIP-2718 / EIP-1559 fork flags)
- Max transaction size (`txMaxSize`)
- Negative value check
- Gas limit vs `pool.currentMaxGas`
- Gas fee cap / tip sanity (non-negative, tip ≤ feecap)
- Gas price floor (`pool.gasPrice`)
- Sender recovery (already cached by `senderCacher`)
- Local/remote slot limits check via `pool.all.slots`

Stateful checks (remain in `validateTx`, called under lock):
- `pool.currentState.GetNonce(from)` — nonce ordering
- `pool.currentState.GetBalance(from)` — fund sufficiency

**In `addTxs`:** Run `validateTxStateless` concurrently for all transactions before acquiring `pool.mu`. Only transactions passing stateless checks proceed to the locked phase.

```
Before:
  lock → [stateless check + state check + insert] per tx sequentially

After:
  [stateless check] per tx concurrently (no lock)
  → filter failures
  → lock → [state check + insert] per tx sequentially
```

This reduces work done under `pool.mu` proportionally to the stateless-failure rate, and allows CPU parallelism for the stateless phase.

---

### Change B — Pre-read nonces and balances before promotion write lock

**Location:** `core/tx_pool.go`, `runReorg` and `promoteExecutables`

**What:** Before `promoteExecutables` runs under `pool.mu`, snapshot the set of queued addresses (under RLock) and perform all `GetNonce` + `GetBalance` reads outside the write lock, storing results in a `map[common.Address]accountState`.

`promoteExecutables` receives this map and uses the pre-read values instead of calling statedb directly.

```
Before:
  pool.mu.Lock()
    for addr in N accounts:
      GetNonce(addr)    ← may hit trie/disk, serialized
      GetBalance(addr)  ← may hit trie/disk, serialized
      promote(addr)
  pool.mu.Unlock()

After:
  pool.mu.RLock()
  addrs := keys(pool.queue)   ← snapshot addresses
  pool.mu.RUnlock()

  prefetched := map[addr]{nonce, balance}
  for addr in addrs:          ← sequential reads, NO write lock held
    prefetched[addr] = {GetNonce(addr), GetBalance(addr)}

  pool.mu.Lock()
    for addr in accounts:
      use prefetched[addr]    ← cache hit, nanoseconds
      promote(addr)
  pool.mu.Unlock()
```

The write lock hold time drops from `O(N × statedb_read_latency)` to `O(N × map_lookup)`. The total I/O work is unchanged; it is simply moved outside the critical section.

---

## What Does NOT Change

- `txNoncer` logic — untouched
- `txPricedList` — untouched
- Signature recovery batching — already optimal
- The queue/pending data structures — untouched
- Any consensus or validation semantics — purely performance

---

## Files Changed

| File | Change |
|------|--------|
| `core/tx_pool_bench_test.go` | New — benchmark tests (Commit 1) |
| `core/TXPOOL_OPTIMIZATION.md` | New — README (Commit 1) |
| `core/tx_pool.go` | Modified — split validateTx, pre-read in runReorg (Commit 2) |
