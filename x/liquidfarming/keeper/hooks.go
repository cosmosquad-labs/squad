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

func (h Hooks) AfterStaked(ctx sdk.Context, reserveAcc sdk.AccAddress, stakingCoinDenom string, newStakingAmt sdk.Int) {
	mintingAmt := sdk.ZeroInt()
	h.k.IterateMatureQueuedFarmings(ctx, ctx.BlockTime(), func(endTime time.Time, farmingCoinDenom string, farmerAcc sdk.AccAddress, queuedFarming types.QueuedFarming) (stop bool) {
		poolCoinDenom := liquiditytypes.PoolCoinDenom(queuedFarming.PoolId)
		if poolCoinDenom == farmingCoinDenom && newStakingAmt.Equal(queuedFarming.Amount) {
			lfCoinDenom := types.LFCoinDenom(queuedFarming.PoolId)
			lfCoinTotalSupply := h.k.bankKeeper.GetSupply(ctx, lfCoinDenom).Amount
			lpCoinTotalStaked, found := h.k.farmingKeeper.GetTotalStakings(ctx, stakingCoinDenom) // TODO: need verification
			if !found || lfCoinTotalSupply.IsZero() {
				// Initial minting amount
				mintingAmt = queuedFarming.Amount
			} else {
				// Minting Amount = TotalSupplyLFCoin / TotalStakedLPCoin * LPCoinFarmed
				mintingAmt = lfCoinTotalSupply.Quo(lpCoinTotalStaked.Amount).Mul(queuedFarming.Amount)
			}
			mintingCoins := sdk.NewCoins(sdk.NewCoin(lfCoinDenom, mintingAmt))

			// Mint liquid farming coin and send it to the farmer
			if err := h.k.bankKeeper.MintCoins(ctx, types.ModuleName, mintingCoins); err != nil {
				panic(err)
			}
			if err := h.k.bankKeeper.SendCoinsFromModuleToAccount(ctx, types.ModuleName, farmerAcc, mintingCoins); err != nil {
				panic(err)
			}

			// Delete queued farming
			h.k.DeleteQueuedFarming(ctx, endTime, farmingCoinDenom, farmerAcc)

			// TODO: emit events?

			return true
		}
		return false
	})
}

func (h Hooks) AfterAllocateRewards(ctx sdk.Context) {
	// TODO: not implemented yet
	// Select winner -> Distribute rewards -> Refund (?)
	// Stake with the pool coin (auto-compounding role)
	// Create RewardsAuction
}
