package common

import "math/big"

const (
	RewardMasterPercent        = 40
	RewardVoterPercent         = 50
	RewardFoundationPercent    = 10
	HexSignMethod              = "e341eaa4"
	HexSetSecret               = "34d38600"
	HexSetOpening              = "e11f5ba2"
	EpocBlockSecret            = 800
	EpocBlockOpening           = 850
	EpocBlockRandomize         = 900
	MaxMasternodes             = 150
	LimitPenaltyEpoch          = 4
	BlocksPerYear              = uint64(15768000)
	LimitThresholdNonceInQueue = 10
	DefaultMinGasPrice         = 2500
	MergeSignRange             = 15
	RangeReturnSigner          = 150
	MinimunMinerBlockPerEpoch  = 1
)

var TIP2019Block = big.NewInt(1050000)
var TIPSigning = big.NewInt(3000000)
var TIPRandomize = big.NewInt(3464000)
var IsTestnet bool = false
var StoreRewardFolder string
var RollbackHash Hash
var MinGasPrice = big.NewInt(DefaultMinGasPrice)
var TRC21IssuerSMCTestNet = HexToAddress("0x7081C72c9DC44686C7B7EAB1d338EA137Fa9f0D3")
var TRC21IssuerSMC = HexToAddress("0x8c0faeb5C6bEd2129b8674F262Fd45c4e9468bee")
var TRC21GasPrice = big.NewInt(DefaultMinGasPrice)
