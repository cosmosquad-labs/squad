package keeper_test

import (
	"fmt"

	sdk "github.com/cosmos/cosmos-sdk/types"

	chain "github.com/cosmosquad-labs/squad/app"
	utils "github.com/cosmosquad-labs/squad/types"
	"github.com/cosmosquad-labs/squad/x/farming"
	"github.com/cosmosquad-labs/squad/x/farming/types"
)

type Action interface {
	Do(*KeeperTestSuite)
}

type StakeAction struct {
	farmerAcc sdk.AccAddress
	amount    sdk.Coins
}

func (sa StakeAction) Do(suite *KeeperTestSuite) {
	fmt.Printf("Stake(%s, %s)\n", sa.farmerAcc, sa.amount)
	err := suite.keeper.Stake(suite.ctx, sa.farmerAcc, sa.amount)
	suite.Require().NoError(err)
}

type UnstakeAction struct {
	farmerAcc sdk.AccAddress
	amount    sdk.Coins
}

func (ua UnstakeAction) Do(suite *KeeperTestSuite) {
	fmt.Printf("Unstake(%s, %s)\n", ua.farmerAcc, ua.amount)
	err := suite.keeper.Unstake(suite.ctx, ua.farmerAcc, ua.amount)
	suite.Require().NoError(err)
}

type HarvestAction struct {
	farmerAcc         sdk.AccAddress
	stakingCoinDenoms []string
}

func (ha HarvestAction) Do(suite *KeeperTestSuite) {
	fmt.Printf("Harvest(%s, %s)\n", ha.farmerAcc, ha.stakingCoinDenoms)
	err := suite.keeper.Harvest(suite.ctx, ha.farmerAcc, ha.stakingCoinDenoms)
	suite.Require().NoError(err)
}

type AdvanceDayAction struct{}

func (AdvanceDayAction) Do(suite *KeeperTestSuite) {
	fmt.Println("AdvanceDay()")
	suite.ctx = suite.ctx.WithBlockTime(suite.ctx.BlockTime().Add(types.Day))
	farming.EndBlocker(suite.ctx, suite.keeper)
}

type BalanceAssertion struct {
	acc    sdk.AccAddress
	denom  string
	amount sdk.Int
}

func (ba BalanceAssertion) Do(suite *KeeperTestSuite) {
	balance := suite.app.BankKeeper.GetBalance(suite.ctx, ba.acc, ba.denom)
	fmt.Printf("BalanceAssertion(%s, %s, %s)\n", ba.acc, ba.denom, ba.amount)
	suite.Require().True(intEq(ba.amount, balance.Amount))
}

type AllUnharvestedRewardsAssertion struct {
	acc     sdk.AccAddress
	rewards sdk.Coins
}

func (assertion AllUnharvestedRewardsAssertion) Do(suite *KeeperTestSuite) {
	fmt.Printf("AllUnharvestedRewardsAssertion(%s, %s)\n", assertion.acc, assertion.rewards)
	rewards := suite.keeper.AllUnharvestedRewards(suite.ctx, assertion.acc)
	suite.Require().True(coinsEq(assertion.rewards, rewards))
}

type AllRewardsAssertion struct {
	acc     sdk.AccAddress
	rewards sdk.Coins
}

func (ara AllRewardsAssertion) Do(suite *KeeperTestSuite) {
	fmt.Printf("AllRewardsAssertion(%s, %s)\n", ara.acc, ara.rewards)
	rewards := suite.keeper.AllRewards(suite.ctx, ara.acc)
	suite.Require().True(coinsEq(ara.rewards, rewards))
}

func (suite *KeeperTestSuite) TestSimulation() {
	suite.ctx = suite.ctx.WithBlockTime(types.ParseTime("2021-09-01T00:00:00Z"))

	for _, plan := range []types.PlanI{
		types.NewFixedAmountPlan(
			types.NewBasePlan(
				1,
				"",
				types.PlanTypePrivate,
				suite.addrs[0].String(),
				suite.addrs[0].String(),
				sdk.NewDecCoins(
					sdk.NewDecCoinFromDec(denom1, sdk.NewDecWithPrec(3, 1)), // 30%
					sdk.NewDecCoinFromDec(denom2, sdk.NewDecWithPrec(7, 1)), // 70%
				),
				types.ParseTime("0001-01-01T00:00:00Z"),
				types.ParseTime("9999-12-31T00:00:00Z"),
			),
			sdk.NewCoins(sdk.NewInt64Coin(denom3, 1000000)),
		),
	} {
		suite.keeper.SetPlan(suite.ctx, plan)
	}

	addrs := chain.AddTestAddrs(suite.app, suite.ctx, 2, sdk.ZeroInt())
	for _, addr := range addrs {
		err := chain.FundAccount(suite.app.BankKeeper, suite.ctx, addr, sdk.NewCoins(
			sdk.NewInt64Coin(denom1, 1_000_000_000_000),
			sdk.NewInt64Coin(denom2, 1_000_000_000_000)))
		suite.Require().NoError(err)
	}

	for i, action := range []Action{
		BalanceAssertion{addrs[0], denom3, sdk.ZeroInt()},
		AllRewardsAssertion{addrs[0], sdk.NewCoins()},
		BalanceAssertion{addrs[1], denom3, sdk.ZeroInt()},
		AllRewardsAssertion{addrs[1], sdk.NewCoins()},

		StakeAction{addrs[0], sdk.NewCoins(sdk.NewInt64Coin(denom1, 1000000))},
		StakeAction{addrs[1], sdk.NewCoins(sdk.NewInt64Coin(denom1, 500000), sdk.NewInt64Coin(denom2, 500000))},
		AdvanceDayAction{},
		BalanceAssertion{addrs[0], denom3, sdk.ZeroInt()},
		AllRewardsAssertion{addrs[0], sdk.NewCoins()},
		BalanceAssertion{addrs[1], denom3, sdk.ZeroInt()},
		AllRewardsAssertion{addrs[1], sdk.NewCoins()},

		AdvanceDayAction{},
		BalanceAssertion{addrs[0], denom3, sdk.ZeroInt()},
		AllRewardsAssertion{addrs[0], sdk.NewCoins(sdk.NewInt64Coin(denom3, 200000))}, // 300000 * 2/3
		BalanceAssertion{addrs[1], denom3, sdk.ZeroInt()},
		AllRewardsAssertion{addrs[1], sdk.NewCoins(sdk.NewInt64Coin(denom3, 800000))}, // 300000 * 1/3 + 700000

		StakeAction{addrs[0], sdk.NewCoins(sdk.NewInt64Coin(denom1, 500000))},
		AdvanceDayAction{},
		BalanceAssertion{addrs[0], denom3, sdk.NewInt(0)},
		AllUnharvestedRewardsAssertion{addrs[0], utils.ParseCoins("400000denom3")},
		AllRewardsAssertion{addrs[0], sdk.NewCoins()},
		BalanceAssertion{addrs[1], denom3, sdk.ZeroInt()},
		AllRewardsAssertion{addrs[1], sdk.NewCoins(sdk.NewInt64Coin(denom3, 1600000))},

		// User can unstake multiple times before the end of the epoch
		UnstakeAction{addrs[1], sdk.NewCoins(sdk.NewInt64Coin(denom1, 200000), sdk.NewInt64Coin(denom2, 200000))},
		UnstakeAction{addrs[1], sdk.NewCoins(sdk.NewInt64Coin(denom1, 50000), sdk.NewInt64Coin(denom2, 50000))},
		// 250000denom1, 250000denom2
		BalanceAssertion{addrs[1], denom3, sdk.NewInt(0)},
		AllUnharvestedRewardsAssertion{addrs[1], utils.ParseCoins("1600000denom3")},
		AllRewardsAssertion{addrs[1], sdk.NewCoins()},
		AdvanceDayAction{},
		AllUnharvestedRewardsAssertion{addrs[0], utils.ParseCoins("400000denom3")},
		AllRewardsAssertion{addrs[0], sdk.NewCoins(sdk.NewInt64Coin(denom3, 257142))}, // 300000 * (6/7)
		AllUnharvestedRewardsAssertion{addrs[1], utils.ParseCoins("1600000denom3")},
		AllRewardsAssertion{addrs[1], sdk.NewCoins(sdk.NewInt64Coin(denom3, 742857))}, // 300000 * (1/7) + 700000
		// 1000000 => 999999

		// User can harvest multiple times, and it does not affect the rewards
		HarvestAction{addrs[0], []string{denom1}},
		HarvestAction{addrs[1], []string{denom1, denom2}},
		HarvestAction{addrs[1], []string{denom1, denom2}},
		HarvestAction{addrs[1], []string{denom1, denom2}},
		BalanceAssertion{addrs[0], denom3, sdk.NewInt(657142)},
		BalanceAssertion{addrs[1], denom3, sdk.NewInt(2342857)},
		AllRewardsAssertion{addrs[0], sdk.NewCoins()},
		AdvanceDayAction{},
		AllRewardsAssertion{addrs[0], sdk.NewCoins(sdk.NewInt64Coin(denom3, 257142))},
		AllRewardsAssertion{addrs[1], sdk.NewCoins(sdk.NewInt64Coin(denom3, 742857))},
	} {
		suite.Run(fmt.Sprintf("%d", i), func() {
			action.Do(suite)
		})
	}
}
