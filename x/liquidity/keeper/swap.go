package keeper

import (
	sdk "github.com/cosmos/cosmos-sdk/types"

	"github.com/crescent-network/crescent/x/liquidity/types"
)

func (k Keeper) SwapBatch(ctx sdk.Context, msg *types.MsgSwapBatch) error {
	panic("not implemented")
}

func (k Keeper) CancelSwapBatch(ctx sdk.Context, msg *types.MsgCancelSwapBatch) error {
	panic("not implemented")
}
