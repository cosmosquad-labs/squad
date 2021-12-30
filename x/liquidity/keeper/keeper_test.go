package keeper_test

import (
	"testing"
	"time"

	"github.com/stretchr/testify/suite"

	sdk "github.com/cosmos/cosmos-sdk/types"
	tmproto "github.com/tendermint/tendermint/proto/tendermint/types"

	simapp "github.com/tendermint/farming/app"
	"github.com/tendermint/farming/x/liquidity/keeper"
	"github.com/tendermint/farming/x/liquidity/types"
)

const (
	denom1 = "denom1"
	denom2 = "denom2"
)

var (
	initialBalances = sdk.NewCoins(
		sdk.NewInt64Coin(sdk.DefaultBondDenom, 100_000_000_000_000),
		sdk.NewInt64Coin(denom1, 100_000_000_000_000),
		sdk.NewInt64Coin(denom2, 100_000_000_000_000),
	)
)

type KeeperTestSuite struct {
	suite.Suite

	app     *simapp.FarmingApp
	ctx     sdk.Context
	keeper  keeper.Keeper
	querier keeper.Querier
	srv     types.MsgServer
	addrs   []sdk.AccAddress
}

func TestKeeperTestSuite(t *testing.T) {
	suite.Run(t, new(KeeperTestSuite))
}

func (suite *KeeperTestSuite) SetupTest() {
	app := simapp.Setup(false)
	ctx := app.BaseApp.NewContext(false, tmproto.Header{})

	suite.app = app
	suite.ctx = ctx
	suite.ctx = suite.ctx.WithBlockTime(time.Now()) // set to current time
	suite.keeper = suite.app.LiquidityKeeper
	suite.querier = keeper.Querier{Keeper: suite.keeper}
	suite.srv = keeper.NewMsgServerImpl(suite.keeper)
	suite.addrs = simapp.AddTestAddrs(suite.app, suite.ctx, 6, sdk.ZeroInt())
	for _, addr := range suite.addrs {
		err := simapp.FundAccount(suite.app.BankKeeper, suite.ctx, addr, initialBalances)
		suite.Require().NoError(err)
	}
	// TODO: not implemented yet
}
