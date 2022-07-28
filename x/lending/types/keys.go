package types

import (
	utils "github.com/cosmosquad-labs/squad/v2/types"
)

const (
	// ModuleName defines the module name
	ModuleName = "lending"

	// StoreKey defines the primary module store key
	StoreKey = ModuleName

	// RouterKey is the message route for slashing
	RouterKey = ModuleName

	// QuerierRoute defines the module's query routing key
	QuerierRoute = ModuleName
)

var (
	LendingAssetKeyPrefix = []byte{0xd0}
)

// GetLendingAssetKey returns the store key for LendingAsset with denom.
func GetLendingAssetKey(denom string) []byte {
	return append(LendingAssetKeyPrefix, utils.LengthPrefixString(denom)...)
}
