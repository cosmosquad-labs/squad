// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: squad/liquidfarming/v1beta1/params.proto

package types

import (
	fmt "fmt"
	github_com_cosmos_cosmos_sdk_types "github.com/cosmos/cosmos-sdk/types"
	types "github.com/cosmos/cosmos-sdk/types"
	_ "github.com/gogo/protobuf/gogoproto"
	proto "github.com/gogo/protobuf/proto"
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

// Params defines the parameters for the module.
type Params struct {
	LiquidFarmCreationFee github_com_cosmos_cosmos_sdk_types.Coins `protobuf:"bytes,1,rep,name=liquid_farm_creation_fee,json=liquidFarmCreationFee,proto3,castrepeated=github.com/cosmos/cosmos-sdk/types.Coins" json:"liquid_farm_creation_fee"`
	// delayed_farm_gas_fee is used to impose gas fee for the farm request
	DelayedFarmGasFee github_com_cosmos_cosmos_sdk_types.Gas `protobuf:"varint,2,opt,name=delayed_farm_gas_fee,json=delayedFarmGasFee,proto3,customtype=github.com/cosmos/cosmos-sdk/types.Gas" json:"delayed_farm_gas_fee" yaml:"delayed_farm_gas_fee"`
	LiquidFarms       []LiquidFarm                           `protobuf:"bytes,3,rep,name=liquid_farms,json=liquidFarms,proto3" json:"liquid_farms"`
}

func (m *Params) Reset()         { *m = Params{} }
func (m *Params) String() string { return proto.CompactTextString(m) }
func (*Params) ProtoMessage()    {}
func (*Params) Descriptor() ([]byte, []int) {
	return fileDescriptor_6012e16b27fcc811, []int{0}
}
func (m *Params) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *Params) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_Params.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *Params) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Params.Merge(m, src)
}
func (m *Params) XXX_Size() int {
	return m.Size()
}
func (m *Params) XXX_DiscardUnknown() {
	xxx_messageInfo_Params.DiscardUnknown(m)
}

var xxx_messageInfo_Params proto.InternalMessageInfo

// LiquidFarm defines ...
type LiquidFarm struct {
	PoolId            uint64                                 `protobuf:"varint,1,opt,name=pool_id,json=poolId,proto3" json:"pool_id,omitempty"`
	MinimumFarmAmount github_com_cosmos_cosmos_sdk_types.Int `protobuf:"bytes,2,opt,name=minimum_farm_amount,json=minimumFarmAmount,proto3,customtype=github.com/cosmos/cosmos-sdk/types.Int" json:"minimum_farm_amount"`
	MinimumBidAmount  github_com_cosmos_cosmos_sdk_types.Int `protobuf:"bytes,3,opt,name=minimum_bid_amount,json=minimumBidAmount,proto3,customtype=github.com/cosmos/cosmos-sdk/types.Int" json:"minimum_bid_amount"`
}

func (m *LiquidFarm) Reset()      { *m = LiquidFarm{} }
func (*LiquidFarm) ProtoMessage() {}
func (*LiquidFarm) Descriptor() ([]byte, []int) {
	return fileDescriptor_6012e16b27fcc811, []int{1}
}
func (m *LiquidFarm) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *LiquidFarm) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_LiquidFarm.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *LiquidFarm) XXX_Merge(src proto.Message) {
	xxx_messageInfo_LiquidFarm.Merge(m, src)
}
func (m *LiquidFarm) XXX_Size() int {
	return m.Size()
}
func (m *LiquidFarm) XXX_DiscardUnknown() {
	xxx_messageInfo_LiquidFarm.DiscardUnknown(m)
}

var xxx_messageInfo_LiquidFarm proto.InternalMessageInfo

func init() {
	proto.RegisterType((*Params)(nil), "squad.liquidfarming.v1beta1.Params")
	proto.RegisterType((*LiquidFarm)(nil), "squad.liquidfarming.v1beta1.LiquidFarm")
}

func init() {
	proto.RegisterFile("squad/liquidfarming/v1beta1/params.proto", fileDescriptor_6012e16b27fcc811)
}

var fileDescriptor_6012e16b27fcc811 = []byte{
	// 456 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x9c, 0x93, 0x3d, 0x6f, 0xd4, 0x30,
	0x1c, 0xc6, 0xe3, 0xde, 0xe9, 0x00, 0x97, 0x81, 0x86, 0x22, 0x42, 0x2b, 0x25, 0x55, 0x06, 0xc8,
	0x52, 0x9b, 0x82, 0xc4, 0x70, 0x1b, 0xa9, 0xd4, 0xaa, 0x12, 0x48, 0xa7, 0x8c, 0x08, 0x11, 0x39,
	0xb1, 0x1b, 0x2c, 0xe2, 0x38, 0x8d, 0x13, 0xc4, 0x2d, 0x4c, 0x0c, 0x8c, 0x8c, 0x8c, 0x9d, 0xf9,
	0x24, 0x1d, 0x3b, 0x22, 0x86, 0x03, 0xdd, 0x0d, 0x4c, 0x2c, 0x7c, 0x02, 0x64, 0x3b, 0xd7, 0xf2,
	0x26, 0x04, 0x4c, 0x79, 0xfb, 0xfb, 0xf7, 0x7b, 0x1e, 0x2b, 0x86, 0x91, 0x3a, 0xea, 0x08, 0xc5,
	0x25, 0x3f, 0xea, 0x38, 0x3d, 0x24, 0x8d, 0xe0, 0x55, 0x81, 0x9f, 0xef, 0x64, 0xac, 0x25, 0x3b,
	0xb8, 0x26, 0x0d, 0x11, 0x0a, 0xd5, 0x8d, 0x6c, 0xa5, 0xbb, 0x69, 0x26, 0xd1, 0x0f, 0x93, 0xa8,
	0x9f, 0xdc, 0x58, 0x2f, 0x64, 0x21, 0xcd, 0x1c, 0xd6, 0x77, 0x76, 0xc9, 0x86, 0x9f, 0x4b, 0x25,
	0xa4, 0xc2, 0x19, 0x51, 0xec, 0x0c, 0x9a, 0x4b, 0x5e, 0xd9, 0xef, 0xe1, 0xe7, 0x15, 0x38, 0x9a,
	0x18, 0x87, 0xfb, 0x0a, 0x40, 0xcf, 0xa2, 0x53, 0xcd, 0x4e, 0xf3, 0x86, 0x91, 0x96, 0xcb, 0x2a,
	0x3d, 0x64, 0xcc, 0x03, 0x5b, 0x83, 0x68, 0xf5, 0xce, 0x0d, 0x64, 0x71, 0x48, 0xe3, 0x96, 0x66,
	0xb4, 0x2b, 0x79, 0x15, 0xdf, 0x3e, 0x99, 0x05, 0xce, 0xbb, 0x8f, 0x41, 0x54, 0xf0, 0xf6, 0x69,
	0x97, 0xa1, 0x5c, 0x0a, 0xdc, 0xbb, 0xed, 0x65, 0x5b, 0xd1, 0x67, 0xb8, 0x9d, 0xd6, 0x4c, 0x99,
	0x05, 0x2a, 0xb9, 0x66, 0x65, 0x7b, 0xa4, 0x11, 0xbb, 0xbd, 0x6a, 0x8f, 0x31, 0xf7, 0x25, 0x5c,
	0xa7, 0xac, 0x24, 0x53, 0xd6, 0xc7, 0x28, 0x88, 0x32, 0x09, 0x56, 0xb6, 0x40, 0x34, 0x8c, 0x1f,
	0x6a, 0xcd, 0x87, 0x59, 0x70, 0xf3, 0x2f, 0x34, 0xfb, 0x44, 0x7d, 0x9d, 0x05, 0x9b, 0x53, 0x22,
	0xca, 0x71, 0xf8, 0x3b, 0x66, 0x98, 0xac, 0xf5, 0xaf, 0x75, 0x88, 0x7d, 0xa2, 0xb4, 0x7f, 0x02,
	0x2f, 0x7f, 0xb7, 0x0b, 0xca, 0x1b, 0x98, 0xe6, 0xb7, 0xd0, 0x1f, 0xf6, 0x1e, 0x3d, 0x38, 0x6b,
	0x12, 0x0f, 0x75, 0xc0, 0x64, 0xf5, 0xbc, 0x9b, 0x1a, 0x0f, 0x5f, 0x1f, 0x07, 0x4e, 0xf8, 0x05,
	0x40, 0x78, 0x3e, 0xe7, 0x5e, 0x87, 0x17, 0x6a, 0x29, 0xcb, 0x94, 0x53, 0x0f, 0xe8, 0x66, 0xc9,
	0x48, 0x3f, 0x1e, 0x50, 0xf7, 0x09, 0xbc, 0x2a, 0x78, 0xc5, 0x45, 0x27, 0x6c, 0x56, 0x22, 0x64,
	0x57, 0xb5, 0xa6, 0xfe, 0xa5, 0x18, 0xfd, 0x43, 0xfd, 0x83, 0xaa, 0x4d, 0xd6, 0x7a, 0x94, 0x56,
	0xde, 0x37, 0x20, 0xf7, 0x31, 0x74, 0x97, 0xfc, 0x8c, 0xd3, 0x25, 0x7e, 0xf0, 0x5f, 0xf8, 0x2b,
	0x3d, 0x29, 0xe6, 0xd4, 0xd2, 0xc7, 0x17, 0x75, 0xd7, 0xb7, 0xc7, 0x81, 0x13, 0x4f, 0x4e, 0xe6,
	0x3e, 0x38, 0x9d, 0xfb, 0xe0, 0xd3, 0xdc, 0x07, 0x6f, 0x16, 0xbe, 0x73, 0xba, 0xf0, 0x9d, 0xf7,
	0x0b, 0xdf, 0x79, 0x74, 0xef, 0x17, 0xba, 0xde, 0xda, 0xed, 0x92, 0x64, 0x0a, 0xdb, 0xb3, 0xf0,
	0xe2, 0xa7, 0xd3, 0x60, 0x8c, 0xd9, 0xc8, 0xfc, 0xb2, 0x77, 0xbf, 0x05, 0x00, 0x00, 0xff, 0xff,
	0x94, 0xf6, 0x87, 0xa2, 0x31, 0x03, 0x00, 0x00,
}

func (m *Params) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *Params) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *Params) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if len(m.LiquidFarms) > 0 {
		for iNdEx := len(m.LiquidFarms) - 1; iNdEx >= 0; iNdEx-- {
			{
				size, err := m.LiquidFarms[iNdEx].MarshalToSizedBuffer(dAtA[:i])
				if err != nil {
					return 0, err
				}
				i -= size
				i = encodeVarintParams(dAtA, i, uint64(size))
			}
			i--
			dAtA[i] = 0x1a
		}
	}
	if m.DelayedFarmGasFee != 0 {
		i = encodeVarintParams(dAtA, i, uint64(m.DelayedFarmGasFee))
		i--
		dAtA[i] = 0x10
	}
	if len(m.LiquidFarmCreationFee) > 0 {
		for iNdEx := len(m.LiquidFarmCreationFee) - 1; iNdEx >= 0; iNdEx-- {
			{
				size, err := m.LiquidFarmCreationFee[iNdEx].MarshalToSizedBuffer(dAtA[:i])
				if err != nil {
					return 0, err
				}
				i -= size
				i = encodeVarintParams(dAtA, i, uint64(size))
			}
			i--
			dAtA[i] = 0xa
		}
	}
	return len(dAtA) - i, nil
}

func (m *LiquidFarm) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *LiquidFarm) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *LiquidFarm) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	{
		size := m.MinimumBidAmount.Size()
		i -= size
		if _, err := m.MinimumBidAmount.MarshalTo(dAtA[i:]); err != nil {
			return 0, err
		}
		i = encodeVarintParams(dAtA, i, uint64(size))
	}
	i--
	dAtA[i] = 0x1a
	{
		size := m.MinimumFarmAmount.Size()
		i -= size
		if _, err := m.MinimumFarmAmount.MarshalTo(dAtA[i:]); err != nil {
			return 0, err
		}
		i = encodeVarintParams(dAtA, i, uint64(size))
	}
	i--
	dAtA[i] = 0x12
	if m.PoolId != 0 {
		i = encodeVarintParams(dAtA, i, uint64(m.PoolId))
		i--
		dAtA[i] = 0x8
	}
	return len(dAtA) - i, nil
}

func encodeVarintParams(dAtA []byte, offset int, v uint64) int {
	offset -= sovParams(v)
	base := offset
	for v >= 1<<7 {
		dAtA[offset] = uint8(v&0x7f | 0x80)
		v >>= 7
		offset++
	}
	dAtA[offset] = uint8(v)
	return base
}
func (m *Params) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	if len(m.LiquidFarmCreationFee) > 0 {
		for _, e := range m.LiquidFarmCreationFee {
			l = e.Size()
			n += 1 + l + sovParams(uint64(l))
		}
	}
	if m.DelayedFarmGasFee != 0 {
		n += 1 + sovParams(uint64(m.DelayedFarmGasFee))
	}
	if len(m.LiquidFarms) > 0 {
		for _, e := range m.LiquidFarms {
			l = e.Size()
			n += 1 + l + sovParams(uint64(l))
		}
	}
	return n
}

func (m *LiquidFarm) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	if m.PoolId != 0 {
		n += 1 + sovParams(uint64(m.PoolId))
	}
	l = m.MinimumFarmAmount.Size()
	n += 1 + l + sovParams(uint64(l))
	l = m.MinimumBidAmount.Size()
	n += 1 + l + sovParams(uint64(l))
	return n
}

func sovParams(x uint64) (n int) {
	return (math_bits.Len64(x|1) + 6) / 7
}
func sozParams(x uint64) (n int) {
	return sovParams(uint64((x << 1) ^ uint64((int64(x) >> 63))))
}
func (m *Params) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowParams
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
			return fmt.Errorf("proto: Params: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: Params: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field LiquidFarmCreationFee", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowParams
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
				return ErrInvalidLengthParams
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthParams
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.LiquidFarmCreationFee = append(m.LiquidFarmCreationFee, types.Coin{})
			if err := m.LiquidFarmCreationFee[len(m.LiquidFarmCreationFee)-1].Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 2:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field DelayedFarmGasFee", wireType)
			}
			m.DelayedFarmGasFee = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowParams
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.DelayedFarmGasFee |= github_com_cosmos_cosmos_sdk_types.Gas(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 3:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field LiquidFarms", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowParams
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
				return ErrInvalidLengthParams
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthParams
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.LiquidFarms = append(m.LiquidFarms, LiquidFarm{})
			if err := m.LiquidFarms[len(m.LiquidFarms)-1].Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipParams(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthParams
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
func (m *LiquidFarm) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowParams
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
			return fmt.Errorf("proto: LiquidFarm: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: LiquidFarm: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field PoolId", wireType)
			}
			m.PoolId = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowParams
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
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field MinimumFarmAmount", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowParams
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
				return ErrInvalidLengthParams
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthParams
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if err := m.MinimumFarmAmount.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 3:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field MinimumBidAmount", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowParams
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
				return ErrInvalidLengthParams
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthParams
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if err := m.MinimumBidAmount.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipParams(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthParams
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
func skipParams(dAtA []byte) (n int, err error) {
	l := len(dAtA)
	iNdEx := 0
	depth := 0
	for iNdEx < l {
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return 0, ErrIntOverflowParams
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
					return 0, ErrIntOverflowParams
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
					return 0, ErrIntOverflowParams
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
				return 0, ErrInvalidLengthParams
			}
			iNdEx += length
		case 3:
			depth++
		case 4:
			if depth == 0 {
				return 0, ErrUnexpectedEndOfGroupParams
			}
			depth--
		case 5:
			iNdEx += 4
		default:
			return 0, fmt.Errorf("proto: illegal wireType %d", wireType)
		}
		if iNdEx < 0 {
			return 0, ErrInvalidLengthParams
		}
		if depth == 0 {
			return iNdEx, nil
		}
	}
	return 0, io.ErrUnexpectedEOF
}

var (
	ErrInvalidLengthParams        = fmt.Errorf("proto: negative length found during unmarshaling")
	ErrIntOverflowParams          = fmt.Errorf("proto: integer overflow")
	ErrUnexpectedEndOfGroupParams = fmt.Errorf("proto: unexpected end of group")
)
