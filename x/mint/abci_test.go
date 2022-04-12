package mint_test

import (
	"fmt"
	"testing"
	"time"

	"github.com/stretchr/testify/require"
	"github.com/stretchr/testify/suite"

	sdk "github.com/cosmos/cosmos-sdk/types"
	authtypes "github.com/cosmos/cosmos-sdk/x/auth/types"
	abcitypes "github.com/tendermint/tendermint/abci/types"
	tmproto "github.com/tendermint/tendermint/proto/tendermint/types"

	chain "github.com/cosmosquad-labs/squad/app"
	utils "github.com/cosmosquad-labs/squad/types"
	"github.com/cosmosquad-labs/squad/x/mint"
	"github.com/cosmosquad-labs/squad/x/mint/keeper"
	"github.com/cosmosquad-labs/squad/x/mint/types"
)

var (
	initialBalances = sdk.NewCoins(
		sdk.NewInt64Coin(sdk.DefaultBondDenom, 1_000_000_000),
	)
)

type ModuleTestSuite struct {
	suite.Suite

	app    *chain.App
	ctx    sdk.Context
	keeper keeper.Keeper
	addrs  []sdk.AccAddress
}

func TestModuleTestSuite(t *testing.T) {
	suite.Run(t, new(ModuleTestSuite))
}

func (suite *ModuleTestSuite) SetupTest() {
	app := chain.Setup(false)
	ctx := app.BaseApp.NewContext(false, tmproto.Header{})

	suite.app = app
	suite.ctx = ctx
	suite.keeper = suite.app.MintKeeper
	suite.addrs = chain.AddTestAddrs(suite.app, suite.ctx, 6, sdk.ZeroInt())
	for _, addr := range suite.addrs {
		err := chain.FundAccount(suite.app.BankKeeper, suite.ctx, addr, initialBalances)
		suite.Require().NoError(err)
	}
}

func (s *ModuleTestSuite) TestInitGenesis() {
	// default gent state case
	genState := types.DefaultGenesisState()
	mint.InitGenesis(s.ctx, s.app.MintKeeper, s.app.AccountKeeper, genState)
	got := mint.ExportGenesis(s.ctx, s.app.MintKeeper)
	s.Require().Equal(*genState, *got)

	// not nil last block time case
	testTime := utils.ParseTime("2023-01-01T00:00:00Z")
	genState.LastBlockTime = &testTime
	mint.InitGenesis(s.ctx, s.app.MintKeeper, s.app.AccountKeeper, genState)
	got = mint.ExportGenesis(s.ctx, s.app.MintKeeper)
	s.Require().Equal(*genState, *got)

	// invalid last block time case
	testTime2 := time.Unix(-62136697901, 0)
	genState.LastBlockTime = &testTime2
	s.Require().Panics(func() {
		mint.InitGenesis(s.ctx, s.app.MintKeeper, s.app.AccountKeeper, genState)
	})
	got = mint.ExportGenesis(s.ctx, s.app.MintKeeper)
	s.Require().NotEqual(*genState, *got)
}

func (s *ModuleTestSuite) TestImportExportGenesis() {
	k, ctx := s.keeper, s.ctx
	genState := mint.ExportGenesis(ctx, k)
	bz := s.app.AppCodec().MustMarshalJSON(genState)

	var genState2, genState5 types.GenesisState
	s.app.AppCodec().MustUnmarshalJSON(bz, &genState2)
	mint.InitGenesis(ctx, s.app.MintKeeper, s.app.AccountKeeper, &genState2)

	genState3 := mint.ExportGenesis(ctx, k)
	s.Require().Equal(*genState, genState2)
	s.Require().Equal(genState2, *genState3)

	ctx = ctx.WithBlockTime(utils.ParseTime("2022-01-01T00:00:00Z"))
	mint.BeginBlocker(ctx, k)
	genState4 := mint.ExportGenesis(ctx, k)
	bz = s.app.AppCodec().MustMarshalJSON(genState4)
	s.app.AppCodec().MustUnmarshalJSON(bz, &genState5)
	s.Require().Equal(*genState5.LastBlockTime, utils.ParseTime("2022-01-01T00:00:00Z"))
	mint.InitGenesis(s.ctx, s.app.MintKeeper, s.app.AccountKeeper, &genState5)
	genState6 := mint.ExportGenesis(ctx, k)
	s.Require().Equal(*genState4, genState5, genState6)
}

func TestConstantInflation(t *testing.T) {
	app := chain.Setup(false)
	ctx := app.BaseApp.NewContext(false, tmproto.Header{})

	app.InitChain(
		abcitypes.RequestInitChain{
			AppStateBytes: []byte("{}"),
			ChainId:       "test-chain-id",
		},
	)

	blockTime := 5 * time.Second

	feeCollector := app.AccountKeeper.GetModuleAddress(authtypes.FeeCollectorName)
	advanceHeight := func() sdk.Int {
		ctx = ctx.WithBlockHeight(ctx.BlockHeight() + 1).WithBlockTime(ctx.BlockTime().Add(blockTime))
		beforeBalance := app.BankKeeper.GetBalance(ctx, feeCollector, sdk.DefaultBondDenom)
		mint.BeginBlocker(ctx, app.MintKeeper)
		afterBalance := app.BankKeeper.GetBalance(ctx, feeCollector, sdk.DefaultBondDenom)
		mintedAmt := afterBalance.Sub(beforeBalance)
		require.False(t, mintedAmt.IsNegative())
		return mintedAmt.Amount
	}

	ctx = ctx.WithBlockHeight(0).WithBlockTime(utils.ParseTime("2022-01-01T00:00:00Z"))

	// skip first block inflation, not set LastBlockTime
	require.EqualValues(t, advanceHeight(), sdk.NewInt(0))

	// after 2022-01-01 00:00:00
	// 47564687 / 5 * (365 * 24 * 60 * 60) / 300000000000000 ~= 1
	// 47564687 ~= 300000000000000 / (365 * 24 * 60 * 60) * 5
	require.EqualValues(t, advanceHeight(), sdk.NewInt(47564687))
	require.EqualValues(t, advanceHeight(), sdk.NewInt(47564687))
	require.EqualValues(t, advanceHeight(), sdk.NewInt(47564687))
	require.EqualValues(t, advanceHeight(), sdk.NewInt(47564687))

	ctx = ctx.WithBlockHeight(100).WithBlockTime(utils.ParseTime("2023-01-01T00:00:00Z"))

	// applied 10sec(params.BlockTimeThreshold) block time due to block time diff is over params.BlockTimeThreshold
	require.EqualValues(t, advanceHeight(), sdk.NewInt(63419583))
	require.EqualValues(t, advanceHeight(), sdk.NewInt(31709791))

	// 317097919 / 5 * (365 * 24 * 60 * 60) / 200000000000000 ~= 1
	// 317097919 ~= 200000000000000 / (365 * 24 * 60 * 60) * 5
	require.EqualValues(t, advanceHeight(), sdk.NewInt(31709791))
	require.EqualValues(t, advanceHeight(), sdk.NewInt(31709791))
	require.EqualValues(t, advanceHeight(), sdk.NewInt(31709791))
	require.EqualValues(t, advanceHeight(), sdk.NewInt(31709791))

	blockTime = 10 * time.Second
	// 634195839 / 10 * (365 * 24 * 60 * 60) / 200000000000000 ~= 1
	// 634195839 ~= 200000000000000 / (365 * 24 * 60 * 60) * 10
	require.EqualValues(t, advanceHeight(), sdk.NewInt(63419583))
	require.EqualValues(t, advanceHeight(), sdk.NewInt(63419583))

	// over BlockTimeThreshold 10sec
	blockTime = 20 * time.Second
	require.EqualValues(t, advanceHeight(), sdk.NewInt(63419583))
	require.EqualValues(t, advanceHeight(), sdk.NewInt(63419583))

	// no inflation
	ctx = ctx.WithBlockHeight(300).WithBlockTime(utils.ParseTime("2030-01-01T01:00:00Z"))
	require.True(t, advanceHeight().IsZero())
	require.True(t, advanceHeight().IsZero())
	require.True(t, advanceHeight().IsZero())
	require.True(t, advanceHeight().IsZero())
}

func (s *ModuleTestSuite) TestDefaultGenesis() {
	genState := *types.DefaultGenesisState()

	mint.InitGenesis(s.ctx, s.app.MintKeeper, s.app.AccountKeeper, &genState)
	got := mint.ExportGenesis(s.ctx, s.app.MintKeeper)
	s.Require().Equal(genState, *got)
}

func (s *ModuleTestSuite) TestImportExportGenesisEmpty() {
	emptyParams := types.DefaultParams()
	emptyParams.InflationSchedules = []types.InflationSchedule{}
	s.app.MintKeeper.SetParams(s.ctx, emptyParams)
	genState := mint.ExportGenesis(s.ctx, s.app.MintKeeper)

	var genState2 types.GenesisState
	bz := s.app.AppCodec().MustMarshalJSON(genState)
	s.app.AppCodec().MustUnmarshalJSON(bz, &genState2)
	mint.InitGenesis(s.ctx, s.app.MintKeeper, s.app.AccountKeeper, &genState2)

	genState3 := mint.ExportGenesis(s.ctx, s.app.MintKeeper)
	s.Require().Equal(*genState, genState2)
	s.Require().EqualValues(genState2, *genState3)
}

/////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////

func (s *ModuleTestSuite) TestSimualte_ConstantInflation() {
	blockTime := 5 * time.Second

	params := s.keeper.GetParams(s.ctx)
	params.BlockTimeThreshold = 10 * time.Second
	params.InflationSchedules = []types.InflationSchedule{
		{
			StartTime: s.ctx.BlockTime(),
			EndTime:   s.ctx.BlockTime().AddDate(1, 0, 0),
			Amount:    sdk.NewInt(108_700000_000000),
		},
		{
			StartTime: s.ctx.BlockTime().AddDate(1, 0, 0),
			EndTime:   s.ctx.BlockTime().AddDate(2, 0, 0),
			Amount:    sdk.NewInt(216_100000_000000),
		},
		{
			StartTime: s.ctx.BlockTime().AddDate(2, 0, 0),
			EndTime:   s.ctx.BlockTime().AddDate(3, 0, 0),
			Amount:    sdk.NewInt(151_300000_000000),
		},
		{
			StartTime: s.ctx.BlockTime().AddDate(3, 0, 0),
			EndTime:   s.ctx.BlockTime().AddDate(4, 0, 0),
			Amount:    sdk.NewInt(105_900000_000000),
		},
		{
			StartTime: s.ctx.BlockTime().AddDate(4, 0, 0),
			EndTime:   s.ctx.BlockTime().AddDate(5, 0, 0),
			Amount:    sdk.NewInt(74_100000_000000),
		},
		{
			StartTime: s.ctx.BlockTime().AddDate(5, 0, 0),
			EndTime:   s.ctx.BlockTime().AddDate(6, 0, 0),
			Amount:    sdk.NewInt(51_900000_000000),
		},
		{
			StartTime: s.ctx.BlockTime().AddDate(6, 0, 0),
			EndTime:   s.ctx.BlockTime().AddDate(7, 0, 0),
			Amount:    sdk.NewInt(36_300000_000000),
		},
		{
			StartTime: s.ctx.BlockTime().AddDate(7, 0, 0),
			EndTime:   s.ctx.BlockTime().AddDate(8, 0, 0),
			Amount:    sdk.NewInt(25_400000_000000),
		},
		{
			StartTime: s.ctx.BlockTime().AddDate(8, 0, 0),
			EndTime:   s.ctx.BlockTime().AddDate(9, 0, 0),
			Amount:    sdk.NewInt(17_800000_000000),
		},
		{
			StartTime: s.ctx.BlockTime().AddDate(9, 0, 0),
			EndTime:   s.ctx.BlockTime().AddDate(10, 0, 0),
			Amount:    sdk.NewInt(12_500000_000000),
		},
	}
	s.keeper.SetParams(s.ctx, params)

	feeCollector := s.app.AccountKeeper.GetModuleAccount(s.ctx, authtypes.FeeCollectorName)

	advanceHeight := func() sdk.Int {
		nextBlockHeight := s.ctx.BlockHeight() + 1
		nextBlockTime := s.ctx.BlockTime().Add(blockTime)
		s.ctx = s.ctx.WithBlockHeight(nextBlockHeight)
		s.ctx = s.ctx.WithBlockTime(nextBlockTime)
		fmt.Printf("Advance to Next Block Height: %d\n", nextBlockHeight)

		beforeBalance := s.app.BankKeeper.GetBalance(s.ctx, feeCollector.GetAddress(), sdk.DefaultBondDenom)
		mint.BeginBlocker(s.ctx, s.app.MintKeeper)
		afterBalance := s.app.BankKeeper.GetBalance(s.ctx, feeCollector.GetAddress(), sdk.DefaultBondDenom)
		mintedAmt := afterBalance.Sub(beforeBalance)
		s.Require().False(mintedAmt.IsNegative())
		fmt.Printf("Minted Amount: %s\n", mintedAmt)
		return mintedAmt.Amount
	}

	// Start a block time
	startTime := s.ctx.BlockTime()
	s.ctx = s.ctx.WithBlockHeight(0).WithBlockTime(startTime)
	fmt.Println("--- ", startTime)

	// Skip first block inflation since LastBlockTime is not set
	s.Require().EqualValues(advanceHeight(), sdk.NewInt(0))

	// after 2022-01-01 00:00:00
	// 47564687 / 5 * (365 * 24 * 60 * 60) / 300000000000000 ~= 1
	// 47564687 ~= 300000000000000 / (365 * 24 * 60 * 60) * 5
	fmt.Println("advanceHeight(): ", advanceHeight())
	fmt.Println("advanceHeight(): ", advanceHeight())
	fmt.Println("advanceHeight(): ", advanceHeight())
	// s.Require().EqualValues(advanceHeight(), sdk.NewInt(47564687))
	// s.Require().EqualValues(advanceHeight(), sdk.NewInt(47564687))
	// s.Require().EqualValues(advanceHeight(), sdk.NewInt(47564687))
	// s.Require().EqualValues(advanceHeight(), sdk.NewInt(47564687))

	startTime = s.ctx.BlockTime().AddDate(1, 0, 0)
	s.ctx = s.ctx.WithBlockHeight(100).WithBlockTime(startTime)
	fmt.Println("--- ", startTime)

	return

	// applied 10sec(params.BlockTimeThreshold) block time due to block time diff is over params.BlockTimeThreshold
	s.Require().EqualValues(advanceHeight(), sdk.NewInt(63419583))
	s.Require().EqualValues(advanceHeight(), sdk.NewInt(31709791))

	// 317097919 / 5 * (365 * 24 * 60 * 60) / 200000000000000 ~= 1
	// 317097919 ~= 200000000000000 / (365 * 24 * 60 * 60) * 5
	s.Require().EqualValues(advanceHeight(), sdk.NewInt(31709791))
	s.Require().EqualValues(advanceHeight(), sdk.NewInt(31709791))
	s.Require().EqualValues(advanceHeight(), sdk.NewInt(31709791))
	s.Require().EqualValues(advanceHeight(), sdk.NewInt(31709791))

	blockTime = 10 * time.Second
	// 634195839 / 10 * (365 * 24 * 60 * 60) / 200000000000000 ~= 1
	// 634195839 ~= 200000000000000 / (365 * 24 * 60 * 60) * 10
	s.Require().EqualValues(advanceHeight(), sdk.NewInt(63419583))
	s.Require().EqualValues(advanceHeight(), sdk.NewInt(63419583))

	// over BlockTimeThreshold 10sec
	blockTime = 20 * time.Second
	s.Require().EqualValues(advanceHeight(), sdk.NewInt(63419583))
	s.Require().EqualValues(advanceHeight(), sdk.NewInt(63419583))

	// no inflation
	startTime = startTime.AddDate(1, 0, 0)
	s.ctx = s.ctx.WithBlockHeight(300).WithBlockTime(startTime)
	fmt.Println("--- ", startTime)

}
