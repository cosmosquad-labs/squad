package keeper_test

import (
	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"

	utils "github.com/cosmosquad-labs/squad/v2/types"
	"github.com/cosmosquad-labs/squad/v2/x/liquidfarming/types"

	_ "github.com/stretchr/testify/suite"
)

func (s *KeeperTestSuite) TestPlaceBid() {
	pair := s.createPair(s.addr(0), "denom1", "denom2", true)
	pool := s.createPool(s.addr(0), pair.Id, utils.ParseCoins("100_000_000denom1, 100_000_000denom2"), true)

	_, err := s.keeper.PlaceBid(s.ctx, &types.MsgPlaceBid{
		PoolId:      pool.Id,
		Bidder:      s.addr(0).String(),
		BiddingCoin: sdk.NewInt64Coin(pool.PoolCoinDenom, 100_000_000)},
	)
	s.Require().ErrorIs(err, sdkerrors.ErrNotFound)

	s.createLiquidFarm(types.NewLiquidFarm(pool.Id, sdk.NewInt(10_000_000), sdk.NewInt(10_000_000)))
	s.createRewardsAuction(pool.Id)

	for _, tc := range []struct {
		name        string
		msg         *types.MsgPlaceBid
		postRun     func(ctx sdk.Context, bid types.Bid)
		expectedErr string
	}{
		{
			"happy case",
			types.NewMsgPlaceBid(
				pool.Id,
				s.addr(0).String(),
				sdk.NewInt64Coin(pool.PoolCoinDenom, 1_000_000_000),
			),
			func(ctx sdk.Context, bid types.Bid) {
				s.Require().Equal(pool.Id, bid.PoolId)
				s.Require().Equal(s.addr(0), bid.GetBidder())
				s.Require().Equal(sdk.NewInt(1_000_000_000), bid.Amount.Amount)
			},
			"",
		},
		{
			"insufficient balance",
			types.NewMsgPlaceBid(
				pool.Id,
				s.addr(0).String(),
				sdk.NewInt64Coin(pool.PoolCoinDenom, 5_000_000_000_000_000),
			),
			nil,
			"1000000000000 is smaller than 5000000000000000: insufficient funds",
		},
		{
			"insufficient bidding amount",
			types.NewMsgPlaceBid(
				pool.Id,
				s.addr(0).String(),
				sdk.NewInt64Coin(pool.PoolCoinDenom, 1),
			),
			nil,
			"1 is smaller than 10000000: insufficient bid amount",
		},
	} {
		s.Run(tc.name, func() {
			s.Require().NoError(tc.msg.ValidateBasic())
			cacheCtx, _ := s.ctx.CacheContext()
			bid, err := s.keeper.PlaceBid(cacheCtx, tc.msg)
			if tc.expectedErr == "" {
				s.Require().NoError(err)
				bid, found := s.keeper.GetBid(cacheCtx, bid.PoolId, bid.GetBidder())
				s.Require().True(found)
				tc.postRun(cacheCtx, bid)

				auctionId := s.keeper.GetLastRewardsAuctionId(cacheCtx, bid.PoolId)
				_, found = s.keeper.GetWinningBid(cacheCtx, bid.PoolId, auctionId)
				s.Require().True(found)
			} else {
				s.Require().EqualError(err, tc.expectedErr)
			}
		})
	}
}

func (s *KeeperTestSuite) TestPlaceBid_EdgeCases() {
	pair := s.createPair(s.addr(0), "denom1", "denom2", true)
	pool := s.createPool(s.addr(0), pair.Id, utils.ParseCoins("100_000_000denom1, 100_000_000denom2"), true)
	s.createLiquidFarm(types.NewLiquidFarm(pool.Id, sdk.NewInt(10_000_000), sdk.NewInt(10_000_000)))
	s.createRewardsAuction(pool.Id)

	// Place a bid successfully
	_, err := s.keeper.PlaceBid(s.ctx, &types.MsgPlaceBid{
		PoolId:      pool.Id,
		Bidder:      s.addr(0).String(),
		BiddingCoin: sdk.NewInt64Coin(pool.PoolCoinDenom, 500_000_000)},
	)
	s.Require().NoError(err)

	// Error: bid already exists
	s.fundAddr(s.addr(0), utils.ParseCoins("500_000_000pool1"))
	_, err = s.keeper.PlaceBid(s.ctx, &types.MsgPlaceBid{
		PoolId:      pool.Id,
		Bidder:      s.addr(0).String(),
		BiddingCoin: sdk.NewInt64Coin(pool.PoolCoinDenom, 500_000_000)},
	)
	s.Require().ErrorIs(err, sdkerrors.ErrInvalidRequest)

	// Error: place a bid with smaller than the winning bid amount with different bidder
	s.fundAddr(s.addr(1), utils.ParseCoins("1_000_000_000pool1"))
	_, err = s.keeper.PlaceBid(s.ctx, &types.MsgPlaceBid{
		PoolId:      pool.Id,
		Bidder:      s.addr(1).String(),
		BiddingCoin: sdk.NewInt64Coin(pool.PoolCoinDenom, 10_000_000)},
	)
	s.Require().ErrorIs(err, sdkerrors.ErrInvalidRequest)

	// Place a bid with more bidding amount with different bidder
	newBiddingAmt := sdk.NewInt(1_000_000_000)
	_, err = s.keeper.PlaceBid(s.ctx, &types.MsgPlaceBid{
		PoolId:      pool.Id,
		Bidder:      s.addr(1).String(),
		BiddingCoin: sdk.NewCoin(pool.PoolCoinDenom, newBiddingAmt)},
	)
	s.Require().NoError(err)

	auctionId := s.keeper.GetLastRewardsAuctionId(s.ctx, pool.Id)
	winningBid, found := s.keeper.GetWinningBid(s.ctx, pool.Id, auctionId)
	s.Require().True(found)
	s.Require().Equal(newBiddingAmt, winningBid.Amount.Amount)
}

func (s *KeeperTestSuite) TestRefundBid() {
	pair := s.createPair(s.addr(0), "denom1", "denom2", true)
	pool := s.createPool(s.addr(0), pair.Id, utils.ParseCoins("100_000_000denom1, 100_000_000denom2"), true)
	s.createLiquidFarm(types.NewLiquidFarm(pool.Id, sdk.NewInt(10_000_000), sdk.NewInt(10_000_000)))
	s.createRewardsAuction(pool.Id)
	s.placeBid(pool.Id, s.addr(0), sdk.NewInt64Coin(pool.PoolCoinDenom, 500_000_000), true)
	s.placeBid(pool.Id, s.addr(1), sdk.NewInt64Coin(pool.PoolCoinDenom, 600_000_000), true)

	for _, tc := range []struct {
		name        string
		msg         *types.MsgRefundBid
		postRun     func(ctx sdk.Context, bid types.Bid)
		expectedErr string
	}{
		{
			"happy case",
			types.NewMsgRefundBid(
				pool.Id,
				s.addr(0).String(),
			),
			func(ctx sdk.Context, bid types.Bid) {
				s.Require().Equal(pool.Id, bid.PoolId)
				s.Require().Equal(s.addr(0), bid.GetBidder())
			},
			"",
		},
		{
			"refund winning bid",
			types.NewMsgRefundBid(
				pool.Id,
				s.addr(1).String(),
			),
			nil,
			"unable to refund winning bid: invalid request",
		},
		{
			"auction not found",
			types.NewMsgRefundBid(
				5,
				s.addr(1).String(),
			),
			nil,
			"auction corresponds to pool 5 not found: not found",
		},
	} {
		s.Run(tc.name, func() {
			s.Require().NoError(tc.msg.ValidateBasic())
			cacheCtx, _ := s.ctx.CacheContext()
			err := s.keeper.RefundBid(cacheCtx, tc.msg)
			if tc.expectedErr == "" {
				s.Require().NoError(err)

				_, found := s.keeper.GetBid(cacheCtx, tc.msg.PoolId, s.addr(0))
				s.Require().False(found)
			} else {
				s.Require().EqualError(err, tc.expectedErr)
			}
		})
	}
}

func (s *KeeperTestSuite) TestAfterAllocateRewards() {
	s.ctx = s.ctx.WithBlockTime(utils.ParseTime("2022-07-01T00:00:00Z"))
	pair := s.createPair(s.addr(0), "denom1", "denom2", true)
	pool := s.createPool(s.addr(0), pair.Id, utils.ParseCoins("100_000_000denom1, 100_000_000denom2"), true)
	s.createLiquidFarm(types.NewLiquidFarm(pool.Id, sdk.ZeroInt(), sdk.ZeroInt()))

	s.farm(pool.Id, s.addr(0), sdk.NewCoin(pool.PoolCoinDenom, sdk.NewInt(50_000_000)), true)
	s.advanceEpochDays()
	s.advanceEpochDays()

	auctionId := s.keeper.GetLastRewardsAuctionId(s.ctx, pool.Id)
	auction, found := s.keeper.GetRewardsAuction(s.ctx, pool.Id, auctionId)
	s.Require().True(found)

	s.placeBid(auction.PoolId, s.addr(0), sdk.NewInt64Coin(pool.PoolCoinDenom, 10_000_000), true)
	s.placeBid(auction.PoolId, s.addr(1), sdk.NewInt64Coin(pool.PoolCoinDenom, 20_000_000), true)
	s.placeBid(auction.PoolId, s.addr(3), sdk.NewInt64Coin(pool.PoolCoinDenom, 30_000_000), true)

	s.advanceEpochDays()

	// stakedCoins := s.app.FarmingKeeper.GetAllStakedCoinsByFarmer(s.ctx, types.LiquidFarmReserveAddress(pool.Id))
	// fmt.Println("StakedCoins: ", stakedCoins)

}
