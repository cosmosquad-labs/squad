package types

import (
	"fmt"

	"gopkg.in/yaml.v2"

	sdk "github.com/cosmos/cosmos-sdk/types"
	paramstypes "github.com/cosmos/cosmos-sdk/x/params/types"
)

// temporary bearing bond denom
const BearingDenom = "bgdex"

// Parameter store keys
var (
	KeyBearingValidators    = []byte("BearingValidators")
	KeyUnstakeFeeRate = []byte("UnstakeFeeRate")

	// DefaultUnstakeFeeRate is the default Unstake Fee Rate.
	DefaultUnstakeFeeRate = sdk.NewDecWithPrec(1, 3) // "0.001000000000000000"


	// Const variables

	MinimumStakingAmount = sdk.NewInt(1000000)
)

var _ paramstypes.ParamSet = (*Params)(nil)

// ParamKeyTable returns the parameter key table.
func ParamKeyTable() paramstypes.KeyTable {
	return paramstypes.NewKeyTable().RegisterParamSet(&Params{})
}

// DefaultParams returns the default bearing module parameters.
func DefaultParams() Params {
	return Params{
		BearingValidators:    []BearingValidator{},
		UnstakeFeeRate:       DefaultUnstakeFeeRate,
	}
}

// ParamSetPairs implements paramstypes.ParamSet.
func (p *Params) ParamSetPairs() paramstypes.ParamSetPairs {
	return paramstypes.ParamSetPairs{
		paramstypes.NewParamSetPair(KeyBearingValidators, &p.BearingValidators, ValidateBearingValidators),
		paramstypes.NewParamSetPair(KeyUnstakeFeeRate, &p.UnstakeFeeRate, validateUnstakeFeeRate),
	}
}

// String returns a human-readable string representation of the parameters.
func (p Params) String() string {
	out, _ := yaml.Marshal(p)
	return string(out)
}

// Validate validates parameters.
func (p Params) Validate() error {
	for _, v := range []struct {
		value     interface{}
		validator func(interface{}) error
	}{
		{p.BearingValidators, ValidateBearingValidators},
		{p.UnstakeFeeRate, validateUnstakeFeeRate},
	} {
		if err := v.validator(v.value); err != nil {
			return err
		}
	}
	return nil
}

// ValidateBearingValidators validates bearing validator and total weight.
func ValidateBearingValidators(i interface{}) error {
	bvs, ok := i.([]BearingValidator)
	if !ok {
		return fmt.Errorf("invalid parameter type: %T", i)
	}
	for _, bv := range bvs {
		_, valErr := sdk.ValAddressFromBech32(bv.ValidatorAddress)
		if valErr != nil {
			return valErr
		}

		if bv.Weight.IsNil() {
			return fmt.Errorf("bearing validator weight must not be nil")
		}

		if bv.Weight.IsNegative() {
			return fmt.Errorf("bearing validator weight must not be negative: %s", bv.Weight)
		}
	}
	// TODO: TBD total weight should be 1 or not
	return nil
}

func validateUnstakeFeeRate(i interface{}) error {
	v, ok := i.(sdk.Dec)
	if !ok {
		return fmt.Errorf("invalid parameter type: %T", i)
	}

	if v.IsNil() {
		return fmt.Errorf("unstake fee rate must not be nil")
	}

	if v.IsNegative() {
		return fmt.Errorf("unstake fee rate must not be negative: %s", v)
	}

	if v.GT(sdk.OneDec()) {
		return fmt.Errorf("unstake fee rate too large: %s", v)
	}

	return nil
}
