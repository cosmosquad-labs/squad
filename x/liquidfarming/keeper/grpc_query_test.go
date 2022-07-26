package keeper_test

import (
	sdk "github.com/cosmos/cosmos-sdk/types"

	utils "github.com/cosmosquad-labs/squad/v2/types"
	"github.com/cosmosquad-labs/squad/v2/x/liquidfarming/types"
	liquiditytypes "github.com/cosmosquad-labs/squad/v2/x/liquidity/types"

	_ "github.com/stretchr/testify/suite"
)

func (s *KeeperTestSuite) TestGRPCParams() {
	resp, err := s.querier.Params(sdk.WrapSDKContext(s.ctx), &types.QueryParamsRequest{})
	s.Require().NoError(err)
	s.Require().Equal(s.keeper.GetParams(s.ctx), resp.Params)
}

func (s *KeeperTestSuite) TestGRPCLiquidFarms() {
	pair1 := s.createPair(s.addr(0), "denom1", "denom2", true)
	pool1 := s.createPool(s.addr(0), pair1.Id, utils.ParseCoins("100_000_000denom1, 100_000_000denom2"), true)
	minFarmAmt1, minBidAmt1 := sdk.NewInt(10_000_000), sdk.NewInt(10_000_000)
	s.createLiquidFarm(types.NewLiquidFarm(pool1.Id, minFarmAmt1, minBidAmt1))

	pair2 := s.createPair(s.addr(1), "denom3", "denom4", true)
	pool2 := s.createPool(s.addr(1), pair2.Id, utils.ParseCoins("100_000_000denom3, 100_000_000denom4"), true)
	minFarmAmt2, minBidAmt2 := sdk.NewInt(30_000_000), sdk.NewInt(30_000_000)
	s.createLiquidFarm(types.NewLiquidFarm(pool2.Id, minFarmAmt2, minBidAmt2))

	for _, tc := range []struct {
		name      string
		req       *types.QueryLiquidFarmsRequest
		expectErr bool
		postRun   func(*types.QueryLiquidFarmsResponse)
	}{
		{
			"query all liquidfarms",
			&types.QueryLiquidFarmsRequest{},
			false,
			func(resp *types.QueryLiquidFarmsResponse) {
				s.Require().Len(resp.LiquidFarms, 2)

				for _, liquidFarm := range resp.LiquidFarms {
					switch liquidFarm.PoolId {
					case 1:
						s.Require().Equal(minFarmAmt1, liquidFarm.MinimumFarmAmount)
						s.Require().Equal(minBidAmt1, liquidFarm.MinimumBidAmount)
					case 2:
						s.Require().Equal(minFarmAmt2, liquidFarm.MinimumFarmAmount)
						s.Require().Equal(minBidAmt2, liquidFarm.MinimumBidAmount)
					}
					reserveAcc, _ := sdk.AccAddressFromBech32(liquidFarm.LiquidFarmReserveAddress)
					poolCoinDenom := liquiditytypes.PoolCoinDenom(liquidFarm.PoolId)
					queuedAmt := s.app.FarmingKeeper.GetAllQueuedCoinsByFarmer(s.ctx, reserveAcc).AmountOf(poolCoinDenom)
					stakedAmt := s.app.FarmingKeeper.GetAllStakedCoinsByFarmer(s.ctx, reserveAcc).AmountOf(poolCoinDenom)
					s.Require().Equal(queuedAmt, liquidFarm.QueuedCoin.Amount)
					s.Require().Equal(stakedAmt, liquidFarm.StakedCoin.Amount)
				}
			},
		},
	} {
		s.Run(tc.name, func() {
			resp, err := s.querier.LiquidFarms(sdk.WrapSDKContext(s.ctx), tc.req)
			if tc.expectErr {
				s.Require().Error(err)
			} else {
				s.Require().NoError(err)
				tc.postRun(resp)
			}
		})
	}
}

func (s *KeeperTestSuite) TestGRPCLiquidFarm() {
	pair := s.createPair(s.addr(0), "denom1", "denom2", true)
	pool := s.createPool(s.addr(0), pair.Id, utils.ParseCoins("100_000_000denom1, 100_000_000denom2"), true)
	minFarmAmt, minBidAmt := sdk.NewInt(10_000_000), sdk.NewInt(10_000_000)
	s.createLiquidFarm(types.NewLiquidFarm(pool.Id, minFarmAmt, minBidAmt))

	for _, tc := range []struct {
		name      string
		req       *types.QueryLiquidFarmRequest
		expectErr bool
		postRun   func(*types.QueryLiquidFarmResponse)
	}{
		{
			"nil request",
			nil,
			true,
			nil,
		},
		{
			"query by invalid pool id",
			&types.QueryLiquidFarmRequest{
				PoolId: 5,
			},
			false,
			func(resp *types.QueryLiquidFarmResponse) {
				s.Require().Empty(resp.LiquidFarm)
			},
		},
		{
			"query by pool id",
			&types.QueryLiquidFarmRequest{
				PoolId: pool.Id,
			},
			false,
			func(resp *types.QueryLiquidFarmResponse) {
				reserveAcc, _ := sdk.AccAddressFromBech32(resp.LiquidFarm.LiquidFarmReserveAddress)
				poolCoinDenom := liquiditytypes.PoolCoinDenom(resp.LiquidFarm.PoolId)
				queuedAmt := s.app.FarmingKeeper.GetAllQueuedCoinsByFarmer(s.ctx, reserveAcc).AmountOf(poolCoinDenom)
				stakedAmt := s.app.FarmingKeeper.GetAllStakedCoinsByFarmer(s.ctx, reserveAcc).AmountOf(poolCoinDenom)
				s.Require().Equal(queuedAmt, resp.LiquidFarm.QueuedCoin.Amount)
				s.Require().Equal(stakedAmt, resp.LiquidFarm.StakedCoin.Amount)
				s.Require().Equal(types.LiquidFarmCoinDenom(pool.Id), resp.LiquidFarm.LFCoinDenom)
				s.Require().Equal(types.LiquidFarmReserveAddress(pool.Id).String(), resp.LiquidFarm.LiquidFarmReserveAddress)
				s.Require().Equal(minFarmAmt, resp.LiquidFarm.MinimumFarmAmount)
				s.Require().Equal(minBidAmt, resp.LiquidFarm.MinimumBidAmount)
			},
		},
	} {
		s.Run(tc.name, func() {
			resp, err := s.querier.LiquidFarm(sdk.WrapSDKContext(s.ctx), tc.req)
			if tc.expectErr {
				s.Require().Error(err)
			} else {
				s.Require().NoError(err)
				tc.postRun(resp)
			}
		})
	}
}

func (s *KeeperTestSuite) TestGRPCQueuedFarmings() {
	pair := s.createPair(s.addr(0), "denom1", "denom2", true)
	pool := s.createPool(s.addr(0), pair.Id, utils.ParseCoins("100_000_000denom1, 100_000_000denom2"), true)
	s.createLiquidFarm(types.NewLiquidFarm(pool.Id, sdk.ZeroInt(), sdk.ZeroInt()))

	s.deposit(s.addr(1), pool.Id, utils.ParseCoins("100_000_000denom1, 100_000_000denom2"), true)
	s.deposit(s.addr(2), pool.Id, utils.ParseCoins("100_000_000denom1, 100_000_000denom2"), true)
	s.deposit(s.addr(3), pool.Id, utils.ParseCoins("100_000_000denom1, 100_000_000denom2"), true)
	s.nextBlock()

	amount1, amount2, amount3 := sdk.NewInt(100_000), sdk.NewInt(500_000), sdk.NewInt(100_000_000)
	s.farm(pool.Id, s.addr(1), sdk.NewCoin(pool.PoolCoinDenom, amount1), true)
	s.farm(pool.Id, s.addr(2), sdk.NewCoin(pool.PoolCoinDenom, amount2), true)
	s.farm(pool.Id, s.addr(3), sdk.NewCoin(pool.PoolCoinDenom, amount3), true)
	s.nextBlock()

	for _, tc := range []struct {
		name      string
		req       *types.QueryQueuedFarmingsRequest
		expectErr bool
		postRun   func(*types.QueryQueuedFarmingsResponse)
	}{
		{
			"nil request",
			nil,
			true,
			nil,
		},
		{
			"query by invalid pool id",
			&types.QueryQueuedFarmingsRequest{
				PoolId: 5,
			},
			false,
			func(resp *types.QueryQueuedFarmingsResponse) {
				s.Require().Len(resp.QueuedFarmings, 0)
			},
		},
		{
			"query all queued farmings",
			&types.QueryQueuedFarmingsRequest{
				PoolId: pool.Id,
			},
			false,
			func(resp *types.QueryQueuedFarmingsResponse) {
				s.Require().Len(resp.QueuedFarmings, 3)
			},
		},
	} {
		s.Run(tc.name, func() {
			resp, err := s.querier.QueuedFarmings(sdk.WrapSDKContext(s.ctx), tc.req)
			if tc.expectErr {
				s.Require().Error(err)
			} else {
				s.Require().NoError(err)
				tc.postRun(resp)
			}
		})
	}
}

func (s *KeeperTestSuite) TestGRPCQueuedFarmingsByFarmer() {
	pair := s.createPair(s.addr(0), "denom1", "denom2", true)
	pool := s.createPool(s.addr(0), pair.Id, utils.ParseCoins("100_000_000denom1, 100_000_000denom2"), true)
	s.createLiquidFarm(types.NewLiquidFarm(pool.Id, sdk.ZeroInt(), sdk.ZeroInt()))

	s.deposit(s.addr(1), pool.Id, utils.ParseCoins("100_000_000denom1, 100_000_000denom2"), true)
	s.nextBlock()

	amount1, amount2, amount3 := sdk.NewInt(100_000), sdk.NewInt(500_000), sdk.NewInt(100_000_000)
	s.farm(pool.Id, s.addr(1), sdk.NewCoin(pool.PoolCoinDenom, amount1), true)
	s.nextBlock()

	s.farm(pool.Id, s.addr(1), sdk.NewCoin(pool.PoolCoinDenom, amount2), true)
	s.nextBlock()

	s.farm(pool.Id, s.addr(1), sdk.NewCoin(pool.PoolCoinDenom, amount3), true)
	s.nextBlock()

	for _, tc := range []struct {
		name      string
		req       *types.QueryQueuedFarmingsByFarmerRequest
		expectErr bool
		postRun   func(*types.QueryQueuedFarmingsByFarmerResponse)
	}{
		{
			"nil request",
			nil,
			true,
			nil,
		},
		{
			"query by empty farmer address",
			&types.QueryQueuedFarmingsByFarmerRequest{
				PoolId: pool.Id,
			},
			true,
			func(resp *types.QueryQueuedFarmingsByFarmerResponse) {},
		},
		{
			"query by invalid pool id",
			&types.QueryQueuedFarmingsByFarmerRequest{
				PoolId:        5,
				FarmerAddress: s.addr(1).String(),
			},
			false,
			func(resp *types.QueryQueuedFarmingsByFarmerResponse) {
				s.Require().Len(resp.QueuedFarmings, 0)
			},
		},
		{
			"query all queued farmings",
			&types.QueryQueuedFarmingsByFarmerRequest{
				PoolId:        pool.Id,
				FarmerAddress: s.addr(1).String(),
			},
			false,
			func(resp *types.QueryQueuedFarmingsByFarmerResponse) {
				s.Require().Len(resp.QueuedFarmings, 3)
				s.Require().Equal(amount1, resp.QueuedFarmings[0].Amount)
				s.Require().Equal(amount2, resp.QueuedFarmings[1].Amount)
				s.Require().Equal(amount3, resp.QueuedFarmings[2].Amount)
			},
		},
	} {
		s.Run(tc.name, func() {
			resp, err := s.querier.QueuedFarmingsByFarmer(sdk.WrapSDKContext(s.ctx), tc.req)
			if tc.expectErr {
				s.Require().Error(err)
			} else {
				s.Require().NoError(err)
				tc.postRun(resp)
			}
		})
	}
}

func (s *KeeperTestSuite) TestGRPCRewardsAuctions() {
	pair := s.createPair(s.addr(0), "denom1", "denom2", true)
	pool := s.createPool(s.addr(0), pair.Id, utils.ParseCoins("100_000_000denom1, 100_000_000denom2"), true)
	s.createLiquidFarm(types.NewLiquidFarm(pool.Id, sdk.ZeroInt(), sdk.ZeroInt()))

	s.deposit(s.addr(1), pool.Id, utils.ParseCoins("100_000_000denom1, 100_000_000denom2"), true)
	s.deposit(s.addr(2), pool.Id, utils.ParseCoins("100_000_000denom1, 100_000_000denom2"), true)
	s.deposit(s.addr(3), pool.Id, utils.ParseCoins("100_000_000denom1, 100_000_000denom2"), true)
	s.nextBlock()

	s.farm(pool.Id, s.addr(1), utils.ParseCoin("100_000pool1"), true)
	s.farm(pool.Id, s.addr(2), utils.ParseCoin("300_000pool1"), true)
	s.farm(pool.Id, s.addr(3), utils.ParseCoin("500_000pool1"), true)
	s.nextBlock()

	s.advanceEpochDays() // trigger AllocateRewards hook to create the first rewards auction

	s.placeBid(pool.Id, s.addr(5), utils.ParseCoin("100_000pool1"), true)
	s.placeBid(pool.Id, s.addr(6), utils.ParseCoin("110_000pool1"), true)
	s.placeBid(pool.Id, s.addr(7), utils.ParseCoin("150_000pool1"), true)

	s.advanceEpochDays() // finish the first auction and create the second rewards auction

	s.placeBid(pool.Id, s.addr(5), utils.ParseCoin("100_000pool1"), true)
	s.placeBid(pool.Id, s.addr(6), utils.ParseCoin("110_000pool1"), true)
	s.placeBid(pool.Id, s.addr(7), utils.ParseCoin("150_000pool1"), true)

	s.advanceEpochDays()

	for _, tc := range []struct {
		name      string
		req       *types.QueryRewardsAuctionsRequest
		expectErr bool
		postRun   func(*types.QueryRewardsAuctionsResponse)
	}{
		{
			"nil request",
			nil,
			true,
			nil,
		},
		{
			"query by invalid pool id",
			&types.QueryRewardsAuctionsRequest{
				PoolId: 0,
			},
			true,
			func(resp *types.QueryRewardsAuctionsResponse) {},
		},
		{
			"query by invalid pool id",
			&types.QueryRewardsAuctionsRequest{
				PoolId: 10,
			},
			false,
			func(resp *types.QueryRewardsAuctionsResponse) {
				s.Require().Len(resp.RewardAuctions, 0)
			},
		},
		{
			"query by pool id",
			&types.QueryRewardsAuctionsRequest{
				PoolId: pool.Id,
			},
			false,
			func(resp *types.QueryRewardsAuctionsResponse) {
				s.Require().Len(resp.RewardAuctions, 3)
			},
		},
	} {
		s.Run(tc.name, func() {
			resp, err := s.querier.RewardsAuctions(sdk.WrapSDKContext(s.ctx), tc.req)
			if tc.expectErr {
				s.Require().Error(err)
			} else {
				s.Require().NoError(err)
				tc.postRun(resp)
			}
		})
	}
}

func (s *KeeperTestSuite) TestGRPCRewardsAuction() {
	pair := s.createPair(s.addr(0), "denom1", "denom2", true)
	pool := s.createPool(s.addr(0), pair.Id, utils.ParseCoins("100_000_000denom1, 100_000_000denom2"), true)
	s.createLiquidFarm(types.NewLiquidFarm(pool.Id, sdk.ZeroInt(), sdk.ZeroInt()))

	s.deposit(s.addr(1), pool.Id, utils.ParseCoins("100_000_000denom1, 100_000_000denom2"), true)
	s.deposit(s.addr(2), pool.Id, utils.ParseCoins("100_000_000denom1, 100_000_000denom2"), true)
	s.deposit(s.addr(3), pool.Id, utils.ParseCoins("100_000_000denom1, 100_000_000denom2"), true)
	s.nextBlock()

	s.farm(pool.Id, s.addr(1), utils.ParseCoin("100_000pool1"), true)
	s.farm(pool.Id, s.addr(2), utils.ParseCoin("300_000pool1"), true)
	s.farm(pool.Id, s.addr(3), utils.ParseCoin("500_000pool1"), true)
	s.nextBlock()

	s.advanceEpochDays() // trigger AllocateRewards hook to create the first rewards auction

	s.placeBid(pool.Id, s.addr(5), utils.ParseCoin("100_000pool1"), true)
	s.placeBid(pool.Id, s.addr(6), utils.ParseCoin("110_000pool1"), true)
	s.placeBid(pool.Id, s.addr(7), utils.ParseCoin("150_000pool1"), true)

	s.advanceEpochDays() // finish the first auction and create the second rewards auction

	for _, tc := range []struct {
		name      string
		req       *types.QueryRewardsAuctionRequest
		expectErr bool
		postRun   func(*types.QueryRewardsAuctionResponse)
	}{
		{
			"nil request",
			nil,
			true,
			nil,
		},
		{
			"query by invalid pool id",
			&types.QueryRewardsAuctionRequest{
				PoolId:    10,
				AuctionId: 1,
			},
			true,
			func(resp *types.QueryRewardsAuctionResponse) {},
		},
		{
			"query by invalid auction id",
			&types.QueryRewardsAuctionRequest{
				PoolId:    1,
				AuctionId: 10,
			},
			true,
			func(resp *types.QueryRewardsAuctionResponse) {},
		},
		{
			"query finished auction",
			&types.QueryRewardsAuctionRequest{
				PoolId:    pool.Id,
				AuctionId: 1,
			},
			false,
			func(resp *types.QueryRewardsAuctionResponse) {
				s.Require().Equal(pool.PoolCoinDenom, resp.RewardAuction.BiddingCoinDenom)
				s.Require().Equal(types.PayingReserveAddress(pool.Id), resp.RewardAuction.GetPayingReserveAddress())
				s.Require().Equal(types.AuctionStatusFinished, resp.RewardAuction.Status)
			},
		},
		{
			"query started auction",
			&types.QueryRewardsAuctionRequest{
				PoolId:    pool.Id,
				AuctionId: 2,
			},
			false,
			func(resp *types.QueryRewardsAuctionResponse) {
				s.Require().Equal(pool.PoolCoinDenom, resp.RewardAuction.BiddingCoinDenom)
				s.Require().Equal(types.PayingReserveAddress(pool.Id), resp.RewardAuction.GetPayingReserveAddress())
				s.Require().Equal(types.AuctionStatusStarted, resp.RewardAuction.Status)
			},
		},
	} {
		s.Run(tc.name, func() {
			resp, err := s.querier.RewardsAuction(sdk.WrapSDKContext(s.ctx), tc.req)
			if tc.expectErr {
				s.Require().Error(err)
			} else {
				s.Require().NoError(err)
				tc.postRun(resp)
			}
		})
	}
}

func (s *KeeperTestSuite) TestGRPCBids() {
	pair := s.createPair(s.addr(0), "denom1", "denom2", true)
	pool := s.createPool(s.addr(0), pair.Id, utils.ParseCoins("100_000_000denom1, 100_000_000denom2"), true)
	s.createLiquidFarm(types.NewLiquidFarm(pool.Id, sdk.ZeroInt(), sdk.ZeroInt()))

	s.deposit(s.addr(1), pool.Id, utils.ParseCoins("100_000_000denom1, 100_000_000denom2"), true)
	s.deposit(s.addr(2), pool.Id, utils.ParseCoins("100_000_000denom1, 100_000_000denom2"), true)
	s.deposit(s.addr(3), pool.Id, utils.ParseCoins("100_000_000denom1, 100_000_000denom2"), true)
	s.nextBlock()

	s.farm(pool.Id, s.addr(1), utils.ParseCoin("100_000pool1"), true)
	s.farm(pool.Id, s.addr(2), utils.ParseCoin("300_000pool1"), true)
	s.farm(pool.Id, s.addr(3), utils.ParseCoin("500_000pool1"), true)
	s.nextBlock()

	s.advanceEpochDays() // trigger AllocateRewards hook to create the first rewards auction

	s.placeBid(pool.Id, s.addr(5), utils.ParseCoin("100_000pool1"), true)
	s.placeBid(pool.Id, s.addr(6), utils.ParseCoin("110_000pool1"), true)
	s.placeBid(pool.Id, s.addr(7), utils.ParseCoin("150_000pool1"), true)

	for _, tc := range []struct {
		name      string
		req       *types.QueryBidsRequest
		expectErr bool
		postRun   func(*types.QueryBidsResponse)
	}{
		{
			"nil request",
			nil,
			true,
			nil,
		},
		{
			"query by invalid pool id",
			&types.QueryBidsRequest{
				PoolId: 0,
			},
			true,
			func(resp *types.QueryBidsResponse) {},
		},
		{
			"query by pool id",
			&types.QueryBidsRequest{
				PoolId: pool.Id,
			},
			false,
			func(resp *types.QueryBidsResponse) {
				s.Require().Len(resp.Bids, 3)
			},
		},
	} {
		s.Run(tc.name, func() {
			resp, err := s.querier.Bids(sdk.WrapSDKContext(s.ctx), tc.req)
			if tc.expectErr {
				s.Require().Error(err)
			} else {
				s.Require().NoError(err)
				tc.postRun(resp)
			}
		})
	}
}
