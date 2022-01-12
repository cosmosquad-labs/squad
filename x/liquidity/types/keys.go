package types

import (
	"bytes"

	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/cosmos/cosmos-sdk/types/address"
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

	PoolKeyPrefix             = []byte{0xab}
	PoolByReverveAccKeyPrefix = []byte{0xac}
	PoolByPairIndexKeyPrefix  = []byte{0xad}

	DepositRequestKeyPrefix  = []byte{0xb0}
	WithdrawRequestKeyPrefix = []byte{0xb1}
	SwapRequestKeyPrefix     = []byte{0xb2}
)

// GetPairKey returns the store key to retrieve pair object from the pair id.
func GetPairKey(pairId uint64) []byte {
	return append(PairKeyPrefix, sdk.Uint64ToBigEndian(pairId)...)
}

// GetPairIndexKey returns the index key to retrieve pair id that is used to iterate pairs.
func GetPairIndexKey(denomA string, denomB string, pairId uint64) []byte {
	return append(append(append(PairIndexKeyPrefix, LengthPrefixString(denomA)...), LengthPrefixString(denomB)...), sdk.Uint64ToBigEndian(pairId)...)
}

// GetReversePairIndexKey returns the index key to retrieve pair id that is used to iterate pairs.
func GetReversePairIndexKey(denomB string, denomA string, pairId uint64) []byte {
	return append(append(append(ReversePairIndexKeyPrefix, LengthPrefixString(denomB)...), LengthPrefixString(denomA)...), sdk.Uint64ToBigEndian(pairId)...)
}

// GetPairByDenomKey returns the single denom index key.
func GetPairByDenomKey(denom string) []byte {
	return append(PairIndexKeyPrefix, LengthPrefixString(denom)...)
}

// GetReversePairByDenomKey returns the single denom index key.
func GetReversePairByDenomKey(denom string) []byte {
	return append(ReversePairIndexKeyPrefix, LengthPrefixString(denom)...)
}

// GetPoolKey returns the store key to retrieve pool object from the pool id.
func GetPoolKey(poolId uint64) []byte {
	return append(PoolKeyPrefix, sdk.Uint64ToBigEndian(poolId)...)
}

// GetPoolByReserveAccKey returns the index key to retrieve the particular pool.
func GetPoolByReserveAccKey(reserveAcc sdk.AccAddress) []byte {
	return append(PoolByReverveAccKeyPrefix, address.MustLengthPrefix(reserveAcc)...)
}

// GetPoolsByPairIndexKey returns the index key to retrieve pool id that is used to iterate pools.
func GetPoolsByPairIndexKey(pairId uint64, poolId uint64) []byte {
	return append(append(PoolByPairIndexKeyPrefix, sdk.Uint64ToBigEndian(pairId)...), sdk.Uint64ToBigEndian(poolId)...)
}

// GetPoolsByPairKey returns the store key to retrieve pool id to iterate pools.
func GetPoolsByPairKey(pairId uint64) []byte {
	return append(PoolByPairIndexKeyPrefix, sdk.Uint64ToBigEndian(pairId)...)
}

// GetDepositRequestKey returns the store key to retrieve deposit request object from the pool id and request id.
func GetDepositRequestKey(poolId uint64, id uint64) []byte {
	return append(append(DepositRequestKeyPrefix, sdk.Uint64ToBigEndian(poolId)...), sdk.Uint64ToBigEndian(id)...)
}

// GetWithdrawRequestKey returns the store key to retrieve withdaw request object from the pool id and request id
func GetWithdrawRequestKey(poolId uint64, id uint64) []byte {
	return append(append(WithdrawRequestKeyPrefix, sdk.Uint64ToBigEndian(poolId)...), sdk.Uint64ToBigEndian(id)...)
}

// GetSwapRequestKey returns the store key to retrieve deposit swap object from the pool id and request id
func GetSwapRequestKey(poolId uint64, id uint64) []byte {
	return append(append(SwapRequestKeyPrefix, sdk.Uint64ToBigEndian(poolId)...), sdk.Uint64ToBigEndian(id)...)
}

// ParsePairByDenomIndexKey parses a pair by denom index key.
func ParsePairByDenomIndexKey(key []byte) (denomB string, pairId uint64) {
	if !bytes.HasPrefix(key, PairIndexKeyPrefix) {
		panic("key does not have proper prefix")
	}

	denomALen := key[1]
	denomBLen := key[2+denomALen]
	denomB = string(key[3+denomALen : 3+denomALen+denomBLen])
	pairId = sdk.BigEndianToUint64(key[3+denomALen+denomBLen:])

	return
}

// ParseReversePairByDenomIndexKey parses a pair by denom index key.
func ParseReversePairByDenomIndexKey(key []byte) (denomA string, pairId uint64) {
	if !bytes.HasPrefix(key, ReversePairIndexKeyPrefix) {
		panic("key does not have proper prefix")
	}

	denomBLen := key[1]
	denomALen := key[2+denomBLen]
	denomA = string(key[3+denomBLen : 3+denomBLen+denomALen])
	pairId = sdk.BigEndianToUint64(key[3+denomBLen+denomALen:])
	return
}

// ParsePoolsByPairIndexKey parses a pool id from the index key.
func ParsePoolsByPairIndexKey(key []byte) (poolId uint64) {
	if !bytes.HasPrefix(key, PoolByPairIndexKeyPrefix) {
		panic("key does not have proper prefix")
	}

	bytesLen := 8
	poolId = sdk.BigEndianToUint64(key[1+bytesLen:])
	return
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
