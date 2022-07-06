package keeper

import (
	sdk "github.com/cosmos/cosmos-sdk/types"

	"github.com/cosmosquad-labs/squad/x/liquidfarming/types"
)

// CreateLiquidFarmProposal is a handler for executing a liquid farm creation proposal.
func CreateLiquidFarmProposal(ctx sdk.Context, k Keeper, proposal *types.LiquidFarmProposal) error {
	// Title       string       `protobuf:"bytes,1,opt,name=title,proto3" json:"title,omitempty"`
	// Description string       `protobuf:"bytes,2,opt,name=description,proto3" json:"description,omitempty"`
	// LiquidFarms []LiquidFarm `protobuf:"bytes,3,rep,name=liquidfarms,proto3" json:"liquidfarms"`

	// LiquidFarm
	// Id             uint64 `protobuf:"varint,1,opt,name=id,proto3" json:"id,omitempty"`
	// PoolId         uint64 `protobuf:"varint,2,opt,name=pool_id,json=poolId,proto3" json:"pool_id,omitempty"`
	// PoolCoinDenom  string `protobuf:"bytes,3,opt,name=pool_coin_denom,json=poolCoinDenom,proto3" json:"pool_coin_denom,omitempty"`
	// LFCoinDenom    string `protobuf:"bytes,4,opt,name=lf_coin_denom,json=lfCoinDenom,proto3" json:"lf_coin_denom,omitempty"`
	// ReserveAddress string `protobuf:"bytes,5,opt,name=reserve_address,json=reserveAddress,proto3" json:"reserve_address,omitempty"`

	for _, lf := range proposal.LiquidFarms {
		lf.
	}

	return nil
}
