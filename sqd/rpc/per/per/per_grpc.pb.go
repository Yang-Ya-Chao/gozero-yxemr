// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.2.0
// - protoc             v3.21.12
// source: per.proto

package per

import (
	context "context"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
)

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
// Requires gRPC-Go v1.32.0 or later.
const _ = grpc.SupportPackageIsVersion7

// PererClient is the client API for Perer service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type PererClient interface {
	Do(ctx context.Context, in *Req, opts ...grpc.CallOption) (*Resp, error)
}

type pererClient struct {
	cc grpc.ClientConnInterface
}

func NewPererClient(cc grpc.ClientConnInterface) PererClient {
	return &pererClient{cc}
}

func (c *pererClient) Do(ctx context.Context, in *Req, opts ...grpc.CallOption) (*Resp, error) {
	out := new(Resp)
	err := c.cc.Invoke(ctx, "/per.perer/Do", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// PererServer is the server API for Perer service.
// All implementations must embed UnimplementedPererServer
// for forward compatibility
type PererServer interface {
	Do(context.Context, *Req) (*Resp, error)
	mustEmbedUnimplementedPererServer()
}

// UnimplementedPererServer must be embedded to have forward compatible implementations.
type UnimplementedPererServer struct {
}

func (UnimplementedPererServer) Do(context.Context, *Req) (*Resp, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Do not implemented")
}
func (UnimplementedPererServer) mustEmbedUnimplementedPererServer() {}

// UnsafePererServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to PererServer will
// result in compilation errors.
type UnsafePererServer interface {
	mustEmbedUnimplementedPererServer()
}

func RegisterPererServer(s grpc.ServiceRegistrar, srv PererServer) {
	s.RegisterService(&Perer_ServiceDesc, srv)
}

func _Perer_Do_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(Req)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(PererServer).Do(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/per.perer/Do",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(PererServer).Do(ctx, req.(*Req))
	}
	return interceptor(ctx, in, info, handler)
}

// Perer_ServiceDesc is the grpc.ServiceDesc for Perer service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var Perer_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "per.perer",
	HandlerType: (*PererServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "Do",
			Handler:    _Perer_Do_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "per.proto",
}
