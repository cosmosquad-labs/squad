package keeper

import (
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
				lfCoinTotalSupply := h.k.bankKeeper.GetSupply(ctx, types.LiquidFarmCoinDenom(poolId)).Amount
				lpCoinTotalStaked := h.k.farmingKeeper.GetAllStakedCoinsByFarmer(ctx, reserveAddr).AmountOf(poolCoinDenom)
				if lfCoinTotalSupply.IsZero() || lpCoinTotalStaked.IsZero() {
					// Initial minting amount
					mintingAmt = queuedFarming.Amount
				} else {
					// Minting Amount = TotalSupplyLFCoin * LPCoinFarmed / TotalStakedLPCoin
					mintingAmt = lfCoinTotalSupply.Mul(queuedFarming.Amount).Quo(lpCoinTotalStaked)
				}
				mintingCoins := sdk.NewCoins(sdk.NewCoin(types.LiquidFarmCoinDenom(poolId), mintingAmt))

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
		poolId := lf.PoolId
		auctionId := h.k.GetLastRewardsAuctionId(ctx, poolId)

		auction, found := h.k.GetRewardsAuction(ctx, poolId, auctionId)
		if !found {
			h.k.CreateRewardsAuction(ctx, poolId)
			continue
		}

		h.k.TerminateRewardsAuction(ctx, auction)
		h.k.CreateRewardsAuction(ctx, poolId)
	}
}
