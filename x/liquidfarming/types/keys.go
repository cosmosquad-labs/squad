package types

import (
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
	LastDepositRequestIdKeyPrefix = []byte{0xe1}
	LastBidIdKeyPrefix            = []byte{0xe2}
	LastRewardsAuctionIdKey       = []byte{0xe3} // key to retrieve the latest auction id

	DepositRequestKeyPrefix      = []byte{0xe4}
	DepositRequestIndexKeyPrefix = []byte{0xe5}

	AuctionKeyPrefix = []byte{0xe8}
)

// GetLastDepositRequestIdKey returns the store key to retrieve the latest deposit request id.
func GetLastDepositRequestIdKey(poolId uint64) []byte {
	return append(LastDepositRequestIdKeyPrefix, sdk.Uint64ToBigEndian(poolId)...)
}

// GetLastBidIdKey returns the store key to retrieve the latest bid id.
func GetLastBidIdKey(auctionId uint64) []byte {
	return append(LastBidIdKeyPrefix, sdk.Uint64ToBigEndian(auctionId)...)
}

// GetDepositRequestKey returns the store key to retrieve deposit request object.
func GetDepositRequestKey(poolId, reqId uint64) []byte {
	return append(append(DepositRequestKeyPrefix, sdk.Uint64ToBigEndian(poolId)...), sdk.Uint64ToBigEndian(reqId)...)
}

// GetDepositRequestIndexKey returns the index key to map deposit requests.
func GetDepositRequestIndexKey(depositor sdk.AccAddress, poolId, reqId uint64) []byte {
	return append(append(append(DepositRequestIndexKeyPrefix, address.MustLengthPrefix(depositor)...),
		sdk.Uint64ToBigEndian(poolId)...), sdk.Uint64ToBigEndian(reqId)...)
}

// GetRewardsAuctionKey returns the store key to retrieve rewards auction object.
func GetRewardsAuctionKey(poolId, auctionId uint64) []byte {
	return append(append(AuctionKeyPrefix, sdk.Uint64ToBigEndian(poolId)...), sdk.Uint64ToBigEndian(auctionId)...)
}
