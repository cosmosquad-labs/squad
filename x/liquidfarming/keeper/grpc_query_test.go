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

				for _, lf := range resp.LiquidFarms {
					switch lf.PoolId {
					case 1:
						s.Require().Equal(minFarmAmt1, lf.MinimumFarmAmount)
						s.Require().Equal(minBidAmt1, lf.MinimumBidAmount)
					case 2:
						s.Require().Equal(minFarmAmt2, lf.MinimumFarmAmount)
						s.Require().Equal(minBidAmt2, lf.MinimumBidAmount)
					}

					reserveAcc, _ := sdk.AccAddressFromBech32(lf.LiquidFarmReserveAddress)
					poolCoinDenom := liquiditytypes.PoolCoinDenom(lf.PoolId)
					queuedAmt := s.app.FarmingKeeper.GetAllQueuedCoinsByFarmer(s.ctx, reserveAcc).AmountOf(poolCoinDenom)
					stakedAmt := s.app.FarmingKeeper.GetAllStakedCoinsByFarmer(s.ctx, reserveAcc).AmountOf(poolCoinDenom)
					s.Require().Equal(queuedAmt, lf.QueuedCoin.Amount)
					s.Require().Equal(stakedAmt, lf.StakedCoin.Amount)
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
