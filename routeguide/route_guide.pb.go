// Code generated by protoc-gen-go. DO NOT EDIT.
// source: route_guide.proto

package routeguide

import proto "github.com/golang/protobuf/proto"
import fmt "fmt"
import math "math"

import (
	context "golang.org/x/net/context"
	grpc "google.golang.org/grpc"
)

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.ProtoPackageIsVersion2 // please upgrade the proto package

// The request message containing the user's name.
type ShardStore struct {
	Hash                 string   `protobuf:"bytes,1,opt,name=hash" json:"hash,omitempty"`
	Size                 int64    `protobuf:"varint,2,opt,name=size" json:"size,omitempty"`
	Ttl                  int64    `protobuf:"varint,3,opt,name=ttl" json:"ttl,omitempty"`
	StoreOffset          int64    `protobuf:"varint,4,opt,name=storeOffset" json:"storeOffset,omitempty"`
	Content              []byte   `protobuf:"bytes,5,opt,name=content,proto3" json:"content,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *ShardStore) Reset()         { *m = ShardStore{} }
func (m *ShardStore) String() string { return proto.CompactTextString(m) }
func (*ShardStore) ProtoMessage()    {}
func (*ShardStore) Descriptor() ([]byte, []int) {
	return fileDescriptor_route_guide_9910fe60dbd44fb3, []int{0}
}
func (m *ShardStore) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ShardStore.Unmarshal(m, b)
}
func (m *ShardStore) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ShardStore.Marshal(b, m, deterministic)
}
func (dst *ShardStore) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ShardStore.Merge(dst, src)
}
func (m *ShardStore) XXX_Size() int {
	return xxx_messageInfo_ShardStore.Size(m)
}
func (m *ShardStore) XXX_DiscardUnknown() {
	xxx_messageInfo_ShardStore.DiscardUnknown(m)
}

var xxx_messageInfo_ShardStore proto.InternalMessageInfo

func (m *ShardStore) GetHash() string {
	if m != nil {
		return m.Hash
	}
	return ""
}

func (m *ShardStore) GetSize() int64 {
	if m != nil {
		return m.Size
	}
	return 0
}

func (m *ShardStore) GetTtl() int64 {
	if m != nil {
		return m.Ttl
	}
	return 0
}

func (m *ShardStore) GetStoreOffset() int64 {
	if m != nil {
		return m.StoreOffset
	}
	return 0
}

func (m *ShardStore) GetContent() []byte {
	if m != nil {
		return m.Content
	}
	return nil
}

type ShardRetrieval struct {
	Hash                 string   `protobuf:"bytes,1,opt,name=hash" json:"hash,omitempty"`
	Size                 int64    `protobuf:"varint,2,opt,name=size" json:"size,omitempty"`
	StoreOffset          int64    `protobuf:"varint,3,opt,name=storeOffset" json:"storeOffset,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *ShardRetrieval) Reset()         { *m = ShardRetrieval{} }
func (m *ShardRetrieval) String() string { return proto.CompactTextString(m) }
func (*ShardRetrieval) ProtoMessage()    {}
func (*ShardRetrieval) Descriptor() ([]byte, []int) {
	return fileDescriptor_route_guide_9910fe60dbd44fb3, []int{1}
}
func (m *ShardRetrieval) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ShardRetrieval.Unmarshal(m, b)
}
func (m *ShardRetrieval) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ShardRetrieval.Marshal(b, m, deterministic)
}
func (dst *ShardRetrieval) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ShardRetrieval.Merge(dst, src)
}
func (m *ShardRetrieval) XXX_Size() int {
	return xxx_messageInfo_ShardRetrieval.Size(m)
}
func (m *ShardRetrieval) XXX_DiscardUnknown() {
	xxx_messageInfo_ShardRetrieval.DiscardUnknown(m)
}

var xxx_messageInfo_ShardRetrieval proto.InternalMessageInfo

func (m *ShardRetrieval) GetHash() string {
	if m != nil {
		return m.Hash
	}
	return ""
}

func (m *ShardRetrieval) GetSize() int64 {
	if m != nil {
		return m.Size
	}
	return 0
}

func (m *ShardRetrieval) GetStoreOffset() int64 {
	if m != nil {
		return m.StoreOffset
	}
	return 0
}

type ShardRetrievalStream struct {
	Size                 int64    `protobuf:"varint,1,opt,name=size" json:"size,omitempty"`
	Content              []byte   `protobuf:"bytes,2,opt,name=content,proto3" json:"content,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *ShardRetrievalStream) Reset()         { *m = ShardRetrievalStream{} }
func (m *ShardRetrievalStream) String() string { return proto.CompactTextString(m) }
func (*ShardRetrievalStream) ProtoMessage()    {}
func (*ShardRetrievalStream) Descriptor() ([]byte, []int) {
	return fileDescriptor_route_guide_9910fe60dbd44fb3, []int{2}
}
func (m *ShardRetrievalStream) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ShardRetrievalStream.Unmarshal(m, b)
}
func (m *ShardRetrievalStream) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ShardRetrievalStream.Marshal(b, m, deterministic)
}
func (dst *ShardRetrievalStream) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ShardRetrievalStream.Merge(dst, src)
}
func (m *ShardRetrievalStream) XXX_Size() int {
	return xxx_messageInfo_ShardRetrievalStream.Size(m)
}
func (m *ShardRetrievalStream) XXX_DiscardUnknown() {
	xxx_messageInfo_ShardRetrievalStream.DiscardUnknown(m)
}

var xxx_messageInfo_ShardRetrievalStream proto.InternalMessageInfo

func (m *ShardRetrievalStream) GetSize() int64 {
	if m != nil {
		return m.Size
	}
	return 0
}

func (m *ShardRetrievalStream) GetContent() []byte {
	if m != nil {
		return m.Content
	}
	return nil
}

type ShardDelete struct {
	Hash                 string   `protobuf:"bytes,1,opt,name=hash" json:"hash,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *ShardDelete) Reset()         { *m = ShardDelete{} }
func (m *ShardDelete) String() string { return proto.CompactTextString(m) }
func (*ShardDelete) ProtoMessage()    {}
func (*ShardDelete) Descriptor() ([]byte, []int) {
	return fileDescriptor_route_guide_9910fe60dbd44fb3, []int{3}
}
func (m *ShardDelete) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ShardDelete.Unmarshal(m, b)
}
func (m *ShardDelete) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ShardDelete.Marshal(b, m, deterministic)
}
func (dst *ShardDelete) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ShardDelete.Merge(dst, src)
}
func (m *ShardDelete) XXX_Size() int {
	return xxx_messageInfo_ShardDelete.Size(m)
}
func (m *ShardDelete) XXX_DiscardUnknown() {
	xxx_messageInfo_ShardDelete.DiscardUnknown(m)
}

var xxx_messageInfo_ShardDelete proto.InternalMessageInfo

func (m *ShardDelete) GetHash() string {
	if m != nil {
		return m.Hash
	}
	return ""
}

type ShardDeleteSummary struct {
	Status               int64    `protobuf:"varint,1,opt,name=status" json:"status,omitempty"`
	Message              string   `protobuf:"bytes,2,opt,name=message" json:"message,omitempty"`
	ElapsedTime          int64    `protobuf:"varint,3,opt,name=elapsedTime" json:"elapsedTime,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *ShardDeleteSummary) Reset()         { *m = ShardDeleteSummary{} }
func (m *ShardDeleteSummary) String() string { return proto.CompactTextString(m) }
func (*ShardDeleteSummary) ProtoMessage()    {}
func (*ShardDeleteSummary) Descriptor() ([]byte, []int) {
	return fileDescriptor_route_guide_9910fe60dbd44fb3, []int{4}
}
func (m *ShardDeleteSummary) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ShardDeleteSummary.Unmarshal(m, b)
}
func (m *ShardDeleteSummary) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ShardDeleteSummary.Marshal(b, m, deterministic)
}
func (dst *ShardDeleteSummary) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ShardDeleteSummary.Merge(dst, src)
}
func (m *ShardDeleteSummary) XXX_Size() int {
	return xxx_messageInfo_ShardDeleteSummary.Size(m)
}
func (m *ShardDeleteSummary) XXX_DiscardUnknown() {
	xxx_messageInfo_ShardDeleteSummary.DiscardUnknown(m)
}

var xxx_messageInfo_ShardDeleteSummary proto.InternalMessageInfo

func (m *ShardDeleteSummary) GetStatus() int64 {
	if m != nil {
		return m.Status
	}
	return 0
}

func (m *ShardDeleteSummary) GetMessage() string {
	if m != nil {
		return m.Message
	}
	return ""
}

func (m *ShardDeleteSummary) GetElapsedTime() int64 {
	if m != nil {
		return m.ElapsedTime
	}
	return 0
}

type ShardStoreSummary struct {
	Status               int64    `protobuf:"varint,1,opt,name=status" json:"status,omitempty"`
	Message              string   `protobuf:"bytes,2,opt,name=message" json:"message,omitempty"`
	TotalReceived        int64    `protobuf:"varint,3,opt,name=totalReceived" json:"totalReceived,omitempty"`
	ElapsedTime          int64    `protobuf:"varint,4,opt,name=elapsedTime" json:"elapsedTime,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *ShardStoreSummary) Reset()         { *m = ShardStoreSummary{} }
func (m *ShardStoreSummary) String() string { return proto.CompactTextString(m) }
func (*ShardStoreSummary) ProtoMessage()    {}
func (*ShardStoreSummary) Descriptor() ([]byte, []int) {
	return fileDescriptor_route_guide_9910fe60dbd44fb3, []int{5}
}
func (m *ShardStoreSummary) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ShardStoreSummary.Unmarshal(m, b)
}
func (m *ShardStoreSummary) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ShardStoreSummary.Marshal(b, m, deterministic)
}
func (dst *ShardStoreSummary) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ShardStoreSummary.Merge(dst, src)
}
func (m *ShardStoreSummary) XXX_Size() int {
	return xxx_messageInfo_ShardStoreSummary.Size(m)
}
func (m *ShardStoreSummary) XXX_DiscardUnknown() {
	xxx_messageInfo_ShardStoreSummary.DiscardUnknown(m)
}

var xxx_messageInfo_ShardStoreSummary proto.InternalMessageInfo

func (m *ShardStoreSummary) GetStatus() int64 {
	if m != nil {
		return m.Status
	}
	return 0
}

func (m *ShardStoreSummary) GetMessage() string {
	if m != nil {
		return m.Message
	}
	return ""
}

func (m *ShardStoreSummary) GetTotalReceived() int64 {
	if m != nil {
		return m.TotalReceived
	}
	return 0
}

func (m *ShardStoreSummary) GetElapsedTime() int64 {
	if m != nil {
		return m.ElapsedTime
	}
	return 0
}

func init() {
	proto.RegisterType((*ShardStore)(nil), "routeguide.ShardStore")
	proto.RegisterType((*ShardRetrieval)(nil), "routeguide.ShardRetrieval")
	proto.RegisterType((*ShardRetrievalStream)(nil), "routeguide.ShardRetrievalStream")
	proto.RegisterType((*ShardDelete)(nil), "routeguide.ShardDelete")
	proto.RegisterType((*ShardDeleteSummary)(nil), "routeguide.ShardDeleteSummary")
	proto.RegisterType((*ShardStoreSummary)(nil), "routeguide.ShardStoreSummary")
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion4

// Client API for RouteGuide service

type RouteGuideClient interface {
	Retrieve(ctx context.Context, in *ShardRetrieval, opts ...grpc.CallOption) (RouteGuide_RetrieveClient, error)
	Store(ctx context.Context, opts ...grpc.CallOption) (RouteGuide_StoreClient, error)
	Delete(ctx context.Context, in *ShardDelete, opts ...grpc.CallOption) (*ShardDeleteSummary, error)
}

type routeGuideClient struct {
	cc *grpc.ClientConn
}

func NewRouteGuideClient(cc *grpc.ClientConn) RouteGuideClient {
	return &routeGuideClient{cc}
}

func (c *routeGuideClient) Retrieve(ctx context.Context, in *ShardRetrieval, opts ...grpc.CallOption) (RouteGuide_RetrieveClient, error) {
	stream, err := grpc.NewClientStream(ctx, &_RouteGuide_serviceDesc.Streams[0], c.cc, "/routeguide.RouteGuide/Retrieve", opts...)
	if err != nil {
		return nil, err
	}
	x := &routeGuideRetrieveClient{stream}
	if err := x.ClientStream.SendMsg(in); err != nil {
		return nil, err
	}
	if err := x.ClientStream.CloseSend(); err != nil {
		return nil, err
	}
	return x, nil
}

type RouteGuide_RetrieveClient interface {
	Recv() (*ShardRetrievalStream, error)
	grpc.ClientStream
}

type routeGuideRetrieveClient struct {
	grpc.ClientStream
}

func (x *routeGuideRetrieveClient) Recv() (*ShardRetrievalStream, error) {
	m := new(ShardRetrievalStream)
	if err := x.ClientStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

func (c *routeGuideClient) Store(ctx context.Context, opts ...grpc.CallOption) (RouteGuide_StoreClient, error) {
	stream, err := grpc.NewClientStream(ctx, &_RouteGuide_serviceDesc.Streams[1], c.cc, "/routeguide.RouteGuide/Store", opts...)
	if err != nil {
		return nil, err
	}
	x := &routeGuideStoreClient{stream}
	return x, nil
}

type RouteGuide_StoreClient interface {
	Send(*ShardStore) error
	CloseAndRecv() (*ShardStoreSummary, error)
	grpc.ClientStream
}

type routeGuideStoreClient struct {
	grpc.ClientStream
}

func (x *routeGuideStoreClient) Send(m *ShardStore) error {
	return x.ClientStream.SendMsg(m)
}

func (x *routeGuideStoreClient) CloseAndRecv() (*ShardStoreSummary, error) {
	if err := x.ClientStream.CloseSend(); err != nil {
		return nil, err
	}
	m := new(ShardStoreSummary)
	if err := x.ClientStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

func (c *routeGuideClient) Delete(ctx context.Context, in *ShardDelete, opts ...grpc.CallOption) (*ShardDeleteSummary, error) {
	out := new(ShardDeleteSummary)
	err := grpc.Invoke(ctx, "/routeguide.RouteGuide/Delete", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// Server API for RouteGuide service

type RouteGuideServer interface {
	Retrieve(*ShardRetrieval, RouteGuide_RetrieveServer) error
	Store(RouteGuide_StoreServer) error
	Delete(context.Context, *ShardDelete) (*ShardDeleteSummary, error)
}

func RegisterRouteGuideServer(s *grpc.Server, srv RouteGuideServer) {
	s.RegisterService(&_RouteGuide_serviceDesc, srv)
}

func _RouteGuide_Retrieve_Handler(srv interface{}, stream grpc.ServerStream) error {
	m := new(ShardRetrieval)
	if err := stream.RecvMsg(m); err != nil {
		return err
	}
	return srv.(RouteGuideServer).Retrieve(m, &routeGuideRetrieveServer{stream})
}

type RouteGuide_RetrieveServer interface {
	Send(*ShardRetrievalStream) error
	grpc.ServerStream
}

type routeGuideRetrieveServer struct {
	grpc.ServerStream
}

func (x *routeGuideRetrieveServer) Send(m *ShardRetrievalStream) error {
	return x.ServerStream.SendMsg(m)
}

func _RouteGuide_Store_Handler(srv interface{}, stream grpc.ServerStream) error {
	return srv.(RouteGuideServer).Store(&routeGuideStoreServer{stream})
}

type RouteGuide_StoreServer interface {
	SendAndClose(*ShardStoreSummary) error
	Recv() (*ShardStore, error)
	grpc.ServerStream
}

type routeGuideStoreServer struct {
	grpc.ServerStream
}

func (x *routeGuideStoreServer) SendAndClose(m *ShardStoreSummary) error {
	return x.ServerStream.SendMsg(m)
}

func (x *routeGuideStoreServer) Recv() (*ShardStore, error) {
	m := new(ShardStore)
	if err := x.ServerStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

func _RouteGuide_Delete_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ShardDelete)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(RouteGuideServer).Delete(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/routeguide.RouteGuide/Delete",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(RouteGuideServer).Delete(ctx, req.(*ShardDelete))
	}
	return interceptor(ctx, in, info, handler)
}

var _RouteGuide_serviceDesc = grpc.ServiceDesc{
	ServiceName: "routeguide.RouteGuide",
	HandlerType: (*RouteGuideServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "Delete",
			Handler:    _RouteGuide_Delete_Handler,
		},
	},
	Streams: []grpc.StreamDesc{
		{
			StreamName:    "Retrieve",
			Handler:       _RouteGuide_Retrieve_Handler,
			ServerStreams: true,
		},
		{
			StreamName:    "Store",
			Handler:       _RouteGuide_Store_Handler,
			ClientStreams: true,
		},
	},
	Metadata: "route_guide.proto",
}

func init() { proto.RegisterFile("route_guide.proto", fileDescriptor_route_guide_9910fe60dbd44fb3) }

var fileDescriptor_route_guide_9910fe60dbd44fb3 = []byte{
	// 365 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x9c, 0x53, 0x4d, 0x4b, 0xeb, 0x40,
	0x14, 0xed, 0x34, 0x6d, 0xdf, 0xeb, 0xed, 0x7b, 0x62, 0x07, 0xa9, 0x21, 0xa0, 0xc4, 0xe0, 0x22,
	0xab, 0x22, 0xfa, 0x0f, 0xb4, 0xe0, 0x46, 0x10, 0x12, 0x57, 0x6e, 0x64, 0x6c, 0x6e, 0x9b, 0x40,
	0xd2, 0x94, 0xcc, 0x4d, 0x41, 0x97, 0xee, 0xfd, 0x95, 0xfe, 0x11, 0x99, 0x49, 0xda, 0x26, 0x4d,
	0x0b, 0xe2, 0x6e, 0xee, 0xb9, 0x97, 0x73, 0xce, 0xfd, 0x18, 0x18, 0x66, 0x69, 0x4e, 0xf8, 0x32,
	0xcf, 0xa3, 0x00, 0xc7, 0xcb, 0x2c, 0xa5, 0x94, 0x83, 0x86, 0x34, 0xe2, 0x7c, 0x30, 0x00, 0x3f,
	0x14, 0x59, 0xe0, 0x53, 0x9a, 0x21, 0xe7, 0xd0, 0x09, 0x85, 0x0c, 0x4d, 0x66, 0x33, 0xb7, 0xef,
	0xe9, 0xb7, 0xc2, 0x64, 0xf4, 0x8e, 0x66, 0xdb, 0x66, 0xae, 0xe1, 0xe9, 0x37, 0x3f, 0x06, 0x83,
	0x28, 0x36, 0x0d, 0x0d, 0xa9, 0x27, 0xb7, 0x61, 0x20, 0x15, 0xc5, 0xe3, 0x6c, 0x26, 0x91, 0xcc,
	0x8e, 0xce, 0x54, 0x21, 0x6e, 0xc2, 0x9f, 0x69, 0xba, 0x20, 0x5c, 0x90, 0xd9, 0xb5, 0x99, 0xfb,
	0xcf, 0x5b, 0x87, 0xce, 0x33, 0x1c, 0x69, 0x0f, 0x1e, 0x52, 0x16, 0xe1, 0x4a, 0xc4, 0x3f, 0xf6,
	0xb1, 0xa3, 0x6a, 0x34, 0x54, 0x9d, 0x09, 0x9c, 0xd4, 0xb9, 0x7d, 0xca, 0x50, 0x24, 0x1b, 0x36,
	0x56, 0x61, 0xab, 0x38, 0x6c, 0xd7, 0x1d, 0x5e, 0xc0, 0x40, 0xb3, 0x4c, 0x30, 0x46, 0xda, 0x3b,
	0x26, 0x27, 0x04, 0x5e, 0x29, 0xf1, 0xf3, 0x24, 0x11, 0xd9, 0x1b, 0x1f, 0x41, 0x4f, 0x92, 0xa0,
	0x5c, 0x96, 0x42, 0x65, 0xa4, 0xa4, 0x12, 0x94, 0x52, 0xcc, 0x8b, 0x7e, 0xfa, 0xde, 0x3a, 0x54,
	0x2d, 0x61, 0x2c, 0x96, 0x12, 0x83, 0xa7, 0x28, 0xc1, 0x75, 0x4b, 0x15, 0xc8, 0xf9, 0x64, 0x30,
	0xdc, 0xee, 0xec, 0xf7, 0x4a, 0x97, 0xf0, 0x9f, 0x52, 0x12, 0xb1, 0x87, 0x53, 0x8c, 0x56, 0x18,
	0x94, 0x5a, 0x75, 0x70, 0xd7, 0x4f, 0xa7, 0xe1, 0xe7, 0xfa, 0x8b, 0x01, 0x78, 0xea, 0xa4, 0xee,
	0xd5, 0x49, 0xf1, 0x07, 0xf8, 0x5b, 0x0e, 0x1b, 0xb9, 0x35, 0xde, 0xde, 0xda, 0xb8, 0xbe, 0x07,
	0xcb, 0x3e, 0x9c, 0x2b, 0x76, 0xe4, 0xb4, 0xae, 0x18, 0xbf, 0x85, 0x6e, 0x71, 0x9a, 0xa3, 0x46,
	0xb9, 0xc6, 0xad, 0xb3, 0xfd, 0x78, 0x39, 0x16, 0xa7, 0xe5, 0x32, 0x7e, 0x07, 0xbd, 0x72, 0x71,
	0xa7, 0x8d, 0xe2, 0x22, 0x61, 0x9d, 0x1f, 0x48, 0x6c, 0x68, 0x5e, 0x7b, 0xfa, 0xf3, 0xdc, 0x7c,
	0x07, 0x00, 0x00, 0xff, 0xff, 0x08, 0x6e, 0x4e, 0x5b, 0x51, 0x03, 0x00, 0x00,
}
