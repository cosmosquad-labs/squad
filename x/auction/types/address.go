package types

import (
	"fmt"

	sdk "github.com/cosmos/cosmos-sdk/types"

	farmingtypes "github.com/cosmosquad-labs/squad/x/farming/types"
)

const (
	SellingReserveAddressPrefix string = "SellingReserveAddress"
	PayingReserveAddressPrefix  string = "PayingReserveAddress"
	VestingReserveAddressPrefix string = "VestingReserveAddress"
	ModuleAddressNameSplitter   string = "|"

	// ReserveAddressType is an address type of reserve for selling, paying, and vesting.
	// The module uses the address type of 32 bytes length, but it can be changed depending on Cosmos SDK's direction.
	ReserveAddressType = farmingtypes.AddressType32Bytes
)

// SellingReserveAddress returns the selling reserve address with the given auction id.
func SellingReserveAddress(auctionId uint64) sdk.AccAddress {
	return farmingtypes.DeriveAddress(ReserveAddressType, ModuleName, SellingReserveAddressPrefix+ModuleAddressNameSplitter+fmt.Sprint(auctionId))
}

// PayingReserveAddress returns the paying reserve address with the given auction id.
func PayingReserveAddress(auctionId uint64) sdk.AccAddress {
	return farmingtypes.DeriveAddress(ReserveAddressType, ModuleName, PayingReserveAddressPrefix+ModuleAddressNameSplitter+fmt.Sprint(auctionId))
}

// VestingReserveAddress returns the vesting reserve address with the given auction id.
func VestingReserveAddress(auctionId uint64) sdk.AccAddress {
	return farmingtypes.DeriveAddress(ReserveAddressType, ModuleName, VestingReserveAddressPrefix+ModuleAddressNameSplitter+fmt.Sprint(auctionId))
}
