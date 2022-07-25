package types_test

import (
	"testing"

	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/stretchr/testify/require"
	"github.com/tendermint/tendermint/crypto"

	utils "github.com/cosmosquad-labs/squad/v2/types"
	"github.com/cosmosquad-labs/squad/v2/x/liquidfarming/types"
)

var testAddr = sdk.AccAddress(crypto.AddressHash([]byte("test")))

func TestMsgFarm(t *testing.T) {
	for _, tc := range []struct {
		name        string
		malleate    func(msg *types.MsgFarm)
		expectedErr string
	}{
		{
			"happy case",
			func(msg *types.MsgFarm) {},
			"",
		},
		// TODO: not implemented yet
	} {
		t.Run(tc.name, func(t *testing.T) {
			msg := types.NewMsgFarm(1, testAddr.String(), utils.ParseCoin("1000000pool1"))
			tc.malleate(msg)
			require.Equal(t, types.TypeMsgFarm, msg.Type())
			require.Equal(t, types.RouterKey, msg.Route())
			err := msg.ValidateBasic()
			if tc.expectedErr == "" {
				require.NoError(t, err)
				signers := msg.GetSigners()
				require.Len(t, signers, 1)
				require.Equal(t, msg.GetFarmer(), signers[0])
			} else {
				require.EqualError(t, err, tc.expectedErr)
			}
		})
	}
}

func TestMsgUnfarm(t *testing.T) {
	for _, tc := range []struct {
		name        string
		malleate    func(msg *types.MsgUnfarm)
		expectedErr string
	}{
		{
			"happy case",
			func(msg *types.MsgUnfarm) {},
			"",
		},
		// TODO: not implemented yet
	} {
		t.Run(tc.name, func(t *testing.T) {
			msg := types.NewMsgUnfarm(1, testAddr.String(), utils.ParseCoin("1000000lf1"))
			tc.malleate(msg)
			require.Equal(t, types.TypeMsgUnfarm, msg.Type())
			require.Equal(t, types.RouterKey, msg.Route())
			err := msg.ValidateBasic()
			if tc.expectedErr == "" {
				require.NoError(t, err)
				signers := msg.GetSigners()
				require.Len(t, signers, 1)
				require.Equal(t, msg.GetFarmer(), signers[0])
			} else {
				require.EqualError(t, err, tc.expectedErr)
			}
		})
	}
}

func TestMsgCancelQueuedFarming(t *testing.T) {
	for _, tc := range []struct {
		name        string
		malleate    func(msg *types.MsgCancelQueuedFarming)
		expectedErr string
	}{
		{
			"happy case",
			func(msg *types.MsgCancelQueuedFarming) {},
			"",
		},
		// TODO: not implemented yet
	} {
		t.Run(tc.name, func(t *testing.T) {
			msg := types.NewMsgCancelQueuedFarming(1, testAddr.String(), utils.ParseCoin("1000000lf1"))
			tc.malleate(msg)
			require.Equal(t, types.TypeMsgCancelQueuedFarming, msg.Type())
			require.Equal(t, types.RouterKey, msg.Route())
			err := msg.ValidateBasic()
			if tc.expectedErr == "" {
				require.NoError(t, err)
				signers := msg.GetSigners()
				require.Len(t, signers, 1)
				require.Equal(t, msg.GetFarmer(), signers[0])
			} else {
				require.EqualError(t, err, tc.expectedErr)
			}
		})
	}
}

func TestMsgPlaceBid(t *testing.T) {
	for _, tc := range []struct {
		name        string
		malleate    func(msg *types.MsgPlaceBid)
		expectedErr string
	}{
		{
			"happy case",
			func(msg *types.MsgPlaceBid) {},
			"",
		},
		// TODO: not implemented yet
	} {
		t.Run(tc.name, func(t *testing.T) {
			msg := types.NewMsgPlaceBid(1, testAddr.String(), utils.ParseCoin("1000000pool1"))
			tc.malleate(msg)
			require.Equal(t, types.TypeMsgPlaceBid, msg.Type())
			require.Equal(t, types.RouterKey, msg.Route())
			err := msg.ValidateBasic()
			if tc.expectedErr == "" {
				require.NoError(t, err)
				signers := msg.GetSigners()
				require.Len(t, signers, 1)
				require.Equal(t, msg.GetBidder(), signers[0])
			} else {
				require.EqualError(t, err, tc.expectedErr)
			}
		})
	}
}

func TestMsgRefundBid(t *testing.T) {
	for _, tc := range []struct {
		name        string
		malleate    func(msg *types.MsgRefundBid)
		expectedErr string
	}{
		{
			"happy case",
			func(msg *types.MsgRefundBid) {},
			"",
		},
		// TODO: not implemented yet
	} {
		t.Run(tc.name, func(t *testing.T) {
			msg := types.NewMsgRefundBid(1, testAddr.String())
			tc.malleate(msg)
			require.Equal(t, types.TypeMsgRefundBid, msg.Type())
			require.Equal(t, types.RouterKey, msg.Route())
			err := msg.ValidateBasic()
			if tc.expectedErr == "" {
				require.NoError(t, err)
				signers := msg.GetSigners()
				require.Len(t, signers, 1)
				require.Equal(t, msg.GetBidder(), signers[0])
			} else {
				require.EqualError(t, err, tc.expectedErr)
			}
		})
	}
}
