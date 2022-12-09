// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.2.0
// - protoc             v3.21.6
// source: grpc/interface.proto

package Bully

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

// PeerClient is the client API for Peer service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type PeerClient interface {
	Election(ctx context.Context, in *Elect, opts ...grpc.CallOption) (*Ack, error)
	Coordinate(ctx context.Context, in *Coord, opts ...grpc.CallOption) (*Ack, error)
}

type peerClient struct {
	cc grpc.ClientConnInterface
}

func NewPeerClient(cc grpc.ClientConnInterface) PeerClient {
	return &peerClient{cc}
}

func (c *peerClient) Election(ctx context.Context, in *Elect, opts ...grpc.CallOption) (*Ack, error) {
	out := new(Ack)
	err := c.cc.Invoke(ctx, "/Bully.Peer/Election", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *peerClient) Coordinate(ctx context.Context, in *Coord, opts ...grpc.CallOption) (*Ack, error) {
	out := new(Ack)
	err := c.cc.Invoke(ctx, "/Bully.Peer/Coordinate", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// PeerServer is the server API for Peer service.
// All implementations must embed UnimplementedPeerServer
// for forward compatibility
type PeerServer interface {
	Election(context.Context, *Elect) (*Ack, error)
	Coordinate(context.Context, *Coord) (*Ack, error)
	mustEmbedUnimplementedPeerServer()
}

// UnimplementedPeerServer must be embedded to have forward compatible implementations.
type UnimplementedPeerServer struct {
}

func (UnimplementedPeerServer) Election(context.Context, *Elect) (*Ack, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Election not implemented")
}
func (UnimplementedPeerServer) Coordinate(context.Context, *Coord) (*Ack, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Coordinate not implemented")
}
func (UnimplementedPeerServer) mustEmbedUnimplementedPeerServer() {}

// UnsafePeerServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to PeerServer will
// result in compilation errors.
type UnsafePeerServer interface {
	mustEmbedUnimplementedPeerServer()
}

func RegisterPeerServer(s grpc.ServiceRegistrar, srv PeerServer) {
	s.RegisterService(&Peer_ServiceDesc, srv)
}

func _Peer_Election_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(Elect)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(PeerServer).Election(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/Bully.Peer/Election",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(PeerServer).Election(ctx, req.(*Elect))
	}
	return interceptor(ctx, in, info, handler)
}

func _Peer_Coordinate_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(Coord)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(PeerServer).Coordinate(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/Bully.Peer/Coordinate",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(PeerServer).Coordinate(ctx, req.(*Coord))
	}
	return interceptor(ctx, in, info, handler)
}

// Peer_ServiceDesc is the grpc.ServiceDesc for Peer service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var Peer_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "Bully.Peer",
	HandlerType: (*PeerServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "Election",
			Handler:    _Peer_Election_Handler,
		},
		{
			MethodName: "Coordinate",
			Handler:    _Peer_Coordinate_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "grpc/interface.proto",
}