package types

import (
	sdk "github.com/cosmos/cosmos-sdk/types"
)

const (
	// ModuleName defines the module name
	ModuleName = "auction"

	// StoreKey defines the primary module store key
	StoreKey = ModuleName

	// RouterKey is the message route for the module
	RouterKey = ModuleName

	// QuerierRoute defines the module's query routing key
	QuerierRoute = ModuleName

	// MemStoreKey defines the in-memory store key
	MemStoreKey = "mem_auction"
)

var (
	LastAuctionIdKey   = []byte{0x11} // key to retrieve the latest auction id
	LastBidIdKeyPrefix = []byte{0x12}

	AuctionKeyPrefix = []byte{0x21}
)

// GetLastBidIdKey returns the store key to retrieve the latest bid id.
func GetLastBidIdKey(auctionId uint64) []byte {
	return append(LastBidIdKeyPrefix, sdk.Uint64ToBigEndian(auctionId)...)
}

// GetAuctionKey returns the store key to retrieve the auction object.
func GetAuctionKey(auctionId uint64) []byte {
	return append(AuctionKeyPrefix, sdk.Uint64ToBigEndian(auctionId)...)
}
