package keeper

import (
	"strconv"

	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"

	"github.com/cosmosquad-labs/squad/x/liquidfarming/types"
)

// GetNextDepositRequestIdWithUpdate increments the last deposit request id and returns it.
func (k Keeper) GetNextDepositRequestIdWithUpdate(ctx sdk.Context, poolId uint64) uint64 {
	nextId := k.GetLastDepositRequestId(ctx, poolId) + 1
	k.SetDepositRequestId(ctx, poolId, nextId)
	return nextId
}

// Deposit handles types.MsgDeposit and makes a deposit.
func (k Keeper) Deposit(ctx sdk.Context, msg *types.MsgDeposit) (types.DepositRequest, error) {
	// Liquid farm that corresponds to the pool id must be registered in params
	params := k.GetParams(ctx)
	poolId := uint64(0)
	for _, liquidFarm := range params.LiquidFarms {
		if liquidFarm.PoolId == msg.PoolId {
			poolId = liquidFarm.PoolId
			break
		}
	}
	if poolId == 0 {
		return types.DepositRequest{}, types.ErrLiquidFarmNotFound
	}

	// Pool with the given pool id must exist in order to proceed
	pool, found := k.liquidityKeeper.GetPool(ctx, poolId)
	if !found {
		return types.DepositRequest{}, types.ErrPoolNotFound
	}

	// Check if the depositor has sufficient balance of deposit coin
	spendable := k.bankKeeper.SpendableCoins(ctx, msg.GetDepositor())
	poolCoinBalance := spendable.AmountOf(pool.PoolCoinDenom)
	if poolCoinBalance.LT(msg.DepositCoin.Amount) {
		return types.DepositRequest{}, sdkerrors.ErrInsufficientFunds
	}

	// Reserve the deposit coin to the liquid farm reserve account
	if err := k.bankKeeper.SendCoins(ctx, msg.GetDepositor(), types.LiquidFarmReserveAddress(poolId), sdk.NewCoins(msg.DepositCoin)); err != nil {
		return types.DepositRequest{}, sdkerrors.Wrap(err, "failed to reserve deposit coin")
	}

	// Store deposit request
	nextId := k.GetNextDepositRequestIdWithUpdate(ctx, msg.PoolId)
	req := types.NewDepositRequest(nextId, msg.PoolId, msg.Depositor, msg.DepositCoin)
	k.SetDepositRequest(ctx, req)
	k.SetDepositRequestIndex(ctx, req)

	ctx.EventManager().EmitEvents(sdk.Events{
		sdk.NewEvent(
			types.EventTypeDeposit,
			sdk.NewAttribute(types.AttributeKeyPoolId, strconv.FormatUint(msg.PoolId, 10)),
			sdk.NewAttribute(types.AttributeKeyDepositor, msg.Depositor),
			sdk.NewAttribute(types.AttributeKeyDepositCoin, msg.DepositCoin.String()),
		),
	})

	return req, nil
}

// Cancel handles types.MsgCancel to cancel the deposit request.
func (k Keeper) Cancel(ctx sdk.Context, msg *types.MsgCancel) error {
	req, found := k.GetDepositRequest(ctx, msg.PoolId, msg.DepositRequestId)
	if !found {
		return types.ErrDepositRequestNotFound
	}

	k.DeleteDepositRequest(ctx, req)

	ctx.EventManager().EmitEvents(sdk.Events{
		sdk.NewEvent(
			types.EventTypeCancel,
			sdk.NewAttribute(types.AttributeKeyPoolId, strconv.FormatUint(msg.PoolId, 10)),
			sdk.NewAttribute(types.AttributeKeyDepositRequestId, strconv.FormatUint(msg.DepositRequestId, 10)),
			sdk.NewAttribute(types.AttributeKeyDepositor, msg.Depositor),
		),
	})

	return nil
}

// Withdraw handles types.MsgWithdraw to withdraw LFCoin to receive corresponding deposited pool coin.
func (k Keeper) Withdraw(ctx sdk.Context, msg *types.MsgWithdraw) error {
	params := k.GetParams(ctx)

	var withdrawnCoin sdk.Coin
	for _, liquidFarm := range params.LiquidFarms {
		poolId := liquidFarm.PoolId
		reserveAddr := types.LiquidFarmReserveAddress(poolId)

		pool, found := k.liquidityKeeper.GetPool(ctx, poolId)
		if !found {
			return types.ErrPoolNotFound
		}

		spendable := k.bankKeeper.SpendableCoins(ctx, reserveAddr)
		poolCoinBalance := spendable.AmountOf(pool.PoolCoinDenom)
		lfCoinBalance := k.bankKeeper.GetSupply(ctx, types.LFCoinDenom(poolId))

		// WithdrawnCoin = LPCoinTotalStaked / LFCoinTotalSupply * WithdrawingAmount * (1 - WithdrawFee)
		withdrawFee := sdk.NewInt(0) // TODO: TBD
		withdrawnAmt := poolCoinBalance.Quo(lfCoinBalance.Amount).Mul(msg.LFCoin.Amount).Mul(sdk.OneInt().Sub(withdrawFee))
		withdrawnCoin = sdk.NewCoin(pool.PoolCoinDenom, withdrawnAmt)

		// Withdraw corresponding pool coin for the LFCoin
		if err := k.bankKeeper.SendCoins(ctx, reserveAddr, msg.GetWithdrawer(), sdk.NewCoins(withdrawnCoin)); err != nil {
			return err
		}
	}

	// Burn the withdrawn LFCoin
	if err := k.bankKeeper.BurnCoins(ctx, types.ModuleName, sdk.NewCoins(msg.LFCoin)); err != nil {
		return err
	}

	ctx.EventManager().EmitEvents(sdk.Events{
		sdk.NewEvent(
			types.EventTypeWithdraw,
			sdk.NewAttribute(types.AttributeKeyPoolId, strconv.FormatUint(msg.PoolId, 10)),
			sdk.NewAttribute(types.AttributeKeyWithdrawer, msg.Withdrawer),
			sdk.NewAttribute(types.AttributeKeyWithdrawingCoin, msg.LFCoin.String()),
			sdk.NewAttribute(types.AttributeKeyWithdrawnCoin, withdrawnCoin.String()),
		),
	})

	return nil
}
