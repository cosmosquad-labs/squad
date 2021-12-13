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
		pool       types.PoolI
		isDepleted bool
	}{
		{
			"empty pool",
			&staticPool{
				poolCoinSupply: sdk.ZeroInt(),
				rx:             sdk.ZeroInt(),
				ry:             sdk.ZeroInt(),
			},
			true,
		},
		{
			"depleted, with some coins from outside",
			&staticPool{
				poolCoinSupply: sdk.ZeroInt(),
				rx:             sdk.NewInt(100),
				ry:             sdk.ZeroInt(),
			},
			true,
		},
		{
			"depleted, with some coins from outside #2",
			&staticPool{
				poolCoinSupply: sdk.ZeroInt(),
				rx:             sdk.NewInt(100),
				ry:             sdk.NewInt(100),
			},
			true,
		},
		{
			"normal pool",
			&staticPool{
				poolCoinSupply: types.DefaultInitialPoolCoinSupply,
				rx:             sdk.NewInt(1000000),
				ry:             sdk.NewInt(1000000),
			},
			false,
		},
		{
			"not depleted, but reserve coins are gone",
			&staticPool{
				initialPoolCoinSupply: types.DefaultInitialPoolCoinSupply,
				poolCoinSupply:        types.DefaultInitialPoolCoinSupply,
				rx:                    sdk.ZeroInt(),
				ry:                    sdk.NewInt(1000000),
			},
			true,
		},
	} {
		t.Run(tc.name, func(t *testing.T) {
			ops := types.NewPoolOperations(tc.pool)
			require.Equal(t, tc.isDepleted, ops.IsDepleted())
		})
	}
}

func TestPoolOperations_PoolPrice(t *testing.T) {
	for _, tc := range []struct {
		name string
		pool types.PoolI
		p    sdk.Dec
	}{
		{
			"depleted pool",
			&staticPool{
				poolCoinSupply: sdk.ZeroInt(),
				rx:             sdk.NewInt(100),
				ry:             sdk.NewInt(100),
			},
			sdk.ZeroDec(),
		},
		{
			"normal pool",
			&staticPool{
				poolCoinSupply: types.DefaultInitialPoolCoinSupply,
				rx:             sdk.NewInt(200000000),
				ry:             sdk.NewInt(1000000),
			},
			sdk.NewDec(200),
		},
		{
			"decimal rounding",
			&staticPool{
				poolCoinSupply: types.DefaultInitialPoolCoinSupply,
				rx:             sdk.NewInt(200),
				ry:             sdk.NewInt(300),
			},
			sdk.MustNewDecFromStr("0.666666666666666667"),
		},
	} {
		t.Run(tc.name, func(t *testing.T) {
			ops := types.NewPoolOperations(tc.pool)
			require.True(sdk.DecEq(t, tc.p, ops.PoolPrice()))
		})
	}
}
