package keeper

import (
	sdk "github.com/cosmos/cosmos-sdk/types"

	"github.com/cosmosquad-labs/squad/v2/x/farming/types"
)

// Implements FarmingHooks interface
var _ types.FarmingHooks = Keeper{}

// AfterStaked - call hook if registered
func (k Keeper) AfterStaked(ctx sdk.Context, farmer sdk.AccAddress, stakingCoinDenom string, stakingAmt sdk.Int) {
	if k.hooks != nil {
		k.hooks.AfterStaked(ctx, farmer, stakingCoinDenom, stakingAmt)
	}
}

// AfterAllocateRewards - call hook if registered
func (k Keeper) AfterAllocateRewards(ctx sdk.Context) {
	if k.hooks != nil {
		k.hooks.AfterAllocateRewards(ctx)
	}
}
