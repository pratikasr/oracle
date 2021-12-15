// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: bandoracle/fetch_price.proto

package types

import (
	fmt "fmt"
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

type FetchPriceCallData struct {
	Symbols    []string `protobuf:"bytes,1,rep,name=symbols,proto3" json:"symbols,omitempty"`
	Multiplier uint64   `protobuf:"varint,2,opt,name=multiplier,proto3" json:"multiplier,omitempty"`
}

func (m *FetchPriceCallData) Reset()         { *m = FetchPriceCallData{} }
func (m *FetchPriceCallData) String() string { return proto.CompactTextString(m) }
func (*FetchPriceCallData) ProtoMessage()    {}
func (*FetchPriceCallData) Descriptor() ([]byte, []int) {
	return fileDescriptor_53a425633703e201, []int{0}
}
func (m *FetchPriceCallData) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *FetchPriceCallData) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_FetchPriceCallData.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *FetchPriceCallData) XXX_Merge(src proto.Message) {
	xxx_messageInfo_FetchPriceCallData.Merge(m, src)
}
func (m *FetchPriceCallData) XXX_Size() int {
	return m.Size()
}
func (m *FetchPriceCallData) XXX_DiscardUnknown() {
	xxx_messageInfo_FetchPriceCallData.DiscardUnknown(m)
}

var xxx_messageInfo_FetchPriceCallData proto.InternalMessageInfo

func (m *FetchPriceCallData) GetSymbols() []string {
	if m != nil {
		return m.Symbols
	}
	return nil
}

func (m *FetchPriceCallData) GetMultiplier() uint64 {
	if m != nil {
		return m.Multiplier
	}
	return 0
}

type FetchPriceResult struct {
	Rates []uint64 `protobuf:"varint,1,rep,packed,name=rates,proto3" json:"rates,omitempty"`
}

func (m *FetchPriceResult) Reset()         { *m = FetchPriceResult{} }
func (m *FetchPriceResult) String() string { return proto.CompactTextString(m) }
func (*FetchPriceResult) ProtoMessage()    {}
func (*FetchPriceResult) Descriptor() ([]byte, []int) {
	return fileDescriptor_53a425633703e201, []int{1}
}
func (m *FetchPriceResult) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *FetchPriceResult) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_FetchPriceResult.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *FetchPriceResult) XXX_Merge(src proto.Message) {
	xxx_messageInfo_FetchPriceResult.Merge(m, src)
}
func (m *FetchPriceResult) XXX_Size() int {
	return m.Size()
}
func (m *FetchPriceResult) XXX_DiscardUnknown() {
	xxx_messageInfo_FetchPriceResult.DiscardUnknown(m)
}

var xxx_messageInfo_FetchPriceResult proto.InternalMessageInfo

func (m *FetchPriceResult) GetRates() []uint64 {
	if m != nil {
		return m.Rates
	}
	return nil
}

func init() {
	proto.RegisterType((*FetchPriceCallData)(nil), "cosmonaut.oracle.bandoracle.FetchPriceCallData")
	proto.RegisterType((*FetchPriceResult)(nil), "cosmonaut.oracle.bandoracle.FetchPriceResult")
}

func init() { proto.RegisterFile("bandoracle/fetch_price.proto", fileDescriptor_53a425633703e201) }

var fileDescriptor_53a425633703e201 = []byte{
	// 218 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xe2, 0x92, 0x49, 0x4a, 0xcc, 0x4b,
	0xc9, 0x2f, 0x4a, 0x4c, 0xce, 0x49, 0xd5, 0x4f, 0x4b, 0x2d, 0x49, 0xce, 0x88, 0x2f, 0x28, 0xca,
	0x4c, 0x4e, 0xd5, 0x2b, 0x28, 0xca, 0x2f, 0xc9, 0x17, 0x92, 0x4e, 0xce, 0x2f, 0xce, 0xcd, 0xcf,
	0x4b, 0x2c, 0x2d, 0xd1, 0x83, 0xa8, 0xd1, 0x43, 0x28, 0x57, 0xf2, 0xe3, 0x12, 0x72, 0x03, 0xe9,
	0x08, 0x00, 0x69, 0x70, 0x4e, 0xcc, 0xc9, 0x71, 0x49, 0x2c, 0x49, 0x14, 0x92, 0xe0, 0x62, 0x2f,
	0xae, 0xcc, 0x4d, 0xca, 0xcf, 0x29, 0x96, 0x60, 0x54, 0x60, 0xd6, 0xe0, 0x0c, 0x82, 0x71, 0x85,
	0xe4, 0xb8, 0xb8, 0x72, 0x4b, 0x73, 0x4a, 0x32, 0x0b, 0x72, 0x32, 0x53, 0x8b, 0x24, 0x98, 0x14,
	0x18, 0x35, 0x58, 0x82, 0x90, 0x44, 0x94, 0x34, 0xb8, 0x04, 0x10, 0xe6, 0x05, 0xa5, 0x16, 0x97,
	0xe6, 0x94, 0x08, 0x89, 0x70, 0xb1, 0x16, 0x25, 0x96, 0xa4, 0x42, 0xcc, 0x62, 0x09, 0x82, 0x70,
	0x9c, 0x3c, 0x4e, 0x3c, 0x92, 0x63, 0xbc, 0xf0, 0x48, 0x8e, 0xf1, 0xc1, 0x23, 0x39, 0xc6, 0x09,
	0x8f, 0xe5, 0x18, 0x2e, 0x3c, 0x96, 0x63, 0xb8, 0xf1, 0x58, 0x8e, 0x21, 0x4a, 0x2f, 0x3d, 0xb3,
	0x24, 0xa3, 0x34, 0x49, 0x2f, 0x39, 0x3f, 0x57, 0x1f, 0xee, 0x76, 0x7d, 0xa8, 0xff, 0x2a, 0xf4,
	0x91, 0x3c, 0x5b, 0x52, 0x59, 0x90, 0x5a, 0x9c, 0xc4, 0x06, 0xf6, 0xa7, 0x31, 0x20, 0x00, 0x00,
	0xff, 0xff, 0x74, 0x4b, 0x6b, 0x75, 0x07, 0x01, 0x00, 0x00,
}

func (m *FetchPriceCallData) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *FetchPriceCallData) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *FetchPriceCallData) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if m.Multiplier != 0 {
		i = encodeVarintFetchPrice(dAtA, i, uint64(m.Multiplier))
		i--
		dAtA[i] = 0x10
	}
	if len(m.Symbols) > 0 {
		for iNdEx := len(m.Symbols) - 1; iNdEx >= 0; iNdEx-- {
			i -= len(m.Symbols[iNdEx])
			copy(dAtA[i:], m.Symbols[iNdEx])
			i = encodeVarintFetchPrice(dAtA, i, uint64(len(m.Symbols[iNdEx])))
			i--
			dAtA[i] = 0xa
		}
	}
	return len(dAtA) - i, nil
}

func (m *FetchPriceResult) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *FetchPriceResult) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *FetchPriceResult) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if len(m.Rates) > 0 {
		dAtA2 := make([]byte, len(m.Rates)*10)
		var j1 int
		for _, num := range m.Rates {
			for num >= 1<<7 {
				dAtA2[j1] = uint8(uint64(num)&0x7f | 0x80)
				num >>= 7
				j1++
			}
			dAtA2[j1] = uint8(num)
			j1++
		}
		i -= j1
		copy(dAtA[i:], dAtA2[:j1])
		i = encodeVarintFetchPrice(dAtA, i, uint64(j1))
		i--
		dAtA[i] = 0xa
	}
	return len(dAtA) - i, nil
}

func encodeVarintFetchPrice(dAtA []byte, offset int, v uint64) int {
	offset -= sovFetchPrice(v)
	base := offset
	for v >= 1<<7 {
		dAtA[offset] = uint8(v&0x7f | 0x80)
		v >>= 7
		offset++
	}
	dAtA[offset] = uint8(v)
	return base
}
func (m *FetchPriceCallData) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	if len(m.Symbols) > 0 {
		for _, s := range m.Symbols {
			l = len(s)
			n += 1 + l + sovFetchPrice(uint64(l))
		}
	}
	if m.Multiplier != 0 {
		n += 1 + sovFetchPrice(uint64(m.Multiplier))
	}
	return n
}

func (m *FetchPriceResult) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	if len(m.Rates) > 0 {
		l = 0
		for _, e := range m.Rates {
			l += sovFetchPrice(uint64(e))
		}
		n += 1 + sovFetchPrice(uint64(l)) + l
	}
	return n
}

func sovFetchPrice(x uint64) (n int) {
	return (math_bits.Len64(x|1) + 6) / 7
}
func sozFetchPrice(x uint64) (n int) {
	return sovFetchPrice(uint64((x << 1) ^ uint64((int64(x) >> 63))))
}
func (m *FetchPriceCallData) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowFetchPrice
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
			return fmt.Errorf("proto: FetchPriceCallData: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: FetchPriceCallData: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Symbols", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowFetchPrice
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
				return ErrInvalidLengthFetchPrice
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthFetchPrice
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Symbols = append(m.Symbols, string(dAtA[iNdEx:postIndex]))
			iNdEx = postIndex
		case 2:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field Multiplier", wireType)
			}
			m.Multiplier = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowFetchPrice
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.Multiplier |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		default:
			iNdEx = preIndex
			skippy, err := skipFetchPrice(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthFetchPrice
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
func (m *FetchPriceResult) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowFetchPrice
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
			return fmt.Errorf("proto: FetchPriceResult: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: FetchPriceResult: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType == 0 {
				var v uint64
				for shift := uint(0); ; shift += 7 {
					if shift >= 64 {
						return ErrIntOverflowFetchPrice
					}
					if iNdEx >= l {
						return io.ErrUnexpectedEOF
					}
					b := dAtA[iNdEx]
					iNdEx++
					v |= uint64(b&0x7F) << shift
					if b < 0x80 {
						break
					}
				}
				m.Rates = append(m.Rates, v)
			} else if wireType == 2 {
				var packedLen int
				for shift := uint(0); ; shift += 7 {
					if shift >= 64 {
						return ErrIntOverflowFetchPrice
					}
					if iNdEx >= l {
						return io.ErrUnexpectedEOF
					}
					b := dAtA[iNdEx]
					iNdEx++
					packedLen |= int(b&0x7F) << shift
					if b < 0x80 {
						break
					}
				}
				if packedLen < 0 {
					return ErrInvalidLengthFetchPrice
				}
				postIndex := iNdEx + packedLen
				if postIndex < 0 {
					return ErrInvalidLengthFetchPrice
				}
				if postIndex > l {
					return io.ErrUnexpectedEOF
				}
				var elementCount int
				var count int
				for _, integer := range dAtA[iNdEx:postIndex] {
					if integer < 128 {
						count++
					}
				}
				elementCount = count
				if elementCount != 0 && len(m.Rates) == 0 {
					m.Rates = make([]uint64, 0, elementCount)
				}
				for iNdEx < postIndex {
					var v uint64
					for shift := uint(0); ; shift += 7 {
						if shift >= 64 {
							return ErrIntOverflowFetchPrice
						}
						if iNdEx >= l {
							return io.ErrUnexpectedEOF
						}
						b := dAtA[iNdEx]
						iNdEx++
						v |= uint64(b&0x7F) << shift
						if b < 0x80 {
							break
						}
					}
					m.Rates = append(m.Rates, v)
				}
			} else {
				return fmt.Errorf("proto: wrong wireType = %d for field Rates", wireType)
			}
		default:
			iNdEx = preIndex
			skippy, err := skipFetchPrice(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthFetchPrice
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
func skipFetchPrice(dAtA []byte) (n int, err error) {
	l := len(dAtA)
	iNdEx := 0
	depth := 0
	for iNdEx < l {
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return 0, ErrIntOverflowFetchPrice
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
					return 0, ErrIntOverflowFetchPrice
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
					return 0, ErrIntOverflowFetchPrice
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
				return 0, ErrInvalidLengthFetchPrice
			}
			iNdEx += length
		case 3:
			depth++
		case 4:
			if depth == 0 {
				return 0, ErrUnexpectedEndOfGroupFetchPrice
			}
			depth--
		case 5:
			iNdEx += 4
		default:
			return 0, fmt.Errorf("proto: illegal wireType %d", wireType)
		}
		if iNdEx < 0 {
			return 0, ErrInvalidLengthFetchPrice
		}
		if depth == 0 {
			return iNdEx, nil
		}
	}
	return 0, io.ErrUnexpectedEOF
}

var (
	ErrInvalidLengthFetchPrice        = fmt.Errorf("proto: negative length found during unmarshaling")
	ErrIntOverflowFetchPrice          = fmt.Errorf("proto: integer overflow")
	ErrUnexpectedEndOfGroupFetchPrice = fmt.Errorf("proto: unexpected end of group")
)
