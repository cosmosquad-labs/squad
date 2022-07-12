package keeper_test

import (
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/cosmosquad-labs/squad/v2/x/liquidfarming/types"

	_ "github.com/stretchr/testify/suite"
)

func (s *KeeperTestSuite) TestLastQueuedFarmingId() {
	poolId := uint64(1)
	reqId := s.keeper.GetLastQueuedFarmingId(s.ctx, poolId)
	s.Require().Equal(uint64(0), reqId)

	cacheCtx, _ := s.ctx.CacheContext()
	nextId := s.keeper.GetNextQueuedFarmingIdWithUpdate(cacheCtx, poolId)
	s.Require().Equal(uint64(1), nextId)

	s.createPair(s.addr(0), "denom1", "denom2", true)
	s.createPool(s.addr(0), 1, sdk.NewCoins(
		sdk.NewInt64Coin("denom1", 100000000),
		sdk.NewInt64Coin("denom2", 100000000)), true)
	s.createLiquidFarm([]types.LiquidFarm{
		{PoolId: poolId, MinimumDepositAmount: sdk.NewInt(1), MinimumBidAmount: sdk.NewInt(1)}})
	s.farm(poolId, s.addr(0), sdk.NewInt64Coin("pool1", 100000000), true)

	nextId = s.keeper.GetNextQueuedFarmingIdWithUpdate(cacheCtx, poolId)
	s.Require().Equal(uint64(2), nextId)
}

func (s *KeeperTestSuite) TestLastRewardsAuctionId() {
	// TODO: not implemented yet
}

func (s *KeeperTestSuite) TestDepositRequest() {
	// TODO: not implemented yet
	// Set | Get | Delete
}

func (s *KeeperTestSuite) TestIterateRewardsAuctions() {
	// TODO: not implemented yet
	// Set | Get | Delete
}
