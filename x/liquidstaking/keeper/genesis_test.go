package keeper_test

import (
	"github.com/tendermint/farming/x/liquidstaking/types"

	_ "github.com/stretchr/testify/suite"
)

func (suite *KeeperTestSuite) TestInitGenesis() {
	suite.SetupTest()
	params := suite.keeper.GetParams(suite.ctx)
	params.Bearings = suite.bearings[:4]
	suite.keeper.SetParams(suite.ctx, params)

	emptyGenState := suite.keeper.ExportGenesis(suite.ctx)
	suite.Require().NotPanics(func() {
		suite.keeper.InitGenesis(suite.ctx, *emptyGenState)
	})
	suite.Require().Equal(emptyGenState, suite.keeper.ExportGenesis(suite.ctx))
	suite.Require().Nil(emptyGenState.BearingRecords)

	err := suite.keeper.CollectBearings(suite.ctx)
	suite.Require().NoError(err)

	var genState *types.GenesisState
	suite.Require().NotPanics(func() {
		genState = suite.keeper.ExportGenesis(suite.ctx)
	})
	err = types.ValidateGenesis(*genState)
	suite.Require().NoError(err)

	suite.Require().NotNil(genState.BearingRecords)
	suite.Require().NotPanics(func() {
		suite.keeper.InitGenesis(suite.ctx, *genState)
	})
	suite.Require().Equal(genState, suite.keeper.ExportGenesis(suite.ctx))
}
