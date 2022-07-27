package types

import (
	fmt "fmt"

	sdk "github.com/cosmos/cosmos-sdk/types"
)

// DefaultGenesis returns the default genesis state.
func DefaultGenesis() *GenesisState {
	return &GenesisState{
		Params:               DefaultParams(),
		QueuedFarmingRecords: []QueuedFarmingRecord{},
		RewardsAuctions:      []RewardsAuction{},
		Bids:                 []Bid{},
		WinningBidRecords:    []WinningBidRecord{},
	}
}

// Validate performs basic genesis state validation returning an error upon any failure.
func (gs GenesisState) Validate() error {
	if err := gs.Params.Validate(); err != nil {
		return fmt.Errorf("invalid params: %w", err)
	}

	for _, record := range gs.QueuedFarmingRecords {
		if _, err := sdk.AccAddressFromBech32(record.Farmer); err != nil {
			return fmt.Errorf("invalid farmer address %s: %w", record.Farmer, err)
		}
		if err := sdk.ValidateDenom(record.FarmingCoinDenom); err != nil {
			return fmt.Errorf("invalid farming coin denom")
		}
		if record.QueuedFarming.PoolId == 0 {
			return fmt.Errorf("pool id must not be 0")
		}
		if !record.QueuedFarming.Amount.IsPositive() {
			return fmt.Errorf("amount must be positive value")
		}
	}

	for _, auction := range gs.RewardsAuctions {
		if err := auction.Validate(); err != nil {
			return err
		}
	}

	for _, bid := range gs.Bids {
		if err := bid.Validate(); err != nil {
			return err
		}
	}

	winningBidMap := map[uint64]Bid{}
	for i, record := range gs.WinningBidRecords {
		if record.AuctionId == 0 {
			return fmt.Errorf("auction id must not be 0")
		}
		if err := record.WinningBid.Validate(); err != nil {
			return fmt.Errorf("invalid bid: %w", err)
		}
		if _, ok := winningBidMap[record.AuctionId]; ok {
			return fmt.Errorf("bid at %d has a duplicate id: %d", i, record.AuctionId)
		}
	}

	return nil
}
