package keeper_test

import (
	"github.com/cosmosquad-labs/squad/v2/x/liquidfarming/types"

	_ "github.com/stretchr/testify/suite"
)

func (suite *KeeperTestSuite) TestDefaultGenesis() {
	genState := *types.DefaultGenesis()

	suite.keeper.InitGenesis(suite.ctx, genState)
	got := suite.keeper.ExportGenesis(suite.ctx)
	suite.Require().Equal(genState, *got)
}

func (s *KeeperTestSuite) TestImportExportGenesis() {
	// TODO: not implemented yet
}
