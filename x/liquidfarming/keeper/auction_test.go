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
	s.createRewardsAuction()

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
	s.createRewardsAuction()

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
