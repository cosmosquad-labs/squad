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

	// MemStoreKey defines the in-memory store key
	MemStoreKey = "mem_liquidfarming"
)

var (
	LastLiquidfarmIdKey     = []byte{0xe0} // key to retrieve the latest liquidfarm id
	LastRewardsAuctionIdKey = []byte{0xe1} // key to retrieve the latest auction id

	LiquidFarmKeyPrefix = []byte{0xe4}

	DepositRequestKeyPrefix      = []byte{0xe6}
	DepositRequestIndexKeyPrefix = []byte{0xe7}

	AuctionKeyPrefix = []byte{0xea}
)

func GetLiquidFarmKey(liquidfarmId uint64) []byte {
	return append(LiquidFarmKeyPrefix, sdk.Uint64ToBigEndian(liquidfarmId)...)
}

func GetDepositRequestKey(liquidfarmId, reqId uint64) []byte {
	return append(append(DepositRequestKeyPrefix, sdk.Uint64ToBigEndian(liquidfarmId)...), sdk.Uint64ToBigEndian(reqId)...)
}

func GetDepositRequestIndexKey(depositor sdk.AccAddress, liquidfarmId, reqId uint64) []byte {
	return append(append(append(DepositRequestIndexKeyPrefix, address.MustLengthPrefix(depositor)...),
		sdk.Uint64ToBigEndian(liquidfarmId)...), sdk.Uint64ToBigEndian(reqId)...)
}

func GetRewardsAuctionKey(liquidfarmId, auctionId uint64) []byte {
	return append(append(AuctionKeyPrefix, sdk.Uint64ToBigEndian(liquidfarmId)...), sdk.Uint64ToBigEndian(auctionId)...)
}
