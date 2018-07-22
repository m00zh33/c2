// Code generated by protoc-gen-go. DO NOT EDIT.
// source: c2.proto

package c2

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

type C2Request_Command int32

const (
	C2Request_NEW_CLIENT          C2Request_Command = 0
	C2Request_REMOVE_CLIENT       C2Request_Command = 1
	C2Request_NEW_TOPIC_CLIENT    C2Request_Command = 2
	C2Request_REMOVE_TOPIC_CLIENT C2Request_Command = 3
	C2Request_RESET_CLIENT        C2Request_Command = 4
	C2Request_NEW_TOPIC           C2Request_Command = 5
	C2Request_REMOVE_TOPIC        C2Request_Command = 6
	C2Request_NEW_CLIENT_KEY      C2Request_Command = 7
	C2Request_SEND_MESSAGE        C2Request_Command = 8
)

var C2Request_Command_name = map[int32]string{
	0: "NEW_CLIENT",
	1: "REMOVE_CLIENT",
	2: "NEW_TOPIC_CLIENT",
	3: "REMOVE_TOPIC_CLIENT",
	4: "RESET_CLIENT",
	5: "NEW_TOPIC",
	6: "REMOVE_TOPIC",
	7: "NEW_CLIENT_KEY",
	8: "SEND_MESSAGE",
}
var C2Request_Command_value = map[string]int32{
	"NEW_CLIENT":          0,
	"REMOVE_CLIENT":       1,
	"NEW_TOPIC_CLIENT":    2,
	"REMOVE_TOPIC_CLIENT": 3,
	"RESET_CLIENT":        4,
	"NEW_TOPIC":           5,
	"REMOVE_TOPIC":        6,
	"NEW_CLIENT_KEY":      7,
	"SEND_MESSAGE":        8,
}

func (x C2Request_Command) String() string {
	return proto.EnumName(C2Request_Command_name, int32(x))
}
func (C2Request_Command) EnumDescriptor() ([]byte, []int) {
	return fileDescriptor_c2_d2621b88c04c2ce9, []int{0, 0}
}

type C2Request struct {
	Command              C2Request_Command `protobuf:"varint,1,opt,name=command,proto3,enum=c2.C2Request_Command" json:"command,omitempty"`
	Id                   []byte            `protobuf:"bytes,2,opt,name=id,proto3" json:"id,omitempty"`
	Key                  []byte            `protobuf:"bytes,3,opt,name=key,proto3" json:"key,omitempty"`
	Topic                string            `protobuf:"bytes,4,opt,name=topic,proto3" json:"topic,omitempty"`
	Msg                  string            `protobuf:"bytes,5,opt,name=msg,proto3" json:"msg,omitempty"`
	XXX_NoUnkeyedLiteral struct{}          `json:"-"`
	XXX_unrecognized     []byte            `json:"-"`
	XXX_sizecache        int32             `json:"-"`
}

func (m *C2Request) Reset()         { *m = C2Request{} }
func (m *C2Request) String() string { return proto.CompactTextString(m) }
func (*C2Request) ProtoMessage()    {}
func (*C2Request) Descriptor() ([]byte, []int) {
	return fileDescriptor_c2_d2621b88c04c2ce9, []int{0}
}
func (m *C2Request) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_C2Request.Unmarshal(m, b)
}
func (m *C2Request) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_C2Request.Marshal(b, m, deterministic)
}
func (dst *C2Request) XXX_Merge(src proto.Message) {
	xxx_messageInfo_C2Request.Merge(dst, src)
}
func (m *C2Request) XXX_Size() int {
	return xxx_messageInfo_C2Request.Size(m)
}
func (m *C2Request) XXX_DiscardUnknown() {
	xxx_messageInfo_C2Request.DiscardUnknown(m)
}

var xxx_messageInfo_C2Request proto.InternalMessageInfo

func (m *C2Request) GetCommand() C2Request_Command {
	if m != nil {
		return m.Command
	}
	return C2Request_NEW_CLIENT
}

func (m *C2Request) GetId() []byte {
	if m != nil {
		return m.Id
	}
	return nil
}

func (m *C2Request) GetKey() []byte {
	if m != nil {
		return m.Key
	}
	return nil
}

func (m *C2Request) GetTopic() string {
	if m != nil {
		return m.Topic
	}
	return ""
}

func (m *C2Request) GetMsg() string {
	if m != nil {
		return m.Msg
	}
	return ""
}

type C2Response struct {
	Success              bool     `protobuf:"varint,1,opt,name=success,proto3" json:"success,omitempty"`
	Err                  string   `protobuf:"bytes,2,opt,name=err,proto3" json:"err,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *C2Response) Reset()         { *m = C2Response{} }
func (m *C2Response) String() string { return proto.CompactTextString(m) }
func (*C2Response) ProtoMessage()    {}
func (*C2Response) Descriptor() ([]byte, []int) {
	return fileDescriptor_c2_d2621b88c04c2ce9, []int{1}
}
func (m *C2Response) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_C2Response.Unmarshal(m, b)
}
func (m *C2Response) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_C2Response.Marshal(b, m, deterministic)
}
func (dst *C2Response) XXX_Merge(src proto.Message) {
	xxx_messageInfo_C2Response.Merge(dst, src)
}
func (m *C2Response) XXX_Size() int {
	return xxx_messageInfo_C2Response.Size(m)
}
func (m *C2Response) XXX_DiscardUnknown() {
	xxx_messageInfo_C2Response.DiscardUnknown(m)
}

var xxx_messageInfo_C2Response proto.InternalMessageInfo

func (m *C2Response) GetSuccess() bool {
	if m != nil {
		return m.Success
	}
	return false
}

func (m *C2Response) GetErr() string {
	if m != nil {
		return m.Err
	}
	return ""
}

func init() {
	proto.RegisterType((*C2Request)(nil), "c2.C2Request")
	proto.RegisterType((*C2Response)(nil), "c2.C2Response")
	proto.RegisterEnum("c2.C2Request_Command", C2Request_Command_name, C2Request_Command_value)
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion4

// C2Client is the client API for C2 service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type C2Client interface {
	C2Command(ctx context.Context, in *C2Request, opts ...grpc.CallOption) (*C2Response, error)
}

type c2Client struct {
	cc *grpc.ClientConn
}

func NewC2Client(cc *grpc.ClientConn) C2Client {
	return &c2Client{cc}
}

func (c *c2Client) C2Command(ctx context.Context, in *C2Request, opts ...grpc.CallOption) (*C2Response, error) {
	out := new(C2Response)
	err := c.cc.Invoke(ctx, "/c2.C2/C2Command", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// C2Server is the server API for C2 service.
type C2Server interface {
	C2Command(context.Context, *C2Request) (*C2Response, error)
}

func RegisterC2Server(s *grpc.Server, srv C2Server) {
	s.RegisterService(&_C2_serviceDesc, srv)
}

func _C2_C2Command_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(C2Request)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(C2Server).C2Command(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/c2.C2/C2Command",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(C2Server).C2Command(ctx, req.(*C2Request))
	}
	return interceptor(ctx, in, info, handler)
}

var _C2_serviceDesc = grpc.ServiceDesc{
	ServiceName: "c2.C2",
	HandlerType: (*C2Server)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "C2Command",
			Handler:    _C2_C2Command_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "c2.proto",
}

func init() { proto.RegisterFile("c2.proto", fileDescriptor_c2_d2621b88c04c2ce9) }

var fileDescriptor_c2_d2621b88c04c2ce9 = []byte{
	// 308 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x54, 0x91, 0x4f, 0x4f, 0x02, 0x31,
	0x10, 0xc5, 0x69, 0xf9, 0xb3, 0xec, 0x04, 0x36, 0x75, 0xc4, 0xd8, 0x78, 0x22, 0x7b, 0xe2, 0x60,
	0xd6, 0xa4, 0x5e, 0xbc, 0x9a, 0xda, 0x18, 0xa2, 0x80, 0xe9, 0x12, 0x8d, 0x27, 0xa2, 0xa5, 0x31,
	0xc4, 0xc0, 0x22, 0x85, 0x83, 0xdf, 0xcb, 0xab, 0xdf, 0xcd, 0x74, 0xd7, 0x42, 0xb8, 0xcd, 0xbc,
	0xf7, 0xdb, 0x97, 0x9d, 0x57, 0x68, 0x1b, 0x91, 0xad, 0x37, 0xc5, 0xb6, 0x40, 0x6a, 0x44, 0xfa,
	0x4b, 0x21, 0x96, 0x42, 0xdb, 0xaf, 0x9d, 0x75, 0x5b, 0xbc, 0x82, 0xc8, 0x14, 0xcb, 0xe5, 0xdb,
	0x6a, 0xce, 0x49, 0x9f, 0x0c, 0x12, 0x71, 0x96, 0x19, 0x91, 0xed, 0xfd, 0x4c, 0x56, 0xa6, 0x0e,
	0x14, 0x26, 0x40, 0x17, 0x73, 0x4e, 0xfb, 0x64, 0xd0, 0xd1, 0x74, 0x31, 0x47, 0x06, 0xf5, 0x4f,
	0xfb, 0xcd, 0xeb, 0xa5, 0xe0, 0x47, 0xec, 0x41, 0x73, 0x5b, 0xac, 0x17, 0x86, 0x37, 0xfa, 0x64,
	0x10, 0xeb, 0x6a, 0xf1, 0xdc, 0xd2, 0x7d, 0xf0, 0x66, 0xa9, 0xf9, 0x31, 0xfd, 0x21, 0x10, 0xc9,
	0x7d, 0x2a, 0x8c, 0xd5, 0xcb, 0x4c, 0x3e, 0x0e, 0xd5, 0x78, 0xca, 0x6a, 0x78, 0x02, 0x5d, 0xad,
	0x46, 0x93, 0x67, 0x15, 0x24, 0x82, 0x3d, 0x60, 0x1e, 0x99, 0x4e, 0x9e, 0x86, 0x32, 0xa8, 0x14,
	0xcf, 0xe1, 0xf4, 0x1f, 0x3c, 0x32, 0xea, 0xc8, 0xa0, 0xa3, 0x55, 0xae, 0xa6, 0x41, 0x69, 0x60,
	0x17, 0xe2, 0x7d, 0x00, 0x6b, 0x56, 0xc0, 0xe1, 0x4b, 0xd6, 0x42, 0x84, 0xe4, 0xf0, 0x13, 0xb3,
	0x07, 0xf5, 0xca, 0x22, 0x4f, 0xe5, 0x6a, 0x7c, 0x37, 0x1b, 0xa9, 0x3c, 0xbf, 0xbd, 0x57, 0xac,
	0x9d, 0xde, 0x00, 0xf8, 0x7a, 0xdc, 0xba, 0x58, 0x39, 0x8b, 0x1c, 0x22, 0xb7, 0x33, 0xc6, 0x3a,
	0x57, 0xf6, 0xd7, 0xd6, 0x61, 0xf5, 0x07, 0xdb, 0xcd, 0xa6, 0x6c, 0x2a, 0xd6, 0x7e, 0x14, 0x02,
	0xa8, 0x14, 0x78, 0xe9, 0xeb, 0x0f, 0x77, 0x77, 0x8f, 0xda, 0xbe, 0x48, 0xc2, 0x5a, 0xa5, 0xa7,
	0xb5, 0xf7, 0x56, 0xf9, 0x70, 0xd7, 0x7f, 0x01, 0x00, 0x00, 0xff, 0xff, 0x3b, 0xff, 0xc5, 0xfc,
	0xc4, 0x01, 0x00, 0x00,
}
