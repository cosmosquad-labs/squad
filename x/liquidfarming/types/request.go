package types

import (
	"github.com/cosmos/cosmos-sdk/codec"
	sdk "github.com/cosmos/cosmos-sdk/types"
)

// NewDepositRequest returns a new DepositRequest.
func NewDepositRequest(poolId uint64, depositReqId uint64, depositor string, depositCoin sdk.Coin) DepositRequest {
	return DepositRequest{
		PoolId:      poolId,
		Id:          depositReqId,
		Depositor:   depositor,
		DepositCoin: depositCoin,
	}
}

// GetDepositor returns depositor int the form of sdk.AccAddress.
func (req DepositRequest) GetDepositor() sdk.AccAddress {
	addr, err := sdk.AccAddressFromBech32(req.Depositor)
	if err != nil {
		panic(err)
	}
	return addr
}

// Validate validates DepositRequest for genesis.
func (req DepositRequest) Validate() error {
	// TODO: not implemented yet
	return nil
}

// MarshalDepositRequest returns the DepositRequest bytes. Panics if fails.
func MarshalDepositRequest(cdc codec.BinaryCodec, msg DepositRequest) ([]byte, error) {
	return cdc.Marshal(&msg)
}

// UnmarshalDepositRequest returns the DepositRequest from bytes.
func UnmarshalDepositRequest(cdc codec.BinaryCodec, value []byte) (msg DepositRequest, err error) {
	err = cdc.Unmarshal(value, &msg)
	return msg, err
}

// MustMarshalDepositRequest returns the DepositRequest bytes. Panics if fails.
func MustMarshalDepositRequest(cdc codec.BinaryCodec, msg DepositRequest) []byte {
	bz, err := MarshalDepositRequest(cdc, msg)
	if err != nil {
		panic(err)
	}
	return bz
}

// MustUnmarshalDepositRequest returns the DepositRequest from bytes.
// It throws panic if it fails.
func MustUnmarshalDepositRequest(cdc codec.BinaryCodec, value []byte) DepositRequest {
	msg, err := UnmarshalDepositRequest(cdc, value)
	if err != nil {
		panic(err)
	}
	return msg
}
