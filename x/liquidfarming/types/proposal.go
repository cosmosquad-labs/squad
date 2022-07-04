package types

import (
	fmt "fmt"
	"strings"

	"gopkg.in/yaml.v2"

	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
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

func NewLiquidFarmProposal(title string, description string, liquidfarms []LiquidFarm) *LiquidFarmProposal {
	return &LiquidFarmProposal{
		Title:       title,
		Description: description,
		Liquidfarms: liquidfarms,
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

	if len(p.Liquidfarms) == 0 {
		return ErrEmptyLiquidFarms
	}

	for _, l := range p.Liquidfarms {
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
`, p.Title, p.Description, p.Liquidfarms)
}

func NewLiquidFarm(id uint64, poolId uint64, poolCoinDenom string, lfCoinDenom string, reserveAddr string) LiquidFarm {
	return LiquidFarm{
		Id:             id,
		PoolId:         poolId,
		PoolCoinDenom:  poolCoinDenom,
		LfCoinDenom:    lfCoinDenom,
		ReserveAddress: reserveAddr,
	}
}

func (l LiquidFarm) String() string {
	out, _ := yaml.Marshal(l)
	return string(out)
}

// TODO: double check with these validity checks
func (l LiquidFarm) Validate() error {
	if l.Id == 0 {
		return sdkerrors.Wrap(sdkerrors.ErrInvalidRequest, "invalid id")
	}
	if l.PoolId == 0 {
		return sdkerrors.Wrap(sdkerrors.ErrInvalidRequest, "invalid pool id")
	}
	if err := sdk.ValidateDenom(l.PoolCoinDenom); err != nil {
		return sdkerrors.Wrapf(sdkerrors.ErrInvalidRequest, "invalid pool coin denom: %v", err)
	}
	if err := sdk.ValidateDenom(l.LfCoinDenom); err != nil {
		return sdkerrors.Wrapf(sdkerrors.ErrInvalidRequest, "invalid liquid farming coin denom: %v", err)
	}
	if !strings.HasPrefix(l.PoolCoinDenom, "pool") {
		return sdkerrors.Wrap(sdkerrors.ErrInvalidRequest, "invalid pool coin denom")
	}
	if !strings.HasPrefix(l.LfCoinDenom, "lf") {
		return sdkerrors.Wrap(sdkerrors.ErrInvalidRequest, "invalid liquid farming coin denom")
	}
	return nil
}
