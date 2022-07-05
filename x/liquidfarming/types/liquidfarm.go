package types

import (
	"strings"

	"github.com/cosmos/cosmos-sdk/codec"
	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
	"gopkg.in/yaml.v2"
)

func NewLiquidFarm(id uint64, poolId uint64, poolCoinDenom string, lfCoinDenom string, reserveAddr string) LiquidFarm {
	return LiquidFarm{
		Id:             id,
		PoolId:         poolId,
		PoolCoinDenom:  poolCoinDenom,
		LfCoinDenom:    lfCoinDenom,
		ReserveAddress: reserveAddr,
	}
}

func (l LiquidFarm) String() string {
	out, _ := yaml.Marshal(l)
	return string(out)
}

// TODO: double check with these validity checks
func (l LiquidFarm) Validate() error {
	if l.Id == 0 {
		return sdkerrors.Wrap(sdkerrors.ErrInvalidRequest, "invalid id")
	}
	if l.PoolId == 0 {
		return sdkerrors.Wrap(sdkerrors.ErrInvalidRequest, "invalid pool id")
	}
	if err := sdk.ValidateDenom(l.PoolCoinDenom); err != nil {
		return sdkerrors.Wrapf(sdkerrors.ErrInvalidRequest, "invalid pool coin denom: %v", err)
	}
	if err := sdk.ValidateDenom(l.LfCoinDenom); err != nil {
		return sdkerrors.Wrapf(sdkerrors.ErrInvalidRequest, "invalid liquid farming coin denom: %v", err)
	}
	if !strings.HasPrefix(l.PoolCoinDenom, "pool") {
		return sdkerrors.Wrap(sdkerrors.ErrInvalidRequest, "invalid pool coin denom")
	}
	if !strings.HasPrefix(l.LfCoinDenom, "lf") {
		return sdkerrors.Wrap(sdkerrors.ErrInvalidRequest, "invalid liquid farming coin denom")
	}
	return nil
}

func UnmarshalLiquidFarm(cdc codec.BinaryCodec, value []byte) (liquidfarm LiquidFarm, err error) {
	err = cdc.Unmarshal(value, &liquidfarm)
	return liquidfarm, err
}

func MustMarshalLiquidFarm(cdc codec.BinaryCodec, liquidfarm LiquidFarm) []byte {
	return cdc.MustMarshal(&liquidfarm)
}

func MustUnmarshalLiquidFarm(cdc codec.BinaryCodec, value []byte) LiquidFarm {
	liquidfarm, err := UnmarshalLiquidFarm(cdc, value)
	if err != nil {
		panic(err)
	}

	return liquidfarm
}
