package types

import "github.com/cosmos/cosmos-sdk/codec"

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
