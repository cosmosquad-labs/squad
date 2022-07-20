package keeper

import (
	"strconv"
	"time"

	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"

	farmingtypes "github.com/cosmosquad-labs/squad/v2/x/farming/types"
	"github.com/cosmosquad-labs/squad/v2/x/liquidfarming/types"
	liquiditytypes "github.com/cosmosquad-labs/squad/v2/x/liquidity/types"
)

func (k Keeper) PlaceBid(ctx sdk.Context, msg *types.MsgPlaceBid) (types.Bid, error) {
	auctionId := k.GetLastRewardsAuctionId(ctx, msg.PoolId)
	auction, found := k.GetRewardsAuction(ctx, msg.PoolId, auctionId)
	if !found {
		return types.Bid{}, sdkerrors.Wrapf(sdkerrors.ErrNotFound, "auction %d not found", auctionId)
	}

	liquidFarm, found := k.GetLiquidFarm(ctx, msg.PoolId)
	if !found {
		return types.Bid{}, sdkerrors.Wrapf(sdkerrors.ErrNotFound, "liquid farm with pool %d not found", msg.PoolId)
	}

	balance := k.bankKeeper.SpendableCoins(ctx, msg.GetBidder()).AmountOf(msg.BiddingCoin.Denom)
	if balance.LT(msg.BiddingCoin.Amount) {
		return types.Bid{}, sdkerrors.Wrapf(sdkerrors.ErrInsufficientFunds, "%s is smaller than %s", balance, msg.BiddingCoin.Amount)
	}

	if msg.BiddingCoin.Amount.LT(liquidFarm.MinimumBidAmount) {
		return types.Bid{}, sdkerrors.Wrapf(types.ErrInsufficientBidAmount, "%s is smaller than %s", msg.BiddingCoin.Amount, liquidFarm.MinimumBidAmount)
	}

	_, found = k.GetBid(ctx, auctionId, msg.GetBidder())
	if found {
		return types.Bid{}, sdkerrors.Wrapf(sdkerrors.ErrInvalidRequest, "bid by %s already exists", msg.Bidder)
	}

	winningBid, found := k.GetWinningBid(ctx, msg.PoolId, auctionId)
	if found {
		if winningBid.Amount.IsGTE(msg.BiddingCoin) {
			return types.Bid{}, sdkerrors.Wrapf(sdkerrors.ErrInvalidRequest, "%s is smaller than winning bid amount %s", msg.BiddingCoin.Amount, winningBid.Amount)
		}
	}

	if err := k.bankKeeper.SendCoins(ctx, msg.GetBidder(), auction.GetPayingReserveAddress(), sdk.NewCoins(msg.BiddingCoin)); err != nil {
		return types.Bid{}, err
	}

	bid := types.NewBid(
		msg.PoolId,
		msg.Bidder,
		msg.BiddingCoin,
	)
	k.SetBid(ctx, bid)
	k.SetWinningBid(ctx, bid, auction.Id)

	ctx.EventManager().EmitEvents(sdk.Events{
		sdk.NewEvent(
			types.EventTypePlaceBid,
			sdk.NewAttribute(types.AttributeKeyPoolId, strconv.FormatUint(msg.PoolId, 10)),
			sdk.NewAttribute(types.AttributeKeyAuctionId, strconv.FormatUint(auctionId, 10)),
			sdk.NewAttribute(types.AttributeKeyBidder, msg.Bidder),
			sdk.NewAttribute(types.AttributeKeyBiddingCoin, msg.BiddingCoin.String()),
		),
	})

	return bid, nil
}

func (k Keeper) RefundBid(ctx sdk.Context, msg *types.MsgRefundBid) error {
	auctionId := k.GetLastRewardsAuctionId(ctx, msg.PoolId)
	auction, found := k.GetRewardsAuction(ctx, msg.PoolId, auctionId)
	if !found {
		return sdkerrors.Wrapf(sdkerrors.ErrNotFound, "auction corresponds to pool %d not found", msg.PoolId)
	}

	winningBid, found := k.GetWinningBid(ctx, msg.PoolId, auctionId)
	if found {
		if winningBid.Bidder == msg.Bidder {
			return sdkerrors.Wrap(sdkerrors.ErrInvalidRequest, "unable to refund winning bid")
		}
	}

	bid, found := k.GetBid(ctx, msg.PoolId, msg.GetBidder())
	if !found {
		return sdkerrors.Wrap(sdkerrors.ErrNotFound, "bid not found")
	}

	if err := k.bankKeeper.SendCoins(ctx, auction.GetPayingReserveAddress(), msg.GetBidder(), sdk.NewCoins(bid.Amount)); err != nil {
		return err
	}

	k.DeleteBid(ctx, bid)

	ctx.EventManager().EmitEvents(sdk.Events{
		sdk.NewEvent(
			types.EventTypeRefundBid,
			sdk.NewAttribute(types.AttributeKeyPoolId, strconv.FormatUint(msg.PoolId, 10)),
			sdk.NewAttribute(types.AttributeKeyBidder, msg.Bidder),
		),
	})

	return nil
}

// getNextAuctionIdWithUpdate increments rewards auction id by one and store it.
func (k Keeper) getNextAuctionIdWithUpdate(ctx sdk.Context, poolId uint64) uint64 {
	id := k.GetLastRewardsAuctionId(ctx, poolId) + 1
	k.SetRewardsAuctionId(ctx, poolId, id)
	return id
}

func (k Keeper) CreateRewardsAuction(ctx sdk.Context, poolId uint64) {
	currentEpochDays := k.farmingKeeper.GetCurrentEpochDays(ctx)
	startTime := ctx.BlockTime()
	endTime := startTime.Add(time.Duration(currentEpochDays) * farmingtypes.Day)
	nextAuctionId := k.getNextAuctionIdWithUpdate(ctx, poolId)
	poolCoinDenom := liquiditytypes.PoolCoinDenom(poolId)

	auction := types.NewRewardsAuction(
		nextAuctionId,
		poolId,
		poolCoinDenom,
		startTime,
		endTime,
	)
	k.SetRewardsAuction(ctx, auction)
}
