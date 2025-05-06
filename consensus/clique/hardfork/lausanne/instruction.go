package lausanne

import (
	"fmt"
	"math/big"

	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/consensus/clique/hardfork"
)

type LausanneParams struct {
	StakeManagerV2        common.Address
	StakeManagerStorageV2 common.Address
	StakeManagerVault     common.Address
	NftContract           common.Address
	SlashManagerV2        common.Address
	KKub                  common.Address
	SlashThreshold        *big.Int
	SlashEpochSize        *big.Int
	SoloSlashRate         *big.Int
}

func New(params LausanneParams) (hardfork.HardForkInstruction, error) {

	if params.StakeManagerV2 == (common.Address{}) {
		return hardfork.HardForkInstruction{}, fmt.Errorf("create Lausanne hardfork requires StakeManagerV2 address")
	}
	if params.StakeManagerStorageV2 == (common.Address{}) {
		return hardfork.HardForkInstruction{}, fmt.Errorf("create Lausanne hardfork requires StakeManagerStorageV2 address")
	}
	if params.StakeManagerVault == (common.Address{}) {
		return hardfork.HardForkInstruction{}, fmt.Errorf("create Lausanne hardfork requires StakeManagerVault address")
	}
	if params.SlashManagerV2 == (common.Address{}) {
		return hardfork.HardForkInstruction{}, fmt.Errorf("create Lausanne hardfork requires SlashManagerV2 address")
	}
	if params.NftContract == (common.Address{}) {
		return hardfork.HardForkInstruction{}, fmt.Errorf("create Lausanne hardfork requires NftContract address")
	}
	if params.KKub == (common.Address{}) {
		return hardfork.HardForkInstruction{}, fmt.Errorf("create Lausanne hardfork requires KKub address")
	}
	if params.SlashThreshold == nil {
		return hardfork.HardForkInstruction{}, fmt.Errorf("create Lausanne hardfork requires SlashThreshold")
	}
	if params.SlashEpochSize == nil {
		return hardfork.HardForkInstruction{}, fmt.Errorf("create Lausanne hardfork requires SlashEpochSize")
	}
	if params.SoloSlashRate == nil {
		return hardfork.HardForkInstruction{}, fmt.Errorf("create Lausanne hardfork requires SoloSlashRate")
	}

	instruction := hardfork.HardForkInstruction{
		Name:    "Lausanne",
		Storage: make(map[common.Address]map[common.Hash]common.Hash),
		Code:    make(map[common.Address][]byte),
	}
	// Set codes
	instruction.Code[params.StakeManagerV2] = StakeManagerV2ByteCode
	instruction.Code[params.StakeManagerStorageV2] = StakeManagerStorageV2ByteCode
	instruction.Code[params.SlashManagerV2] = SlashManagerV2ByteCode

	// Set storages
	// This will set the storage of the contract at stakeManagerV2 with the same value as the v1
	// contract, but with the new variable which can be adjusted after.
	instruction.Storage[params.StakeManagerStorageV2] = map[common.Hash]common.Hash{
		hardfork.IntToHash(24): common.HexToHash("0x64"),                   // soloSlashRate = 100
		hardfork.IntToHash(25): common.HexToHash("0x152D02C7E14AF6800000"), // minimumPoolStake = 100_000 ether
		hardfork.IntToHash(26): common.HexToHash("0x56BC75E2D63100000"),    // minimumPoolDelegate = 100 ether
		hardfork.IntToHash(27): common.HexToHash("0x8AC7230489E80000"),     // minimumSoloStake = 10 ether
		hardfork.IntToHash(28): common.BigToHash(params.SlashThreshold),
		hardfork.IntToHash(29): common.BigToHash(params.SlashEpochSize),
	}
	instruction.Storage[params.StakeManagerV2] = map[common.Hash]common.Hash{
		hardfork.IntToHash(5): params.StakeManagerStorageV2.Hash(), // stakeManagerStorage
		hardfork.IntToHash(6): params.StakeManagerVault.Hash(),     // stakeManagerVault
		hardfork.IntToHash(7): params.NftContract.Hash(),           // nftContract
		hardfork.IntToHash(8): params.KKub.Hash(),                  // kkub
	}
	instruction.Storage[params.SlashManagerV2] = map[common.Hash]common.Hash{
		hardfork.IntToHash(3): params.StakeManagerStorageV2.Hash(), // stakeManagerStorage
	}

	// !! Important notes !!
	// This hard fork is the first time we set the storage of the contract
	// For the best practice, if we want to set the storage slot that exceeds 64 bits,
	// please consider to big number for calculations and use the function `common.BigToHash`
	// to convert them to hash.

	return instruction, nil
}
