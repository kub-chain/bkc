package hardfork

import (
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/state"
	"github.com/ethereum/go-ethereum/log"
)

type HardForkInstruction struct {
	Name    string
	Storage map[common.Address]map[common.Hash]common.Hash
	Code    map[common.Address][]byte
}

func ApplyHardfork(db *state.StateDB, instruction HardForkInstruction) {
	for address, storage := range instruction.Storage {
		for key, value := range storage {
			db.SetState(address, key, value)
			log.Debug("Set storage", "address", address.Hex, "key", key.Hex(), "value", value.Hex())
		}
	}

	for address, bytecode := range instruction.Code {
		db.SetCode(address, bytecode)
	}
}

// IntToHash function is a utility function that allows us to convert
// slot numer to hash easily. Take a note that int is max at 64 bits.
func IntToHash(storageSlot int) common.Hash {
	return common.BytesToHash([]byte{byte(storageSlot)})
}
