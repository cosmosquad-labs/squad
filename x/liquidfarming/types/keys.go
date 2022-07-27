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
	LastRewardsAuctionIdKey = []byte{0xe1} // key to retrieve the latest rewards auction id

	QueuedFarmingKeyPrefix      = []byte{0xe4}
	QueuedFarmingIndexKeyPrefix = []byte{0xe5}

	RewardsAuctionKeyPrefix = []byte{0xe7}

	BidKeyPrefix        = []byte{0xea}
	WinningBidKeyPrefix = []byte{0xeb}
)

// GetLastRewardsAuctionIdKey returns the store key to retrieve the last rewards auction
// by the given pool id.
func GetLastRewardsAuctionIdKey(poolId uint64) []byte {
	return append(LastRewardsAuctionIdKey, sdk.Uint64ToBigEndian(poolId)...)
}

// GetQueuedFarmingKey returns the store key to retrieve queued farming object
// by the given end time, farming coin denom, and farmer address.
func GetQueuedFarmingKey(endTime time.Time, farmingCoinDenom string, farmerAcc sdk.AccAddress) []byte {
	return append(append(append(QueuedFarmingKeyPrefix,
		LengthPrefixTimeBytes(endTime)...),
		LengthPrefixString(farmingCoinDenom)...),
		farmerAcc...)
}

// GetQueuedFarmingIndexKey returns the index key to retrieve queued farming object
// by the given farmer address, farming coin denom, and end time.
func GetQueuedFarmingIndexKey(farmerAcc sdk.AccAddress, farmingCoinDenom string, endTime time.Time) []byte {
	return append(append(append(QueuedFarmingIndexKeyPrefix,
		address.MustLengthPrefix(farmerAcc)...),
		LengthPrefixString(farmingCoinDenom)...),
		sdk.FormatTimeBytes(endTime)...)
}

// GetQueuedFarmingsByFarmerPrefix returns the index key prefix to iterate queued farming objects
// by the given farmer address.
func GetQueuedFarmingsByFarmerPrefix(farmerAcc sdk.AccAddress) []byte {
	return append(QueuedFarmingIndexKeyPrefix, address.MustLengthPrefix(farmerAcc)...)
}

// GetQueuedFarmingsByFarmerAndDenomPrefix returns the index key prefix to iterate queued farming objects
// by the given farmer address and farming coin denom.
func GetQueuedFarmingsByFarmerAndDenomPrefix(farmerAcc sdk.AccAddress, farmingCoinDenom string) []byte {
	return append(append(QueuedFarmingIndexKeyPrefix,
		address.MustLengthPrefix(farmerAcc)...),
		LengthPrefixString(farmingCoinDenom)...)
}

// GetQueuedFarmingEndBytes returns end time bytes to iterate queued farming objects
// by the given end time.
// By adding 1 to the given end time, the returned end bytes are inclusive of the endTime.
func GetQueuedFarmingEndBytes(endTime time.Time) []byte {
	return append(QueuedFarmingKeyPrefix, LengthPrefixTimeBytes(endTime.Add(1))...)
}

// GetRewardsAuctionKey returns the store key to retrieve rewards auction object
// by the given pool id and auction id.
func GetRewardsAuctionKey(poolId, auctionId uint64) []byte {
	return append(append(RewardsAuctionKeyPrefix, sdk.Uint64ToBigEndian(poolId)...), sdk.Uint64ToBigEndian(auctionId)...)
}

// GetBidKey returns the store key to retrieve the bid
// by the given pool id and bidder address.
func GetBidKey(poolId uint64, bidder sdk.AccAddress) []byte {
	return append(append(BidKeyPrefix, sdk.Uint64ToBigEndian(poolId)...), address.MustLengthPrefix(bidder)...)
}

// GetBidByPoolIdPrefix returns the prefix to iterate all bids
// by the given pool id.
func GetBidByPoolIdPrefix(poolId uint64) []byte {
	return append(BidKeyPrefix, sdk.Uint64ToBigEndian(poolId)...)
}

// GetBidByBidderPrefix returns the prefix to iterate all bids
// by the given bidder address.
func GetBidByBidderPrefix(bidder sdk.AccAddress) []byte {
	return append(BidKeyPrefix, address.MustLengthPrefix(bidder)...)
}

// GetWinningBidKey returns the store key to retrieve the winning bid
// by the given pool id and auction id.
func GetWinningBidKey(poolId uint64, auctionId uint64) []byte {
	return append(append(WinningBidKeyPrefix, sdk.Uint64ToBigEndian(poolId)...), sdk.Uint64ToBigEndian(auctionId)...)
}

// ParseQueuedFarmingKey parses a queued farming key bytes.
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

// ParseQueuedFarmingIndexKey parses a queued farming index key bytes.
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
