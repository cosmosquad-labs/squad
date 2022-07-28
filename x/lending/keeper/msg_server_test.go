package keeper_test

import (
	"fmt"

	utils "github.com/cosmosquad-labs/squad/v2/types"
	"github.com/cosmosquad-labs/squad/v2/x/lending/types"
)

func (s *KeeperTestSuite) TestLend() {
	s.keeper.SetLendingAssetParams(s.ctx, []types.LendingAssetParam{
		types.NewLendingAssetParam("stake", utils.ParseDec("0.01"), utils.ParseDec("0.1"), 1),
	})

	lender := utils.TestAddress(0)
	err := s.lend(lender, utils.ParseCoin("1000000stake"), true)
	s.Require().NoError(err)

	fmt.Println(s.getBalances(lender))

	err = s.lend(lender, utils.ParseCoin("555555stake"), true)
	s.Require().NoError(err)

	fmt.Println(s.getBalances(lender))
}
