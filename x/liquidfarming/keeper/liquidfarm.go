package keeper

import (
	"strconv"
	"time"

	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"

	farmingtypes "github.com/cosmosquad-labs/squad/v2/x/farming/types"
	"github.com/cosmosquad-labs/squad/v2/x/liquidfarming/types"
)

// Farm handles types.MsgFarm to liquid farm.
func (k Keeper) Farm(ctx sdk.Context, msg *types.MsgFarm) error {
	params := k.GetParams(ctx)

	poolId := uint64(0)
	minFarmAmt := sdk.ZeroInt()
	for _, liquidFarm := range params.LiquidFarms {
		if liquidFarm.PoolId == msg.PoolId {
			poolId = liquidFarm.PoolId
			minFarmAmt = liquidFarm.MinimumFarmAmount
			break
		}
	}
	if poolId == 0 {
		return sdkerrors.Wrapf(sdkerrors.ErrNotFound, "liquid farm by pool %d not found", msg.PoolId)
	}

	if msg.FarmingCoin.Amount.LT(minFarmAmt) {
		return sdkerrors.Wrapf(types.ErrSmallerThanMinimumAmount, "%s is smaller than %s", msg.FarmingCoin.Amount, minFarmAmt)
	}

	pool, found := k.liquidityKeeper.GetPool(ctx, poolId)
	if !found {
		return sdkerrors.Wrapf(sdkerrors.ErrNotFound, "pool %d not found", poolId)
	}

	if pool.PoolCoinDenom != msg.FarmingCoin.Denom {
		return sdkerrors.Wrapf(sdkerrors.ErrInvalidRequest, "expected denom %s, but got %s", pool.PoolCoinDenom, msg.FarmingCoin.Denom)
	}

	farmerAddr := msg.GetFarmer()
	farmingCoin := msg.FarmingCoin
	reserveAddr := types.LiquidFarmReserveAddress(poolId)

	poolCoinBalance := k.bankKeeper.SpendableCoins(ctx, farmerAddr).AmountOf(pool.PoolCoinDenom)
	if poolCoinBalance.LT(farmingCoin.Amount) {
		return sdkerrors.Wrapf(sdkerrors.ErrInsufficientFunds, "%s is smaller than %s", poolCoinBalance, minFarmAmt)
	}

	// Reserve farming coins
	if err := k.bankKeeper.SendCoins(ctx, farmerAddr, reserveAddr, sdk.NewCoins(farmingCoin)); err != nil {
		return sdkerrors.Wrap(err, "reserve farming coin")
	}

	// Impose more gas in relative to a number of queued farmings farmed by the farmer
	// This prevents from potential spamming attack
	numQueuedFarmings := 0
	for range k.GetQueuedFarmingsByFarmer(ctx, farmerAddr) {
		numQueuedFarmings++
	}
	if numQueuedFarmings > 0 {
		ctx.GasMeter().ConsumeGas(sdk.Gas(numQueuedFarmings)*params.DelayedFarmGasFee, "DelayedFarmGasFee")
	}

	// Stake with the reserve account in the farming module
	if err := k.farmingKeeper.Stake(ctx, reserveAddr, sdk.NewCoins(farmingCoin)); err != nil {
		return err
	}

	currentEpochDays := k.farmingKeeper.GetCurrentEpochDays(ctx)
	endTime := ctx.BlockTime().Add(time.Duration(currentEpochDays) * farmingtypes.Day) // current time + epoch days

	k.SetQueuedFarming(ctx, endTime, pool.PoolCoinDenom, farmerAddr, types.QueuedFarming{
		PoolId: msg.PoolId,
		Amount: msg.FarmingCoin.Amount,
	})

	ctx.EventManager().EmitEvents(sdk.Events{
		sdk.NewEvent(
			types.EventTypeFarm,
			sdk.NewAttribute(types.AttributeKeyPoolId, strconv.FormatUint(msg.PoolId, 10)),
			sdk.NewAttribute(types.AttributeKeyFarmer, farmerAddr.String()),
			sdk.NewAttribute(types.AttributeKeyFarmingCoin, farmingCoin.String()),
		),
	})

	return nil
}

// Unfarm handles types.MsgUnfarm to unfarm LFCoin.
func (k Keeper) Unfarm(ctx sdk.Context, msg *types.MsgUnfarm) error {
	for _, liquidFarm := range k.GetParams(ctx).LiquidFarms {
		if msg.PoolId == liquidFarm.PoolId {
			reserveAddr := types.LiquidFarmReserveAddress(liquidFarm.PoolId)
			lfCoinDenom := types.LiquidFarmCoinDenom(liquidFarm.PoolId)

			lfCoinBalance := k.bankKeeper.SpendableCoins(ctx, msg.GetFarmer()).AmountOf(lfCoinDenom)
			if lfCoinBalance.LT(msg.LFCoin.Amount) {
				return sdkerrors.Wrapf(types.ErrInsufficientUnfarmingAmount, "%s is smaller than %s", lfCoinBalance, msg.LFCoin.Amount)
			}

			pool, found := k.liquidityKeeper.GetPool(ctx, liquidFarm.PoolId)
			if !found {
				return sdkerrors.Wrapf(sdkerrors.ErrNotFound, "pool %d not found", liquidFarm.PoolId)
			}

			lfCoinTotalSupply := k.bankKeeper.GetSupply(ctx, lfCoinDenom).Amount
			lpCoinTotalStaked := k.farmingKeeper.GetAllStakedCoinsByFarmer(ctx, reserveAddr).AmountOf(pool.PoolCoinDenom)
			unfarmFee := sdk.ZeroInt() // TODO: TBD

			// UnfarmedAmount = TotalStakedLPAmount / TotalSupplyLFAmount * UnfarmingLFAmount * (1 - UnfarmFee)
			unfarmedAmt := lpCoinTotalStaked.Quo(lfCoinTotalSupply).Mul(msg.LFCoin.Amount).Mul(sdk.OneInt().Sub(unfarmFee))
			unfarmedCoin := sdk.NewCoin(pool.PoolCoinDenom, unfarmedAmt)

			// Send the unfarming LFCoin to module account and burn them
			if err := k.bankKeeper.SendCoinsFromAccountToModule(ctx, msg.GetFarmer(), types.ModuleName, sdk.NewCoins(msg.LFCoin)); err != nil {
				return err
			}
			if err := k.bankKeeper.BurnCoins(ctx, types.ModuleName, sdk.NewCoins(msg.LFCoin)); err != nil {
				return err
			}

			// Unstake with the reserve account and release corresponding pool coin amount back to the farmer
			if err := k.farmingKeeper.Unstake(ctx, reserveAddr, sdk.NewCoins(unfarmedCoin)); err != nil {
				return err
			}
			if err := k.bankKeeper.SendCoins(ctx, reserveAddr, msg.GetFarmer(), sdk.NewCoins(unfarmedCoin)); err != nil {
				return err
			}

			ctx.EventManager().EmitEvents(sdk.Events{
				sdk.NewEvent(
					types.EventTypeUnfarm,
					sdk.NewAttribute(types.AttributeKeyPoolId, strconv.FormatUint(msg.PoolId, 10)),
					sdk.NewAttribute(types.AttributeKeyFarmer, msg.Farmer),
					sdk.NewAttribute(types.AttributeKeyUnfarmingCoin, msg.LFCoin.String()),
					sdk.NewAttribute(types.AttributeKeyUnfarmedCoin, unfarmedCoin.String()),
				),
			})

			// Break out of the loop to stop from executing further
			break
		}
	}

	return nil
}

// CancelQueuedFarming handles types.MsgCancelQueuedFarming to cancel queued farming.
func (k Keeper) CancelQueuedFarming(ctx sdk.Context, msg *types.MsgCancelQueuedFarming) error {
	queuedFarmings := k.GetQueuedFarmingsByFarmer(ctx, msg.GetFarmer())
	if len(queuedFarmings) == 0 {
		return sdkerrors.Wrapf(sdkerrors.ErrNotFound, "queued farming by %s not found", msg.Farmer)
	}

	farmerAddr := msg.GetFarmer()
	farmingCoin := msg.UnfarmingCoin

	canceled := sdk.ZeroInt()
	k.IterateQueuedFarmingsByFarmerAndDenomReverse(ctx, farmerAddr, farmingCoin.Denom, func(endTime time.Time, queuedFarming types.QueuedFarming) (stop bool) {
		if endTime.After(ctx.BlockTime()) { // sanity check
			amtToCancel := sdk.MinInt(farmingCoin.Amount.Sub(canceled), queuedFarming.Amount)
			queuedFarming.Amount = queuedFarming.Amount.Sub(amtToCancel)
			if queuedFarming.Amount.IsZero() {
				k.DeleteQueuedFarming(ctx, endTime, farmingCoin.Denom, farmerAddr)
			} else {
				k.SetQueuedFarming(ctx, endTime, farmingCoin.Denom, farmerAddr, queuedFarming)
			}

			canceled = canceled.Add(amtToCancel)
			if canceled.Equal(farmingCoin.Amount) { // fully canceled from queued farmings, so stop
				return true
			}
		}
		return false
	})

	if farmingCoin.Amount.GT(canceled) {
		return sdkerrors.Wrapf(sdkerrors.ErrInsufficientFunds, "%s is smaller than %s", farmingCoin.Amount, canceled)
	}

	reserveAddr := types.LiquidFarmReserveAddress(msg.PoolId)
	canceledCoin := sdk.NewCoin(farmingCoin.Denom, canceled)

	// Unstake the canceled amount with the reserve account in the farming module
	if err := k.farmingKeeper.Unstake(ctx, reserveAddr, sdk.NewCoins(canceledCoin)); err != nil {
		return err
	}

	// Release the corresponding pool coin amount back to the farmer
	if err := k.bankKeeper.SendCoins(ctx, reserveAddr, msg.GetFarmer(), sdk.NewCoins(canceledCoin)); err != nil {
		return err
	}

	ctx.EventManager().EmitEvents(sdk.Events{
		sdk.NewEvent(
			types.EventTypeCancelQueuedFarming,
			sdk.NewAttribute(types.AttributeKeyPoolId, strconv.FormatUint(msg.PoolId, 10)),
			sdk.NewAttribute(types.AttributeKeyFarmer, msg.Farmer),
			sdk.NewAttribute(types.AttributeKeyCanceledCoin, canceledCoin.String()),
		),
	})

	return nil
}
