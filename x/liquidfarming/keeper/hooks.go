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

// Hooks creates new hooks
func (k Keeper) Hooks() Hooks {
	return Hooks{k}
}

// AfterStaked hook is triggered in the farming module when an epoch is advanced and
// queued coins are moved into staked coins in ProcessQueuedCoins function.
//
// It iterates through all the queued farming objects that are mature at current time and
// proceeds to mint LFCoin in proportion to the staked amount for every queued farming and
// deletes the queued farming object when it is minted and sent to the farmer successfully.
func (h Hooks) AfterStaked(ctx sdk.Context, stakingAcc sdk.AccAddress, stakingCoinDenom string, newStakingAmt sdk.Int) {
	mintingAmt := sdk.ZeroInt()
	h.k.IterateMatureQueuedFarmings(ctx, ctx.BlockTime(), func(endTime time.Time, farmingCoinDenom string, farmerAcc sdk.AccAddress, queuedFarming types.QueuedFarming) (stop bool) {
		poolId := queuedFarming.PoolId
		reserveAddr := types.LiquidFarmReserveAddress(poolId)
		poolCoinDenom := liquiditytypes.PoolCoinDenom(poolId)

		// Mint LFCoin when staking address is equal to the LiquidFarm reserve address and pool coin denom is the same as farming coin denom.
		// The logic considers a case when multiple farmers farm their amounts in the same block that makes the LiquidFarm reserve account
		// to stake with all the farmers' amounts which ends up with the value of newStakingAmt.
		if stakingAcc.Equals(reserveAddr) && poolCoinDenom == farmingCoinDenom {
			remainingAmt := newStakingAmt
			if newStakingAmt.GTE(queuedFarming.Amount) {
				remainingAmt = remainingAmt.Sub(queuedFarming.Amount)
				lfCoinTotalSupply := h.k.bankKeeper.GetSupply(ctx, types.LiquidFarmCoinDenom(poolId)).Amount
				lpCoinTotalStaked := h.k.farmingKeeper.GetAllStakedCoinsByFarmer(ctx, reserveAddr).AmountOf(poolCoinDenom)

				// Use QueuedFarming amount for the initial minting when either
				// total supply or total staked amount is zero; otherwise use the following formula:
				// Minting Amount = TotalSupplyLFCoin * LPCoinFarmed / TotalStakedLPCoin
				if lfCoinTotalSupply.IsZero() || lpCoinTotalStaked.IsZero() {
					mintingAmt = queuedFarming.Amount
				} else {
					mintingAmt = lfCoinTotalSupply.Mul(queuedFarming.Amount).Quo(lpCoinTotalStaked)
				}
				mintingCoins := sdk.NewCoins(sdk.NewCoin(types.LiquidFarmCoinDenom(poolId), mintingAmt))

				if err := h.k.bankKeeper.MintCoins(ctx, types.ModuleName, mintingCoins); err != nil {
					panic(err)
				}
				if err := h.k.bankKeeper.SendCoinsFromModuleToAccount(ctx, types.ModuleName, farmerAcc, mintingCoins); err != nil {
					panic(err)
				}

				// Delete the queued farming when it is successfully minted and sent to the farmer
				h.k.DeleteQueuedFarming(ctx, endTime, farmingCoinDenom, farmerAcc)

				if remainingAmt.IsZero() {
					return true // stop the iteration
				}
			}
		}
		return false
	})
}

// AfterAllocateRewards hook is triggered in the farming module when an epoch is advanced and
// AllocateRewards is successfully executed till the end logic.
//
// It creates the first rewards auction if liquid farm doesn't have any auction before.
// If there is an ongoing rewards auction, finish the auction and create the next one.
func (h Hooks) AfterAllocateRewards(ctx sdk.Context) {
	for _, liquidFarm := range h.k.GetParams(ctx).LiquidFarms {
		poolId := liquidFarm.PoolId
		auctionId := h.k.GetLastRewardsAuctionId(ctx, poolId)

		// Create the first rewards auction if not exists; otherwise
		// finish the rewards auction and create the next one
		auction, found := h.k.GetRewardsAuction(ctx, poolId, auctionId)
		if !found {
			h.k.CreateRewardsAuction(ctx, poolId)
		} else {
			h.k.FinishRewardsAuction(ctx, auction)
		}
	}
}
