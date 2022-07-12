package keeper

import (
	sdk "github.com/cosmos/cosmos-sdk/types"

	"github.com/cosmosquad-labs/squad/v2/x/farming/types"
)

// Implements FarmingHooks interface
var _ types.FarmingHooks = Keeper{}

// AfterStaked - call hook if registered
func (k Keeper) AfterStaked(ctx sdk.Context) {
	if k.hooks != nil {
		k.hooks.AfterStaked(ctx)
	}
}

// AfterAllocateRewards - call hook if registered
func (k Keeper) AfterAllocateRewards(ctx sdk.Context) {
	if k.hooks != nil {
		k.hooks.AfterAllocateRewards(ctx)
	}
}
