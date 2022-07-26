// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.2.0
// - protoc             v3.21.2
// source: oauth.proto

package oauthpb

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

// OAuthServiceClient is the client API for OAuthService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type OAuthServiceClient interface {
	//rpc Token (OAuthTokenRequest) returns (OAuthTokenResponse);
	Verify(ctx context.Context, in *OAuthTokenVerifyRequest, opts ...grpc.CallOption) (*OAuthTokenVerifyResponse, error)
	RemoveToken(ctx context.Context, in *OAuthRemoveTokenRequest, opts ...grpc.CallOption) (*OAuthRemoveTokenResponse, error)
}

type oAuthServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewOAuthServiceClient(cc grpc.ClientConnInterface) OAuthServiceClient {
	return &oAuthServiceClient{cc}
}

func (c *oAuthServiceClient) Verify(ctx context.Context, in *OAuthTokenVerifyRequest, opts ...grpc.CallOption) (*OAuthTokenVerifyResponse, error) {
	out := new(OAuthTokenVerifyResponse)
	err := c.cc.Invoke(ctx, "/oauth.OAuthService/Verify", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *oAuthServiceClient) RemoveToken(ctx context.Context, in *OAuthRemoveTokenRequest, opts ...grpc.CallOption) (*OAuthRemoveTokenResponse, error) {
	out := new(OAuthRemoveTokenResponse)
	err := c.cc.Invoke(ctx, "/oauth.OAuthService/RemoveToken", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// OAuthServiceServer is the server API for OAuthService service.
// All implementations must embed UnimplementedOAuthServiceServer
// for forward compatibility
type OAuthServiceServer interface {
	//rpc Token (OAuthTokenRequest) returns (OAuthTokenResponse);
	Verify(context.Context, *OAuthTokenVerifyRequest) (*OAuthTokenVerifyResponse, error)
	RemoveToken(context.Context, *OAuthRemoveTokenRequest) (*OAuthRemoveTokenResponse, error)
	mustEmbedUnimplementedOAuthServiceServer()
}

// UnimplementedOAuthServiceServer must be embedded to have forward compatible implementations.
type UnimplementedOAuthServiceServer struct {
}

func (UnimplementedOAuthServiceServer) Verify(context.Context, *OAuthTokenVerifyRequest) (*OAuthTokenVerifyResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Verify not implemented")
}
func (UnimplementedOAuthServiceServer) RemoveToken(context.Context, *OAuthRemoveTokenRequest) (*OAuthRemoveTokenResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method RemoveToken not implemented")
}
func (UnimplementedOAuthServiceServer) mustEmbedUnimplementedOAuthServiceServer() {}

// UnsafeOAuthServiceServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to OAuthServiceServer will
// result in compilation errors.
type UnsafeOAuthServiceServer interface {
	mustEmbedUnimplementedOAuthServiceServer()
}

func RegisterOAuthServiceServer(s grpc.ServiceRegistrar, srv OAuthServiceServer) {
	s.RegisterService(&OAuthService_ServiceDesc, srv)
}

func _OAuthService_Verify_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(OAuthTokenVerifyRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(OAuthServiceServer).Verify(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/oauth.OAuthService/Verify",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(OAuthServiceServer).Verify(ctx, req.(*OAuthTokenVerifyRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _OAuthService_RemoveToken_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(OAuthRemoveTokenRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(OAuthServiceServer).RemoveToken(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/oauth.OAuthService/RemoveToken",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(OAuthServiceServer).RemoveToken(ctx, req.(*OAuthRemoveTokenRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// OAuthService_ServiceDesc is the grpc.ServiceDesc for OAuthService service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var OAuthService_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "oauth.OAuthService",
	HandlerType: (*OAuthServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "Verify",
			Handler:    _OAuthService_Verify_Handler,
		},
		{
			MethodName: "RemoveToken",
			Handler:    _OAuthService_RemoveToken_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "oauth.proto",
}

// UserServiceClient is the client API for UserService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type UserServiceClient interface {
	GetUserByUsername(ctx context.Context, in *GetUserRequest, opts ...grpc.CallOption) (*User, error)
}

type userServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewUserServiceClient(cc grpc.ClientConnInterface) UserServiceClient {
	return &userServiceClient{cc}
}

func (c *userServiceClient) GetUserByUsername(ctx context.Context, in *GetUserRequest, opts ...grpc.CallOption) (*User, error) {
	out := new(User)
	err := c.cc.Invoke(ctx, "/oauth.UserService/GetUserByUsername", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// UserServiceServer is the server API for UserService service.
// All implementations must embed UnimplementedUserServiceServer
// for forward compatibility
type UserServiceServer interface {
	GetUserByUsername(context.Context, *GetUserRequest) (*User, error)
	mustEmbedUnimplementedUserServiceServer()
}

// UnimplementedUserServiceServer must be embedded to have forward compatible implementations.
type UnimplementedUserServiceServer struct {
}

func (UnimplementedUserServiceServer) GetUserByUsername(context.Context, *GetUserRequest) (*User, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetUserByUsername not implemented")
}
func (UnimplementedUserServiceServer) mustEmbedUnimplementedUserServiceServer() {}

// UnsafeUserServiceServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to UserServiceServer will
// result in compilation errors.
type UnsafeUserServiceServer interface {
	mustEmbedUnimplementedUserServiceServer()
}

func RegisterUserServiceServer(s grpc.ServiceRegistrar, srv UserServiceServer) {
	s.RegisterService(&UserService_ServiceDesc, srv)
}

func _UserService_GetUserByUsername_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetUserRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(UserServiceServer).GetUserByUsername(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/oauth.UserService/GetUserByUsername",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(UserServiceServer).GetUserByUsername(ctx, req.(*GetUserRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// UserService_ServiceDesc is the grpc.ServiceDesc for UserService service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var UserService_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "oauth.UserService",
	HandlerType: (*UserServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "GetUserByUsername",
			Handler:    _UserService_GetUserByUsername_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "oauth.proto",
}
