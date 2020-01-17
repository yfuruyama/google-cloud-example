// Code generated by protoc-gen-go. DO NOT EDIT.
// source: proto/hello.proto

package hello

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

type EchoRequest struct {
	Msg                  string   `protobuf:"bytes,1,opt,name=msg,proto3" json:"msg,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *EchoRequest) Reset()         { *m = EchoRequest{} }
func (m *EchoRequest) String() string { return proto.CompactTextString(m) }
func (*EchoRequest) ProtoMessage()    {}
func (*EchoRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_f7116fbf19896b02, []int{0}
}

func (m *EchoRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_EchoRequest.Unmarshal(m, b)
}
func (m *EchoRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_EchoRequest.Marshal(b, m, deterministic)
}
func (m *EchoRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_EchoRequest.Merge(m, src)
}
func (m *EchoRequest) XXX_Size() int {
	return xxx_messageInfo_EchoRequest.Size(m)
}
func (m *EchoRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_EchoRequest.DiscardUnknown(m)
}

var xxx_messageInfo_EchoRequest proto.InternalMessageInfo

func (m *EchoRequest) GetMsg() string {
	if m != nil {
		return m.Msg
	}
	return ""
}

type EchoReply struct {
	Msg                  string   `protobuf:"bytes,1,opt,name=msg,proto3" json:"msg,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *EchoReply) Reset()         { *m = EchoReply{} }
func (m *EchoReply) String() string { return proto.CompactTextString(m) }
func (*EchoReply) ProtoMessage()    {}
func (*EchoReply) Descriptor() ([]byte, []int) {
	return fileDescriptor_f7116fbf19896b02, []int{1}
}

func (m *EchoReply) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_EchoReply.Unmarshal(m, b)
}
func (m *EchoReply) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_EchoReply.Marshal(b, m, deterministic)
}
func (m *EchoReply) XXX_Merge(src proto.Message) {
	xxx_messageInfo_EchoReply.Merge(m, src)
}
func (m *EchoReply) XXX_Size() int {
	return xxx_messageInfo_EchoReply.Size(m)
}
func (m *EchoReply) XXX_DiscardUnknown() {
	xxx_messageInfo_EchoReply.DiscardUnknown(m)
}

var xxx_messageInfo_EchoReply proto.InternalMessageInfo

func (m *EchoReply) GetMsg() string {
	if m != nil {
		return m.Msg
	}
	return ""
}

func init() {
	proto.RegisterType((*EchoRequest)(nil), "hello.EchoRequest")
	proto.RegisterType((*EchoReply)(nil), "hello.EchoReply")
}

func init() { proto.RegisterFile("proto/hello.proto", fileDescriptor_f7116fbf19896b02) }

var fileDescriptor_f7116fbf19896b02 = []byte{
	// 120 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xe2, 0x12, 0x2c, 0x28, 0xca, 0x2f,
	0xc9, 0xd7, 0xcf, 0x48, 0xcd, 0xc9, 0xc9, 0xd7, 0x03, 0xb3, 0x85, 0x58, 0xc1, 0x1c, 0x25, 0x79,
	0x2e, 0x6e, 0xd7, 0xe4, 0x8c, 0xfc, 0xa0, 0xd4, 0xc2, 0xd2, 0xd4, 0xe2, 0x12, 0x21, 0x01, 0x2e,
	0xe6, 0xdc, 0xe2, 0x74, 0x09, 0x46, 0x05, 0x46, 0x0d, 0xce, 0x20, 0x10, 0x53, 0x49, 0x96, 0x8b,
	0x13, 0xa2, 0xa0, 0x20, 0xa7, 0x12, 0x53, 0xda, 0xc8, 0x9c, 0x8b, 0xd5, 0x03, 0x64, 0x90, 0x90,
	0x1e, 0x17, 0x0b, 0x48, 0x9d, 0x90, 0x90, 0x1e, 0xc4, 0x16, 0x24, 0x53, 0xa5, 0x04, 0x50, 0xc4,
	0x0a, 0x72, 0x2a, 0x95, 0x18, 0x92, 0xd8, 0xc0, 0xce, 0x30, 0x06, 0x04, 0x00, 0x00, 0xff, 0xff,
	0x35, 0x91, 0x14, 0x75, 0x9b, 0x00, 0x00, 0x00,
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion4

// HelloClient is the client API for Hello service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type HelloClient interface {
	Echo(ctx context.Context, in *EchoRequest, opts ...grpc.CallOption) (*EchoReply, error)
}

type helloClient struct {
	cc *grpc.ClientConn
}

func NewHelloClient(cc *grpc.ClientConn) HelloClient {
	return &helloClient{cc}
}

func (c *helloClient) Echo(ctx context.Context, in *EchoRequest, opts ...grpc.CallOption) (*EchoReply, error) {
	out := new(EchoReply)
	err := c.cc.Invoke(ctx, "/hello.Hello/Echo", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// HelloServer is the server API for Hello service.
type HelloServer interface {
	Echo(context.Context, *EchoRequest) (*EchoReply, error)
}

// UnimplementedHelloServer can be embedded to have forward compatible implementations.
type UnimplementedHelloServer struct {
}

func (*UnimplementedHelloServer) Echo(ctx context.Context, req *EchoRequest) (*EchoReply, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Echo not implemented")
}

func RegisterHelloServer(s *grpc.Server, srv HelloServer) {
	s.RegisterService(&_Hello_serviceDesc, srv)
}

func _Hello_Echo_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(EchoRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(HelloServer).Echo(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/hello.Hello/Echo",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(HelloServer).Echo(ctx, req.(*EchoRequest))
	}
	return interceptor(ctx, in, info, handler)
}

var _Hello_serviceDesc = grpc.ServiceDesc{
	ServiceName: "hello.Hello",
	HandlerType: (*HelloServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "Echo",
			Handler:    _Hello_Echo_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "proto/hello.proto",
}
