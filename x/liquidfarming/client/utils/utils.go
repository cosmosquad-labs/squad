package utils

import (
	"io/ioutil"

	"github.com/cosmos/cosmos-sdk/codec"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/cosmos/cosmos-sdk/types/rest"

	"github.com/cosmosquad-labs/squad/x/liquidfarming/types"
)

type (
	// LiquidFarmProposalJSON defines a LiquidFarmProposal with a deposit used
	// to parse an array of liquid farm from a JSON file.
	LiquidFarmProposalJSON struct {
		Title       string             `json:"title" yaml:"title"`
		Description string             `json:"description" yaml:"description"`
		LiquidFarms []types.LiquidFarm `json:"liquidfarms" yaml:"liquidfarms"`
		Deposit     string             `json:"deposit" yaml:"deposit"`
	}

	// LiquidFarmProposalReq defines a liquidfarm proposal request body.
	LiquidFarmProposalReq struct {
		BaseReq rest.BaseReq `json:"base_req" yaml:"base_req"`

		Title       string             `json:"title" yaml:"title"`
		Description string             `json:"description" yaml:"description"`
		LiquidFarms []types.LiquidFarm `json:"liquidfarms" yaml:"liquidfarms"`
		Proposer    sdk.AccAddress     `json:"proposer" yaml:"proposer"`
		Deposit     sdk.Coins          `json:"deposit" yaml:"deposit"`
	}
)

// ParseLiquidFarmProposalJSON reads and parses a LiquidFarmProposalJSON from file.
func ParseLiquidFarmProposalJSON(cdc *codec.LegacyAmino, proposalFile string) (LiquidFarmProposalJSON, error) {
	proposal := LiquidFarmProposalJSON{}

	contents, err := ioutil.ReadFile(proposalFile)
	if err != nil {
		return proposal, err
	}

	if err := cdc.UnmarshalJSON(contents, &proposal); err != nil {
		return proposal, err
	}

	return proposal, nil
}
