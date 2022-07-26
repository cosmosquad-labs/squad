package types

// DefaultGenesis returns the default genesis state for the lending module.
func DefaultGenesis() *GenesisState {
	return &GenesisState{
		Params: DefaultParams(),
	}
}

// Validate performs basic genesis state validation returning an error upon any
// failure.
func (genState GenesisState) Validate() error {
	return nil
}
