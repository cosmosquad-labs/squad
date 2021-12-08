package types

import (
	"fmt"

	"gopkg.in/yaml.v2"

	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
	paramstypes "github.com/cosmos/cosmos-sdk/x/params/types"
)

const (
	// MaxBearingNameLength is the maximum length of the name of each bearing.
	MaxBearingNameLength int = 50
	// DefaultEpochBlocks is the default epoch blocks.
	DefaultEpochBlocks uint32 = 1
)

// Parameter store keys
var (
	KeyBearings    = []byte("Bearings")
	KeyEpochBlocks = []byte("EpochBlocks")
)

var _ paramstypes.ParamSet = (*Params)(nil)

// ParamKeyTable returns the parameter key table.
func ParamKeyTable() paramstypes.KeyTable {
	return paramstypes.NewKeyTable().RegisterParamSet(&Params{})
}

// DefaultParams returns the default bearing module parameters.
func DefaultParams() Params {
	return Params{
		Bearings:    []Bearing{},
		EpochBlocks: DefaultEpochBlocks,
	}
}

// ParamSetPairs implements paramstypes.ParamSet.
func (p *Params) ParamSetPairs() paramstypes.ParamSetPairs {
	return paramstypes.ParamSetPairs{
		paramstypes.NewParamSetPair(KeyBearings, &p.Bearings, ValidateBearings),
		paramstypes.NewParamSetPair(KeyEpochBlocks, &p.EpochBlocks, ValidateEpochBlocks),
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
		{p.Bearings, ValidateBearings},
	} {
		if err := v.validator(v.value); err != nil {
			return err
		}
	}
	return nil
}

// ValidateBearings validates bearing name and total rate.
// The total rate of bearings with the same source address must not exceed 1.
func ValidateBearings(i interface{}) error {
	bearings, ok := i.([]Bearing)
	if !ok {
		return fmt.Errorf("invalid parameter type: %T", i)
	}
	names := make(map[string]bool)
	for _, bearing := range bearings {
		err := bearing.Validate()
		if err != nil {
			return err
		}
		if _, ok := names[bearing.Name]; ok {
			return sdkerrors.Wrap(ErrDuplicateBearingName, bearing.Name)
		}
		names[bearing.Name] = true
	}
	bearingsBySourceMap := GetBearingsBySourceMap(bearings)
	for addr, bearingsBySource := range bearingsBySourceMap {
		if bearingsBySource.TotalRate.GT(sdk.OneDec()) {
			// If the TotalRate of Bearings with the same source address exceeds 1,
			// recalculate and verify the TotalRate of Bearings with overlapping time ranges.
			for _, bearing := range bearingsBySource.Bearings {
				totalRate := sdk.ZeroDec()
				for _, bearingToCheck := range bearingsBySource.Bearings {
					if DateRangesOverlap(bearing.StartTime, bearing.EndTime, bearingToCheck.StartTime, bearingToCheck.EndTime) {
						totalRate = totalRate.Add(bearingToCheck.Rate)
					}
				}
				if totalRate.GT(sdk.OneDec()) {
					return sdkerrors.Wrapf(
						ErrInvalidTotalBearingRate,
						"total rate for source address %s must not exceed 1: %v", addr, totalRate)
				}
			}

		}
	}
	return nil
}

// ValidateEpochBlocks validates epoch blocks.
func ValidateEpochBlocks(i interface{}) error {
	_, ok := i.(uint32)
	if !ok {
		return fmt.Errorf("invalid parameter type: %T", i)
	}
	return nil
}
