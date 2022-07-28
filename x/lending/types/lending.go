package types

import (
	sdk "github.com/cosmos/cosmos-sdk/types"
)

// NewLendingAsset returns a newly initialized LendingAsset.
func NewLendingAsset(denom string) LendingAsset {
	return LendingAsset{
		Denom:                 denom,
		BondDenom:             DeriveBondDenom(denom),
		ReserveAddress:        DeriveLendingAssetReserveAddress(denom).String(),
		TotalLentAmount:       sdk.ZeroInt(),
		TotalBorrowedAmount:   sdk.ZeroInt(),
		AccruedInterestAmount: sdk.ZeroDec(),
	}
}

// GetReserveAddress is a convenient helper for getting the reserve address as
// sdk.AccAddress.
func (asset LendingAsset) GetReserveAddress() sdk.AccAddress {
	addr, err := sdk.AccAddressFromBech32(asset.ReserveAddress)
	if err != nil {
		panic(err)
	}
	return addr
}

func CalculateMintingBondAmount(lendingAmt, totalLentAmt sdk.Int, accruedInterestAmt sdk.Dec, bondSupply sdk.Int) sdk.Int {
	if totalLentAmt.IsZero() {
		return lendingAmt
	}
	// exchangeRatio = bondSupply / (totalLentAmt + accruedInterestAmt)
	exchangeRatio := sdk.NewDecFromInt(bondSupply).QuoTruncate(sdk.NewDecFromInt(totalLentAmt).Add(accruedInterestAmt))
	return exchangeRatio.MulInt(lendingAmt).TruncateInt()
}

func CalculateWithdrawingLendingAssetAmount() sdk.Int {
	panic("not implemented")
}
