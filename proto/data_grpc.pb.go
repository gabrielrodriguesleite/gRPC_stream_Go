// Code generated by protoc-gen-go-grpc. DO NOT EDIT.

package data

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

// ServiceUpvoteClient is the client API for ServiceUpvote service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type ServiceUpvoteClient interface {
	MethodVote(ctx context.Context, in *RequestChose, opts ...grpc.CallOption) (ServiceUpvote_MethodVoteClient, error)
}

type serviceUpvoteClient struct {
	cc grpc.ClientConnInterface
}

func NewServiceUpvoteClient(cc grpc.ClientConnInterface) ServiceUpvoteClient {
	return &serviceUpvoteClient{cc}
}

func (c *serviceUpvoteClient) MethodVote(ctx context.Context, in *RequestChose, opts ...grpc.CallOption) (ServiceUpvote_MethodVoteClient, error) {
	stream, err := c.cc.NewStream(ctx, &ServiceUpvote_ServiceDesc.Streams[0], "/data.ServiceUpvote/MethodVote", opts...)
	if err != nil {
		return nil, err
	}
	x := &serviceUpvoteMethodVoteClient{stream}
	if err := x.ClientStream.SendMsg(in); err != nil {
		return nil, err
	}
	if err := x.ClientStream.CloseSend(); err != nil {
		return nil, err
	}
	return x, nil
}

type ServiceUpvote_MethodVoteClient interface {
	Recv() (*ResponseOption, error)
	grpc.ClientStream
}

type serviceUpvoteMethodVoteClient struct {
	grpc.ClientStream
}

func (x *serviceUpvoteMethodVoteClient) Recv() (*ResponseOption, error) {
	m := new(ResponseOption)
	if err := x.ClientStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

// ServiceUpvoteServer is the server API for ServiceUpvote service.
// All implementations must embed UnimplementedServiceUpvoteServer
// for forward compatibility
type ServiceUpvoteServer interface {
	MethodVote(*RequestChose, ServiceUpvote_MethodVoteServer) error
	mustEmbedUnimplementedServiceUpvoteServer()
}

// UnimplementedServiceUpvoteServer must be embedded to have forward compatible implementations.
type UnimplementedServiceUpvoteServer struct {
}

func (UnimplementedServiceUpvoteServer) MethodVote(*RequestChose, ServiceUpvote_MethodVoteServer) error {
	return status.Errorf(codes.Unimplemented, "method MethodVote not implemented")
}
func (UnimplementedServiceUpvoteServer) mustEmbedUnimplementedServiceUpvoteServer() {}

// UnsafeServiceUpvoteServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to ServiceUpvoteServer will
// result in compilation errors.
type UnsafeServiceUpvoteServer interface {
	mustEmbedUnimplementedServiceUpvoteServer()
}

func RegisterServiceUpvoteServer(s grpc.ServiceRegistrar, srv ServiceUpvoteServer) {
	s.RegisterService(&ServiceUpvote_ServiceDesc, srv)
}

func _ServiceUpvote_MethodVote_Handler(srv interface{}, stream grpc.ServerStream) error {
	m := new(RequestChose)
	if err := stream.RecvMsg(m); err != nil {
		return err
	}
	return srv.(ServiceUpvoteServer).MethodVote(m, &serviceUpvoteMethodVoteServer{stream})
}

type ServiceUpvote_MethodVoteServer interface {
	Send(*ResponseOption) error
	grpc.ServerStream
}

type serviceUpvoteMethodVoteServer struct {
	grpc.ServerStream
}

func (x *serviceUpvoteMethodVoteServer) Send(m *ResponseOption) error {
	return x.ServerStream.SendMsg(m)
}

// ServiceUpvote_ServiceDesc is the grpc.ServiceDesc for ServiceUpvote service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var ServiceUpvote_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "data.ServiceUpvote",
	HandlerType: (*ServiceUpvoteServer)(nil),
	Methods:     []grpc.MethodDesc{},
	Streams: []grpc.StreamDesc{
		{
			StreamName:    "MethodVote",
			Handler:       _ServiceUpvote_MethodVote_Handler,
			ServerStreams: true,
		},
	},
	Metadata: "proto/data.proto",
}
