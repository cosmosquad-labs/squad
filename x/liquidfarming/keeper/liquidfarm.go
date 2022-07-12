package keeper

import (
	"strconv"

	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"

	"github.com/cosmosquad-labs/squad/v2/x/liquidfarming/types"
)

// GetNextQueuedFarmingIdWithUpdate increments the last deposit request id and returns it.
func (k Keeper) GetNextQueuedFarmingIdWithUpdate(ctx sdk.Context, poolId uint64) uint64 {
	nextId := k.GetLastQueuedFarmingId(ctx, poolId) + 1
	k.SetQueuedFarmingId(ctx, poolId, nextId)
	return nextId
}

// Farm handles types.MsgFarm
func (k Keeper) Farm(ctx sdk.Context, msg *types.MsgFarm) (types.QueuedFarming, error) {
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
	if poolId == 0 {
		return types.QueuedFarming{}, types.ErrLiquidFarmNotFound
	}
	if msg.FarmingCoin.Amount.LT(minDepositAmt) {
		return types.QueuedFarming{}, types.ErrInsufficientFarmingCoinAmount
	}

	// Check that the corresponding pool must exist
	pool, found := k.liquidityKeeper.GetPool(ctx, poolId)
	if !found {
		return types.QueuedFarming{}, types.ErrPoolNotFound
	}

	farmerAcc := msg.GetFarmer()
	farmingCoin := msg.FarmingCoin
	reserveAcc := types.LiquidFarmReserveAddress(poolId)

	// Check if the depositor has sufficient balance of deposit coin
	poolCoinBalance := k.bankKeeper.SpendableCoins(ctx, farmerAcc).AmountOf(pool.PoolCoinDenom)
	if poolCoinBalance.LT(farmingCoin.Amount) {
		return types.QueuedFarming{}, types.ErrInsufficientFarmingCoinAmount
	}

	// Reserve farming coins to reserve account
	if err := k.bankKeeper.SendCoins(ctx, farmerAcc, reserveAcc, sdk.NewCoins(farmingCoin)); err != nil {
		return types.QueuedFarming{}, sdkerrors.Wrap(err, "failed to reserve deposit coin")
	}

	// Impose more gas in relative to a number of deposit requests made by the depositor
	numQueuedFarmings := 0
	for _, req := range k.GetQueuedFarmingsByDepositor(ctx, farmerAcc) {
		if req.FarmingCoin.Denom == farmingCoin.Denom {
			numQueuedFarmings++
		}
	}
	if numQueuedFarmings > 0 {
		ctx.GasMeter().ConsumeGas(sdk.Gas(numQueuedFarmings)*params.DelayedDepositGasFee, "DelayedDepositGasFee")
	}

	// Stake in the farming module
	if err := k.farmingKeeper.Stake(ctx, reserveAcc, sdk.NewCoins(farmingCoin)); err != nil {
		return types.QueuedFarming{}, err
	}

	nextId := k.GetNextQueuedFarmingIdWithUpdate(ctx, msg.PoolId)
	queued := types.NewQueuedFarming(nextId, msg.PoolId, farmerAcc.String(), farmingCoin)
	k.SetQueuedFarming(ctx, queued)
	k.SetQueuedFarmingIndex(ctx, queued)

	ctx.EventManager().EmitEvents(sdk.Events{
		sdk.NewEvent(
			types.EventTypeFarm,
			sdk.NewAttribute(types.AttributeKeyPoolId, strconv.FormatUint(msg.PoolId, 10)),
			sdk.NewAttribute(types.AttributeKeyFarmer, farmerAcc.String()),
			sdk.NewAttribute(types.AttributeKeyFarmingCoin, farmingCoin.String()),
		),
	})

	return queued, nil
}

// CancelQueuedFarming handles types.MsgCancel to cancel the queued farming.
func (k Keeper) CancelQueuedFarming(ctx sdk.Context, msg *types.MsgCancelQueuedFarming) error {
	requests := k.GetQueuedFarmingsByDepositor(ctx, msg.GetFarmer())
	if len(requests) == 0 {
		return types.ErrQueuedFarmingNotFound
	}

	req, found := k.GetQueuedFarming(ctx, msg.PoolId, msg.QueuedFarmingId)
	if !found {
		return types.ErrQueuedFarmingNotFound
	}

	k.DeleteQueuedFarming(ctx, req)

	ctx.EventManager().EmitEvents(sdk.Events{
		sdk.NewEvent(
			types.EventTypeCancelQueuedFarming,
			sdk.NewAttribute(types.AttributeKeyPoolId, strconv.FormatUint(msg.PoolId, 10)),
			sdk.NewAttribute(types.AttributeKeyQueuedFarmingId, strconv.FormatUint(msg.QueuedFarmingId, 10)),
			sdk.NewAttribute(types.AttributeKeyFarmer, msg.Farmer),
		),
	})

	return nil
}

// Unfarm handles types.MsgUnfarm to unfarm LFCoin to receive the corresponding staked pool coin.
func (k Keeper) Unfarm(ctx sdk.Context, msg *types.MsgUnfarm) error {
	params := k.GetParams(ctx)

	var unfarmedCoin sdk.Coin
	for _, lf := range params.LiquidFarms {
		if msg.PoolId == lf.PoolId {
			reserveAddr := types.LiquidFarmReserveAddress(lf.PoolId)
			lfCoinDenom := types.LFCoinDenom(lf.PoolId)

			lfCoinBalance := k.bankKeeper.SpendableCoins(ctx, msg.GetFarmer()).AmountOf(lfCoinDenom)
			if lfCoinBalance.LT(msg.LFCoin.Amount) {
				return types.ErrInsufficientWithdrawingAmount
			}

			pool, found := k.liquidityKeeper.GetPool(ctx, lf.PoolId)
			if !found {
				return types.ErrPoolNotFound
			}

			poolCoinBalance := k.bankKeeper.SpendableCoins(ctx, reserveAddr).AmountOf(pool.PoolCoinDenom)
			lfCoinTotalSupply := k.bankKeeper.GetSupply(ctx, lfCoinDenom).Amount
			unfarmingAmt := msg.LFCoin.Amount
			fee := sdk.ZeroInt() // TODO: TBD

			// UnfarmedCoin = LPCoinTotalStaked / LFCoinTotalSupply * UnfarmingAmt * (1 - Fee)
			unfarmedAmt := poolCoinBalance.Quo(lfCoinTotalSupply).Mul(unfarmingAmt).Mul(sdk.OneInt().Sub(fee))
			unfarmedCoin = sdk.NewCoin(pool.PoolCoinDenom, unfarmedAmt)

			// Withdraw corresponding pool coin for the LFCoin
			if err := k.bankKeeper.SendCoins(ctx, reserveAddr, msg.GetFarmer(), sdk.NewCoins(unfarmedCoin)); err != nil {
				return err
			}

			break
		}
	}

	// Unstake from the farming module
	if err := k.farmingKeeper.Unstake(ctx, msg.GetFarmer(), sdk.NewCoins(unfarmedCoin)); err != nil {
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

	return nil
}
