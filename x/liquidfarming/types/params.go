package types

import (
	"fmt"

	sdk "github.com/cosmos/cosmos-sdk/types"
	paramtypes "github.com/cosmos/cosmos-sdk/x/params/types"
)

// Parameter store keys
var (
	KeyLiquidFarmCreationFee = []byte("LiquidFarmCreationFee")
	KeyDelayedFarmGasFee     = []byte("DelayedFarmGasFee")
	KeyLiquidFarms           = []byte("LiquidFarms")

	DefaultLiquidFarmCreationFee = sdk.NewCoins(sdk.NewCoin(sdk.DefaultBondDenom, sdk.NewInt(100_000_000)))
	DefaultDelayedFarmGasFee     = sdk.Gas(60000)
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
		DelayedFarmGasFee:     DefaultDelayedFarmGasFee,
		LiquidFarms:           DefaultLiquidFarms,
	}
}

// ParamSetPairs get the params.ParamSet.
func (p *Params) ParamSetPairs() paramtypes.ParamSetPairs {
	return paramtypes.ParamSetPairs{
		paramtypes.NewParamSetPair(KeyLiquidFarmCreationFee, &p.LiquidFarmCreationFee, validateLiquidFarmCreationFee),
		paramtypes.NewParamSetPair(KeyDelayedFarmGasFee, &p.DelayedFarmGasFee, validateDelayedFarmGasFee),
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
		{p.DelayedFarmGasFee, validateDelayedFarmGasFee},
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

	for _, liquidFarm := range liquidFarms {
		if liquidFarm.MinimumBidAmount.IsNegative() {
			return fmt.Errorf("minimum bid amount can't be negative value: %s", liquidFarm.MinimumBidAmount)
		}
		if liquidFarm.MinimumDepositAmount.IsNegative() {
			return fmt.Errorf("minimum deposit amount can't be negative value: %s", liquidFarm.MinimumBidAmount)
		}
	}

	return nil
}

func validateDelayedFarmGasFee(i interface{}) error {
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
