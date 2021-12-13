package types_test

import (
	"testing"

	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/stretchr/testify/require"

	"github.com/tendermint/farming/x/liquidity/types"
)

var _ types.PoolI = (*staticPool)(nil)

type staticPool struct {
	initialPoolCoinSupply sdk.Int
	poolCoinSupply        sdk.Int
	rx, ry                sdk.Int
	swapRequests          []types.SwapRequest
}

func (pool *staticPool) InitialPoolCoinSupply() sdk.Int {
	return pool.initialPoolCoinSupply
}

func (pool *staticPool) PoolCoinSupply() sdk.Int {
	return pool.poolCoinSupply
}

func (pool *staticPool) ReserveBalance() (x, y sdk.Int) {
	return pool.rx, pool.ry
}

func (pool *staticPool) SwapRequests() []types.SwapRequest {
	return pool.swapRequests
}

func TestPoolOperations_IsDepleted(t *testing.T) {
	for _, tc := range []struct {
		name       string
		ps         sdk.Int // pool coin supply
		rx, ry     sdk.Int // reserve balance
		isDepleted bool
	}{
		{
			name:       "empty pool",
			ps:         sdk.ZeroInt(),
			rx:         sdk.ZeroInt(),
			ry:         sdk.ZeroInt(),
			isDepleted: true,
		},
		{
			name:       "depleted, with some coins from outside",
			ps:         sdk.ZeroInt(),
			rx:         sdk.NewInt(100),
			ry:         sdk.ZeroInt(),
			isDepleted: true,
		},
		{
			name:       "depleted, with some coins from outside #2",
			ps:         sdk.ZeroInt(),
			rx:         sdk.NewInt(100),
			ry:         sdk.NewInt(100),
			isDepleted: true,
		},
		{
			name:       "normal pool",
			ps:         types.DefaultInitialPoolCoinSupply,
			rx:         sdk.NewInt(1000000),
			ry:         sdk.NewInt(1000000),
			isDepleted: false,
		},
		{
			name:       "not depleted, but reserve coins are gone",
			ps:         types.DefaultInitialPoolCoinSupply,
			rx:         sdk.ZeroInt(),
			ry:         sdk.NewInt(1000000),
			isDepleted: true,
		},
	} {
		t.Run(tc.name, func(t *testing.T) {
			ops := types.NewPoolOperations(&staticPool{
				poolCoinSupply: tc.ps,
				rx:             tc.rx,
				ry:             tc.ry,
			})
			require.Equal(t, tc.isDepleted, ops.IsDepleted())
		})
	}
}

func TestPoolOperations_PoolPrice(t *testing.T) {
	for _, tc := range []struct {
		name   string
		ps     sdk.Int // pool coin supply
		rx, ry sdk.Int // reserve balance
		p      sdk.Dec // expected pool price
	}{
		{
			name: "depleted pool",
			ps:   sdk.ZeroInt(),
			rx:   sdk.NewInt(100),
			ry:   sdk.NewInt(100),
			p:    sdk.ZeroDec(),
		},
		{
			name: "normal pool",
			ps:   types.DefaultInitialPoolCoinSupply,
			rx:   sdk.NewInt(200000000),
			ry:   sdk.NewInt(1000000),
			p:    sdk.NewDec(200),
		},
		{
			name: "decimal rounding",
			ps:   types.DefaultInitialPoolCoinSupply,
			rx:   sdk.NewInt(200),
			ry:   sdk.NewInt(300),
			p:    sdk.MustNewDecFromStr("0.666666666666666667"),
		},
	} {
		t.Run(tc.name, func(t *testing.T) {
			ops := types.NewPoolOperations(&staticPool{
				poolCoinSupply: tc.ps,
				rx:             tc.rx,
				ry:             tc.ry,
			})
			require.True(sdk.DecEq(t, tc.p, ops.PoolPrice()))
		})
	}
}
