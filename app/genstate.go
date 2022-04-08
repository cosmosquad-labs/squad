package app

import (
	"time"

	sdk "github.com/cosmos/cosmos-sdk/types"
	budgettypes "github.com/tendermint/budget/x/budget/types"
	tmproto "github.com/tendermint/tendermint/proto/tendermint/types"

	utils "github.com/cosmosquad-labs/squad/types"
	claimtypes "github.com/cosmosquad-labs/squad/x/claim/types"
	farmingtypes "github.com/cosmosquad-labs/squad/x/farming/types"
	liquiditytypes "github.com/cosmosquad-labs/squad/x/liquidity/types"
	liquidstakingtypes "github.com/cosmosquad-labs/squad/x/liquidstaking/types"
	minttypes "github.com/cosmosquad-labs/squad/x/mint/types"
)

var (
	FoundationAddress             = ""
	AirdropSourceAddress          = ""
	DevTeamAddress                = ""
	FarmingFeeCollector           = ""
	LiquidityFeeCollectorAddress  = ""
	LiquidityDustCollectorAddress = ""
	InflationFeeCollector         = ""
	EcosystemIncentive            = ""
	EcosystemIncentiveLP          = ""
	EcosystemIncentiveMM          = ""
	EcosystemIncentiveBoost       = ""

	BondDenom       = "stake"
	LiquidBondDenom = "bstake"
	DEXDropSupply   = sdk.NewInt(50_000_000_000_000) // 50mil
	BoostDropSupply = sdk.NewInt(50_000_000_000_000) // 50mil
)

type GenesisStates struct {
	DEXdropSupply   sdk.Coin
	BoostdropSupply sdk.Coin
	BondDenom       string
	LiquidBondDenom string
	GenesisTime     time.Time

	ConsensusParams     *tmproto.ConsensusParams
	MintParams          minttypes.Params
	LiquidityParams     liquiditytypes.Params
	LiquidStakingParams liquidstakingtypes.Params
	FarmingParams       farmingtypes.Params
	BudgetParams        budgettypes.Params
	ClaimGenesisState   claimtypes.GenesisState
}

func NewGenesisState() *GenesisStates {
	genParams := &GenesisStates{}
	genParams.BondDenom = BondDenom
	genParams.LiquidBondDenom = LiquidBondDenom
	genParams.DEXdropSupply = sdk.NewCoin(genParams.BondDenom, DEXDropSupply)
	genParams.BoostdropSupply = sdk.NewCoin(genParams.BondDenom, BoostDropSupply)

	// Set genesis time
	genParams.GenesisTime = utils.ParseTime("2022-04-13T00:00:00Z")

	// Set consensus params
	genParams.ConsensusParams = &tmproto.ConsensusParams{
		Block: tmproto.BlockParams{
			MaxBytes:   10000000,
			MaxGas:     100000000,
			TimeIotaMs: 1000,
		},
		Evidence: tmproto.EvidenceParams{
			MaxAgeNumBlocks: 201600,
			MaxAgeDuration:  1209600000000000,
			MaxBytes:        1000000,
		},
		Validator: tmproto.ValidatorParams{
			PubKeyTypes: []string{"ed25519"},
		},
		Version: tmproto.VersionParams{},
	}

	// Set mint params
	genParams.MintParams = minttypes.Params{
		MintDenom:          genParams.BondDenom,
		BlockTimeThreshold: 10 * time.Second,
		InflationSchedules: []minttypes.InflationSchedule{
			{
				StartTime: genParams.GenesisTime,
				EndTime:   genParams.GenesisTime.AddDate(1, 0, 0),
				Amount:    sdk.NewInt(108_700000_000000),
			},
			{
				StartTime: genParams.GenesisTime.AddDate(1, 0, 0),
				EndTime:   genParams.GenesisTime.AddDate(2, 0, 0),
				Amount:    sdk.NewInt(216_100000_000000),
			},
			{
				StartTime: genParams.GenesisTime.AddDate(2, 0, 0),
				EndTime:   genParams.GenesisTime.AddDate(3, 0, 0),
				Amount:    sdk.NewInt(151_300000_000000),
			},
			{
				StartTime: genParams.GenesisTime.AddDate(3, 0, 0),
				EndTime:   genParams.GenesisTime.AddDate(4, 0, 0),
				Amount:    sdk.NewInt(105_900000_000000),
			},
			{
				StartTime: genParams.GenesisTime.AddDate(4, 0, 0),
				EndTime:   genParams.GenesisTime.AddDate(5, 0, 0),
				Amount:    sdk.NewInt(74_100000_000000),
			},
			{
				StartTime: genParams.GenesisTime.AddDate(5, 0, 0),
				EndTime:   genParams.GenesisTime.AddDate(6, 0, 0),
				Amount:    sdk.NewInt(51_900000_000000),
			},
			{
				StartTime: genParams.GenesisTime.AddDate(6, 0, 0),
				EndTime:   genParams.GenesisTime.AddDate(7, 0, 0),
				Amount:    sdk.NewInt(36_300000_000000),
			},
			{
				StartTime: genParams.GenesisTime.AddDate(7, 0, 0),
				EndTime:   genParams.GenesisTime.AddDate(8, 0, 0),
				Amount:    sdk.NewInt(25_400000_000000),
			},
			{
				StartTime: genParams.GenesisTime.AddDate(8, 0, 0),
				EndTime:   genParams.GenesisTime.AddDate(9, 0, 0),
				Amount:    sdk.NewInt(17_800000_000000),
			},
			{
				StartTime: genParams.GenesisTime.AddDate(9, 0, 0),
				EndTime:   genParams.GenesisTime.AddDate(10, 0, 0),
				Amount:    sdk.NewInt(12_500000_000000),
			},
		},
	}

	// Set farming params
	genParams.FarmingParams = farmingtypes.Params{
		PrivatePlanCreationFee: sdk.NewCoins(sdk.NewInt64Coin(genParams.BondDenom, 100000000)),
		NextEpochDays:          1,
		FarmingFeeCollector:    FarmingFeeCollector,
		DelayedStakingGasFee:   sdk.Gas(100000),
		MaxNumPrivatePlans:     10000,
	}

	// Set liquidstaking params
	genParams.LiquidStakingParams = liquidstakingtypes.Params{
		LiquidBondDenom: genParams.LiquidBondDenom,
		WhitelistedValidators: []liquidstakingtypes.WhitelistedValidator{
			{
				ValidatorAddress: "",
				TargetWeight:     sdk.NewInt(10),
			},
		},
		UnstakeFeeRate:         sdk.MustNewDecFromStr("0.000000000000000000"),
		MinLiquidStakingAmount: sdk.NewInt(1000000),
	}

	// Set liquidity params
	genParams.LiquidityParams = liquiditytypes.Params{
		BatchSize:                1,
		TickPrecision:            3,
		FeeCollectorAddress:      LiquidityFeeCollectorAddress,
		DustCollectorAddress:     LiquidityDustCollectorAddress,
		MinInitialPoolCoinSupply: sdk.NewInt(1_000000_000000),
		PairCreationFee:          sdk.NewCoins(sdk.NewInt64Coin(genParams.BondDenom, 100_000_000)),
		PoolCreationFee:          sdk.NewCoins(sdk.NewInt64Coin(genParams.BondDenom, 100_000_000)),
		MinInitialDepositAmount:  sdk.NewInt(100000000),
		DepositExtraGas:          sdk.Gas(60000),
		WithdrawExtraGas:         sdk.Gas(64000),
		OrderExtraGas:            sdk.Gas(37000),
		MaxPriceLimitRatio:       sdk.MustNewDecFromStr("0.100000000000000000"),
		MaxOrderLifespan:         86400 * time.Second,
		SwapFeeRate:              sdk.MustNewDecFromStr("0.000000000000000000"),
		WithdrawFeeRate:          sdk.MustNewDecFromStr("0.000000000000000000"),
	}

	// Set budget params
	genParams.BudgetParams = budgettypes.Params{
		EpochBlocks: 1,
		Budgets: []budgettypes.Budget{
			{
				Name:               "budget-ecosystem-incentive",
				Rate:               sdk.MustNewDecFromStr("0.662500000000000000"),
				SourceAddress:      InflationFeeCollector,
				DestinationAddress: EcosystemIncentive,
				StartTime:          genParams.GenesisTime,
				EndTime:            genParams.GenesisTime.AddDate(10, 0, 0),
			},
			{
				Name:               "budget-dev-team",
				Rate:               sdk.MustNewDecFromStr("0.250000000000000000"),
				SourceAddress:      InflationFeeCollector,
				DestinationAddress: DevTeamAddress,
				StartTime:          genParams.GenesisTime,
				EndTime:            genParams.GenesisTime.AddDate(10, 0, 0),
			},
			{
				Name:               "budget-ecosystem-incentive-lp-1",
				Rate:               sdk.MustNewDecFromStr("0.600000000000000000"),
				SourceAddress:      EcosystemIncentive,
				DestinationAddress: EcosystemIncentiveLP,
				StartTime:          genParams.GenesisTime,
				EndTime:            genParams.GenesisTime.AddDate(1, 0, 0),
			},
			{
				Name:               "budget-ecosystem-incentive-mm-1",
				Rate:               sdk.MustNewDecFromStr("0.200000000000000000"),
				SourceAddress:      EcosystemIncentive,
				DestinationAddress: EcosystemIncentiveMM,
				StartTime:          genParams.GenesisTime,
				EndTime:            genParams.GenesisTime.AddDate(1, 0, 0),
			},
			{
				Name:               "budget-ecosystem-incentive-boost-1",
				Rate:               sdk.MustNewDecFromStr("0.200000000000000000"),
				SourceAddress:      EcosystemIncentive,
				DestinationAddress: EcosystemIncentiveBoost,
				StartTime:          genParams.GenesisTime,
				EndTime:            genParams.GenesisTime.AddDate(1, 0, 0),
			},

			{
				Name:               "budget-ecosystem-incentive-lp-2",
				Rate:               sdk.MustNewDecFromStr("0.200000000000000000"),
				SourceAddress:      EcosystemIncentive,
				DestinationAddress: EcosystemIncentiveLP,
				StartTime:          genParams.GenesisTime.AddDate(1, 0, 0),
				EndTime:            genParams.GenesisTime.AddDate(2, 0, 0),
			},
			{
				Name:               "budget-ecosystem-incentive-mm-2",
				Rate:               sdk.MustNewDecFromStr("0.300000000000000000"),
				SourceAddress:      EcosystemIncentive,
				DestinationAddress: EcosystemIncentiveMM,
				StartTime:          genParams.GenesisTime.AddDate(1, 0, 0),
				EndTime:            genParams.GenesisTime.AddDate(2, 0, 0),
			},
			{
				Name:               "budget-ecosystem-incentive-boost-2",
				Rate:               sdk.MustNewDecFromStr("0.500000000000000000"),
				SourceAddress:      EcosystemIncentive,
				DestinationAddress: EcosystemIncentiveBoost,
				StartTime:          genParams.GenesisTime.AddDate(1, 0, 0),
				EndTime:            genParams.GenesisTime.AddDate(2, 0, 0),
			},
			{
				Name:               "budget-ecosystem-incentive-lp-3-10",
				Rate:               sdk.MustNewDecFromStr("0.100000000000000000"),
				SourceAddress:      EcosystemIncentive,
				DestinationAddress: EcosystemIncentiveLP,
				StartTime:          genParams.GenesisTime.AddDate(2, 0, 0),
				EndTime:            genParams.GenesisTime.AddDate(10, 0, 0),
			},
			{
				Name:               "budget-ecosystem-incentive-mm-3-10",
				Rate:               sdk.MustNewDecFromStr("0.300000000000000000"),
				SourceAddress:      EcosystemIncentive,
				DestinationAddress: EcosystemIncentiveMM,
				StartTime:          genParams.GenesisTime.AddDate(2, 0, 0),
				EndTime:            genParams.GenesisTime.AddDate(10, 0, 0),
			},
			{
				Name:               "budget-ecosystem-incentive-boost-3-10",
				Rate:               sdk.MustNewDecFromStr("0.600000000000000000"),
				SourceAddress:      EcosystemIncentive,
				DestinationAddress: EcosystemIncentiveBoost,
				StartTime:          genParams.GenesisTime.AddDate(2, 0, 0),
				EndTime:            genParams.GenesisTime.AddDate(10, 0, 0),
			},
		},
	}

	// Set claim genesis states
	genParams.ClaimGenesisState.Airdrops = []claimtypes.Airdrop{
		{
			Id:            1,
			SourceAddress: AirdropSourceAddress, // airdrop source address
			Conditions: []claimtypes.ConditionType{
				claimtypes.ConditionTypeDeposit,
				claimtypes.ConditionTypeSwap,
				claimtypes.ConditionTypeLiquidStake,
				claimtypes.ConditionTypeVote,
			},
			StartTime: genParams.GenesisTime,
			EndTime:   genParams.GenesisTime.AddDate(0, 6, 0),
		},
	}
	genParams.ClaimGenesisState.ClaimRecords = []claimtypes.ClaimRecord{
		{
			AirdropId:             1,
			Recipient:             "",
			InitialClaimableCoins: sdk.NewCoins(),
			ClaimableCoins:        sdk.NewCoins(),
			ClaimedConditions:     []claimtypes.ConditionType{},
		},
	}

	return genParams
}
