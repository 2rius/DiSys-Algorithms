// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.2.0
// - protoc             v3.21.6
// source: api/interface.proto

package raft

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
	Heartbeat(ctx context.Context, in *AppendRequest, opts ...grpc.CallOption) (*AppendReply, error)
	RequestVote(ctx context.Context, in *VoteRequest, opts ...grpc.CallOption) (*VoteReply, error)
	AppendEntries(ctx context.Context, in *AppendRequest, opts ...grpc.CallOption) (*AppendReply, error)
}

type peerClient struct {
	cc grpc.ClientConnInterface
}

func NewPeerClient(cc grpc.ClientConnInterface) PeerClient {
	return &peerClient{cc}
}

func (c *peerClient) Heartbeat(ctx context.Context, in *AppendRequest, opts ...grpc.CallOption) (*AppendReply, error) {
	out := new(AppendReply)
	err := c.cc.Invoke(ctx, "/raft.Peer/Heartbeat", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *peerClient) RequestVote(ctx context.Context, in *VoteRequest, opts ...grpc.CallOption) (*VoteReply, error) {
	out := new(VoteReply)
	err := c.cc.Invoke(ctx, "/raft.Peer/RequestVote", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *peerClient) AppendEntries(ctx context.Context, in *AppendRequest, opts ...grpc.CallOption) (*AppendReply, error) {
	out := new(AppendReply)
	err := c.cc.Invoke(ctx, "/raft.Peer/AppendEntries", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// PeerServer is the server API for Peer service.
// All implementations must embed UnimplementedPeerServer
// for forward compatibility
type PeerServer interface {
	Heartbeat(context.Context, *AppendRequest) (*AppendReply, error)
	RequestVote(context.Context, *VoteRequest) (*VoteReply, error)
	AppendEntries(context.Context, *AppendRequest) (*AppendReply, error)
	mustEmbedUnimplementedPeerServer()
}

// UnimplementedPeerServer must be embedded to have forward compatible implementations.
type UnimplementedPeerServer struct {
}

func (UnimplementedPeerServer) Heartbeat(context.Context, *AppendRequest) (*AppendReply, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Heartbeat not implemented")
}
func (UnimplementedPeerServer) RequestVote(context.Context, *VoteRequest) (*VoteReply, error) {
	return nil, status.Errorf(codes.Unimplemented, "method RequestVote not implemented")
}
func (UnimplementedPeerServer) AppendEntries(context.Context, *AppendRequest) (*AppendReply, error) {
	return nil, status.Errorf(codes.Unimplemented, "method AppendEntries not implemented")
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

func _Peer_Heartbeat_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(AppendRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(PeerServer).Heartbeat(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/raft.Peer/Heartbeat",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(PeerServer).Heartbeat(ctx, req.(*AppendRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Peer_RequestVote_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(VoteRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(PeerServer).RequestVote(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/raft.Peer/RequestVote",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(PeerServer).RequestVote(ctx, req.(*VoteRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Peer_AppendEntries_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(AppendRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(PeerServer).AppendEntries(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/raft.Peer/AppendEntries",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(PeerServer).AppendEntries(ctx, req.(*AppendRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// Peer_ServiceDesc is the grpc.ServiceDesc for Peer service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var Peer_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "raft.Peer",
	HandlerType: (*PeerServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "Heartbeat",
			Handler:    _Peer_Heartbeat_Handler,
		},
		{
			MethodName: "RequestVote",
			Handler:    _Peer_RequestVote_Handler,
		},
		{
			MethodName: "AppendEntries",
			Handler:    _Peer_AppendEntries_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "api/interface.proto",
}
