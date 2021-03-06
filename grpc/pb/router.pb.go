// Code generated by protoc-gen-go. DO NOT EDIT.
// source: router.proto

package pb

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

type SSRouter_TransferType int32

const (
	SSRouter_TRANSTYPHEARTBEAT    SSRouter_TransferType = 0
	SSRouter_TRANSTYPEBYP2P       SSRouter_TransferType = 1
	SSRouter_TRANSTYPEBYP2G       SSRouter_TransferType = 2
	SSRouter_TRANSTYPEBYBROADCAST SSRouter_TransferType = 3
	SSRouter_TRANSTYPEBYKEY       SSRouter_TransferType = 4
)

var SSRouter_TransferType_name = map[int32]string{
	0: "TRANSTYPHEARTBEAT",
	1: "TRANSTYPEBYP2P",
	2: "TRANSTYPEBYP2G",
	3: "TRANSTYPEBYBROADCAST",
	4: "TRANSTYPEBYKEY",
}
var SSRouter_TransferType_value = map[string]int32{
	"TRANSTYPHEARTBEAT":    0,
	"TRANSTYPEBYP2P":       1,
	"TRANSTYPEBYP2G":       2,
	"TRANSTYPEBYBROADCAST": 3,
	"TRANSTYPEBYKEY":       4,
}

func (x SSRouter_TransferType) String() string {
	return proto.EnumName(SSRouter_TransferType_name, int32(x))
}
func (SSRouter_TransferType) EnumDescriptor() ([]byte, []int) { return fileDescriptor1, []int{0, 0} }

// SSRouter Server Message结构
type SSRouter struct {
	SrcSID    int32                 `protobuf:"varint,1,opt,name=srcSID" json:"srcSID,omitempty"`
	SrcType   int32                 `protobuf:"varint,2,opt,name=srcType" json:"srcType,omitempty"`
	DestSID   int32                 `protobuf:"varint,3,opt,name=destSID" json:"destSID,omitempty"`
	DestType  int32                 `protobuf:"varint,4,opt,name=destType" json:"destType,omitempty"`
	TransType SSRouter_TransferType `protobuf:"varint,5,opt,name=transType,enum=pb.SSRouter_TransferType" json:"transType,omitempty"`
	Uid       int64                 `protobuf:"varint,6,opt,name=uid" json:"uid,omitempty"`
	Body      []byte                `protobuf:"bytes,7,opt,name=body,proto3" json:"body,omitempty"`
}

func (m *SSRouter) Reset()                    { *m = SSRouter{} }
func (m *SSRouter) String() string            { return proto.CompactTextString(m) }
func (*SSRouter) ProtoMessage()               {}
func (*SSRouter) Descriptor() ([]byte, []int) { return fileDescriptor1, []int{0} }

func (m *SSRouter) GetSrcSID() int32 {
	if m != nil {
		return m.SrcSID
	}
	return 0
}

func (m *SSRouter) GetSrcType() int32 {
	if m != nil {
		return m.SrcType
	}
	return 0
}

func (m *SSRouter) GetDestSID() int32 {
	if m != nil {
		return m.DestSID
	}
	return 0
}

func (m *SSRouter) GetDestType() int32 {
	if m != nil {
		return m.DestType
	}
	return 0
}

func (m *SSRouter) GetTransType() SSRouter_TransferType {
	if m != nil {
		return m.TransType
	}
	return SSRouter_TRANSTYPHEARTBEAT
}

func (m *SSRouter) GetUid() int64 {
	if m != nil {
		return m.Uid
	}
	return 0
}

func (m *SSRouter) GetBody() []byte {
	if m != nil {
		return m.Body
	}
	return nil
}

func init() {
	proto.RegisterType((*SSRouter)(nil), "pb.SSRouter")
	proto.RegisterEnum("pb.SSRouter_TransferType", SSRouter_TransferType_name, SSRouter_TransferType_value)
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion4

// Client API for Router service

type RouterClient interface {
	// 定义Proxy方法
	Proxy(ctx context.Context, opts ...grpc.CallOption) (Router_ProxyClient, error)
}

type routerClient struct {
	cc *grpc.ClientConn
}

func NewRouterClient(cc *grpc.ClientConn) RouterClient {
	return &routerClient{cc}
}

func (c *routerClient) Proxy(ctx context.Context, opts ...grpc.CallOption) (Router_ProxyClient, error) {
	stream, err := grpc.NewClientStream(ctx, &_Router_serviceDesc.Streams[0], c.cc, "/pb.Router/Proxy", opts...)
	if err != nil {
		return nil, err
	}
	x := &routerProxyClient{stream}
	return x, nil
}

type Router_ProxyClient interface {
	Send(*SSRouter) error
	Recv() (*SSRouter, error)
	grpc.ClientStream
}

type routerProxyClient struct {
	grpc.ClientStream
}

func (x *routerProxyClient) Send(m *SSRouter) error {
	return x.ClientStream.SendMsg(m)
}

func (x *routerProxyClient) Recv() (*SSRouter, error) {
	m := new(SSRouter)
	if err := x.ClientStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

// Server API for Router service

type RouterServer interface {
	// 定义Proxy方法
	Proxy(Router_ProxyServer) error
}

func RegisterRouterServer(s *grpc.Server, srv RouterServer) {
	s.RegisterService(&_Router_serviceDesc, srv)
}

func _Router_Proxy_Handler(srv interface{}, stream grpc.ServerStream) error {
	return srv.(RouterServer).Proxy(&routerProxyServer{stream})
}

type Router_ProxyServer interface {
	Send(*SSRouter) error
	Recv() (*SSRouter, error)
	grpc.ServerStream
}

type routerProxyServer struct {
	grpc.ServerStream
}

func (x *routerProxyServer) Send(m *SSRouter) error {
	return x.ServerStream.SendMsg(m)
}

func (x *routerProxyServer) Recv() (*SSRouter, error) {
	m := new(SSRouter)
	if err := x.ServerStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

var _Router_serviceDesc = grpc.ServiceDesc{
	ServiceName: "pb.Router",
	HandlerType: (*RouterServer)(nil),
	Methods:     []grpc.MethodDesc{},
	Streams: []grpc.StreamDesc{
		{
			StreamName:    "Proxy",
			Handler:       _Router_Proxy_Handler,
			ServerStreams: true,
			ClientStreams: true,
		},
	},
	Metadata: "router.proto",
}

func init() { proto.RegisterFile("router.proto", fileDescriptor1) }

var fileDescriptor1 = []byte{
	// 282 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x64, 0x91, 0xcf, 0x4e, 0xb4, 0x30,
	0x14, 0xc5, 0xa7, 0xfc, 0x9b, 0xf9, 0x6e, 0xc8, 0xa4, 0xdf, 0x8d, 0x9a, 0x3a, 0x2b, 0xc2, 0x0a,
	0x37, 0xc4, 0x30, 0x0b, 0xd7, 0xe0, 0x10, 0x35, 0x26, 0x4a, 0x4a, 0x37, 0xb8, 0x93, 0x01, 0x13,
	0x37, 0x42, 0x0a, 0x93, 0x48, 0x7c, 0x47, 0x9f, 0xc9, 0xb4, 0x82, 0xa2, 0xee, 0xce, 0x39, 0xbf,
	0x73, 0x9b, 0xb6, 0x17, 0x5c, 0xd9, 0x1c, 0xfa, 0x5a, 0x86, 0xad, 0x6c, 0xfa, 0x06, 0x8d, 0xb6,
	0xf4, 0xdf, 0x0d, 0x58, 0xe5, 0x39, 0xd7, 0x31, 0x9e, 0x80, 0xd3, 0xc9, 0x7d, 0x7e, 0xb3, 0x63,
	0xc4, 0x23, 0x81, 0xcd, 0x47, 0x87, 0x0c, 0x96, 0x9d, 0xdc, 0x8b, 0xa1, 0xad, 0x99, 0xa1, 0xc1,
	0x64, 0x15, 0xa9, 0xea, 0xae, 0x57, 0x23, 0xe6, 0x27, 0x19, 0x2d, 0x6e, 0x60, 0xa5, 0xa4, 0x1e,
	0xb2, 0x34, 0xfa, 0xf2, 0x78, 0x01, 0xff, 0x7a, 0xf9, 0xf8, 0xd2, 0x69, 0x68, 0x7b, 0x24, 0x58,
	0x47, 0xa7, 0x61, 0x5b, 0x86, 0xd3, 0x45, 0x42, 0xa1, 0xe8, 0x53, 0x2d, 0x55, 0x81, 0x7f, 0x77,
	0x91, 0x82, 0x79, 0x78, 0xae, 0x98, 0xe3, 0x91, 0xc0, 0xe4, 0x4a, 0x22, 0x82, 0x55, 0x36, 0xd5,
	0xc0, 0x96, 0x1e, 0x09, 0x5c, 0xae, 0xb5, 0xff, 0x06, 0xee, 0xfc, 0x00, 0x3c, 0x86, 0xff, 0x82,
	0xc7, 0x77, 0xb9, 0x28, 0xb2, 0xeb, 0x34, 0xe6, 0x22, 0x49, 0x63, 0x41, 0x17, 0x88, 0xb0, 0x9e,
	0xe2, 0x34, 0x29, 0xb2, 0x28, 0xa3, 0xe4, 0x4f, 0x76, 0x45, 0x0d, 0x64, 0x70, 0x34, 0xcb, 0x12,
	0x7e, 0x1f, 0xef, 0x2e, 0xe3, 0x5c, 0x50, 0xf3, 0x57, 0xfb, 0x36, 0x2d, 0xa8, 0x15, 0x6d, 0xc1,
	0x19, 0x7f, 0xf3, 0x0c, 0xec, 0x4c, 0x36, 0xaf, 0x03, 0xba, 0xf3, 0xb7, 0x6d, 0x7e, 0x38, 0x7f,
	0x11, 0x90, 0x73, 0x92, 0x58, 0x0f, 0x46, 0x5b, 0x96, 0x8e, 0x5e, 0xcb, 0xf6, 0x23, 0x00, 0x00,
	0xff, 0xff, 0x58, 0x31, 0x4f, 0x5b, 0xa6, 0x01, 0x00, 0x00,
}
