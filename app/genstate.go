package app

import (
	"time"

	sdk "github.com/cosmos/cosmos-sdk/types"
	authtypes "github.com/cosmos/cosmos-sdk/x/auth/types"
	banktypes "github.com/cosmos/cosmos-sdk/x/bank/types"
	crisistypes "github.com/cosmos/cosmos-sdk/x/crisis/types"
	distrtypes "github.com/cosmos/cosmos-sdk/x/distribution/types"
	govtypes "github.com/cosmos/cosmos-sdk/x/gov/types"
	slashingtypes "github.com/cosmos/cosmos-sdk/x/slashing/types"
	stakingtypes "github.com/cosmos/cosmos-sdk/x/staking/types"
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
)

type GenesisStates struct {
	DEXdropSupply   sdk.Coin
	BoostdropSupply sdk.Coin
	BondDenom       string
	GenesisTime     time.Time

	ConsensusParams     *tmproto.ConsensusParams
	AuthParams          authtypes.Params
	BankParams          banktypes.Params
	DistributionParams  distrtypes.Params
	StakingParams       stakingtypes.Params
	GovParams           govtypes.Params
	SlashingParams      slashingtypes.Params
	MintParams          minttypes.Params
	LiquidityParams     liquiditytypes.Params
	LiquidStakingParams liquidstakingtypes.Params
	FarmingParams       farmingtypes.Params
	BudgetParams        budgettypes.Params
	BankGenesisStates   banktypes.GenesisState
	CrisisStates        crisistypes.GenesisState
	ClaimGenesisState   claimtypes.GenesisState
}

func NewGenesisState() *GenesisStates {
	genParams := &GenesisStates{}
	genParams.BondDenom = "ucre"
	genParams.DEXdropSupply = sdk.NewInt64Coin(genParams.BondDenom, 50_000_000_000_000)   // 50mil
	genParams.BoostdropSupply = sdk.NewInt64Coin(genParams.BondDenom, 50_000_000_000_000) // 50mil

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

	// Set auth params
	genParams.AuthParams = authtypes.DefaultParams()
	genParams.AuthParams.MaxMemoCharacters = 512

	// Set bank params
	genParams.BankParams = banktypes.DefaultParams()

	// Set crisis genesis states
	genParams.CrisisStates = crisistypes.GenesisState{
		ConstantFee: sdk.NewInt64Coin(genParams.BondDenom, 1000),
	}

	// Set distribution params
	genParams.DistributionParams = distrtypes.Params{
		CommunityTax:        sdk.MustNewDecFromStr("0.285714285700000000"),
		BaseProposerReward:  sdk.MustNewDecFromStr("0.007142857143000000"),
		BonusProposerReward: sdk.MustNewDecFromStr("0.028571428570000000"),
		WithdrawAddrEnabled: true,
	}

	// Set staking params
	genParams.StakingParams = stakingtypes.Params{
		UnbondingTime:     1209600 * time.Second, // 2 weeks
		MaxValidators:     50,
		MaxEntries:        28,
		HistoricalEntries: 10000,
		BondDenom:         genParams.BondDenom,
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

	// Set slashing params
	genParams.SlashingParams = slashingtypes.Params{
		SignedBlocksWindow:      30000,
		MinSignedPerWindow:      sdk.MustNewDecFromStr("0.050000000000000000"),
		DowntimeJailDuration:    60 * time.Second,
		SlashFractionDoubleSign: sdk.MustNewDecFromStr("0.050000000000000000"),
		SlashFractionDowntime:   sdk.MustNewDecFromStr("0.000000000000000000"),
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
		LiquidBondDenom:       "ubcre",
		WhitelistedValidators: []liquidstakingtypes.WhitelistedValidator{
			// {
			// 	ValidatorAddress: "crevaloper1s96rxwvhrv4zn39v8haulhexflvjjp50j596ug",
			// 	TargetWeight:     sdk.NewInt(10),
			// },
			// {
			// 	ValidatorAddress: "crevaloper1jwjph8k3933uuejyhvnptmnxf4afve876vnx6k",
			// 	TargetWeight:     sdk.NewInt(10),
			// },
			// {
			// 	ValidatorAddress: "crevaloper1ckn4wlv5repm4lj62y9nwyvyvk63ydrxqt5t6q",
			// 	TargetWeight:     sdk.NewInt(10),
			// },
			// {
			// 	ValidatorAddress: "crevaloper1g7lz8463vkmdjtzj2a8s4lwz2xksfnk3838quf",
			// 	TargetWeight:     sdk.NewInt(10),
			// },
			// {
			// 	ValidatorAddress: "crevaloper1fksh8k3dhggajvm2mm433c2dr0jeq8kun5eqcg",
			// 	TargetWeight:     sdk.NewInt(10),
			// },
			// {
			// 	ValidatorAddress: "crevaloper1scdg75uqv3j5kcsh089ksqmyx590mjz4n4ep9s",
			// 	TargetWeight:     sdk.NewInt(10),
			// },
			// {
			// 	ValidatorAddress: "crevaloper10tzu9srek0masgefjsgqpyyvm5jywgwwj8nwen",
			// 	TargetWeight:     sdk.NewInt(10),
			// },
			// {
			// 	ValidatorAddress: "crevaloper1x5wgh6vwye60wv3dtshs9dmqggwfx2ld4uln5g",
			// 	TargetWeight:     sdk.NewInt(10),
			// },
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

	// Set gov params
	genParams.GovParams = govtypes.Params{
		DepositParams: govtypes.DepositParams{
			MinDeposit: sdk.NewCoins(
				sdk.NewInt64Coin(genParams.BondDenom, 500000000),
			),
			MaxDepositPeriod: 432000 * time.Second, // 5 days
		},
		VotingParams: govtypes.VotingParams{
			VotingPeriod: 432000 * time.Second,
		},
		TallyParams: govtypes.TallyParams{
			Quorum:        sdk.MustNewDecFromStr("0.400000000000000000"),
			Threshold:     sdk.MustNewDecFromStr("0.500000000000000000"),
			VetoThreshold: sdk.MustNewDecFromStr("0.334000000000000000"),
		},
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
	airdrop := claimtypes.Airdrop{
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
	}
	genParams.ClaimGenesisState.Airdrops = []claimtypes.Airdrop{airdrop}

	// // Parse claim records, balances, and total initial genesis coin from the airdrop result file
	// records, balances, totalInitialGenesisCoin := parseClaimRecords(genParams)

	// // Deduct 20% initial airdrop amount
	// dexDropSupply := genParams.DEXdropSupply.Sub(totalInitialGenesisCoin)

	// // Set source account balances
	// balances = append(balances, banktypes.Balance{
	// 	Address: airdrop.SourceAddress,
	// 	Coins:   sdk.NewCoins(dexDropSupply.Add(genParams.BoostdropSupply)), // DEXDropSupply + BoostDropSupply
	// })

	// // Add accounts
	// newBalances, totalCoins := addAccounts(genParams)
	// balances = append(balances, newBalances...)

	// // Set claim genesis states
	// genParams.ClaimGenesisState.ClaimRecords = records
	// genParams.BankGenesisStates.Balances = balances

	// // Set supply genesis states
	// // Total supply = DEXDropSupply + BoostDropSupply + TotalCoins
	// genParams.BankGenesisStates.Supply = sdk.NewCoins(genParams.DEXdropSupply.Add(genParams.BoostdropSupply)).Add(totalCoins...)

	return genParams
}
