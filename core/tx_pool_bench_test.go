package core

import (
	"fmt"
	"math/big"
	"testing"
	"time"

	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/rawdb"
	"github.com/ethereum/go-ethereum/core/state"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/crypto"
	"github.com/ethereum/go-ethereum/event"
	"github.com/ethereum/go-ethereum/params"
)

// benchFundAccount sets a large balance on addr in statedb so all txs pass
// the balance check during validation.
func benchFundAccount(statedb *state.StateDB, addr common.Address) {
	statedb.AddBalance(addr, big.NewInt(0).Mul(big.NewInt(1e18), big.NewInt(1000)))
}

// newBenchPool creates a minimal TxPool backed by an in-memory statedb.
// Callers receive the pool and the underlying statedb so they can fund accounts.
func newBenchPool(t *testing.T) (*TxPool, *state.StateDB) {
	t.Helper()
	statedb, err := state.New(common.Hash{}, state.NewDatabase(rawdb.NewMemoryDatabase()), nil)
	if err != nil {
		t.Fatalf("failed to create statedb: %v", err)
	}
	blockchain := &testBlockChain{10_000_000, statedb, new(event.Feed)}
	cfg := testTxPoolConfig
	cfg.GlobalSlots = 8192
	cfg.GlobalQueue = 4096
	pool := NewTxPool(cfg, params.TestChainConfig, blockchain)
	<-pool.initDoneCh
	return pool, statedb
}

// printResult logs timing and throughput in a consistent format.
func printResult(t *testing.T, scenario string, txCount int, elapsed time.Duration) {
	t.Helper()
	perTx := elapsed / time.Duration(txCount)
	throughput := float64(txCount) / elapsed.Seconds()
	t.Logf("[BENCH] %-40s | txs=%-4d | total=%-12s | per-tx=%-10s | throughput=%.0f tx/s",
		scenario, txCount, elapsed.Round(time.Microsecond), perTx.Round(time.Microsecond), throughput)
	fmt.Printf("[BENCH] %-40s | txs=%-4d | total=%-12s | per-tx=%-10s | throughput=%.0f tx/s\n",
		scenario, txCount, elapsed.Round(time.Microsecond), perTx.Round(time.Microsecond), throughput)
}

// ----------------------------------------------------------------------------
// Scenario A: Single Sender — warm/hot path baseline
//
// One account sends 200 sequential transactions. The state for this account
// is cached after the first lookup, so this represents the best-case path
// (minimal cold storage reads).
// ----------------------------------------------------------------------------
func TestTxPoolBenchmark_SingleSender(t *testing.T) {
	const txCount = 200

	pool, statedb := newBenchPool(t)
	defer pool.Stop()

	key, _ := crypto.GenerateKey()
	addr := crypto.PubkeyToAddress(key.PublicKey)
	benchFundAccount(statedb, addr)

	// Pre-sign all transactions — excluded from timer.
	txs := make([]*types.Transaction, txCount)
	for i := 0; i < txCount; i++ {
		txs[i] = pricedTransaction(uint64(i), 21000, big.NewInt(1), key)
	}

	start := time.Now()
	pool.AddRemotesSync(txs)
	elapsed := time.Since(start)

	printResult(t, "SingleSender (hot path)", txCount, elapsed)
}

// ----------------------------------------------------------------------------
// Scenario B: Many Unique Senders — cold path (distributed traffic)
//
// 200 distinct accounts each send 1 transaction. Every sender is a fresh
// address, so each nonce and balance check hits a cold statedb object.
// This is the worst-case scenario that the prefetch optimisation targets.
// ----------------------------------------------------------------------------
func TestTxPoolBenchmark_ManyUniqueSenders(t *testing.T) {
	const senderCount = 200

	pool, statedb := newBenchPool(t)
	defer pool.Stop()

	// Pre-generate keys, fund accounts, and sign transactions — all excluded
	// from the timer.
	txs := make([]*types.Transaction, senderCount)
	for i := 0; i < senderCount; i++ {
		key, _ := crypto.GenerateKey()
		addr := crypto.PubkeyToAddress(key.PublicKey)
		benchFundAccount(statedb, addr)
		txs[i] = pricedTransaction(0, 21000, big.NewInt(1), key)
	}

	start := time.Now()
	pool.AddRemotesSync(txs)
	elapsed := time.Since(start)

	printResult(t, "ManyUniqueSenders (cold path)", senderCount, elapsed)
}

// ----------------------------------------------------------------------------
// Scenario C: Mixed Load — realistic traffic pattern
//
// 50 unique senders each submit 4 transactions (sequential nonces).
// Represents a realistic mix of returning and new accounts.
// ----------------------------------------------------------------------------
func TestTxPoolBenchmark_MixedLoad(t *testing.T) {
	const (
		senderCount = 50
		txPerSender = 4
		txCount     = senderCount * txPerSender
	)

	pool, statedb := newBenchPool(t)
	defer pool.Stop()

	txs := make([]*types.Transaction, 0, txCount)
	for i := 0; i < senderCount; i++ {
		key, _ := crypto.GenerateKey()
		addr := crypto.PubkeyToAddress(key.PublicKey)
		benchFundAccount(statedb, addr)
		for n := 0; n < txPerSender; n++ {
			txs = append(txs, pricedTransaction(uint64(n), 21000, big.NewInt(1), key))
		}
	}

	start := time.Now()
	pool.AddRemotesSync(txs)
	elapsed := time.Since(start)

	printResult(t, "MixedLoad (50 senders x 4 txs)", txCount, elapsed)
}
