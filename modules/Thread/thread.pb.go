// Code generated by protoc-gen-go. DO NOT EDIT.
// source: thread/thread.proto

package thread

import (
	context "context"
	fmt "fmt"
	proto "github.com/golang/protobuf/proto"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
	math "math"
)

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.ProtoPackageIsVersion3 // please upgrade the proto package

// thread 请求数据
type TheadRequest struct {
	Id                   string   `protobuf:"bytes,1,opt,name=Id,proto3" json:"Id,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *TheadRequest) Reset()         { *m = TheadRequest{} }
func (m *TheadRequest) String() string { return proto.CompactTextString(m) }
func (*TheadRequest) ProtoMessage()    {}
func (*TheadRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_7d712470c362cd36, []int{0}
}

func (m *TheadRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_TheadRequest.Unmarshal(m, b)
}
func (m *TheadRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_TheadRequest.Marshal(b, m, deterministic)
}
func (m *TheadRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_TheadRequest.Merge(m, src)
}
func (m *TheadRequest) XXX_Size() int {
	return xxx_messageInfo_TheadRequest.Size(m)
}
func (m *TheadRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_TheadRequest.DiscardUnknown(m)
}

var xxx_messageInfo_TheadRequest proto.InternalMessageInfo

func (m *TheadRequest) GetId() string {
	if m != nil {
		return m.Id
	}
	return ""
}

// thread 响应数据
type TheadReply struct {
	Id                   string   `protobuf:"bytes,1,opt,name=Id,proto3" json:"Id,omitempty"`
	Title                string   `protobuf:"bytes,2,opt,name=Title,proto3" json:"Title,omitempty"`
	Image                string   `protobuf:"bytes,3,opt,name=Image,proto3" json:"Image,omitempty"`
	Content              string   `protobuf:"bytes,4,opt,name=Content,proto3" json:"Content,omitempty"`
	CreatedAt            string   `protobuf:"bytes,5,opt,name=CreatedAt,proto3" json:"CreatedAt,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *TheadReply) Reset()         { *m = TheadReply{} }
func (m *TheadReply) String() string { return proto.CompactTextString(m) }
func (*TheadReply) ProtoMessage()    {}
func (*TheadReply) Descriptor() ([]byte, []int) {
	return fileDescriptor_7d712470c362cd36, []int{1}
}

func (m *TheadReply) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_TheadReply.Unmarshal(m, b)
}
func (m *TheadReply) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_TheadReply.Marshal(b, m, deterministic)
}
func (m *TheadReply) XXX_Merge(src proto.Message) {
	xxx_messageInfo_TheadReply.Merge(m, src)
}
func (m *TheadReply) XXX_Size() int {
	return xxx_messageInfo_TheadReply.Size(m)
}
func (m *TheadReply) XXX_DiscardUnknown() {
	xxx_messageInfo_TheadReply.DiscardUnknown(m)
}

var xxx_messageInfo_TheadReply proto.InternalMessageInfo

func (m *TheadReply) GetId() string {
	if m != nil {
		return m.Id
	}
	return ""
}

func (m *TheadReply) GetTitle() string {
	if m != nil {
		return m.Title
	}
	return ""
}

func (m *TheadReply) GetImage() string {
	if m != nil {
		return m.Image
	}
	return ""
}

func (m *TheadReply) GetContent() string {
	if m != nil {
		return m.Content
	}
	return ""
}

func (m *TheadReply) GetCreatedAt() string {
	if m != nil {
		return m.CreatedAt
	}
	return ""
}

// thread list request 数据
type ThreadListRequest struct {
	Order                string   `protobuf:"bytes,1,opt,name=Order,proto3" json:"Order,omitempty"`
	CreatedAt            string   `protobuf:"bytes,2,opt,name=CreatedAt,proto3" json:"CreatedAt,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *ThreadListRequest) Reset()         { *m = ThreadListRequest{} }
func (m *ThreadListRequest) String() string { return proto.CompactTextString(m) }
func (*ThreadListRequest) ProtoMessage()    {}
func (*ThreadListRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_7d712470c362cd36, []int{2}
}

func (m *ThreadListRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ThreadListRequest.Unmarshal(m, b)
}
func (m *ThreadListRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ThreadListRequest.Marshal(b, m, deterministic)
}
func (m *ThreadListRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ThreadListRequest.Merge(m, src)
}
func (m *ThreadListRequest) XXX_Size() int {
	return xxx_messageInfo_ThreadListRequest.Size(m)
}
func (m *ThreadListRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_ThreadListRequest.DiscardUnknown(m)
}

var xxx_messageInfo_ThreadListRequest proto.InternalMessageInfo

func (m *ThreadListRequest) GetOrder() string {
	if m != nil {
		return m.Order
	}
	return ""
}

func (m *ThreadListRequest) GetCreatedAt() string {
	if m != nil {
		return m.CreatedAt
	}
	return ""
}

// thread list response 数据
type ThreadListReply struct {
	List                 []*TheadReply `protobuf:"bytes,1,rep,name=list,proto3" json:"list,omitempty"`
	XXX_NoUnkeyedLiteral struct{}      `json:"-"`
	XXX_unrecognized     []byte        `json:"-"`
	XXX_sizecache        int32         `json:"-"`
}

func (m *ThreadListReply) Reset()         { *m = ThreadListReply{} }
func (m *ThreadListReply) String() string { return proto.CompactTextString(m) }
func (*ThreadListReply) ProtoMessage()    {}
func (*ThreadListReply) Descriptor() ([]byte, []int) {
	return fileDescriptor_7d712470c362cd36, []int{3}
}

func (m *ThreadListReply) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ThreadListReply.Unmarshal(m, b)
}
func (m *ThreadListReply) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ThreadListReply.Marshal(b, m, deterministic)
}
func (m *ThreadListReply) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ThreadListReply.Merge(m, src)
}
func (m *ThreadListReply) XXX_Size() int {
	return xxx_messageInfo_ThreadListReply.Size(m)
}
func (m *ThreadListReply) XXX_DiscardUnknown() {
	xxx_messageInfo_ThreadListReply.DiscardUnknown(m)
}

var xxx_messageInfo_ThreadListReply proto.InternalMessageInfo

func (m *ThreadListReply) GetList() []*TheadReply {
	if m != nil {
		return m.List
	}
	return nil
}

func init() {
	proto.RegisterType((*TheadRequest)(nil), "thread.TheadRequest")
	proto.RegisterType((*TheadReply)(nil), "thread.TheadReply")
	proto.RegisterType((*ThreadListRequest)(nil), "thread.ThreadListRequest")
	proto.RegisterType((*ThreadListReply)(nil), "thread.ThreadListReply")
}

func init() { proto.RegisterFile("thread/thread.proto", fileDescriptor_7d712470c362cd36) }

var fileDescriptor_7d712470c362cd36 = []byte{
	// 265 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x6c, 0x51, 0x4d, 0x4b, 0xc3, 0x40,
	0x10, 0x75, 0xd3, 0x26, 0xd2, 0xf1, 0x0b, 0xc7, 0x82, 0x6b, 0x11, 0x29, 0x39, 0x48, 0x4f, 0x15,
	0xea, 0x49, 0x3c, 0x69, 0x91, 0x12, 0x10, 0x84, 0x90, 0x3f, 0x10, 0xd9, 0xc1, 0x06, 0x62, 0x53,
	0x37, 0xe3, 0xa1, 0x37, 0x7f, 0x80, 0x3f, 0x5a, 0x32, 0x9b, 0x92, 0xba, 0xf4, 0xb4, 0xbc, 0xf7,
	0x66, 0x1e, 0xef, 0xcd, 0xc2, 0x05, 0x2f, 0x2d, 0xe5, 0xe6, 0xce, 0x3d, 0xd3, 0xb5, 0xad, 0xb8,
	0xc2, 0xc8, 0xa1, 0xf8, 0x06, 0x8e, 0xb3, 0x25, 0xe5, 0x26, 0xa5, 0xaf, 0x6f, 0xaa, 0x19, 0x4f,
	0x21, 0x48, 0x8c, 0x56, 0x63, 0x35, 0x19, 0xa4, 0x41, 0x62, 0xe2, 0x1f, 0x05, 0xd0, 0x0e, 0xac,
	0xcb, 0x8d, 0x2f, 0xe3, 0x10, 0xc2, 0xac, 0xe0, 0x92, 0x74, 0x20, 0x94, 0x03, 0x0d, 0x9b, 0x7c,
	0xe6, 0x1f, 0xa4, 0x7b, 0x8e, 0x15, 0x80, 0x1a, 0x0e, 0xe7, 0xd5, 0x8a, 0x69, 0xc5, 0xba, 0x2f,
	0xfc, 0x16, 0xe2, 0x35, 0x0c, 0xe6, 0x96, 0x72, 0x26, 0xf3, 0xc4, 0x3a, 0x14, 0xad, 0x23, 0xe2,
	0x05, 0x9c, 0x67, 0x12, 0xf6, 0xb5, 0xa8, 0x79, 0x9b, 0x73, 0x08, 0xe1, 0x9b, 0x35, 0x64, 0xdb,
	0x2c, 0x0e, 0xfc, 0x37, 0x0a, 0x7c, 0xa3, 0x07, 0x38, 0xdb, 0x35, 0x6a, 0xfa, 0xdc, 0x42, 0xbf,
	0x2c, 0x6a, 0xd6, 0x6a, 0xdc, 0x9b, 0x1c, 0xcd, 0x70, 0xda, 0xde, 0xa8, 0x6b, 0x9c, 0x8a, 0x3e,
	0xfb, 0x55, 0x10, 0xb9, 0x5d, 0x7c, 0x84, 0x93, 0x05, 0xb1, 0x03, 0xcf, 0x9b, 0xe6, 0x06, 0xde,
	0x96, 0x04, 0x1c, 0xed, 0xf1, 0x8a, 0x0f, 0xf0, 0x65, 0x67, 0xb9, 0x49, 0x81, 0x57, 0xdd, 0x98,
	0x57, 0x71, 0x74, 0xb9, 0x4f, 0x12, 0x9b, 0xf7, 0x48, 0x3e, 0xf1, 0xfe, 0x2f, 0x00, 0x00, 0xff,
	0xff, 0x42, 0x7a, 0xd7, 0x3a, 0xdb, 0x01, 0x00, 0x00,
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion4

// ThreadClient is the client API for Thread service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type ThreadClient interface {
	GetThreadById(ctx context.Context, in *TheadRequest, opts ...grpc.CallOption) (*TheadReply, error)
	GetThreadList(ctx context.Context, in *ThreadListRequest, opts ...grpc.CallOption) (*ThreadListReply, error)
}

type threadClient struct {
	cc *grpc.ClientConn
}

func NewThreadClient(cc *grpc.ClientConn) ThreadClient {
	return &threadClient{cc}
}

func (c *threadClient) GetThreadById(ctx context.Context, in *TheadRequest, opts ...grpc.CallOption) (*TheadReply, error) {
	out := new(TheadReply)
	err := c.cc.Invoke(ctx, "/thread.Thread/GetThreadById", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *threadClient) GetThreadList(ctx context.Context, in *ThreadListRequest, opts ...grpc.CallOption) (*ThreadListReply, error) {
	out := new(ThreadListReply)
	err := c.cc.Invoke(ctx, "/thread.Thread/GetThreadList", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// ThreadServer is the server API for Thread service.
type ThreadServer interface {
	GetThreadById(context.Context, *TheadRequest) (*TheadReply, error)
	GetThreadList(context.Context, *ThreadListRequest) (*ThreadListReply, error)
}

// UnimplementedThreadServer can be embedded to have forward compatible implementations.
type UnimplementedThreadServer struct {
}

func (*UnimplementedThreadServer) GetThreadById(ctx context.Context, req *TheadRequest) (*TheadReply, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetThreadById not implemented")
}
func (*UnimplementedThreadServer) GetThreadList(ctx context.Context, req *ThreadListRequest) (*ThreadListReply, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetThreadList not implemented")
}

func RegisterThreadServer(s *grpc.Server, srv ThreadServer) {
	s.RegisterService(&_Thread_serviceDesc, srv)
}

func _Thread_GetThreadById_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(TheadRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ThreadServer).GetThreadById(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/thread.Thread/GetThreadById",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ThreadServer).GetThreadById(ctx, req.(*TheadRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Thread_GetThreadList_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ThreadListRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ThreadServer).GetThreadList(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/thread.Thread/GetThreadList",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ThreadServer).GetThreadList(ctx, req.(*ThreadListRequest))
	}
	return interceptor(ctx, in, info, handler)
}

var _Thread_serviceDesc = grpc.ServiceDesc{
	ServiceName: "thread.Thread",
	HandlerType: (*ThreadServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "GetThreadById",
			Handler:    _Thread_GetThreadById_Handler,
		},
		{
			MethodName: "GetThreadList",
			Handler:    _Thread_GetThreadList_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "thread/thread.proto",
}
