package types

import (
	fmt "fmt"

	"github.com/cosmos/cosmos-sdk/codec"
	sdk "github.com/cosmos/cosmos-sdk/types"

	farmingtypes "github.com/cosmosquad-labs/squad/x/farming/types"
)

const (
	SellingReserveAddressPrefix string = "SellingReserveAddress"
	PayingReserveAddressPrefix  string = "PayingReserveAddress"
	ModuleAddressNameSplitter   string = "|"

	ReserveAddressType = farmingtypes.AddressType32Bytes
)

func MustMarshalRewardsAuction(cdc codec.BinaryCodec, auction RewardsAuction) []byte {
	return cdc.MustMarshal(&auction)
}

func MustUnmarshalRewardsAuction(cdc codec.BinaryCodec, value []byte) RewardsAuction {
	pair, err := UnmarshalRewardsAuction(cdc, value)
	if err != nil {
		panic(err)
	}
	return pair
}

func UnmarshalRewardsAuction(cdc codec.BinaryCodec, value []byte) (auction RewardsAuction, err error) {
	err = cdc.UnmarshalInterface(value, &auction)
	return auction, err
}

// SellingReserveAddress returns the selling reserve address with the given pool id.
func SellingReserveAddress(poolId uint64) sdk.AccAddress {
	return farmingtypes.DeriveAddress(ReserveAddressType, ModuleName, SellingReserveAddressPrefix+ModuleAddressNameSplitter+fmt.Sprint(poolId))
}

// PayingReserveAddress returns the paying reserve address with the given pool id.
func PayingReserveAddress(poolId uint64) sdk.AccAddress {
	return farmingtypes.DeriveAddress(ReserveAddressType, ModuleName, PayingReserveAddressPrefix+ModuleAddressNameSplitter+fmt.Sprint(poolId))
}
