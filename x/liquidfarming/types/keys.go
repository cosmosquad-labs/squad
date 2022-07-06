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
	LastRewardsAuctionIdKey = []byte{0xe1} // key to retrieve the latest auction id

	DepositRequestKeyPrefix      = []byte{0xe4}
	DepositRequestIndexKeyPrefix = []byte{0xe5}

	AuctionKeyPrefix = []byte{0xe8}
)

// GetDepositRequestKey returns the store key to retrieve deposit request object.
func GetDepositRequestKey(reqId, poolId uint64) []byte {
	return append(append(DepositRequestKeyPrefix, sdk.Uint64ToBigEndian(reqId)...), sdk.Uint64ToBigEndian(poolId)...)
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
