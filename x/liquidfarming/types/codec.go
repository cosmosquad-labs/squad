package types

import (
	"github.com/cosmos/cosmos-sdk/codec"
	"github.com/cosmos/cosmos-sdk/codec/types"

	auctiontypes "github.com/cosmosquad-labs/squad/x/auction/types"
)

func RegisterCodec(cdc *codec.LegacyAmino) {
	cdc.RegisterConcrete(&FixedPriceAuction{}, "squad/FixedPriceAuction", nil)
}

func RegisterInterfaces(registry types.InterfaceRegistry) {
	registry.RegisterImplementations(
		(*auctiontypes.Custom)(nil),
		&FixedPriceAuction{},
	)
}
