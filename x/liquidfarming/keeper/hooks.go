package keeper

import (
	"fmt"
	"time"

	sdk "github.com/cosmos/cosmos-sdk/types"

	farmingtypes "github.com/cosmosquad-labs/squad/v2/x/farming/types"
	"github.com/cosmosquad-labs/squad/v2/x/liquidfarming/types"
	liquiditytypes "github.com/cosmosquad-labs/squad/v2/x/liquidity/types"
)

// Wrapper struct
type Hooks struct {
	k Keeper
}

var _ farmingtypes.FarmingHooks = Hooks{}

// Create new liquidfarming hooks
func (k Keeper) Hooks() Hooks { return Hooks{k} }

func (h Hooks) AfterStaked(ctx sdk.Context, stakingAcc sdk.AccAddress, stakingCoinDenom string, newStakingAmt sdk.Int) {
	mintingAmt := sdk.ZeroInt()
	h.k.IterateMatureQueuedFarmings(ctx, ctx.BlockTime(), func(endTime time.Time, farmingCoinDenom string, farmerAcc sdk.AccAddress, queuedFarming types.QueuedFarming) (stop bool) {
		poolId := queuedFarming.PoolId
		reserveAddr := types.LiquidFarmReserveAddress(poolId)
		poolCoinDenom := liquiditytypes.PoolCoinDenom(poolId)
		if stakingAcc.Equals(reserveAddr) && poolCoinDenom == farmingCoinDenom {
			// Consider a case when multiple farmers farm their respective amount in the same block and
			// the same reserve account stakes the total amount in the farming module, which results to newStakingAmt.
			remainingAmt := newStakingAmt
			if newStakingAmt.GTE(queuedFarming.Amount) {
				remainingAmt = remainingAmt.Sub(queuedFarming.Amount)
				lfCoinTotalSupply := h.k.bankKeeper.GetSupply(ctx, types.LFCoinDenom(poolId)).Amount
				lpCoinTotalStaked := h.k.farmingKeeper.GetAllStakedCoinsByFarmer(ctx, reserveAddr).AmountOf(poolCoinDenom)
				if lfCoinTotalSupply.IsZero() || lpCoinTotalStaked.IsZero() {
					// Initial minting amount
					mintingAmt = queuedFarming.Amount
				} else {
					// Minting Amount = TotalSupplyLFCoin * LPCoinFarmed / TotalStakedLPCoin
					mintingAmt = lfCoinTotalSupply.Mul(queuedFarming.Amount).Quo(lpCoinTotalStaked)
				}
				mintingCoins := sdk.NewCoins(sdk.NewCoin(types.LFCoinDenom(poolId), mintingAmt))

				if err := h.k.bankKeeper.MintCoins(ctx, types.ModuleName, mintingCoins); err != nil {
					panic(err)
				}
				if err := h.k.bankKeeper.SendCoinsFromModuleToAccount(ctx, types.ModuleName, farmerAcc, mintingCoins); err != nil {
					panic(err)
				}

				// Delete queued farming after the module successfully mints LFCoin and send it to the farmer
				h.k.DeleteQueuedFarming(ctx, endTime, farmingCoinDenom, farmerAcc)

				if remainingAmt.IsZero() {
					return true
				}
			}
		}
		return false
	})
}

func (h Hooks) AfterAllocateRewards(ctx sdk.Context) {
	for _, lf := range h.k.GetParams(ctx).LiquidFarms {
		auctionId := h.k.GetLastRewardsAuctionId(ctx, lf.PoolId)
		auction, found := h.k.GetRewardsAuction(ctx, lf.PoolId, auctionId)
		if !found {
			h.k.CreateRewardsAuction(ctx, lf.PoolId) // initial
			continue
		}

		// Sanity checks
		if auction.EndTime.Before(ctx.BlockTime()) && auction.Status != types.AuctionStatusStarted {
			continue
		}

		winningBid, found := h.k.GetWinningBid(ctx, auction.PoolId, auction.Id)
		if !found {
			h.k.DeleteAndCreateRewardsAuction(ctx, auction)
			continue
		}

		reserveAddr := types.LiquidFarmReserveAddress(winningBid.PoolId)
		poolCoinDenom := liquiditytypes.PoolCoinDenom(winningBid.PoolId)

		stakedCoins := h.k.farmingKeeper.GetAllStakedCoinsByFarmer(ctx, reserveAddr)
		rewards := h.k.farmingKeeper.Rewards(ctx, reserveAddr, poolCoinDenom)
		unharvestedRewards, found := h.k.farmingKeeper.GetUnharvestedRewards(ctx, reserveAddr, poolCoinDenom)
		if !found {
			fmt.Println("NOT FOUND")
		}

		fmt.Println("UnharvestedRewards: ", unharvestedRewards)
		fmt.Println("Balances: ", h.k.bankKeeper.GetAllBalances(ctx, reserveAddr))
		fmt.Println("Balances: ", h.k.bankKeeper.GetAllBalances(ctx, farmingtypes.RewardsReserveAcc))

		fmt.Println("winningBid: ", winningBid)
		fmt.Println("stakedCoins: ", stakedCoins)
		fmt.Println("rewards: ", rewards)

		// Harvest
		if err := h.k.farmingKeeper.Harvest(ctx, reserveAddr, []string{poolCoinDenom}); err != nil {
			panic(err)
		}

		fmt.Println("Balances (Harvest): ", h.k.bankKeeper.GetAllBalances(ctx, reserveAddr))

		// Distribute
		if err := h.k.bankKeeper.SendCoins(ctx, reserveAddr, winningBid.GetBidder(), rewards); err != nil {
			panic(err)
		}

		// Refund existing bids by the pool id
		for _, bid := range h.k.GetBidsByPoolId(ctx, winningBid.PoolId) {
			if err := h.k.bankKeeper.SendCoins(ctx, auction.GetPayingReserveAddress(), bid.GetBidder(), sdk.NewCoins(bid.Amount)); err != nil {
				panic(err)
			}
			h.k.DeleteBid(ctx, bid)
		}

		h.k.DeleteAndCreateRewardsAuction(ctx, auction)
	}
}

func (k Keeper) DeleteAndCreateRewardsAuction(ctx sdk.Context, auction types.RewardsAuction) {
	k.DeleteRewardsAuction(ctx, auction)
	k.CreateRewardsAuction(ctx, auction.PoolId)
}
