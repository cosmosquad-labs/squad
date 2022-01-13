package keeper

import (
	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"

	"github.com/crescent-network/crescent/x/liquidity/types"
)

// GetNextPoolIdWithUpdate increments pool id by one and set it.
func (k Keeper) GetNextPoolIdWithUpdate(ctx sdk.Context) uint64 {
	id := k.GetLastPoolId(ctx) + 1
	k.SetLastPoolId(ctx, id)
	return id
}

// CreatePool creates a liquidity pool.
func (k Keeper) CreatePool(ctx sdk.Context, msg *types.MsgCreatePool) error {
	creator := msg.GetCreator()

	params := k.GetParams(ctx)
	if msg.XCoin.Amount.LT(params.MinInitialDepositAmount) || msg.YCoin.Amount.LT(params.MinInitialDepositAmount) {
		return types.ErrInsufficientDepositAmount // TODO: more detail error?
	}

	pair, found := k.GetPairByDenoms(ctx, msg.XCoin.Denom, msg.YCoin.Denom)
	if found {
		// If there is a pair with given denoms, check if there is a pool with
		// the pair.
		// Current version disallows to create multiple pools with same pair,
		// but later this can be changed(in v2).
		found := false
		k.IteratePoolsByPair(ctx, pair.Id, func(pool types.Pool) (stop bool) {
			// TODO: check if pool isn't disabled
			// if !pool.Disabled {
			//     found = true
			//     return true
			// }
			// return false
			found = true
			return true
		})
		if found {
			return types.ErrPoolAlreadyExists
		}
	} else {
		// If there is no such pair, create one and store it to the variable.
		pair = k.CreatePair(ctx, msg.XCoin.Denom, msg.YCoin.Denom)
	}

	// Create and save the new pool object.
	poolId := k.GetNextPoolIdWithUpdate(ctx)
	pool := types.NewPool(poolId, pair.Id, msg.XCoin.Denom, msg.YCoin.Denom)
	k.SetPool(ctx, pool)

	// Send deposit coins to the pool's reserve account.
	depositCoins := sdk.NewCoins(msg.XCoin, msg.YCoin)
	// TODO: can we use multi-send?
	if err := k.bankKeeper.SendCoins(ctx, creator, pool.GetReserveAddress(), depositCoins); err != nil {
		return err
	}
	if err := k.bankKeeper.SendCoinsFromAccountToModule(ctx, creator, types.ModuleName, params.PoolCreationFee); err != nil {
		return sdkerrors.Wrap(err, "insufficient pool creation fee")
	}
	// Mint and send pool coin to the creator.
	poolCoin := sdk.NewCoin(pool.PoolCoinDenom, params.InitialPoolCoinSupply)
	if err := k.bankKeeper.MintCoins(ctx, types.ModuleName, sdk.NewCoins(poolCoin)); err != nil {
		return err
	}
	if err := k.bankKeeper.SendCoinsFromModuleToAccount(ctx, types.ModuleName, creator, sdk.NewCoins(poolCoin)); err != nil {
		return err
	}

	ctx.EventManager().EmitEvents(sdk.Events{
		sdk.NewEvent(
			types.EventTypeCreatePool,
			sdk.NewAttribute(types.AttributeKeyCreator, creator.String()),
			sdk.NewAttribute(types.AttributeKeyXCoin, msg.XCoin.String()),
			sdk.NewAttribute(types.AttributeKeyYCoin, msg.YCoin.String()),
		),
	})

	return nil
}

func (k Keeper) DepositBatch(ctx sdk.Context, msg *types.MsgDepositBatch) error {
	panic("not implemented")
}

func (k Keeper) WithdrawBatch(ctx sdk.Context, msg *types.MsgWithdrawBatch) error {
	panic("not implemented")
}
