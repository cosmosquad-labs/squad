package keeper

import (
	"fmt"
	"time"

	sdk "github.com/cosmos/cosmos-sdk/types"

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
	// TODO: not implemented yet
	// Check for minimum deposit amount
	// Check if the bidder has enough amount of coin to place a bid
	// Check if the auction exists
	// Check if the bid amount is greater than the currently winning bid amount

	// Get the last auction id

	// params := k.GetParams(ctx)
	// poolId := uint64(0)
	// minBidAmt := sdk.ZeroInt()
	// for _, lf := range params.LiquidFarms {
	// 	if lf.PoolId == msg.PoolId {
	// 		poolId = lf.PoolId
	// 		minBidAmt = lf.MinimumBidAmount
	// 		break
	// 	}
	// }

	// // Check if the amount exceeds minimum bid amount
	// if msg.Amount.Amount.LT(minBidAmt) {
	// 	return sdkerrors.Wrapf(types.ErrInsufficientBidAmount, "%s is smaller than %s", msg.Amount.Amount, minBidAmt)
	// }

	// Check if the bidder has sufficient balance to place a bid
	// poolCoinBalance := k.bankKeeper.SpendableCoins(ctx, bidder).AmountOf(pool.PoolCoinDenom)
	// if poolCoinBalance.LT(farmingCoin.Amount) {
	// return sdkerrors.Wrapf(types.ErrInsufficientFarmingCoinAmount, "%s is smaller than %s", poolCoinBalance, minDepositAmt)
	// }

	// k.GetRewardsAuction(ctx, poolId uint64, auctionId uint64)

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

func (k Keeper) CreateRewardsAuction(ctx sdk.Context) error {
	// Harvest and create rewards auction for every liquid farm
	for _, lf := range k.GetParams(ctx).LiquidFarms { // looping LiquidFarms?
		reserveAddr := types.LiquidFarmReserveAddress(lf.PoolId)
		stakingCoinDenom := liquiditytypes.PoolCoinDenom(lf.PoolId)
		// if err := k.farmingKeeper.Harvest(ctx, reserveAddr, []string{stakingCoinDenom}); err != nil {
		// 	return err
		// }

		// TODO: staking coin is already staked, so i believe the reserve account only has rewards
		// but verify this with test code
		spendable := k.bankKeeper.SpendableCoins(ctx, reserveAddr)
		fmt.Println("spendable: ", spendable)

		startTime := ctx.BlockTime()
		currentEpochDays := k.farmingKeeper.GetCurrentEpochDays(ctx)
		endTime := startTime.Add(time.Duration(currentEpochDays) * farmingtypes.Day)

		auction := types.NewRewardsAuction(
			k.GetNextAuctionIdWithUpdate(ctx),
			lf.PoolId,
			spendable, // Remove this field
			stakingCoinDenom,
			startTime,
			endTime,
		)
		k.SetRewardsAuction(ctx, auction)
	}
	return nil
}

func (k Keeper) DistributeRewards(ctx sdk.Context) {

}
