package types

import (
	"fmt"

	sdk "github.com/cosmos/cosmos-sdk/types"
	paramstypes "github.com/cosmos/cosmos-sdk/x/params/types"
)

var _ paramstypes.ParamSet = (*Params)(nil)

var (
	KeyLendingAssetParams = []byte("LendingAssetParams")
)

func ParamKeyTable() paramstypes.KeyTable {
	return paramstypes.NewKeyTable().RegisterParamSet(&Params{})
}

// DefaultParams returns a default params for the lending module.
func DefaultParams() Params {
	return Params{
		LendingAssetParams: []LendingAssetParam{},
	}
}

// ParamSetPairs implements ParamSet.
func (params *Params) ParamSetPairs() paramstypes.ParamSetPairs {
	return paramstypes.ParamSetPairs{
		paramstypes.NewParamSetPair(KeyLendingAssetParams, &params.LendingAssetParams, validateLendingAssetParams),
	}
}

// Validate validates Params.
func (params Params) Validate() error {
	for _, field := range []struct {
		val          interface{}
		validateFunc func(i interface{}) error
	}{
		{params.LendingAssetParams, validateLendingAssetParams},
	} {
		if err := field.validateFunc(field.val); err != nil {
			return err
		}
	}
	return nil
}

func validateLendingAssetParams(i interface{}) error {
	v, ok := i.([]LendingAssetParam)
	if !ok {
		return fmt.Errorf("invalid parameter type: %T", i)
	}

	for _, p := range v {
		if err := sdk.ValidateDenom(p.Denom); err != nil {
			return fmt.Errorf("invalid lending asset denom: %w", err)
		}
		if !p.MinInterestRate.IsPositive() {
			return fmt.Errorf("min interest rate must be positive: %s", p.MinInterestRate)
		}
		if !p.MaxInterestRate.IsPositive() {
			return fmt.Errorf("max interest rate must be positive: %s", p.MaxInterestRate)
		}
		if p.UtilizationRatioPower == 0 {
			return fmt.Errorf("utilization ratio power must be positive: %d", p.UtilizationRatioPower)
		}
	}

	return nil
}

// NewLendingAssetParam returns a new LendingAssetParam from given parameters.
func NewLendingAssetParam(denom string, minInterestRate, maxInterestRate sdk.Dec, utilRatioPower uint32) LendingAssetParam {
	return LendingAssetParam{
		Denom:                 denom,
		MinInterestRate:       minInterestRate,
		MaxInterestRate:       maxInterestRate,
		UtilizationRatioPower: utilRatioPower,
	}
}
