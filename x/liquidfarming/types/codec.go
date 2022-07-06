package types

import (
	"github.com/cosmos/cosmos-sdk/codec"
	cdctypes "github.com/cosmos/cosmos-sdk/codec/types"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/cosmos/cosmos-sdk/types/msgservice"
	govtypes "github.com/cosmos/cosmos-sdk/x/gov/types"
)

// RegisterLegacyAminoCodec registers the necessary x/liquidfarming interfaces and concrete types
// on the provided LegacyAmino codec. These types are used for Amino JSON serialization.
func RegisterLegacyAminoCodec(cdc *codec.LegacyAmino) {
	cdc.RegisterConcrete(&MsgDeposit{}, "liquidfarming/MsgDeposit", nil)
	cdc.RegisterConcrete(&MsgCancel{}, "liquidfarming/MsgCancel", nil)
	cdc.RegisterConcrete(&MsgWithdraw{}, "liquidfarming/MsgWithdraw", nil)
	cdc.RegisterConcrete(&MsgPlaceBid{}, "liquidfarming/MsgPlaceBid", nil)
}

// RegisterInterfaces registers the x/liquidfarming interfaces types with the interface registry
func RegisterInterfaces(registry cdctypes.InterfaceRegistry) {
	registry.RegisterImplementations(
		(*sdk.Msg)(nil),
		&MsgDeposit{},
		&MsgCancel{},
		&MsgWithdraw{},
		&MsgPlaceBid{},
	)

	registry.RegisterImplementations(
		(*govtypes.Content)(nil),
		&LiquidFarmProposal{},
	)

	msgservice.RegisterMsgServiceDesc(registry, &_Msg_serviceDesc)
}

var (
	Amino     = codec.NewLegacyAmino()
	ModuleCdc = codec.NewProtoCodec(cdctypes.NewInterfaceRegistry())
)
