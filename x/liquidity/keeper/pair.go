package keeper

import (
	gogotypes "github.com/gogo/protobuf/types"

	sdk "github.com/cosmos/cosmos-sdk/types"

	"github.com/crescent-network/crescent/x/liquidity/types"
)

// GetLastPairId returns the global pair id counter.
func (k Keeper) GetLastPairId(ctx sdk.Context) uint64 {
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
	id := k.GetLastPairId(ctx) + 1
	k.SetLastPairId(ctx, id)
	return id
}

// SetLastPairId stores the global pair id counter.
func (k Keeper) SetLastPairId(ctx sdk.Context, id uint64) {
	store := ctx.KVStore(k.storeKey)
	bz := k.cdc.MustMarshal(&gogotypes.UInt64Value{Value: id})
	store.Set(types.PairIdKey, bz)
}

// SetPair stores the particular pair.
func (k Keeper) SetPair(ctx sdk.Context, pair types.Pair) {
	store := ctx.KVStore(k.storeKey)
	bz := types.MustMarshalPair(k.cdc, pair)
	store.Set(types.GetPairKey(pair.Id), bz)
}

// SetPairDenom stores the particular denom pair.
func (k Keeper) SetPairDenom(ctx sdk.Context, denomA string, denomB string, pairId uint64) {
	store := ctx.KVStore(k.storeKey)
	store.Set(types.GetPairIndexKey(denomA, denomB, pairId), []byte{})
}

// SetReversePairDenom stores the particular denom pair in reverse order.
func (k Keeper) SetReversePairDenom(ctx sdk.Context, denomB string, denomA string, pairId uint64) {
	store := ctx.KVStore(k.storeKey)
	store.Set(types.GetReversePairIndexKey(denomB, denomA, pairId), []byte{})
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

// IteratePairsByDenom iterates over all the stored pairs by particular denomination and
// performs a callback function. Stops iteration when callback returns true.
func (k Keeper) IteratePairsByDenom(ctx sdk.Context, denom string, cb func(pair types.Pair) (stop bool)) {
	store := ctx.KVStore(k.storeKey)

	iterator := sdk.KVStorePrefixIterator(store, types.GetPairByDenomKey(denom))
	defer iterator.Close()

	for ; iterator.Valid(); iterator.Next() {
		_, pairId := types.ParsePairByDenomIndexKey(iterator.Key())
		pair, _ := k.GetPair(ctx, pairId)
		if cb(pair) {
			break
		}
	}
}

// IterateAllPairs iterates over all the stored pairs by reverse demoniation and
// performs a callback function. Stops iteration when callback returns true.
func (k Keeper) IterateReversePairsByDenom(ctx sdk.Context, denom string, cb func(pair types.Pair) (stop bool)) {
	store := ctx.KVStore(k.storeKey)

	iterator := sdk.KVStorePrefixIterator(store, types.GetReversePairByDenomKey(denom))
	defer iterator.Close()

	for ; iterator.Valid(); iterator.Next() {
		_, pairId := types.ParseReversePairByDenomIndexKey(iterator.Key())
		pair, _ := k.GetPair(ctx, pairId)
		if cb(pair) {
			break
		}
	}
}
