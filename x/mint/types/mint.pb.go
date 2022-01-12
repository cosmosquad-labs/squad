// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: crescent/mint/v1beta1/mint.proto

package types

import (
	fmt "fmt"
	github_com_cosmos_cosmos_sdk_types "github.com/cosmos/cosmos-sdk/types"
	_ "github.com/gogo/protobuf/gogoproto"
	proto "github.com/gogo/protobuf/proto"
	github_com_gogo_protobuf_types "github.com/gogo/protobuf/types"
	_ "google.golang.org/protobuf/types/known/durationpb"
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

// Params holds parameters for the mint module.
type Params struct {
	// type of coin to mint
	MintDenom          string        `protobuf:"bytes,1,opt,name=mint_denom,json=mintDenom,proto3" json:"mint_denom,omitempty"`
	BlockTimeThreshold time.Duration `protobuf:"bytes,2,opt,name=block_time_threshold,json=blockTimeThreshold,proto3,stdduration" json:"block_time_threshold,omitempty" yaml:"block_time_threshold"`
}

func (m *Params) Reset()      { *m = Params{} }
func (*Params) ProtoMessage() {}
func (*Params) Descriptor() ([]byte, []int) {
	return fileDescriptor_fe08af702efa1523, []int{0}
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

func (m *Params) GetMintDenom() string {
	if m != nil {
		return m.MintDenom
	}
	return ""
}

func (m *Params) GetBlockTimeThreshold() time.Duration {
	if m != nil {
		return m.BlockTimeThreshold
	}
	return 0
}

type InflationPeriod struct {
	StartTime time.Time                              `protobuf:"bytes,1,opt,name=start_time,json=startTime,proto3,stdtime" json:"start_time" yaml:"start_time"`
	EndTime   time.Time                              `protobuf:"bytes,2,opt,name=end_time,json=endTime,proto3,stdtime" json:"end_time" yaml:"end_time"`
	Amount    github_com_cosmos_cosmos_sdk_types.Int `protobuf:"bytes,3,opt,name=amount,proto3,customtype=github.com/cosmos/cosmos-sdk/types.Int" json:"amount"`
}

func (m *InflationPeriod) Reset()         { *m = InflationPeriod{} }
func (m *InflationPeriod) String() string { return proto.CompactTextString(m) }
func (*InflationPeriod) ProtoMessage()    {}
func (*InflationPeriod) Descriptor() ([]byte, []int) {
	return fileDescriptor_fe08af702efa1523, []int{1}
}
func (m *InflationPeriod) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *InflationPeriod) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_InflationPeriod.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *InflationPeriod) XXX_Merge(src proto.Message) {
	xxx_messageInfo_InflationPeriod.Merge(m, src)
}
func (m *InflationPeriod) XXX_Size() int {
	return m.Size()
}
func (m *InflationPeriod) XXX_DiscardUnknown() {
	xxx_messageInfo_InflationPeriod.DiscardUnknown(m)
}

var xxx_messageInfo_InflationPeriod proto.InternalMessageInfo

func (m *InflationPeriod) GetStartTime() time.Time {
	if m != nil {
		return m.StartTime
	}
	return time.Time{}
}

func (m *InflationPeriod) GetEndTime() time.Time {
	if m != nil {
		return m.EndTime
	}
	return time.Time{}
}

func init() {
	proto.RegisterType((*Params)(nil), "crescent.mint.v1beta1.Params")
	proto.RegisterType((*InflationPeriod)(nil), "crescent.mint.v1beta1.InflationPeriod")
}

func init() { proto.RegisterFile("crescent/mint/v1beta1/mint.proto", fileDescriptor_fe08af702efa1523) }

var fileDescriptor_fe08af702efa1523 = []byte{
	// 440 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x84, 0x92, 0xbf, 0x6f, 0xd3, 0x40,
	0x14, 0xc7, 0x7d, 0x01, 0x05, 0x72, 0x1d, 0x2a, 0xac, 0x22, 0x85, 0x54, 0xb5, 0x23, 0x0f, 0xa8,
	0x03, 0xb5, 0x95, 0xb2, 0x75, 0x8c, 0x2a, 0xa4, 0x88, 0xa5, 0x8a, 0x32, 0x20, 0x96, 0xc8, 0x3f,
	0xae, 0x8e, 0x15, 0xdf, 0xbd, 0xe8, 0xee, 0x19, 0xc8, 0x5f, 0xc0, 0xda, 0x09, 0x75, 0xe4, 0xaf,
	0x41, 0x1d, 0x3b, 0x22, 0x06, 0x83, 0x92, 0x8d, 0xb1, 0x7f, 0x01, 0xba, 0x3b, 0x1b, 0x10, 0x44,
	0xea, 0x64, 0xdf, 0x7b, 0xdf, 0xef, 0xe7, 0xfb, 0xee, 0xf4, 0xe8, 0x30, 0x95, 0x4c, 0xa5, 0x4c,
	0x60, 0xc4, 0x0b, 0x81, 0xd1, 0xbb, 0x51, 0xc2, 0x30, 0x1e, 0x99, 0x43, 0xb8, 0x92, 0x80, 0xe0,
	0x3e, 0x6d, 0x15, 0xa1, 0x29, 0x36, 0x8a, 0xc1, 0x41, 0x0e, 0x39, 0x18, 0x45, 0xa4, 0xff, 0xac,
	0x78, 0xe0, 0xe5, 0x00, 0x79, 0xc9, 0x22, 0x73, 0x4a, 0xaa, 0xcb, 0x28, 0xab, 0x64, 0x8c, 0x05,
	0x88, 0xa6, 0xef, 0xff, 0xdb, 0xc7, 0x82, 0x33, 0x85, 0x31, 0x5f, 0x59, 0x41, 0xf0, 0x85, 0xd0,
	0xee, 0x45, 0x2c, 0x63, 0xae, 0xdc, 0x23, 0x4a, 0x75, 0xe2, 0x3c, 0x63, 0x02, 0x78, 0x9f, 0x0c,
	0xc9, 0x71, 0x6f, 0xda, 0xd3, 0x95, 0x73, 0x5d, 0x70, 0x3f, 0x11, 0x7a, 0x90, 0x94, 0x90, 0x2e,
	0xe7, 0x9a, 0x31, 0xc7, 0x85, 0x64, 0x6a, 0x01, 0x65, 0xd6, 0xef, 0x0c, 0xc9, 0xf1, 0xde, 0xe9,
	0xb3, 0xd0, 0x46, 0x85, 0x6d, 0x54, 0x78, 0xde, 0x8c, 0x32, 0x9e, 0xdc, 0xd4, 0xbe, 0xf3, 0xb3,
	0xf6, 0xbd, 0x5d, 0xf6, 0x17, 0xc0, 0x0b, 0x64, 0x7c, 0x85, 0xeb, 0xbb, 0xda, 0x3f, 0x5c, 0xc7,
	0xbc, 0x3c, 0x0b, 0x76, 0xe9, 0x82, 0xeb, 0xef, 0x3e, 0x99, 0xba, 0xa6, 0x35, 0x2b, 0x38, 0x9b,
	0xb5, 0x8d, 0xb3, 0x87, 0xd7, 0x9f, 0x7d, 0x27, 0xf8, 0xd8, 0xa1, 0xfb, 0x13, 0x71, 0x59, 0x9a,
	0xc8, 0x0b, 0x26, 0x0b, 0xc8, 0xdc, 0x37, 0x94, 0x2a, 0x8c, 0x25, 0x1a, 0x94, 0xb9, 0xd1, 0xde,
	0xe9, 0xe0, 0xbf, 0x39, 0x67, 0xed, 0x93, 0x8c, 0x8f, 0xf4, 0xa0, 0x77, 0xb5, 0xff, 0xc4, 0x8e,
	0xf1, 0xc7, 0x1b, 0x5c, 0xe9, 0xf0, 0x9e, 0x29, 0x68, 0xb9, 0x3b, 0xa5, 0x8f, 0x99, 0xc8, 0x2c,
	0xb7, 0x73, 0x2f, 0xf7, 0xb0, 0xe1, 0xee, 0x5b, 0x6e, 0xeb, 0xb4, 0xd4, 0x47, 0x4c, 0x64, 0x86,
	0xf9, 0x8a, 0x76, 0x63, 0x0e, 0x95, 0xc0, 0xfe, 0x03, 0xfd, 0xf6, 0xe3, 0x50, 0xbb, 0xbe, 0xd5,
	0xfe, 0xf3, 0xbc, 0xc0, 0x45, 0x95, 0x84, 0x29, 0xf0, 0x28, 0x05, 0xc5, 0x41, 0x35, 0x9f, 0x13,
	0x95, 0x2d, 0x23, 0x5c, 0xaf, 0x98, 0x0a, 0x27, 0x02, 0xa7, 0x8d, 0x7b, 0xfc, 0xfa, 0x66, 0xe3,
	0x91, 0xdb, 0x8d, 0x47, 0x7e, 0x6c, 0x3c, 0x72, 0xb5, 0xf5, 0x9c, 0xdb, 0xad, 0xe7, 0x7c, 0xdd,
	0x7a, 0xce, 0xdb, 0xd1, 0xdf, 0xa4, 0x66, 0xcb, 0x4e, 0x04, 0xc3, 0xf7, 0x20, 0x97, 0xbf, 0x0b,
	0xd1, 0x07, 0xbb, 0x9a, 0x06, 0x9c, 0x74, 0xcd, 0x75, 0x5e, 0xfe, 0x0a, 0x00, 0x00, 0xff, 0xff,
	0x6e, 0x86, 0x9b, 0x6e, 0xb8, 0x02, 0x00, 0x00,
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
	n1, err1 := github_com_gogo_protobuf_types.StdDurationMarshalTo(m.BlockTimeThreshold, dAtA[i-github_com_gogo_protobuf_types.SizeOfStdDuration(m.BlockTimeThreshold):])
	if err1 != nil {
		return 0, err1
	}
	i -= n1
	i = encodeVarintMint(dAtA, i, uint64(n1))
	i--
	dAtA[i] = 0x12
	if len(m.MintDenom) > 0 {
		i -= len(m.MintDenom)
		copy(dAtA[i:], m.MintDenom)
		i = encodeVarintMint(dAtA, i, uint64(len(m.MintDenom)))
		i--
		dAtA[i] = 0xa
	}
	return len(dAtA) - i, nil
}

func (m *InflationPeriod) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *InflationPeriod) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *InflationPeriod) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	{
		size := m.Amount.Size()
		i -= size
		if _, err := m.Amount.MarshalTo(dAtA[i:]); err != nil {
			return 0, err
		}
		i = encodeVarintMint(dAtA, i, uint64(size))
	}
	i--
	dAtA[i] = 0x1a
	n2, err2 := github_com_gogo_protobuf_types.StdTimeMarshalTo(m.EndTime, dAtA[i-github_com_gogo_protobuf_types.SizeOfStdTime(m.EndTime):])
	if err2 != nil {
		return 0, err2
	}
	i -= n2
	i = encodeVarintMint(dAtA, i, uint64(n2))
	i--
	dAtA[i] = 0x12
	n3, err3 := github_com_gogo_protobuf_types.StdTimeMarshalTo(m.StartTime, dAtA[i-github_com_gogo_protobuf_types.SizeOfStdTime(m.StartTime):])
	if err3 != nil {
		return 0, err3
	}
	i -= n3
	i = encodeVarintMint(dAtA, i, uint64(n3))
	i--
	dAtA[i] = 0xa
	return len(dAtA) - i, nil
}

func encodeVarintMint(dAtA []byte, offset int, v uint64) int {
	offset -= sovMint(v)
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
	l = len(m.MintDenom)
	if l > 0 {
		n += 1 + l + sovMint(uint64(l))
	}
	l = github_com_gogo_protobuf_types.SizeOfStdDuration(m.BlockTimeThreshold)
	n += 1 + l + sovMint(uint64(l))
	return n
}

func (m *InflationPeriod) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	l = github_com_gogo_protobuf_types.SizeOfStdTime(m.StartTime)
	n += 1 + l + sovMint(uint64(l))
	l = github_com_gogo_protobuf_types.SizeOfStdTime(m.EndTime)
	n += 1 + l + sovMint(uint64(l))
	l = m.Amount.Size()
	n += 1 + l + sovMint(uint64(l))
	return n
}

func sovMint(x uint64) (n int) {
	return (math_bits.Len64(x|1) + 6) / 7
}
func sozMint(x uint64) (n int) {
	return sovMint(uint64((x << 1) ^ uint64((int64(x) >> 63))))
}
func (m *Params) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowMint
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
				return fmt.Errorf("proto: wrong wireType = %d for field MintDenom", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowMint
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
				return ErrInvalidLengthMint
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthMint
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.MintDenom = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field BlockTimeThreshold", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowMint
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
				return ErrInvalidLengthMint
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthMint
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if err := github_com_gogo_protobuf_types.StdDurationUnmarshal(&m.BlockTimeThreshold, dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipMint(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthMint
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
func (m *InflationPeriod) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowMint
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
			return fmt.Errorf("proto: InflationPeriod: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: InflationPeriod: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field StartTime", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowMint
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
				return ErrInvalidLengthMint
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthMint
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if err := github_com_gogo_protobuf_types.StdTimeUnmarshal(&m.StartTime, dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field EndTime", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowMint
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
				return ErrInvalidLengthMint
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthMint
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if err := github_com_gogo_protobuf_types.StdTimeUnmarshal(&m.EndTime, dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 3:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Amount", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowMint
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
				return ErrInvalidLengthMint
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthMint
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if err := m.Amount.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipMint(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthMint
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
func skipMint(dAtA []byte) (n int, err error) {
	l := len(dAtA)
	iNdEx := 0
	depth := 0
	for iNdEx < l {
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return 0, ErrIntOverflowMint
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
					return 0, ErrIntOverflowMint
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
					return 0, ErrIntOverflowMint
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
				return 0, ErrInvalidLengthMint
			}
			iNdEx += length
		case 3:
			depth++
		case 4:
			if depth == 0 {
				return 0, ErrUnexpectedEndOfGroupMint
			}
			depth--
		case 5:
			iNdEx += 4
		default:
			return 0, fmt.Errorf("proto: illegal wireType %d", wireType)
		}
		if iNdEx < 0 {
			return 0, ErrInvalidLengthMint
		}
		if depth == 0 {
			return iNdEx, nil
		}
	}
	return 0, io.ErrUnexpectedEOF
}

var (
	ErrInvalidLengthMint        = fmt.Errorf("proto: negative length found during unmarshaling")
	ErrIntOverflowMint          = fmt.Errorf("proto: integer overflow")
	ErrUnexpectedEndOfGroupMint = fmt.Errorf("proto: unexpected end of group")
)
