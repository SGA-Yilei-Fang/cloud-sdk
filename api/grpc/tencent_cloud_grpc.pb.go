// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.2.0
// - protoc             (unknown)
// source: api/grpc/tencent_cloud.proto

package api

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

// TencentCloudClient is the client API for TencentCloud service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type TencentCloudClient interface {
	SendMessage(ctx context.Context, in *TencentSendMessageRequest, opts ...grpc.CallOption) (*TencentSendMessageResponse, error)
}

type tencentCloudClient struct {
	cc grpc.ClientConnInterface
}

func NewTencentCloudClient(cc grpc.ClientConnInterface) TencentCloudClient {
	return &tencentCloudClient{cc}
}

func (c *tencentCloudClient) SendMessage(ctx context.Context, in *TencentSendMessageRequest, opts ...grpc.CallOption) (*TencentSendMessageResponse, error) {
	out := new(TencentSendMessageResponse)
	err := c.cc.Invoke(ctx, "/api.TencentCloud/SendMessage", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// TencentCloudServer is the server API for TencentCloud service.
// All implementations must embed UnimplementedTencentCloudServer
// for forward compatibility
type TencentCloudServer interface {
	SendMessage(context.Context, *TencentSendMessageRequest) (*TencentSendMessageResponse, error)
	mustEmbedUnimplementedTencentCloudServer()
}

// UnimplementedTencentCloudServer must be embedded to have forward compatible implementations.
type UnimplementedTencentCloudServer struct {
}

func (UnimplementedTencentCloudServer) SendMessage(context.Context, *TencentSendMessageRequest) (*TencentSendMessageResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method SendMessage not implemented")
}
func (UnimplementedTencentCloudServer) mustEmbedUnimplementedTencentCloudServer() {}

// UnsafeTencentCloudServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to TencentCloudServer will
// result in compilation errors.
type UnsafeTencentCloudServer interface {
	mustEmbedUnimplementedTencentCloudServer()
}

func RegisterTencentCloudServer(s grpc.ServiceRegistrar, srv TencentCloudServer) {
	s.RegisterService(&TencentCloud_ServiceDesc, srv)
}

func _TencentCloud_SendMessage_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(TencentSendMessageRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(TencentCloudServer).SendMessage(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/api.TencentCloud/SendMessage",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(TencentCloudServer).SendMessage(ctx, req.(*TencentSendMessageRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// TencentCloud_ServiceDesc is the grpc.ServiceDesc for TencentCloud service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var TencentCloud_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "api.TencentCloud",
	HandlerType: (*TencentCloudServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "SendMessage",
			Handler:    _TencentCloud_SendMessage_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "api/grpc/tencent_cloud.proto",
}
