package keeper_test

import (
	"testing"
	"time"

	govtypes "github.com/cosmos/cosmos-sdk/x/gov/types"
	minttypes "github.com/cosmos/cosmos-sdk/x/mint/types"
	"github.com/stretchr/testify/suite"
	abcitypes "github.com/tendermint/tendermint/abci/types"
	tmproto "github.com/tendermint/tendermint/proto/tendermint/types"

	sdk "github.com/cosmos/cosmos-sdk/types"

	chain "github.com/cosmosquad-labs/squad/v2/app"

	"github.com/cosmosquad-labs/squad/v2/x/marketmaker"
	"github.com/cosmosquad-labs/squad/v2/x/marketmaker/keeper"
	"github.com/cosmosquad-labs/squad/v2/x/marketmaker/types"
)

const (
	denom1 = "denom1"
	denom2 = "denom2"
	denom3 = "denom3"
)

var (
	initialBalances = sdk.NewCoins(
		sdk.NewInt64Coin(sdk.DefaultBondDenom, 1_000_000_000_000),
		sdk.NewInt64Coin(denom1, 1_000_000_000),
		sdk.NewInt64Coin(denom2, 1_000_000_000),
		sdk.NewInt64Coin(denom3, 1_000_000_000))
)

type KeeperTestSuite struct {
	suite.Suite

	app        *chain.App
	ctx        sdk.Context
	keeper     keeper.Keeper
	querier    keeper.Querier
	msgServer  types.MsgServer
	govHandler govtypes.Handler
	addrs      []sdk.AccAddress
}

func TestKeeperTestSuite(t *testing.T) {
	suite.Run(t, new(KeeperTestSuite))
}

func (suite *KeeperTestSuite) SetupTest() {
	app := chain.Setup(false)
	ctx := app.BaseApp.NewContext(false, tmproto.Header{})

	suite.app = app
	suite.ctx = ctx
	suite.keeper = suite.app.MarketMakerKeeper
	suite.querier = keeper.Querier{Keeper: suite.keeper}
	suite.msgServer = keeper.NewMsgServerImpl(suite.keeper)
	suite.govHandler = marketmaker.NewMarketMakerProposalHandler(suite.keeper)
	suite.addrs = chain.AddTestAddrs(suite.app, suite.ctx, 6, sdk.ZeroInt())
	for _, addr := range suite.addrs {
		err := chain.FundAccount(suite.app.BankKeeper, suite.ctx, addr, initialBalances)
		suite.Require().NoError(err)
	}
	suite.SetIncentivePairs()
}

func (suite *KeeperTestSuite) AddTestAddrs(num int, coins sdk.Coins) []sdk.AccAddress {
	addrs := chain.AddTestAddrs(suite.app, suite.ctx, num, sdk.ZeroInt())
	for _, addr := range addrs {
		err := chain.FundAccount(suite.app.BankKeeper, suite.ctx, addr, coins)
		suite.Require().NoError(err)
	}
	return addrs
}

func (suite *KeeperTestSuite) SetIncentivePairs() {
	params := suite.keeper.GetParams(suite.ctx)
	params.IncentivePairs = []types.IncentivePair{
		{
			PairName: "pair1",
			PairId:   uint64(1),
		},
		{
			PairName: "pair2",
			PairId:   uint64(2),
		},
		{
			PairName: "pair3",
			PairId:   uint64(3),
		},
		{
			PairName: "pair4",
			PairId:   uint64(4),
		},
		{
			PairName: "pair5",
			PairId:   uint64(5),
		},
		{
			PairName: "pair6",
			PairId:   uint64(6),
		},
		{
			PairName: "pair7",
			PairId:   uint64(7),
		},
	}
	suite.keeper.SetParams(suite.ctx, params)
}

func (suite *KeeperTestSuite) ResetIncentivePairs() {
	params := suite.keeper.GetParams(suite.ctx)
	params.IncentivePairs = []types.IncentivePair{}
	suite.keeper.SetParams(suite.ctx, params)
}

func (suite *KeeperTestSuite) handleProposal(content govtypes.Content) {
	suite.T().Helper()
	err := content.ValidateBasic()
	suite.Require().NoError(err)
	err = suite.govHandler(suite.ctx, content)
	suite.Require().NoError(err)
}

func (suite *KeeperTestSuite) addDenoms(denoms ...string) {
	suite.T().Helper()
	coins := sdk.Coins{}
	for _, denom := range denoms {
		coins = coins.Add(sdk.NewInt64Coin(denom, 1))
	}
	err := suite.app.BankKeeper.MintCoins(suite.ctx, minttypes.ModuleName, coins)
	suite.Require().NoError(err)
}

func (suite *KeeperTestSuite) addDenomsFromCoins(coins sdk.Coins) {
	var denoms []string
	for _, coin := range coins {
		denoms = append(denoms, coin.Denom)
	}
	suite.addDenoms(denoms...)
}

func (suite *KeeperTestSuite) addDenomsFromDecCoins(coins sdk.DecCoins) {
	var denoms []string
	for _, coin := range coins {
		denoms = append(denoms, coin.Denom)
	}
	suite.addDenoms(denoms...)
}

func (suite *KeeperTestSuite) executeBlock(blockTime time.Time, f func()) {
	suite.T().Helper()
	suite.ctx = suite.ctx.WithBlockTime(blockTime)
	suite.app.BeginBlocker(suite.ctx, abcitypes.RequestBeginBlock{})
	if f != nil {
		f()
	}
	suite.app.EndBlocker(suite.ctx, abcitypes.RequestEndBlock{})
}

func intEq(exp, got sdk.Int) (bool, string, string, string) {
	return exp.Equal(got), "expected:\t%v\ngot:\t\t%v", exp.String(), got.String()
}

func decEq(exp, got sdk.Dec) (bool, string, string, string) {
	return exp.Equal(got), "expected:\t%v\ngot:\t\t%v", exp.String(), got.String()
}

func coinsEq(exp, got sdk.Coins) (bool, string, string, string) {
	return exp.IsEqual(got), "expected:\t%v\ngot:\t\t%v", exp.String(), got.String()
}

func decCoinsEq(exp, got sdk.DecCoins) (bool, string, string, string) {
	return exp.IsEqual(got), "expected:\t%v\ngot:\t\t%v", exp.String(), got.String()
}

func parseCoins(s string) sdk.Coins {
	coins, err := sdk.ParseCoinsNormalized(s)
	if err != nil {
		panic(err)
	}
	return coins
}

func parseDecCoins(s string) sdk.DecCoins {
	decCoins, err := sdk.ParseDecCoins(s)
	if err != nil {
		panic(err)
	}
	return decCoins
}

func parseDec(s string) sdk.Dec {
	dec, err := sdk.NewDecFromStr(s)
	if err != nil {
		panic(err)
	}
	return dec
}
