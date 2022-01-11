package types

import (
	sdk "github.com/cosmos/cosmos-sdk/types"
)

const (
	// ModuleName defines the module name
	ModuleName = "liquidity"

	// StoreKey defines the primary module store key
	StoreKey = ModuleName

	// RouterKey is the message route for slashing
	RouterKey = ModuleName

	// QuerierRoute defines the module's query routing key
	QuerierRoute = ModuleName
)

var (
	PairIdKey = []byte{0xa0} // key for the lastest pair id
	PoolIdKey = []byte{0xa1} // key for the latest pool id

	PairKeyPrefix             = []byte{0xa5}
	PairIndexKeyPrefix        = []byte{0xa6}
	ReversePairIndexKeyPrefix = []byte{0xa7}

	PoolKeyPrefix                  = []byte{0xab}
	PoolByReverveAccIndexKeyPrefix = []byte{0xac}
	PoolByPairIndexKeyPrefix       = []byte{0xad}

	DepositRequestKeyPrefix  = []byte{0xb0}
	WithdrawRequestKeyPrefix = []byte{0xb1}
	SwapRequestKeyPrefix     = []byte{0xb2}
)

// GetPairKey returns the store key to retrieve pair object from the pair id.
func GetPairKey(pairId uint64) []byte {
	return append(PairKeyPrefix, sdk.Uint64ToBigEndian(pairId)...)
}

// GetPairIndexKey returns the store key ...
func GetPairIndexKey(denomA string, denomB string) []byte {
	return append(append(PairIndexKeyPrefix, LengthPrefixString(denomA)...), LengthPrefixString(denomB)...)
}

// GetReversePairIndexKey returns the store key ...
func GetReversePairIndexKey(denomB string, denomA string) []byte {
	return append(append(PairIndexKeyPrefix, LengthPrefixString(denomB)...), LengthPrefixString(denomA)...)
}

// GetPoolKey returns the store key to retrieve pool object from the pool id.
func GetPoolKey(poolId uint64) []byte {
	return append(PairKeyPrefix, sdk.Uint64ToBigEndian(poolId)...)
}

// LengthPrefixString returns length-prefixed bytes representation
// of a string.
func LengthPrefixString(s string) []byte {
	bz := []byte(s)
	bzLen := len(bz)
	if bzLen == 0 {
		return bz
	}
	return append([]byte{byte(bzLen)}, bz...)
}
