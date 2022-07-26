// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.2.0
// - protoc             v3.21.2
// source: follow.proto

package followpb

import (
	context "context"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
)

// This is a compile-times assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
// Requires gRPC-Go v1.32.0 or later.
const _ = grpc.SupportPackageIsVersion7

// FollowServiceClient is the client API for FollowService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type FollowServiceClient interface {
	IsFollow(ctx context.Context, in *IsFollowRequest, opts ...grpc.CallOption) (*IsFollowResponse, error)
	Follow(ctx context.Context, in *FollowRequest, opts ...grpc.CallOption) (*FollowResponse, error)
	Following(ctx context.Context, in *FollowingRequest, opts ...grpc.CallOption) (*FollowingResponse, error)
	Followers(ctx context.Context, in *FollowersRequest, opts ...grpc.CallOption) (*FollowersResponse, error)
}

type followServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewFollowServiceClient(cc grpc.ClientConnInterface) FollowServiceClient {
	return &followServiceClient{cc}
}

func (c *followServiceClient) IsFollow(ctx context.Context, in *IsFollowRequest, opts ...grpc.CallOption) (*IsFollowResponse, error) {
	out := new(IsFollowResponse)
	err := c.cc.Invoke(ctx, "/follow.FollowService/IsFollow", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *followServiceClient) Follow(ctx context.Context, in *FollowRequest, opts ...grpc.CallOption) (*FollowResponse, error) {
	out := new(FollowResponse)
	err := c.cc.Invoke(ctx, "/follow.FollowService/Follow", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *followServiceClient) Following(ctx context.Context, in *FollowingRequest, opts ...grpc.CallOption) (*FollowingResponse, error) {
	out := new(FollowingResponse)
	err := c.cc.Invoke(ctx, "/follow.FollowService/Following", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *followServiceClient) Followers(ctx context.Context, in *FollowersRequest, opts ...grpc.CallOption) (*FollowersResponse, error) {
	out := new(FollowersResponse)
	err := c.cc.Invoke(ctx, "/follow.FollowService/Followers", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// FollowServiceServer is the server API for FollowService service.
// All implementations must embed UnimplementedFollowServiceServer
// for forward compatibility
type FollowServiceServer interface {
	IsFollow(context.Context, *IsFollowRequest) (*IsFollowResponse, error)
	Follow(context.Context, *FollowRequest) (*FollowResponse, error)
	Following(context.Context, *FollowingRequest) (*FollowingResponse, error)
	Followers(context.Context, *FollowersRequest) (*FollowersResponse, error)
	mustEmbedUnimplementedFollowServiceServer()
}

// UnimplementedFollowServiceServer must be embedded to have forward compatible implementations.
type UnimplementedFollowServiceServer struct {
}

func (UnimplementedFollowServiceServer) IsFollow(context.Context, *IsFollowRequest) (*IsFollowResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method IsFollow not implemented")
}
func (UnimplementedFollowServiceServer) Follow(context.Context, *FollowRequest) (*FollowResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Follow not implemented")
}
func (UnimplementedFollowServiceServer) Following(context.Context, *FollowingRequest) (*FollowingResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Following not implemented")
}
func (UnimplementedFollowServiceServer) Followers(context.Context, *FollowersRequest) (*FollowersResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Followers not implemented")
}
func (UnimplementedFollowServiceServer) mustEmbedUnimplementedFollowServiceServer() {}

// UnsafeFollowServiceServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to FollowServiceServer will
// result in compilation errors.
type UnsafeFollowServiceServer interface {
	mustEmbedUnimplementedFollowServiceServer()
}

func RegisterFollowServiceServer(s grpc.ServiceRegistrar, srv FollowServiceServer) {
	s.RegisterService(&FollowService_ServiceDesc, srv)
}

func _FollowService_IsFollow_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(IsFollowRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(FollowServiceServer).IsFollow(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/follow.FollowService/IsFollow",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(FollowServiceServer).IsFollow(ctx, req.(*IsFollowRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _FollowService_Follow_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(FollowRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(FollowServiceServer).Follow(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/follow.FollowService/Follow",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(FollowServiceServer).Follow(ctx, req.(*FollowRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _FollowService_Following_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(FollowingRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(FollowServiceServer).Following(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/follow.FollowService/Following",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(FollowServiceServer).Following(ctx, req.(*FollowingRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _FollowService_Followers_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(FollowersRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(FollowServiceServer).Followers(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/follow.FollowService/Followers",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(FollowServiceServer).Followers(ctx, req.(*FollowersRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// FollowService_ServiceDesc is the grpc.ServiceDesc for FollowService service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var FollowService_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "follow.FollowService",
	HandlerType: (*FollowServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "IsFollow",
			Handler:    _FollowService_IsFollow_Handler,
		},
		{
			MethodName: "Follow",
			Handler:    _FollowService_Follow_Handler,
		},
		{
			MethodName: "Following",
			Handler:    _FollowService_Following_Handler,
		},
		{
			MethodName: "Followers",
			Handler:    _FollowService_Followers_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "follow.proto",
}
