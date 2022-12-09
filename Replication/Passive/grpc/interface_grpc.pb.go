// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.2.0
// - protoc             v3.21.6
// source: grpc/interface.proto

package Passive

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

// ManagerClient is the client API for Manager service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type ManagerClient interface {
	Heartbeat(ctx context.Context, in *Void, opts ...grpc.CallOption) (*Void, error)
	Update(ctx context.Context, in *Data, opts ...grpc.CallOption) (*Void, error)
	Set(ctx context.Context, in *Data, opts ...grpc.CallOption) (*Void, error)
	Get(ctx context.Context, in *Void, opts ...grpc.CallOption) (*Data, error)
	Election(ctx context.Context, in *Elect, opts ...grpc.CallOption) (*Void, error)
	Coordinate(ctx context.Context, in *Coord, opts ...grpc.CallOption) (*Void, error)
}

type managerClient struct {
	cc grpc.ClientConnInterface
}

func NewManagerClient(cc grpc.ClientConnInterface) ManagerClient {
	return &managerClient{cc}
}

func (c *managerClient) Heartbeat(ctx context.Context, in *Void, opts ...grpc.CallOption) (*Void, error) {
	out := new(Void)
	err := c.cc.Invoke(ctx, "/manager.Manager/Heartbeat", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *managerClient) Update(ctx context.Context, in *Data, opts ...grpc.CallOption) (*Void, error) {
	out := new(Void)
	err := c.cc.Invoke(ctx, "/manager.Manager/Update", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *managerClient) Set(ctx context.Context, in *Data, opts ...grpc.CallOption) (*Void, error) {
	out := new(Void)
	err := c.cc.Invoke(ctx, "/manager.Manager/Set", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *managerClient) Get(ctx context.Context, in *Void, opts ...grpc.CallOption) (*Data, error) {
	out := new(Data)
	err := c.cc.Invoke(ctx, "/manager.Manager/Get", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *managerClient) Election(ctx context.Context, in *Elect, opts ...grpc.CallOption) (*Void, error) {
	out := new(Void)
	err := c.cc.Invoke(ctx, "/manager.Manager/Election", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *managerClient) Coordinate(ctx context.Context, in *Coord, opts ...grpc.CallOption) (*Void, error) {
	out := new(Void)
	err := c.cc.Invoke(ctx, "/manager.Manager/Coordinate", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// ManagerServer is the server API for Manager service.
// All implementations must embed UnimplementedManagerServer
// for forward compatibility
type ManagerServer interface {
	Heartbeat(context.Context, *Void) (*Void, error)
	Update(context.Context, *Data) (*Void, error)
	Set(context.Context, *Data) (*Void, error)
	Get(context.Context, *Void) (*Data, error)
	Election(context.Context, *Elect) (*Void, error)
	Coordinate(context.Context, *Coord) (*Void, error)
	mustEmbedUnimplementedManagerServer()
}

// UnimplementedManagerServer must be embedded to have forward compatible implementations.
type UnimplementedManagerServer struct {
}

func (UnimplementedManagerServer) Heartbeat(context.Context, *Void) (*Void, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Heartbeat not implemented")
}
func (UnimplementedManagerServer) Update(context.Context, *Data) (*Void, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Update not implemented")
}
func (UnimplementedManagerServer) Set(context.Context, *Data) (*Void, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Set not implemented")
}
func (UnimplementedManagerServer) Get(context.Context, *Void) (*Data, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Get not implemented")
}
func (UnimplementedManagerServer) Election(context.Context, *Elect) (*Void, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Election not implemented")
}
func (UnimplementedManagerServer) Coordinate(context.Context, *Coord) (*Void, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Coordinate not implemented")
}
func (UnimplementedManagerServer) mustEmbedUnimplementedManagerServer() {}

// UnsafeManagerServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to ManagerServer will
// result in compilation errors.
type UnsafeManagerServer interface {
	mustEmbedUnimplementedManagerServer()
}

func RegisterManagerServer(s grpc.ServiceRegistrar, srv ManagerServer) {
	s.RegisterService(&Manager_ServiceDesc, srv)
}

func _Manager_Heartbeat_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(Void)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ManagerServer).Heartbeat(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/manager.Manager/Heartbeat",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ManagerServer).Heartbeat(ctx, req.(*Void))
	}
	return interceptor(ctx, in, info, handler)
}

func _Manager_Update_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(Data)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ManagerServer).Update(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/manager.Manager/Update",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ManagerServer).Update(ctx, req.(*Data))
	}
	return interceptor(ctx, in, info, handler)
}

func _Manager_Set_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(Data)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ManagerServer).Set(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/manager.Manager/Set",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ManagerServer).Set(ctx, req.(*Data))
	}
	return interceptor(ctx, in, info, handler)
}

func _Manager_Get_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(Void)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ManagerServer).Get(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/manager.Manager/Get",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ManagerServer).Get(ctx, req.(*Void))
	}
	return interceptor(ctx, in, info, handler)
}

func _Manager_Election_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(Elect)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ManagerServer).Election(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/manager.Manager/Election",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ManagerServer).Election(ctx, req.(*Elect))
	}
	return interceptor(ctx, in, info, handler)
}

func _Manager_Coordinate_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(Coord)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ManagerServer).Coordinate(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/manager.Manager/Coordinate",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ManagerServer).Coordinate(ctx, req.(*Coord))
	}
	return interceptor(ctx, in, info, handler)
}

// Manager_ServiceDesc is the grpc.ServiceDesc for Manager service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var Manager_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "manager.Manager",
	HandlerType: (*ManagerServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "Heartbeat",
			Handler:    _Manager_Heartbeat_Handler,
		},
		{
			MethodName: "Update",
			Handler:    _Manager_Update_Handler,
		},
		{
			MethodName: "Set",
			Handler:    _Manager_Set_Handler,
		},
		{
			MethodName: "Get",
			Handler:    _Manager_Get_Handler,
		},
		{
			MethodName: "Election",
			Handler:    _Manager_Election_Handler,
		},
		{
			MethodName: "Coordinate",
			Handler:    _Manager_Coordinate_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "grpc/interface.proto",
}

// FrontendClient is the client API for Frontend service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type FrontendClient interface {
	Heartbeat(ctx context.Context, in *Primary, opts ...grpc.CallOption) (*Void, error)
}

type frontendClient struct {
	cc grpc.ClientConnInterface
}

func NewFrontendClient(cc grpc.ClientConnInterface) FrontendClient {
	return &frontendClient{cc}
}

func (c *frontendClient) Heartbeat(ctx context.Context, in *Primary, opts ...grpc.CallOption) (*Void, error) {
	out := new(Void)
	err := c.cc.Invoke(ctx, "/manager.Frontend/Heartbeat", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// FrontendServer is the server API for Frontend service.
// All implementations must embed UnimplementedFrontendServer
// for forward compatibility
type FrontendServer interface {
	Heartbeat(context.Context, *Primary) (*Void, error)
	mustEmbedUnimplementedFrontendServer()
}

// UnimplementedFrontendServer must be embedded to have forward compatible implementations.
type UnimplementedFrontendServer struct {
}

func (UnimplementedFrontendServer) Heartbeat(context.Context, *Primary) (*Void, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Heartbeat not implemented")
}
func (UnimplementedFrontendServer) mustEmbedUnimplementedFrontendServer() {}

// UnsafeFrontendServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to FrontendServer will
// result in compilation errors.
type UnsafeFrontendServer interface {
	mustEmbedUnimplementedFrontendServer()
}

func RegisterFrontendServer(s grpc.ServiceRegistrar, srv FrontendServer) {
	s.RegisterService(&Frontend_ServiceDesc, srv)
}

func _Frontend_Heartbeat_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(Primary)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(FrontendServer).Heartbeat(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/manager.Frontend/Heartbeat",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(FrontendServer).Heartbeat(ctx, req.(*Primary))
	}
	return interceptor(ctx, in, info, handler)
}

// Frontend_ServiceDesc is the grpc.ServiceDesc for Frontend service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var Frontend_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "manager.Frontend",
	HandlerType: (*FrontendServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "Heartbeat",
			Handler:    _Frontend_Heartbeat_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "grpc/interface.proto",
}
