package types

import (
	fmt "fmt"
	"strconv"
	"strings"

	"github.com/cosmos/cosmos-sdk/codec"
	codectypes "github.com/cosmos/cosmos-sdk/codec/types"
	sdk "github.com/cosmos/cosmos-sdk/types"

	farmingtypes "github.com/cosmosquad-labs/squad/v2/x/farming/types"
)

const (
	LiquidFarmReserveAccPrefix string = "LiquidFarmReserveAcc"
)

// NewLiquidFarm returns a new LiquidFarm.
func NewLiquidFarm(poolId uint64, minFarmAmt, minBidAmount sdk.Int) LiquidFarm {
	return LiquidFarm{
		PoolId:            poolId,
		MinimumFarmAmount: minFarmAmt,
		MinimumBidAmount:  minBidAmount,
	}
}

// String returns a human-readable string representation of the LiquidFarm.
func (l LiquidFarm) String() string {
	out, _ := l.MarshalYAML()
	return out.(string)
}

// MarshalYAML returns the YAML representation of a LiquidFarm.
func (l LiquidFarm) MarshalYAML() (interface{}, error) {
	bz, err := codec.MarshalYAML(codec.NewProtoCodec(codectypes.NewInterfaceRegistry()), &l)
	if err != nil {
		return nil, err
	}
	return string(bz), err
}

// LiquidFarmCoinDenom returns a unique liquid farming coin denom for a LiquidFarm.
func LiquidFarmCoinDenom(poolId uint64) string {
	return fmt.Sprintf("lf%d", poolId)
}

// LiquidFarmReserveAddress returns the reserve address for a liquid farm with the given pool id.
func LiquidFarmReserveAddress(poolId uint64) sdk.AccAddress {
	return farmingtypes.DeriveAddress(
		ReserveAddressType,
		ModuleName,
		strings.Join([]string{LiquidFarmReserveAccPrefix, strconv.FormatUint(poolId, 10)}, ModuleAddressNameSplitter),
	)
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
