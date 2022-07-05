package keeper

import (
	sdk "github.com/cosmos/cosmos-sdk/types"

	"github.com/cosmosquad-labs/squad/x/liquidfarming/types"
)

func (k Keeper) Deposit(ctx sdk.Context, msg *types.MsgDeposit) error {
	// TODO: not implemented yet

	return nil
}

func (k Keeper) Cancel(ctx sdk.Context, msg *types.MsgCancel) error {
	// TODO: not implemented yet

	return nil
}

func (k Keeper) Withdraw(ctx sdk.Context, msg *types.MsgWithdraw) error {
	// TODO: not implemented yet

	return nil
}
