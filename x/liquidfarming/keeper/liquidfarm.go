package keeper

import (
	"strconv"
	"time"

	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"

	farmingtypes "github.com/cosmosquad-labs/squad/v2/x/farming/types"
	"github.com/cosmosquad-labs/squad/v2/x/liquidfarming/types"
)

// Farm handles types.MsgFarm to farm for liquid farm
func (k Keeper) Farm(ctx sdk.Context, msg *types.MsgFarm) error {
	params := k.GetParams(ctx)
	poolId := uint64(0)
	minDepositAmt := sdk.ZeroInt()
	for _, lf := range params.LiquidFarms {
		if lf.PoolId == msg.PoolId {
			poolId = lf.PoolId
			minDepositAmt = lf.MinimumDepositAmount
			break
		}
	}

	// Check if liquid farm that corresponds to the pool id exists
	if poolId == 0 {
		return sdkerrors.Wrap(sdkerrors.ErrNotFound, "liquid farm not found")
	}

	// Check if the farming coin amount exceeds minimum deposit amount
	if msg.FarmingCoin.Amount.LT(minDepositAmt) {
		return sdkerrors.Wrapf(types.ErrInsufficientFarmingCoinAmount, "%s is smaller than %s", msg.FarmingCoin, minDepositAmt)
	}

	// Check that the corresponding pool must exist
	pool, found := k.liquidityKeeper.GetPool(ctx, poolId)
	if !found {
		return sdkerrors.Wrapf(sdkerrors.ErrNotFound, "pool %d not found", poolId)
	}

	farmerAcc := msg.GetFarmer()
	farmingCoin := msg.FarmingCoin
	reserveAcc := types.LiquidFarmReserveAddress(poolId)

	// Check if the depositor has sufficient balance of deposit coin
	poolCoinBalance := k.bankKeeper.SpendableCoins(ctx, farmerAcc).AmountOf(pool.PoolCoinDenom)
	if poolCoinBalance.LT(farmingCoin.Amount) {
		return sdkerrors.Wrapf(types.ErrInsufficientFarmingCoinAmount, "%s is smaller than %s", poolCoinBalance, minDepositAmt)
	}

	// Reserve farming coins to reserve account
	if err := k.bankKeeper.SendCoins(ctx, farmerAcc, reserveAcc, sdk.NewCoins(farmingCoin)); err != nil {
		return sdkerrors.Wrap(err, "reserve farming coin")
	}

	// Impose more gas in relative to a number of queued farmings farmed by the farmer
	numQueuedFarmings := 0
	for range k.GetQueuedFarmingsByFarmer(ctx, farmerAcc) {
		numQueuedFarmings++
	}
	if numQueuedFarmings > 0 {
		ctx.GasMeter().ConsumeGas(sdk.Gas(numQueuedFarmings)*params.DelayedFarmGasFee, "DelayedFarmGasFee")
	}

	// Stake with the reserve account in the farming module
	if err := k.farmingKeeper.Stake(ctx, reserveAcc, sdk.NewCoins(farmingCoin)); err != nil {
		return err
	}

	currentEpochDays := k.farmingKeeper.GetCurrentEpochDays(ctx)
	endTime := ctx.BlockTime().Add(time.Duration(currentEpochDays) * farmingtypes.Day)

	// Store queued farming
	k.SetQueuedFarming(ctx, endTime, pool.PoolCoinDenom, farmerAcc, types.QueuedFarming{
		PoolId: msg.PoolId,
		Amount: msg.FarmingCoin.Amount,
	})

	ctx.EventManager().EmitEvents(sdk.Events{
		sdk.NewEvent(
			types.EventTypeFarm,
			sdk.NewAttribute(types.AttributeKeyPoolId, strconv.FormatUint(msg.PoolId, 10)),
			sdk.NewAttribute(types.AttributeKeyFarmer, farmerAcc.String()),
			sdk.NewAttribute(types.AttributeKeyFarmingCoin, farmingCoin.String()),
		),
	})

	return nil
}

// CancelQueuedFarming handles types.MsgCancelQueuedFarming to cancel the queued farming.
func (k Keeper) CancelQueuedFarming(ctx sdk.Context, msg *types.MsgCancelQueuedFarming) error {
	queuedFarmings := k.GetQueuedFarmingsByFarmer(ctx, msg.GetFarmer())
	if len(queuedFarmings) == 0 {
		return sdkerrors.Wrapf(sdkerrors.ErrNotFound, "queued farming by %s not found", msg.Farmer)
	}

	// TODO: since key must have endTime, it is mandatory to receive endTime to delete particular QueuedFarming.

	// qf, found := k.GetQueuedFarming(ctx, msg.PoolId, msg.QueuedFarmingId)
	// if !found {
	// 	return sdkerrors.Wrapf(sdkerrors.ErrNotFound, "queued farming by pool id %d and queued farming id %d not found", msg.PoolId, msg.QueuedFarmingId)
	// }

	// reserveAddr := types.LiquidFarmReserveAddress(qf.PoolId)
	// canceledCoin := qf.FarmingCoin

	// // Unstake with the reserve account from the farming module
	// if err := k.farmingKeeper.Unstake(ctx, reserveAddr, sdk.NewCoins(canceledCoin)); err != nil {
	// 	return err
	// }

	// // Release corresponding pool coin to the farmer
	// if err := k.bankKeeper.SendCoins(ctx, reserveAddr, msg.GetFarmer(), sdk.NewCoins(canceledCoin)); err != nil {
	// 	return err
	// }

	// // Delete the store
	// k.DeleteQueuedFarming(ctx, qf)

	// ctx.EventManager().EmitEvents(sdk.Events{
	// 	sdk.NewEvent(
	// 		types.EventTypeCancelQueuedFarming,
	// 		sdk.NewAttribute(types.AttributeKeyPoolId, strconv.FormatUint(msg.PoolId, 10)),
	// 		sdk.NewAttribute(types.AttributeKeyFarmer, msg.Farmer),
	// 		sdk.NewAttribute(types.AttributeKeyCanceledCoin, canceledCoin.String()),
	// 	),
	// })

	return nil
}

// Unfarm handles types.MsgUnfarm to unfarm LFCoin.
// It unstakes pool coin from the farming module, burns the LFCoin, and releases the corresponding amount of pool coin.
func (k Keeper) Unfarm(ctx sdk.Context, msg *types.MsgUnfarm) error {
	params := k.GetParams(ctx)
	for _, lf := range params.LiquidFarms {
		if msg.PoolId == lf.PoolId {
			reserveAddr := types.LiquidFarmReserveAddress(lf.PoolId)
			lfCoinDenom := types.LFCoinDenom(lf.PoolId)

			lfCoinBalance := k.bankKeeper.SpendableCoins(ctx, msg.GetFarmer()).AmountOf(lfCoinDenom)
			if lfCoinBalance.LT(msg.LFCoin.Amount) {
				return sdkerrors.Wrapf(types.ErrInsufficientUnfarmingAmount, "%s is smaller than %s", lfCoinBalance, msg.LFCoin.Amount)
			}

			pool, found := k.liquidityKeeper.GetPool(ctx, lf.PoolId)
			if !found {
				return sdkerrors.Wrapf(sdkerrors.ErrNotFound, "pool %d not found", lf.PoolId)
			}

			poolCoinBalance := k.bankKeeper.SpendableCoins(ctx, reserveAddr).AmountOf(pool.PoolCoinDenom)
			lfCoinTotalSupply := k.bankKeeper.GetSupply(ctx, lfCoinDenom).Amount
			unfarmingAmt := msg.LFCoin.Amount
			fee := sdk.ZeroInt() // TODO: TBD

			// UnfarmedCoin = LPCoinTotalStaked / LFCoinTotalSupply * UnfarmingAmt * (1 - Fee)
			unfarmedAmt := poolCoinBalance.Quo(lfCoinTotalSupply).Mul(unfarmingAmt).Mul(sdk.OneInt().Sub(fee))
			unfarmedCoin := sdk.NewCoin(pool.PoolCoinDenom, unfarmedAmt)

			// Unstake with the reserve account from the farming module
			if err := k.farmingKeeper.Unstake(ctx, reserveAddr, sdk.NewCoins(unfarmedCoin)); err != nil {
				return err
			}

			// Release corresponding pool coin to the farmer
			if err := k.bankKeeper.SendCoins(ctx, reserveAddr, msg.GetFarmer(), sdk.NewCoins(unfarmedCoin)); err != nil {
				return err
			}

			// Burn the withdrawn LFCoin
			if err := k.bankKeeper.BurnCoins(ctx, types.ModuleName, sdk.NewCoins(msg.LFCoin)); err != nil {
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

			break
		}
	}

	return nil
}
