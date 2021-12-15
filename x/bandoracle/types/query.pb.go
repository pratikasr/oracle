// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: bandoracle/query.proto

package types

import (
	context "context"
	fmt "fmt"
	_ "github.com/cosmos/cosmos-sdk/types/query"
	_ "github.com/gogo/protobuf/gogoproto"
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

// QueryParamsRequest is request type for the Query/Params RPC method.
type QueryParamsRequest struct {
}

func (m *QueryParamsRequest) Reset()         { *m = QueryParamsRequest{} }
func (m *QueryParamsRequest) String() string { return proto.CompactTextString(m) }
func (*QueryParamsRequest) ProtoMessage()    {}
func (*QueryParamsRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_84eba62bfb351d5a, []int{0}
}
func (m *QueryParamsRequest) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *QueryParamsRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_QueryParamsRequest.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *QueryParamsRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_QueryParamsRequest.Merge(m, src)
}
func (m *QueryParamsRequest) XXX_Size() int {
	return m.Size()
}
func (m *QueryParamsRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_QueryParamsRequest.DiscardUnknown(m)
}

var xxx_messageInfo_QueryParamsRequest proto.InternalMessageInfo

// QueryParamsResponse is response type for the Query/Params RPC method.
type QueryParamsResponse struct {
	// params holds all the parameters of this module.
	Params Params `protobuf:"bytes,1,opt,name=params,proto3" json:"params"`
}

func (m *QueryParamsResponse) Reset()         { *m = QueryParamsResponse{} }
func (m *QueryParamsResponse) String() string { return proto.CompactTextString(m) }
func (*QueryParamsResponse) ProtoMessage()    {}
func (*QueryParamsResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_84eba62bfb351d5a, []int{1}
}
func (m *QueryParamsResponse) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *QueryParamsResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_QueryParamsResponse.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *QueryParamsResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_QueryParamsResponse.Merge(m, src)
}
func (m *QueryParamsResponse) XXX_Size() int {
	return m.Size()
}
func (m *QueryParamsResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_QueryParamsResponse.DiscardUnknown(m)
}

var xxx_messageInfo_QueryParamsResponse proto.InternalMessageInfo

func (m *QueryParamsResponse) GetParams() Params {
	if m != nil {
		return m.Params
	}
	return Params{}
}

type QueryFetchPriceRequest struct {
	RequestId int64 `protobuf:"varint,1,opt,name=request_id,json=requestId,proto3" json:"request_id,omitempty"`
}

func (m *QueryFetchPriceRequest) Reset()         { *m = QueryFetchPriceRequest{} }
func (m *QueryFetchPriceRequest) String() string { return proto.CompactTextString(m) }
func (*QueryFetchPriceRequest) ProtoMessage()    {}
func (*QueryFetchPriceRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_84eba62bfb351d5a, []int{2}
}
func (m *QueryFetchPriceRequest) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *QueryFetchPriceRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_QueryFetchPriceRequest.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *QueryFetchPriceRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_QueryFetchPriceRequest.Merge(m, src)
}
func (m *QueryFetchPriceRequest) XXX_Size() int {
	return m.Size()
}
func (m *QueryFetchPriceRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_QueryFetchPriceRequest.DiscardUnknown(m)
}

var xxx_messageInfo_QueryFetchPriceRequest proto.InternalMessageInfo

func (m *QueryFetchPriceRequest) GetRequestId() int64 {
	if m != nil {
		return m.RequestId
	}
	return 0
}

type QueryFetchPriceResponse struct {
	Result *FetchPriceResult `protobuf:"bytes,1,opt,name=result,proto3" json:"result,omitempty"`
}

func (m *QueryFetchPriceResponse) Reset()         { *m = QueryFetchPriceResponse{} }
func (m *QueryFetchPriceResponse) String() string { return proto.CompactTextString(m) }
func (*QueryFetchPriceResponse) ProtoMessage()    {}
func (*QueryFetchPriceResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_84eba62bfb351d5a, []int{3}
}
func (m *QueryFetchPriceResponse) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *QueryFetchPriceResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_QueryFetchPriceResponse.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *QueryFetchPriceResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_QueryFetchPriceResponse.Merge(m, src)
}
func (m *QueryFetchPriceResponse) XXX_Size() int {
	return m.Size()
}
func (m *QueryFetchPriceResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_QueryFetchPriceResponse.DiscardUnknown(m)
}

var xxx_messageInfo_QueryFetchPriceResponse proto.InternalMessageInfo

func (m *QueryFetchPriceResponse) GetResult() *FetchPriceResult {
	if m != nil {
		return m.Result
	}
	return nil
}

type QueryLastFetchPriceIdRequest struct {
}

func (m *QueryLastFetchPriceIdRequest) Reset()         { *m = QueryLastFetchPriceIdRequest{} }
func (m *QueryLastFetchPriceIdRequest) String() string { return proto.CompactTextString(m) }
func (*QueryLastFetchPriceIdRequest) ProtoMessage()    {}
func (*QueryLastFetchPriceIdRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_84eba62bfb351d5a, []int{4}
}
func (m *QueryLastFetchPriceIdRequest) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *QueryLastFetchPriceIdRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_QueryLastFetchPriceIdRequest.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *QueryLastFetchPriceIdRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_QueryLastFetchPriceIdRequest.Merge(m, src)
}
func (m *QueryLastFetchPriceIdRequest) XXX_Size() int {
	return m.Size()
}
func (m *QueryLastFetchPriceIdRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_QueryLastFetchPriceIdRequest.DiscardUnknown(m)
}

var xxx_messageInfo_QueryLastFetchPriceIdRequest proto.InternalMessageInfo

type QueryLastFetchPriceIdResponse struct {
	RequestId int64 `protobuf:"varint,1,opt,name=request_id,json=requestId,proto3" json:"request_id,omitempty"`
}

func (m *QueryLastFetchPriceIdResponse) Reset()         { *m = QueryLastFetchPriceIdResponse{} }
func (m *QueryLastFetchPriceIdResponse) String() string { return proto.CompactTextString(m) }
func (*QueryLastFetchPriceIdResponse) ProtoMessage()    {}
func (*QueryLastFetchPriceIdResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_84eba62bfb351d5a, []int{5}
}
func (m *QueryLastFetchPriceIdResponse) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *QueryLastFetchPriceIdResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_QueryLastFetchPriceIdResponse.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *QueryLastFetchPriceIdResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_QueryLastFetchPriceIdResponse.Merge(m, src)
}
func (m *QueryLastFetchPriceIdResponse) XXX_Size() int {
	return m.Size()
}
func (m *QueryLastFetchPriceIdResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_QueryLastFetchPriceIdResponse.DiscardUnknown(m)
}

var xxx_messageInfo_QueryLastFetchPriceIdResponse proto.InternalMessageInfo

func (m *QueryLastFetchPriceIdResponse) GetRequestId() int64 {
	if m != nil {
		return m.RequestId
	}
	return 0
}

func init() {
	proto.RegisterType((*QueryParamsRequest)(nil), "cosmonaut.oracle.bandoracle.QueryParamsRequest")
	proto.RegisterType((*QueryParamsResponse)(nil), "cosmonaut.oracle.bandoracle.QueryParamsResponse")
	proto.RegisterType((*QueryFetchPriceRequest)(nil), "cosmonaut.oracle.bandoracle.QueryFetchPriceRequest")
	proto.RegisterType((*QueryFetchPriceResponse)(nil), "cosmonaut.oracle.bandoracle.QueryFetchPriceResponse")
	proto.RegisterType((*QueryLastFetchPriceIdRequest)(nil), "cosmonaut.oracle.bandoracle.QueryLastFetchPriceIdRequest")
	proto.RegisterType((*QueryLastFetchPriceIdResponse)(nil), "cosmonaut.oracle.bandoracle.QueryLastFetchPriceIdResponse")
}

func init() { proto.RegisterFile("bandoracle/query.proto", fileDescriptor_84eba62bfb351d5a) }

var fileDescriptor_84eba62bfb351d5a = []byte{
	// 483 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x9c, 0x93, 0x4f, 0x8b, 0xd3, 0x40,
	0x18, 0xc6, 0x13, 0x5d, 0x03, 0x8e, 0x17, 0x19, 0x97, 0x5d, 0x89, 0xdd, 0x28, 0x59, 0x94, 0x45,
	0x71, 0xc6, 0xed, 0x2e, 0x88, 0x7b, 0x10, 0x5c, 0x50, 0x5c, 0xf0, 0xb0, 0xe6, 0x24, 0x5e, 0xe2,
	0x24, 0x19, 0xb3, 0x81, 0x34, 0x93, 0x66, 0x26, 0x62, 0x11, 0x2f, 0x7e, 0x02, 0xc1, 0x8b, 0xf8,
	0x4d, 0xfc, 0x06, 0x3d, 0x16, 0xbc, 0x78, 0x12, 0x69, 0xfd, 0x20, 0x92, 0x99, 0xa9, 0x4d, 0x9b,
	0x36, 0x55, 0x6f, 0x61, 0xde, 0xf7, 0x79, 0xde, 0xdf, 0xfb, 0x27, 0x60, 0x2b, 0x20, 0x59, 0xc4,
	0x0a, 0x12, 0xa6, 0x14, 0xf7, 0x4b, 0x5a, 0x0c, 0x50, 0x5e, 0x30, 0xc1, 0xe0, 0xb5, 0x90, 0xf1,
	0x1e, 0xcb, 0x48, 0x29, 0x90, 0x8a, 0xa2, 0x59, 0xa2, 0xbd, 0x19, 0xb3, 0x98, 0xc9, 0x3c, 0x5c,
	0x7d, 0x29, 0x89, 0xdd, 0x89, 0x19, 0x8b, 0x53, 0x8a, 0x49, 0x9e, 0x60, 0x92, 0x65, 0x4c, 0x10,
	0x91, 0xb0, 0x8c, 0xeb, 0xe8, 0x6d, 0x69, 0xc8, 0x71, 0x40, 0xb8, 0xae, 0x84, 0xdf, 0xec, 0x07,
	0x54, 0x90, 0x7d, 0x9c, 0x93, 0x38, 0xc9, 0x64, 0xb2, 0xce, 0xdd, 0xae, 0x41, 0xe5, 0xa4, 0x20,
	0xbd, 0xa9, 0x49, 0xa7, 0x16, 0x78, 0x4d, 0x45, 0x78, 0xe6, 0xe7, 0x45, 0x12, 0x52, 0x15, 0x75,
	0x37, 0x01, 0x7c, 0x5e, 0x19, 0x9f, 0x4a, 0x89, 0x47, 0xfb, 0x25, 0xe5, 0xc2, 0x7d, 0x01, 0xae,
	0xcc, 0xbd, 0xf2, 0x9c, 0x65, 0x9c, 0xc2, 0x47, 0xc0, 0x52, 0xd6, 0x57, 0xcd, 0x1b, 0xe6, 0xde,
	0xa5, 0xee, 0x2e, 0x6a, 0xe9, 0x18, 0x29, 0xf1, 0xf1, 0xc6, 0xf0, 0xc7, 0x75, 0xc3, 0xd3, 0x42,
	0xf7, 0x3e, 0xd8, 0x92, 0xce, 0x4f, 0x2a, 0x92, 0xd3, 0x0a, 0x44, 0xd7, 0x84, 0x3b, 0x00, 0x14,
	0xea, 0xd3, 0x4f, 0x22, 0x59, 0xe0, 0xbc, 0x77, 0x51, 0xbf, 0x9c, 0x44, 0xee, 0x2b, 0xb0, 0xdd,
	0x10, 0x6a, 0xac, 0xc7, 0xc0, 0x2a, 0x28, 0x2f, 0x53, 0xa1, 0xb1, 0xee, 0xb6, 0x62, 0xcd, 0x19,
	0x94, 0xa9, 0xf0, 0xb4, 0xd8, 0x75, 0x40, 0x47, 0x56, 0x78, 0x46, 0xb8, 0x98, 0x25, 0x9d, 0x44,
	0xd3, 0xa1, 0x3c, 0x04, 0x3b, 0x2b, 0xe2, 0x9a, 0xa3, 0xbd, 0x83, 0xee, 0x97, 0x0d, 0x70, 0x41,
	0x1a, 0xc0, 0xcf, 0x26, 0xb0, 0xd4, 0x74, 0x20, 0x6e, 0x65, 0x6d, 0xae, 0xc6, 0xbe, 0xf7, 0xf7,
	0x02, 0x85, 0xe5, 0xde, 0xf9, 0xf0, 0xed, 0xd7, 0xa7, 0x73, 0x37, 0xe1, 0x2e, 0xfe, 0xa3, 0xc4,
	0xfa, 0x1e, 0x1a, 0x37, 0x03, 0xbf, 0x9a, 0xe0, 0xf2, 0xe2, 0x84, 0xe0, 0xc1, 0xfa, 0x9a, 0x8d,
	0x7d, 0xda, 0x87, 0xff, 0x26, 0xd2, 0xb0, 0x47, 0x12, 0xf6, 0x10, 0x76, 0x97, 0x20, 0xd6, 0xae,
	0xd7, 0x57, 0x3b, 0xc3, 0xef, 0x66, 0x03, 0x7f, 0x2f, 0xd9, 0x17, 0x97, 0x03, 0x1f, 0xac, 0xc7,
	0x58, 0xb1, 0x70, 0xfb, 0xe8, 0x7f, 0xa4, 0xba, 0x0f, 0x24, 0xfb, 0xd8, 0x83, 0xb7, 0x96, 0xf4,
	0x91, 0x12, 0x2e, 0xfc, 0x7a, 0x33, 0x49, 0x74, 0xfc, 0x74, 0x38, 0x76, 0xcc, 0xd1, 0xd8, 0x31,
	0x7f, 0x8e, 0x1d, 0xf3, 0xe3, 0xc4, 0x31, 0x46, 0x13, 0xc7, 0xf8, 0x3e, 0x71, 0x8c, 0x97, 0x28,
	0x4e, 0xc4, 0x59, 0x19, 0xa0, 0x90, 0xf5, 0x9a, 0x0b, 0x7c, 0x5b, 0xf7, 0x15, 0x83, 0x9c, 0xf2,
	0xc0, 0x92, 0x3f, 0xf6, 0xc1, 0xef, 0x00, 0x00, 0x00, 0xff, 0xff, 0x1d, 0xdd, 0x87, 0xd7, 0xa6,
	0x04, 0x00, 0x00,
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
	// Parameters queries the parameters of the module.
	Params(ctx context.Context, in *QueryParamsRequest, opts ...grpc.CallOption) (*QueryParamsResponse, error)
	// FetchPriceResult defines a rpc handler method for MsgFetchPriceData.
	FetchPriceResult(ctx context.Context, in *QueryFetchPriceRequest, opts ...grpc.CallOption) (*QueryFetchPriceResponse, error)
	// LastFetchPriceId query the last FetchPrice result id
	LastFetchPriceId(ctx context.Context, in *QueryLastFetchPriceIdRequest, opts ...grpc.CallOption) (*QueryLastFetchPriceIdResponse, error)
}

type queryClient struct {
	cc grpc1.ClientConn
}

func NewQueryClient(cc grpc1.ClientConn) QueryClient {
	return &queryClient{cc}
}

func (c *queryClient) Params(ctx context.Context, in *QueryParamsRequest, opts ...grpc.CallOption) (*QueryParamsResponse, error) {
	out := new(QueryParamsResponse)
	err := c.cc.Invoke(ctx, "/cosmonaut.oracle.bandoracle.Query/Params", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *queryClient) FetchPriceResult(ctx context.Context, in *QueryFetchPriceRequest, opts ...grpc.CallOption) (*QueryFetchPriceResponse, error) {
	out := new(QueryFetchPriceResponse)
	err := c.cc.Invoke(ctx, "/cosmonaut.oracle.bandoracle.Query/FetchPriceResult", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *queryClient) LastFetchPriceId(ctx context.Context, in *QueryLastFetchPriceIdRequest, opts ...grpc.CallOption) (*QueryLastFetchPriceIdResponse, error) {
	out := new(QueryLastFetchPriceIdResponse)
	err := c.cc.Invoke(ctx, "/cosmonaut.oracle.bandoracle.Query/LastFetchPriceId", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// QueryServer is the server API for Query service.
type QueryServer interface {
	// Parameters queries the parameters of the module.
	Params(context.Context, *QueryParamsRequest) (*QueryParamsResponse, error)
	// FetchPriceResult defines a rpc handler method for MsgFetchPriceData.
	FetchPriceResult(context.Context, *QueryFetchPriceRequest) (*QueryFetchPriceResponse, error)
	// LastFetchPriceId query the last FetchPrice result id
	LastFetchPriceId(context.Context, *QueryLastFetchPriceIdRequest) (*QueryLastFetchPriceIdResponse, error)
}

// UnimplementedQueryServer can be embedded to have forward compatible implementations.
type UnimplementedQueryServer struct {
}

func (*UnimplementedQueryServer) Params(ctx context.Context, req *QueryParamsRequest) (*QueryParamsResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Params not implemented")
}
func (*UnimplementedQueryServer) FetchPriceResult(ctx context.Context, req *QueryFetchPriceRequest) (*QueryFetchPriceResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method FetchPriceResult not implemented")
}
func (*UnimplementedQueryServer) LastFetchPriceId(ctx context.Context, req *QueryLastFetchPriceIdRequest) (*QueryLastFetchPriceIdResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method LastFetchPriceId not implemented")
}

func RegisterQueryServer(s grpc1.Server, srv QueryServer) {
	s.RegisterService(&_Query_serviceDesc, srv)
}

func _Query_Params_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(QueryParamsRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(QueryServer).Params(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/cosmonaut.oracle.bandoracle.Query/Params",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(QueryServer).Params(ctx, req.(*QueryParamsRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Query_FetchPriceResult_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(QueryFetchPriceRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(QueryServer).FetchPriceResult(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/cosmonaut.oracle.bandoracle.Query/FetchPriceResult",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(QueryServer).FetchPriceResult(ctx, req.(*QueryFetchPriceRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Query_LastFetchPriceId_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(QueryLastFetchPriceIdRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(QueryServer).LastFetchPriceId(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/cosmonaut.oracle.bandoracle.Query/LastFetchPriceId",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(QueryServer).LastFetchPriceId(ctx, req.(*QueryLastFetchPriceIdRequest))
	}
	return interceptor(ctx, in, info, handler)
}

var _Query_serviceDesc = grpc.ServiceDesc{
	ServiceName: "cosmonaut.oracle.bandoracle.Query",
	HandlerType: (*QueryServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "Params",
			Handler:    _Query_Params_Handler,
		},
		{
			MethodName: "FetchPriceResult",
			Handler:    _Query_FetchPriceResult_Handler,
		},
		{
			MethodName: "LastFetchPriceId",
			Handler:    _Query_LastFetchPriceId_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "bandoracle/query.proto",
}

func (m *QueryParamsRequest) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *QueryParamsRequest) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *QueryParamsRequest) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	return len(dAtA) - i, nil
}

func (m *QueryParamsResponse) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *QueryParamsResponse) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *QueryParamsResponse) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	{
		size, err := m.Params.MarshalToSizedBuffer(dAtA[:i])
		if err != nil {
			return 0, err
		}
		i -= size
		i = encodeVarintQuery(dAtA, i, uint64(size))
	}
	i--
	dAtA[i] = 0xa
	return len(dAtA) - i, nil
}

func (m *QueryFetchPriceRequest) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *QueryFetchPriceRequest) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *QueryFetchPriceRequest) MarshalToSizedBuffer(dAtA []byte) (int, error) {
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

func (m *QueryFetchPriceResponse) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *QueryFetchPriceResponse) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *QueryFetchPriceResponse) MarshalToSizedBuffer(dAtA []byte) (int, error) {
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

func (m *QueryLastFetchPriceIdRequest) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *QueryLastFetchPriceIdRequest) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *QueryLastFetchPriceIdRequest) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	return len(dAtA) - i, nil
}

func (m *QueryLastFetchPriceIdResponse) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *QueryLastFetchPriceIdResponse) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *QueryLastFetchPriceIdResponse) MarshalToSizedBuffer(dAtA []byte) (int, error) {
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
func (m *QueryParamsRequest) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	return n
}

func (m *QueryParamsResponse) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	l = m.Params.Size()
	n += 1 + l + sovQuery(uint64(l))
	return n
}

func (m *QueryFetchPriceRequest) Size() (n int) {
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

func (m *QueryFetchPriceResponse) Size() (n int) {
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

func (m *QueryLastFetchPriceIdRequest) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	return n
}

func (m *QueryLastFetchPriceIdResponse) Size() (n int) {
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
func (m *QueryParamsRequest) Unmarshal(dAtA []byte) error {
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
			return fmt.Errorf("proto: QueryParamsRequest: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: QueryParamsRequest: illegal tag %d (wire type %d)", fieldNum, wire)
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
func (m *QueryParamsResponse) Unmarshal(dAtA []byte) error {
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
			return fmt.Errorf("proto: QueryParamsResponse: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: QueryParamsResponse: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Params", wireType)
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
			if err := m.Params.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
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
func (m *QueryFetchPriceRequest) Unmarshal(dAtA []byte) error {
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
			return fmt.Errorf("proto: QueryFetchPriceRequest: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: QueryFetchPriceRequest: illegal tag %d (wire type %d)", fieldNum, wire)
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
func (m *QueryFetchPriceResponse) Unmarshal(dAtA []byte) error {
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
			return fmt.Errorf("proto: QueryFetchPriceResponse: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: QueryFetchPriceResponse: illegal tag %d (wire type %d)", fieldNum, wire)
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
				m.Result = &FetchPriceResult{}
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
func (m *QueryLastFetchPriceIdRequest) Unmarshal(dAtA []byte) error {
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
			return fmt.Errorf("proto: QueryLastFetchPriceIdRequest: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: QueryLastFetchPriceIdRequest: illegal tag %d (wire type %d)", fieldNum, wire)
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
func (m *QueryLastFetchPriceIdResponse) Unmarshal(dAtA []byte) error {
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
			return fmt.Errorf("proto: QueryLastFetchPriceIdResponse: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: QueryLastFetchPriceIdResponse: illegal tag %d (wire type %d)", fieldNum, wire)
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