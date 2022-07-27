package keeper

import (
	"time"

	gogotypes "github.com/gogo/protobuf/types"

	sdk "github.com/cosmos/cosmos-sdk/types"

	"github.com/cosmosquad-labs/squad/v2/x/liquidfarming/types"
)

// GetQueuedFarming returns a queued farming object for the given end time, farming coin denom and farmer.
func (k Keeper) GetQueuedFarming(ctx sdk.Context, endTime time.Time, farmingCoinDenom string, farmerAcc sdk.AccAddress) (queuedFarming types.QueuedFarming, found bool) {
	store := ctx.KVStore(k.storeKey)
	bz := store.Get(types.GetQueuedFarmingKey(endTime, farmingCoinDenom, farmerAcc))
	if bz == nil {
		return
	}
	k.cdc.MustUnmarshal(bz, &queuedFarming)
	found = true
	return
}

// GetQueuedFarmingsByFarmer returns all queued farming objects by a farmer.
func (k Keeper) GetQueuedFarmingsByFarmer(ctx sdk.Context, farmerAcc sdk.AccAddress) []types.QueuedFarming {
	queuedFarmings := []types.QueuedFarming{}
	k.IterateQueuedFarmingsByFarmer(ctx, farmerAcc, func(farmingCoinDenom string, endTime time.Time, queuedFarming types.QueuedFarming) (stop bool) {
		queuedFarmings = append(queuedFarmings, queuedFarming)
		return false
	})
	return queuedFarmings
}

// SetQueuedFarming stores a queued farming with the given end time, farming coin denom, and farmer address.
func (k Keeper) SetQueuedFarming(ctx sdk.Context, endTime time.Time, farmingCoinDenom string, farmerAcc sdk.AccAddress, queuedFarming types.QueuedFarming) {
	store := ctx.KVStore(k.storeKey)
	bz := k.cdc.MustMarshal(&queuedFarming)
	store.Set(types.GetQueuedFarmingKey(endTime, farmingCoinDenom, farmerAcc), bz)
	store.Set(types.GetQueuedFarmingIndexKey(farmerAcc, farmingCoinDenom, endTime), []byte{})
}

// DeleteQueuedFarming deletes a queued farming for given farming coin denom
// and farmer.
func (k Keeper) DeleteQueuedFarming(ctx sdk.Context, endTime time.Time, farmingCoinDenom string, farmerAcc sdk.AccAddress) {
	store := ctx.KVStore(k.storeKey)
	store.Delete(types.GetQueuedFarmingKey(endTime, farmingCoinDenom, farmerAcc))
	store.Delete(types.GetQueuedFarmingIndexKey(farmerAcc, farmingCoinDenom, endTime))
}

// GetLastRewardsAuctionId returns the last rewards auction id.
func (k Keeper) GetLastRewardsAuctionId(ctx sdk.Context, poolId uint64) uint64 {
	var id uint64
	store := ctx.KVStore(k.storeKey)
	bz := store.Get(types.GetLastRewardsAuctionIdKey(poolId))
	if bz == nil {
		id = 0 // initialize the auction id
	} else {
		val := gogotypes.UInt64Value{}
		k.cdc.MustUnmarshal(bz, &val)
		id = val.GetValue()
	}
	return id
}

// SetRewardsAuctionId stores the last rewards auction id.
func (k Keeper) SetRewardsAuctionId(ctx sdk.Context, poolId uint64, id uint64) {
	store := ctx.KVStore(k.storeKey)
	bz := k.cdc.MustMarshal(&gogotypes.UInt64Value{Value: id})
	store.Set(types.GetLastRewardsAuctionIdKey(poolId), bz)
}

// GetRewardsAuction returns the reward auction object by the given pool id and auction id.
func (k Keeper) GetRewardsAuction(ctx sdk.Context, poolId, auctionId uint64) (auction types.RewardsAuction, found bool) {
	store := ctx.KVStore(k.storeKey)
	bz := store.Get(types.GetRewardsAuctionKey(poolId, auctionId))
	if bz == nil {
		return auction, false
	}

	auction = types.MustUnmarshalRewardsAuction(k.cdc, bz)

	return auction, true
}

// GetAllRewardsAuctions returns all rewards auctions in the store.
func (k Keeper) GetAllRewardsAuctions(ctx sdk.Context) (auctions []types.RewardsAuction) {
	k.IterateRewardsAuctions(ctx, func(auction types.RewardsAuction) (stop bool) {
		auctions = append(auctions, auction)
		return false
	})
	return auctions
}

// SetRewardsAuction stores rewards auction.
func (k Keeper) SetRewardsAuction(ctx sdk.Context, auction types.RewardsAuction) {
	store := ctx.KVStore(k.storeKey)
	bz := types.MustMarshalRewardsAuction(k.cdc, auction)
	store.Set(types.GetRewardsAuctionKey(auction.PoolId, auction.Id), bz)
}

// GetBid returns the bid object by the given pool id and bidder address.
func (k Keeper) GetBid(ctx sdk.Context, poolId uint64, bidder sdk.AccAddress) (bid types.Bid, found bool) {
	store := ctx.KVStore(k.storeKey)
	bz := store.Get(types.GetBidKey(poolId, bidder))
	if bz == nil {
		return bid, false
	}
	k.cdc.MustUnmarshal(bz, &bid)
	return bid, true
}

// SetBid stores a bid object.
func (k Keeper) SetBid(ctx sdk.Context, bid types.Bid) {
	store := ctx.KVStore(k.storeKey)
	bz := k.cdc.MustMarshal(&bid)
	store.Set(types.GetBidKey(bid.PoolId, bid.GetBidder()), bz)
}

// DeleteBid deletes the bid object.
func (k Keeper) DeleteBid(ctx sdk.Context, bid types.Bid) {
	store := ctx.KVStore(k.storeKey)
	store.Delete(types.GetBidKey(bid.PoolId, bid.GetBidder()))
}

// GetBidsByPoolId returns all bid objects by the pool id.
func (k Keeper) GetBidsByPoolId(ctx sdk.Context, poolId uint64) []types.Bid {
	bids := []types.Bid{}
	k.IterateBidsByPoolId(ctx, poolId, func(bid types.Bid) (stop bool) {
		bids = append(bids, bid)
		return false
	})
	return bids
}

// GetWinningBid returns the winning bid object by the given pool id and auction id.
func (k Keeper) GetWinningBid(ctx sdk.Context, poolId uint64, auctionId uint64) (bid types.Bid, found bool) {
	store := ctx.KVStore(k.storeKey)
	bz := store.Get(types.GetWinningBidKey(poolId, auctionId))
	if bz == nil {
		return bid, false
	}
	k.cdc.MustUnmarshal(bz, &bid)
	return bid, true
}

// SetWinningBid stores the winning bid with the auction id.
func (k Keeper) SetWinningBid(ctx sdk.Context, bid types.Bid, auctionId uint64) {
	store := ctx.KVStore(k.storeKey)
	bz := k.cdc.MustMarshal(&bid)
	store.Set(types.GetWinningBidKey(bid.PoolId, auctionId), bz)
}

// IterateQueuedFarmings iterates through all queued farming objects
// stored in the store and invokes callback function for each item.
// Stops the iteration when the callback function returns true.
func (k Keeper) IterateQueuedFarmings(ctx sdk.Context, cb func(endTime time.Time, farmingCoinDenom string, farmerAcc sdk.AccAddress, queuedFarming types.QueuedFarming) (stop bool)) {
	store := ctx.KVStore(k.storeKey)
	iter := sdk.KVStorePrefixIterator(store, types.QueuedFarmingKeyPrefix)
	defer iter.Close()
	for ; iter.Valid(); iter.Next() {
		var queuedFarming types.QueuedFarming
		k.cdc.MustUnmarshal(iter.Value(), &queuedFarming)
		endTime, farmingCoinDenom, farmerAcc := types.ParseQueuedFarmingKey(iter.Key())
		if cb(endTime, farmingCoinDenom, farmerAcc, queuedFarming) {
			break
		}
	}
}

// IterateQueuedFarmingsByFarmer iterates through all queued farming objects
// by farmer stored in the store and invokes callback function for each item.
// Stops the iteration when the callback function returns true.
func (k Keeper) IterateQueuedFarmingsByFarmer(ctx sdk.Context, farmerAcc sdk.AccAddress, cb func(farmingCoinDenom string, endTime time.Time, queuedFarming types.QueuedFarming) (stop bool)) {
	store := ctx.KVStore(k.storeKey)
	iter := sdk.KVStorePrefixIterator(store, types.GetQueuedFarmingsByFarmerPrefix(farmerAcc))
	defer iter.Close()
	for ; iter.Valid(); iter.Next() {
		_, farmingCoinDenom, endTime := types.ParseQueuedFarmingIndexKey(iter.Key())
		queuedFarming, _ := k.GetQueuedFarming(ctx, endTime, farmingCoinDenom, farmerAcc)
		if cb(farmingCoinDenom, endTime, queuedFarming) {
			break
		}
	}
}

// IterateMatureQueuedFarmings iterates through all the queued farming objects
// that are mature at current time.
// Stops the iteration when the callback function returns true.
func (k Keeper) IterateMatureQueuedFarmings(ctx sdk.Context, currTime time.Time, cb func(endTime time.Time, farmingCoinDenom string, farmerAcc sdk.AccAddress, queuedFarming types.QueuedFarming) (stop bool)) {
	store := ctx.KVStore(k.storeKey)
	iter := store.Iterator(types.QueuedFarmingKeyPrefix, types.GetQueuedFarmingEndBytes(currTime))
	defer iter.Close()
	for ; iter.Valid(); iter.Next() {
		var queuedFarming types.QueuedFarming
		k.cdc.MustUnmarshal(iter.Value(), &queuedFarming)
		endTime, farmingCoinDenom, farmerAcc := types.ParseQueuedFarmingKey(iter.Key())
		if cb(endTime, farmingCoinDenom, farmerAcc, queuedFarming) {
			break
		}
	}
}

// IterateQueuedFarmingsByFarmerAndDenomReverse iterates through all the queued farming objects
// by farmer address and staking coin denom in reverse order.
// Stops the iteration when the callback function returns true.
func (k Keeper) IterateQueuedFarmingsByFarmerAndDenomReverse(ctx sdk.Context, farmerAcc sdk.AccAddress, farmingCoinDenom string, cb func(endTime time.Time, queuedFarming types.QueuedFarming) (stop bool)) {
	store := ctx.KVStore(k.storeKey)
	iter := sdk.KVStoreReversePrefixIterator(store, types.GetQueuedFarmingsByFarmerAndDenomPrefix(farmerAcc, farmingCoinDenom))
	defer iter.Close()
	for ; iter.Valid(); iter.Next() {
		_, _, endTime := types.ParseQueuedFarmingIndexKey(iter.Key())
		queuedFarming, _ := k.GetQueuedFarming(ctx, endTime, farmingCoinDenom, farmerAcc)
		if cb(endTime, queuedFarming) {
			break
		}
	}
}

// IterateRewardsAuctions iterates over all the stored auctions and performs a callback function.
// Stops iteration when callback returns true.
func (k Keeper) IterateRewardsAuctions(ctx sdk.Context, cb func(auction types.RewardsAuction) (stop bool)) {
	store := ctx.KVStore(k.storeKey)
	iterator := sdk.KVStorePrefixIterator(store, types.RewardsAuctionKeyPrefix)
	defer iterator.Close()
	for ; iterator.Valid(); iterator.Next() {
		auction := types.MustUnmarshalRewardsAuction(k.cdc, iterator.Value())
		if cb(auction) {
			break
		}
	}
}

// IterateBidsBy PoolId iterates through all bids by pool id stored in the store and
// invokes callback function for each item.
// Stops the iteration when the callback function returns true.
func (k Keeper) IterateBidsByPoolId(ctx sdk.Context, poolId uint64, cb func(bid types.Bid) (stop bool)) {
	store := ctx.KVStore(k.storeKey)
	iter := sdk.KVStorePrefixIterator(store, types.GetBidByPoolIdPrefix(poolId))
	defer iter.Close()
	for ; iter.Valid(); iter.Next() {
		var bid types.Bid
		k.cdc.MustUnmarshal(iter.Value(), &bid)
		if cb(bid) {
			break
		}
	}
}
