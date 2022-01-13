package types

import (
	"fmt"

	sdk "github.com/cosmos/cosmos-sdk/types"
	paramstypes "github.com/cosmos/cosmos-sdk/x/params/types"

	farmingtypes "github.com/crescent-network/crescent/x/farming/types"
)

const (
	PoolReserveAccPrefix = "PoolReserveAcc"
	AccNameSplitter      = "|"
	AddressType          = farmingtypes.AddressType32Bytes
)

// TODO: sort keys
var (
	KeyInitialPoolCoinSupply   = []byte("InitialPoolCoinSupply")
	KeyBatchSize               = []byte("BatchSize")
	KeyTickPrecision           = []byte("TickPrecision")
	KeyMinInitialDepositAmount = []byte("MinInitialDepositAmount")
	KeyPoolCreationFee         = []byte("PoolCreationFee")
)

var (
	DefaultInitialPoolCoinSupply          = sdk.NewInt(1_000_000_000_000)
	DefaultBatchSize               uint32 = 1
	DefaultTickPrecision           uint32 = 3
	DefaultMinInitialDepositAmount        = sdk.NewInt(1000000)
	DefaultPoolCreationFee                = sdk.NewCoins(sdk.NewCoin(sdk.DefaultBondDenom, sdk.NewInt(1000000)))
)

var (
	MinOfferCoinAmount = sdk.NewInt(100) // This value can be modified in the future

	GlobalEscrowAddr = farmingtypes.DeriveAddress(AddressType, ModuleName, "GlobalEscrow")
)

var _ paramstypes.ParamSet = (*Params)(nil)

func ParamKeyTable() paramstypes.KeyTable {
	return paramstypes.NewKeyTable().RegisterParamSet(&Params{})
}

func DefaultParams() Params {
	return Params{
		InitialPoolCoinSupply:   DefaultInitialPoolCoinSupply,
		BatchSize:               DefaultBatchSize,
		TickPrecision:           DefaultTickPrecision,
		MinInitialDepositAmount: DefaultMinInitialDepositAmount,
		PoolCreationFee:         DefaultPoolCreationFee,
	}
}

func (params *Params) ParamSetPairs() paramstypes.ParamSetPairs {
	return paramstypes.ParamSetPairs{
		paramstypes.NewParamSetPair(KeyInitialPoolCoinSupply, &params.InitialPoolCoinSupply, validateInitialPoolCoinSupply),
		paramstypes.NewParamSetPair(KeyBatchSize, &params.BatchSize, validateBatchSize),
		paramstypes.NewParamSetPair(KeyTickPrecision, &params.TickPrecision, validateTickPrecision),
		paramstypes.NewParamSetPair(KeyMinInitialDepositAmount, &params.MinInitialDepositAmount, validateMinInitialDepositAmount),
		paramstypes.NewParamSetPair(KeyPoolCreationFee, &params.PoolCreationFee, validatePoolCreationFee),
	}
}

func (params Params) Validate() error {
	for _, field := range []struct {
		val          interface{}
		validateFunc func(i interface{}) error
	}{
		{params.InitialPoolCoinSupply, validateInitialPoolCoinSupply},
		{params.BatchSize, validateBatchSize},
		{params.TickPrecision, validateTickPrecision},
		{params.MinInitialDepositAmount, validateMinInitialDepositAmount},
		{params.PoolCreationFee, validatePoolCreationFee},
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

func validateTickPrecision(i interface{}) error {
	_, ok := i.(uint32)
	if !ok {
		return fmt.Errorf("invalid parameter type: %T", i)
	}

	return nil
}

func validateMinInitialDepositAmount(i interface{}) error {
	v, ok := i.(sdk.Int)
	if !ok {
		return fmt.Errorf("invalid parameter type: %T", i)
	}

	if v.IsNegative() {
		return fmt.Errorf("minimum initial deposit amount must not be negative: %s", v)
	}

	return nil
}

func validatePoolCreationFee(i interface{}) error {
	v, ok := i.(sdk.Coins)
	if !ok {
		return fmt.Errorf("invalid parameter type: %T", i)
	}

	if err := v.Validate(); err != nil {
		return fmt.Errorf("invalid pool creation fee: %w", err)
	}

	return nil
}
