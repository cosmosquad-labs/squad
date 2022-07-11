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
	// depositor specifies the bech32-encoded address that makes a deposit to the pool
	Depositor string `protobuf:"bytes,3,opt,name=depositor,proto3" json:"depositor,omitempty"`
	// deposit_coin specifies the amount of pool coin to deposit
	DepositCoin types.Coin `protobuf:"bytes,4,opt,name=deposit_coin,json=depositCoin,proto3,castrepeated=github.com/cosmos/cosmos-sdk/types.Coin" json:"deposit_coin"`
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
	// 772 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xb4, 0x94, 0xcf, 0x6e, 0xe3, 0x44,
	0x1c, 0xc7, 0xed, 0x24, 0xa4, 0xcd, 0x94, 0x66, 0x83, 0x95, 0xdd, 0x75, 0xbd, 0xac, 0x63, 0xe5,
	0x42, 0xb4, 0x62, 0x6d, 0x76, 0xa9, 0x7a, 0xe8, 0x2d, 0x69, 0x12, 0xe1, 0x03, 0x01, 0x9c, 0x54,
	0x42, 0x5c, 0x2c, 0x3b, 0x33, 0x4d, 0x47, 0xd8, 0x9e, 0xd4, 0x63, 0xb7, 0xf4, 0x0d, 0xaa, 0x9e,
	0xfa, 0x02, 0x95, 0x90, 0xb8, 0xf1, 0x06, 0x3c, 0x00, 0x52, 0x8e, 0xbd, 0x20, 0x71, 0x6a, 0x69,
	0xfb, 0x06, 0x3c, 0x01, 0x9a, 0x3f, 0x69, 0x48, 0x8a, 0x0a, 0x17, 0x4e, 0xc9, 0xef, 0xf7, 0xf3,
	0xe7, 0x3b, 0xdf, 0x19, 0x7f, 0xc7, 0xc0, 0xa1, 0x47, 0x79, 0x00, 0x9d, 0x08, 0x1f, 0xe5, 0x18,
	0x1e, 0x04, 0x69, 0x8c, 0x93, 0x89, 0x73, 0xfc, 0x2e, 0x44, 0x59, 0xf0, 0x6e, 0xb9, 0x6b, 0x4f,
	0x53, 0x92, 0x11, 0xed, 0x15, 0x07, 0xec, 0xe5, 0x91, 0x04, 0x8c, 0xfa, 0x84, 0x4c, 0x08, 0x7f,
	0xce, 0x61, 0xff, 0x04, 0x62, 0x6c, 0x8d, 0x09, 0x8d, 0x09, 0xf5, 0xc5, 0x40, 0x14, 0x72, 0x64,
	0x8a, 0xca, 0x09, 0x03, 0x8a, 0x1e, 0x96, 0x1d, 0x13, 0x9c, 0xc8, 0x79, 0x63, 0x42, 0xc8, 0x24,
	0x42, 0x0e, 0xaf, 0xc2, 0xfc, 0xc0, 0xc9, 0x70, 0x8c, 0x68, 0x16, 0xc4, 0x53, 0xf1, 0x40, 0xf3,
	0x57, 0x15, 0x6c, 0x7e, 0x93, 0xa3, 0x1c, 0xc1, 0xbe, 0xf0, 0xa2, 0xbd, 0x04, 0x6b, 0x53, 0x42,
	0x22, 0x1f, 0x43, 0x5d, 0xb5, 0xd4, 0x56, 0xc9, 0x2b, 0xb3, 0xd2, 0x85, 0x5a, 0x15, 0x14, 0x30,
	0xd4, 0x0b, 0xbc, 0x57, 0xc0, 0x50, 0xfb, 0x18, 0x54, 0x20, 0x9a, 0x12, 0x8a, 0x33, 0x92, 0xea,
	0x45, 0x4b, 0x6d, 0x55, 0xbc, 0x45, 0x43, 0x8b, 0xc1, 0x87, 0xb2, 0xf0, 0x99, 0x1f, 0xbd, 0x64,
	0xa9, 0xad, 0x8d, 0xf7, 0x5b, 0xb6, 0xb4, 0xcf, 0x0c, 0xcf, 0xb7, 0x6d, 0xef, 0x11, 0x9c, 0x74,
	0x9c, 0xd9, 0x75, 0x43, 0xf9, 0xf9, 0xa6, 0xf1, 0xc9, 0x04, 0x67, 0x87, 0x79, 0x68, 0x8f, 0x49,
	0x2c, 0xf7, 0x2a, 0x7f, 0xde, 0x52, 0xf8, 0xbd, 0x93, 0x9d, 0x4e, 0x11, 0xe5, 0x80, 0xb7, 0x21,
	0xf5, 0x59, 0xd1, 0xbc, 0x2d, 0x81, 0xaa, 0x87, 0x4e, 0x82, 0x14, 0xd2, 0x76, 0x3e, 0xce, 0x30,
	0x49, 0xa4, 0x5f, 0xf5, 0xc1, 0xef, 0xdf, 0x36, 0x56, 0x58, 0xda, 0x58, 0x06, 0x9e, 0x51, 0x14,
	0x45, 0x38, 0x99, 0xf8, 0xa9, 0x90, 0xd0, 0x8b, 0x56, 0xf1, 0x69, 0xb7, 0x9f, 0x49, 0xb7, 0xad,
	0xff, 0xe8, 0x96, 0x7a, 0x55, 0xb9, 0x86, 0x74, 0xa9, 0x7d, 0x0a, 0xb4, 0x10, 0x43, 0xc8, 0x56,
	0x65, 0x07, 0xe4, 0x43, 0x94, 0x90, 0x98, 0x1f, 0x53, 0xc5, 0xab, 0xc9, 0x09, 0x23, 0xbb, 0xac,
	0xaf, 0xed, 0x80, 0x97, 0x0b, 0x8f, 0x14, 0xa5, 0xc7, 0xc8, 0x0f, 0x20, 0x4c, 0x11, 0xa5, 0xfa,
	0x07, 0x1c, 0x79, 0xfe, 0x20, 0xcf, 0xa7, 0x6d, 0x31, 0xd4, 0xb6, 0xc1, 0x8b, 0x69, 0x70, 0xfa,
	0x4f, 0x58, 0x99, 0x63, 0x75, 0x31, 0x5d, 0xa1, 0xbe, 0x05, 0x80, 0x66, 0x41, 0x9a, 0xf9, 0x2c,
	0x2e, 0xfa, 0x1a, 0x7f, 0x75, 0x86, 0x2d, 0xb2, 0x64, 0xcf, 0xb3, 0x64, 0x8f, 0xe6, 0x59, 0xea,
	0xbc, 0x9e, 0x5d, 0x37, 0xd4, 0x3f, 0xaf, 0x1b, 0x1f, 0x9d, 0x06, 0x71, 0xb4, 0xdb, 0x5c, 0xb0,
	0xcd, 0x8b, 0x9b, 0x86, 0xea, 0x55, 0x78, 0x83, 0x3d, 0xae, 0x79, 0x60, 0x1d, 0x25, 0x50, 0xe8,
	0xae, 0xff, 0xab, 0xee, 0x2b, 0xa9, 0xfb, 0x4c, 0xe8, 0xce, 0x49, 0xa1, 0xba, 0x86, 0x12, 0xc8,
	0x35, 0x3b, 0xa0, 0x4c, 0xb3, 0x20, 0xcb, 0xa9, 0x5e, 0xb1, 0xd4, 0x56, 0xf5, 0xfd, 0x1b, 0xfb,
	0x89, 0x3b, 0x66, 0xcb, 0x78, 0x0c, 0x39, 0xe1, 0x49, 0x52, 0x6b, 0x82, 0xcd, 0x13, 0x9c, 0x24,
	0x28, 0xf5, 0x43, 0x0c, 0x59, 0x44, 0x00, 0x8f, 0xc8, 0x86, 0x68, 0x76, 0x30, 0x74, 0x61, 0xf3,
	0x37, 0x15, 0x14, 0x3b, 0x18, 0x3e, 0x0a, 0xd6, 0x6b, 0x00, 0x02, 0x21, 0xba, 0xc8, 0x56, 0x45,
	0x76, 0x5c, 0xa8, 0xbd, 0x00, 0x65, 0xf6, 0x3a, 0xd1, 0xfc, 0x92, 0xc8, 0x4a, 0x0b, 0x41, 0x39,
	0x88, 0x49, 0x9e, 0x64, 0xff, 0xc3, 0xdd, 0x90, 0xca, 0xcc, 0x1a, 0xa6, 0x7e, 0x1c, 0x64, 0xe3,
	0x43, 0x04, 0x79, 0x52, 0xd6, 0xbd, 0x0a, 0xa6, 0x5f, 0x8a, 0xc6, 0x6e, 0xe9, 0xec, 0xc7, 0x86,
	0xf2, 0xe6, 0x17, 0x15, 0x6c, 0x2e, 0x9d, 0x8a, 0xb6, 0x0d, 0x8c, 0xf6, 0xfe, 0xde, 0xc8, 0xfd,
	0x6a, 0xe0, 0x0f, 0x47, 0xed, 0xd1, 0xfe, 0xd0, 0xdf, 0x1f, 0x0c, 0xbf, 0xee, 0xed, 0xb9, 0x7d,
	0xb7, 0xd7, 0xad, 0x29, 0x46, 0xfd, 0xfc, 0xd2, 0xaa, 0x2d, 0x21, 0x03, 0x1c, 0xb1, 0xac, 0xad,
	0x50, 0xc3, 0x51, 0xdb, 0x1b, 0xf5, 0xba, 0x35, 0xd5, 0xd0, 0xcf, 0x2f, 0xad, 0xfa, 0x12, 0x31,
	0x64, 0x99, 0x40, 0x90, 0x25, 0x7b, 0x85, 0xea, 0xbb, 0x03, 0x77, 0xf8, 0x45, 0xaf, 0x5b, 0x2b,
	0x18, 0x5b, 0xe7, 0x97, 0xd6, 0xf3, 0x25, 0xac, 0x8f, 0x13, 0x4c, 0x0f, 0x11, 0x34, 0x4a, 0x67,
	0x3f, 0x99, 0x4a, 0x67, 0x34, 0xbb, 0x35, 0x95, 0xd9, 0x9d, 0xa9, 0x5e, 0xdd, 0x99, 0xea, 0x1f,
	0x77, 0xa6, 0x7a, 0x71, 0x6f, 0x2a, 0x57, 0xf7, 0xa6, 0xf2, 0xfb, 0xbd, 0xa9, 0x7c, 0xb7, 0xf3,
	0xe8, 0xbc, 0x58, 0x30, 0xde, 0x46, 0x41, 0x48, 0xe5, 0x87, 0xfb, 0x87, 0x95, 0x4f, 0x37, 0x3f,
	0xc3, 0xb0, 0xcc, 0xb3, 0xf8, 0xf9, 0x5f, 0x01, 0x00, 0x00, 0xff, 0xff, 0xae, 0xe0, 0x47, 0x70,
	0xde, 0x05, 0x00, 0x00,
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
		size, err := m.DepositCoin.MarshalToSizedBuffer(dAtA[:i])
		if err != nil {
			return 0, err
		}
		i -= size
		i = encodeVarintLiquidfarming(dAtA, i, uint64(size))
	}
	i--
	dAtA[i] = 0x22
	if len(m.Depositor) > 0 {
		i -= len(m.Depositor)
		copy(dAtA[i:], m.Depositor)
		i = encodeVarintLiquidfarming(dAtA, i, uint64(len(m.Depositor)))
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
	l = len(m.Depositor)
	if l > 0 {
		n += 1 + l + sovLiquidfarming(uint64(l))
	}
	l = m.DepositCoin.Size()
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
				return fmt.Errorf("proto: wrong wireType = %d for field Depositor", wireType)
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
			m.Depositor = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 4:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field DepositCoin", wireType)
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
			if err := m.DepositCoin.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
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
