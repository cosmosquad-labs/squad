package types

import (
	"fmt"

	sdk "github.com/cosmos/cosmos-sdk/types"
	paramstypes "github.com/cosmos/cosmos-sdk/x/params/types"
)

var (
	KeyInitialPoolCoinSupply = []byte("InitialPoolCoinSupply")
	KeyBatchSize             = []byte("BatchSize")
)

var (
	DefaultInitialPoolCoinSupply        = sdk.NewInt(1_000_000_000_000)
	DefaultBatchSize             uint32 = 1
)

var _ paramstypes.ParamSet = (*Params)(nil)

func ParamKeyTable() paramstypes.KeyTable {
	return paramstypes.NewKeyTable().RegisterParamSet(&Params{})
}

func DefaultParams() Params {
	return Params{
		InitialPoolCoinSupply: DefaultInitialPoolCoinSupply,
		BatchSize:             DefaultBatchSize,
	}
}

func (params *Params) ParamSetPairs() paramstypes.ParamSetPairs {
	return paramstypes.ParamSetPairs{
		paramstypes.NewParamSetPair(KeyInitialPoolCoinSupply, &params.InitialPoolCoinSupply, validateInitialPoolCoinSupply),
		paramstypes.NewParamSetPair(KeyBatchSize, &params.BatchSize, validateBatchSize),
	}
}

func (params Params) Validate() error {
	for _, field := range []struct {
		val          interface{}
		validateFunc func(i interface{}) error
	}{
		{params.InitialPoolCoinSupply, validateInitialPoolCoinSupply},
		{params.BatchSize, validateBatchSize},
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

func validateBatchSize(i interface{}) error {
	v, ok := i.(uint32)
	if !ok {
		return fmt.Errorf("invalid parameter type: %T", i)
	}

	if v == 0 {
		return fmt.Errorf("batch size must be positive: %d", v)
	}

	return nil
}
