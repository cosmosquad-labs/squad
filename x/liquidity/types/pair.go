package types

import (
	"github.com/cosmos/cosmos-sdk/codec"
)

// NewPair returns a new pair object.
func NewPair(xCoinDenom, yCoinDenom string) Pair {
	return Pair{
		Id:                0, // id not set yet, a keeper will set it
		XCoinDenom:        xCoinDenom,
		YCoinDenom:        yCoinDenom,
		LastSwapRequestId: 0,
		LastPrice:         nil,
	}
}

// MustMarshalPair returns the pair bytes.
// It throws panic if it fails.
func MustMarshalPair(cdc codec.BinaryCodec, pair Pair) []byte {
	return cdc.MustMarshal(&pair)
}

// MustUnmarshalPair return the unmarshaled pair from bytes.
// It throws panic if it fails.
func MustUnmarshalPair(cdc codec.BinaryCodec, value []byte) Pair {
	pair, err := UnmarshalPair(cdc, value)
	if err != nil {
		panic(err)
	}

	return pair
}

// UnmarshalPair returns the pair from bytes.
func UnmarshalPair(cdc codec.BinaryCodec, value []byte) (pair Pair, err error) {
	err = cdc.Unmarshal(value, &pair)
	return pair, err
}
