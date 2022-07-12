package types

import (
	"bytes"

	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/cosmos/cosmos-sdk/types/address"
)

const (
	// ModuleName defines the module name
	ModuleName = "liquidfarming"

	// StoreKey defines the primary module store key
	StoreKey = ModuleName

	// RouterKey is the message route for slashing
	RouterKey = ModuleName

	// QuerierRoute defines the module's query routing key
	QuerierRoute = ModuleName
)

// keys for the store prefixes
var (
	LastQueuedFarmingIdKeyPrefix = []byte{0xe1}
	LastBidIdKeyPrefix           = []byte{0xe2}
	LastRewardsAuctionIdKey      = []byte{0xe3} // key to retrieve the latest auction id

	QueuedFarmingKeyPrefix      = []byte{0xe4}
	QueuedFarmingIndexKeyPrefix = []byte{0xe5}

	AuctionKeyPrefix = []byte{0xe8}
)

// GetLastQueuedFarmingIdKey returns the store key to retrieve the latest deposit request id.
func GetLastQueuedFarmingIdKey(poolId uint64) []byte {
	return append(LastQueuedFarmingIdKeyPrefix, sdk.Uint64ToBigEndian(poolId)...)
}

// GetLastBidIdKey returns the store key to retrieve the latest bid id.
func GetLastBidIdKey(auctionId uint64) []byte {
	return append(LastBidIdKeyPrefix, sdk.Uint64ToBigEndian(auctionId)...)
}

// GetQueuedFarmingKey returns the store key to retrieve deposit request object.
func GetQueuedFarmingKey(poolId, reqId uint64) []byte {
	return append(append(QueuedFarmingKeyPrefix, sdk.Uint64ToBigEndian(poolId)...), sdk.Uint64ToBigEndian(reqId)...)
}

// GetQueuedFarmingIndexKey returns the index key to map deposit requests.
func GetQueuedFarmingIndexKey(depositor sdk.AccAddress, poolId, reqId uint64) []byte {
	return append(append(append(QueuedFarmingIndexKeyPrefix, address.MustLengthPrefix(depositor)...),
		sdk.Uint64ToBigEndian(poolId)...), sdk.Uint64ToBigEndian(reqId)...)
}

// GetQueuedFarmingIndexKeyPrefix returns the index key prefix to iterate
// deposit requests by a depositor.
func GetQueuedFarmingIndexKeyPrefix(depositor sdk.AccAddress) []byte {
	return append(QueuedFarmingIndexKeyPrefix, address.MustLengthPrefix(depositor)...)
}

// GetRewardsAuctionKey returns the store key to retrieve rewards auction object.
func GetRewardsAuctionKey(poolId, auctionId uint64) []byte {
	return append(append(AuctionKeyPrefix, sdk.Uint64ToBigEndian(poolId)...), sdk.Uint64ToBigEndian(auctionId)...)
}

// ParseQueuedFarmingIndexKey parses a deposit request index key.
func ParseQueuedFarmingIndexKey(key []byte) (depositor sdk.AccAddress, poolId, reqId uint64) {
	if !bytes.HasPrefix(key, QueuedFarmingIndexKeyPrefix) {
		panic("key does not have proper prefix")
	}

	addrLen := key[1]
	depositor = key[2 : 2+addrLen]
	poolId = sdk.BigEndianToUint64(key[2+addrLen : 2+addrLen+8])
	reqId = sdk.BigEndianToUint64(key[2+addrLen+8:])
	return
}
