package keeper

import (
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
	targetMap := map[string]sdk.Dec{}
	for _, val := range activeVals {
		if val.Status == 1 {
			targetWeight = sdk.OneDec()
		} else {
			targetWeight = sdk.ZeroDec()
		}
		targetMap[val.OperatorAddress] = totalLiquidTokens.ToDec().Mul(targetWeight).Quo(totalWeight)
	}

	for i := 0; i < len(activeVals); i++ {
		maxVal, minVal, amountNeeded := activeVals.MinMaxGap(targetMap)
		if amountNeeded.LT(threshold) {
			break
		} else {
			for idx := range activeVals {
				if activeVals[idx].OperatorAddress == maxVal.OperatorAddress {
					activeVals[idx].LiquidTokens = activeVals[idx].LiquidTokens.Add(amountNeeded.TruncateInt())
				}
				if activeVals[idx].OperatorAddress == minVal.OperatorAddress {
					activeVals[idx].LiquidTokens = activeVals[idx].LiquidTokens.Sub(amountNeeded.TruncateInt())
				}
			}

		}
	}

	// time, _ := k.stakingKeeper.BeginRedelegation(ctx, moduleAcc, maxVal.GetOperator(), minVal.GetOperator(), amountNeeded)
	// coins, _ := k.stakingKeeper.CompleteRedelegation(ctx, moduleAcc, maxVal.GetOperator(), minVal.GetOperator())

	//for _, val := range activeVals {
	//	//TODO: add rebalancing logic

	//val.Weight
	//val.OperatorAddress
	//val.LiquidTokens
	//val.Status

	return activeVals
}
