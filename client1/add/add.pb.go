// Code generated by protoc-gen-go.
// source: add.proto
// DO NOT EDIT!

/*
Package add is a generated protocol buffer package.

It is generated from these files:
	add.proto

It has these top-level messages:
	AddRequest
	AddResponse
*/
package add

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

type AddRequest struct {
	First  int32 `protobuf:"varint,1,opt,name=first" json:"first,omitempty"`
	Second int32 `protobuf:"varint,2,opt,name=second" json:"second,omitempty"`
}

func (m *AddRequest) Reset()                    { *m = AddRequest{} }
func (m *AddRequest) String() string            { return proto.CompactTextString(m) }
func (*AddRequest) ProtoMessage()               {}
func (*AddRequest) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{0} }

type AddResponse struct {
	Sum int32 `protobuf:"varint,1,opt,name=sum" json:"sum,omitempty"`
}

func (m *AddResponse) Reset()                    { *m = AddResponse{} }
func (m *AddResponse) String() string            { return proto.CompactTextString(m) }
func (*AddResponse) ProtoMessage()               {}
func (*AddResponse) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{1} }

func init() {
	proto.RegisterType((*AddRequest)(nil), "AddRequest")
	proto.RegisterType((*AddResponse)(nil), "AddResponse")
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion3

// Client API for AddService service

type AddServiceClient interface {
	Add(ctx context.Context, in *AddRequest, opts ...grpc.CallOption) (*AddResponse, error)
}

type addServiceClient struct {
	cc *grpc.ClientConn
}

func NewAddServiceClient(cc *grpc.ClientConn) AddServiceClient {
	return &addServiceClient{cc}
}

func (c *addServiceClient) Add(ctx context.Context, in *AddRequest, opts ...grpc.CallOption) (*AddResponse, error) {
	out := new(AddResponse)
	err := grpc.Invoke(ctx, "/AddService/Add", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// Server API for AddService service

type AddServiceServer interface {
	Add(context.Context, *AddRequest) (*AddResponse, error)
}

func RegisterAddServiceServer(s *grpc.Server, srv AddServiceServer) {
	s.RegisterService(&_AddService_serviceDesc, srv)
}

func _AddService_Add_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(AddRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(AddServiceServer).Add(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/AddService/Add",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(AddServiceServer).Add(ctx, req.(*AddRequest))
	}
	return interceptor(ctx, in, info, handler)
}

var _AddService_serviceDesc = grpc.ServiceDesc{
	ServiceName: "AddService",
	HandlerType: (*AddServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "Add",
			Handler:    _AddService_Add_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: fileDescriptor0,
}

func init() { proto.RegisterFile("add.proto", fileDescriptor0) }

var fileDescriptor0 = []byte{
	// 138 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x09, 0x6e, 0x88, 0x02, 0xff, 0xe2, 0xe2, 0x4c, 0x4c, 0x49, 0xd1,
	0x2b, 0x28, 0xca, 0x2f, 0xc9, 0x57, 0xb2, 0xe2, 0xe2, 0x72, 0x4c, 0x49, 0x09, 0x4a, 0x2d, 0x2c,
	0x4d, 0x2d, 0x2e, 0x11, 0x12, 0xe1, 0x62, 0x4d, 0xcb, 0x2c, 0x2a, 0x2e, 0x91, 0x60, 0x54, 0x60,
	0xd4, 0x60, 0x0d, 0x82, 0x70, 0x84, 0xc4, 0xb8, 0xd8, 0x8a, 0x53, 0x93, 0xf3, 0xf3, 0x52, 0x24,
	0x98, 0xc0, 0xc2, 0x50, 0x9e, 0x92, 0x3c, 0x17, 0x37, 0x58, 0x6f, 0x71, 0x41, 0x7e, 0x5e, 0x71,
	0xaa, 0x90, 0x00, 0x17, 0x73, 0x71, 0x69, 0x2e, 0x54, 0x2b, 0x88, 0x69, 0xa4, 0x07, 0x36, 0x3c,
	0x38, 0xb5, 0xa8, 0x2c, 0x33, 0x39, 0x55, 0x48, 0x81, 0x8b, 0x19, 0xc8, 0x13, 0xe2, 0xd6, 0x43,
	0x58, 0x28, 0xc5, 0xa3, 0x87, 0x64, 0x42, 0x12, 0x1b, 0xd8, 0x4d, 0xc6, 0x80, 0x00, 0x00, 0x00,
	0xff, 0xff, 0x72, 0x98, 0x70, 0x02, 0xa0, 0x00, 0x00, 0x00,
}
