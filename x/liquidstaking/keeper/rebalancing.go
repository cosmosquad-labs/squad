package keeper

import (
	"time"

	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/cosmosquad-labs/squad/x/liquidstaking/types"
)

func (k Keeper) GetProxyAccBalance(ctx sdk.Context, proxyAcc sdk.AccAddress) (balance sdk.Int) {
	return k.bankKeeper.GetBalance(ctx, proxyAcc, k.stakingKeeper.BondDenom(ctx)).Amount
}

// TryRedelegation attempts redelegation, which is applied only when successful through cached context because there is a constraint that fails if already receiving redelegation.
func (k Keeper) TryRedelegation(ctx sdk.Context, re types.Redelegation, last bool) (completionTime time.Time, err error) {
	cachedCtx, writeCache := ctx.CacheContext()
	srcVal := re.SrcValidator.GetOperator()
	dstVal := re.DstValidator.GetOperator()
	// calculate delShares from tokens with validation
	shares, err := k.stakingKeeper.ValidateUnbondAmount(
		cachedCtx, re.Delegator, srcVal, re.Amount,
	)
	if err != nil {
		return time.Time{}, err
	}

	// when last, full redelegation of shares from delegation
	if last {
		shares = re.SrcValidator.GetDelShares(ctx, k.stakingKeeper)
	}
	completionTime, err = k.stakingKeeper.BeginRedelegation(cachedCtx, re.Delegator, srcVal, dstVal, shares)
	if err != nil {
		return time.Time{}, err
	}
	writeCache()
	return completionTime, nil
}

// DivideByCurrentWeight divide the input value by the ratio of the weight of the liquid validator's liquid token and return it with crumb
// which is may occur while dividing according to the weight of liquid validators by decimal error.
func (k Keeper) DivideByCurrentWeight(ctx sdk.Context, avs types.ActiveLiquidValidators, input sdk.Dec) (outputs []sdk.Dec, crumb sdk.Dec) {
	totalLiquidTokens := avs.TotalLiquidTokens(ctx, k.stakingKeeper)
	if !totalLiquidTokens.IsPositive() {
		return []sdk.Dec{}, sdk.ZeroDec()
	}
	totalOutput := sdk.ZeroDec()
	unitInput := input.QuoTruncate(totalLiquidTokens.ToDec())
	for _, val := range avs {
		output := unitInput.MulTruncate(val.GetLiquidTokens(ctx, k.stakingKeeper).ToDec())
		totalOutput = totalOutput.Add(output)
		outputs = append(outputs, output)
	}
	return outputs, input.Sub(totalOutput)
}

// Rebalancing argument liquidVals containing ValidatorStatusActive which is containing just added on whitelist(liquidToken 0) and ValidatorStatusInActive to delist
func (k Keeper) Rebalancing(ctx sdk.Context, proxyAcc sdk.AccAddress, liquidVals types.LiquidValidators, whitelistedValMap types.WhitelistedValMap, rebalancingTrigger sdk.Dec) (redelegations []types.Redelegation) {
	logger := k.Logger(ctx)
	totalLiquidTokens := liquidVals.TotalLiquidTokens(ctx, k.stakingKeeper)
	if !totalLiquidTokens.IsPositive() {
		return []types.Redelegation{}
	}

	weightMap, totalWeight := k.GetWeightMap(ctx, liquidVals, whitelistedValMap)

	// no active liquid validators
	if !totalWeight.IsPositive() {
		return []types.Redelegation{}
	}

	// calculate rebalancing target map
	targetMap := map[string]sdk.Int{}
	totalTargetMap := sdk.ZeroInt()
	for _, val := range liquidVals {
		targetMap[val.OperatorAddress] = totalLiquidTokens.Mul(weightMap[val.OperatorAddress]).Quo(totalWeight)
		totalTargetMap = totalTargetMap.Add(targetMap[val.OperatorAddress])
	}
	crumb := totalLiquidTokens.Sub(totalTargetMap)
	if !totalTargetMap.IsPositive() {
		return []types.Redelegation{}
	}
	// crumb to first non zero liquid validator
	for _, val := range liquidVals {
		if targetMap[val.OperatorAddress].IsPositive() {
			targetMap[val.OperatorAddress] = targetMap[val.OperatorAddress].Add(crumb)
			break
		}
	}

	lenLiquidVals := liquidVals.Len()
	rebalancingThreshold := rebalancingTrigger.Mul(totalLiquidTokens.ToDec()).TruncateInt()
	for i := 0; i < lenLiquidVals; i++ {
		minVal, maxVal, amountNeeded, last := liquidVals.MinMaxGap(ctx, k.stakingKeeper, targetMap, rebalancingThreshold, crumb)
		if amountNeeded.IsZero() || amountNeeded.LT(rebalancingThreshold) {
			break
		}

		redelegation := types.Redelegation{
			Delegator:    proxyAcc,
			SrcValidator: maxVal,
			DstValidator: minVal,
			Amount:       amountNeeded,
		}
		redelegations = append(redelegations, redelegation)
		_, err := k.TryRedelegation(ctx, redelegation, last)
		if err != nil {
			logger.Error("rebalancing failed due to redelegation restriction", "redelegations", redelegations, "error", err)
		}
	}
	return redelegations
}

// WithdrawRewardsAndReStaking withdraw rewards and re-staking when over threshold
func (k Keeper) WithdrawRewardsAndReStaking(ctx sdk.Context, whitelistedValMap types.WhitelistedValMap) {
	activeVals := k.GetActiveLiquidValidators(ctx, whitelistedValMap)
	totalLiquidTokens := activeVals.TotalLiquidTokens(ctx, k.stakingKeeper)
	// skip when invalid totalLiquidTokens or totalWeight
	if !totalLiquidTokens.IsPositive() || !activeVals.TotalWeight(whitelistedValMap).IsPositive() {
		return
	}
	// Withdraw rewards of LiquidStakingProxyAcc and re-staking
	totalRewards, _, _ := k.CheckTotalRewards(ctx, types.LiquidStakingProxyAcc)
	// checking over types.RewardTrigger and execute GetRewards
	balance := k.GetProxyAccBalance(ctx, types.LiquidStakingProxyAcc)
	rewardsThreshold := types.RewardTrigger.Mul(totalLiquidTokens.ToDec())
	if balance.ToDec().Add(totalRewards).GTE(rewardsThreshold) {
		// re-staking with balance, due to auto-withdraw on add staking by f1
		k.WithdrawLiquidRewards(ctx, types.LiquidStakingProxyAcc)
		balance = k.GetProxyAccBalance(ctx, types.LiquidStakingProxyAcc)
		_, err := k.LiquidDelegate(ctx, types.LiquidStakingProxyAcc, activeVals, balance, whitelistedValMap)
		if err != nil {
			panic(err)
		}
	}
}

func (k Keeper) UpdateLiquidValidatorSet(ctx sdk.Context) []types.Redelegation {
	params := k.GetParams(ctx)
	liquidValidators := k.GetAllLiquidValidators(ctx)
	liquidValsMap := liquidValidators.Map()
	whitelistedValMap := types.GetWhitelistedValMap(params.WhitelistedValidators)

	// Set Liquid validators for added whitelist validators
	for _, wv := range params.WhitelistedValidators {
		if _, ok := liquidValsMap[wv.ValidatorAddress]; !ok {
			lv := types.LiquidValidator{
				OperatorAddress: wv.ValidatorAddress,
			}
			if k.ActiveCondition(ctx, lv, true) {
				k.SetLiquidValidator(ctx, lv)
				liquidValidators = append(liquidValidators, lv)
			}
		}
	}

	// rebalancing based updated liquid validators status with threshold, try by cachedCtx
	// tombstone status also handled on Rebalancing
	reds := k.Rebalancing(ctx, types.LiquidStakingProxyAcc, liquidValidators, whitelistedValMap, types.RebalancingTrigger)

	// remove inactive with zero liquidToken liquidvalidator
	for _, lv := range liquidValidators {
		if !k.ActiveCondition(ctx, lv, whitelistedValMap.IsListed(lv.OperatorAddress)) && !lv.GetDelShares(ctx, k.stakingKeeper).IsPositive() {
			k.RemoveLiquidValidator(ctx, lv)
		}
	}

	// withdraw rewards and re-staking when over threshold
	k.WithdrawRewardsAndReStaking(ctx, whitelistedValMap)
	return reds
}

//// Deprecated: AddStakingTargetMap is make add staking target map for one-way rebalancing, it can be called recursively, not using on this version for simplify.
//func (k Keeper) AddStakingTargetMap(ctx sdk.Context, activeVals types.ActiveLiquidValidators, addStakingAmt sdk.Int) map[string]sdk.Int {
//	targetMap := make(map[string]sdk.Int)
//	if addStakingAmt.IsNil() || !addStakingAmt.IsPositive() || activeVals.Len() == 0 {
//		return targetMap
//	}
//	params := k.GetParams(ctx)
//	whitelistedValMap := types.GetWhitelistedValMap(params.WhitelistedValidators)
//	totalLiquidTokens := activeVals.TotalLiquidTokens(ctx, k.stakingKeeper)
//	totalWeight := activeVals.TotalWeight(whitelistedValMap)
//	ToBeTotalDelShares := totalLiquidTokens.TruncateInt().Add(addStakingAmt)
//	existOverWeightedVal := false
//
//	sharePerWeight := ToBeTotalDelShares.Quo(totalWeight)
//	crumb := ToBeTotalDelShares.Sub(sharePerWeight.Mul(totalWeight))
//
//	i := 0
//	for _, val := range activeVals {
//		weightedShare := val.GetWeight(whitelistedValMap, true).Mul(sharePerWeight)
//		if val.GetLiquidTokens(ctx, k.stakingKeeper).TruncateInt().GT(weightedShare) {
//			existOverWeightedVal = true
//		} else {
//			activeVals[i] = val
//			i++
//			targetMap[val.OperatorAddress] = weightedShare.Sub(val.GetLiquidTokens(ctx, k.stakingKeeper).TruncateInt())
//		}
//	}
//	// remove overWeightedVals for recursive call
//	activeVals = activeVals[:i]
//
//	if !existOverWeightedVal {
//		if v, ok := targetMap[activeVals[0].OperatorAddress]; ok {
//			targetMap[activeVals[0].OperatorAddress] = v.Add(crumb)
//		} else {
//			targetMap[activeVals[0].OperatorAddress] = crumb
//		}
//		return targetMap
//	} else {
//		return k.AddStakingTargetMap(ctx, activeVals, addStakingAmt)
//	}
//}
