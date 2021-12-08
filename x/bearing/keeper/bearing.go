package keeper

import (
	sdk "github.com/cosmos/cosmos-sdk/types"
	banktypes "github.com/cosmos/cosmos-sdk/x/bank/types"

	"github.com/tendermint/farming/x/bearing/types"
)

// CollectBearings collects all the valid bearings registered in params.Bearings and
// distributes the total collected coins to destination address.
func (k Keeper) CollectBearings(ctx sdk.Context) error {
	params := k.GetParams(ctx)
	var bearings []types.Bearing
	if params.EpochBlocks > 0 && ctx.BlockHeight()%int64(params.EpochBlocks) == 0 {
		bearings = types.CollectibleBearings(params.Bearings, ctx.BlockTime())
	}
	if len(bearings) == 0 {
		return nil
	}

	// Get a map GetBearingsBySourceMap that has a list of bearings and their total rate, which
	// contain the same SourceAddress
	bearingsBySourceMap := types.GetBearingsBySourceMap(bearings)
	for source, bearingsBySource := range bearingsBySourceMap {
		sourceAcc, err := sdk.AccAddressFromBech32(source)
		if err != nil {
			return err
		}
		sourceBalances := sdk.NewDecCoinsFromCoins(k.bankKeeper.GetAllBalances(ctx, sourceAcc)...)
		if sourceBalances.IsZero() {
			continue
		}

		var inputs []banktypes.Input
		var outputs []banktypes.Output
		bearingsBySource.CollectionCoins = make([]sdk.Coins, len(bearingsBySource.Bearings))
		for i, bearing := range bearingsBySource.Bearings {
			destinationAcc, err := sdk.AccAddressFromBech32(bearing.DestinationAddress)
			if err != nil {
				return err
			}

			collectionCoins, _ := sourceBalances.MulDecTruncate(bearing.Rate).TruncateDecimal()
			if collectionCoins.Empty() || !collectionCoins.IsValid() {
				continue
			}

			inputs = append(inputs, banktypes.NewInput(sourceAcc, collectionCoins))
			outputs = append(outputs, banktypes.NewOutput(destinationAcc, collectionCoins))
			bearingsBySource.CollectionCoins[i] = collectionCoins
		}

		if err := k.bankKeeper.InputOutputCoins(ctx, inputs, outputs); err != nil {
			return err
		}

		for i, bearing := range bearingsBySource.Bearings {
			k.AddTotalCollectedCoins(ctx, bearing.Name, bearingsBySource.CollectionCoins[i])
			ctx.EventManager().EmitEvents(sdk.Events{
				sdk.NewEvent(
					types.EventTypeBearingCollected,
					sdk.NewAttribute(types.AttributeValueName, bearing.Name),
					sdk.NewAttribute(types.AttributeValueDestinationAddress, bearing.DestinationAddress),
					sdk.NewAttribute(types.AttributeValueSourceAddress, bearing.SourceAddress),
					sdk.NewAttribute(types.AttributeValueRate, bearing.Rate.String()),
					sdk.NewAttribute(types.AttributeValueAmount, bearingsBySource.CollectionCoins[i].String()),
				),
			})
		}
	}
	return nil
}

// GetTotalCollectedCoins returns total collected coins for a bearing.
func (k Keeper) GetTotalCollectedCoins(ctx sdk.Context, bearingName string) sdk.Coins {
	store := ctx.KVStore(k.storeKey)
	bz := store.Get(types.GetTotalCollectedCoinsKey(bearingName))
	if bz == nil {
		return nil
	}
	var collectedCoins types.TotalCollectedCoins
	k.cdc.MustUnmarshal(bz, &collectedCoins)
	return collectedCoins.TotalCollectedCoins
}

// IterateAllTotalCollectedCoins iterates over all the stored TotalCollectedCoins and performs a callback function.
// Stops iteration when callback returns true.
func (k Keeper) IterateAllTotalCollectedCoins(ctx sdk.Context, cb func(record types.BearingRecord) (stop bool)) {
	store := ctx.KVStore(k.storeKey)
	iterator := sdk.KVStorePrefixIterator(store, types.TotalCollectedCoinsKeyPrefix)

	defer iterator.Close()
	for ; iterator.Valid(); iterator.Next() {
		var record types.BearingRecord
		var collectedCoins types.TotalCollectedCoins
		k.cdc.MustUnmarshal(iterator.Value(), &collectedCoins)
		record.Name = types.ParseTotalCollectedCoinsKey(iterator.Key())
		record.TotalCollectedCoins = collectedCoins.TotalCollectedCoins
		if cb(record) {
			break
		}
	}
}

// SetTotalCollectedCoins sets total collected coins for a bearing.
func (k Keeper) SetTotalCollectedCoins(ctx sdk.Context, bearingName string, amount sdk.Coins) {
	store := ctx.KVStore(k.storeKey)
	collectedCoins := types.TotalCollectedCoins{TotalCollectedCoins: amount}
	bz := k.cdc.MustMarshal(&collectedCoins)
	store.Set(types.GetTotalCollectedCoinsKey(bearingName), bz)
}

// AddTotalCollectedCoins increases total collected coins for a bearing.
func (k Keeper) AddTotalCollectedCoins(ctx sdk.Context, bearingName string, amount sdk.Coins) {
	collectedCoins := k.GetTotalCollectedCoins(ctx, bearingName)
	collectedCoins = collectedCoins.Add(amount...)
	k.SetTotalCollectedCoins(ctx, bearingName, collectedCoins)
}
