// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: squad/liquidfarming/v1beta1/liquidfarming.proto

package types

import (
	fmt "fmt"
	github_com_cosmos_cosmos_sdk_types "github.com/cosmos/cosmos-sdk/types"
	types "github.com/cosmos/cosmos-sdk/types"
	_ "github.com/gogo/protobuf/gogoproto"
	proto "github.com/gogo/protobuf/proto"
	github_com_gogo_protobuf_types "github.com/gogo/protobuf/types"
	_ "github.com/regen-network/cosmos-proto"
	_ "google.golang.org/protobuf/types/known/timestamppb"
	io "io"
	math "math"
	math_bits "math/bits"
	time "time"
)

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf
var _ = time.Kitchen

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.GoGoProtoPackageIsVersion3 // please upgrade the proto package

// AuctionStatus enumerates the valid status of an auction.
type AuctionStatus int32

const (
	// AUCTION_STATUS_UNSPECIFIED defines the default auction status
	AuctionStatusNil AuctionStatus = 0
	// AUCTION_STATUS_STARTED defines the started auction status
	AuctionStatusStarted AuctionStatus = 1
	// AUCTION_STATUS_FINISHED defines the finished auction status
	AuctionStatusFinished AuctionStatus = 2
)

var AuctionStatus_name = map[int32]string{
	0: "AUCTION_STATUS_UNSPECIFIED",
	1: "AUCTION_STATUS_STARTED",
	2: "AUCTION_STATUS_FINISHED",
}

var AuctionStatus_value = map[string]int32{
	"AUCTION_STATUS_UNSPECIFIED": 0,
	"AUCTION_STATUS_STARTED":     1,
	"AUCTION_STATUS_FINISHED":    2,
}

func (x AuctionStatus) String() string {
	return proto.EnumName(AuctionStatus_name, int32(x))
}

func (AuctionStatus) EnumDescriptor() ([]byte, []int) {
	return fileDescriptor_b3445e3599d3c045, []int{0}
}

// QueuedFarming defines queued farming.
type QueuedFarming struct {
	// pool_id specifies the pool id
	PoolId uint64 `protobuf:"varint,1,opt,name=pool_id,json=poolId,proto3" json:"pool_id,omitempty"`
	// id specifies the id for the request
	Id uint64 `protobuf:"varint,2,opt,name=id,proto3" json:"id,omitempty"`
	// farmer specifies the bech32-encoded address that makes a deposit to the pool
	Farmer string `protobuf:"bytes,3,opt,name=farmer,proto3" json:"farmer,omitempty"`
	// farming_coin specifies the amount of pool coin to deposit
	FarmingCoin types.Coin `protobuf:"bytes,4,opt,name=farming_coin,json=farmingCoin,proto3,castrepeated=github.com/cosmos/cosmos-sdk/types.Coin" json:"farming_coin"`
}

func (m *QueuedFarming) Reset()         { *m = QueuedFarming{} }
func (m *QueuedFarming) String() string { return proto.CompactTextString(m) }
func (*QueuedFarming) ProtoMessage()    {}
func (*QueuedFarming) Descriptor() ([]byte, []int) {
	return fileDescriptor_b3445e3599d3c045, []int{0}
}
func (m *QueuedFarming) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *QueuedFarming) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_QueuedFarming.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *QueuedFarming) XXX_Merge(src proto.Message) {
	xxx_messageInfo_QueuedFarming.Merge(m, src)
}
func (m *QueuedFarming) XXX_Size() int {
	return m.Size()
}
func (m *QueuedFarming) XXX_DiscardUnknown() {
	xxx_messageInfo_QueuedFarming.DiscardUnknown(m)
}

var xxx_messageInfo_QueuedFarming proto.InternalMessageInfo

// RewardsAuction defines rewards auction that is created by the module account
// at an end block for every epoch.
type RewardsAuction struct {
	// id specifies the id for the auction
	Id uint64 `protobuf:"varint,1,opt,name=id,proto3" json:"id,omitempty"`
	// pool_id specifies the pool id
	PoolId                uint64                                   `protobuf:"varint,2,opt,name=pool_id,json=poolId,proto3" json:"pool_id,omitempty"`
	SellingRewards        github_com_cosmos_cosmos_sdk_types.Coins `protobuf:"bytes,3,rep,name=selling_rewards,json=sellingRewards,proto3,castrepeated=github.com/cosmos/cosmos-sdk/types.Coins" json:"selling_rewards"`
	BiddingCoinDenom      string                                   `protobuf:"bytes,4,opt,name=bidding_coin_denom,json=biddingCoinDenom,proto3" json:"bidding_coin_denom,omitempty"`
	SellingReserveAddress string                                   `protobuf:"bytes,5,opt,name=selling_reserve_address,json=sellingReserveAddress,proto3" json:"selling_reserve_address,omitempty"`
	PayingReserveAddress  string                                   `protobuf:"bytes,6,opt,name=paying_reserve_address,json=payingReserveAddress,proto3" json:"paying_reserve_address,omitempty"`
	StartTime             *time.Time                               `protobuf:"bytes,7,opt,name=start_time,json=startTime,proto3,stdtime" json:"start_time,omitempty" yaml:"start_time"`
	EndTime               *time.Time                               `protobuf:"bytes,8,opt,name=end_time,json=endTime,proto3,stdtime" json:"end_time,omitempty" yaml:"end_time"`
	Status                AuctionStatus                            `protobuf:"varint,9,opt,name=status,proto3,enum=squad.liquidfarming.v1beta1.AuctionStatus" json:"status,omitempty"`
	WinnerBidId           uint64                                   `protobuf:"varint,10,opt,name=winner_bid_id,json=winnerBidId,proto3" json:"winner_bid_id,omitempty"`
}

func (m *RewardsAuction) Reset()         { *m = RewardsAuction{} }
func (m *RewardsAuction) String() string { return proto.CompactTextString(m) }
func (*RewardsAuction) ProtoMessage()    {}
func (*RewardsAuction) Descriptor() ([]byte, []int) {
	return fileDescriptor_b3445e3599d3c045, []int{1}
}
func (m *RewardsAuction) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *RewardsAuction) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_RewardsAuction.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *RewardsAuction) XXX_Merge(src proto.Message) {
	xxx_messageInfo_RewardsAuction.Merge(m, src)
}
func (m *RewardsAuction) XXX_Size() int {
	return m.Size()
}
func (m *RewardsAuction) XXX_DiscardUnknown() {
	xxx_messageInfo_RewardsAuction.DiscardUnknown(m)
}

var xxx_messageInfo_RewardsAuction proto.InternalMessageInfo

// Bid defines ...
type Bid struct {
	Id        uint64     `protobuf:"varint,1,opt,name=id,proto3" json:"id,omitempty"`
	AuctionId uint64     `protobuf:"varint,2,opt,name=auction_id,json=auctionId,proto3" json:"auction_id,omitempty"`
	Bidder    string     `protobuf:"bytes,3,opt,name=bidder,proto3" json:"bidder,omitempty"`
	Amount    types.Coin `protobuf:"bytes,4,opt,name=amount,proto3,castrepeated=github.com/cosmos/cosmos-sdk/types.Coin" json:"amount"`
	IsMatched bool       `protobuf:"varint,5,opt,name=is_matched,json=isMatched,proto3" json:"is_matched,omitempty"`
}

func (m *Bid) Reset()         { *m = Bid{} }
func (m *Bid) String() string { return proto.CompactTextString(m) }
func (*Bid) ProtoMessage()    {}
func (*Bid) Descriptor() ([]byte, []int) {
	return fileDescriptor_b3445e3599d3c045, []int{2}
}
func (m *Bid) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *Bid) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_Bid.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *Bid) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Bid.Merge(m, src)
}
func (m *Bid) XXX_Size() int {
	return m.Size()
}
func (m *Bid) XXX_DiscardUnknown() {
	xxx_messageInfo_Bid.DiscardUnknown(m)
}

var xxx_messageInfo_Bid proto.InternalMessageInfo

func init() {
	proto.RegisterEnum("squad.liquidfarming.v1beta1.AuctionStatus", AuctionStatus_name, AuctionStatus_value)
	proto.RegisterType((*QueuedFarming)(nil), "squad.liquidfarming.v1beta1.QueuedFarming")
	proto.RegisterType((*RewardsAuction)(nil), "squad.liquidfarming.v1beta1.RewardsAuction")
	proto.RegisterType((*Bid)(nil), "squad.liquidfarming.v1beta1.Bid")
}

func init() {
	proto.RegisterFile("squad/liquidfarming/v1beta1/liquidfarming.proto", fileDescriptor_b3445e3599d3c045)
}

var fileDescriptor_b3445e3599d3c045 = []byte{
	// 763 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xb4, 0x94, 0x4f, 0x6f, 0xe3, 0x44,
	0x18, 0xc6, 0xed, 0x24, 0xa4, 0xcd, 0x94, 0x66, 0x83, 0x95, 0xdd, 0x75, 0xbd, 0x5a, 0xc7, 0xca,
	0x85, 0x68, 0xc5, 0xda, 0xec, 0x52, 0xf5, 0xd0, 0x5b, 0xd2, 0x24, 0xc2, 0x07, 0x02, 0x38, 0xa9,
	0x84, 0xb8, 0x58, 0x76, 0x66, 0x9a, 0x8e, 0xb0, 0x3d, 0xa9, 0xc7, 0x6e, 0xe9, 0x37, 0xa8, 0x7a,
	0xea, 0x17, 0xa8, 0x84, 0xc4, 0x8d, 0x6f, 0xc0, 0x89, 0x6b, 0x8e, 0xbd, 0x20, 0x71, 0x6a, 0x69,
	0xfb, 0x0d, 0xf8, 0x04, 0x68, 0xfe, 0xa4, 0x21, 0x29, 0x2a, 0x5c, 0xf6, 0x94, 0xbc, 0xef, 0x3b,
	0xbf, 0xc7, 0xcf, 0x8c, 0x9f, 0x31, 0x70, 0xe8, 0x51, 0x1e, 0x40, 0x27, 0xc2, 0x47, 0x39, 0x86,
	0x07, 0x41, 0x1a, 0xe3, 0x64, 0xe2, 0x1c, 0xbf, 0x0b, 0x51, 0x16, 0xbc, 0x5b, 0xee, 0xda, 0xd3,
	0x94, 0x64, 0x44, 0x7b, 0xc5, 0x01, 0x7b, 0x79, 0x24, 0x01, 0xa3, 0x3e, 0x21, 0x13, 0xc2, 0xd7,
	0x39, 0xec, 0x9f, 0x40, 0x8c, 0xad, 0x31, 0xa1, 0x31, 0xa1, 0xbe, 0x18, 0x88, 0x42, 0x8e, 0x4c,
	0x51, 0x39, 0x61, 0x40, 0xd1, 0xc3, 0x63, 0xc7, 0x04, 0x27, 0x72, 0xde, 0x98, 0x10, 0x32, 0x89,
	0x90, 0xc3, 0xab, 0x30, 0x3f, 0x70, 0x32, 0x1c, 0x23, 0x9a, 0x05, 0xf1, 0x54, 0x2c, 0x68, 0xfe,
	0xa6, 0x82, 0xcd, 0x6f, 0x73, 0x94, 0x23, 0xd8, 0x17, 0x5e, 0xb4, 0x97, 0x60, 0x6d, 0x4a, 0x48,
	0xe4, 0x63, 0xa8, 0xab, 0x96, 0xda, 0x2a, 0x79, 0x65, 0x56, 0xba, 0x50, 0xab, 0x82, 0x02, 0x86,
	0x7a, 0x81, 0xf7, 0x0a, 0x18, 0x6a, 0x2f, 0x40, 0x99, 0xf9, 0x47, 0xa9, 0x5e, 0xb4, 0xd4, 0x56,
	0xc5, 0x93, 0x95, 0x16, 0x83, 0x8f, 0xe5, 0xbe, 0x7c, 0xe6, 0x44, 0x2f, 0x59, 0x6a, 0x6b, 0xe3,
	0xfd, 0x96, 0x2d, 0x8d, 0x33, 0xab, 0xf3, 0x0d, 0xdb, 0x7b, 0x04, 0x27, 0x1d, 0x67, 0x76, 0xdd,
	0x50, 0x7e, 0xb9, 0x69, 0x7c, 0x3a, 0xc1, 0xd9, 0x61, 0x1e, 0xda, 0x63, 0x12, 0xcb, 0x5d, 0xca,
	0x9f, 0xb7, 0x14, 0xfe, 0xe0, 0x64, 0xa7, 0x53, 0x44, 0x39, 0xe0, 0x6d, 0x48, 0x7d, 0x56, 0x34,
	0x6f, 0x4b, 0xa0, 0xea, 0xa1, 0x93, 0x20, 0x85, 0xb4, 0x9d, 0x8f, 0x33, 0x4c, 0x12, 0xe9, 0x54,
	0x7d, 0x70, 0xfa, 0x8f, 0x2d, 0x15, 0x96, 0xb6, 0x94, 0x81, 0x67, 0x14, 0x45, 0x11, 0xb3, 0x9a,
	0x0a, 0x09, 0xbd, 0x68, 0x15, 0x9f, 0x76, 0xfb, 0xb9, 0x74, 0xdb, 0xfa, 0x9f, 0x6e, 0xa9, 0x57,
	0x95, 0xcf, 0x90, 0x2e, 0xb5, 0xcf, 0x80, 0x16, 0x62, 0x08, 0xe7, 0x07, 0xe4, 0x43, 0x94, 0x90,
	0x98, 0x1f, 0x53, 0xc5, 0xab, 0xc9, 0x09, 0x23, 0xbb, 0xac, 0xaf, 0xed, 0x80, 0x97, 0x0b, 0x8f,
	0x14, 0xa5, 0xc7, 0xc8, 0x0f, 0x20, 0x4c, 0x11, 0xa5, 0xfa, 0x47, 0x1c, 0x79, 0xfe, 0x20, 0xcf,
	0xa7, 0x6d, 0x31, 0xd4, 0xb6, 0xc1, 0x8b, 0x69, 0x70, 0xfa, 0x6f, 0x58, 0x99, 0x63, 0x75, 0x31,
	0x5d, 0xa1, 0xbe, 0x03, 0x80, 0x66, 0x41, 0x9a, 0xf9, 0x2c, 0x28, 0xfa, 0x1a, 0x7f, 0x75, 0x86,
	0x2d, 0x52, 0x64, 0xcf, 0x53, 0x64, 0x8f, 0xe6, 0x29, 0xea, 0xbc, 0x9e, 0x5d, 0x37, 0xd4, 0xbf,
	0xae, 0x1b, 0x9f, 0x9c, 0x06, 0x71, 0xb4, 0xdb, 0x5c, 0xb0, 0xcd, 0x8b, 0x9b, 0x86, 0xea, 0x55,
	0x78, 0x83, 0x2d, 0xd7, 0x3c, 0xb0, 0x8e, 0x12, 0x28, 0x74, 0xd7, 0xff, 0x53, 0xf7, 0x95, 0xd4,
	0x7d, 0x26, 0x74, 0xe7, 0xa4, 0x50, 0x5d, 0x43, 0x09, 0xe4, 0x9a, 0x1d, 0x50, 0xa6, 0x59, 0x90,
	0xe5, 0x54, 0xaf, 0x58, 0x6a, 0xab, 0xfa, 0xfe, 0x8d, 0xfd, 0xc4, 0xed, 0xb2, 0x65, 0x3c, 0x86,
	0x9c, 0xf0, 0x24, 0xa9, 0x35, 0xc1, 0xe6, 0x09, 0x4e, 0x12, 0x94, 0xfa, 0x21, 0x86, 0x2c, 0x22,
	0x80, 0x47, 0x64, 0x43, 0x34, 0x3b, 0x18, 0xba, 0xb0, 0xf9, 0xbb, 0x0a, 0x8a, 0x1d, 0x0c, 0x1f,
	0x05, 0xeb, 0x35, 0x00, 0x81, 0x10, 0x5d, 0x64, 0xab, 0x22, 0x3b, 0x2e, 0xbf, 0x21, 0xec, 0x75,
	0x2e, 0x6e, 0x88, 0xa8, 0xb4, 0x10, 0x94, 0x83, 0x98, 0xe4, 0x49, 0xf6, 0x01, 0xee, 0x86, 0x54,
	0x66, 0xd6, 0x30, 0xf5, 0xe3, 0x20, 0x1b, 0x1f, 0x22, 0xc8, 0x93, 0xb2, 0xee, 0x55, 0x30, 0xfd,
	0x4a, 0x34, 0x76, 0x4b, 0x67, 0x3f, 0x35, 0x94, 0x37, 0xbf, 0xaa, 0x60, 0x73, 0xe9, 0x54, 0xb4,
	0x6d, 0x60, 0xb4, 0xf7, 0xf7, 0x46, 0xee, 0xd7, 0x03, 0x7f, 0x38, 0x6a, 0x8f, 0xf6, 0x87, 0xfe,
	0xfe, 0x60, 0xf8, 0x4d, 0x6f, 0xcf, 0xed, 0xbb, 0xbd, 0x6e, 0x4d, 0x31, 0xea, 0xe7, 0x97, 0x56,
	0x6d, 0x09, 0x19, 0xe0, 0x88, 0x65, 0x6d, 0x85, 0x1a, 0x8e, 0xda, 0xde, 0xa8, 0xd7, 0xad, 0xa9,
	0x86, 0x7e, 0x7e, 0x69, 0xd5, 0x97, 0x88, 0x21, 0xcb, 0x04, 0x82, 0x2c, 0xd9, 0x2b, 0x54, 0xdf,
	0x1d, 0xb8, 0xc3, 0x2f, 0x7b, 0xdd, 0x5a, 0xc1, 0xd8, 0x3a, 0xbf, 0xb4, 0x9e, 0x2f, 0x61, 0x7d,
	0x9c, 0x60, 0x7a, 0x88, 0xa0, 0x51, 0x3a, 0xfb, 0xd9, 0x54, 0x3a, 0xa3, 0xd9, 0xad, 0xa9, 0xcc,
	0xee, 0x4c, 0xf5, 0xea, 0xce, 0x54, 0xff, 0xbc, 0x33, 0xd5, 0x8b, 0x7b, 0x53, 0xb9, 0xba, 0x37,
	0x95, 0x3f, 0xee, 0x4d, 0xe5, 0xfb, 0x9d, 0x47, 0xe7, 0xc5, 0x82, 0xf1, 0x36, 0x0a, 0x42, 0x2a,
	0x3f, 0xd9, 0x3f, 0xae, 0x7c, 0xb4, 0xf9, 0x19, 0x86, 0x65, 0x9e, 0xc5, 0x2f, 0xfe, 0x0e, 0x00,
	0x00, 0xff, 0xff, 0x10, 0x74, 0x6f, 0x6c, 0xd8, 0x05, 0x00, 0x00,
}

func (m *QueuedFarming) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *QueuedFarming) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *QueuedFarming) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	{
		size, err := m.FarmingCoin.MarshalToSizedBuffer(dAtA[:i])
		if err != nil {
			return 0, err
		}
		i -= size
		i = encodeVarintLiquidfarming(dAtA, i, uint64(size))
	}
	i--
	dAtA[i] = 0x22
	if len(m.Farmer) > 0 {
		i -= len(m.Farmer)
		copy(dAtA[i:], m.Farmer)
		i = encodeVarintLiquidfarming(dAtA, i, uint64(len(m.Farmer)))
		i--
		dAtA[i] = 0x1a
	}
	if m.Id != 0 {
		i = encodeVarintLiquidfarming(dAtA, i, uint64(m.Id))
		i--
		dAtA[i] = 0x10
	}
	if m.PoolId != 0 {
		i = encodeVarintLiquidfarming(dAtA, i, uint64(m.PoolId))
		i--
		dAtA[i] = 0x8
	}
	return len(dAtA) - i, nil
}

func (m *RewardsAuction) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *RewardsAuction) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *RewardsAuction) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if m.WinnerBidId != 0 {
		i = encodeVarintLiquidfarming(dAtA, i, uint64(m.WinnerBidId))
		i--
		dAtA[i] = 0x50
	}
	if m.Status != 0 {
		i = encodeVarintLiquidfarming(dAtA, i, uint64(m.Status))
		i--
		dAtA[i] = 0x48
	}
	if m.EndTime != nil {
		n2, err2 := github_com_gogo_protobuf_types.StdTimeMarshalTo(*m.EndTime, dAtA[i-github_com_gogo_protobuf_types.SizeOfStdTime(*m.EndTime):])
		if err2 != nil {
			return 0, err2
		}
		i -= n2
		i = encodeVarintLiquidfarming(dAtA, i, uint64(n2))
		i--
		dAtA[i] = 0x42
	}
	if m.StartTime != nil {
		n3, err3 := github_com_gogo_protobuf_types.StdTimeMarshalTo(*m.StartTime, dAtA[i-github_com_gogo_protobuf_types.SizeOfStdTime(*m.StartTime):])
		if err3 != nil {
			return 0, err3
		}
		i -= n3
		i = encodeVarintLiquidfarming(dAtA, i, uint64(n3))
		i--
		dAtA[i] = 0x3a
	}
	if len(m.PayingReserveAddress) > 0 {
		i -= len(m.PayingReserveAddress)
		copy(dAtA[i:], m.PayingReserveAddress)
		i = encodeVarintLiquidfarming(dAtA, i, uint64(len(m.PayingReserveAddress)))
		i--
		dAtA[i] = 0x32
	}
	if len(m.SellingReserveAddress) > 0 {
		i -= len(m.SellingReserveAddress)
		copy(dAtA[i:], m.SellingReserveAddress)
		i = encodeVarintLiquidfarming(dAtA, i, uint64(len(m.SellingReserveAddress)))
		i--
		dAtA[i] = 0x2a
	}
	if len(m.BiddingCoinDenom) > 0 {
		i -= len(m.BiddingCoinDenom)
		copy(dAtA[i:], m.BiddingCoinDenom)
		i = encodeVarintLiquidfarming(dAtA, i, uint64(len(m.BiddingCoinDenom)))
		i--
		dAtA[i] = 0x22
	}
	if len(m.SellingRewards) > 0 {
		for iNdEx := len(m.SellingRewards) - 1; iNdEx >= 0; iNdEx-- {
			{
				size, err := m.SellingRewards[iNdEx].MarshalToSizedBuffer(dAtA[:i])
				if err != nil {
					return 0, err
				}
				i -= size
				i = encodeVarintLiquidfarming(dAtA, i, uint64(size))
			}
			i--
			dAtA[i] = 0x1a
		}
	}
	if m.PoolId != 0 {
		i = encodeVarintLiquidfarming(dAtA, i, uint64(m.PoolId))
		i--
		dAtA[i] = 0x10
	}
	if m.Id != 0 {
		i = encodeVarintLiquidfarming(dAtA, i, uint64(m.Id))
		i--
		dAtA[i] = 0x8
	}
	return len(dAtA) - i, nil
}

func (m *Bid) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *Bid) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *Bid) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if m.IsMatched {
		i--
		if m.IsMatched {
			dAtA[i] = 1
		} else {
			dAtA[i] = 0
		}
		i--
		dAtA[i] = 0x28
	}
	{
		size, err := m.Amount.MarshalToSizedBuffer(dAtA[:i])
		if err != nil {
			return 0, err
		}
		i -= size
		i = encodeVarintLiquidfarming(dAtA, i, uint64(size))
	}
	i--
	dAtA[i] = 0x22
	if len(m.Bidder) > 0 {
		i -= len(m.Bidder)
		copy(dAtA[i:], m.Bidder)
		i = encodeVarintLiquidfarming(dAtA, i, uint64(len(m.Bidder)))
		i--
		dAtA[i] = 0x1a
	}
	if m.AuctionId != 0 {
		i = encodeVarintLiquidfarming(dAtA, i, uint64(m.AuctionId))
		i--
		dAtA[i] = 0x10
	}
	if m.Id != 0 {
		i = encodeVarintLiquidfarming(dAtA, i, uint64(m.Id))
		i--
		dAtA[i] = 0x8
	}
	return len(dAtA) - i, nil
}

func encodeVarintLiquidfarming(dAtA []byte, offset int, v uint64) int {
	offset -= sovLiquidfarming(v)
	base := offset
	for v >= 1<<7 {
		dAtA[offset] = uint8(v&0x7f | 0x80)
		v >>= 7
		offset++
	}
	dAtA[offset] = uint8(v)
	return base
}
func (m *QueuedFarming) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	if m.PoolId != 0 {
		n += 1 + sovLiquidfarming(uint64(m.PoolId))
	}
	if m.Id != 0 {
		n += 1 + sovLiquidfarming(uint64(m.Id))
	}
	l = len(m.Farmer)
	if l > 0 {
		n += 1 + l + sovLiquidfarming(uint64(l))
	}
	l = m.FarmingCoin.Size()
	n += 1 + l + sovLiquidfarming(uint64(l))
	return n
}

func (m *RewardsAuction) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	if m.Id != 0 {
		n += 1 + sovLiquidfarming(uint64(m.Id))
	}
	if m.PoolId != 0 {
		n += 1 + sovLiquidfarming(uint64(m.PoolId))
	}
	if len(m.SellingRewards) > 0 {
		for _, e := range m.SellingRewards {
			l = e.Size()
			n += 1 + l + sovLiquidfarming(uint64(l))
		}
	}
	l = len(m.BiddingCoinDenom)
	if l > 0 {
		n += 1 + l + sovLiquidfarming(uint64(l))
	}
	l = len(m.SellingReserveAddress)
	if l > 0 {
		n += 1 + l + sovLiquidfarming(uint64(l))
	}
	l = len(m.PayingReserveAddress)
	if l > 0 {
		n += 1 + l + sovLiquidfarming(uint64(l))
	}
	if m.StartTime != nil {
		l = github_com_gogo_protobuf_types.SizeOfStdTime(*m.StartTime)
		n += 1 + l + sovLiquidfarming(uint64(l))
	}
	if m.EndTime != nil {
		l = github_com_gogo_protobuf_types.SizeOfStdTime(*m.EndTime)
		n += 1 + l + sovLiquidfarming(uint64(l))
	}
	if m.Status != 0 {
		n += 1 + sovLiquidfarming(uint64(m.Status))
	}
	if m.WinnerBidId != 0 {
		n += 1 + sovLiquidfarming(uint64(m.WinnerBidId))
	}
	return n
}

func (m *Bid) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	if m.Id != 0 {
		n += 1 + sovLiquidfarming(uint64(m.Id))
	}
	if m.AuctionId != 0 {
		n += 1 + sovLiquidfarming(uint64(m.AuctionId))
	}
	l = len(m.Bidder)
	if l > 0 {
		n += 1 + l + sovLiquidfarming(uint64(l))
	}
	l = m.Amount.Size()
	n += 1 + l + sovLiquidfarming(uint64(l))
	if m.IsMatched {
		n += 2
	}
	return n
}

func sovLiquidfarming(x uint64) (n int) {
	return (math_bits.Len64(x|1) + 6) / 7
}
func sozLiquidfarming(x uint64) (n int) {
	return sovLiquidfarming(uint64((x << 1) ^ uint64((int64(x) >> 63))))
}
func (m *QueuedFarming) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowLiquidfarming
			}
			if iNdEx >= l {
				return io.ErrUnexpectedEOF
			}
			b := dAtA[iNdEx]
			iNdEx++
			wire |= uint64(b&0x7F) << shift
			if b < 0x80 {
				break
			}
		}
		fieldNum := int32(wire >> 3)
		wireType := int(wire & 0x7)
		if wireType == 4 {
			return fmt.Errorf("proto: QueuedFarming: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: QueuedFarming: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field PoolId", wireType)
			}
			m.PoolId = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowLiquidfarming
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.PoolId |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 2:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field Id", wireType)
			}
			m.Id = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowLiquidfarming
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.Id |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 3:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Farmer", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowLiquidfarming
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				stringLen |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			intStringLen := int(stringLen)
			if intStringLen < 0 {
				return ErrInvalidLengthLiquidfarming
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthLiquidfarming
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Farmer = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 4:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field FarmingCoin", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowLiquidfarming
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				msglen |= int(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			if msglen < 0 {
				return ErrInvalidLengthLiquidfarming
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthLiquidfarming
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if err := m.FarmingCoin.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipLiquidfarming(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthLiquidfarming
			}
			if (iNdEx + skippy) > l {
				return io.ErrUnexpectedEOF
			}
			iNdEx += skippy
		}
	}

	if iNdEx > l {
		return io.ErrUnexpectedEOF
	}
	return nil
}
func (m *RewardsAuction) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowLiquidfarming
			}
			if iNdEx >= l {
				return io.ErrUnexpectedEOF
			}
			b := dAtA[iNdEx]
			iNdEx++
			wire |= uint64(b&0x7F) << shift
			if b < 0x80 {
				break
			}
		}
		fieldNum := int32(wire >> 3)
		wireType := int(wire & 0x7)
		if wireType == 4 {
			return fmt.Errorf("proto: RewardsAuction: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: RewardsAuction: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field Id", wireType)
			}
			m.Id = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowLiquidfarming
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.Id |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 2:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field PoolId", wireType)
			}
			m.PoolId = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowLiquidfarming
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.PoolId |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 3:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field SellingRewards", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowLiquidfarming
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				msglen |= int(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			if msglen < 0 {
				return ErrInvalidLengthLiquidfarming
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthLiquidfarming
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.SellingRewards = append(m.SellingRewards, types.Coin{})
			if err := m.SellingRewards[len(m.SellingRewards)-1].Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 4:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field BiddingCoinDenom", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowLiquidfarming
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				stringLen |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			intStringLen := int(stringLen)
			if intStringLen < 0 {
				return ErrInvalidLengthLiquidfarming
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthLiquidfarming
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.BiddingCoinDenom = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 5:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field SellingReserveAddress", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowLiquidfarming
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				stringLen |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			intStringLen := int(stringLen)
			if intStringLen < 0 {
				return ErrInvalidLengthLiquidfarming
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthLiquidfarming
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.SellingReserveAddress = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 6:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field PayingReserveAddress", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowLiquidfarming
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				stringLen |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			intStringLen := int(stringLen)
			if intStringLen < 0 {
				return ErrInvalidLengthLiquidfarming
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthLiquidfarming
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.PayingReserveAddress = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 7:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field StartTime", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowLiquidfarming
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				msglen |= int(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			if msglen < 0 {
				return ErrInvalidLengthLiquidfarming
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthLiquidfarming
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if m.StartTime == nil {
				m.StartTime = new(time.Time)
			}
			if err := github_com_gogo_protobuf_types.StdTimeUnmarshal(m.StartTime, dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 8:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field EndTime", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowLiquidfarming
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				msglen |= int(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			if msglen < 0 {
				return ErrInvalidLengthLiquidfarming
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthLiquidfarming
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if m.EndTime == nil {
				m.EndTime = new(time.Time)
			}
			if err := github_com_gogo_protobuf_types.StdTimeUnmarshal(m.EndTime, dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 9:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field Status", wireType)
			}
			m.Status = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowLiquidfarming
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.Status |= AuctionStatus(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 10:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field WinnerBidId", wireType)
			}
			m.WinnerBidId = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowLiquidfarming
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.WinnerBidId |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		default:
			iNdEx = preIndex
			skippy, err := skipLiquidfarming(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthLiquidfarming
			}
			if (iNdEx + skippy) > l {
				return io.ErrUnexpectedEOF
			}
			iNdEx += skippy
		}
	}

	if iNdEx > l {
		return io.ErrUnexpectedEOF
	}
	return nil
}
func (m *Bid) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowLiquidfarming
			}
			if iNdEx >= l {
				return io.ErrUnexpectedEOF
			}
			b := dAtA[iNdEx]
			iNdEx++
			wire |= uint64(b&0x7F) << shift
			if b < 0x80 {
				break
			}
		}
		fieldNum := int32(wire >> 3)
		wireType := int(wire & 0x7)
		if wireType == 4 {
			return fmt.Errorf("proto: Bid: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: Bid: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field Id", wireType)
			}
			m.Id = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowLiquidfarming
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.Id |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 2:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field AuctionId", wireType)
			}
			m.AuctionId = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowLiquidfarming
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.AuctionId |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 3:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Bidder", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowLiquidfarming
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				stringLen |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			intStringLen := int(stringLen)
			if intStringLen < 0 {
				return ErrInvalidLengthLiquidfarming
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthLiquidfarming
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Bidder = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 4:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Amount", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowLiquidfarming
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				msglen |= int(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			if msglen < 0 {
				return ErrInvalidLengthLiquidfarming
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthLiquidfarming
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if err := m.Amount.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 5:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field IsMatched", wireType)
			}
			var v int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowLiquidfarming
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				v |= int(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			m.IsMatched = bool(v != 0)
		default:
			iNdEx = preIndex
			skippy, err := skipLiquidfarming(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthLiquidfarming
			}
			if (iNdEx + skippy) > l {
				return io.ErrUnexpectedEOF
			}
			iNdEx += skippy
		}
	}

	if iNdEx > l {
		return io.ErrUnexpectedEOF
	}
	return nil
}
func skipLiquidfarming(dAtA []byte) (n int, err error) {
	l := len(dAtA)
	iNdEx := 0
	depth := 0
	for iNdEx < l {
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return 0, ErrIntOverflowLiquidfarming
			}
			if iNdEx >= l {
				return 0, io.ErrUnexpectedEOF
			}
			b := dAtA[iNdEx]
			iNdEx++
			wire |= (uint64(b) & 0x7F) << shift
			if b < 0x80 {
				break
			}
		}
		wireType := int(wire & 0x7)
		switch wireType {
		case 0:
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return 0, ErrIntOverflowLiquidfarming
				}
				if iNdEx >= l {
					return 0, io.ErrUnexpectedEOF
				}
				iNdEx++
				if dAtA[iNdEx-1] < 0x80 {
					break
				}
			}
		case 1:
			iNdEx += 8
		case 2:
			var length int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return 0, ErrIntOverflowLiquidfarming
				}
				if iNdEx >= l {
					return 0, io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				length |= (int(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			if length < 0 {
				return 0, ErrInvalidLengthLiquidfarming
			}
			iNdEx += length
		case 3:
			depth++
		case 4:
			if depth == 0 {
				return 0, ErrUnexpectedEndOfGroupLiquidfarming
			}
			depth--
		case 5:
			iNdEx += 4
		default:
			return 0, fmt.Errorf("proto: illegal wireType %d", wireType)
		}
		if iNdEx < 0 {
			return 0, ErrInvalidLengthLiquidfarming
		}
		if depth == 0 {
			return iNdEx, nil
		}
	}
	return 0, io.ErrUnexpectedEOF
}

var (
	ErrInvalidLengthLiquidfarming        = fmt.Errorf("proto: negative length found during unmarshaling")
	ErrIntOverflowLiquidfarming          = fmt.Errorf("proto: integer overflow")
	ErrUnexpectedEndOfGroupLiquidfarming = fmt.Errorf("proto: unexpected end of group")
)
