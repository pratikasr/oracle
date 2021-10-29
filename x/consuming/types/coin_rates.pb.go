// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: consuming/coin_rates.proto

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

type CoinRatesCallData struct {
	Symbols    []string `protobuf:"bytes,1,rep,name=symbols,proto3" json:"symbols,omitempty"`
	Multiplier uint64   `protobuf:"varint,2,opt,name=multiplier,proto3" json:"multiplier,omitempty"`
}

func (m *CoinRatesCallData) Reset()         { *m = CoinRatesCallData{} }
func (m *CoinRatesCallData) String() string { return proto.CompactTextString(m) }
func (*CoinRatesCallData) ProtoMessage()    {}
func (*CoinRatesCallData) Descriptor() ([]byte, []int) {
	return fileDescriptor_18d98c2f0f662644, []int{0}
}
func (m *CoinRatesCallData) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *CoinRatesCallData) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_CoinRatesCallData.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *CoinRatesCallData) XXX_Merge(src proto.Message) {
	xxx_messageInfo_CoinRatesCallData.Merge(m, src)
}
func (m *CoinRatesCallData) XXX_Size() int {
	return m.Size()
}
func (m *CoinRatesCallData) XXX_DiscardUnknown() {
	xxx_messageInfo_CoinRatesCallData.DiscardUnknown(m)
}

var xxx_messageInfo_CoinRatesCallData proto.InternalMessageInfo

func (m *CoinRatesCallData) GetSymbols() []string {
	if m != nil {
		return m.Symbols
	}
	return nil
}

func (m *CoinRatesCallData) GetMultiplier() uint64 {
	if m != nil {
		return m.Multiplier
	}
	return 0
}

type CoinRatesResult struct {
	Rates []uint64 `protobuf:"varint,1,rep,packed,name=rates,proto3" json:"rates,omitempty"`
}

func (m *CoinRatesResult) Reset()         { *m = CoinRatesResult{} }
func (m *CoinRatesResult) String() string { return proto.CompactTextString(m) }
func (*CoinRatesResult) ProtoMessage()    {}
func (*CoinRatesResult) Descriptor() ([]byte, []int) {
	return fileDescriptor_18d98c2f0f662644, []int{1}
}
func (m *CoinRatesResult) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *CoinRatesResult) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_CoinRatesResult.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *CoinRatesResult) XXX_Merge(src proto.Message) {
	xxx_messageInfo_CoinRatesResult.Merge(m, src)
}
func (m *CoinRatesResult) XXX_Size() int {
	return m.Size()
}
func (m *CoinRatesResult) XXX_DiscardUnknown() {
	xxx_messageInfo_CoinRatesResult.DiscardUnknown(m)
}

var xxx_messageInfo_CoinRatesResult proto.InternalMessageInfo

func (m *CoinRatesResult) GetRates() []uint64 {
	if m != nil {
		return m.Rates
	}
	return nil
}

func init() {
	proto.RegisterType((*CoinRatesCallData)(nil), "cosmonaut.oracle.consuming.CoinRatesCallData")
	proto.RegisterType((*CoinRatesResult)(nil), "cosmonaut.oracle.consuming.CoinRatesResult")
}

func init() { proto.RegisterFile("consuming/coin_rates.proto", fileDescriptor_18d98c2f0f662644) }

var fileDescriptor_18d98c2f0f662644 = []byte{
	// 221 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xe2, 0x92, 0x4a, 0xce, 0xcf, 0x2b,
	0x2e, 0xcd, 0xcd, 0xcc, 0x4b, 0xd7, 0x4f, 0xce, 0xcf, 0xcc, 0x8b, 0x2f, 0x4a, 0x2c, 0x49, 0x2d,
	0xd6, 0x2b, 0x28, 0xca, 0x2f, 0xc9, 0x17, 0x92, 0x4a, 0xce, 0x2f, 0xce, 0xcd, 0xcf, 0x4b, 0x2c,
	0x2d, 0xd1, 0xcb, 0x2f, 0x4a, 0x4c, 0xce, 0x49, 0xd5, 0x83, 0x2b, 0x56, 0xf2, 0xe5, 0x12, 0x74,
	0xce, 0xcf, 0xcc, 0x0b, 0x02, 0x29, 0x77, 0x4e, 0xcc, 0xc9, 0x71, 0x49, 0x2c, 0x49, 0x14, 0x92,
	0xe0, 0x62, 0x2f, 0xae, 0xcc, 0x4d, 0xca, 0xcf, 0x29, 0x96, 0x60, 0x54, 0x60, 0xd6, 0xe0, 0x0c,
	0x82, 0x71, 0x85, 0xe4, 0xb8, 0xb8, 0x72, 0x4b, 0x73, 0x4a, 0x32, 0x0b, 0x72, 0x32, 0x53, 0x8b,
	0x24, 0x98, 0x14, 0x18, 0x35, 0x58, 0x82, 0x90, 0x44, 0x94, 0xd4, 0xb9, 0xf8, 0xe1, 0xc6, 0x05,
	0xa5, 0x16, 0x97, 0xe6, 0x94, 0x08, 0x89, 0x70, 0xb1, 0x82, 0x1d, 0x03, 0x36, 0x8a, 0x25, 0x08,
	0xc2, 0x71, 0x72, 0x3f, 0xf1, 0x48, 0x8e, 0xf1, 0xc2, 0x23, 0x39, 0xc6, 0x07, 0x8f, 0xe4, 0x18,
	0x27, 0x3c, 0x96, 0x63, 0xb8, 0xf0, 0x58, 0x8e, 0xe1, 0xc6, 0x63, 0x39, 0x86, 0x28, 0xdd, 0xf4,
	0xcc, 0x92, 0x8c, 0xd2, 0x24, 0xbd, 0xe4, 0xfc, 0x5c, 0x7d, 0xb8, 0xc3, 0xf5, 0x21, 0x0e, 0xd7,
	0xaf, 0xd0, 0x47, 0xf8, 0xb3, 0xa4, 0xb2, 0x20, 0xb5, 0x38, 0x89, 0x0d, 0xec, 0x47, 0x63, 0x40,
	0x00, 0x00, 0x00, 0xff, 0xff, 0x44, 0x45, 0x44, 0x93, 0x01, 0x01, 0x00, 0x00,
}

func (m *CoinRatesCallData) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *CoinRatesCallData) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *CoinRatesCallData) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if m.Multiplier != 0 {
		i = encodeVarintCoinRates(dAtA, i, uint64(m.Multiplier))
		i--
		dAtA[i] = 0x10
	}
	if len(m.Symbols) > 0 {
		for iNdEx := len(m.Symbols) - 1; iNdEx >= 0; iNdEx-- {
			i -= len(m.Symbols[iNdEx])
			copy(dAtA[i:], m.Symbols[iNdEx])
			i = encodeVarintCoinRates(dAtA, i, uint64(len(m.Symbols[iNdEx])))
			i--
			dAtA[i] = 0xa
		}
	}
	return len(dAtA) - i, nil
}

func (m *CoinRatesResult) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *CoinRatesResult) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *CoinRatesResult) MarshalToSizedBuffer(dAtA []byte) (int, error) {
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
		i = encodeVarintCoinRates(dAtA, i, uint64(j1))
		i--
		dAtA[i] = 0xa
	}
	return len(dAtA) - i, nil
}

func encodeVarintCoinRates(dAtA []byte, offset int, v uint64) int {
	offset -= sovCoinRates(v)
	base := offset
	for v >= 1<<7 {
		dAtA[offset] = uint8(v&0x7f | 0x80)
		v >>= 7
		offset++
	}
	dAtA[offset] = uint8(v)
	return base
}
func (m *CoinRatesCallData) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	if len(m.Symbols) > 0 {
		for _, s := range m.Symbols {
			l = len(s)
			n += 1 + l + sovCoinRates(uint64(l))
		}
	}
	if m.Multiplier != 0 {
		n += 1 + sovCoinRates(uint64(m.Multiplier))
	}
	return n
}

func (m *CoinRatesResult) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	if len(m.Rates) > 0 {
		l = 0
		for _, e := range m.Rates {
			l += sovCoinRates(uint64(e))
		}
		n += 1 + sovCoinRates(uint64(l)) + l
	}
	return n
}

func sovCoinRates(x uint64) (n int) {
	return (math_bits.Len64(x|1) + 6) / 7
}
func sozCoinRates(x uint64) (n int) {
	return sovCoinRates(uint64((x << 1) ^ uint64((int64(x) >> 63))))
}
func (m *CoinRatesCallData) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowCoinRates
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
			return fmt.Errorf("proto: CoinRatesCallData: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: CoinRatesCallData: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Symbols", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowCoinRates
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
				return ErrInvalidLengthCoinRates
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthCoinRates
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
					return ErrIntOverflowCoinRates
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
			skippy, err := skipCoinRates(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthCoinRates
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
func (m *CoinRatesResult) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowCoinRates
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
			return fmt.Errorf("proto: CoinRatesResult: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: CoinRatesResult: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType == 0 {
				var v uint64
				for shift := uint(0); ; shift += 7 {
					if shift >= 64 {
						return ErrIntOverflowCoinRates
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
						return ErrIntOverflowCoinRates
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
					return ErrInvalidLengthCoinRates
				}
				postIndex := iNdEx + packedLen
				if postIndex < 0 {
					return ErrInvalidLengthCoinRates
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
							return ErrIntOverflowCoinRates
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
			skippy, err := skipCoinRates(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthCoinRates
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
func skipCoinRates(dAtA []byte) (n int, err error) {
	l := len(dAtA)
	iNdEx := 0
	depth := 0
	for iNdEx < l {
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return 0, ErrIntOverflowCoinRates
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
					return 0, ErrIntOverflowCoinRates
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
					return 0, ErrIntOverflowCoinRates
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
				return 0, ErrInvalidLengthCoinRates
			}
			iNdEx += length
		case 3:
			depth++
		case 4:
			if depth == 0 {
				return 0, ErrUnexpectedEndOfGroupCoinRates
			}
			depth--
		case 5:
			iNdEx += 4
		default:
			return 0, fmt.Errorf("proto: illegal wireType %d", wireType)
		}
		if iNdEx < 0 {
			return 0, ErrInvalidLengthCoinRates
		}
		if depth == 0 {
			return iNdEx, nil
		}
	}
	return 0, io.ErrUnexpectedEOF
}

var (
	ErrInvalidLengthCoinRates        = fmt.Errorf("proto: negative length found during unmarshaling")
	ErrIntOverflowCoinRates          = fmt.Errorf("proto: integer overflow")
	ErrUnexpectedEndOfGroupCoinRates = fmt.Errorf("proto: unexpected end of group")
)