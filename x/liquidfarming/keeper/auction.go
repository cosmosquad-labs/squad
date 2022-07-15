package keeper

import (
	"time"

	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"

	farmingtypes "github.com/cosmosquad-labs/squad/v2/x/farming/types"
	"github.com/cosmosquad-labs/squad/v2/x/liquidfarming/types"
	liquiditytypes "github.com/cosmosquad-labs/squad/v2/x/liquidity/types"
)

// GetNextAuctionIdWithUpdate increments rewards auction id by one and store it.
func (k Keeper) GetNextAuctionIdWithUpdate(ctx sdk.Context) uint64 {
	id := k.GetRewardsAuctionId(ctx) + 1
	k.SetRewardsAuctionId(ctx, id)
	return id
}

func (k Keeper) PlaceBid(ctx sdk.Context, msg *types.MsgPlaceBid) error {
	auctionId := k.GetRewardsAuctionId(ctx)
	auction, found := k.GetRewardsAuction(ctx, msg.PoolId, auctionId)
	if !found {
		return sdkerrors.Wrapf(sdkerrors.ErrNotFound, "auction %d not found", auctionId)
	}

	if auction.Status != types.AuctionStatusStarted {
		return sdkerrors.Wrapf(types.ErrInvalidAuctionStatus, "auction status is not %s", auction.Status.String())
	}

	liquidFarm, found := k.GetLiquidFarm(ctx, msg.PoolId)
	if !found {
		return sdkerrors.Wrapf(sdkerrors.ErrNotFound, "liquid farm with pool %d not found", msg.PoolId)
	}

	balance := k.bankKeeper.SpendableCoins(ctx, msg.GetBidder()).AmountOf(msg.BiddingCoin.Denom)
	if balance.LT(msg.BiddingCoin.Amount) {
		return sdkerrors.Wrapf(types.ErrInsufficientFarmingCoinAmount, "%s is smaller than %s", balance, msg.BiddingCoin.Amount)
	}

	if msg.BiddingCoin.Amount.LT(liquidFarm.MinimumBidAmount) {
		return sdkerrors.Wrapf(types.ErrInsufficientBidAmount, "%s is smaller than %s", msg.BiddingCoin.Amount, liquidFarm.MinimumBidAmount)
	}

	bidId := k.GetBidId(ctx, auctionId)
	bid, found := k.GetBid(ctx, auctionId, bidId)
	if found {
		if bid.Amount.IsGTE(msg.BiddingCoin) { // bidding amount must be greater than the winning bid amount
			return sdkerrors.Wrapf(sdkerrors.ErrInvalidRequest, "%s is smaller than winning bid amount %s", msg.BiddingCoin.Amount, bid.Amount)
		}
	}

	k.SetBid(ctx, types.Bid{
		Id:        k.GetNextAuctionIdWithUpdate(ctx),
		AuctionId: auctionId,
		Bidder:    msg.Bidder,
		Amount:    msg.BiddingCoin,
	})

	return nil
}

func (k Keeper) RefundBid(ctx sdk.Context, msg *types.MsgRefundBid) error {
	// TODO: not implemented yet
	// Winning bid can't be refunded

	return nil
}

func (k Keeper) TerminateRewardsAuction(ctx sdk.Context) error {
	/*
		Loop through all existing RewardsAuctions
		Get winning bidder
		Harvest and distribute rewards to the winner
		Set WinningBidId in the RewardsAuction
		Stake bid amounts (auto-compound)
	*/
	// for _, auction := range k.GetRewardsAuctions(ctx) {

	// }

	return nil
}

// CreateRewardsAuction ...
func (k Keeper) CreateRewardsAuction(ctx sdk.Context) error {
	currentEpochDays := k.farmingKeeper.GetCurrentEpochDays(ctx)

	for _, lf := range k.GetParams(ctx).LiquidFarms { // looping LiquidFarms?
		startTime := ctx.BlockTime()
		endTime := startTime.Add(time.Duration(currentEpochDays) * farmingtypes.Day)

		auction := types.NewRewardsAuction(
			k.GetNextAuctionIdWithUpdate(ctx),
			lf.PoolId,
			liquiditytypes.PoolCoinDenom(lf.PoolId),
			startTime,
			endTime,
		)
		k.SetRewardsAuction(ctx, auction)
	}
	return nil
}

func (k Keeper) DistributeRewards(ctx sdk.Context) {

}
