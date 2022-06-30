// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: squad/auction/v1beta1/tx.proto

package types

import (
	context "context"
	fmt "fmt"
	types "github.com/cosmos/cosmos-sdk/codec/types"
	_ "github.com/cosmos/cosmos-sdk/types"
	_ "github.com/gogo/protobuf/gogoproto"
	grpc1 "github.com/gogo/protobuf/grpc"
	proto "github.com/gogo/protobuf/proto"
	_ "github.com/regen-network/cosmos-proto"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
	_ "google.golang.org/protobuf/types/known/timestamppb"
	io "io"
	math "math"
	math_bits "math/bits"
)

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.GoGoProtoPackageIsVersion3 // please upgrade the proto package

type MsgCreateAuction struct {
	Custom     *types.Any `protobuf:"bytes,1,opt,name=custom,proto3" json:"custom,omitempty"`
	Auctioneer string     `protobuf:"bytes,2,opt,name=auctioneer,proto3" json:"auctioneer,omitempty"`
}

func (m *MsgCreateAuction) Reset()      { *m = MsgCreateAuction{} }
func (*MsgCreateAuction) ProtoMessage() {}
func (*MsgCreateAuction) Descriptor() ([]byte, []int) {
	return fileDescriptor_5715cbe414447688, []int{0}
}
func (m *MsgCreateAuction) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *MsgCreateAuction) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_MsgCreateAuction.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *MsgCreateAuction) XXX_Merge(src proto.Message) {
	xxx_messageInfo_MsgCreateAuction.Merge(m, src)
}
func (m *MsgCreateAuction) XXX_Size() int {
	return m.Size()
}
func (m *MsgCreateAuction) XXX_DiscardUnknown() {
	xxx_messageInfo_MsgCreateAuction.DiscardUnknown(m)
}

var xxx_messageInfo_MsgCreateAuction proto.InternalMessageInfo

type MsgCreateAuctionResponse struct {
}

func (m *MsgCreateAuctionResponse) Reset()         { *m = MsgCreateAuctionResponse{} }
func (m *MsgCreateAuctionResponse) String() string { return proto.CompactTextString(m) }
func (*MsgCreateAuctionResponse) ProtoMessage()    {}
func (*MsgCreateAuctionResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_5715cbe414447688, []int{1}
}
func (m *MsgCreateAuctionResponse) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *MsgCreateAuctionResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_MsgCreateAuctionResponse.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *MsgCreateAuctionResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_MsgCreateAuctionResponse.Merge(m, src)
}
func (m *MsgCreateAuctionResponse) XXX_Size() int {
	return m.Size()
}
func (m *MsgCreateAuctionResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_MsgCreateAuctionResponse.DiscardUnknown(m)
}

var xxx_messageInfo_MsgCreateAuctionResponse proto.InternalMessageInfo

func init() {
	proto.RegisterType((*MsgCreateAuction)(nil), "squad.auction.v1beta1.MsgCreateAuction")
	proto.RegisterType((*MsgCreateAuctionResponse)(nil), "squad.auction.v1beta1.MsgCreateAuctionResponse")
}

func init() { proto.RegisterFile("squad/auction/v1beta1/tx.proto", fileDescriptor_5715cbe414447688) }

var fileDescriptor_5715cbe414447688 = []byte{
	// 347 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x8c, 0x91, 0x31, 0x4f, 0xf2, 0x40,
	0x18, 0xc7, 0xaf, 0xef, 0x9b, 0x90, 0x78, 0xc6, 0x84, 0x34, 0x98, 0x60, 0x87, 0x2b, 0xc1, 0x41,
	0x16, 0xee, 0x04, 0x17, 0xe3, 0x06, 0x6c, 0x26, 0x2c, 0x8c, 0x2e, 0xe6, 0x5a, 0xcf, 0xb3, 0x09,
	0xed, 0x53, 0xb9, 0xab, 0x81, 0xc5, 0x38, 0x3a, 0x3a, 0x3a, 0xf2, 0x21, 0xfc, 0x10, 0xc6, 0x89,
	0xd1, 0xc1, 0xc1, 0xc0, 0xe2, 0xc7, 0x30, 0xf4, 0xae, 0x44, 0x1b, 0x06, 0xb7, 0x3e, 0xfd, 0xfd,
	0xfa, 0xfc, 0xaf, 0xff, 0xc3, 0x44, 0xdd, 0x66, 0xfc, 0x8a, 0xf1, 0x2c, 0xd4, 0x11, 0x24, 0xec,
	0xae, 0x13, 0x08, 0xcd, 0x3b, 0x4c, 0x4f, 0x69, 0x3a, 0x01, 0x0d, 0xee, 0x7e, 0xce, 0xa9, 0xe5,
	0xd4, 0x72, 0xaf, 0x26, 0x41, 0x42, 0x6e, 0xb0, 0xf5, 0x93, 0x91, 0xbd, 0x83, 0x10, 0x54, 0x0c,
	0xea, 0xd2, 0x00, 0x33, 0x58, 0x44, 0xcc, 0xc4, 0x02, 0xae, 0xc4, 0x26, 0x25, 0x84, 0x28, 0x29,
	0x3e, 0x95, 0x00, 0x72, 0x2c, 0x58, 0x3e, 0x05, 0xd9, 0x35, 0xe3, 0xc9, 0xcc, 0x22, 0xbf, 0x8c,
	0x74, 0x14, 0x0b, 0xa5, 0x79, 0x9c, 0x5a, 0xe1, 0x70, 0xfb, 0x3f, 0x14, 0x67, 0xce, 0xa5, 0xe6,
	0x3d, 0xae, 0x0e, 0x95, 0x1c, 0x4c, 0x04, 0xd7, 0xa2, 0x67, 0x88, 0x7b, 0x8a, 0x2b, 0x61, 0xa6,
	0x34, 0xc4, 0x75, 0xa7, 0xe1, 0xb4, 0x76, 0xbb, 0x35, 0x6a, 0xa2, 0x68, 0x11, 0x45, 0x7b, 0xc9,
	0xac, 0x8f, 0xdf, 0x5e, 0xda, 0x95, 0x41, 0xee, 0x8d, 0xac, 0xef, 0x12, 0x8c, 0xed, 0x7a, 0x21,
	0x26, 0xf5, 0x7f, 0x0d, 0xa7, 0xb5, 0x33, 0xfa, 0xf1, 0xe6, 0xac, 0xfa, 0x38, 0xf7, 0xd1, 0xf3,
	0xdc, 0x47, 0x5f, 0x73, 0x1f, 0x3d, 0x7c, 0x34, 0x50, 0xd3, 0xc3, 0xf5, 0x72, 0xfe, 0x48, 0xa8,
	0x14, 0x12, 0x25, 0xba, 0x29, 0xfe, 0x3f, 0x54, 0xd2, 0x8d, 0xf0, 0xde, 0xef, 0xf3, 0x1d, 0xd1,
	0xad, 0xed, 0xd3, 0xf2, 0x22, 0x8f, 0xfd, 0x51, 0x2c, 0x12, 0xfb, 0xe7, 0xaf, 0x4b, 0xe2, 0x2c,
	0x96, 0xc4, 0xf9, 0x5c, 0x12, 0xe7, 0x69, 0x45, 0xd0, 0x62, 0x45, 0xd0, 0xfb, 0x8a, 0xa0, 0x8b,
	0x63, 0x19, 0xe9, 0x9b, 0x2c, 0xa0, 0x21, 0xc4, 0xf6, 0x06, 0xd7, 0x9b, 0xdb, 0x63, 0x1e, 0x28,
	0x66, 0x7a, 0x9e, 0x6e, 0x9a, 0xd6, 0xb3, 0x54, 0xa8, 0xa0, 0x92, 0xb7, 0x75, 0xf2, 0x1d, 0x00,
	0x00, 0xff, 0xff, 0x4f, 0x02, 0x73, 0x1a, 0x4b, 0x02, 0x00, 0x00,
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion4

// MsgClient is the client API for Msg service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type MsgClient interface {
	CreateAuction(ctx context.Context, in *MsgCreateAuction, opts ...grpc.CallOption) (*MsgCreateAuctionResponse, error)
}

type msgClient struct {
	cc grpc1.ClientConn
}

func NewMsgClient(cc grpc1.ClientConn) MsgClient {
	return &msgClient{cc}
}

func (c *msgClient) CreateAuction(ctx context.Context, in *MsgCreateAuction, opts ...grpc.CallOption) (*MsgCreateAuctionResponse, error) {
	out := new(MsgCreateAuctionResponse)
	err := c.cc.Invoke(ctx, "/squad.auction.v1beta1.Msg/CreateAuction", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// MsgServer is the server API for Msg service.
type MsgServer interface {
	CreateAuction(context.Context, *MsgCreateAuction) (*MsgCreateAuctionResponse, error)
}

// UnimplementedMsgServer can be embedded to have forward compatible implementations.
type UnimplementedMsgServer struct {
}

func (*UnimplementedMsgServer) CreateAuction(ctx context.Context, req *MsgCreateAuction) (*MsgCreateAuctionResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method CreateAuction not implemented")
}

func RegisterMsgServer(s grpc1.Server, srv MsgServer) {
	s.RegisterService(&_Msg_serviceDesc, srv)
}

func _Msg_CreateAuction_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(MsgCreateAuction)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(MsgServer).CreateAuction(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/squad.auction.v1beta1.Msg/CreateAuction",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(MsgServer).CreateAuction(ctx, req.(*MsgCreateAuction))
	}
	return interceptor(ctx, in, info, handler)
}

var _Msg_serviceDesc = grpc.ServiceDesc{
	ServiceName: "squad.auction.v1beta1.Msg",
	HandlerType: (*MsgServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "CreateAuction",
			Handler:    _Msg_CreateAuction_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "squad/auction/v1beta1/tx.proto",
}

func (m *MsgCreateAuction) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *MsgCreateAuction) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *MsgCreateAuction) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if len(m.Auctioneer) > 0 {
		i -= len(m.Auctioneer)
		copy(dAtA[i:], m.Auctioneer)
		i = encodeVarintTx(dAtA, i, uint64(len(m.Auctioneer)))
		i--
		dAtA[i] = 0x12
	}
	if m.Custom != nil {
		{
			size, err := m.Custom.MarshalToSizedBuffer(dAtA[:i])
			if err != nil {
				return 0, err
			}
			i -= size
			i = encodeVarintTx(dAtA, i, uint64(size))
		}
		i--
		dAtA[i] = 0xa
	}
	return len(dAtA) - i, nil
}

func (m *MsgCreateAuctionResponse) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *MsgCreateAuctionResponse) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *MsgCreateAuctionResponse) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	return len(dAtA) - i, nil
}

func encodeVarintTx(dAtA []byte, offset int, v uint64) int {
	offset -= sovTx(v)
	base := offset
	for v >= 1<<7 {
		dAtA[offset] = uint8(v&0x7f | 0x80)
		v >>= 7
		offset++
	}
	dAtA[offset] = uint8(v)
	return base
}
func (m *MsgCreateAuction) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	if m.Custom != nil {
		l = m.Custom.Size()
		n += 1 + l + sovTx(uint64(l))
	}
	l = len(m.Auctioneer)
	if l > 0 {
		n += 1 + l + sovTx(uint64(l))
	}
	return n
}

func (m *MsgCreateAuctionResponse) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	return n
}

func sovTx(x uint64) (n int) {
	return (math_bits.Len64(x|1) + 6) / 7
}
func sozTx(x uint64) (n int) {
	return sovTx(uint64((x << 1) ^ uint64((int64(x) >> 63))))
}
func (m *MsgCreateAuction) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowTx
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
			return fmt.Errorf("proto: MsgCreateAuction: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: MsgCreateAuction: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Custom", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowTx
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
				return ErrInvalidLengthTx
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthTx
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if m.Custom == nil {
				m.Custom = &types.Any{}
			}
			if err := m.Custom.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Auctioneer", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowTx
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
				return ErrInvalidLengthTx
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthTx
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Auctioneer = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipTx(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthTx
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
func (m *MsgCreateAuctionResponse) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowTx
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
			return fmt.Errorf("proto: MsgCreateAuctionResponse: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: MsgCreateAuctionResponse: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		default:
			iNdEx = preIndex
			skippy, err := skipTx(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthTx
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
func skipTx(dAtA []byte) (n int, err error) {
	l := len(dAtA)
	iNdEx := 0
	depth := 0
	for iNdEx < l {
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return 0, ErrIntOverflowTx
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
					return 0, ErrIntOverflowTx
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
					return 0, ErrIntOverflowTx
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
				return 0, ErrInvalidLengthTx
			}
			iNdEx += length
		case 3:
			depth++
		case 4:
			if depth == 0 {
				return 0, ErrUnexpectedEndOfGroupTx
			}
			depth--
		case 5:
			iNdEx += 4
		default:
			return 0, fmt.Errorf("proto: illegal wireType %d", wireType)
		}
		if iNdEx < 0 {
			return 0, ErrInvalidLengthTx
		}
		if depth == 0 {
			return iNdEx, nil
		}
	}
	return 0, io.ErrUnexpectedEOF
}

var (
	ErrInvalidLengthTx        = fmt.Errorf("proto: negative length found during unmarshaling")
	ErrIntOverflowTx          = fmt.Errorf("proto: integer overflow")
	ErrUnexpectedEndOfGroupTx = fmt.Errorf("proto: unexpected end of group")
)
