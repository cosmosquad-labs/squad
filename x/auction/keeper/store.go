package keeper

import (
	gogotypes "github.com/gogo/protobuf/types"

	sdk "github.com/cosmos/cosmos-sdk/types"

	"github.com/cosmosquad-labs/squad/x/auction/types"
)

// GetLastAuctionId returns the last auction id.
func (k Keeper) GetLastAuctionId(ctx sdk.Context) uint64 {
	var id uint64
	store := ctx.KVStore(k.storeKey)
	bz := store.Get(types.LastAuctionIdKey)
	if bz == nil {
		id = 0 // initialize the auction id
	} else {
		val := gogotypes.UInt64Value{}
		k.cdc.MustUnmarshal(bz, &val)
		id = val.GetValue()
	}
	return id
}

// SetAuctionId stores the last auction id.
func (k Keeper) SetAuctionId(ctx sdk.Context, id uint64) {
	store := ctx.KVStore(k.storeKey)
	bz := k.cdc.MustMarshal(&gogotypes.UInt64Value{Value: id})
	store.Set(types.LastAuctionIdKey, bz)
}

// GetAuction returns an auction interface from the given auction id.
func (k Keeper) GetAuction(ctx sdk.Context, id uint64) (auction types.Auction, found bool) {
	store := ctx.KVStore(k.storeKey)
	bz := store.Get(types.GetAuctionKey(id))
	if bz == nil {
		return types.Auction{}, false
	}

	k.MustUnmarshalAuction(bz, &auction)

	return auction, true
}

func (k Keeper) SetAuction(ctx sdk.Context, auction types.Auction) {
	store := ctx.KVStore(k.storeKey)
	bz := k.MustMarshalAuction(auction)
	store.Set(types.GetAuctionKey(auction.Id), bz)
}

// GetAuctions returns all auctions in the store.
func (k Keeper) GetAuctions(ctx sdk.Context) (auctions []types.Auction) {
	k.IterateAuctions(ctx, func(auction types.Auction) (stop bool) {
		auctions = append(auctions, auction)
		return false
	})
	return auctions
}

// IterateAuctions iterates over all the stored auctions and performs a callback function.
// Stops iteration when callback returns true.
func (k Keeper) IterateAuctions(ctx sdk.Context, cb func(auction types.Auction) (stop bool)) {
	store := ctx.KVStore(k.storeKey)
	iterator := sdk.KVStorePrefixIterator(store, types.AuctionKeyPrefix)
	defer iterator.Close()

	for ; iterator.Valid(); iterator.Next() {
		var auction types.Auction
		err := k.UnmarshalAuction(iterator.Value(), &auction)
		if err != nil {
			panic(err)
		}
		if cb(auction) {
			break
		}
	}
}
