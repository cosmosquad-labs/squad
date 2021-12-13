package keeper

import (
	sdk "github.com/cosmos/cosmos-sdk/types"

	"github.com/tendermint/farming/x/liquidity/types"
)

var (
	_ types.PoolI = (*wiredPool)(nil)
	_ types.PoolI = (*cachedWiredPool)(nil)
)

type wiredPool struct {
	types.Pool
	k   Keeper
	ctx sdk.Context
}

//nolint
func (k Keeper) wiredPool(ctx sdk.Context, pool types.Pool) *wiredPool {
	return &wiredPool{pool, k, ctx}
}

//nolint
func (k Keeper) cachedWiredPool(ctx sdk.Context, pool types.Pool) *cachedWiredPool {
	return &cachedWiredPool{wiredPool: wiredPool{pool, k, ctx}}
}

func (pool *wiredPool) InitialPoolCoinSupply() sdk.Int {
	k, ctx := pool.k, pool.ctx
	return k.GetParams(ctx).InitialPoolCoinSupply
}

func (pool *wiredPool) PoolCoinSupply() sdk.Int {
	k, ctx := pool.k, pool.ctx
	return k.bankKeeper.GetSupply(ctx, pool.PoolCoinDenom).Amount
}

func (pool *wiredPool) ReserveBalance() (x, y sdk.Int) {
	k, ctx := pool.k, pool.ctx
	x = k.bankKeeper.GetBalance(ctx, pool.GetReserveAddress(), pool.ReserveCoinDenoms[0]).Amount
	y = k.bankKeeper.GetBalance(ctx, pool.GetReserveAddress(), pool.ReserveCoinDenoms[1]).Amount
	return
}

func (pool *wiredPool) SwapRequests() []types.SwapRequest {
	return nil
}

type cachedWiredPool struct {
	wiredPool
	initPoolCoinSupply sdk.Int
	poolCoinSupply     sdk.Int
	rx, ry             sdk.Int
	swapRequests       []types.SwapRequest
}

func (pool *cachedWiredPool) InitialPoolCoinSupply() sdk.Int {
	if pool.initPoolCoinSupply.IsNil() {
		pool.initPoolCoinSupply = pool.wiredPool.InitialPoolCoinSupply()
	}
	return pool.initPoolCoinSupply
}

func (pool *cachedWiredPool) PoolCoinSupply() sdk.Int {
	if pool.poolCoinSupply.IsNil() {
		pool.poolCoinSupply = pool.wiredPool.PoolCoinSupply()
	}
	return pool.poolCoinSupply
}

func (pool *cachedWiredPool) ReserveBalance() (x, y sdk.Int) {
	if pool.rx.IsNil() || pool.ry.IsNil() {
		pool.rx, pool.ry = pool.wiredPool.ReserveBalance()
	}
	return pool.rx, pool.ry
}

func (pool *cachedWiredPool) SwapRequests() []types.SwapRequest {
	if pool.swapRequests == nil {
		pool.swapRequests = pool.wiredPool.SwapRequests()
	}
	return pool.swapRequests
}
