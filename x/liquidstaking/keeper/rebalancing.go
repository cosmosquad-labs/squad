package keeper

import (
	"fmt"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/tendermint/farming/x/liquidstaking/types"
)

func (k Keeper) UpdateLiquidValidators(ctx sdk.Context) {
	//liquidVals := k.GetAllLiquidValidators()
	// TODO: GET, SET, GETALL, Iterate LiquidValidators, indexing
	//for _, val := range liquidVals {
	//	if val.
	//}
	//k.stakingKeeper.GetLastTotalPower()
}

// activeVals containing ValidatorStatusWhiteListed which is containing just added on whitelist(power 0) and ValidatorStatusDelisting
func (k Keeper) Rebalancing(ctx sdk.Context, moduleAcc sdk.AccAddress, activeVals types.LiquidValidators, threshold sdk.Dec) (rebalancedLiquidVals types.LiquidValidators) {
	totalLiquidTokens := sdk.ZeroInt()
	totalWeight := sdk.ZeroDec()
	for _, val := range activeVals {
		totalLiquidTokens = totalLiquidTokens.Add(val.LiquidTokens)
		if val.Status == 1 {
			totalWeight = totalWeight.Add(val.Weight)
		}
	}

	var targetWeight sdk.Dec
	targetMap := map[string]sdk.Int{}
	for _, val := range activeVals {
		if val.Status == 1 {
			targetWeight = val.Weight
		} else {
			targetWeight = sdk.ZeroDec()
		}
		targetMap[val.OperatorAddress] = totalLiquidTokens.ToDec().MulTruncate(targetWeight).QuoTruncate(totalWeight).TruncateInt()
	}
	fmt.Println(targetMap)

	for i := 0; i < len(activeVals); i++ {
		maxVal, minVal, amountNeeded := activeVals.MinMaxGap(targetMap)
		if amountNeeded.IsZero() || (i == 0 && amountNeeded.LT(threshold.TruncateInt())) {
			break
		}
		for idx := range activeVals {
			if activeVals[idx].OperatorAddress == maxVal.OperatorAddress {
				activeVals[idx].LiquidTokens = activeVals[idx].LiquidTokens.Add(amountNeeded)
			}
			if activeVals[idx].OperatorAddress == minVal.OperatorAddress {
				activeVals[idx].LiquidTokens = activeVals[idx].LiquidTokens.Sub(amountNeeded)
			}
		}
	}
	fmt.Println(activeVals)

	return activeVals
}

func (k Keeper) ProcessStaking(moduleAcc sdk.AccAddress, activeVals types.LiquidValidators, addStakingTokens sdk.Int, unstakingTokens sdk.Int) (rebalancedLiquidVals types.LiquidValidators) {
	// suppose that when unstaking process starts, the required amount of unstaking is transferred to the moduleAcc
	// and when the unstaking queue matures, finally the tokens are transferred to staker's address
	// addStakingTokens : additional staking amount to be considered when rebalancing
	// unstakingTokens : unstaking amount to be considered when rebalancing

	type accountBalance struct {
		Address      string
		LiquidTokens sdk.Int
	}
	moduleAccBalance := accountBalance{
		Address:      moduleAcc.String(),
		LiquidTokens: sdk.ZeroInt(),
	}
	// temporary struct for tokens in moduleAcc (require to fix)

	if addStakingTokens.GT(unstakingTokens) {
		moduleAccBalance.LiquidTokens = moduleAccBalance.LiquidTokens.Add(unstakingTokens)
		addStakingTokens = addStakingTokens.Sub(unstakingTokens) // must change to send action
		distrbutionAmount := addStakingTokens.Quo(sdk.NewInt(int64(len(activeVals))))
		for _, val := range activeVals {
			if val.Status == 1 {
				val.LiquidTokens = val.LiquidTokens.Add(distrbutionAmount)
				addStakingTokens = addStakingTokens.Sub(distrbutionAmount) // must change to delegate action
			}
		}
	} else {
		moduleAccBalance.LiquidTokens = moduleAccBalance.LiquidTokens.Add(addStakingTokens)
		unstakingTokens = unstakingTokens.Sub(addStakingTokens) // must change to send action

	}
	fmt.Println(moduleAccBalance)

	return
}
