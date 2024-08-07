// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.4.0
// - protoc             (unknown)
// source: kessel/inventory/v1beta1/rhel_hosts_service.proto

package v1beta1

import (
	context "context"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
)

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
// Requires gRPC-Go v1.62.0 or later.
const _ = grpc.SupportPackageIsVersion8

const (
	KesselRhelHostService_CreateRhelHost_FullMethodName = "/kessel.inventory.v1beta1.KesselRhelHostService/CreateRhelHost"
	KesselRhelHostService_UpdateRhelHost_FullMethodName = "/kessel.inventory.v1beta1.KesselRhelHostService/UpdateRhelHost"
	KesselRhelHostService_DeleteRhelHost_FullMethodName = "/kessel.inventory.v1beta1.KesselRhelHostService/DeleteRhelHost"
)

// KesselRhelHostServiceClient is the client API for KesselRhelHostService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type KesselRhelHostServiceClient interface {
	CreateRhelHost(ctx context.Context, in *CreateRhelHostRequest, opts ...grpc.CallOption) (*CreateRhelHostResponse, error)
	UpdateRhelHost(ctx context.Context, in *UpdateRhelHostRequest, opts ...grpc.CallOption) (*UpdateRhelHostResponse, error)
	DeleteRhelHost(ctx context.Context, in *DeleteRhelHostRequest, opts ...grpc.CallOption) (*DeleteRhelHostResponse, error)
}

type kesselRhelHostServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewKesselRhelHostServiceClient(cc grpc.ClientConnInterface) KesselRhelHostServiceClient {
	return &kesselRhelHostServiceClient{cc}
}

func (c *kesselRhelHostServiceClient) CreateRhelHost(ctx context.Context, in *CreateRhelHostRequest, opts ...grpc.CallOption) (*CreateRhelHostResponse, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(CreateRhelHostResponse)
	err := c.cc.Invoke(ctx, KesselRhelHostService_CreateRhelHost_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *kesselRhelHostServiceClient) UpdateRhelHost(ctx context.Context, in *UpdateRhelHostRequest, opts ...grpc.CallOption) (*UpdateRhelHostResponse, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(UpdateRhelHostResponse)
	err := c.cc.Invoke(ctx, KesselRhelHostService_UpdateRhelHost_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *kesselRhelHostServiceClient) DeleteRhelHost(ctx context.Context, in *DeleteRhelHostRequest, opts ...grpc.CallOption) (*DeleteRhelHostResponse, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(DeleteRhelHostResponse)
	err := c.cc.Invoke(ctx, KesselRhelHostService_DeleteRhelHost_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// KesselRhelHostServiceServer is the server API for KesselRhelHostService service.
// All implementations must embed UnimplementedKesselRhelHostServiceServer
// for forward compatibility
type KesselRhelHostServiceServer interface {
	CreateRhelHost(context.Context, *CreateRhelHostRequest) (*CreateRhelHostResponse, error)
	UpdateRhelHost(context.Context, *UpdateRhelHostRequest) (*UpdateRhelHostResponse, error)
	DeleteRhelHost(context.Context, *DeleteRhelHostRequest) (*DeleteRhelHostResponse, error)
	mustEmbedUnimplementedKesselRhelHostServiceServer()
}

// UnimplementedKesselRhelHostServiceServer must be embedded to have forward compatible implementations.
type UnimplementedKesselRhelHostServiceServer struct {
}

func (UnimplementedKesselRhelHostServiceServer) CreateRhelHost(context.Context, *CreateRhelHostRequest) (*CreateRhelHostResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method CreateRhelHost not implemented")
}
func (UnimplementedKesselRhelHostServiceServer) UpdateRhelHost(context.Context, *UpdateRhelHostRequest) (*UpdateRhelHostResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method UpdateRhelHost not implemented")
}
func (UnimplementedKesselRhelHostServiceServer) DeleteRhelHost(context.Context, *DeleteRhelHostRequest) (*DeleteRhelHostResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method DeleteRhelHost not implemented")
}
func (UnimplementedKesselRhelHostServiceServer) mustEmbedUnimplementedKesselRhelHostServiceServer() {}

// UnsafeKesselRhelHostServiceServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to KesselRhelHostServiceServer will
// result in compilation errors.
type UnsafeKesselRhelHostServiceServer interface {
	mustEmbedUnimplementedKesselRhelHostServiceServer()
}

func RegisterKesselRhelHostServiceServer(s grpc.ServiceRegistrar, srv KesselRhelHostServiceServer) {
	s.RegisterService(&KesselRhelHostService_ServiceDesc, srv)
}

func _KesselRhelHostService_CreateRhelHost_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CreateRhelHostRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(KesselRhelHostServiceServer).CreateRhelHost(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: KesselRhelHostService_CreateRhelHost_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(KesselRhelHostServiceServer).CreateRhelHost(ctx, req.(*CreateRhelHostRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _KesselRhelHostService_UpdateRhelHost_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(UpdateRhelHostRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(KesselRhelHostServiceServer).UpdateRhelHost(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: KesselRhelHostService_UpdateRhelHost_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(KesselRhelHostServiceServer).UpdateRhelHost(ctx, req.(*UpdateRhelHostRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _KesselRhelHostService_DeleteRhelHost_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(DeleteRhelHostRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(KesselRhelHostServiceServer).DeleteRhelHost(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: KesselRhelHostService_DeleteRhelHost_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(KesselRhelHostServiceServer).DeleteRhelHost(ctx, req.(*DeleteRhelHostRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// KesselRhelHostService_ServiceDesc is the grpc.ServiceDesc for KesselRhelHostService service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var KesselRhelHostService_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "kessel.inventory.v1beta1.KesselRhelHostService",
	HandlerType: (*KesselRhelHostServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "CreateRhelHost",
			Handler:    _KesselRhelHostService_CreateRhelHost_Handler,
		},
		{
			MethodName: "UpdateRhelHost",
			Handler:    _KesselRhelHostService_UpdateRhelHost_Handler,
		},
		{
			MethodName: "DeleteRhelHost",
			Handler:    _KesselRhelHostService_DeleteRhelHost_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "kessel/inventory/v1beta1/rhel_hosts_service.proto",
}