package types

import (
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

// GetPairIndexKey returns the index key to retrieve denomB that is used to iterate pairs.
func GetPairIndexKey(denomA string, denomB string) []byte {
	return append(append(PairIndexKeyPrefix, LengthPrefixString(denomA)...), LengthPrefixString(denomB)...)
}

// GetReversePairIndexKey returns the index key to retrieve denomA that is used to iterate pairs.
func GetReversePairIndexKey(denomB string, denomA string) []byte {
	return append(append(ReversePairIndexKeyPrefix, LengthPrefixString(denomB)...), LengthPrefixString(denomA)...)
}

// GetPoolKey returns the store key to retrieve pool object from the pool id.
func GetPoolKey(poolId uint64) []byte {
	return append(PoolKeyPrefix, sdk.Uint64ToBigEndian(poolId)...)
}

// GetPoolByReserveAccIndexKey returns the index key to retrieve poolIds to that is used to iterate pools.
func GetPoolByReserveAccIndexKey(reserveAcc sdk.AccAddress, poolId uint64) []byte {
	return append(append(PoolByReverveAccIndexKeyPrefix, address.MustLengthPrefix(reserveAcc)...), sdk.Uint64ToBigEndian(poolId)...)
}

// GetPoolsByPairIndexKey returns the index key to retrieve pool id that is used to iterate pools.
func GetPoolsByPairIndexKey(pairId uint64, poolId uint64) []byte {
	return append(append(PoolByPairIndexKeyPrefix, sdk.Uint64ToBigEndian(pairId)...), sdk.Uint64ToBigEndian(poolId)...)
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
