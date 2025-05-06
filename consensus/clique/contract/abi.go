package contract

const stakeManageABI = `[
  {
    "inputs": [],
    "name": "distributeReward",
    "outputs": [],
    "stateMutability": "payable",
    "type": "function"
  },
  {
    "type": "function",
    "name": "stakeManagerStorage",
    "inputs": [],
    "outputs": [{ "name": "", "type": "address", "internalType": "contract IStakeManagerStorage" }],
    "stateMutability": "view"
  },
  {
    "type": "function",
    "name": "stakeManagerVault",
    "inputs": [],
    "outputs": [{ "name": "", "type": "address", "internalType": "contract IStakeManagerVault" }],
    "stateMutability": "view"
  },
  {
    "type": "function",
    "name": "nftContract",
    "inputs": [],
    "outputs": [{ "name": "", "type": "address", "internalType": "contract StakingNFT" }],
    "stateMutability": "view"
  },
  {
    "type": "function",
    "name": "kkub",
    "inputs": [],
    "outputs": [{ "name": "", "type": "address", "internalType": "contract IKKUB" }],
    "stateMutability": "view"
  }
]`
const validatorSetABI = `[
    {
      "inputs": [
        {
          "internalType": "address",
          "name": "stakeManagerStorage_",
          "type": "address"
        },
        {
          "internalType": "uint256",
          "name": "initialSpanBlock_",
          "type": "uint256"
        },
        {
          "internalType": "bytes",
          "name": "initialVbytes_",
          "type": "bytes"
        }
      ],
      "stateMutability": "nonpayable",
      "type": "constructor"
    },
    {
      "anonymous": false,
      "inputs": [
        {
          "indexed": true,
          "internalType": "uint256",
          "name": "span",
          "type": "uint256"
        },
        {
          "indexed": true,
          "internalType": "uint256",
          "name": "start",
          "type": "uint256"
        },
        {
          "indexed": true,
          "internalType": "uint256",
          "name": "end",
          "type": "uint256"
        }
      ],
      "name": "CommitSpan",
      "type": "event"
    },
    {
      "anonymous": false,
      "inputs": [
        {
          "indexed": true,
          "internalType": "uint256",
          "name": "span",
          "type": "uint256"
        },
        {
          "indexed": false,
          "internalType": "bytes",
          "name": "vbytes",
          "type": "bytes"
        }
      ],
      "name": "SetValidators",
      "type": "event"
    },
    {
      "inputs": [],
      "name": "SPAN",
      "outputs": [
        {
          "internalType": "uint256",
          "name": "",
          "type": "uint256"
        }
      ],
      "stateMutability": "view",
      "type": "function"
    },
    {
      "inputs": [],
      "name": "calculateTransitionSpanCommitmentBlock",
      "outputs": [
        {
          "internalType": "uint256",
          "name": "",
          "type": "uint256"
        }
      ],
      "stateMutability": "view",
      "type": "function"
    },
    {
      "inputs": [
        {
          "internalType": "bytes",
          "name": "validatorBytes_",
          "type": "bytes"
        }
      ],
      "name": "commitSpan",
      "outputs": [],
      "stateMutability": "nonpayable",
      "type": "function"
    },
    {
      "inputs": [],
      "name": "currentSpanNumber",
      "outputs": [
        {
          "internalType": "uint256",
          "name": "",
          "type": "uint256"
        }
      ],
      "stateMutability": "view",
      "type": "function"
    },
    {
      "inputs": [
        {
          "internalType": "uint256",
          "name": "number_",
          "type": "uint256"
        }
      ],
      "name": "getCommitmentBlock",
      "outputs": [
        {
          "internalType": "uint256",
          "name": "",
          "type": "uint256"
        }
      ],
      "stateMutability": "view",
      "type": "function"
    },
    {
      "inputs": [],
      "name": "getEligibleValidators",
      "outputs": [
        {
          "components": [
            {
              "internalType": "address",
              "name": "signer",
              "type": "address"
            },
            {
              "internalType": "uint256",
              "name": "power",
              "type": "uint256"
            }
          ],
          "internalType": "struct MinimalValidator[]",
          "name": "",
          "type": "tuple[]"
        }
      ],
      "stateMutability": "view",
      "type": "function"
    },
    {
      "inputs": [],
      "name": "getOfficialPool",
      "outputs": [
        {
          "internalType": "address",
          "name": "",
          "type": "address"
        }
      ],
      "stateMutability": "view",
      "type": "function"
    },
    {
      "inputs": [],
      "name": "getSlashManager",
      "outputs": [
        {
          "internalType": "address",
          "name": "",
          "type": "address"
        }
      ],
      "stateMutability": "view",
      "type": "function"
    },
    {
      "inputs": [
        {
          "internalType": "uint256",
          "name": "number_",
          "type": "uint256"
        }
      ],
      "name": "getSpanByBlock",
      "outputs": [
        {
          "internalType": "uint256",
          "name": "",
          "type": "uint256"
        }
      ],
      "stateMutability": "view",
      "type": "function"
    },
    {
      "inputs": [
        {
          "internalType": "uint256",
          "name": "number_",
          "type": "uint256"
        }
      ],
      "name": "getSpanRange",
      "outputs": [
        {
          "internalType": "uint256",
          "name": "startBlock",
          "type": "uint256"
        },
        {
          "internalType": "uint256",
          "name": "endBlock",
          "type": "uint256"
        }
      ],
      "stateMutability": "view",
      "type": "function"
    },
    {
      "inputs": [],
      "name": "getStakeManager",
      "outputs": [
        {
          "internalType": "address",
          "name": "",
          "type": "address"
        }
      ],
      "stateMutability": "view",
      "type": "function"
    },
    {
      "inputs": [
        {
          "internalType": "uint256",
          "name": "number_",
          "type": "uint256"
        }
      ],
      "name": "getValidators",
      "outputs": [
        {
          "internalType": "address[]",
          "name": "",
          "type": "address[]"
        },
        {
          "internalType": "uint256[]",
          "name": "",
          "type": "uint256[]"
        },
        {
          "internalType": "address[3]",
          "name": "",
          "type": "address[3]"
        }
      ],
      "stateMutability": "view",
      "type": "function"
    },
    {
      "inputs": [],
      "name": "initialSpanBlock",
      "outputs": [
        {
          "internalType": "uint256",
          "name": "",
          "type": "uint256"
        }
      ],
      "stateMutability": "view",
      "type": "function"
    },
    {
      "inputs": [
        {
          "internalType": "address",
          "name": "signer_",
          "type": "address"
        }
      ],
      "name": "isCurrentValidator",
      "outputs": [
        {
          "internalType": "bool",
          "name": "",
          "type": "bool"
        }
      ],
      "stateMutability": "view",
      "type": "function"
    },
    {
      "inputs": [
        {
          "internalType": "uint256",
          "name": "span_",
          "type": "uint256"
        },
        {
          "internalType": "address",
          "name": "signer_",
          "type": "address"
        }
      ],
      "name": "isValidator",
      "outputs": [
        {
          "internalType": "bool",
          "name": "",
          "type": "bool"
        }
      ],
      "stateMutability": "view",
      "type": "function"
    },
    {
      "inputs": [
        {
          "internalType": "uint256",
          "name": "",
          "type": "uint256"
        }
      ],
      "name": "spans",
      "outputs": [
        {
          "internalType": "uint256",
          "name": "number",
          "type": "uint256"
        },
        {
          "internalType": "uint256",
          "name": "startBlock",
          "type": "uint256"
        },
        {
          "internalType": "uint256",
          "name": "endBlock",
          "type": "uint256"
        }
      ],
      "stateMutability": "view",
      "type": "function"
    },
    {
      "inputs": [],
      "name": "stableSpanBlock",
      "outputs": [
        {
          "internalType": "uint256",
          "name": "",
          "type": "uint256"
        }
      ],
      "stateMutability": "view",
      "type": "function"
    },
    {
      "inputs": [],
      "name": "stakeManagerStorage",
      "outputs": [
        {
          "internalType": "contract IStakeManagerStorage",
          "name": "",
          "type": "address"
        }
      ],
      "stateMutability": "view",
      "type": "function"
    },
    {
      "inputs": [
        {
          "internalType": "uint256",
          "name": "",
          "type": "uint256"
        }
      ],
      "name": "validators",
      "outputs": [
        {
          "internalType": "bytes",
          "name": "",
          "type": "bytes"
        }
      ],
      "stateMutability": "view",
      "type": "function"
    }
  ]`
const slashABI = `[{"type":"constructor","inputs":[{"name":"stakeManagerStorage_","type":"address","internalType":"address"},{"name":"initialSpanBlock_","type":"uint256","internalType":"uint256"}],"stateMutability":"nonpayable"},{"type":"function","name":"SPAN","inputs":[],"outputs":[{"name":"","type":"uint256","internalType":"uint256"}],"stateMutability":"view"},{"type":"function","name":"calculateTransitionSpanCommitmentBlock","inputs":[],"outputs":[{"name":"","type":"uint256","internalType":"uint256"}],"stateMutability":"view"},{"type":"function","name":"currentSpanNumber","inputs":[],"outputs":[{"name":"","type":"uint256","internalType":"uint256"}],"stateMutability":"view"},{"type":"function","name":"epoch","inputs":[{"name":"","type":"address","internalType":"address"}],"outputs":[{"name":"lastEpochStart","type":"uint256","internalType":"uint256"},{"name":"lastSlash","type":"uint256","internalType":"uint256"},{"name":"counter","type":"uint256","internalType":"uint256"}],"stateMutability":"view"},{"type":"function","name":"getCommitmentBlock","inputs":[{"name":"number_","type":"uint256","internalType":"uint256"}],"outputs":[{"name":"","type":"uint256","internalType":"uint256"}],"stateMutability":"view"},{"type":"function","name":"getSpanByBlock","inputs":[{"name":"number_","type":"uint256","internalType":"uint256"}],"outputs":[{"name":"","type":"uint256","internalType":"uint256"}],"stateMutability":"view"},{"type":"function","name":"getSpanRange","inputs":[{"name":"number_","type":"uint256","internalType":"uint256"}],"outputs":[{"name":"startBlock","type":"uint256","internalType":"uint256"},{"name":"endBlock","type":"uint256","internalType":"uint256"}],"stateMutability":"view"},{"type":"function","name":"initialSpanBlock","inputs":[],"outputs":[{"name":"","type":"uint256","internalType":"uint256"}],"stateMutability":"view"},{"type":"function","name":"isSignerSlashed","inputs":[{"name":"signer_","type":"address","internalType":"address"},{"name":"span_","type":"uint256","internalType":"uint256"}],"outputs":[{"name":"","type":"bool","internalType":"bool"}],"stateMutability":"view"},{"type":"function","name":"maxEpochSize","inputs":[],"outputs":[{"name":"","type":"uint256","internalType":"uint256"}],"stateMutability":"view"},{"type":"function","name":"slash","inputs":[{"name":"signer_","type":"address","internalType":"address"},{"name":"currentSpan_","type":"uint256","internalType":"uint256"}],"outputs":[{"name":"","type":"bool","internalType":"bool"}],"stateMutability":"nonpayable"},{"type":"function","name":"span","inputs":[{"name":"","type":"address","internalType":"address"},{"name":"","type":"uint256","internalType":"uint256"}],"outputs":[{"name":"","type":"bool","internalType":"bool"}],"stateMutability":"view"},{"type":"function","name":"stableSpanBlock","inputs":[],"outputs":[{"name":"","type":"uint256","internalType":"uint256"}],"stateMutability":"view"},{"type":"function","name":"stakeManagerStorage","inputs":[],"outputs":[{"name":"","type":"address","internalType":"contract IStakeManagerStorageV2"}],"stateMutability":"view"},{"type":"function","name":"threshold","inputs":[],"outputs":[{"name":"","type":"uint256","internalType":"uint256"}],"stateMutability":"view"},{"type":"event","name":"Slash","inputs":[{"name":"signer","type":"address","indexed":true,"internalType":"address"},{"name":"span","type":"uint256","indexed":true,"internalType":"uint256"},{"name":"counter","type":"uint256","indexed":true,"internalType":"uint256"}],"anonymous":false},{"type":"event","name":"StartEpoch","inputs":[{"name":"signer","type":"address","indexed":true,"internalType":"address"},{"name":"span","type":"uint256","indexed":true,"internalType":"uint256"}],"anonymous":false},{"type":"event","name":"Warn","inputs":[{"name":"signer","type":"address","indexed":true,"internalType":"address"},{"name":"span","type":"uint256","indexed":true,"internalType":"uint256"},{"name":"counter","type":"uint256","indexed":true,"internalType":"uint256"}],"anonymous":false}]`
const stakeManagerStorageABI = `[{"type":"constructor","inputs":[{"name":"input_","type":"tuple","internalType":"struct StakeManagerStorageConstructorInput","components":[{"name":"committee","type":"address","internalType":"address"},{"name":"defaultInfraCommissionRate","type":"uint256","internalType":"uint256"},{"name":"soloSlashRate","type":"uint256","internalType":"uint256"},{"name":"poolSlashAmount","type":"uint256","internalType":"uint256"},{"name":"officialPool","type":"address","internalType":"address"}]}],"stateMutability":"nonpayable"},{"type":"function","name":"DEFAULT_SOLO_SLASH_RATE","inputs":[],"outputs":[{"name":"","type":"uint256","internalType":"uint256"}],"stateMutability":"view"},{"type":"function","name":"INCORRECT_VALIDATOR_ID","inputs":[],"outputs":[{"name":"","type":"uint256","internalType":"uint256"}],"stateMutability":"view"},{"type":"function","name":"MAX_COMMISSION_RATE","inputs":[],"outputs":[{"name":"","type":"uint16","internalType":"uint16"}],"stateMutability":"view"},{"type":"function","name":"MAX_INFRA_COMMISSION_RATE","inputs":[],"outputs":[{"name":"","type":"uint16","internalType":"uint16"}],"stateMutability":"view"},{"type":"function","name":"MAX_RATE","inputs":[],"outputs":[{"name":"","type":"uint16","internalType":"uint16"}],"stateMutability":"view"},{"type":"function","name":"activateValidator","inputs":[{"name":"_id","type":"uint256","internalType":"uint256"}],"outputs":[],"stateMutability":"nonpayable"},{"type":"function","name":"addMinimalValidator","inputs":[{"name":"val_","type":"address","internalType":"address"}],"outputs":[],"stateMutability":"nonpayable"},{"type":"function","name":"addValidator","inputs":[{"name":"key_","type":"address","internalType":"address"},{"name":"val_","type":"tuple","internalType":"struct Validator","components":[{"name":"amount","type":"uint128","internalType":"uint128"},{"name":"delegatedAmount","type":"uint128","internalType":"uint128"},{"name":"reward","type":"uint128","internalType":"uint128"},{"name":"delegatorsReward","type":"uint128","internalType":"uint128"},{"name":"infraCommissionAmount","type":"uint128","internalType":"uint128"},{"name":"validatorCommissionAmount","type":"uint128","internalType":"uint128"},{"name":"delegatorCommissionAmount","type":"uint128","internalType":"uint128"},{"name":"minDeposit","type":"uint128","internalType":"uint128"},{"name":"signer","type":"address","internalType":"address"},{"name":"validatorShareContract","type":"address","internalType":"address"},{"name":"status","type":"uint8","internalType":"enum Status"},{"name":"infraCommissionRate","type":"uint16","internalType":"uint16"},{"name":"commissionRate","type":"uint16","internalType":"uint16"}]}],"outputs":[{"name":"","type":"uint256","internalType":"uint256"}],"stateMutability":"nonpayable"},{"type":"function","name":"changeOfficialSigner","inputs":[{"name":"newOfficialSigner_","type":"address","internalType":"address"},{"name":"checksummedNewOfficialSigner_","type":"string","internalType":"string"}],"outputs":[],"stateMutability":"nonpayable"},{"type":"function","name":"committee","inputs":[],"outputs":[{"name":"","type":"address","internalType":"address"}],"stateMutability":"view"},{"type":"function","name":"defaultInfraCommissionRate","inputs":[],"outputs":[{"name":"","type":"uint256","internalType":"uint256"}],"stateMutability":"view"},{"type":"function","name":"getAllMinimalValidators","inputs":[],"outputs":[{"name":"","type":"address[]","internalType":"address[]"}],"stateMutability":"view"},{"type":"function","name":"getAllValidator","inputs":[],"outputs":[{"name":"validatorAddresses","type":"address[]","internalType":"address[]"}],"stateMutability":"view"},{"type":"function","name":"getMinimalValidatorByIndex","inputs":[{"name":"index_","type":"uint256","internalType":"uint256"}],"outputs":[{"name":"","type":"address","internalType":"address"}],"stateMutability":"view"},{"type":"function","name":"getMinimalValidatorIndex","inputs":[{"name":"val_","type":"address","internalType":"address"},{"name":"revertIfNotFound_","type":"bool","internalType":"bool"}],"outputs":[{"name":"","type":"uint256","internalType":"uint256"},{"name":"","type":"bool","internalType":"bool"}],"stateMutability":"view"},{"type":"function","name":"getMinimalValidatorListLength","inputs":[],"outputs":[{"name":"","type":"uint256","internalType":"uint256"}],"stateMutability":"view"},{"type":"function","name":"getMinimalValidatorsByPage","inputs":[{"name":"page","type":"uint256","internalType":"uint256"},{"name":"limit","type":"uint256","internalType":"uint256"}],"outputs":[{"name":"","type":"address[]","internalType":"address[]"}],"stateMutability":"view"},{"type":"function","name":"getMinimalValidatorsWithValidatorPowerByPage","inputs":[{"name":"page_","type":"uint256","internalType":"uint256"},{"name":"limit_","type":"uint256","internalType":"uint256"}],"outputs":[{"name":"minimalValidators","type":"address[]","internalType":"address[]"},{"name":"","type":"uint256[]","internalType":"uint256[]"}],"stateMutability":"view"},{"type":"function","name":"getNewOfficialPoolValidBlock","inputs":[],"outputs":[{"name":"","type":"uint256","internalType":"uint256"}],"stateMutability":"view"},{"type":"function","name":"getOldNewOfficialSignerAndValidBlock","inputs":[],"outputs":[{"name":"","type":"address","internalType":"address"},{"name":"","type":"address","internalType":"address"},{"name":"","type":"uint256","internalType":"uint256"}],"stateMutability":"view"},{"type":"function","name":"getValidatorByIndex","inputs":[{"name":"index_","type":"uint256","internalType":"uint256"}],"outputs":[{"name":"","type":"address","internalType":"address"}],"stateMutability":"view"},{"type":"function","name":"getValidatorByPage","inputs":[{"name":"page_","type":"uint256","internalType":"uint256"},{"name":"limit_","type":"uint256","internalType":"uint256"}],"outputs":[{"name":"validatorAddresses","type":"address[]","internalType":"address[]"}],"stateMutability":"view"},{"type":"function","name":"getValidatorCurrentIndex","inputs":[{"name":"val_","type":"address","internalType":"address"},{"name":"revertIfNotFound_","type":"bool","internalType":"bool"}],"outputs":[{"name":"","type":"uint256","internalType":"uint256"},{"name":"","type":"bool","internalType":"bool"}],"stateMutability":"view"},{"type":"function","name":"getValidatorIndexByIndex","inputs":[{"name":"val_","type":"address","internalType":"address"},{"name":"index_","type":"uint256","internalType":"uint256"}],"outputs":[{"name":"","type":"uint256","internalType":"uint256"}],"stateMutability":"view"},{"type":"function","name":"getValidatorIndexLength","inputs":[{"name":"val_","type":"address","internalType":"address"}],"outputs":[{"name":"","type":"uint256","internalType":"uint256"}],"stateMutability":"view"},{"type":"function","name":"getValidatorInfo","inputs":[{"name":"key_","type":"address","internalType":"address"}],"outputs":[{"name":"","type":"tuple","internalType":"struct Validator","components":[{"name":"amount","type":"uint128","internalType":"uint128"},{"name":"delegatedAmount","type":"uint128","internalType":"uint128"},{"name":"reward","type":"uint128","internalType":"uint128"},{"name":"delegatorsReward","type":"uint128","internalType":"uint128"},{"name":"infraCommissionAmount","type":"uint128","internalType":"uint128"},{"name":"validatorCommissionAmount","type":"uint128","internalType":"uint128"},{"name":"delegatorCommissionAmount","type":"uint128","internalType":"uint128"},{"name":"minDeposit","type":"uint128","internalType":"uint128"},{"name":"signer","type":"address","internalType":"address"},{"name":"validatorShareContract","type":"address","internalType":"address"},{"name":"status","type":"uint8","internalType":"enum Status"},{"name":"infraCommissionRate","type":"uint16","internalType":"uint16"},{"name":"commissionRate","type":"uint16","internalType":"uint16"}]}],"stateMutability":"view"},{"type":"function","name":"getValidatorInfoAmount","inputs":[{"name":"key_","type":"address","internalType":"address"}],"outputs":[{"name":"","type":"uint256","internalType":"uint256"}],"stateMutability":"view"},{"type":"function","name":"getValidatorInfoAmountByIndex","inputs":[{"name":"index_","type":"uint256","internalType":"uint256"}],"outputs":[{"name":"","type":"uint256","internalType":"uint256"}],"stateMutability":"view"},{"type":"function","name":"getValidatorInfoByIndex","inputs":[{"name":"index_","type":"uint256","internalType":"uint256"}],"outputs":[{"name":"","type":"tuple","internalType":"struct Validator","components":[{"name":"amount","type":"uint128","internalType":"uint128"},{"name":"delegatedAmount","type":"uint128","internalType":"uint128"},{"name":"reward","type":"uint128","internalType":"uint128"},{"name":"delegatorsReward","type":"uint128","internalType":"uint128"},{"name":"infraCommissionAmount","type":"uint128","internalType":"uint128"},{"name":"validatorCommissionAmount","type":"uint128","internalType":"uint128"},{"name":"delegatorCommissionAmount","type":"uint128","internalType":"uint128"},{"name":"minDeposit","type":"uint128","internalType":"uint128"},{"name":"signer","type":"address","internalType":"address"},{"name":"validatorShareContract","type":"address","internalType":"address"},{"name":"status","type":"uint8","internalType":"enum Status"},{"name":"infraCommissionRate","type":"uint16","internalType":"uint16"},{"name":"commissionRate","type":"uint16","internalType":"uint16"}]}],"stateMutability":"view"},{"type":"function","name":"getValidatorInfoCommissionRate","inputs":[{"name":"key_","type":"address","internalType":"address"}],"outputs":[{"name":"","type":"uint256","internalType":"uint256"}],"stateMutability":"view"},{"type":"function","name":"getValidatorInfoCommissionRateByIndex","inputs":[{"name":"index_","type":"uint256","internalType":"uint256"}],"outputs":[{"name":"","type":"uint256","internalType":"uint256"}],"stateMutability":"view"},{"type":"function","name":"getValidatorInfoDelegatedAmount","inputs":[{"name":"key_","type":"address","internalType":"address"}],"outputs":[{"name":"","type":"uint256","internalType":"uint256"}],"stateMutability":"view"},{"type":"function","name":"getValidatorInfoDelegatedAmountByIndex","inputs":[{"name":"index_","type":"uint256","internalType":"uint256"}],"outputs":[{"name":"","type":"uint256","internalType":"uint256"}],"stateMutability":"view"},{"type":"function","name":"getValidatorInfoDelegatorCommissionAmount","inputs":[{"name":"key_","type":"address","internalType":"address"}],"outputs":[{"name":"","type":"uint256","internalType":"uint256"}],"stateMutability":"view"},{"type":"function","name":"getValidatorInfoDelegatorCommissionAmountByIndex","inputs":[{"name":"index_","type":"uint256","internalType":"uint256"}],"outputs":[{"name":"","type":"uint256","internalType":"uint256"}],"stateMutability":"view"},{"type":"function","name":"getValidatorInfoDelegatorsReward","inputs":[{"name":"key_","type":"address","internalType":"address"}],"outputs":[{"name":"","type":"uint256","internalType":"uint256"}],"stateMutability":"view"},{"type":"function","name":"getValidatorInfoDelegatorsRewardByIndex","inputs":[{"name":"index_","type":"uint256","internalType":"uint256"}],"outputs":[{"name":"","type":"uint256","internalType":"uint256"}],"stateMutability":"view"},{"type":"function","name":"getValidatorInfoInfraCommissionAmount","inputs":[{"name":"key_","type":"address","internalType":"address"}],"outputs":[{"name":"","type":"uint256","internalType":"uint256"}],"stateMutability":"view"},{"type":"function","name":"getValidatorInfoInfraCommissionAmountByIndex","inputs":[{"name":"index_","type":"uint256","internalType":"uint256"}],"outputs":[{"name":"","type":"uint256","internalType":"uint256"}],"stateMutability":"view"},{"type":"function","name":"getValidatorInfoInfraCommissionRate","inputs":[{"name":"key_","type":"address","internalType":"address"}],"outputs":[{"name":"","type":"uint256","internalType":"uint256"}],"stateMutability":"view"},{"type":"function","name":"getValidatorInfoInfraCommissionRateByIndex","inputs":[{"name":"index_","type":"uint256","internalType":"uint256"}],"outputs":[{"name":"","type":"uint256","internalType":"uint256"}],"stateMutability":"view"},{"type":"function","name":"getValidatorInfoMinDeposit","inputs":[{"name":"key_","type":"address","internalType":"address"}],"outputs":[{"name":"","type":"uint256","internalType":"uint256"}],"stateMutability":"view"},{"type":"function","name":"getValidatorInfoMinDepositByIndex","inputs":[{"name":"index_","type":"uint256","internalType":"uint256"}],"outputs":[{"name":"","type":"uint256","internalType":"uint256"}],"stateMutability":"view"},{"type":"function","name":"getValidatorInfoReward","inputs":[{"name":"key_","type":"address","internalType":"address"}],"outputs":[{"name":"","type":"uint256","internalType":"uint256"}],"stateMutability":"view"},{"type":"function","name":"getValidatorInfoRewardByIndex","inputs":[{"name":"index_","type":"uint256","internalType":"uint256"}],"outputs":[{"name":"","type":"uint256","internalType":"uint256"}],"stateMutability":"view"},{"type":"function","name":"getValidatorInfoSigner","inputs":[{"name":"key_","type":"address","internalType":"address"}],"outputs":[{"name":"","type":"address","internalType":"address"}],"stateMutability":"view"},{"type":"function","name":"getValidatorInfoSignerByIndex","inputs":[{"name":"index_","type":"uint256","internalType":"uint256"}],"outputs":[{"name":"","type":"address","internalType":"address"}],"stateMutability":"view"},{"type":"function","name":"getValidatorInfoStatus","inputs":[{"name":"key_","type":"address","internalType":"address"}],"outputs":[{"name":"","type":"uint8","internalType":"enum Status"}],"stateMutability":"view"},{"type":"function","name":"getValidatorInfoStatusByIndex","inputs":[{"name":"index_","type":"uint256","internalType":"uint256"}],"outputs":[{"name":"","type":"uint8","internalType":"enum Status"}],"stateMutability":"view"},{"type":"function","name":"getValidatorInfoValidatorCommissionAmount","inputs":[{"name":"key_","type":"address","internalType":"address"}],"outputs":[{"name":"","type":"uint256","internalType":"uint256"}],"stateMutability":"view"},{"type":"function","name":"getValidatorInfoValidatorCommissionAmountByIndex","inputs":[{"name":"index_","type":"uint256","internalType":"uint256"}],"outputs":[{"name":"","type":"uint256","internalType":"uint256"}],"stateMutability":"view"},{"type":"function","name":"getValidatorInfoValidatorShareContract","inputs":[{"name":"key_","type":"address","internalType":"address"}],"outputs":[{"name":"","type":"address","internalType":"address"}],"stateMutability":"view"},{"type":"function","name":"getValidatorInfoValidatorShareContractByIndex","inputs":[{"name":"index_","type":"uint256","internalType":"uint256"}],"outputs":[{"name":"","type":"address","internalType":"address"}],"stateMutability":"view"},{"type":"function","name":"getValidatorListLength","inputs":[],"outputs":[{"name":"","type":"uint256","internalType":"uint256"}],"stateMutability":"view"},{"type":"function","name":"incDescDefaultInfraCommissionRate","inputs":[{"name":"val_","type":"uint256","internalType":"uint256"},{"name":"mode_","type":"uint8","internalType":"enum ChangeMode"}],"outputs":[],"stateMutability":"nonpayable"},{"type":"function","name":"incDescOfficialAmount","inputs":[{"name":"val_","type":"uint256","internalType":"uint256"},{"name":"mode_","type":"uint8","internalType":"enum ChangeMode"}],"outputs":[],"stateMutability":"nonpayable"},{"type":"function","name":"incDescOfficialLimit","inputs":[{"name":"val_","type":"uint256","internalType":"uint256"},{"name":"mode_","type":"uint8","internalType":"enum ChangeMode"}],"outputs":[],"stateMutability":"nonpayable"},{"type":"function","name":"incDescPoolAmount","inputs":[{"name":"val_","type":"uint256","internalType":"uint256"},{"name":"mode_","type":"uint8","internalType":"enum ChangeMode"}],"outputs":[],"stateMutability":"nonpayable"},{"type":"function","name":"incDescPoolLimit","inputs":[{"name":"val_","type":"uint256","internalType":"uint256"},{"name":"mode_","type":"uint8","internalType":"enum ChangeMode"}],"outputs":[],"stateMutability":"nonpayable"},{"type":"function","name":"incDescPoolSlashAmount","inputs":[{"name":"val_","type":"uint256","internalType":"uint256"},{"name":"mode_","type":"uint8","internalType":"enum ChangeMode"}],"outputs":[],"stateMutability":"nonpayable"},{"type":"function","name":"incDescSoloAmount","inputs":[{"name":"val_","type":"uint256","internalType":"uint256"},{"name":"mode_","type":"uint8","internalType":"enum ChangeMode"}],"outputs":[],"stateMutability":"nonpayable"},{"type":"function","name":"incDescSoloLimit","inputs":[{"name":"val_","type":"uint256","internalType":"uint256"},{"name":"mode_","type":"uint8","internalType":"enum ChangeMode"}],"outputs":[],"stateMutability":"nonpayable"},{"type":"function","name":"incDescTotalRewards","inputs":[{"name":"val_","type":"uint256","internalType":"uint256"},{"name":"mode_","type":"uint8","internalType":"enum ChangeMode"}],"outputs":[],"stateMutability":"nonpayable"},{"type":"function","name":"incDescTotalRewardsLiquidated","inputs":[{"name":"val_","type":"uint256","internalType":"uint256"},{"name":"mode_","type":"uint8","internalType":"enum ChangeMode"}],"outputs":[],"stateMutability":"nonpayable"},{"type":"function","name":"incDescTotalStaked","inputs":[{"name":"val_","type":"uint256","internalType":"uint256"},{"name":"mode_","type":"uint8","internalType":"enum ChangeMode"}],"outputs":[],"stateMutability":"nonpayable"},{"type":"function","name":"incDescUnallocatedReward","inputs":[{"name":"val_","type":"uint256","internalType":"uint256"},{"name":"mode_","type":"uint8","internalType":"enum ChangeMode"}],"outputs":[],"stateMutability":"nonpayable"},{"type":"function","name":"incDescValidatorInfoAmount","inputs":[{"name":"key_","type":"address","internalType":"address"},{"name":"val_","type":"uint256","internalType":"uint256"},{"name":"mode_","type":"uint8","internalType":"enum ChangeMode"}],"outputs":[],"stateMutability":"nonpayable"},{"type":"function","name":"incDescValidatorInfoAmountByIndex","inputs":[{"name":"index_","type":"uint256","internalType":"uint256"},{"name":"val_","type":"uint256","internalType":"uint256"},{"name":"mode_","type":"uint8","internalType":"enum ChangeMode"}],"outputs":[],"stateMutability":"nonpayable"},{"type":"function","name":"incDescValidatorInfoCommissionRate","inputs":[{"name":"key_","type":"address","internalType":"address"},{"name":"val_","type":"uint256","internalType":"uint256"},{"name":"mode_","type":"uint8","internalType":"enum ChangeMode"}],"outputs":[],"stateMutability":"nonpayable"},{"type":"function","name":"incDescValidatorInfoCommissionRateByIndex","inputs":[{"name":"index_","type":"uint256","internalType":"uint256"},{"name":"val_","type":"uint256","internalType":"uint256"},{"name":"mode_","type":"uint8","internalType":"enum ChangeMode"}],"outputs":[],"stateMutability":"nonpayable"},{"type":"function","name":"incDescValidatorInfoDelegatedAmount","inputs":[{"name":"key_","type":"address","internalType":"address"},{"name":"val_","type":"uint256","internalType":"uint256"},{"name":"mode_","type":"uint8","internalType":"enum ChangeMode"}],"outputs":[],"stateMutability":"nonpayable"},{"type":"function","name":"incDescValidatorInfoDelegatedAmountByIndex","inputs":[{"name":"index_","type":"uint256","internalType":"uint256"},{"name":"val_","type":"uint256","internalType":"uint256"},{"name":"mode_","type":"uint8","internalType":"enum ChangeMode"}],"outputs":[],"stateMutability":"nonpayable"},{"type":"function","name":"incDescValidatorInfoDelegatorCommissionAmount","inputs":[{"name":"key_","type":"address","internalType":"address"},{"name":"val_","type":"uint256","internalType":"uint256"},{"name":"mode_","type":"uint8","internalType":"enum ChangeMode"}],"outputs":[],"stateMutability":"nonpayable"},{"type":"function","name":"incDescValidatorInfoDelegatorCommissionAmountByIndex","inputs":[{"name":"index_","type":"uint256","internalType":"uint256"},{"name":"val_","type":"uint256","internalType":"uint256"},{"name":"mode_","type":"uint8","internalType":"enum ChangeMode"}],"outputs":[],"stateMutability":"nonpayable"},{"type":"function","name":"incDescValidatorInfoDelegatorsReward","inputs":[{"name":"key_","type":"address","internalType":"address"},{"name":"val_","type":"uint256","internalType":"uint256"},{"name":"mode_","type":"uint8","internalType":"enum ChangeMode"}],"outputs":[],"stateMutability":"nonpayable"},{"type":"function","name":"incDescValidatorInfoDelegatorsRewardByIndex","inputs":[{"name":"index_","type":"uint256","internalType":"uint256"},{"name":"val_","type":"uint256","internalType":"uint256"},{"name":"mode_","type":"uint8","internalType":"enum ChangeMode"}],"outputs":[],"stateMutability":"nonpayable"},{"type":"function","name":"incDescValidatorInfoInfraCommissionAmount","inputs":[{"name":"key_","type":"address","internalType":"address"},{"name":"val_","type":"uint256","internalType":"uint256"},{"name":"mode_","type":"uint8","internalType":"enum ChangeMode"}],"outputs":[],"stateMutability":"nonpayable"},{"type":"function","name":"incDescValidatorInfoInfraCommissionAmountByIndex","inputs":[{"name":"index_","type":"uint256","internalType":"uint256"},{"name":"val_","type":"uint256","internalType":"uint256"},{"name":"mode_","type":"uint8","internalType":"enum ChangeMode"}],"outputs":[],"stateMutability":"nonpayable"},{"type":"function","name":"incDescValidatorInfoInfraCommissionRate","inputs":[{"name":"key_","type":"address","internalType":"address"},{"name":"val_","type":"uint256","internalType":"uint256"},{"name":"mode_","type":"uint8","internalType":"enum ChangeMode"}],"outputs":[],"stateMutability":"nonpayable"},{"type":"function","name":"incDescValidatorInfoInfraCommissionRateByIndex","inputs":[{"name":"index_","type":"uint256","internalType":"uint256"},{"name":"val_","type":"uint256","internalType":"uint256"},{"name":"mode_","type":"uint8","internalType":"enum ChangeMode"}],"outputs":[],"stateMutability":"nonpayable"},{"type":"function","name":"incDescValidatorInfoMinDeposit","inputs":[{"name":"key_","type":"address","internalType":"address"},{"name":"val_","type":"uint256","internalType":"uint256"},{"name":"mode_","type":"uint8","internalType":"enum ChangeMode"}],"outputs":[],"stateMutability":"nonpayable"},{"type":"function","name":"incDescValidatorInfoMinDepositByIndex","inputs":[{"name":"index_","type":"uint256","internalType":"uint256"},{"name":"val_","type":"uint256","internalType":"uint256"},{"name":"mode_","type":"uint8","internalType":"enum ChangeMode"}],"outputs":[],"stateMutability":"nonpayable"},{"type":"function","name":"incDescValidatorInfoReward","inputs":[{"name":"key_","type":"address","internalType":"address"},{"name":"val_","type":"uint256","internalType":"uint256"},{"name":"mode_","type":"uint8","internalType":"enum ChangeMode"}],"outputs":[],"stateMutability":"nonpayable"},{"type":"function","name":"incDescValidatorInfoRewardByIndex","inputs":[{"name":"index_","type":"uint256","internalType":"uint256"},{"name":"val_","type":"uint256","internalType":"uint256"},{"name":"mode_","type":"uint8","internalType":"enum ChangeMode"}],"outputs":[],"stateMutability":"nonpayable"},{"type":"function","name":"incDescValidatorInfoValidatorCommissionAmount","inputs":[{"name":"key_","type":"address","internalType":"address"},{"name":"val_","type":"uint256","internalType":"uint256"},{"name":"mode_","type":"uint8","internalType":"enum ChangeMode"}],"outputs":[],"stateMutability":"nonpayable"},{"type":"function","name":"incDescValidatorInfoValidatorCommissionAmountByIndex","inputs":[{"name":"index_","type":"uint256","internalType":"uint256"},{"name":"val_","type":"uint256","internalType":"uint256"},{"name":"mode_","type":"uint8","internalType":"enum ChangeMode"}],"outputs":[],"stateMutability":"nonpayable"},{"type":"function","name":"initialize1","inputs":[{"name":"nftContract_","type":"address","internalType":"address"},{"name":"validatorShareFactory_","type":"address","internalType":"address"},{"name":"slashManager_","type":"address","internalType":"address"}],"outputs":[],"stateMutability":"nonpayable"},{"type":"function","name":"initialize2","inputs":[{"name":"stakeManager_","type":"address","internalType":"address"}],"outputs":[],"stateMutability":"nonpayable"},{"type":"function","name":"initialized","inputs":[{"name":"","type":"uint256","internalType":"uint256"}],"outputs":[{"name":"","type":"bool","internalType":"bool"}],"stateMutability":"view"},{"type":"function","name":"isInMinimalValidatorList","inputs":[{"name":"val_","type":"address","internalType":"address"}],"outputs":[{"name":"","type":"bool","internalType":"bool"}],"stateMutability":"view"},{"type":"function","name":"isInMinimalValidatorListByValidatorId","inputs":[{"name":"validatorId_","type":"uint256","internalType":"uint256"}],"outputs":[{"name":"","type":"bool","internalType":"bool"}],"stateMutability":"view"},{"type":"function","name":"isInValidatorList","inputs":[{"name":"val_","type":"address","internalType":"address"}],"outputs":[{"name":"","type":"bool","internalType":"bool"}],"stateMutability":"view"},{"type":"function","name":"isValidatorValid","inputs":[{"name":"_signer","type":"address","internalType":"address"}],"outputs":[{"name":"","type":"bool","internalType":"bool"}],"stateMutability":"view"},{"type":"function","name":"minimumPoolDelegate","inputs":[],"outputs":[{"name":"","type":"uint256","internalType":"uint256"}],"stateMutability":"view"},{"type":"function","name":"minimumPoolStake","inputs":[],"outputs":[{"name":"","type":"uint256","internalType":"uint256"}],"stateMutability":"view"},{"type":"function","name":"minimumSoloStake","inputs":[],"outputs":[{"name":"","type":"uint256","internalType":"uint256"}],"stateMutability":"view"},{"type":"function","name":"nftContract","inputs":[],"outputs":[{"name":"","type":"address","internalType":"address"}],"stateMutability":"view"},{"type":"function","name":"officialAmount","inputs":[],"outputs":[{"name":"","type":"uint256","internalType":"uint256"}],"stateMutability":"view"},{"type":"function","name":"officialLimit","inputs":[],"outputs":[{"name":"","type":"uint256","internalType":"uint256"}],"stateMutability":"view"},{"type":"function","name":"officialPool","inputs":[],"outputs":[{"name":"","type":"address","internalType":"address"}],"stateMutability":"view"},{"type":"function","name":"poolAmount","inputs":[],"outputs":[{"name":"","type":"uint256","internalType":"uint256"}],"stateMutability":"view"},{"type":"function","name":"poolLimit","inputs":[],"outputs":[{"name":"","type":"uint256","internalType":"uint256"}],"stateMutability":"view"},{"type":"function","name":"poolSlashAmount","inputs":[],"outputs":[{"name":"","type":"uint256","internalType":"uint256"}],"stateMutability":"view"},{"type":"function","name":"removeMinimalValidator","inputs":[{"name":"val_","type":"address","internalType":"address"}],"outputs":[],"stateMutability":"nonpayable"},{"type":"function","name":"removeMinimalValidatorByIndex","inputs":[{"name":"index_","type":"uint256","internalType":"uint256"}],"outputs":[],"stateMutability":"nonpayable"},{"type":"function","name":"setCommittee","inputs":[{"name":"_committee","type":"address","internalType":"address"}],"outputs":[],"stateMutability":"nonpayable"},{"type":"function","name":"setDefaultInfraCommissionRate","inputs":[{"name":"val_","type":"uint256","internalType":"uint256"}],"outputs":[],"stateMutability":"nonpayable"},{"type":"function","name":"setMinimumPoolDelegate","inputs":[{"name":"_value","type":"uint256","internalType":"uint256"}],"outputs":[],"stateMutability":"nonpayable"},{"type":"function","name":"setMinimumPoolStake","inputs":[{"name":"_value","type":"uint256","internalType":"uint256"}],"outputs":[],"stateMutability":"nonpayable"},{"type":"function","name":"setMinimumSoloStake","inputs":[{"name":"_value","type":"uint256","internalType":"uint256"}],"outputs":[],"stateMutability":"nonpayable"},{"type":"function","name":"setOfficialAmount","inputs":[{"name":"val_","type":"uint256","internalType":"uint256"}],"outputs":[],"stateMutability":"nonpayable"},{"type":"function","name":"setOfficialLimit","inputs":[{"name":"val_","type":"uint256","internalType":"uint256"}],"outputs":[],"stateMutability":"nonpayable"},{"type":"function","name":"setPoolAmount","inputs":[{"name":"val_","type":"uint256","internalType":"uint256"}],"outputs":[],"stateMutability":"nonpayable"},{"type":"function","name":"setPoolLimit","inputs":[{"name":"val_","type":"uint256","internalType":"uint256"}],"outputs":[],"stateMutability":"nonpayable"},{"type":"function","name":"setPoolSlashAmount","inputs":[{"name":"val_","type":"uint256","internalType":"uint256"}],"outputs":[],"stateMutability":"nonpayable"},{"type":"function","name":"setSlashEpochSize","inputs":[{"name":"_size","type":"uint256","internalType":"uint256"}],"outputs":[],"stateMutability":"nonpayable"},{"type":"function","name":"setSlashEpochThreshold","inputs":[{"name":"_threshold","type":"uint256","internalType":"uint256"}],"outputs":[],"stateMutability":"nonpayable"},{"type":"function","name":"setSlashManager","inputs":[{"name":"slashManager_","type":"address","internalType":"address"}],"outputs":[],"stateMutability":"nonpayable"},{"type":"function","name":"setSoloAmount","inputs":[{"name":"val_","type":"uint256","internalType":"uint256"}],"outputs":[],"stateMutability":"nonpayable"},{"type":"function","name":"setSoloLimit","inputs":[{"name":"val_","type":"uint256","internalType":"uint256"}],"outputs":[],"stateMutability":"nonpayable"},{"type":"function","name":"setSoloSlashRate","inputs":[{"name":"val_","type":"uint256","internalType":"uint256"}],"outputs":[],"stateMutability":"nonpayable"},{"type":"function","name":"setStakeManager","inputs":[{"name":"stakeManager_","type":"address","internalType":"address"}],"outputs":[],"stateMutability":"nonpayable"},{"type":"function","name":"setTotalRewards","inputs":[{"name":"val_","type":"uint256","internalType":"uint256"}],"outputs":[],"stateMutability":"nonpayable"},{"type":"function","name":"setTotalRewardsLiquidated","inputs":[{"name":"val_","type":"uint256","internalType":"uint256"}],"outputs":[],"stateMutability":"nonpayable"},{"type":"function","name":"setTotalStaked","inputs":[{"name":"val_","type":"uint256","internalType":"uint256"}],"outputs":[],"stateMutability":"nonpayable"},{"type":"function","name":"setUnallocatedReward","inputs":[{"name":"val_","type":"uint256","internalType":"uint256"}],"outputs":[],"stateMutability":"nonpayable"},{"type":"function","name":"setValidatorIds","inputs":[{"name":"key_","type":"address","internalType":"address"},{"name":"validatorIds_","type":"uint256[]","internalType":"uint256[]"}],"outputs":[],"stateMutability":"nonpayable"},{"type":"function","name":"setValidatorInfo","inputs":[{"name":"key_","type":"address","internalType":"address"},{"name":"val_","type":"tuple","internalType":"struct Validator","components":[{"name":"amount","type":"uint128","internalType":"uint128"},{"name":"delegatedAmount","type":"uint128","internalType":"uint128"},{"name":"reward","type":"uint128","internalType":"uint128"},{"name":"delegatorsReward","type":"uint128","internalType":"uint128"},{"name":"infraCommissionAmount","type":"uint128","internalType":"uint128"},{"name":"validatorCommissionAmount","type":"uint128","internalType":"uint128"},{"name":"delegatorCommissionAmount","type":"uint128","internalType":"uint128"},{"name":"minDeposit","type":"uint128","internalType":"uint128"},{"name":"signer","type":"address","internalType":"address"},{"name":"validatorShareContract","type":"address","internalType":"address"},{"name":"status","type":"uint8","internalType":"enum Status"},{"name":"infraCommissionRate","type":"uint16","internalType":"uint16"},{"name":"commissionRate","type":"uint16","internalType":"uint16"}]}],"outputs":[],"stateMutability":"nonpayable"},{"type":"function","name":"setValidatorInfoAmount","inputs":[{"name":"key_","type":"address","internalType":"address"},{"name":"val_","type":"uint256","internalType":"uint256"}],"outputs":[],"stateMutability":"nonpayable"},{"type":"function","name":"setValidatorInfoAmountByIndex","inputs":[{"name":"index_","type":"uint256","internalType":"uint256"},{"name":"val_","type":"uint256","internalType":"uint256"}],"outputs":[],"stateMutability":"nonpayable"},{"type":"function","name":"setValidatorInfoByIndex","inputs":[{"name":"index_","type":"uint256","internalType":"uint256"},{"name":"val_","type":"tuple","internalType":"struct Validator","components":[{"name":"amount","type":"uint128","internalType":"uint128"},{"name":"delegatedAmount","type":"uint128","internalType":"uint128"},{"name":"reward","type":"uint128","internalType":"uint128"},{"name":"delegatorsReward","type":"uint128","internalType":"uint128"},{"name":"infraCommissionAmount","type":"uint128","internalType":"uint128"},{"name":"validatorCommissionAmount","type":"uint128","internalType":"uint128"},{"name":"delegatorCommissionAmount","type":"uint128","internalType":"uint128"},{"name":"minDeposit","type":"uint128","internalType":"uint128"},{"name":"signer","type":"address","internalType":"address"},{"name":"validatorShareContract","type":"address","internalType":"address"},{"name":"status","type":"uint8","internalType":"enum Status"},{"name":"infraCommissionRate","type":"uint16","internalType":"uint16"},{"name":"commissionRate","type":"uint16","internalType":"uint16"}]}],"outputs":[],"stateMutability":"nonpayable"},{"type":"function","name":"setValidatorInfoCommissionRate","inputs":[{"name":"key_","type":"address","internalType":"address"},{"name":"val_","type":"uint256","internalType":"uint256"}],"outputs":[],"stateMutability":"nonpayable"},{"type":"function","name":"setValidatorInfoCommissionRateByIndex","inputs":[{"name":"index_","type":"uint256","internalType":"uint256"},{"name":"val_","type":"uint256","internalType":"uint256"}],"outputs":[],"stateMutability":"nonpayable"},{"type":"function","name":"setValidatorInfoDelegatedAmount","inputs":[{"name":"key_","type":"address","internalType":"address"},{"name":"val_","type":"uint256","internalType":"uint256"}],"outputs":[],"stateMutability":"nonpayable"},{"type":"function","name":"setValidatorInfoDelegatedAmountByIndex","inputs":[{"name":"index_","type":"uint256","internalType":"uint256"},{"name":"val_","type":"uint256","internalType":"uint256"}],"outputs":[],"stateMutability":"nonpayable"},{"type":"function","name":"setValidatorInfoDelegatorCommissionAmount","inputs":[{"name":"key_","type":"address","internalType":"address"},{"name":"val_","type":"uint256","internalType":"uint256"}],"outputs":[],"stateMutability":"nonpayable"},{"type":"function","name":"setValidatorInfoDelegatorCommissionAmountByIndex","inputs":[{"name":"index_","type":"uint256","internalType":"uint256"},{"name":"val_","type":"uint256","internalType":"uint256"}],"outputs":[],"stateMutability":"nonpayable"},{"type":"function","name":"setValidatorInfoDelegatorsReward","inputs":[{"name":"key_","type":"address","internalType":"address"},{"name":"val_","type":"uint256","internalType":"uint256"}],"outputs":[],"stateMutability":"nonpayable"},{"type":"function","name":"setValidatorInfoDelegatorsRewardByIndex","inputs":[{"name":"index_","type":"uint256","internalType":"uint256"},{"name":"val_","type":"uint256","internalType":"uint256"}],"outputs":[],"stateMutability":"nonpayable"},{"type":"function","name":"setValidatorInfoInfraCommissionAmount","inputs":[{"name":"key_","type":"address","internalType":"address"},{"name":"val_","type":"uint256","internalType":"uint256"}],"outputs":[],"stateMutability":"nonpayable"},{"type":"function","name":"setValidatorInfoInfraCommissionAmountByIndex","inputs":[{"name":"index_","type":"uint256","internalType":"uint256"},{"name":"val_","type":"uint256","internalType":"uint256"}],"outputs":[],"stateMutability":"nonpayable"},{"type":"function","name":"setValidatorInfoInfraCommissionRate","inputs":[{"name":"key_","type":"address","internalType":"address"},{"name":"val_","type":"uint256","internalType":"uint256"}],"outputs":[],"stateMutability":"nonpayable"},{"type":"function","name":"setValidatorInfoInfraCommissionRateByIndex","inputs":[{"name":"index_","type":"uint256","internalType":"uint256"},{"name":"val_","type":"uint256","internalType":"uint256"}],"outputs":[],"stateMutability":"nonpayable"},{"type":"function","name":"setValidatorInfoMinDeposit","inputs":[{"name":"key_","type":"address","internalType":"address"},{"name":"val_","type":"uint256","internalType":"uint256"}],"outputs":[],"stateMutability":"nonpayable"},{"type":"function","name":"setValidatorInfoMinDepositByIndex","inputs":[{"name":"index_","type":"uint256","internalType":"uint256"},{"name":"val_","type":"uint256","internalType":"uint256"}],"outputs":[],"stateMutability":"nonpayable"},{"type":"function","name":"setValidatorInfoReward","inputs":[{"name":"key_","type":"address","internalType":"address"},{"name":"val_","type":"uint256","internalType":"uint256"}],"outputs":[],"stateMutability":"nonpayable"},{"type":"function","name":"setValidatorInfoRewardByIndex","inputs":[{"name":"index_","type":"uint256","internalType":"uint256"},{"name":"val_","type":"uint256","internalType":"uint256"}],"outputs":[],"stateMutability":"nonpayable"},{"type":"function","name":"setValidatorInfoSigner","inputs":[{"name":"key_","type":"address","internalType":"address"},{"name":"val_","type":"address","internalType":"address"}],"outputs":[],"stateMutability":"nonpayable"},{"type":"function","name":"setValidatorInfoSignerByIndex","inputs":[{"name":"index_","type":"uint256","internalType":"uint256"},{"name":"val_","type":"address","internalType":"address"}],"outputs":[],"stateMutability":"nonpayable"},{"type":"function","name":"setValidatorInfoStatus","inputs":[{"name":"key_","type":"address","internalType":"address"},{"name":"val_","type":"uint8","internalType":"enum Status"}],"outputs":[],"stateMutability":"nonpayable"},{"type":"function","name":"setValidatorInfoStatusByIndex","inputs":[{"name":"index_","type":"uint256","internalType":"uint256"},{"name":"val_","type":"uint8","internalType":"enum Status"}],"outputs":[],"stateMutability":"nonpayable"},{"type":"function","name":"setValidatorInfoValidatorCommissionAmount","inputs":[{"name":"key_","type":"address","internalType":"address"},{"name":"val_","type":"uint256","internalType":"uint256"}],"outputs":[],"stateMutability":"nonpayable"},{"type":"function","name":"setValidatorInfoValidatorCommissionAmountByIndex","inputs":[{"name":"index_","type":"uint256","internalType":"uint256"},{"name":"val_","type":"uint256","internalType":"uint256"}],"outputs":[],"stateMutability":"nonpayable"},{"type":"function","name":"setValidatorInfoValidatorShareContract","inputs":[{"name":"key_","type":"address","internalType":"address"},{"name":"val_","type":"address","internalType":"address"}],"outputs":[],"stateMutability":"nonpayable"},{"type":"function","name":"setValidatorInfoValidatorShareContractByIndex","inputs":[{"name":"index_","type":"uint256","internalType":"uint256"},{"name":"val_","type":"address","internalType":"address"}],"outputs":[],"stateMutability":"nonpayable"},{"type":"function","name":"setValidatorShareFactory","inputs":[{"name":"validatorShareFactory_","type":"address","internalType":"address"}],"outputs":[],"stateMutability":"nonpayable"},{"type":"function","name":"slashEpochSize","inputs":[],"outputs":[{"name":"","type":"uint256","internalType":"uint256"}],"stateMutability":"view"},{"type":"function","name":"slashEpochThreshold","inputs":[],"outputs":[{"name":"","type":"uint256","internalType":"uint256"}],"stateMutability":"view"},{"type":"function","name":"slashManager","inputs":[],"outputs":[{"name":"","type":"address","internalType":"address"}],"stateMutability":"view"},{"type":"function","name":"soloAmount","inputs":[],"outputs":[{"name":"","type":"uint256","internalType":"uint256"}],"stateMutability":"view"},{"type":"function","name":"soloLimit","inputs":[],"outputs":[{"name":"","type":"uint256","internalType":"uint256"}],"stateMutability":"view"},{"type":"function","name":"soloSlashRate","inputs":[],"outputs":[{"name":"","type":"uint256","internalType":"uint256"}],"stateMutability":"view"},{"type":"function","name":"stakeManager","inputs":[],"outputs":[{"name":"","type":"address","internalType":"address"}],"stateMutability":"view"},{"type":"function","name":"totalRewards","inputs":[],"outputs":[{"name":"","type":"uint256","internalType":"uint256"}],"stateMutability":"view"},{"type":"function","name":"totalRewardsLiquidated","inputs":[],"outputs":[{"name":"","type":"uint256","internalType":"uint256"}],"stateMutability":"view"},{"type":"function","name":"totalStaked","inputs":[],"outputs":[{"name":"","type":"uint256","internalType":"uint256"}],"stateMutability":"view"},{"type":"function","name":"unallocatedReward","inputs":[],"outputs":[{"name":"","type":"uint256","internalType":"uint256"}],"stateMutability":"view"},{"type":"function","name":"validatorShareFactory","inputs":[],"outputs":[{"name":"","type":"address","internalType":"address"}],"stateMutability":"view"},{"type":"event","name":"CommitteeSet","inputs":[{"name":"oldCommittee","type":"address","indexed":true,"internalType":"address"},{"name":"newCommittee","type":"address","indexed":true,"internalType":"address"},{"name":"caller","type":"address","indexed":true,"internalType":"address"}],"anonymous":false},{"type":"event","name":"MinimalValidatorAdded","inputs":[{"name":"val","type":"address","indexed":true,"internalType":"address"},{"name":"length","type":"uint256","indexed":false,"internalType":"uint256"}],"anonymous":false},{"type":"event","name":"MinimalValidatorRemoved","inputs":[{"name":"val","type":"address","indexed":true,"internalType":"address"},{"name":"length","type":"uint256","indexed":false,"internalType":"uint256"}],"anonymous":false},{"type":"event","name":"OfficialSignerChanged","inputs":[{"name":"oldOfficialSigner","type":"address","indexed":true,"internalType":"address"},{"name":"newOfficialSigner","type":"address","indexed":true,"internalType":"address"},{"name":"currentValidatorId","type":"uint256","indexed":true,"internalType":"uint256"},{"name":"newValidBlock","type":"uint256","indexed":false,"internalType":"uint256"},{"name":"caller","type":"address","indexed":false,"internalType":"address"}],"anonymous":false},{"type":"event","name":"SlashManagerSet","inputs":[{"name":"oldSlashManager","type":"address","indexed":true,"internalType":"address"},{"name":"newSlashManager","type":"address","indexed":true,"internalType":"address"},{"name":"caller","type":"address","indexed":true,"internalType":"address"}],"anonymous":false},{"type":"event","name":"StakeManagerSet","inputs":[{"name":"oldStakeManager","type":"address","indexed":true,"internalType":"address"},{"name":"newStakeManager","type":"address","indexed":true,"internalType":"address"},{"name":"caller","type":"address","indexed":true,"internalType":"address"}],"anonymous":false},{"type":"event","name":"ValidatorAdded","inputs":[{"name":"val","type":"address","indexed":true,"internalType":"address"},{"name":"length","type":"uint256","indexed":false,"internalType":"uint256"}],"anonymous":false},{"type":"event","name":"ValidatorRemoved","inputs":[{"name":"val","type":"address","indexed":true,"internalType":"address"},{"name":"length","type":"uint256","indexed":false,"internalType":"uint256"}],"anonymous":false},{"type":"event","name":"ValidatorShareFactorySet","inputs":[{"name":"oldValidatorShareFactory","type":"address","indexed":true,"internalType":"address"},{"name":"newValidatorShareFactory","type":"address","indexed":true,"internalType":"address"},{"name":"caller","type":"address","indexed":true,"internalType":"address"}],"anonymous":false},{"type":"event","name":"ValueSet","inputs":[{"name":"hash","type":"bytes32","indexed":true,"internalType":"bytes32"},{"name":"oldValue","type":"bytes32","indexed":false,"internalType":"bytes32"},{"name":"newValue","type":"bytes32","indexed":false,"internalType":"bytes32"},{"name":"caller","type":"address","indexed":true,"internalType":"address"}],"anonymous":false}]`