package keeper

import (
	"strconv"
	"time"

	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
	banktypes "github.com/cosmos/cosmos-sdk/x/bank/types"

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
		return types.Bid{}, sdkerrors.Wrapf(sdkerrors.ErrInvalidRequest, "bid already exists by %s; refund bid is required to place new bid", msg.Bidder)
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

	k.SetRewardsAuction(ctx, types.NewRewardsAuction(
		nextAuctionId,
		poolId,
		poolCoinDenom,
		startTime,
		endTime,
	))
}

func (k Keeper) TerminateRewardsAuction(ctx sdk.Context, auction types.RewardsAuction) {
	winningBid, found := k.GetWinningBid(ctx, auction.PoolId, auction.Id)
	if !found {
		k.CreateRewardsAuction(ctx, auction.PoolId)
		return
	}

	poolId := winningBid.PoolId
	auctionPayingReserveAddr := auction.GetPayingReserveAddress()
	liquidFarmReserveAddr := types.LiquidFarmReserveAddress(poolId)
	poolCoinDenom := liquiditytypes.PoolCoinDenom(poolId)
	rewards := k.farmingKeeper.Rewards(ctx, liquidFarmReserveAddr, poolCoinDenom)

	// Harvest farming rewards and send them to the auction winner
	if err := k.farmingKeeper.Harvest(ctx, liquidFarmReserveAddr, []string{poolCoinDenom}); err != nil {
		panic(err)
	}
	if err := k.bankKeeper.SendCoins(ctx, liquidFarmReserveAddr, winningBid.GetBidder(), rewards); err != nil {
		panic(err)
	}

	// Refund all at once and delete all bids
	inputs := []banktypes.Input{}
	outputs := []banktypes.Output{}
	for _, bid := range k.GetBidsByPoolId(ctx, poolId) {
		if bid.Bidder == winningBid.Bidder {
			continue
		}

		inputs = append(inputs, banktypes.NewInput(auctionPayingReserveAddr, sdk.NewCoins(bid.Amount)))
		outputs = append(outputs, banktypes.NewOutput(bid.GetBidder(), sdk.NewCoins(bid.Amount)))

		k.DeleteBid(ctx, bid)
	}
	if err := k.bankKeeper.InputOutputCoins(ctx, inputs, outputs); err != nil {
		panic(err)
	}

	// Reserve winning bid amount
	if err := k.bankKeeper.SendCoins(ctx, auctionPayingReserveAddr, liquidFarmReserveAddr, sdk.NewCoins(winningBid.Amount)); err != nil {
		panic(err)
	}

	// Stake with the winning bid amount for auto compounding
	if err := k.farmingKeeper.Stake(ctx, liquidFarmReserveAddr, sdk.NewCoins(winningBid.Amount)); err != nil {
		panic(err)
	}

	// Update auction fields
	auction.SetWinner(winningBid.Bidder)
	auction.SetRewards(rewards)
	auction.SetStatus(types.AuctionStatusFinished)
	k.SetRewardsAuction(ctx, auction)
}
