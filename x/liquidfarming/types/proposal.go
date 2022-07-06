package types

import (
	fmt "fmt"

	govtypes "github.com/cosmos/cosmos-sdk/x/gov/types"
)

const (
	ProposalTypeLiquidFarm string = "LiquidFarm"
)

var _ govtypes.Content = &LiquidFarmProposal{}

func init() {
	govtypes.RegisterProposalType(ProposalTypeLiquidFarm)
	govtypes.RegisterProposalTypeCodec(&LiquidFarmProposal{}, "squad/LiquidFarmProposal")
}

// NewLiquidFarmProposal creates a new LiquidFarmProposal
func NewLiquidFarmProposal(title string, description string, liquidFarms []LiquidFarm) *LiquidFarmProposal {
	return &LiquidFarmProposal{
		Title:       title,
		Description: description,
		LiquidFarms: liquidFarms,
	}
}

func (p *LiquidFarmProposal) GetTitle() string { return p.Title }

func (p *LiquidFarmProposal) GetDescription() string { return p.Description }

func (p *LiquidFarmProposal) ProposalRoute() string { return RouterKey }

func (p *LiquidFarmProposal) ProposalType() string { return ProposalTypeLiquidFarm }

func (p *LiquidFarmProposal) ValidateBasic() error {
	if err := govtypes.ValidateAbstract(p); err != nil {
		return err
	}

	if len(p.LiquidFarms) == 0 {
		return ErrEmptyLiquidFarms
	}

	for _, l := range p.LiquidFarms {
		if err := l.Validate(); err != nil {
			return err
		}
	}
	return nil
}

func (p *LiquidFarmProposal) String() string {
	return fmt.Sprintf(`LiquidFarm Proposal:
  Title:              %s
  Description:        %s
  LiquidFarms:    %v
`, p.Title, p.Description, p.LiquidFarms)
}
