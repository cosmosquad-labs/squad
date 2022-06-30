package types

import (
	"fmt"

	"github.com/cosmos/cosmos-sdk/codec/legacy"
	"github.com/cosmos/cosmos-sdk/codec/types"
	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
	"github.com/gogo/protobuf/proto"
	"gopkg.in/yaml.v2"
)

var (
	_ sdk.Msg                       = (*MsgCreateAuction)(nil)
	_ types.UnpackInterfacesMessage = &MsgCreateAuction{}
)

const (
	TypeMsgCreateAuction = "create_auction"
)

func NewMsgCreateAuction(custom Custom, auctioneer string) (*MsgCreateAuction, error) {
	m := &MsgCreateAuction{
		Auctioneer: auctioneer,
	}

	err := m.SetCustom(custom)
	if err != nil {
		return nil, err
	}

	return m, nil
}

func (m *MsgCreateAuction) GetCustom() Custom {
	custom, ok := m.Custom.GetCachedValue().(Custom)
	if !ok {
		return nil
	}
	return custom
}

func (msg MsgCreateAuction) GetAuctioneer() sdk.AccAddress {
	addr, err := sdk.AccAddressFromBech32(msg.Auctioneer)
	if err != nil {
		panic(err)
	}
	return addr
}

func (m *MsgCreateAuction) SetCustom(custom Custom) error {
	msg, ok := custom.(proto.Message)
	if !ok {
		return fmt.Errorf("can't proto marshal %T", msg)
	}
	any, err := types.NewAnyWithValue(msg)
	if err != nil {
		return err
	}
	m.Custom = any
	return nil
}

func (msg MsgCreateAuction) Route() string { return RouterKey }

func (msg MsgCreateAuction) Type() string { return TypeMsgCreateAuction }

func (msg MsgCreateAuction) ValidateBasic() error {
	if _, err := sdk.AccAddressFromBech32(msg.Auctioneer); err != nil {
		return sdkerrors.Wrapf(sdkerrors.ErrInvalidAddress, "invalid auctioneer address: %v", err)
	}
	return nil
}

func (msg MsgCreateAuction) GetSignBytes() []byte {
	return sdk.MustSortJSON(legacy.Cdc.MustMarshalJSON(&msg))
}

func (msg MsgCreateAuction) GetSigners() []sdk.AccAddress {
	addr, err := sdk.AccAddressFromBech32(msg.Auctioneer)
	if err != nil {
		panic(err)
	}
	return []sdk.AccAddress{addr}
}

// String implements the Stringer interface
func (m MsgCreateAuction) String() string {
	out, _ := yaml.Marshal(m)
	return string(out)
}

// UnpackInterfaces implements UnpackInterfacesMessage.UnpackInterfaces
func (m MsgCreateAuction) UnpackInterfaces(unpacker types.AnyUnpacker) error {
	var custom Custom
	return unpacker.UnpackAny(m.Custom, &custom)
}
