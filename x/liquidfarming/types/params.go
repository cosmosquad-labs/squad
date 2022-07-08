package types

import (
	"fmt"

	sdk "github.com/cosmos/cosmos-sdk/types"
	paramtypes "github.com/cosmos/cosmos-sdk/x/params/types"
)

// Parameter store keys
var (
	KeyLiquidFarmCreationFee = []byte("LiquidFarmCreationFee")
	KeyDelayedDepositGasFee  = []byte("DelayedDepositGasFee")
	KeyLiquidFarms           = []byte("LiquidFarms")

	DefaultLiquidFarmCreationFee = sdk.NewCoins(sdk.NewCoin(sdk.DefaultBondDenom, sdk.NewInt(100_000_000)))
	DefaultDelayedDepositGasFee  = sdk.Gas(60000)
	DefaultLiquidFarms           = []LiquidFarm{}
)

var _ paramtypes.ParamSet = (*Params)(nil)

// ParamKeyTable the param key table for launch module.
func ParamKeyTable() paramtypes.KeyTable {
	return paramtypes.NewKeyTable().RegisterParamSet(&Params{})
}

// DefaultParams returns a default set of parameters.
func DefaultParams() Params {
	return Params{
		LiquidFarmCreationFee: DefaultLiquidFarmCreationFee,
		DelayedDepositGasFee:  DefaultDelayedDepositGasFee,
		LiquidFarms:           DefaultLiquidFarms,
	}
}

// ParamSetPairs get the params.ParamSet.
func (p *Params) ParamSetPairs() paramtypes.ParamSetPairs {
	return paramtypes.ParamSetPairs{
		paramtypes.NewParamSetPair(KeyLiquidFarmCreationFee, &p.LiquidFarmCreationFee, validateLiquidFarmCreationFee),
		paramtypes.NewParamSetPair(KeyDelayedDepositGasFee, &p.DelayedDepositGasFee, validateDelayedDepositGasFee),
		paramtypes.NewParamSetPair(KeyLiquidFarms, &p.LiquidFarms, validateLiquidFarms),
	}
}

// Validate validates the set of parameters.
func (p Params) Validate() error {
	for _, v := range []struct {
		value     interface{}
		validator func(interface{}) error
	}{
		{p.LiquidFarmCreationFee, validateLiquidFarmCreationFee},
		{p.DelayedDepositGasFee, validateDelayedDepositGasFee},
		{p.LiquidFarms, validateLiquidFarms},
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
		fmt.Println("validate: ", lf)
	}

	return nil
}

func validateDelayedDepositGasFee(i interface{}) error {
	_, ok := i.(sdk.Gas)
	if !ok {
		return fmt.Errorf("invalid parameter type: %T", i)
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
