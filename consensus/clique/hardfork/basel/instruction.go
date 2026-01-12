package basel

import (
	"fmt"
	"math/big"

	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/consensus/clique/hardfork"
	"github.com/ethereum/go-ethereum/core/state"
	"github.com/ethereum/go-ethereum/crypto"
)

type BaselParams struct {
	StakeManagerV3             common.Address
	StakeManagerStorageV3      common.Address
	BKCValidatorSetV3          common.Address
	StakeManagerVaultV3        common.Address
	NftContractV3              common.Address
	OfficialNodeValidatorShare common.Address
	SlashManagerV3             common.Address
	SuperNodeAddress           common.Address
	SuperNodeOwnerAddress      common.Address
	BaselBlock                 *big.Int
}

// StakeManager storage slots
const (
	SLOT_VALIDATOR_ADDRESS_TO_IDS = 20 // mapping(address => uint256[])
	SLOT_VALIDATOR_LIST           = 21 // Validator[] array
	SLOT_MINIMAL_LIST             = 22 // EnumerableSet array
	SLOT_SUPER_NODE_OLD           = 30 // Old super node address
	SLOT_SUPER_NODE_NEW           = 31 // New super node address
	SLOT_POOL_AMOUNT              = 9  // pool amount
)

// NFT Contract storage slots
const (
	SLOT_HOLDER_TOKENS        = 1
	SLOT_TOKEN_OWNERS         = 2
	SLOT_TOKEN_OWNERS_INDEXES = 3
)

const (
	SLOT_IS_OFFICAL_POOL = 19
)

func New(state *state.StateDB, params BaselParams) (hardfork.HardForkInstruction, error) {

	if params.StakeManagerV3 == (common.Address{}) {
		return hardfork.HardForkInstruction{}, fmt.Errorf("create Lausanne hardfork requires StakeManagerV3 address")
	}
	if params.StakeManagerStorageV3 == (common.Address{}) {
		return hardfork.HardForkInstruction{}, fmt.Errorf("create Lausanne hardfork requires StakeManagerStorageV3 address")
	}
	if params.SlashManagerV3 == (common.Address{}) {
		return hardfork.HardForkInstruction{}, fmt.Errorf("create Lausanne hardfork requires SlashManagerV2 address")
	}
	if params.NftContractV3 == (common.Address{}) {
		return hardfork.HardForkInstruction{}, fmt.Errorf("create Lausanne hardfork requires NftContract address")
	}
	if params.BKCValidatorSetV3 == (common.Address{}) {
		return hardfork.HardForkInstruction{}, fmt.Errorf("create Lausanne hardfork requires BKCValidatorSetV3 address")
	}

	instruction := hardfork.HardForkInstruction{
		Name:    "Basel",
		Storage: make(map[common.Address]map[common.Hash]common.Hash),
		Code:    make(map[common.Address][]byte),
	}
	// Deploy V3 contracts with user's modifications
	instruction.Code[params.StakeManagerV3] = StakeManagerV3ByteCode
	instruction.Code[params.StakeManagerStorageV3] = StakeManagerStorageV3ByteCode
	instruction.Code[params.SlashManagerV3] = SlashManagerV3ByteCode
	instruction.Code[params.NftContractV3] = NftContractV3ByteCode
	instruction.Code[params.BKCValidatorSetV3] = BKCValidatorSetV3ByteCode

	totalSupplySolt := common.BigToHash(big.NewInt(SLOT_TOKEN_OWNERS))
	nextValidatorId := state.GetState(params.NftContractV3, totalSupplySolt).Big().Uint64()

	nextLength := nextValidatorId + 1

	// Read current poolAmount and increment it for official node conversion
	poolAmount := state.GetState(params.StakeManagerStorageV3, common.BigToHash(big.NewInt(SLOT_POOL_AMOUNT))).Big().Uint64()
	newPoolAmount := poolAmount + 1 // Convert official node (validator 0) to pool

	// Get validator 0 slot to update its validatorShareContract
	validator0SlotInt := new(big.Int).SetBytes(getValidatorSlot(SLOT_VALIDATOR_LIST, 0).Bytes())
	// Read current validator 0 slot 5 to preserve status and commission rates
	validator0Slot5 := state.GetState(params.StakeManagerStorageV3, common.BigToHash(new(big.Int).Add(validator0SlotInt, big.NewInt(5))))
	// Extract current status and rates (packed with address in slot 5)
	validator0Slot5Int := validator0Slot5.Big()
	validator0Status := uint8(new(big.Int).Rsh(validator0Slot5Int, 160).Uint64() & 0xFF)
	validator0InfraRate := uint16(new(big.Int).Rsh(validator0Slot5Int, 168).Uint64() & 0xFFFF)
	validator0CommissionRate := uint16(new(big.Int).Rsh(validator0Slot5Int, 184).Uint64() & 0xFFFF)

	nextValidatorSoltInt := new(big.Int).SetBytes(getValidatorSlot(SLOT_VALIDATOR_LIST, nextValidatorId).Bytes()) // For validator ID 1

	validatorIdsArrayLengthSlot := getMappingSlot(common.BytesToHash(params.SuperNodeAddress.Bytes()), SLOT_VALIDATOR_ADDRESS_TO_IDS)

	validatorIdsArrayDataSlot := crypto.Keccak256Hash(validatorIdsArrayLengthSlot.Bytes())
	validatorIdsArrayElement0Slot := common.BigToHash(new(big.Int).SetBytes(validatorIdsArrayDataSlot.Bytes()))

	minimalValidatorListLengthSlot := common.BigToHash(big.NewInt(SLOT_MINIMAL_LIST))
	currentMinimalListLength := state.GetState(params.StakeManagerStorageV3, minimalValidatorListLengthSlot).Big().Uint64()
	minimalValidatorListDataSlotInt := crypto.Keccak256Hash(minimalValidatorListLengthSlot.Bytes()).Big()

	// EnumerableSet._indexes mapping is at base slot + 1 (slot 23 for SLOT_MINIMAL_LIST=22)
	minimalValidatorListIndexesSlot := SLOT_MINIMAL_LIST + 1
	minimalValidatorListDataIndexes := getMappingSlot(common.BytesToHash(params.SuperNodeAddress.Bytes()), uint64(minimalValidatorListIndexesSlot))

	// Calculate superNodeAndValidBlock according to the formula:
	// uint256 tmpUint256 = uint256(uint160(superNode_));
	// tmpUint256 += blockNumber_ << 160;
	tmpUint256 := new(big.Int).SetBytes(params.SuperNodeAddress.Bytes())       // uint256(uint160(superNode_))
	shiftedBlockNumber := new(big.Int).Lsh(params.BaselBlock, 160)             // blockNumber_ << 160
	superNodeAndValidBlock := new(big.Int).Add(tmpUint256, shiftedBlockNumber) // tmpUint256 += blockNumber_ << 160

	instruction.Storage[params.StakeManagerStorageV3] = map[common.Hash]common.Hash{
		// StakeManager reference
		common.BigToHash(big.NewInt(1)): common.BytesToHash(params.StakeManagerV3.Bytes()),

		// Super node configuration
		common.BigToHash(big.NewInt(SLOT_SUPER_NODE_OLD)): common.HexToHash("0x0000000000000000000000000000000000000000000000000000000000000000"),
		common.BigToHash(big.NewInt(SLOT_SUPER_NODE_NEW)): common.BigToHash(superNodeAndValidBlock),

		// New validator struct (6 slots)
		common.BigToHash(new(big.Int).Add(nextValidatorSoltInt, big.NewInt(0))): common.BigToHash(big.NewInt(0)),
		common.BigToHash(new(big.Int).Add(nextValidatorSoltInt, big.NewInt(1))): common.BigToHash(big.NewInt(0)),
		common.BigToHash(new(big.Int).Add(nextValidatorSoltInt, big.NewInt(2))): common.BigToHash(big.NewInt(0)),
		common.BigToHash(new(big.Int).Add(nextValidatorSoltInt, big.NewInt(3))): common.BigToHash(big.NewInt(0)),
		common.BigToHash(new(big.Int).Add(nextValidatorSoltInt, big.NewInt(4))): common.BytesToHash(params.SuperNodeAddress.Bytes()),
		common.BigToHash(new(big.Int).Add(nextValidatorSoltInt, big.NewInt(5))): packValidatorShareAndRates(common.Address{}, 1, 0, 0),

		// Validator list
		common.BigToHash(big.NewInt(SLOT_VALIDATOR_LIST)): common.BigToHash(new(big.Int).SetUint64(nextLength)),

		// Validator IDs mapping
		validatorIdsArrayLengthSlot:   common.BigToHash(big.NewInt(1)),
		validatorIdsArrayElement0Slot: common.BigToHash(new(big.Int).SetUint64(nextValidatorId)),

		// Minimal validator list
		common.BigToHash(minimalValidatorListLengthSlot.Big()):                                                           common.BigToHash(big.NewInt(int64(currentMinimalListLength + 1))),
		common.BigToHash(new(big.Int).Add(minimalValidatorListDataSlotInt, big.NewInt(int64(currentMinimalListLength)))): common.BytesToHash(params.SuperNodeAddress.Bytes()),
		common.BigToHash(minimalValidatorListDataIndexes.Big()):                                                          common.BigToHash(big.NewInt(int64(currentMinimalListLength + 1))),

		// Pool amount
		common.BigToHash(big.NewInt(SLOT_POOL_AMOUNT)): common.BigToHash(big.NewInt(int64(newPoolAmount))),

		// Convert validator 0 to pool
		common.BigToHash(new(big.Int).Add(validator0SlotInt, big.NewInt(5))): packValidatorShareAndRates(params.OfficialNodeValidatorShare, validator0Status, validator0InfraRate, validator0CommissionRate),
	}

	tokenOwnerSlotLength := common.BigToHash(big.NewInt(SLOT_TOKEN_OWNERS))
	tokenOwnerDataSlot := crypto.Keccak256Hash(tokenOwnerSlotLength.Bytes())
	tokenOwnerDataSlotInt := new(big.Int).Add(
		tokenOwnerDataSlot.Big(),
		big.NewInt(int64(nextValidatorId)*2))

	tokenIdHash := common.BigToHash(big.NewInt(int64(nextValidatorId)))

	indexesSlot := getMappingSlot(tokenIdHash, SLOT_TOKEN_OWNERS_INDEXES)

	// Get existing holder token count
	holderArraySlot := getMappingSlot(common.BytesToHash(params.SuperNodeOwnerAddress.Bytes()), SLOT_HOLDER_TOKENS)
	currentHolderTokenCount := state.GetState(params.NftContractV3, holderArraySlot).Big().Uint64()
	newHolderTokenCount := currentHolderTokenCount + 1

	holderArraySlotInt := holderArraySlot.Big()
	holderArrayDataSlotInt := crypto.Keccak256Hash(holderArraySlot.Bytes()).Big()

	// Add new token at the current count position (0-indexed)
	newTokenPosition := currentHolderTokenCount
	holderTokenArrayElementSlot := new(big.Int).Add(holderArrayDataSlotInt, big.NewInt(int64(newTokenPosition)))

	holderIndexesSlot := GetHolderTokenIndexSlot(holderArraySlot, uint64(nextValidatorId))

	// Configure BKCValidatorSet storage
	instruction.Storage[params.BKCValidatorSetV3] = map[common.Hash]common.Hash{
		common.BigToHash(big.NewInt(3)): common.BytesToHash(params.StakeManagerStorageV3.Bytes()),
	}

	// Configure NFT contract storage
	instruction.Storage[params.NftContractV3] = map[common.Hash]common.Hash{
		// StakeManager reference
		common.BigToHash(big.NewInt(12)): common.BytesToHash(params.StakeManagerStorageV3.Bytes()),

		// Token supply
		common.BigToHash(big.NewInt(SLOT_TOKEN_OWNERS)): common.BigToHash(big.NewInt(int64(nextLength))),

		// Token ownership entries
		common.BigToHash(new(big.Int).Add(tokenOwnerDataSlotInt, big.NewInt(0))): common.BigToHash(big.NewInt(int64(nextValidatorId))),
		common.BigToHash(new(big.Int).Add(tokenOwnerDataSlotInt, big.NewInt(1))): common.BytesToHash(params.SuperNodeOwnerAddress.Bytes()),
		common.BigToHash(indexesSlot.Big()):                                      common.BigToHash(big.NewInt(int64(nextLength))),

		// Holder token array
		common.BigToHash(holderArraySlotInt):          common.BigToHash(big.NewInt(int64(newHolderTokenCount))),
		common.BigToHash(holderTokenArrayElementSlot): common.BigToHash(big.NewInt(int64(nextValidatorId))),
		common.BigToHash(holderIndexesSlot.Big()):     common.BigToHash(big.NewInt(int64(newHolderTokenCount))),
	}

	// Configure official node validator share
	isOfficialData := state.GetState(params.OfficialNodeValidatorShare, common.BigToHash(big.NewInt(SLOT_IS_OFFICAL_POOL)))

	// Set isOfficalPool to false
	isOfficialDataBytes := isOfficialData.Bytes()
	isOfficialDataBytes[11] = 0x00

	instruction.Storage[params.OfficialNodeValidatorShare] = map[common.Hash]common.Hash{
		common.BigToHash(big.NewInt(SLOT_IS_OFFICAL_POOL)): common.BytesToHash(isOfficialDataBytes),
	}

	// Preserve existing StakeManagerV3 addresses
	committeeAddress := state.GetState(params.StakeManagerV3, common.BigToHash(big.NewInt(1)))
	callHelperAddress := state.GetState(params.StakeManagerV3, common.BigToHash(big.NewInt(2)))
	transferRouterAddress := state.GetState(params.StakeManagerV3, common.BigToHash(big.NewInt(3)))
	officialPoolStakerAddress := state.GetState(params.StakeManagerV3, common.BigToHash(big.NewInt(4)))
	kkubAddress := state.GetState(params.StakeManagerV3, common.BigToHash(big.NewInt(8)))

	// Configure StakeManagerV3 storage
	instruction.Storage[params.StakeManagerV3] = map[common.Hash]common.Hash{
		// Reentrancy guard
		common.BigToHash(big.NewInt(0)): common.BigToHash(big.NewInt(1)),

		// Preserved addresses
		common.BigToHash(big.NewInt(1)): committeeAddress,
		common.BigToHash(big.NewInt(2)): callHelperAddress,
		common.BigToHash(big.NewInt(3)): transferRouterAddress,
		common.BigToHash(big.NewInt(4)): officialPoolStakerAddress,
		common.BigToHash(big.NewInt(8)): kkubAddress,

		// New contract references
		common.BigToHash(big.NewInt(5)): common.BytesToHash(params.StakeManagerStorageV3.Bytes()),
		common.BigToHash(big.NewInt(6)): common.BytesToHash(params.StakeManagerVaultV3.Bytes()),
		common.BigToHash(big.NewInt(7)): common.BytesToHash(params.NftContractV3.Bytes()),
	}

	return instruction, nil
}

func getMappingSlot(key common.Hash, baseSlot uint64) common.Hash {
	slotBytes := common.LeftPadBytes(new(big.Int).SetUint64(baseSlot).Bytes(), 32)
	data := append(key.Bytes(), slotBytes...)
	return crypto.Keccak256Hash(data)
}

func getValidatorSlot(baseSlot uint64, validatorId uint64) common.Hash {
	slotBytes := common.LeftPadBytes(new(big.Int).SetUint64(baseSlot).Bytes(), 32)
	arrayBase := crypto.Keccak256Hash(slotBytes)

	arrayBaseInt := new(big.Int).SetBytes(arrayBase.Bytes())
	// Each Validator struct takes 6 storage slots (0-5)
	offset := new(big.Int).Mul(new(big.Int).SetUint64(validatorId), big.NewInt(6))
	validatorSlot := new(big.Int).Add(arrayBaseInt, offset)

	return common.BytesToHash(validatorSlot.Bytes())
}

// GetHolderTokenIndexSlot returns the storage slot for _holderTokens[holder]._inner._indexes[tokenId]
// EnumerableSet.UintSet structure:
// - _values array at slot 0 (relative to set base)
// - _indexes mapping at slot 1 (relative to set base)
func GetHolderTokenIndexSlot(holderSetBaseSlot common.Hash, tokenId uint64) common.Hash {
	// _indexes is at holderSetBaseSlot + 1
	indexesSlotBase := new(big.Int).Add(holderSetBaseSlot.Big(), big.NewInt(1))

	// mapping(uint256 => uint256) lookup for tokenId
	data := make([]byte, 64)
	tokenIdBytes := common.LeftPadBytes(new(big.Int).SetUint64(tokenId).Bytes(), 32)
	copy(data[0:32], tokenIdBytes)
	indexesSlotBytes := common.LeftPadBytes(indexesSlotBase.Bytes(), 32)
	copy(data[32:64], indexesSlotBytes)

	return crypto.Keccak256Hash(data)
}

// GetTokenOwnersIndexSlot returns the storage slot for _indexes[tokenId] at SLOT_TOKEN_OWNERS_INDEXES
// This is for the separate _indexes mapping at slot 3
func GetTokenOwnersIndexSlot(tokenId uint64) common.Hash {
	// mapping(uint256 => uint256) at SLOT_TOKEN_OWNERS_INDEXES (slot 3)
	data := make([]byte, 64)
	tokenIdBytes := common.LeftPadBytes(new(big.Int).SetUint64(tokenId).Bytes(), 32)
	copy(data[0:32], tokenIdBytes)
	slotBytes := common.LeftPadBytes(new(big.Int).SetUint64(SLOT_TOKEN_OWNERS_INDEXES).Bytes(), 32)
	copy(data[32:64], slotBytes)

	return crypto.Keccak256Hash(data)
}

// packValidatorShareAndRates packs validatorShareContract, status, and commission rates into one slot (DEPRECATED - now split into slot 5 and 6)
// Layout: validatorShareContract (160 bits) + status (8 bits) + infraCommissionRate (16 bits) + commissionRate (16 bits) = 200 bits
func packValidatorShareAndRates(addr common.Address, status uint8, infraRate uint16, commissionRate uint16) common.Hash {
	result := new(big.Int)
	// validatorShareContract address in lower 160 bits
	result.Or(result, new(big.Int).SetBytes(addr.Bytes()))
	// Status in next 8 bits (shift left by 160)
	result.Or(result, new(big.Int).Lsh(new(big.Int).SetUint64(uint64(status)), 160))
	// infraCommissionRate in next 16 bits (shift left by 168)
	result.Or(result, new(big.Int).Lsh(new(big.Int).SetUint64(uint64(infraRate)), 168))
	// commissionRate in next 16 bits (shift left by 184)
	result.Or(result, new(big.Int).Lsh(new(big.Int).SetUint64(uint64(commissionRate)), 184))
	return common.BigToHash(result)
}
