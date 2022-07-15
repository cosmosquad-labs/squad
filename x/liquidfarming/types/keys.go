package types

import (
	"bytes"
	fmt "fmt"
	time "time"

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
	LastBidIdKeyPrefix      = []byte{0xe1}
	LastRewardsAuctionIdKey = []byte{0xe2} // key to retrieve the latest auction id

	QueuedFarmingKeyPrefix      = []byte{0xe4}
	QueuedFarmingIndexKeyPrefix = []byte{0xe5}

	AuctionKeyPrefix = []byte{0xe7}

	BidKeyPrefix = []byte{0xea}
)

// GetLastBidIdKey returns the store key to retrieve the latest bid id.
func GetLastBidIdKey(auctionId uint64) []byte {
	return append(LastBidIdKeyPrefix, sdk.Uint64ToBigEndian(auctionId)...)
}

// GetQueuedFarmingKey returns a key for a queued farming.
func GetQueuedFarmingKey(endTime time.Time, farmingCoinDenom string, farmerAcc sdk.AccAddress) []byte {
	return append(append(append(QueuedFarmingKeyPrefix,
		LengthPrefixTimeBytes(endTime)...),
		LengthPrefixString(farmingCoinDenom)...),
		farmerAcc...)
}

// GetQueuedFarmingIndexKey returns an indexing key for a queued farming.
func GetQueuedFarmingIndexKey(farmerAcc sdk.AccAddress, farmingCoinDenom string, endTime time.Time) []byte {
	return append(append(append(QueuedFarmingIndexKeyPrefix,
		address.MustLengthPrefix(farmerAcc)...),
		LengthPrefixString(farmingCoinDenom)...),
		sdk.FormatTimeBytes(endTime)...)
}

// GetQueuedFarmingsByFarmerAndDenomPrefix returns a key prefix used to
// iterate queued farmings by farmer address and farming coin denom.
func GetQueuedFarmingsByFarmerAndDenomPrefix(farmerAcc sdk.AccAddress, farmingCoinDenom string) []byte {
	return append(append(QueuedFarmingIndexKeyPrefix,
		address.MustLengthPrefix(farmerAcc)...),
		LengthPrefixString(farmingCoinDenom)...)
}

// GetQueuedFarmingsByFarmerPrefix returns a key prefix used to iterate
// queued farmings by a farmer.
func GetQueuedFarmingsByFarmerPrefix(farmerAcc sdk.AccAddress) []byte {
	return append(QueuedFarmingIndexKeyPrefix, address.MustLengthPrefix(farmerAcc)...)
}

// GetQueuedFarmingEndBytes returns end bytes for iteration of queued farmings.
// The returned end bytes should be used directly, not through
// sdk.InclusiveEndBytes.
// The range this end bytes form includes queued farmings with same endTime.
func GetQueuedFarmingEndBytes(endTime time.Time) []byte {
	return append(QueuedFarmingKeyPrefix, LengthPrefixTimeBytes(endTime.Add(1))...)
}

// GetRewardsAuctionKey returns the store key to retrieve rewards auction object.
func GetRewardsAuctionKey(poolId, auctionId uint64) []byte {
	return append(append(AuctionKeyPrefix, sdk.Uint64ToBigEndian(poolId)...), sdk.Uint64ToBigEndian(auctionId)...)
}

// GetBidKey returns the store key to retrieve the bid object.
func GetBidKey(auctionId uint64, bidId uint64) []byte {
	return append(append(BidKeyPrefix, sdk.Uint64ToBigEndian(auctionId)...), sdk.Uint64ToBigEndian(bidId)...)
}

// ParseQueuedFarmingKey parses a queued farming key.
func ParseQueuedFarmingKey(key []byte) (endTime time.Time, farmingCoinDenom string, farmerAcc sdk.AccAddress) {
	if !bytes.HasPrefix(key, QueuedFarmingKeyPrefix) {
		panic("key does not have proper prefix")
	}
	timeLen := key[1]
	var err error
	endTime, err = sdk.ParseTimeBytes(key[2 : 2+timeLen])
	if err != nil {
		panic(fmt.Errorf("parse end time: %w", err))
	}
	denomLen := key[2+timeLen]
	farmingCoinDenom = string(key[3+timeLen : 3+timeLen+denomLen])
	farmerAcc = key[3+timeLen+denomLen:]
	return
}

// ParseQueuedFarmingIndexKey parses a queued farming index key.
func ParseQueuedFarmingIndexKey(key []byte) (farmerAcc sdk.AccAddress, farmingCoinDenom string, endTime time.Time) {
	if !bytes.HasPrefix(key, QueuedFarmingIndexKeyPrefix) {
		panic("key does not have proper prefix")
	}
	addrLen := key[1]
	farmerAcc = key[2 : 2+addrLen]
	denomLen := key[2+addrLen]
	farmingCoinDenom = string(key[3+addrLen : 3+addrLen+denomLen])
	var err error
	endTime, err = sdk.ParseTimeBytes(key[3+addrLen+denomLen:])
	if err != nil {
		panic(fmt.Errorf("parse end time: %w", err))
	}
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

// LengthPrefixTimeBytes returns length-prefixed bytes representation
// of time.Time.
func LengthPrefixTimeBytes(t time.Time) []byte {
	bz := sdk.FormatTimeBytes(t)
	return append([]byte{byte(len(bz))}, bz...)
}
