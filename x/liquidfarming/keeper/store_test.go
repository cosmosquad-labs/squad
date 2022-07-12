package keeper_test

import (
	"fmt"

	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/cosmosquad-labs/squad/v2/x/liquidfarming/types"

	_ "github.com/stretchr/testify/suite"
)

func (s *KeeperTestSuite) TestQueuedFarming() {
	// TODO: not implemented yet
	// Set | Get | Delete
}

// Test
func (s *KeeperTestSuite) TestQueuedFarmingByFarmerAndDenom() {
	poolId := uint64(1)
	liquidFarm := types.LiquidFarm{
		PoolId:               poolId,
		MinimumDepositAmount: sdk.ZeroInt(),
		MinimumBidAmount:     sdk.ZeroInt(),
	}
	farmerAcc := s.addr(0)

	s.createPair(farmerAcc, "denom1", "denom2", true)
	s.createPool(farmerAcc, 1, sdk.NewCoins(sdk.NewInt64Coin("denom1", 100000000), sdk.NewInt64Coin("denom2", 100000000)), true)
	s.createLiquidFarm([]types.LiquidFarm{liquidFarm})
	s.farm(poolId, farmerAcc, sdk.NewInt64Coin("pool1", 100000000), true)

	queuedFarming, found := s.keeper.GetQueuedFarmingByFarmer(s.ctx, farmerAcc, "lf1")
	fmt.Println("found: ", found)
	fmt.Println("queuedFarming: ", queuedFarming)
}
