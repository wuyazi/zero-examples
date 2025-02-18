// Code generated by protoc-gen-go.
// source: transform.proto
// DO NOT EDIT!

/*
Package transform is a generated protocol buffer package.

It is generated from these files:
	transform.proto

It has these top-level messages:
	ExpandReq
	ExpandResp
	ShortenReq
	ShortenResp
*/
package transform

import (
	proto "github.com/golang/protobuf/proto"
	fmt "fmt"
	math "math"
)

import (
	context "golang.org/x/net/context"
	grpc "google.golang.org/grpc"
)

// Reference imports to suppress errors if they are not otherwise used.
var (
	_ = proto.Marshal
	_ = fmt.Errorf
	_ = math.Inf
)

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.ProtoPackageIsVersion2 // please upgrade the proto package

type ExpandReq struct {
	Shorten string `protobuf:"bytes,1,opt,name=shorten" json:"shorten,omitempty"`
}

func (m *ExpandReq) Reset()                    { *m = ExpandReq{} }
func (m *ExpandReq) String() string            { return proto.CompactTextString(m) }
func (*ExpandReq) ProtoMessage()               {}
func (*ExpandReq) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{0} }

func (m *ExpandReq) GetShorten() string {
	if m != nil {
		return m.Shorten
	}
	return ""
}

type ExpandResp struct {
	Url string `protobuf:"bytes,1,opt,name=url" json:"url,omitempty"`
}

func (m *ExpandResp) Reset()                    { *m = ExpandResp{} }
func (m *ExpandResp) String() string            { return proto.CompactTextString(m) }
func (*ExpandResp) ProtoMessage()               {}
func (*ExpandResp) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{1} }

func (m *ExpandResp) GetUrl() string {
	if m != nil {
		return m.Url
	}
	return ""
}

type ShortenReq struct {
	Url string `protobuf:"bytes,1,opt,name=url" json:"url,omitempty"`
}

func (m *ShortenReq) Reset()                    { *m = ShortenReq{} }
func (m *ShortenReq) String() string            { return proto.CompactTextString(m) }
func (*ShortenReq) ProtoMessage()               {}
func (*ShortenReq) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{2} }

func (m *ShortenReq) GetUrl() string {
	if m != nil {
		return m.Url
	}
	return ""
}

type ShortenResp struct {
	Shorten string `protobuf:"bytes,1,opt,name=shorten" json:"shorten,omitempty"`
}

func (m *ShortenResp) Reset()                    { *m = ShortenResp{} }
func (m *ShortenResp) String() string            { return proto.CompactTextString(m) }
func (*ShortenResp) ProtoMessage()               {}
func (*ShortenResp) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{3} }

func (m *ShortenResp) GetShorten() string {
	if m != nil {
		return m.Shorten
	}
	return ""
}

func init() {
	proto.RegisterType((*ExpandReq)(nil), "transform.expandReq")
	proto.RegisterType((*ExpandResp)(nil), "transform.expandResp")
	proto.RegisterType((*ShortenReq)(nil), "transform.shortenReq")
	proto.RegisterType((*ShortenResp)(nil), "transform.shortenResp")
}

// Reference imports to suppress errors if they are not otherwise used.
var (
	_ context.Context
	_ grpc.ClientConn
)

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion4

// Client API for Transformer service

type TransformerClient interface {
	Expand(ctx context.Context, in *ExpandReq, opts ...grpc.CallOption) (*ExpandResp, error)
	Shorten(ctx context.Context, in *ShortenReq, opts ...grpc.CallOption) (*ShortenResp, error)
}

type transformerClient struct {
	cc *grpc.ClientConn
}

func NewTransformerClient(cc *grpc.ClientConn) TransformerClient {
	return &transformerClient{cc}
}

func (c *transformerClient) Expand(ctx context.Context, in *ExpandReq, opts ...grpc.CallOption) (*ExpandResp, error) {
	out := new(ExpandResp)
	err := grpc.Invoke(ctx, "/transform.transformer/expand", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *transformerClient) Shorten(ctx context.Context, in *ShortenReq, opts ...grpc.CallOption) (*ShortenResp, error) {
	out := new(ShortenResp)
	err := grpc.Invoke(ctx, "/transform.transformer/shorten", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// Server API for Transformer service

type TransformerServer interface {
	Expand(context.Context, *ExpandReq) (*ExpandResp, error)
	Shorten(context.Context, *ShortenReq) (*ShortenResp, error)
}

func RegisterTransformerServer(s *grpc.Server, srv TransformerServer) {
	s.RegisterService(&_Transformer_serviceDesc, srv)
}

func _Transformer_Expand_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ExpandReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(TransformerServer).Expand(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/transform.transformer/Expand",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(TransformerServer).Expand(ctx, req.(*ExpandReq))
	}
	return interceptor(ctx, in, info, handler)
}

func _Transformer_Shorten_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ShortenReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(TransformerServer).Shorten(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/transform.transformer/Shorten",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(TransformerServer).Shorten(ctx, req.(*ShortenReq))
	}
	return interceptor(ctx, in, info, handler)
}

var _Transformer_serviceDesc = grpc.ServiceDesc{
	ServiceName: "transform.transformer",
	HandlerType: (*TransformerServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "expand",
			Handler:    _Transformer_Expand_Handler,
		},
		{
			MethodName: "shorten",
			Handler:    _Transformer_Shorten_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "transform.proto",
}

func init() { proto.RegisterFile("transform.proto", fileDescriptor0) }

var fileDescriptor0 = []byte{
	// 163 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x09, 0x6e, 0x88, 0x02, 0xff, 0xe2, 0xe2, 0x2f, 0x29, 0x4a, 0xcc,
	0x2b, 0x4e, 0xcb, 0x2f, 0xca, 0xd5, 0x2b, 0x28, 0xca, 0x2f, 0xc9, 0x17, 0xe2, 0x84, 0x0b, 0x28,
	0xa9, 0x72, 0x71, 0xa6, 0x56, 0x14, 0x24, 0xe6, 0xa5, 0x04, 0xa5, 0x16, 0x0a, 0x49, 0x70, 0xb1,
	0x17, 0x67, 0xe4, 0x17, 0x95, 0xa4, 0xe6, 0x49, 0x30, 0x2a, 0x30, 0x6a, 0x70, 0x06, 0xc1, 0xb8,
	0x4a, 0x72, 0x5c, 0x5c, 0x30, 0x65, 0xc5, 0x05, 0x42, 0x02, 0x5c, 0xcc, 0xa5, 0x45, 0x39, 0x50,
	0x35, 0x20, 0x26, 0x48, 0x1e, 0xaa, 0x14, 0x64, 0x0e, 0xa6, 0xbc, 0x3a, 0x17, 0x37, 0x5c, 0xbe,
	0xb8, 0x00, 0xb7, 0x45, 0x46, 0x75, 0x5c, 0xdc, 0x70, 0xc7, 0xa5, 0x16, 0x09, 0x99, 0x72, 0xb1,
	0x41, 0xec, 0x15, 0x12, 0xd1, 0x43, 0xf8, 0x02, 0xee, 0x62, 0x29, 0x51, 0x2c, 0xa2, 0xc5, 0x05,
	0x42, 0x16, 0x70, 0xf3, 0x85, 0x90, 0x55, 0x20, 0x9c, 0x28, 0x25, 0x86, 0x4d, 0xb8, 0xb8, 0x20,
	0x89, 0x0d, 0x1c, 0x42, 0xc6, 0x80, 0x00, 0x00, 0x00, 0xff, 0xff, 0x8a, 0x27, 0x78, 0x05, 0x34,
	0x01, 0x00, 0x00,
}
