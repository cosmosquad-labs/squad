package keeper

// import (
// 	sdk "github.com/cosmos/cosmos-sdk/types"

// 	"github.com/crescent-network/crescent/x/liquidity/types"
// )

// // SetDepositRequest stores ...
// func (k Keeper) SetDepositRequest(ctx sdk.Context, poolId uint64, id uint64, state types.MsgDepositRequest) {
// 	store := ctx.KVStore(k.storeKey)
// 	// bz := types.MustMarshalDepositMsgState(k.cdc, state)
// 	store.Set(types.GetDepositRequestKey(poolId, id), bz)
// }

// // SetWithdrawRequest stores ...
// func (k Keeper) SetWithdrawRequest(ctx sdk.Context, poolId uint64, id uint64, state types.MsgWithdrawRequest) {
// 	store := ctx.KVStore(k.storeKey)
// 	// bz := types.MustMarshalDepositMsgState(k.cdc, state)
// 	store.Set(types.GetWithdrawRequestKey(poolId, id), bz)
// }

// // SetSwapRequest stores ...
// func (k Keeper) SetSwapRequest(ctx sdk.Context, poolId uint64, id uint64, state types.MsgDepositRequest) {
// 	store := ctx.KVStore(k.storeKey)
// 	// bz := types.MustMarshalDepositMsgState(k.cdc, state)
// 	store.Set(types.GetDepositRequestKey(poolId, id), bz)
// }
