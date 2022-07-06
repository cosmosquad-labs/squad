package types

import (
	"fmt"

	sdk "github.com/cosmos/cosmos-sdk/types"
	paramtypes "github.com/cosmos/cosmos-sdk/x/params/types"
)

// Parameter store keys
var (
	KeyLiquidFarms           = []byte("LiquidFarms")
	KeyLiquidFarmCreationFee = []byte("LiquidFarmCreationFee")

	DefaultLiquidFarms           = []LiquidFarm{}
	DefaultLiquidFarmCreationFee = sdk.NewCoins(sdk.NewCoin(sdk.DefaultBondDenom, sdk.NewInt(100_000_000)))
)

var _ paramtypes.ParamSet = (*Params)(nil)

// ParamKeyTable the param key table for launch module.
func ParamKeyTable() paramtypes.KeyTable {
	return paramtypes.NewKeyTable().RegisterParamSet(&Params{})
}

// DefaultParams returns a default set of parameters.
func DefaultParams() Params {
	return Params{
		LiquidFarms:           DefaultLiquidFarms,
		LiquidFarmCreationFee: DefaultLiquidFarmCreationFee,
	}
}

// ParamSetPairs get the params.ParamSet.
func (p *Params) ParamSetPairs() paramtypes.ParamSetPairs {
	return paramtypes.ParamSetPairs{
		paramtypes.NewParamSetPair(KeyLiquidFarms, &p.LiquidFarms, validateLiquidFarms),
		paramtypes.NewParamSetPair(KeyLiquidFarmCreationFee, &p.LiquidFarmCreationFee, validateLiquidFarmCreationFee),
	}
}

// Validate validates the set of parameters.
func (p Params) Validate() error {
	for _, v := range []struct {
		value     interface{}
		validator func(interface{}) error
	}{
		{p.LiquidFarms, validateLiquidFarms},
		{p.LiquidFarmCreationFee, validateLiquidFarmCreationFee},
	} {
		if err := v.validator(v.value); err != nil {
			return err
		}
	}
	return nil
}

func validateLiquidFarms(i interface{}) error {
	liquidFarms, ok := i.([]LiquidFarm)
	if !ok {
		return fmt.Errorf("invalid parameter type: %T", i)
	}

	// TODO: not implemented yet
	// Do we allow 0 for minimum params?
	for _, lf := range liquidFarms {
		fmt.Println(lf)
	}

	return nil
}

func validateLiquidFarmCreationFee(i interface{}) error {
	v, ok := i.(sdk.Coins)
	if !ok {
		return fmt.Errorf("invalid parameter type: %T", i)
	}

	if err := v.Validate(); err != nil {
		return err
	}

	return nil
}
