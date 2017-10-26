// Code generated by protoc-gen-go. DO NOT EDIT.
// source: github.com/ewanvalentine/gateway-test/proto/greeter/greeter.proto

/*
Package greeter is a generated protocol buffer package.

It is generated from these files:
	github.com/ewanvalentine/gateway-test/proto/greeter/greeter.proto

It has these top-level messages:
	HelloRequest
	HelloResponse
*/
package greeter

import (
	fmt "fmt"
	"log"

	proto "github.com/golang/protobuf/proto"

	math "math"

	_ "google.golang.org/genproto/googleapis/api/annotations"

	client "github.com/micro/go-micro/client"

	server "github.com/micro/go-micro/server"

	context "golang.org/x/net/context"
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

type HelloRequest struct {
	Name string `protobuf:"bytes,1,opt,name=name" json:"name,omitempty"`
}

func (m *HelloRequest) Reset()                    { *m = HelloRequest{} }
func (m *HelloRequest) String() string            { return proto.CompactTextString(m) }
func (*HelloRequest) ProtoMessage()               {}
func (*HelloRequest) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{0} }

func (m *HelloRequest) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

type HelloResponse struct {
	Greeting string `protobuf:"bytes,2,opt,name=greeting" json:"greeting,omitempty"`
}

func (m *HelloResponse) Reset()                    { *m = HelloResponse{} }
func (m *HelloResponse) String() string            { return proto.CompactTextString(m) }
func (*HelloResponse) ProtoMessage()               {}
func (*HelloResponse) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{1} }

func (m *HelloResponse) GetGreeting() string {
	if m != nil {
		return m.Greeting
	}
	return ""
}

func init() {
	proto.RegisterType((*HelloRequest)(nil), "greeter.HelloRequest")
	proto.RegisterType((*HelloResponse)(nil), "greeter.HelloResponse")
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ client.Option
var _ server.Option

// Client API for Greeter service

type GreeterClient interface {
	Hello(ctx context.Context, in *HelloRequest, opts ...client.CallOption) (*HelloResponse, error)
}

type greeterClient struct {
	c           client.Client
	serviceName string
}

func NewGreeterClient(serviceName string, c client.Client) GreeterClient {
	if c == nil {
		c = client.NewClient()
	}
	if len(serviceName) == 0 {
		serviceName = "greeter"
	}
	return &greeterClient{
		c:           c,
		serviceName: serviceName,
	}
}

func (c *greeterClient) Hello(ctx context.Context, in *HelloRequest, opts ...client.CallOption) (*HelloResponse, error) {
	log.Println("Shit")
	req := c.c.NewRequest(c.serviceName, "Greeter.Hello", in)
	out := new(HelloResponse)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// Server API for Greeter service

type GreeterHandler interface {
	Hello(context.Context, *HelloRequest, *HelloResponse) error
}

func RegisterGreeterHandler(s server.Server, hdlr GreeterHandler, opts ...server.HandlerOption) {
	s.Handle(s.NewHandler(&Greeter{hdlr}, opts...))
}

type Greeter struct {
	GreeterHandler
}

func (h *Greeter) Hello(ctx context.Context, in *HelloRequest, out *HelloResponse) error {
	return h.GreeterHandler.Hello(ctx, in, out)
}

func init() {
	proto.RegisterFile("github.com/ewanvalentine/gateway-test/proto/greeter/greeter.proto", fileDescriptor0)
}

var fileDescriptor0 = []byte{
	// 227 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xe2, 0x72, 0x4c, 0xcf, 0x2c, 0xc9,
	0x28, 0x4d, 0xd2, 0x4b, 0xce, 0xcf, 0xd5, 0x4f, 0x2d, 0x4f, 0xcc, 0x2b, 0x4b, 0xcc, 0x49, 0xcd,
	0x2b, 0xc9, 0xcc, 0x4b, 0xd5, 0x4f, 0x4f, 0x2c, 0x49, 0x2d, 0x4f, 0xac, 0xd4, 0x2d, 0x49, 0x2d,
	0x2e, 0xd1, 0x2f, 0x28, 0xca, 0x2f, 0xc9, 0xd7, 0x4f, 0x2f, 0x4a, 0x4d, 0x2d, 0x49, 0x2d, 0x82,
	0xd1, 0x7a, 0x60, 0x51, 0x21, 0x76, 0x28, 0x57, 0x4a, 0x26, 0x3d, 0x3f, 0x3f, 0x3d, 0x27, 0x55,
	0x3f, 0xb1, 0x20, 0x53, 0x3f, 0x31, 0x2f, 0x2f, 0xbf, 0x24, 0xb1, 0x24, 0x33, 0x3f, 0xaf, 0x18,
	0xa2, 0x4c, 0x49, 0x89, 0x8b, 0xc7, 0x23, 0x35, 0x27, 0x27, 0x3f, 0x28, 0xb5, 0xb0, 0x34, 0xb5,
	0xb8, 0x44, 0x48, 0x88, 0x8b, 0x25, 0x2f, 0x31, 0x37, 0x55, 0x82, 0x51, 0x81, 0x51, 0x83, 0x33,
	0x08, 0xcc, 0x56, 0xd2, 0xe6, 0xe2, 0x85, 0xaa, 0x29, 0x2e, 0xc8, 0xcf, 0x2b, 0x4e, 0x15, 0x92,
	0xe2, 0xe2, 0x00, 0x9b, 0x9e, 0x99, 0x97, 0x2e, 0xc1, 0x04, 0x56, 0x08, 0xe7, 0x1b, 0xc5, 0x73,
	0xb1, 0xbb, 0x43, 0x6c, 0x16, 0x0a, 0xe1, 0x62, 0x05, 0xeb, 0x13, 0x12, 0xd5, 0x83, 0xb9, 0x0d,
	0xd9, 0x2e, 0x29, 0x31, 0x74, 0x61, 0x88, 0xf1, 0x4a, 0x32, 0x4d, 0x97, 0x9f, 0x4c, 0x66, 0x12,
	0x13, 0x12, 0x01, 0xbb, 0xb9, 0xcc, 0x50, 0x3f, 0x03, 0x24, 0xad, 0x5f, 0x0d, 0x72, 0x4c, 0x6d,
	0x12, 0x1b, 0xd8, 0xe1, 0xc6, 0x80, 0x00, 0x00, 0x00, 0xff, 0xff, 0x8d, 0x02, 0x9f, 0xc7, 0x24,
	0x01, 0x00, 0x00,
}
