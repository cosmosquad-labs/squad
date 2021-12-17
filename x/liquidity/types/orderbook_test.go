package types_test

import (
	"testing"

	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/stretchr/testify/require"
	"github.com/tendermint/tendermint/crypto"

	"github.com/tendermint/farming/x/liquidity/types"
)

func BenchmarkOrderBook_Add(b *testing.B) {
	orderer := sdk.AccAddress(crypto.AddressHash([]byte("addr1")))
	var ob types.OrderBook
	b.ResetTimer()

	for i := 0; i < b.N; i++ {
		ob.Add(types.Order{
			Orderer:         orderer,
			Direction:       types.SwapDirectionXToY,
			Price:           sdk.OneDec(),
			RemainingAmount: sdk.OneInt(),
			ReceivedAmount:  sdk.ZeroInt(),
		})
	}
}

func TestOrderBook_Add(t *testing.T) {
	orderer := sdk.AccAddress(crypto.AddressHash([]byte("addr1")))

	var ob types.OrderBook
	// Only calling `Add` will not modify the order book itself.
	// To modify the order book, caller should assign it to the old
	// variable, just like slices.
	for i := 0; i < 10; i++ {
		ob.Add(types.Order{
			Orderer:         orderer,
			Direction:       types.SwapDirectionXToY,
			Price:           sdk.OneDec(),
			RemainingAmount: sdk.OneInt(),
			ReceivedAmount:  sdk.ZeroInt(),
		})
	}
	require.Len(t, ob, 0)

	// Doing this 10 times to demonstrate how orders are
	// grouped together based on their price.
	for i := 0; i < 10; i++ {
		price := sdk.OneDec()
		// Add orders for 10 different prices.
		for j := 0; j < 10; j++ {
			ob = ob.Add(types.Order{
				Orderer:         orderer,
				Direction:       types.SwapDirectionXToY,
				Price:           price,
				RemainingAmount: sdk.NewInt(100),
				ReceivedAmount:  sdk.ZeroInt(),
			})
			price = price.Add(sdk.MustNewDecFromStr("0.1"))
		}
	}
	require.Len(t, ob, 10)

	// See if the orders are well grouped and sorted in descending order.
	price := sdk.OneDec()
	for i := 0; i < 10; i++ {
		og := ob[len(ob)-i-1]
		require.True(sdk.DecEq(t, price, og.Price))
		require.Len(t, og.XToYOrders, 10)
		require.Len(t, og.YToXOrders, 0)
		price = price.Add(sdk.MustNewDecFromStr("0.1"))
	}
}
