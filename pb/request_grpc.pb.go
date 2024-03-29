// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.2.0
// - protoc             v3.19.4
// source: request.proto

package pb

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

// NodeClient is the client API for Node service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type NodeClient interface {
	Receive(ctx context.Context, in *Data, opts ...grpc.CallOption) (*Response, error)
	Delete(ctx context.Context, in *Metadata, opts ...grpc.CallOption) (*Response, error)
	Healthcheck(ctx context.Context, opts ...grpc.CallOption) (Node_HealthcheckClient, error)
}

type nodeClient struct {
	cc grpc.ClientConnInterface
}

func NewNodeClient(cc grpc.ClientConnInterface) NodeClient {
	return &nodeClient{cc}
}

func (c *nodeClient) Receive(ctx context.Context, in *Data, opts ...grpc.CallOption) (*Response, error) {
	out := new(Response)
	err := c.cc.Invoke(ctx, "/Node/Receive", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *nodeClient) Delete(ctx context.Context, in *Metadata, opts ...grpc.CallOption) (*Response, error) {
	out := new(Response)
	err := c.cc.Invoke(ctx, "/Node/Delete", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *nodeClient) Healthcheck(ctx context.Context, opts ...grpc.CallOption) (Node_HealthcheckClient, error) {
	stream, err := c.cc.NewStream(ctx, &Node_ServiceDesc.Streams[0], "/Node/Healthcheck", opts...)
	if err != nil {
		return nil, err
	}
	x := &nodeHealthcheckClient{stream}
	return x, nil
}

type Node_HealthcheckClient interface {
	Send(*HealthcheckMetadata) error
	CloseAndRecv() (*Empty, error)
	grpc.ClientStream
}

type nodeHealthcheckClient struct {
	grpc.ClientStream
}

func (x *nodeHealthcheckClient) Send(m *HealthcheckMetadata) error {
	return x.ClientStream.SendMsg(m)
}

func (x *nodeHealthcheckClient) CloseAndRecv() (*Empty, error) {
	if err := x.ClientStream.CloseSend(); err != nil {
		return nil, err
	}
	m := new(Empty)
	if err := x.ClientStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

// NodeServer is the server API for Node service.
// All implementations must embed UnimplementedNodeServer
// for forward compatibility
type NodeServer interface {
	Receive(context.Context, *Data) (*Response, error)
	Delete(context.Context, *Metadata) (*Response, error)
	Healthcheck(Node_HealthcheckServer) error
	mustEmbedUnimplementedNodeServer()
}

// UnimplementedNodeServer must be embedded to have forward compatible implementations.
type UnimplementedNodeServer struct {
}

func (UnimplementedNodeServer) Receive(context.Context, *Data) (*Response, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Receive not implemented")
}
func (UnimplementedNodeServer) Delete(context.Context, *Metadata) (*Response, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Delete not implemented")
}
func (UnimplementedNodeServer) Healthcheck(Node_HealthcheckServer) error {
	return status.Errorf(codes.Unimplemented, "method Healthcheck not implemented")
}
func (UnimplementedNodeServer) mustEmbedUnimplementedNodeServer() {}

// UnsafeNodeServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to NodeServer will
// result in compilation errors.
type UnsafeNodeServer interface {
	mustEmbedUnimplementedNodeServer()
}

func RegisterNodeServer(s grpc.ServiceRegistrar, srv NodeServer) {
	s.RegisterService(&Node_ServiceDesc, srv)
}

func _Node_Receive_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(Data)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(NodeServer).Receive(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/Node/Receive",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(NodeServer).Receive(ctx, req.(*Data))
	}
	return interceptor(ctx, in, info, handler)
}

func _Node_Delete_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(Metadata)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(NodeServer).Delete(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/Node/Delete",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(NodeServer).Delete(ctx, req.(*Metadata))
	}
	return interceptor(ctx, in, info, handler)
}

func _Node_Healthcheck_Handler(srv interface{}, stream grpc.ServerStream) error {
	return srv.(NodeServer).Healthcheck(&nodeHealthcheckServer{stream})
}

type Node_HealthcheckServer interface {
	SendAndClose(*Empty) error
	Recv() (*HealthcheckMetadata, error)
	grpc.ServerStream
}

type nodeHealthcheckServer struct {
	grpc.ServerStream
}

func (x *nodeHealthcheckServer) SendAndClose(m *Empty) error {
	return x.ServerStream.SendMsg(m)
}

func (x *nodeHealthcheckServer) Recv() (*HealthcheckMetadata, error) {
	m := new(HealthcheckMetadata)
	if err := x.ServerStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

// Node_ServiceDesc is the grpc.ServiceDesc for Node service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var Node_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "Node",
	HandlerType: (*NodeServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "Receive",
			Handler:    _Node_Receive_Handler,
		},
		{
			MethodName: "Delete",
			Handler:    _Node_Delete_Handler,
		},
	},
	Streams: []grpc.StreamDesc{
		{
			StreamName:    "Healthcheck",
			Handler:       _Node_Healthcheck_Handler,
			ClientStreams: true,
		},
	},
	Metadata: "request.proto",
}
