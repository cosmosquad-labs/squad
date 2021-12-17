package types_test

import (
	"testing"

	"github.com/stretchr/testify/require"
	"github.com/tendermint/tendermint/crypto"

	sdk "github.com/cosmos/cosmos-sdk/types"

	"github.com/tendermint/farming/x/liquidity/types"
)

// staticPool is only for testing purpose.
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

func (pool *staticPool) ReserveAddress() sdk.AccAddress {
	return sdk.AccAddress(crypto.AddressHash([]byte("staticPool")))
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
		ps         int64 // pool coin supply
		rx, ry     int64 // reserve balance
		isDepleted bool
	}{
		{
			name:       "empty pool",
			ps:         0,
			rx:         0,
			ry:         0,
			isDepleted: true,
		},
		{
			name:       "depleted, with some coins from outside",
			ps:         0,
			rx:         100,
			ry:         0,
			isDepleted: true,
		},
		{
			name:       "depleted, with some coins from outside #2",
			ps:         0,
			rx:         100,
			ry:         100,
			isDepleted: true,
		},
		{
			name:       "normal pool",
			ps:         10000,
			rx:         10000,
			ry:         10000,
			isDepleted: false,
		},
		{
			name:       "not depleted, but reserve coins are gone",
			ps:         10000,
			rx:         0,
			ry:         10000,
			isDepleted: true,
		},
	} {
		t.Run(tc.name, func(t *testing.T) {
			ops := types.NewPoolOperations(&staticPool{
				poolCoinSupply: sdk.NewInt(tc.ps),
				rx:             sdk.NewInt(tc.rx),
				ry:             sdk.NewInt(tc.ry),
			})
			require.Equal(t, tc.isDepleted, ops.IsDepleted())
		})
	}
}

func TestPoolOperations_PoolPrice(t *testing.T) {
	for _, tc := range []struct {
		name   string
		ps     int64   // pool coin supply
		rx, ry int64   // reserve balance
		p      sdk.Dec // expected pool price
	}{
		{
			name: "depleted pool",
			ps:   0,
			rx:   100,
			ry:   100,
			p:    sdk.ZeroDec(),
		},
		{
			name: "normal pool",
			ps:   10000,
			rx:   20000,
			ry:   100,
			p:    sdk.NewDec(200),
		},
		{
			name: "decimal rounding",
			ps:   10000,
			rx:   200,
			ry:   300,
			p:    sdk.MustNewDecFromStr("0.666666666666666667"),
		},
	} {
		t.Run(tc.name, func(t *testing.T) {
			ops := types.NewPoolOperations(&staticPool{
				poolCoinSupply: sdk.NewInt(tc.ps),
				rx:             sdk.NewInt(tc.rx),
				ry:             sdk.NewInt(tc.ry),
			})
			require.True(sdk.DecEq(t, tc.p, ops.PoolPrice()))
		})
	}
}

func TestPoolOperations_Deposit(t *testing.T) {
	for _, tc := range []struct {
		name   string
		ips    int64 // initial pool coin supply
		ps     int64 // pool coin supply
		rx, ry int64 // reserve balance
		x, y   int64 // depositing coin amount
		ax, ay int64 // expected accepted coin amount
		pc     int64 // expected minted pool coin amount
	}{
		{
			name: "creating a pool",
			ips:  10000,
			ps:   0,
			rx:   0,
			ry:   0,
			x:    100,
			y:    100,
			ax:   100,
			ay:   100,
			pc:   10000,
		},
		{
			name: "reinitialize a depleted pool",
			ips:  10000,
			ps:   0,
			rx:   100,
			ry:   50,
			x:    100,
			y:    100,
			ax:   100,
			ay:   100,
			pc:   10000,
		},
		// TODO: what if a pool has positive pool coin supply
		//       but has zero reserve balance?
		{
			name: "ideal deposit",
			ps:   10000,
			rx:   2000,
			ry:   100,
			x:    200,
			y:    10,
			ax:   200,
			ay:   10,
			pc:   1000,
		},
		{
			name: "unbalanced deposit",
			ps:   10000,
			rx:   2000,
			ry:   100,
			x:    100,
			y:    2000,
			ax:   100,
			ay:   5,
			pc:   500,
		},
		{
			name: "decimal truncation",
			ps:   333,
			rx:   222,
			ry:   333,
			x:    100,
			y:    100,
			ax:   66,
			ay:   99,
			pc:   99,
		},
		{
			name: "decimal truncation #2",
			ps:   333,
			rx:   200,
			ry:   300,
			x:    80,
			y:    80,
			ax:   53,
			ay:   80,
			pc:   88,
		},
		{
			name: "zero minting amount",
			ps:   100,
			rx:   10000,
			ry:   10000,
			x:    99,
			y:    99,
			ax:   0,
			ay:   0,
			pc:   0,
		},
		{
			name: "tiny minting amount",
			ps:   100,
			rx:   10000,
			ry:   10000,
			x:    100,
			y:    100,
			ax:   100,
			ay:   100,
			pc:   1,
		},
		{
			name: "tiny minting amount #2",
			ps:   100,
			rx:   10000,
			ry:   10000,
			x:    199,
			y:    199,
			ax:   100,
			ay:   100,
			pc:   1,
		},
	} {
		t.Run(tc.name, func(t *testing.T) {
			ops := types.NewPoolOperations(&staticPool{
				initialPoolCoinSupply: sdk.NewInt(tc.ips),
				poolCoinSupply:        sdk.NewInt(tc.ps),
				rx:                    sdk.NewInt(tc.rx),
				ry:                    sdk.NewInt(tc.ry),
			})
			ax, ay, pc := ops.Deposit(sdk.NewInt(tc.x), sdk.NewInt(tc.y))
			require.True(sdk.IntEq(t, sdk.NewInt(tc.ax), ax))
			require.True(sdk.IntEq(t, sdk.NewInt(tc.ay), ay))
			require.True(sdk.IntEq(t, sdk.NewInt(tc.pc), pc))
			// Additional assertions
			if !ops.IsDepleted() {
				require.True(t, (ax.Int64()*tc.ps) >= (pc.Int64()*tc.rx)) // (ax / rx) > (pc / ps)
				require.True(t, (ay.Int64()*tc.ps) >= (pc.Int64()*tc.ry)) // (ay / ry) > (pc / ps)
			}
		})
	}
}

func TestPoolOperations_Withdraw(t *testing.T) {
	for _, tc := range []struct {
		name    string
		ps      int64 // pool coin supply
		rx, ry  int64 // reserve balance
		pc      int64 // redeeming pool coin amount
		feeRate sdk.Dec
		x, y    int64 // withdrawn coin amount
	}{
		{
			name:    "ideal withdraw",
			ps:      10000,
			rx:      2000,
			ry:      100,
			pc:      1000,
			feeRate: sdk.ZeroDec(),
			x:       200,
			y:       10,
		},
		{
			name:    "ideal withdraw - with fee",
			ps:      10000,
			rx:      2000,
			ry:      100,
			pc:      1000,
			feeRate: sdk.MustNewDecFromStr("0.003"),
			x:       199,
			y:       9,
		},
		{
			name:    "withdraw all",
			ps:      10,
			rx:      123,
			ry:      567,
			pc:      10,
			feeRate: sdk.MustNewDecFromStr("0.003"),
			x:       123,
			y:       567,
		},
		{
			name:    "advantageous for pool",
			ps:      10000,
			rx:      100,
			ry:      100,
			pc:      99,
			feeRate: sdk.ZeroDec(),
			x:       0,
			y:       0,
		},
	} {
		t.Run(tc.name, func(t *testing.T) {
			ops := types.NewPoolOperations(&staticPool{
				poolCoinSupply: sdk.NewInt(tc.ps),
				rx:             sdk.NewInt(tc.rx),
				ry:             sdk.NewInt(tc.ry),
			})
			x, y := ops.Withdraw(sdk.NewInt(tc.pc), tc.feeRate)
			require.True(sdk.IntEq(t, sdk.NewInt(tc.x), x))
			require.True(sdk.IntEq(t, sdk.NewInt(tc.y), y))
			// Additional assertions
			require.True(t, (tc.pc*tc.rx) >= (x.Int64()*tc.ps))
			require.True(t, (tc.pc*tc.ry) >= (y.Int64()*tc.ps))
		})
	}
}
