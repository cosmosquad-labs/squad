package keeper_test

import (
	"math/rand"
	"time"

	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/cosmosquad-labs/squad/v2/x/liquidfarming/types"

	_ "github.com/stretchr/testify/suite"
)

func (s *KeeperTestSuite) TestQueuedFarming() {
	// TODO: not implemented yet
	// Set | Get | Delete
}

func (s *KeeperTestSuite) TestIterateQueuedFarmingsByFarmerAndDenomReverse() {
	poolId := uint64(1)
	poolCoinDenom := "pool1"
	farmerAcc := s.addr(0)

	s.createPair(farmerAcc, "denom1", "denom2", true)
	s.createPool(farmerAcc, 1, sdk.NewCoins(sdk.NewInt64Coin("denom1", 100000000), sdk.NewInt64Coin("denom2", 100000000)), true)
	s.createLiquidFarm(types.NewLiquidFarm(poolId, sdk.ZeroInt(), sdk.ZeroInt()))
	s.Require().Len(s.keeper.GetParams(s.ctx).LiquidFarms, 1)

	for seed := int64(0); seed <= 5; seed++ {
		r := rand.New(rand.NewSource(seed))

		s.farm(poolId, farmerAcc, sdk.NewInt64Coin(poolCoinDenom, r.Int63()+1), true)
		s.nextBlock()
	}

	skip := true // first item
	prevEndTime := time.Time{}
	s.keeper.IterateQueuedFarmingsByFarmerAndDenomReverse(s.ctx, farmerAcc, poolCoinDenom, func(endTime time.Time, queuedFarming types.QueuedFarming) (stop bool) {
		if skip {
			skip = false
		} else {
			s.Require().True(prevEndTime.After(endTime))
		}
		prevEndTime = endTime

		return false
	})
}
