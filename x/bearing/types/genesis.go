package types

import (
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
)

// NewGenesisState returns new GenesisState instance.
func NewGenesisState(params Params, records []BearingRecord) *GenesisState {
	return &GenesisState{
		Params:         params,
		BearingRecords: records,
	}
}

// DefaultGenesisState returns the default genesis state.
func DefaultGenesisState() *GenesisState {
	return NewGenesisState(
		DefaultParams(),
		[]BearingRecord{},
	)
}

// ValidateGenesis validates GenesisState.
func ValidateGenesis(data GenesisState) error {
	if err := data.Params.Validate(); err != nil {
		return err
	}
	for _, record := range data.BearingRecords {
		if err := record.TotalCollectedCoins.Validate(); err != nil {
			return sdkerrors.Wrapf(
				sdkerrors.ErrInvalidCoins,
				"invalid total collected coins %s: %v", record.TotalCollectedCoins, err)
		}
		if err := ValidateName(record.Name); err != nil {
			return err
		}
	}
	return nil
}
