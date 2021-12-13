package types

import (
	"fmt"

	sdk "github.com/cosmos/cosmos-sdk/types"
	paramstypes "github.com/cosmos/cosmos-sdk/x/params/types"
)

var (
	KeyInitialPoolCoinSupply = []byte("InitialPoolCoinSupply")
)

var (
	DefaultInitialPoolCoinSupply = sdk.NewInt(1_000_000_000_000)
)

var _ paramstypes.ParamSet = (*Params)(nil)

func ParamKeyTable() paramstypes.KeyTable {
	return paramstypes.NewKeyTable().RegisterParamSet(&Params{})
}

func DefaultParams() Params {
	return Params{
		InitialPoolCoinSupply: DefaultInitialPoolCoinSupply,
	}
}

func (params *Params) ParamSetPairs() paramstypes.ParamSetPairs {
	return paramstypes.ParamSetPairs{
		paramstypes.NewParamSetPair(KeyInitialPoolCoinSupply, &params.InitialPoolCoinSupply, validateInitialPoolCoinSupply),
	}
}

func (params Params) Validate() error {
	for _, field := range []struct {
		val          interface{}
		validateFunc func(i interface{}) error
	}{
		{params.InitialPoolCoinSupply, validateInitialPoolCoinSupply},
	} {
		if err := field.validateFunc(field.val); err != nil {
			return err
		}
	}
	return nil
}

func validateInitialPoolCoinSupply(i interface{}) error {
	v, ok := i.(sdk.Int)
	if !ok {
		return fmt.Errorf("invalid parameter type: %T", i)
	}

	if v.IsNil() {
		return fmt.Errorf("initial pool coin supply must not be nil")
	}

	if !v.IsPositive() {
		return fmt.Errorf("initial pool coin supply must be positive: %s", v)
	}

	return nil
}
