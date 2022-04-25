package types

import (
	fmt "fmt"

	sdk "github.com/cosmos/cosmos-sdk/types"
)

func (a Airdrop) GetSourceAddress() sdk.AccAddress {
	addr, err := sdk.AccAddressFromBech32(a.SourceAddress)
	if err != nil {
		panic(err)
	}
	return addr
}

func (r ClaimRecord) GetRecipient() sdk.AccAddress {
	addr, err := sdk.AccAddressFromBech32(r.Recipient)
	if err != nil {
		panic(err)
	}
	return addr
}

// GetClaimableCoinsForCondition uses unclaimed # of conditions as divisor to
// calculate a proportionate claimable amount of coins for the condition.
func (r ClaimRecord) GetClaimableCoinsForCondition(airdropConditions []ConditionType) sdk.Coins {
	conditionSet := map[ConditionType]struct{}{}
	for _, ac := range airdropConditions {
		conditionSet[ac] = struct{}{}
	}
	for _, c := range r.ClaimedConditions {
		delete(conditionSet, c)
	}
	unclaimedNum := sdk.NewInt(int64(len(conditionSet)))

	claimableCoins := sdk.Coins{}
	for _, c := range r.ClaimableCoins {
		claimableAmt := c.Amount.Quo(unclaimedNum)
		claimableCoins = claimableCoins.Add(sdk.NewCoin(c.Denom, claimableAmt))
	}
	return claimableCoins
}

func CodeQLAnalysisTest() {
	// Directly using the bech32 constants
	fmt.Println(sdk.Bech32MainPrefix)
	fmt.Println(sdk.Bech32PrefixValAddr)
	fmt.Println(sdk.Bech32PrefixConsAddr)

	// Iteration over map
	delegatorByAmount := map[string]sdk.Int{}
	delegatorByAmount["cosmos1"] = sdk.NewInt(100000)
	delegatorByAmount["cosmos2"] = sdk.NewInt(200000)
	delegatorByAmount["cosmos3"] = sdk.NewInt(300000)

	for _, amount := range delegatorByAmount {
		fmt.Println(amount)
	}

	// Floating point arithmetic
	result := float64(1) - float64(0.05)
	fmt.Println(result)
}
