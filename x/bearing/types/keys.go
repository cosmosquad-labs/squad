package types

import (
	"bytes"
)

const (
	// ModuleName is the name of the bearing module
	ModuleName = "bearing"

	// RouterKey is the message router key for the bearing module
	RouterKey = ModuleName

	// StoreKey is the default store key for the bearing module
	StoreKey = ModuleName

	// QuerierRoute is the querier route for the bearing module
	QuerierRoute = ModuleName
)

var (
	// Keys for store prefixes
	TotalCollectedCoinsKeyPrefix = []byte{0x11}
)

// GetTotalCollectedCoinsKey creates the key for the total collected coins for a bearing.
func GetTotalCollectedCoinsKey(bearingName string) []byte {
	return append(TotalCollectedCoinsKeyPrefix, []byte(bearingName)...)
}

// ParseTotalCollectedCoinsKey parses the total collected coins key and returns the bearing name.
func ParseTotalCollectedCoinsKey(key []byte) (bearingName string) {
	if !bytes.HasPrefix(key, TotalCollectedCoinsKeyPrefix) {
		panic("key does not have proper prefix")
	}
	return string(key[1:])
}
