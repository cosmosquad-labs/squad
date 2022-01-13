package keeper

import (
	sdk "github.com/cosmos/cosmos-sdk/types"

	"github.com/crescent-network/crescent/x/liquidity/types"
)

// GetNextPairIdWithUpdate increments pair id by one and set it.
func (k Keeper) GetNextPairIdWithUpdate(ctx sdk.Context) uint64 {
	id := k.GetLastPairId(ctx) + 1
	k.SetLastPairId(ctx, id)
	return id
}

// CreatePair creates a new pair.
func (k Keeper) CreatePair(ctx sdk.Context, xCoinDenom, yCoinDenom string) types.Pair {
	id := k.GetNextPairIdWithUpdate(ctx)
	pair := types.NewPair(id, xCoinDenom, yCoinDenom)
	k.SetPair(ctx, pair)
	return pair
}

// GetNextSwapRequestIdWithUpdate increments the pair's last swap request
// id and returns it.
func (k Keeper) GetNextSwapRequestIdWithUpdate(ctx sdk.Context, pair types.Pair) uint64 {
	id := pair.LastSwapRequestId + 1
	pair.LastSwapRequestId = id
	k.SetPair(ctx, pair)
	return id
}
