// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.2.0
// - protoc             v4.25.3
// source: proto/hello.proto

package proto

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

// HelloServiceClient is the client API for HelloService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type HelloServiceClient interface {
	HelloUnary(ctx context.Context, in *HelloRequest, opts ...grpc.CallOption) (*HelloResponse, error)
	HelloServerStream(ctx context.Context, in *Names, opts ...grpc.CallOption) (HelloService_HelloServerStreamClient, error)
	HelloClientStream(ctx context.Context, opts ...grpc.CallOption) (HelloService_HelloClientStreamClient, error)
	HelloBiDirectionalStream(ctx context.Context, opts ...grpc.CallOption) (HelloService_HelloBiDirectionalStreamClient, error)
}

type helloServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewHelloServiceClient(cc grpc.ClientConnInterface) HelloServiceClient {
	return &helloServiceClient{cc}
}

func (c *helloServiceClient) HelloUnary(ctx context.Context, in *HelloRequest, opts ...grpc.CallOption) (*HelloResponse, error) {
	out := new(HelloResponse)
	err := c.cc.Invoke(ctx, "/hello_service.HelloService/HelloUnary", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *helloServiceClient) HelloServerStream(ctx context.Context, in *Names, opts ...grpc.CallOption) (HelloService_HelloServerStreamClient, error) {
	stream, err := c.cc.NewStream(ctx, &HelloService_ServiceDesc.Streams[0], "/hello_service.HelloService/HelloServerStream", opts...)
	if err != nil {
		return nil, err
	}
	x := &helloServiceHelloServerStreamClient{stream}
	if err := x.ClientStream.SendMsg(in); err != nil {
		return nil, err
	}
	if err := x.ClientStream.CloseSend(); err != nil {
		return nil, err
	}
	return x, nil
}

type HelloService_HelloServerStreamClient interface {
	Recv() (*HelloResponse, error)
	grpc.ClientStream
}

type helloServiceHelloServerStreamClient struct {
	grpc.ClientStream
}

func (x *helloServiceHelloServerStreamClient) Recv() (*HelloResponse, error) {
	m := new(HelloResponse)
	if err := x.ClientStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

func (c *helloServiceClient) HelloClientStream(ctx context.Context, opts ...grpc.CallOption) (HelloService_HelloClientStreamClient, error) {
	stream, err := c.cc.NewStream(ctx, &HelloService_ServiceDesc.Streams[1], "/hello_service.HelloService/HelloClientStream", opts...)
	if err != nil {
		return nil, err
	}
	x := &helloServiceHelloClientStreamClient{stream}
	return x, nil
}

type HelloService_HelloClientStreamClient interface {
	Send(*HelloRequest) error
	CloseAndRecv() (*Messages, error)
	grpc.ClientStream
}

type helloServiceHelloClientStreamClient struct {
	grpc.ClientStream
}

func (x *helloServiceHelloClientStreamClient) Send(m *HelloRequest) error {
	return x.ClientStream.SendMsg(m)
}

func (x *helloServiceHelloClientStreamClient) CloseAndRecv() (*Messages, error) {
	if err := x.ClientStream.CloseSend(); err != nil {
		return nil, err
	}
	m := new(Messages)
	if err := x.ClientStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

func (c *helloServiceClient) HelloBiDirectionalStream(ctx context.Context, opts ...grpc.CallOption) (HelloService_HelloBiDirectionalStreamClient, error) {
	stream, err := c.cc.NewStream(ctx, &HelloService_ServiceDesc.Streams[2], "/hello_service.HelloService/HelloBiDirectionalStream", opts...)
	if err != nil {
		return nil, err
	}
	x := &helloServiceHelloBiDirectionalStreamClient{stream}
	return x, nil
}

type HelloService_HelloBiDirectionalStreamClient interface {
	Send(*HelloRequest) error
	Recv() (*HelloResponse, error)
	grpc.ClientStream
}

type helloServiceHelloBiDirectionalStreamClient struct {
	grpc.ClientStream
}

func (x *helloServiceHelloBiDirectionalStreamClient) Send(m *HelloRequest) error {
	return x.ClientStream.SendMsg(m)
}

func (x *helloServiceHelloBiDirectionalStreamClient) Recv() (*HelloResponse, error) {
	m := new(HelloResponse)
	if err := x.ClientStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

// HelloServiceServer is the server API for HelloService service.
// All implementations must embed UnimplementedHelloServiceServer
// for forward compatibility
type HelloServiceServer interface {
	HelloUnary(context.Context, *HelloRequest) (*HelloResponse, error)
	HelloServerStream(*Names, HelloService_HelloServerStreamServer) error
	HelloClientStream(HelloService_HelloClientStreamServer) error
	HelloBiDirectionalStream(HelloService_HelloBiDirectionalStreamServer) error
	mustEmbedUnimplementedHelloServiceServer()
}

// UnimplementedHelloServiceServer must be embedded to have forward compatible implementations.
type UnimplementedHelloServiceServer struct {
}

func (UnimplementedHelloServiceServer) HelloUnary(context.Context, *HelloRequest) (*HelloResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method HelloUnary not implemented")
}
func (UnimplementedHelloServiceServer) HelloServerStream(*Names, HelloService_HelloServerStreamServer) error {
	return status.Errorf(codes.Unimplemented, "method HelloServerStream not implemented")
}
func (UnimplementedHelloServiceServer) HelloClientStream(HelloService_HelloClientStreamServer) error {
	return status.Errorf(codes.Unimplemented, "method HelloClientStream not implemented")
}
func (UnimplementedHelloServiceServer) HelloBiDirectionalStream(HelloService_HelloBiDirectionalStreamServer) error {
	return status.Errorf(codes.Unimplemented, "method HelloBiDirectionalStream not implemented")
}
func (UnimplementedHelloServiceServer) mustEmbedUnimplementedHelloServiceServer() {}

// UnsafeHelloServiceServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to HelloServiceServer will
// result in compilation errors.
type UnsafeHelloServiceServer interface {
	mustEmbedUnimplementedHelloServiceServer()
}

func RegisterHelloServiceServer(s grpc.ServiceRegistrar, srv HelloServiceServer) {
	s.RegisterService(&HelloService_ServiceDesc, srv)
}

func _HelloService_HelloUnary_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(HelloRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(HelloServiceServer).HelloUnary(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/hello_service.HelloService/HelloUnary",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(HelloServiceServer).HelloUnary(ctx, req.(*HelloRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _HelloService_HelloServerStream_Handler(srv interface{}, stream grpc.ServerStream) error {
	m := new(Names)
	if err := stream.RecvMsg(m); err != nil {
		return err
	}
	return srv.(HelloServiceServer).HelloServerStream(m, &helloServiceHelloServerStreamServer{stream})
}

type HelloService_HelloServerStreamServer interface {
	Send(*HelloResponse) error
	grpc.ServerStream
}

type helloServiceHelloServerStreamServer struct {
	grpc.ServerStream
}

func (x *helloServiceHelloServerStreamServer) Send(m *HelloResponse) error {
	return x.ServerStream.SendMsg(m)
}

func _HelloService_HelloClientStream_Handler(srv interface{}, stream grpc.ServerStream) error {
	return srv.(HelloServiceServer).HelloClientStream(&helloServiceHelloClientStreamServer{stream})
}

type HelloService_HelloClientStreamServer interface {
	SendAndClose(*Messages) error
	Recv() (*HelloRequest, error)
	grpc.ServerStream
}

type helloServiceHelloClientStreamServer struct {
	grpc.ServerStream
}

func (x *helloServiceHelloClientStreamServer) SendAndClose(m *Messages) error {
	return x.ServerStream.SendMsg(m)
}

func (x *helloServiceHelloClientStreamServer) Recv() (*HelloRequest, error) {
	m := new(HelloRequest)
	if err := x.ServerStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

func _HelloService_HelloBiDirectionalStream_Handler(srv interface{}, stream grpc.ServerStream) error {
	return srv.(HelloServiceServer).HelloBiDirectionalStream(&helloServiceHelloBiDirectionalStreamServer{stream})
}

type HelloService_HelloBiDirectionalStreamServer interface {
	Send(*HelloResponse) error
	Recv() (*HelloRequest, error)
	grpc.ServerStream
}

type helloServiceHelloBiDirectionalStreamServer struct {
	grpc.ServerStream
}

func (x *helloServiceHelloBiDirectionalStreamServer) Send(m *HelloResponse) error {
	return x.ServerStream.SendMsg(m)
}

func (x *helloServiceHelloBiDirectionalStreamServer) Recv() (*HelloRequest, error) {
	m := new(HelloRequest)
	if err := x.ServerStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

// HelloService_ServiceDesc is the grpc.ServiceDesc for HelloService service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var HelloService_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "hello_service.HelloService",
	HandlerType: (*HelloServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "HelloUnary",
			Handler:    _HelloService_HelloUnary_Handler,
		},
	},
	Streams: []grpc.StreamDesc{
		{
			StreamName:    "HelloServerStream",
			Handler:       _HelloService_HelloServerStream_Handler,
			ServerStreams: true,
		},
		{
			StreamName:    "HelloClientStream",
			Handler:       _HelloService_HelloClientStream_Handler,
			ClientStreams: true,
		},
		{
			StreamName:    "HelloBiDirectionalStream",
			Handler:       _HelloService_HelloBiDirectionalStream_Handler,
			ServerStreams: true,
			ClientStreams: true,
		},
	},
	Metadata: "proto/hello.proto",
}
