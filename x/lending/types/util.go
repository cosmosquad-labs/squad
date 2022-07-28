package types

import (
	"strings"

	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/cosmos/cosmos-sdk/types/address"
)

const LendingAssetReserveAddressPrefix = "LendingAssetReserveAddress"

// DeriveBondDenom derives and returns the bond denom for the lending asset denom.
func DeriveBondDenom(lendingAssetDenom string) string {
	return "lending/" + lendingAssetDenom // TODO: use different denom prefix?
}

// DeriveLendingAssetReserveAddress derives and returns the reserve address of
// a lending asset.
func DeriveLendingAssetReserveAddress(lendingAssetDenom string) sdk.AccAddress {
	// TODO: use DeriveAddress like other modules?
	return address.Module(ModuleName, []byte(strings.Join([]string{
		LendingAssetReserveAddressPrefix,
		lendingAssetDenom,
	}, "|")))
}
