package types_test

import (
	"testing"

	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/stretchr/testify/require"
	"github.com/tendermint/tendermint/crypto"

	"github.com/tendermint/farming/x/liquidity/types"
)

func BenchmarkOrderBook_Add(b *testing.B) {
	orderer := sdk.AccAddress(crypto.AddressHash([]byte("addr1"))).String()
	var orderBook types.OrderBook
	b.ResetTimer()

	for i := 0; i < b.N; i++ {
		orderBook.Add(types.NewOrder(orderer, types.SwapDirectionXToY, sdk.OneDec(), sdk.OneInt()))
	}
}

func TestOrderBook_Add(t *testing.T) {
	orderer := sdk.AccAddress(crypto.AddressHash([]byte("addr1"))).String()

	var orderBook types.OrderBook
	// Only calling `Add` will not modify the order book itself.
	// To modify the order book, caller should assign it to the old
	// variable, just like slices.
	for i := 0; i < 10; i++ {
		orderBook.Add(types.NewOrder(orderer, types.SwapDirectionXToY, sdk.OneDec(), sdk.OneInt()))
	}
	require.Len(t, orderBook, 0)

	// Doing this 10 times to demonstrate how orders are
	// grouped together based on their price.
	for i := 0; i < 10; i++ {
		price := sdk.OneDec()
		// Add orders for 10 different prices.
		for j := 0; j < 10; j++ {
			orderBook = orderBook.Add(types.NewOrder(orderer, types.SwapDirectionXToY, price, sdk.NewInt(100)))
			price = price.Add(sdk.MustNewDecFromStr("0.1"))
		}
	}
	require.Len(t, orderBook, 10)

	// See if the orders are well grouped and sorted in descending order.
	price := sdk.OneDec()
	for i := 0; i < 10; i++ {
		g := orderBook[len(orderBook)-i-1]
		require.True(sdk.DecEq(t, price, g.Price))
		require.Len(t, g.XToYOrders, 10)
		require.Len(t, g.YToXOrders, 0)
		price = price.Add(sdk.MustNewDecFromStr("0.1"))
	}
}
