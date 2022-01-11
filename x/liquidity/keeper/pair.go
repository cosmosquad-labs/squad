package keeper

import (
	gogotypes "github.com/gogo/protobuf/types"

	sdk "github.com/cosmos/cosmos-sdk/types"

	"github.com/crescent-network/crescent/x/liquidity/types"
)

// GetPairId returns the global pair id counter.
func (k Keeper) GetPairId(ctx sdk.Context) uint64 {
	var id uint64
	store := ctx.KVStore(k.storeKey)
	bz := store.Get(types.PairIdKey)
	if bz == nil {
		id = 0 // initialize the pair id
	} else {
		val := gogotypes.UInt64Value{}
		err := k.cdc.Unmarshal(bz, &val)
		if err != nil {
			panic(err)
		}
		id = val.GetValue()
	}
	return id
}

// GetPair returns pair object for the given pair id.
func (k Keeper) GetPair(ctx sdk.Context, id uint64) (pair types.Pair, found bool) {
	store := ctx.KVStore(k.storeKey)
	key := types.GetPairKey(id)

	value := store.Get(key)
	if value == nil {
		return pair, false
	}

	pair = types.MustUnmarshalPair(k.cdc, value)

	return pair, true
}

// GetAllPairs returns all pairs in the store.
func (k Keeper) GetAllPairs(ctx sdk.Context) (pairs []types.Pair) {
	k.IterateAllPairs(ctx, func(pair types.Pair) (stop bool) {
		pairs = append(pairs, pair)
		return false
	})

	return pairs
}

// GetNextPairIdWithUpdate increments pair id by one and set it.
func (k Keeper) GetNextPairIdWithUpdate(ctx sdk.Context) uint64 {
	id := k.GetPairId(ctx) + 1
	k.SetPairId(ctx, id)
	return id
}

// SetPairId sets the global pair id counter.
func (k Keeper) SetPairId(ctx sdk.Context, id uint64) {
	store := ctx.KVStore(k.storeKey)
	bz := k.cdc.MustMarshal(&gogotypes.UInt64Value{Value: id})
	store.Set(types.PairIdKey, bz)
}

// IterateAllPairs iterates over all the stored pairs and performs a callback function.
// Stops iteration when callback returns true.
func (k Keeper) IterateAllPairs(ctx sdk.Context, cb func(pair types.Pair) (stop bool)) {
	store := ctx.KVStore(k.storeKey)

	iterator := sdk.KVStorePrefixIterator(store, types.PairKeyPrefix)
	defer iterator.Close()

	for ; iterator.Valid(); iterator.Next() {
		pair := types.MustUnmarshalPair(k.cdc, iterator.Value())
		if cb(pair) {
			break
		}
	}
}
