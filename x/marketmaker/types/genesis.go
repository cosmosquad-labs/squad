package types

// NewGenesisState returns new GenesisState.
func NewGenesisState(
	params Params, marketMakers []MarketMaker, incentives []Incentive, depositRecords []DepositRecord,
) *GenesisState {
	return &GenesisState{
		Params:         params,
		MarketMakers:   marketMakers,
		Incentives:     incentives,
		DepositRecords: depositRecords,
	}
}

// DefaultGenesisState returns the default genesis state.
func DefaultGenesisState() *GenesisState {
	return NewGenesisState(
		DefaultParams(),
		[]MarketMaker{},
		[]Incentive{},
		[]DepositRecord{},
	)
}

// ValidateGenesis validates GenesisState.
func ValidateGenesis(data GenesisState) error {
	if err := data.Params.Validate(); err != nil {
		return err
	}

	for _, record := range data.MarketMakers {
		if err := record.Validate(); err != nil {
			panic(err)
		}
	}

	for _, record := range data.Incentives {
		if err := record.Validate(); err != nil {
			panic(err)
		}
	}

	for _, record := range data.DepositRecords {
		if err := record.Validate(); err != nil {
			panic(err)
		}
	}
	return nil
}
