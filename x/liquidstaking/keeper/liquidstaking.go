package keeper

import (
	"fmt"
	"time"

	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
	stakingtypes "github.com/cosmos/cosmos-sdk/x/staking/types"

	"github.com/cosmosquad-labs/squad/x/liquidstaking/types"
)

func (k Keeper) LiquidBondDenom(ctx sdk.Context) (res string) {
	k.paramSpace.Get(ctx, types.KeyLiquidBondDenom, &res)
	return
}

// NetAmountState calculates the sum of bondedDenom balance, total delegation tokens(slash applied LiquidTokens), total remaining reward of types.LiquidStakingProxyAcc
// During Liquid Unstacking, btoken immediately burns and the unbonding queue belongs to the requester, so the liquid staker's unbonding values are excluded on netAmount
// It is used only for calculation and query and is not stored in kv.
func (k Keeper) NetAmountState(ctx sdk.Context) (nas types.NetAmountState) {
	nas.ProxyAccBalance = k.bankKeeper.GetBalance(ctx, types.LiquidStakingProxyAcc, k.stakingKeeper.BondDenom(ctx)).Amount

	totalRemainingRewards, totalDelShares, totalLiquidTokens := k.CheckDelegationStates(ctx, types.LiquidStakingProxyAcc)
	nas.TotalRemainingRewards = totalRemainingRewards
	nas.TotalDelShares = totalDelShares
	nas.TotalLiquidTokens = totalLiquidTokens

	totalUnbondingBalance := sdk.ZeroInt()
	ubds := k.stakingKeeper.GetAllUnbondingDelegations(ctx, types.LiquidStakingProxyAcc)
	for _, ubd := range ubds {
		for _, entry := range ubd.Entries {
			// use Balance(slashing applied) not InitialBalance(without slashing)
			totalUnbondingBalance = totalUnbondingBalance.Add(entry.Balance)
		}
	}
	nas.TotalUnbondingBalance = totalUnbondingBalance
	nas.BtokenTotalSupply = k.bankKeeper.GetSupply(ctx, k.LiquidBondDenom(ctx)).Amount
	nas.NetAmount = nas.CalcNetAmount()
	nas.MintRate = nas.CalcMintRate()
	return
}

// LiquidStaking mints bToken worth of staking coin value according to NetAmount and performs LiquidDelegate.
func (k Keeper) LiquidStaking(
	ctx sdk.Context, proxyAcc, liquidStaker sdk.AccAddress, stakingCoin sdk.Coin) (newShares sdk.Dec, bTokenMintAmount sdk.Int, err error) {

	// check bond denomination
	bondDenom := k.stakingKeeper.BondDenom(ctx)
	if stakingCoin.Denom != bondDenom {
		return sdk.ZeroDec(), bTokenMintAmount, sdkerrors.Wrapf(
			types.ErrInvalidBondDenom, "invalid coin denomination: got %s, expected %s", stakingCoin.Denom, bondDenom,
		)
	}

	params := k.GetParams(ctx)
	whitelistedValMap := types.GetWhitelistedValMap(params.WhitelistedValidators)
	activeVals := k.GetActiveLiquidValidators(ctx, whitelistedValMap)
	if activeVals.Len() == 0 || !activeVals.TotalWeight(whitelistedValMap).IsPositive() {
		return sdk.ZeroDec(), bTokenMintAmount, types.ErrActiveLiquidValidatorsNotExists
	}

	// NetAmount must be calculated before send
	nas := k.NetAmountState(ctx)

	// send staking coin to liquid staking proxy account to proxy delegation, need sufficient spendable balances
	err = k.bankKeeper.SendCoins(ctx, liquidStaker, proxyAcc, sdk.NewCoins(stakingCoin))
	if err != nil {
		return sdk.ZeroDec(), bTokenMintAmount, err
	}

	// mint btoken, MintAmount = TotalSupply * StakeAmount/NetAmount
	liquidBondDenom := k.LiquidBondDenom(ctx)
	bTokenMintAmount = stakingCoin.Amount
	if nas.BtokenTotalSupply.IsPositive() {
		bTokenMintAmount = types.NativeTokenToBToken(stakingCoin.Amount, nas.BtokenTotalSupply, nas.NetAmount)
	}

	if !bTokenMintAmount.IsPositive() {
		return sdk.ZeroDec(), sdk.ZeroInt(), types.ErrTooSmallLiquidStakingAmount
	}

	// mint on module acc and send
	mintCoin := sdk.NewCoins(sdk.NewCoin(liquidBondDenom, bTokenMintAmount))
	err = k.bankKeeper.MintCoins(ctx, types.ModuleName, mintCoin)
	if err != nil {
		return sdk.ZeroDec(), bTokenMintAmount, err
	}
	err = k.bankKeeper.SendCoinsFromModuleToAccount(ctx, types.ModuleName, liquidStaker, mintCoin)
	if err != nil {
		return sdk.ZeroDec(), bTokenMintAmount, err
	}
	newShares, err = k.LiquidDelegate(ctx, proxyAcc, activeVals, stakingCoin.Amount, whitelistedValMap)
	return newShares, bTokenMintAmount, err
}

// LiquidDelegate delegates staking amount to active validators by proxy account.
func (k Keeper) LiquidDelegate(ctx sdk.Context, proxyAcc sdk.AccAddress, activeVals types.ActiveLiquidValidators, stakingAmt sdk.Int, whitelistedValMap types.WhitelistedValMap) (newShares sdk.Dec, err error) {
	totalNewShares := sdk.ZeroDec()
	// crumb may occur due to a decimal point error in dividing the staking amount into the weight of liquid validators, It added on first active liquid validator
	weightedAmt, crumb := types.DivideByWeight(activeVals, stakingAmt, whitelistedValMap)
	if len(weightedAmt) == 0 {
		return sdk.ZeroDec(), types.ErrInvalidActiveLiquidValidators
	}
	weightedAmt[0] = weightedAmt[0].Add(crumb)
	for i, val := range activeVals {
		if !weightedAmt[i].IsPositive() {
			continue
		}
		validator, _ := k.stakingKeeper.GetValidator(ctx, val.GetOperator())
		newShares, err = k.stakingKeeper.Delegate(ctx, proxyAcc, weightedAmt[i], stakingtypes.Unbonded, validator, true)
		if err != nil {
			return sdk.ZeroDec(), err
		}
		totalNewShares = totalNewShares.Add(newShares)
	}
	return totalNewShares, nil
}

// LiquidUnstaking burns unstakingBtoken and performs LiquidUnbond to active liquid validators with del shares worth of shares according to NetAmount with each validators current weight.
func (k Keeper) LiquidUnstaking(
	ctx sdk.Context, proxyAcc, liquidStaker sdk.AccAddress, unstakingBtoken sdk.Coin,
) (time.Time, sdk.Int, []stakingtypes.UnbondingDelegation, sdk.Int, error) {

	// check bond denomination
	params := k.GetParams(ctx)
	liquidBondDenom := k.LiquidBondDenom(ctx)
	if unstakingBtoken.Denom != liquidBondDenom {
		return time.Time{}, sdk.ZeroInt(), []stakingtypes.UnbondingDelegation{}, sdk.ZeroInt(), sdkerrors.Wrapf(
			types.ErrInvalidLiquidBondDenom, "invalid coin denomination: got %s, expected %s", unstakingBtoken.Denom, liquidBondDenom,
		)
	}

	// Get NetAmount states
	nas := k.NetAmountState(ctx)

	if unstakingBtoken.Amount.GT(nas.BtokenTotalSupply) {
		return time.Time{}, sdk.ZeroInt(), []stakingtypes.UnbondingDelegation{}, sdk.ZeroInt(), types.ErrInvalidBTokenSupply
	}

	// UnstakeAmount = NetAmount * BTokenAmount/TotalSupply * (1-UnstakeFeeRate)
	unbondingAmount := types.BTokenToNativeToken(unstakingBtoken.Amount, nas.BtokenTotalSupply, nas.NetAmount)
	unbondingAmount = types.DeductFeeRate(unbondingAmount, params.UnstakeFeeRate)
	unbondingAmountInt := unbondingAmount.TruncateInt()

	if !unbondingAmountInt.IsPositive() {
		return time.Time{}, sdk.ZeroInt(), []stakingtypes.UnbondingDelegation{}, sdk.ZeroInt(), types.ErrTooSmallLiquidUnstakingAmount
	}

	// burn btoken
	err := k.bankKeeper.SendCoinsFromAccountToModule(ctx, liquidStaker, types.ModuleName, sdk.NewCoins(unstakingBtoken))
	if err != nil {
		return time.Time{}, sdk.ZeroInt(), []stakingtypes.UnbondingDelegation{}, sdk.ZeroInt(), err
	}
	err = k.bankKeeper.BurnCoins(ctx, types.ModuleName, sdk.NewCoins(sdk.NewCoin(liquidBondDenom, unstakingBtoken.Amount)))
	if err != nil {
		return time.Time{}, sdk.ZeroInt(), []stakingtypes.UnbondingDelegation{}, sdk.ZeroInt(), err
	}

	liquidVals := k.GetAllLiquidValidators(ctx)
	totalLiquidTokens, liquidTokenMap := liquidVals.TotalLiquidTokens(ctx, k.stakingKeeper, false)

	// if no totalLiquidTokens, withdraw directly from balance of proxy acc
	if !totalLiquidTokens.IsPositive() {
		if nas.ProxyAccBalance.GTE(unbondingAmountInt) {
			err = k.bankKeeper.SendCoins(ctx, types.LiquidStakingProxyAcc, liquidStaker, sdk.NewCoins(sdk.NewCoin(k.stakingKeeper.BondDenom(ctx), unbondingAmountInt)))
			if err != nil {
				return time.Time{}, sdk.ZeroInt(), []stakingtypes.UnbondingDelegation{}, sdk.ZeroInt(), err
			} else {
				return time.Time{}, sdk.ZeroInt(), []stakingtypes.UnbondingDelegation{}, unbondingAmountInt, nil
			}
		} else {
			// error case where there is a quantity that are unbonding balance or remaining rewards that is not re-stake or withdrawn in netAmount.
			return time.Time{}, sdk.ZeroInt(), []stakingtypes.UnbondingDelegation{}, sdk.ZeroInt(), types.ErrInsufficientProxyAccBalance
		}
	}
	// fail when no liquid validators to unbond
	if liquidVals.Len() == 0 {
		return time.Time{}, sdk.ZeroInt(), []stakingtypes.UnbondingDelegation{}, sdk.ZeroInt(), types.ErrLiquidValidatorsNotExists
	}

	// crumb may occur due to a decimal error in dividing the unstaking bToken into the weight of liquid validators, it will remain in the NetAmount
	unbondingAmounts, _ := types.DivideByCurrentWeight(liquidVals, unbondingAmount, totalLiquidTokens, liquidTokenMap)
	if len(unbondingAmounts) == 0 {
		return time.Time{}, sdk.ZeroInt(), []stakingtypes.UnbondingDelegation{}, sdk.ZeroInt(), types.ErrInvalidActiveLiquidValidators
	}
	totalReturnAmount := sdk.ZeroInt()
	var ubdTime time.Time
	var ubds []stakingtypes.UnbondingDelegation
	for i, val := range liquidVals {
		var ubd stakingtypes.UnbondingDelegation
		var returnAmount sdk.Int
		var weightedShare sdk.Dec
		// calculate delShares from tokens with validation
		weightedShare, err = k.stakingKeeper.ValidateUnbondAmount(ctx, proxyAcc, val.GetOperator(), unbondingAmounts[i].TruncateInt())
		if err != nil {
			return time.Time{}, sdk.ZeroInt(), []stakingtypes.UnbondingDelegation{}, sdk.ZeroInt(), err
		}
		if !weightedShare.IsPositive() {
			continue
		}
		// unbond with weightedShare
		ubdTime, returnAmount, ubd, err = k.LiquidUnbond(ctx, proxyAcc, liquidStaker, val.GetOperator(), weightedShare)
		if err != nil {
			return time.Time{}, sdk.ZeroInt(), []stakingtypes.UnbondingDelegation{}, sdk.ZeroInt(), err
		}
		ubds = append(ubds, ubd)
		totalReturnAmount = totalReturnAmount.Add(returnAmount)
	}
	return ubdTime, totalReturnAmount, ubds, sdk.ZeroInt(), nil
}

// LiquidUnbond unbond delegation shares to active validators by proxy account.
func (k Keeper) LiquidUnbond(
	ctx sdk.Context, proxyAcc, liquidStaker sdk.AccAddress, valAddr sdk.ValAddress, shares sdk.Dec,
) (time.Time, sdk.Int, stakingtypes.UnbondingDelegation, error) {
	validator, found := k.stakingKeeper.GetValidator(ctx, valAddr)
	if !found {
		return time.Time{}, sdk.ZeroInt(), stakingtypes.UnbondingDelegation{}, stakingtypes.ErrNoDelegatorForAddress
	}

	returnAmount, err := k.stakingKeeper.Unbond(ctx, proxyAcc, valAddr, shares)
	if err != nil {
		return time.Time{}, sdk.ZeroInt(), stakingtypes.UnbondingDelegation{}, err
	}

	// transfer the validator tokens to the not bonded pool
	if validator.IsBonded() {
		coins := sdk.NewCoins(sdk.NewCoin(k.stakingKeeper.BondDenom(ctx), returnAmount))
		if err = k.bankKeeper.SendCoinsFromModuleToModule(ctx, stakingtypes.BondedPoolName, stakingtypes.NotBondedPoolName, coins); err != nil {
			panic(err)
		}
	}

	completionTime := ctx.BlockHeader().Time.Add(k.stakingKeeper.UnbondingTime(ctx))
	ubd := k.stakingKeeper.SetUnbondingDelegationEntry(ctx, liquidStaker, valAddr, ctx.BlockHeight(), completionTime, returnAmount)
	k.stakingKeeper.InsertUBDQueue(ctx, ubd, completionTime)

	return completionTime, returnAmount, ubd, nil
}

// CheckDelegationStates returns total remaining rewards, delshares, liquid tokens of delegations by proxy account
func (k Keeper) CheckDelegationStates(ctx sdk.Context, proxyAcc sdk.AccAddress) (sdk.Dec, sdk.Dec, sdk.Int) {
	bondDenom := k.stakingKeeper.BondDenom(ctx)
	totalRewards := sdk.ZeroDec()
	totalDelShares := sdk.ZeroDec()
	totalLiquidTokens := sdk.ZeroInt()

	// Cache ctx for calculate rewards
	cachedCtx, _ := ctx.CacheContext()
	k.stakingKeeper.IterateDelegations(
		cachedCtx, proxyAcc,
		func(_ int64, del stakingtypes.DelegationI) (stop bool) {
			valAddr := del.GetValidatorAddr()
			val := k.stakingKeeper.Validator(cachedCtx, valAddr)
			endingPeriod := k.distrKeeper.IncrementValidatorPeriod(cachedCtx, val)
			delReward := k.distrKeeper.CalculateDelegationRewards(cachedCtx, val, del, endingPeriod)
			delShares := del.GetShares()
			if delShares.IsPositive() {
				totalDelShares = totalDelShares.Add(delShares)
				liquidTokens := val.TokensFromSharesTruncated(delShares).TruncateInt()
				totalLiquidTokens = totalLiquidTokens.Add(liquidTokens)
				totalRewards = totalRewards.Add(delReward.AmountOf(bondDenom))
			}
			return false
		},
	)

	return totalRewards, totalDelShares, totalLiquidTokens
}

func (k Keeper) WithdrawLiquidRewards(ctx sdk.Context, proxyAcc sdk.AccAddress) sdk.Int {
	totalRewards := sdk.ZeroInt()
	bondDenom := k.stakingKeeper.BondDenom(ctx)
	k.stakingKeeper.IterateDelegations(
		ctx, proxyAcc,
		func(_ int64, del stakingtypes.DelegationI) (stop bool) {
			valAddr := del.GetValidatorAddr()
			reward, err := k.distrKeeper.WithdrawDelegationRewards(ctx, proxyAcc, valAddr)
			if err != nil {
				panic(err)
			}
			totalRewards = totalRewards.Add(reward.AmountOf(bondDenom))
			return false
		},
	)
	if totalRewards.IsPositive() {
		fmt.Println("[WithdrawLiquidRewards]", totalRewards)
	}
	return totalRewards
}

// GetLiquidValidator get a single liquid validator
func (k Keeper) GetLiquidValidator(ctx sdk.Context, addr sdk.ValAddress) (val types.LiquidValidator, found bool) {
	store := ctx.KVStore(k.storeKey)

	value := store.Get(types.GetLiquidValidatorKey(addr))
	if value == nil {
		return val, false
	}

	val = types.MustUnmarshalLiquidValidator(k.cdc, value)
	return val, true
}

// SetLiquidValidator set the main record holding liquid validator details
func (k Keeper) SetLiquidValidator(ctx sdk.Context, val types.LiquidValidator) {
	store := ctx.KVStore(k.storeKey)
	bz := types.MustMarshalLiquidValidator(k.cdc, &val)
	store.Set(types.GetLiquidValidatorKey(val.GetOperator()), bz)
}

// RemoveLiquidValidator remove a liquid validator on kv store
func (k Keeper) RemoveLiquidValidator(ctx sdk.Context, val types.LiquidValidator) {
	store := ctx.KVStore(k.storeKey)
	store.Delete(types.GetLiquidValidatorKey(val.GetOperator()))
}

// GetAllLiquidValidators get the set of all liquid validators with no limits, used during genesis dump
func (k Keeper) GetAllLiquidValidators(ctx sdk.Context) (vals types.LiquidValidators) {
	store := ctx.KVStore(k.storeKey)

	iterator := sdk.KVStorePrefixIterator(store, types.LiquidValidatorsKey)
	defer iterator.Close()

	for ; iterator.Valid(); iterator.Next() {
		val := types.MustUnmarshalLiquidValidator(k.cdc, iterator.Value())
		vals = append(vals, val)
	}

	return vals
}

// GetActiveLiquidValidators get the set of active liquid validators.
func (k Keeper) GetActiveLiquidValidators(ctx sdk.Context, whitelistedValMap types.WhitelistedValMap) (vals types.ActiveLiquidValidators) {
	store := ctx.KVStore(k.storeKey)
	iterator := sdk.KVStorePrefixIterator(store, types.LiquidValidatorsKey)
	defer iterator.Close()

	for ; iterator.Valid(); iterator.Next() {
		val := types.MustUnmarshalLiquidValidator(k.cdc, iterator.Value())
		if k.ActiveCondition(ctx, val, whitelistedValMap.IsListed(val.OperatorAddress)) {
			vals = append(vals, val)
		}
	}
	return vals
}

func (k Keeper) GetAllLiquidValidatorStates(ctx sdk.Context) (liquidValidatorStates []types.LiquidValidatorState) {
	lvs := k.GetAllLiquidValidators(ctx)
	whitelistedValMap := k.GetParams(ctx).WhitelistedValMap()
	for _, lv := range lvs {
		active := k.ActiveCondition(ctx, lv, whitelistedValMap.IsListed(lv.OperatorAddress))
		lvState := types.LiquidValidatorState{
			OperatorAddress: lv.OperatorAddress,
			Weight:          lv.GetWeight(whitelistedValMap, active),
			Status:          lv.GetStatus(active),
			DelShares:       lv.GetDelShares(ctx, k.stakingKeeper),
			LiquidTokens:    lv.GetLiquidTokens(ctx, k.stakingKeeper, false),
		}
		liquidValidatorStates = append(liquidValidatorStates, lvState)
	}
	return
}

func (k Keeper) GetLiquidValidatorState(ctx sdk.Context, addr sdk.ValAddress) (liquidValidatorState types.LiquidValidatorState, found bool) {
	lv, found := k.GetLiquidValidator(ctx, addr)
	if !found {
		return types.LiquidValidatorState{
			OperatorAddress: addr.String(),
			Weight:          sdk.ZeroInt(),
			Status:          types.ValidatorStatusUnspecified,
			DelShares:       sdk.ZeroDec(),
			LiquidTokens:    sdk.ZeroInt(),
		}, false
	}
	whitelistedValMap := k.GetParams(ctx).WhitelistedValMap()
	active := k.ActiveCondition(ctx, lv, whitelistedValMap.IsListed(lv.OperatorAddress))
	return types.LiquidValidatorState{
		OperatorAddress: lv.OperatorAddress,
		Weight:          lv.GetWeight(whitelistedValMap, active),
		Status:          lv.GetStatus(active),
		DelShares:       lv.GetDelShares(ctx, k.stakingKeeper),
		LiquidTokens:    lv.GetLiquidTokens(ctx, k.stakingKeeper, false),
	}, true
}

// Name suggestion: IsActiveLiquidValidator
// Remove whitelisted parameter, and use that value directly in the caller.
// For example, if you're gonna call this function with whitelisted = false,
// then you can write code like this:
// if whitelisted && k.ActiveCondition(ctx, v) { ... }
func (k Keeper) ActiveCondition(ctx sdk.Context, v types.LiquidValidator, whitelisted bool) bool {
	val, found := k.stakingKeeper.GetValidator(ctx, v.GetOperator())
	if !found {
		return false
	}
	// This can be simplified to:
	//consPk, err := val.ConsPubKey()
	//if err != nil {
	//	panic(err)
	//}
	//tombstoned := k.slashingKeeper.IsTombstoned(ctx, sdk.ConsAddress(consPk.Address()))
	//return whitelisted &&
	//	!tombstoned &&
	//	// !Unspecified ==> Bonded, Unbonding, Unbonded
	//	val.GetStatus() != stakingtypes.Unspecified &&
	//	!val.GetTokens().IsNil() &&
	//	!val.GetDelegatorShares().IsNil() &&
	//	!val.InvalidExRate()
	tombstoned := v.IsTombstoned(ctx, k.stakingKeeper, k.slashingKeeper)
	return types.ActiveCondition(val, whitelisted, tombstoned)
}

func (k Keeper) GetWeightMap(ctx sdk.Context, liquidVals types.LiquidValidators, whitelistedValMap types.WhitelistedValMap) (map[string]sdk.Int, sdk.Int) {
	weightMap := make(map[string]sdk.Int)
	totalWeight := sdk.ZeroInt()
	for _, val := range liquidVals {
		weight := val.GetWeight(whitelistedValMap, k.ActiveCondition(ctx, val, whitelistedValMap.IsListed(val.OperatorAddress)))
		totalWeight = totalWeight.Add(weight)
		weightMap[val.OperatorAddress] = weight
	}
	return weightMap, totalWeight
}
