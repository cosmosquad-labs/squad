package keeper_test

import (
	sdk "github.com/cosmos/cosmos-sdk/types"

	"github.com/cosmosquad-labs/squad/v2/x/farming/keeper"
	"github.com/cosmosquad-labs/squad/v2/x/farming/types"

	_ "github.com/stretchr/testify/suite"
)

var _ types.FarmingHooks = (*MockFarmingHooksReceiver)(nil)

// MockFarmingHooksReceiver event hooks for farming object (noalias)
type MockFarmingHooksReceiver struct {
	AfterStakedValid          bool
	AfterAllocateRewardsValid bool
}

func (h *MockFarmingHooksReceiver) AfterStaked(ctx sdk.Context, farmer sdk.AccAddress, stakingCoinDenom string, stakingAmt sdk.Int) {
	h.AfterStakedValid = true
}

func (h *MockFarmingHooksReceiver) AfterAllocateRewards(ctx sdk.Context) {
	h.AfterAllocateRewardsValid = true
}

func (s *KeeperTestSuite) TestHooks() {
	farmingHooksReceiver := MockFarmingHooksReceiver{}

	// Set hooks
	keeper.UnsafeSetHooks(
		&s.keeper, types.NewMultiFarmingHooks(&farmingHooksReceiver),
	)

	// Default must be false
	s.Require().False(farmingHooksReceiver.AfterStakedValid)
	s.Require().False(farmingHooksReceiver.AfterAllocateRewardsValid)

	// Create sample farming plan
	s.CreateFixedAmountPlan(s.addrs[5], map[string]string{denom1: "1"}, map[string]int64{denom3: 700000})

	// Stake
	s.Stake(s.addrs[0], sdk.NewCoins(sdk.NewInt64Coin(denom1, 1000000)))

	// Advanced epoch twice to trigger AllocateRewards function
	s.advanceEpochDays()
	s.advanceEpochDays()

	// Must be true
	s.Require().True(farmingHooksReceiver.AfterStakedValid)
	s.Require().True(farmingHooksReceiver.AfterAllocateRewardsValid)
}
