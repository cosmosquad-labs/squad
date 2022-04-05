package keeper

import (
	"fmt"

	sdk "github.com/cosmos/cosmos-sdk/types"

	utils "github.com/cosmosquad-labs/squad/types"
	"github.com/cosmosquad-labs/squad/x/liquidity/types"
)

const (
	EscrowCoinsAfterBatchStateKey = "escrow_coins_after_batch"
	CanceledOrderCoinsStateKey    = "canceled_order_coins"
	ExpiredOrderCoinsStateKey     = "expired_order_coins"
)

func (k Keeper) CheckEscrowBalances(ctx sdk.Context) {
	_ = k.IterateAllPairs(ctx, func(pair types.Pair) (stop bool, err error) {
		remainingOfferCoins := sdk.Coins{}
		_ = k.IterateOrdersByPair(ctx, pair.Id, func(order types.Order) (stop bool, err error) {
			if !order.Status.ShouldBeDeleted() {
				remainingOfferCoins = remainingOfferCoins.Add(order.RemainingOfferCoin)
			}
			return false, nil
		})
		escrowBalances := k.bankKeeper.GetAllBalances(ctx, pair.GetEscrowAddress())
		if !escrowBalances.IsEqual(remainingOfferCoins) {
			panic(fmt.Errorf("pair %d, escrow balances %s != remaining offer coins %s", pair.Id, escrowBalances, remainingOfferCoins))
		}
		return false, nil
	})

	globalEscrowBalances := k.bankKeeper.GetAllBalances(ctx, types.GlobalEscrowAddress)
	if !globalEscrowBalances.IsZero() {
		panic(fmt.Errorf("global escrow balances %s != 0", globalEscrowBalances))
	}
}

func (k Keeper) GetAllOrderEscrowCoins(ctx sdk.Context) sdk.Coins {
	escrowCoins := sdk.Coins{}
	_ = k.IterateAllPairs(ctx, func(pair types.Pair) (stop bool, err error) {
		escrowCoins = escrowCoins.Add(k.bankKeeper.GetAllBalances(ctx, pair.GetEscrowAddress())...)
		return false, nil
	})
	return escrowCoins
}

func (k Keeper) AfterOrderCanceled(ctx sdk.Context, order types.Order) {
	utils.State.AddCoins(CanceledOrderCoinsStateKey, sdk.NewCoins(order.RemainingOfferCoin))
}

func (k Keeper) AfterOrderExpired(ctx sdk.Context, order types.Order) {
	utils.State.AddCoins(ExpiredOrderCoinsStateKey, sdk.NewCoins(order.RemainingOfferCoin))
}

func (k Keeper) BeforeMatching(ctx sdk.Context) {
	if val, ok := utils.State.Get(EscrowCoinsAfterBatchStateKey); ok {
		escrowCoinsBefore := val.(sdk.Coins)
		escrowCoins := k.GetAllOrderEscrowCoins(ctx)

		newOrderCoins := sdk.Coins{}
		_ = k.IterateAllOrders(ctx, func(order types.Order) (stop bool, err error) {
			if order.Status == types.OrderStatusNotExecuted {
				newOrderCoins = newOrderCoins.Add(order.RemainingOfferCoin)
			}
			return false, nil
		})

		canceledCoins := utils.State.GetCoins(CanceledOrderCoinsStateKey)
		expiredCoins := utils.State.GetCoins(ExpiredOrderCoinsStateKey)

		if !escrowCoins.IsEqual(escrowCoinsBefore.Add(newOrderCoins...).Sub(canceledCoins).Sub(expiredCoins)) {
			panic(fmt.Errorf(
				"invalid changes in escrow balances; escrowBefore=%s, newOrders=%s, canceled=%s, expired=%s, escrowAfter=%s",
				escrowCoinsBefore, newOrderCoins, canceledCoins, expiredCoins, escrowCoins))
		}

		utils.State.Delete(EscrowCoinsAfterBatchStateKey)
	}
}

func (k Keeper) AfterMatching(ctx sdk.Context) {
	utils.State.Set(EscrowCoinsAfterBatchStateKey, k.GetAllOrderEscrowCoins(ctx))
	utils.State.Delete(CanceledOrderCoinsStateKey)
	utils.State.Delete(ExpiredOrderCoinsStateKey)
}
