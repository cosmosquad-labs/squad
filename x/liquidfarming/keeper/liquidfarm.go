package keeper

import (
	"strconv"

	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"

	"github.com/cosmosquad-labs/squad/v2/x/liquidfarming/types"
)

// GetNextDepositRequestIdWithUpdate increments the last deposit request id and returns it.
func (k Keeper) GetNextDepositRequestIdWithUpdate(ctx sdk.Context, poolId uint64) uint64 {
	nextId := k.GetLastDepositRequestId(ctx, poolId) + 1
	k.SetDepositRequestId(ctx, poolId, nextId)
	return nextId
}

// Deposit handles types.MsgDeposit and makes a deposit.
func (k Keeper) Deposit(ctx sdk.Context, msg *types.MsgDeposit) (types.DepositRequest, error) {
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
		return types.DepositRequest{}, types.ErrLiquidFarmNotFound
	}
	if msg.DepositCoin.Amount.LT(minDepositAmt) {
		return types.DepositRequest{}, types.ErrInsufficientDepositAmount
	}

	// Check that the corresponding pool must exist
	pool, found := k.liquidityKeeper.GetPool(ctx, poolId)
	if !found {
		return types.DepositRequest{}, types.ErrPoolNotFound
	}

	// Check if the depositor has sufficient balance of deposit coin
	poolCoinBalance := k.bankKeeper.SpendableCoins(ctx, msg.GetDepositor()).AmountOf(pool.PoolCoinDenom)
	if poolCoinBalance.LT(msg.DepositCoin.Amount) {
		return types.DepositRequest{}, types.ErrInsufficientDepositAmount
	}

	if err := k.bankKeeper.SendCoins(ctx, msg.GetDepositor(), types.LiquidFarmReserveAddress(poolId), sdk.NewCoins(msg.DepositCoin)); err != nil {
		return types.DepositRequest{}, sdkerrors.Wrap(err, "failed to reserve deposit coin")
	}

	// Impose more gas in relative to a number of deposit requests made by the depositor
	numDepositRequests := 0
	for _, req := range k.GetDepositRequestsByDepositor(ctx, msg.GetDepositor()) {
		if req.DepositCoin.Denom == msg.DepositCoin.Denom {
			numDepositRequests++
		}
	}
	if numDepositRequests > 0 {
		ctx.GasMeter().ConsumeGas(sdk.Gas(numDepositRequests)*params.DelayedDepositGasFee, "DelayedDepositGasFee")
	}

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
	requests := k.GetDepositRequestsByDepositor(ctx, msg.GetDepositor())
	if len(requests) == 0 {
		return types.ErrDepositRequestNotFound
	}

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

// Withdraw handles types.MsgWithdraw to withdraw LFCoin to receive the corresponding deposited pool coin.
func (k Keeper) Withdraw(ctx sdk.Context, msg *types.MsgWithdraw) error {
	params := k.GetParams(ctx)

	var withdrawnCoin sdk.Coin
	for _, lf := range params.LiquidFarms {
		if msg.PoolId == lf.PoolId {
			reserveAddr := types.LiquidFarmReserveAddress(lf.PoolId)
			lfCoinDenom := types.LFCoinDenom(lf.PoolId)

			lfCoinBalance := k.bankKeeper.SpendableCoins(ctx, msg.GetWithdrawer()).AmountOf(lfCoinDenom)
			if lfCoinBalance.LT(msg.LFCoin.Amount) {
				return types.ErrInsufficientWithdrawingAmount
			}

			pool, found := k.liquidityKeeper.GetPool(ctx, lf.PoolId)
			if !found {
				return types.ErrPoolNotFound
			}

			poolCoinBalance := k.bankKeeper.SpendableCoins(ctx, reserveAddr).AmountOf(pool.PoolCoinDenom)
			lfCoinTotalSupply := k.bankKeeper.GetSupply(ctx, lfCoinDenom).Amount
			withdrawingAmt := msg.LFCoin.Amount
			withdrawFee := sdk.ZeroInt() // TODO: TBD

			// WithdrawnCoin = LPCoinTotalStaked / LFCoinTotalSupply * WithdrawingAmount * (1 - WithdrawFee)
			withdrawnAmt := poolCoinBalance.Quo(lfCoinTotalSupply).Mul(withdrawingAmt).Mul(sdk.OneInt().Sub(withdrawFee))
			withdrawnCoin = sdk.NewCoin(pool.PoolCoinDenom, withdrawnAmt)

			// Withdraw corresponding pool coin for the LFCoin
			if err := k.bankKeeper.SendCoins(ctx, reserveAddr, msg.GetWithdrawer(), sdk.NewCoins(withdrawnCoin)); err != nil {
				return err
			}

			break
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
