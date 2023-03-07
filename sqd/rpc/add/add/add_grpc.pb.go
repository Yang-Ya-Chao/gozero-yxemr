// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.2.0
// - protoc             v3.21.12
// source: add.proto

package add

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

// AdderClient is the client API for Adder service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type AdderClient interface {
	Do(ctx context.Context, in *Req, opts ...grpc.CallOption) (*Resp, error)
}

type adderClient struct {
	cc grpc.ClientConnInterface
}

func NewAdderClient(cc grpc.ClientConnInterface) AdderClient {
	return &adderClient{cc}
}

func (c *adderClient) Do(ctx context.Context, in *Req, opts ...grpc.CallOption) (*Resp, error) {
	out := new(Resp)
	err := c.cc.Invoke(ctx, "/add.adder/Do", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// AdderServer is the server API for Adder service.
// All implementations must embed UnimplementedAdderServer
// for forward compatibility
type AdderServer interface {
	Do(context.Context, *Req) (*Resp, error)
	mustEmbedUnimplementedAdderServer()
}

// UnimplementedAdderServer must be embedded to have forward compatible implementations.
type UnimplementedAdderServer struct {
}

func (UnimplementedAdderServer) Do(context.Context, *Req) (*Resp, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Do not implemented")
}
func (UnimplementedAdderServer) mustEmbedUnimplementedAdderServer() {}

// UnsafeAdderServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to AdderServer will
// result in compilation errors.
type UnsafeAdderServer interface {
	mustEmbedUnimplementedAdderServer()
}

func RegisterAdderServer(s grpc.ServiceRegistrar, srv AdderServer) {
	s.RegisterService(&Adder_ServiceDesc, srv)
}

func _Adder_Do_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(Req)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(AdderServer).Do(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/add.adder/Do",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(AdderServer).Do(ctx, req.(*Req))
	}
	return interceptor(ctx, in, info, handler)
}

// Adder_ServiceDesc is the grpc.ServiceDesc for Adder service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var Adder_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "add.adder",
	HandlerType: (*AdderServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "Do",
			Handler:    _Adder_Do_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "add.proto",
}