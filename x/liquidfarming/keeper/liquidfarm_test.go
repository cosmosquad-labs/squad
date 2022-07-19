package keeper_test

import (
	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"

	utils "github.com/cosmosquad-labs/squad/v2/types"
	"github.com/cosmosquad-labs/squad/v2/x/liquidfarming/types"

	_ "github.com/stretchr/testify/suite"
)

func (s *KeeperTestSuite) TestFarm() {
	farmerAcc := s.addr(0)

	s.createPair(farmerAcc, "denom1", "denom2", true)
	s.createPool(s.addr(0), 1, utils.ParseCoins("100_000_000denom1, 100_000_000denom2"), true)
	s.createLiquidFarm(types.NewLiquidFarm(1, sdk.ZeroInt(), sdk.ZeroInt()))
	s.Require().Len(s.keeper.GetParams(s.ctx).LiquidFarms, 1)

	pool, found := s.app.LiquidityKeeper.GetPool(s.ctx, 1)
	s.Require().True(found)

	amount1, amount2, amount3 := sdk.NewInt(100000000), sdk.NewInt(200000000), sdk.NewInt(300000000)

	s.farm(pool.Id, farmerAcc, sdk.NewCoin(pool.PoolCoinDenom, amount1), true)
	s.nextBlock()

	s.farm(pool.Id, farmerAcc, sdk.NewCoin(pool.PoolCoinDenom, amount2), true)
	s.nextBlock()

	s.farm(pool.Id, farmerAcc, sdk.NewCoin(pool.PoolCoinDenom, amount3), true)
	s.nextBlock()

	// Check if the liquid farm reserve account staked in the farming module
	reserveAcc := types.LiquidFarmReserveAddress(pool.Id)
	queuedCoins := s.app.FarmingKeeper.GetAllQueuedCoinsByFarmer(s.ctx, reserveAcc)
	s.Require().Equal(amount1.Add(amount2).Add(amount3), queuedCoins.AmountOf(pool.PoolCoinDenom))

	// Check queued farmings farmed by the farmer
	queuedFarmings := s.keeper.GetQueuedFarmingsByFarmer(s.ctx, farmerAcc)
	s.Require().Len(queuedFarmings, 3)
}

func (s *KeeperTestSuite) TestFarm_AfterStaked() {
	s.createPair(s.addr(0), "denom1", "denom2", true)
	s.createPool(s.addr(0), 1, utils.ParseCoins("100_000_000denom1, 100_000_000denom2"), true)
	s.createLiquidFarm(types.NewLiquidFarm(1, sdk.ZeroInt(), sdk.ZeroInt()))
	s.Require().Len(s.keeper.GetParams(s.ctx).LiquidFarms, 1)

	pool, found := s.app.LiquidityKeeper.GetPool(s.ctx, 1)
	s.Require().True(found)

	farmerAcc := s.addr(1)
	amount1, amount2 := sdk.NewInt(100_000_000), sdk.NewInt(400_000_000)

	s.farm(pool.Id, farmerAcc, sdk.NewCoin(pool.PoolCoinDenom, amount1), true)
	s.nextBlock()
	s.advanceEpochDays()

	// Staked amount must be 100
	totalStakings, found := s.app.FarmingKeeper.GetTotalStakings(s.ctx, pool.PoolCoinDenom)
	s.Require().True(found)
	s.Require().Equal(amount1, totalStakings.Amount)

	s.farm(pool.Id, farmerAcc, sdk.NewCoin(pool.PoolCoinDenom, amount2), true)
	s.advanceEpochDays()

	// Staked amount must be 500 (amount1 + amount2)
	totalStakings, found = s.app.FarmingKeeper.GetTotalStakings(s.ctx, pool.PoolCoinDenom)
	s.Require().True(found)
	s.Require().Equal(amount1.Add(amount2), totalStakings.Amount)

	// Check minted LFCoin
	lfCoinDenom := types.LFCoinDenom(pool.Id)
	lfCoinBalance := s.getBalance(farmerAcc, lfCoinDenom)
	s.Require().Equal(sdk.NewCoin(lfCoinDenom, amount1.Add(amount2)), lfCoinBalance)

	// Queued farmings must be deleted
	queuedFarmings := s.keeper.GetQueuedFarmingsByFarmer(s.ctx, farmerAcc)
	s.Require().Len(queuedFarmings, 0)
}

func (s *KeeperTestSuite) TestFarm_AfterStaked_MultipleFarmers() {
	s.createPair(s.addr(0), "denom1", "denom2", true)
	s.createPool(s.addr(0), 1, utils.ParseCoins("100_000_000denom1, 100_000_000denom2"), true)
	s.createLiquidFarm(types.NewLiquidFarm(1, sdk.ZeroInt(), sdk.ZeroInt()))

	pool, found := s.app.LiquidityKeeper.GetPool(s.ctx, 1)
	s.Require().True(found)

	farmerAcc1, farmerAcc2, farmerAcc3 := s.addr(1), s.addr(2), s.addr(3)
	amount1, amount2, amount3 := sdk.NewInt(100000000), sdk.NewInt(200000000), sdk.NewInt(666)

	// Multiple farmers farm in the same block
	s.farm(pool.Id, farmerAcc1, sdk.NewCoin(pool.PoolCoinDenom, amount1), true)
	s.farm(pool.Id, farmerAcc2, sdk.NewCoin(pool.PoolCoinDenom, amount2), true)
	s.advanceEpochDays()

	queuedFarmings := s.keeper.GetQueuedFarmingsByFarmer(s.ctx, farmerAcc1)
	s.Require().Len(queuedFarmings, 0)

	queuedFarmings = s.keeper.GetQueuedFarmingsByFarmer(s.ctx, farmerAcc2)
	s.Require().Len(queuedFarmings, 0)

	s.Require().Equal(amount1, s.getBalance(farmerAcc1, types.LFCoinDenom(pool.Id)).Amount)
	s.Require().Equal(amount2, s.getBalance(farmerAcc2, types.LFCoinDenom(pool.Id)).Amount)

	s.farm(pool.Id, farmerAcc3, sdk.NewCoin(pool.PoolCoinDenom, amount3), true)
	s.advanceEpochDays()

	queuedFarmings = s.keeper.GetQueuedFarmingsByFarmer(s.ctx, farmerAcc3)
	s.Require().Len(queuedFarmings, 0)

	reserveAddr := types.LiquidFarmReserveAddress(pool.Id)
	lfCoinTotalSupply := s.app.BankKeeper.GetSupply(s.ctx, types.LFCoinDenom(pool.Id)).Amount
	lpCoinTotalStaked := s.app.FarmingKeeper.GetAllStakedCoinsByFarmer(s.ctx, reserveAddr).AmountOf(pool.PoolCoinDenom)
	mintingAmt := lfCoinTotalSupply.Mul(amount3).Quo(lpCoinTotalStaked)
	s.Require().Equal(mintingAmt, s.getBalance(farmerAcc3, types.LFCoinDenom(pool.Id)).Amount)
}

func (s *KeeperTestSuite) TestCancelQueuedFarming_All() {
	s.createPair(s.addr(0), "denom1", "denom2", true)
	s.createPool(s.addr(0), 1, utils.ParseCoins("100_000_000denom1, 100_000_000denom2"), true)
	s.createLiquidFarm(types.NewLiquidFarm(1, sdk.ZeroInt(), sdk.ZeroInt()))
	s.Require().Len(s.keeper.GetParams(s.ctx).LiquidFarms, 1)

	pool, found := s.app.LiquidityKeeper.GetPool(s.ctx, 1)
	s.Require().True(found)

	farmerAcc := s.addr(1)

	s.farm(pool.Id, farmerAcc, sdk.NewCoin(pool.PoolCoinDenom, sdk.NewInt(500_000_000)), true)
	s.nextBlock()

	s.farm(pool.Id, farmerAcc, sdk.NewCoin(pool.PoolCoinDenom, sdk.NewInt(500_000_000)), true)
	s.nextBlock()

	queuedFarmings := s.keeper.GetQueuedFarmingsByFarmer(s.ctx, farmerAcc)
	s.Require().Len(queuedFarmings, 2)

	// Cancel all amounts
	unfarmingCoin := sdk.NewInt64Coin(pool.PoolCoinDenom, 1_000_000_000)
	s.cancelQueuedFarming(pool.Id, farmerAcc, unfarmingCoin)

	// QueuedFarmings must be deleted
	queuedFarmings = s.keeper.GetQueuedFarmingsByFarmer(s.ctx, farmerAcc)
	s.Require().Len(queuedFarmings, 0)

	// Unfarming amount must be returned to the farmer
	balance := s.getBalance(farmerAcc, pool.PoolCoinDenom)
	s.Require().Equal(unfarmingCoin, balance)
}

func (s *KeeperTestSuite) TestCancelQueuedFarming_Partial() {
	s.createPair(s.addr(0), "denom1", "denom2", true)
	s.createPool(s.addr(0), 1, utils.ParseCoins("100_000_000denom1, 100_000_000denom2"), true)
	s.createLiquidFarm(types.NewLiquidFarm(1, sdk.ZeroInt(), sdk.ZeroInt()))
	s.Require().Len(s.keeper.GetParams(s.ctx).LiquidFarms, 1)

	pool, found := s.app.LiquidityKeeper.GetPool(s.ctx, 1)
	s.Require().True(found)

	farmerAcc := s.addr(1)

	// Farm 1000 in total (400, 300, 400)
	amount1, amount2, amount3 := sdk.NewInt(400_000_000), sdk.NewInt(300_000_000), sdk.NewInt(300_000_000)

	s.farm(pool.Id, farmerAcc, sdk.NewCoin(pool.PoolCoinDenom, amount1), true)
	s.nextBlock()

	s.farm(pool.Id, farmerAcc, sdk.NewCoin(pool.PoolCoinDenom, amount2), true)
	s.nextBlock()

	s.farm(pool.Id, farmerAcc, sdk.NewCoin(pool.PoolCoinDenom, amount3), true)
	s.nextBlock()

	queuedFarmings := s.keeper.GetQueuedFarmingsByFarmer(s.ctx, farmerAcc)
	s.Require().Len(queuedFarmings, 3)

	// Cancel partial amounts
	unfarmingCoin := sdk.NewInt64Coin(pool.PoolCoinDenom, 650_000_000)
	s.cancelQueuedFarming(pool.Id, farmerAcc, unfarmingCoin)

	// Two queuedFarmings must be deleted and the last one must be in store
	queuedFarmings = s.keeper.GetQueuedFarmingsByFarmer(s.ctx, farmerAcc)
	s.Require().Len(queuedFarmings, 1)

	// Check the remaining queued coin
	total := amount1.Add(amount2).Add(amount3)
	s.Require().Equal(total.Sub(unfarmingCoin.Amount), queuedFarmings[0].Amount)

	// Unfarming amount must be returned to the farmer
	balance := s.getBalance(farmerAcc, pool.PoolCoinDenom)
	s.Require().Equal(unfarmingCoin, balance)
}

func (s *KeeperTestSuite) TestCancelQueuedFarming_Exceed() {
	s.createPair(s.addr(0), "denom1", "denom2", true)
	s.createPool(s.addr(0), 1, utils.ParseCoins("100_000_000denom1, 100_000_000denom2"), true)
	s.createLiquidFarm(types.NewLiquidFarm(1, sdk.ZeroInt(), sdk.ZeroInt()))
	s.Require().Len(s.keeper.GetParams(s.ctx).LiquidFarms, 1)

	pool, found := s.app.LiquidityKeeper.GetPool(s.ctx, 1)
	s.Require().True(found)

	farmerAcc := s.addr(1)

	// Farm 1000 in total (400, 300, 400)
	amount1, amount2, amount3 := sdk.NewInt(400_000_000), sdk.NewInt(300_000_000), sdk.NewInt(300_000_000)

	s.farm(pool.Id, farmerAcc, sdk.NewCoin(pool.PoolCoinDenom, amount1), true)
	s.nextBlock()

	s.farm(pool.Id, farmerAcc, sdk.NewCoin(pool.PoolCoinDenom, amount2), true)
	s.nextBlock()

	s.farm(pool.Id, farmerAcc, sdk.NewCoin(pool.PoolCoinDenom, amount3), true)
	s.nextBlock()

	queuedFarmings := s.keeper.GetQueuedFarmingsByFarmer(s.ctx, farmerAcc)
	s.Require().Len(queuedFarmings, 3)

	// Cancel exceeding amount
	unfarmingCoin := sdk.NewInt64Coin(pool.PoolCoinDenom, 1_500_000_000)
	err := s.keeper.CancelQueuedFarming(s.ctx, &types.MsgCancelQueuedFarming{
		PoolId:        pool.Id,
		Farmer:        farmerAcc.String(),
		UnfarmingCoin: unfarmingCoin,
	})
	s.Require().ErrorIs(err, sdkerrors.ErrInsufficientFunds)
}

func (s *KeeperTestSuite) TestUnfarm_All() {
	s.createPair(s.addr(0), "denom1", "denom2", true)
	s.createPool(s.addr(0), 1, utils.ParseCoins("100_000_000denom1, 100_000_000denom2"), true)
	s.createLiquidFarm(types.NewLiquidFarm(1, sdk.ZeroInt(), sdk.ZeroInt()))
	s.Require().Len(s.keeper.GetParams(s.ctx).LiquidFarms, 1)

	pool, found := s.app.LiquidityKeeper.GetPool(s.ctx, 1)
	s.Require().True(found)

	farmerAcc := s.addr(1)
	amount1 := sdk.NewInt(100_000_000)

	s.farm(pool.Id, farmerAcc, sdk.NewCoin(pool.PoolCoinDenom, amount1), true)
	s.nextBlock()
	s.advanceEpochDays()

	// Staked amount must be 100
	totalStakings, found := s.app.FarmingKeeper.GetTotalStakings(s.ctx, pool.PoolCoinDenom)
	s.Require().True(found)
	s.Require().Equal(amount1, totalStakings.Amount)

	queuedFarmings := s.keeper.GetQueuedFarmingsByFarmer(s.ctx, farmerAcc)
	s.Require().Len(queuedFarmings, 0)

	// Check minted LFCoin
	lfCoinDenom := types.LFCoinDenom(pool.Id)
	lfCoinBalance := s.getBalance(farmerAcc, lfCoinDenom)
	s.Require().Equal(sdk.NewCoin(lfCoinDenom, amount1), lfCoinBalance)

	// Unfarm all amounts
	err := s.keeper.Unfarm(s.ctx, types.NewMsgUnfarm(pool.Id, farmerAcc.String(), lfCoinBalance))
	s.Require().NoError(err)

	// Verify
	_, found = s.app.FarmingKeeper.GetTotalStakings(s.ctx, pool.PoolCoinDenom)
	s.Require().False(found)

	supply := s.app.BankKeeper.GetSupply(s.ctx, lfCoinDenom)
	s.Require().Equal(sdk.ZeroInt(), supply.Amount)
}
