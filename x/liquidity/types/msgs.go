package types

import (
	sdk "github.com/cosmos/cosmos-sdk/types"
)

func (req SwapRequest) GetRequester() sdk.AccAddress {
	addr, err := sdk.AccAddressFromBech32(req.Requester)
	if err != nil {
		panic(err)
	}
	return addr
}
