// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: consuming/query.proto

package types

import (
	context "context"
	fmt "fmt"
	_ "github.com/cosmos/cosmos-sdk/types/query"
	grpc1 "github.com/gogo/protobuf/grpc"
	proto "github.com/gogo/protobuf/proto"
	_ "google.golang.org/genproto/googleapis/api/annotations"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
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

type QueryCoinRatesRequest struct {
	RequestId int64 `protobuf:"varint,1,opt,name=request_id,json=requestId,proto3" json:"request_id,omitempty"`
}

func (m *QueryCoinRatesRequest) Reset()         { *m = QueryCoinRatesRequest{} }
func (m *QueryCoinRatesRequest) String() string { return proto.CompactTextString(m) }
func (*QueryCoinRatesRequest) ProtoMessage()    {}
func (*QueryCoinRatesRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_38b0adbddb76ede3, []int{0}
}
func (m *QueryCoinRatesRequest) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *QueryCoinRatesRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_QueryCoinRatesRequest.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *QueryCoinRatesRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_QueryCoinRatesRequest.Merge(m, src)
}
func (m *QueryCoinRatesRequest) XXX_Size() int {
	return m.Size()
}
func (m *QueryCoinRatesRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_QueryCoinRatesRequest.DiscardUnknown(m)
}

var xxx_messageInfo_QueryCoinRatesRequest proto.InternalMessageInfo

func (m *QueryCoinRatesRequest) GetRequestId() int64 {
	if m != nil {
		return m.RequestId
	}
	return 0
}

type QueryCoinRatesResponse struct {
	Result *CoinRatesResult `protobuf:"bytes,1,opt,name=result,proto3" json:"result,omitempty"`
}

func (m *QueryCoinRatesResponse) Reset()         { *m = QueryCoinRatesResponse{} }
func (m *QueryCoinRatesResponse) String() string { return proto.CompactTextString(m) }
func (*QueryCoinRatesResponse) ProtoMessage()    {}
func (*QueryCoinRatesResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_38b0adbddb76ede3, []int{1}
}
func (m *QueryCoinRatesResponse) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *QueryCoinRatesResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_QueryCoinRatesResponse.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *QueryCoinRatesResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_QueryCoinRatesResponse.Merge(m, src)
}
func (m *QueryCoinRatesResponse) XXX_Size() int {
	return m.Size()
}
func (m *QueryCoinRatesResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_QueryCoinRatesResponse.DiscardUnknown(m)
}

var xxx_messageInfo_QueryCoinRatesResponse proto.InternalMessageInfo

func (m *QueryCoinRatesResponse) GetResult() *CoinRatesResult {
	if m != nil {
		return m.Result
	}
	return nil
}

type QueryLastCoinRatesIdRequest struct {
}

func (m *QueryLastCoinRatesIdRequest) Reset()         { *m = QueryLastCoinRatesIdRequest{} }
func (m *QueryLastCoinRatesIdRequest) String() string { return proto.CompactTextString(m) }
func (*QueryLastCoinRatesIdRequest) ProtoMessage()    {}
func (*QueryLastCoinRatesIdRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_38b0adbddb76ede3, []int{2}
}
func (m *QueryLastCoinRatesIdRequest) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *QueryLastCoinRatesIdRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_QueryLastCoinRatesIdRequest.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *QueryLastCoinRatesIdRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_QueryLastCoinRatesIdRequest.Merge(m, src)
}
func (m *QueryLastCoinRatesIdRequest) XXX_Size() int {
	return m.Size()
}
func (m *QueryLastCoinRatesIdRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_QueryLastCoinRatesIdRequest.DiscardUnknown(m)
}

var xxx_messageInfo_QueryLastCoinRatesIdRequest proto.InternalMessageInfo

type QueryLastCoinRatesIdResponse struct {
	RequestId int64 `protobuf:"varint,1,opt,name=request_id,json=requestId,proto3" json:"request_id,omitempty"`
}

func (m *QueryLastCoinRatesIdResponse) Reset()         { *m = QueryLastCoinRatesIdResponse{} }
func (m *QueryLastCoinRatesIdResponse) String() string { return proto.CompactTextString(m) }
func (*QueryLastCoinRatesIdResponse) ProtoMessage()    {}
func (*QueryLastCoinRatesIdResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_38b0adbddb76ede3, []int{3}
}
func (m *QueryLastCoinRatesIdResponse) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *QueryLastCoinRatesIdResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_QueryLastCoinRatesIdResponse.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *QueryLastCoinRatesIdResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_QueryLastCoinRatesIdResponse.Merge(m, src)
}
func (m *QueryLastCoinRatesIdResponse) XXX_Size() int {
	return m.Size()
}
func (m *QueryLastCoinRatesIdResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_QueryLastCoinRatesIdResponse.DiscardUnknown(m)
}

var xxx_messageInfo_QueryLastCoinRatesIdResponse proto.InternalMessageInfo

func (m *QueryLastCoinRatesIdResponse) GetRequestId() int64 {
	if m != nil {
		return m.RequestId
	}
	return 0
}

func init() {
	proto.RegisterType((*QueryCoinRatesRequest)(nil), "cosmonaut.oracle.consuming.QueryCoinRatesRequest")
	proto.RegisterType((*QueryCoinRatesResponse)(nil), "cosmonaut.oracle.consuming.QueryCoinRatesResponse")
	proto.RegisterType((*QueryLastCoinRatesIdRequest)(nil), "cosmonaut.oracle.consuming.QueryLastCoinRatesIdRequest")
	proto.RegisterType((*QueryLastCoinRatesIdResponse)(nil), "cosmonaut.oracle.consuming.QueryLastCoinRatesIdResponse")
}

func init() { proto.RegisterFile("consuming/query.proto", fileDescriptor_38b0adbddb76ede3) }

var fileDescriptor_38b0adbddb76ede3 = []byte{
	// 391 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x94, 0x92, 0x4f, 0x4a, 0xf3, 0x40,
	0x18, 0xc6, 0x9b, 0x7e, 0x7c, 0x05, 0xc7, 0x85, 0x30, 0x50, 0x91, 0xd8, 0x06, 0x09, 0x22, 0xe2,
	0x9f, 0x8c, 0x8d, 0xa0, 0xdd, 0xb8, 0xb1, 0x0b, 0x29, 0xb8, 0x31, 0x4b, 0x41, 0xc2, 0x24, 0x19,
	0xe2, 0x40, 0x3a, 0x93, 0x66, 0x26, 0x62, 0x11, 0x37, 0x9e, 0x40, 0xf0, 0x26, 0x9e, 0xc2, 0x65,
	0xc1, 0x8d, 0x4b, 0x69, 0xbd, 0x80, 0x37, 0x90, 0x4e, 0x62, 0x5b, 0x6a, 0xac, 0x74, 0x17, 0x26,
	0xef, 0xef, 0xc9, 0xef, 0x79, 0x33, 0xa0, 0xea, 0x73, 0x26, 0xd2, 0x0e, 0x65, 0x21, 0xea, 0xa6,
	0x24, 0xe9, 0x59, 0x71, 0xc2, 0x25, 0x87, 0xba, 0xcf, 0x45, 0x87, 0x33, 0x9c, 0x4a, 0x8b, 0x27,
	0xd8, 0x8f, 0x88, 0x35, 0x9e, 0xd3, 0x6b, 0x21, 0xe7, 0x61, 0x44, 0x10, 0x8e, 0x29, 0xc2, 0x8c,
	0x71, 0x89, 0x25, 0xe5, 0x4c, 0x64, 0xa4, 0xbe, 0xa3, 0x48, 0x81, 0x3c, 0x2c, 0x48, 0x16, 0x89,
	0x6e, 0x1a, 0x1e, 0x91, 0xb8, 0x81, 0x62, 0x1c, 0x52, 0xa6, 0x86, 0xf3, 0x59, 0x7d, 0xf2, 0x71,
	0x9f, 0x53, 0xe6, 0x26, 0x58, 0x92, 0x3c, 0xc7, 0x3c, 0x02, 0xd5, 0x8b, 0x11, 0xdd, 0xe2, 0x94,
	0x39, 0xa3, 0x73, 0x87, 0x74, 0x53, 0x22, 0x24, 0xac, 0x03, 0x90, 0x64, 0x8f, 0x2e, 0x0d, 0xd6,
	0xb4, 0x0d, 0x6d, 0xfb, 0x9f, 0xb3, 0x94, 0x9f, 0xb4, 0x03, 0xf3, 0x0a, 0xac, 0xce, 0x72, 0x22,
	0xe6, 0x4c, 0x10, 0xd8, 0x02, 0x95, 0x84, 0x88, 0x34, 0x92, 0x0a, 0x5a, 0xb6, 0x77, 0xad, 0xdf,
	0x4b, 0x5a, 0xd3, 0x78, 0x1a, 0x49, 0x27, 0x47, 0xcd, 0x3a, 0x58, 0x57, 0xf1, 0xe7, 0x58, 0xc8,
	0xf1, 0x4c, 0x3b, 0xc8, 0xe5, 0xcc, 0x13, 0x50, 0x2b, 0x7e, 0x9d, 0x3b, 0xcc, 0x97, 0xb7, 0x3f,
	0xcb, 0xe0, 0xbf, 0xe2, 0xe1, 0xb3, 0x06, 0x56, 0x66, 0x1c, 0x60, 0x63, 0x9e, 0x70, 0xe1, 0xb2,
	0x74, 0x7b, 0x11, 0x24, 0x73, 0x34, 0x9b, 0x0f, 0xaf, 0x1f, 0x4f, 0x65, 0x1b, 0x1e, 0xa0, 0x8c,
	0x40, 0x45, 0x7f, 0xc9, 0xcd, 0xf6, 0x81, 0xee, 0x26, 0x75, 0xee, 0x95, 0xf4, 0x4c, 0x73, 0x78,
	0xfc, 0xa7, 0x41, 0xf1, 0x2a, 0xf5, 0xe6, 0xe2, 0x60, 0x5e, 0x60, 0x4f, 0x15, 0xd8, 0x82, 0x9b,
	0x3f, 0x0b, 0x44, 0x58, 0x48, 0x77, 0xaa, 0x05, 0x0d, 0x4e, 0xcf, 0x5e, 0x06, 0x86, 0xd6, 0x1f,
	0x18, 0xda, 0xfb, 0xc0, 0xd0, 0x1e, 0x87, 0x46, 0xa9, 0x3f, 0x34, 0x4a, 0x6f, 0x43, 0xa3, 0x74,
	0xb9, 0x1f, 0x52, 0x79, 0x9d, 0x7a, 0x96, 0xcf, 0x3b, 0x68, 0xec, 0xf2, 0x9d, 0x79, 0x3b, 0x95,
	0x2a, 0x7b, 0x31, 0x11, 0x5e, 0x45, 0x5d, 0xdc, 0xc3, 0xaf, 0x00, 0x00, 0x00, 0xff, 0xff, 0xe1,
	0x5f, 0x69, 0x49, 0x53, 0x03, 0x00, 0x00,
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion4

// QueryClient is the client API for Query service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type QueryClient interface {
	// CoinRatesResult defines a rpc handler method for MsgCoinRatesData.
	CoinRatesResult(ctx context.Context, in *QueryCoinRatesRequest, opts ...grpc.CallOption) (*QueryCoinRatesResponse, error)
	// LastCoinRatesId query the last CoinRates result id
	LastCoinRatesId(ctx context.Context, in *QueryLastCoinRatesIdRequest, opts ...grpc.CallOption) (*QueryLastCoinRatesIdResponse, error)
}

type queryClient struct {
	cc grpc1.ClientConn
}

func NewQueryClient(cc grpc1.ClientConn) QueryClient {
	return &queryClient{cc}
}

func (c *queryClient) CoinRatesResult(ctx context.Context, in *QueryCoinRatesRequest, opts ...grpc.CallOption) (*QueryCoinRatesResponse, error) {
	out := new(QueryCoinRatesResponse)
	err := c.cc.Invoke(ctx, "/cosmonaut.oracle.consuming.Query/CoinRatesResult", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *queryClient) LastCoinRatesId(ctx context.Context, in *QueryLastCoinRatesIdRequest, opts ...grpc.CallOption) (*QueryLastCoinRatesIdResponse, error) {
	out := new(QueryLastCoinRatesIdResponse)
	err := c.cc.Invoke(ctx, "/cosmonaut.oracle.consuming.Query/LastCoinRatesId", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// QueryServer is the server API for Query service.
type QueryServer interface {
	// CoinRatesResult defines a rpc handler method for MsgCoinRatesData.
	CoinRatesResult(context.Context, *QueryCoinRatesRequest) (*QueryCoinRatesResponse, error)
	// LastCoinRatesId query the last CoinRates result id
	LastCoinRatesId(context.Context, *QueryLastCoinRatesIdRequest) (*QueryLastCoinRatesIdResponse, error)
}

// UnimplementedQueryServer can be embedded to have forward compatible implementations.
type UnimplementedQueryServer struct {
}

func (*UnimplementedQueryServer) CoinRatesResult(ctx context.Context, req *QueryCoinRatesRequest) (*QueryCoinRatesResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method CoinRatesResult not implemented")
}
func (*UnimplementedQueryServer) LastCoinRatesId(ctx context.Context, req *QueryLastCoinRatesIdRequest) (*QueryLastCoinRatesIdResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method LastCoinRatesId not implemented")
}

func RegisterQueryServer(s grpc1.Server, srv QueryServer) {
	s.RegisterService(&_Query_serviceDesc, srv)
}

func _Query_CoinRatesResult_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(QueryCoinRatesRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(QueryServer).CoinRatesResult(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/cosmonaut.oracle.consuming.Query/CoinRatesResult",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(QueryServer).CoinRatesResult(ctx, req.(*QueryCoinRatesRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Query_LastCoinRatesId_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(QueryLastCoinRatesIdRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(QueryServer).LastCoinRatesId(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/cosmonaut.oracle.consuming.Query/LastCoinRatesId",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(QueryServer).LastCoinRatesId(ctx, req.(*QueryLastCoinRatesIdRequest))
	}
	return interceptor(ctx, in, info, handler)
}

var _Query_serviceDesc = grpc.ServiceDesc{
	ServiceName: "cosmonaut.oracle.consuming.Query",
	HandlerType: (*QueryServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "CoinRatesResult",
			Handler:    _Query_CoinRatesResult_Handler,
		},
		{
			MethodName: "LastCoinRatesId",
			Handler:    _Query_LastCoinRatesId_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "consuming/query.proto",
}

func (m *QueryCoinRatesRequest) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *QueryCoinRatesRequest) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *QueryCoinRatesRequest) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if m.RequestId != 0 {
		i = encodeVarintQuery(dAtA, i, uint64(m.RequestId))
		i--
		dAtA[i] = 0x8
	}
	return len(dAtA) - i, nil
}

func (m *QueryCoinRatesResponse) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *QueryCoinRatesResponse) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *QueryCoinRatesResponse) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if m.Result != nil {
		{
			size, err := m.Result.MarshalToSizedBuffer(dAtA[:i])
			if err != nil {
				return 0, err
			}
			i -= size
			i = encodeVarintQuery(dAtA, i, uint64(size))
		}
		i--
		dAtA[i] = 0xa
	}
	return len(dAtA) - i, nil
}

func (m *QueryLastCoinRatesIdRequest) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *QueryLastCoinRatesIdRequest) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *QueryLastCoinRatesIdRequest) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	return len(dAtA) - i, nil
}

func (m *QueryLastCoinRatesIdResponse) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *QueryLastCoinRatesIdResponse) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *QueryLastCoinRatesIdResponse) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if m.RequestId != 0 {
		i = encodeVarintQuery(dAtA, i, uint64(m.RequestId))
		i--
		dAtA[i] = 0x8
	}
	return len(dAtA) - i, nil
}

func encodeVarintQuery(dAtA []byte, offset int, v uint64) int {
	offset -= sovQuery(v)
	base := offset
	for v >= 1<<7 {
		dAtA[offset] = uint8(v&0x7f | 0x80)
		v >>= 7
		offset++
	}
	dAtA[offset] = uint8(v)
	return base
}
func (m *QueryCoinRatesRequest) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	if m.RequestId != 0 {
		n += 1 + sovQuery(uint64(m.RequestId))
	}
	return n
}

func (m *QueryCoinRatesResponse) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	if m.Result != nil {
		l = m.Result.Size()
		n += 1 + l + sovQuery(uint64(l))
	}
	return n
}

func (m *QueryLastCoinRatesIdRequest) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	return n
}

func (m *QueryLastCoinRatesIdResponse) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	if m.RequestId != 0 {
		n += 1 + sovQuery(uint64(m.RequestId))
	}
	return n
}

func sovQuery(x uint64) (n int) {
	return (math_bits.Len64(x|1) + 6) / 7
}
func sozQuery(x uint64) (n int) {
	return sovQuery(uint64((x << 1) ^ uint64((int64(x) >> 63))))
}
func (m *QueryCoinRatesRequest) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowQuery
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
			return fmt.Errorf("proto: QueryCoinRatesRequest: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: QueryCoinRatesRequest: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field RequestId", wireType)
			}
			m.RequestId = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowQuery
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.RequestId |= int64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		default:
			iNdEx = preIndex
			skippy, err := skipQuery(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthQuery
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
func (m *QueryCoinRatesResponse) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowQuery
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
			return fmt.Errorf("proto: QueryCoinRatesResponse: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: QueryCoinRatesResponse: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Result", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowQuery
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
				return ErrInvalidLengthQuery
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthQuery
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if m.Result == nil {
				m.Result = &CoinRatesResult{}
			}
			if err := m.Result.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipQuery(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthQuery
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
func (m *QueryLastCoinRatesIdRequest) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowQuery
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
			return fmt.Errorf("proto: QueryLastCoinRatesIdRequest: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: QueryLastCoinRatesIdRequest: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		default:
			iNdEx = preIndex
			skippy, err := skipQuery(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthQuery
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
func (m *QueryLastCoinRatesIdResponse) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowQuery
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
			return fmt.Errorf("proto: QueryLastCoinRatesIdResponse: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: QueryLastCoinRatesIdResponse: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field RequestId", wireType)
			}
			m.RequestId = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowQuery
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.RequestId |= int64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		default:
			iNdEx = preIndex
			skippy, err := skipQuery(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthQuery
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
func skipQuery(dAtA []byte) (n int, err error) {
	l := len(dAtA)
	iNdEx := 0
	depth := 0
	for iNdEx < l {
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return 0, ErrIntOverflowQuery
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
					return 0, ErrIntOverflowQuery
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
					return 0, ErrIntOverflowQuery
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
				return 0, ErrInvalidLengthQuery
			}
			iNdEx += length
		case 3:
			depth++
		case 4:
			if depth == 0 {
				return 0, ErrUnexpectedEndOfGroupQuery
			}
			depth--
		case 5:
			iNdEx += 4
		default:
			return 0, fmt.Errorf("proto: illegal wireType %d", wireType)
		}
		if iNdEx < 0 {
			return 0, ErrInvalidLengthQuery
		}
		if depth == 0 {
			return iNdEx, nil
		}
	}
	return 0, io.ErrUnexpectedEOF
}

var (
	ErrInvalidLengthQuery        = fmt.Errorf("proto: negative length found during unmarshaling")
	ErrIntOverflowQuery          = fmt.Errorf("proto: integer overflow")
	ErrUnexpectedEndOfGroupQuery = fmt.Errorf("proto: unexpected end of group")
)
