package keeper_test

import (
	"context"
	"testing"

	"github.com/stretchr/testify/suite"
	abci "github.com/tendermint/tendermint/abci/types"
	tmproto "github.com/tendermint/tendermint/proto/tendermint/types"

	sdk "github.com/cosmos/cosmos-sdk/types"

	chain "github.com/cosmosquad-labs/squad/v2/app"
	utils "github.com/cosmosquad-labs/squad/v2/types"
	"github.com/cosmosquad-labs/squad/v2/x/lending/keeper"
	"github.com/cosmosquad-labs/squad/v2/x/lending/types"
)

type KeeperTestSuite struct {
	suite.Suite

	app       *chain.App
	ctx       sdk.Context
	goCtx     context.Context
	keeper    keeper.Keeper
	msgServer types.MsgServer
}

func TestKeeperTestSuite(t *testing.T) {
	suite.Run(t, new(KeeperTestSuite))
}

func (s *KeeperTestSuite) SetupTest() {
	s.app = chain.Setup(false)
	hdr := tmproto.Header{
		Height: 1,
		Time:   utils.ParseTime("2022-01-01T00:00:00Z"),
	}
	s.app.BeginBlock(abci.RequestBeginBlock{Header: hdr})
	s.ctx = s.app.BaseApp.NewContext(false, hdr)
	s.goCtx = sdk.WrapSDKContext(s.ctx)
	s.app.BeginBlocker(s.ctx, abci.RequestBeginBlock{Header: s.ctx.BlockHeader()})
	s.keeper = s.app.LendingKeeper
	s.msgServer = keeper.NewMsgServerImpl(s.keeper)
}

func (s *KeeperTestSuite) lend(lender sdk.AccAddress, coin sdk.Coin, fund bool) error {
	s.T().Helper()
	if fund {
		s.fundAddr(lender, sdk.NewCoins(coin))
	}
	_, err := s.msgServer.Lend(s.goCtx, types.NewMsgLend(lender, coin))
	return err
}

func (s *KeeperTestSuite) fundAddr(addr sdk.AccAddress, amt sdk.Coins) {
	s.T().Helper()
	err := s.app.BankKeeper.MintCoins(s.ctx, types.ModuleName, amt)
	s.Require().NoError(err)
	err = s.app.BankKeeper.SendCoinsFromModuleToAccount(s.ctx, types.ModuleName, addr, amt)
	s.Require().NoError(err)
}

func (s *KeeperTestSuite) getBalances(addr sdk.AccAddress) sdk.Coins {
	s.T().Helper()
	return s.app.BankKeeper.GetAllBalances(s.ctx, addr)
}

//nolint
func (s *KeeperTestSuite) getBalance(addr sdk.AccAddress, denom string) sdk.Coin {
	s.T().Helper()
	return s.app.BankKeeper.GetBalance(s.ctx, addr, denom)
}
