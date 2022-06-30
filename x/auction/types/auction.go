package types

import (
	"fmt"
	time "time"

	"github.com/cosmos/cosmos-sdk/codec/types"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/gogo/protobuf/proto"
)

const DefaultStartingAuctionID uint64 = 1

func NewAuction(custom Custom, id uint64, auctioneer string, status AuctionStatus, startTime, endTime time.Time) (Auction, error) {
	msg, ok := custom.(proto.Message)
	if !ok {
		return Auction{}, fmt.Errorf("%T does not implement proto.Message", custom)
	}

	any, err := types.NewAnyWithValue(msg)
	if err != nil {
		return Auction{}, err
	}

	auction := Auction{
		Custom:     any,
		Id:         id,
		Auctioneer: auctioneer,
		Status:     status,
		StartTime:  startTime,
		EndTime:    endTime,
	}

	return auction, nil
}

// GetCustom returns the auction Custom
func (a Auction) GetCustom() Custom {
	custom, ok := a.Custom.GetCachedValue().(Custom)
	if !ok {
		return nil
	}
	return custom
}

func (a Auction) AuctionId() uint64 {
	return a.Id
}

func (a Auction) AuctionType() string {
	custom := a.GetCustom()
	if custom == nil {
		return ""
	}
	return custom.AuctionType()
}

func (a Auction) GetAuctioneer() string {
	addr, err := sdk.AccAddressFromBech32(a.Auctioneer)
	if err != nil {
		panic(err)
	}
	return addr.String()
}

func (a Auction) GetStartTime() time.Time {
	return a.StartTime
}

func (a Auction) GetEndTime() time.Time {
	return a.EndTime
}

// TODO: proto file must be updated to getters all false
// func (a Auction) String() string {
// 	out, _ := yaml.Marshal(p)
// 	return string(out)
// }
