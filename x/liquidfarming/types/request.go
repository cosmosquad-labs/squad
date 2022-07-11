package types

import (
	"github.com/cosmos/cosmos-sdk/codec"
	sdk "github.com/cosmos/cosmos-sdk/types"
)

// NewQueuedFarming returns a new QueuedFarming.
func NewQueuedFarming(poolId uint64, depositReqId uint64, depositor string, depositCoin sdk.Coin) QueuedFarming {
	return QueuedFarming{
		PoolId:      poolId,
		Id:          depositReqId,
		Depositor:   depositor,
		DepositCoin: depositCoin,
	}
}

// GetDepositor returns depositor int the form of sdk.AccAddress.
func (req QueuedFarming) GetDepositor() sdk.AccAddress {
	addr, err := sdk.AccAddressFromBech32(req.Depositor)
	if err != nil {
		panic(err)
	}
	return addr
}

// Validate validates QueuedFarming for genesis.
func (req QueuedFarming) Validate() error {
	// TODO: not implemented yet
	return nil
}

// MarshalQueuedFarming returns the QueuedFarming bytes. Panics if fails.
func MarshalQueuedFarming(cdc codec.BinaryCodec, msg QueuedFarming) ([]byte, error) {
	return cdc.Marshal(&msg)
}

// UnmarshalQueuedFarming returns the QueuedFarming from bytes.
func UnmarshalQueuedFarming(cdc codec.BinaryCodec, value []byte) (msg QueuedFarming, err error) {
	err = cdc.Unmarshal(value, &msg)
	return msg, err
}

// MustMarshalQueuedFarming returns the QueuedFarming bytes. Panics if fails.
func MustMarshalQueuedFarming(cdc codec.BinaryCodec, msg QueuedFarming) []byte {
	bz, err := MarshalQueuedFarming(cdc, msg)
	if err != nil {
		panic(err)
	}
	return bz
}

// MustUnmarshalQueuedFarming returns the QueuedFarming from bytes.
// It throws panic if it fails.
func MustUnmarshalQueuedFarming(cdc codec.BinaryCodec, value []byte) QueuedFarming {
	msg, err := UnmarshalQueuedFarming(cdc, value)
	if err != nil {
		panic(err)
	}
	return msg
}
