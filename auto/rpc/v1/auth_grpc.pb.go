// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.3.0
// - protoc             (unknown)
// source: v1/auth.proto

package v1

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

const (
	Authenticate_PreLogin_FullMethodName = "/auth.Authenticate/preLogin"
	Authenticate_Login_FullMethodName    = "/auth.Authenticate/login"
	Authenticate_Logout_FullMethodName   = "/auth.Authenticate/logout"
)

// AuthenticateClient is the client API for Authenticate service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type AuthenticateClient interface {
	PreLogin(ctx context.Context, in *User, opts ...grpc.CallOption) (*ActionReply, error)
	Login(ctx context.Context, in *User, opts ...grpc.CallOption) (*LoginReply, error)
	Logout(ctx context.Context, in *User, opts ...grpc.CallOption) (*ActionReply, error)
}

type authenticateClient struct {
	cc grpc.ClientConnInterface
}

func NewAuthenticateClient(cc grpc.ClientConnInterface) AuthenticateClient {
	return &authenticateClient{cc}
}

func (c *authenticateClient) PreLogin(ctx context.Context, in *User, opts ...grpc.CallOption) (*ActionReply, error) {
	out := new(ActionReply)
	err := c.cc.Invoke(ctx, Authenticate_PreLogin_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *authenticateClient) Login(ctx context.Context, in *User, opts ...grpc.CallOption) (*LoginReply, error) {
	out := new(LoginReply)
	err := c.cc.Invoke(ctx, Authenticate_Login_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *authenticateClient) Logout(ctx context.Context, in *User, opts ...grpc.CallOption) (*ActionReply, error) {
	out := new(ActionReply)
	err := c.cc.Invoke(ctx, Authenticate_Logout_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// AuthenticateServer is the server API for Authenticate service.
// All implementations must embed UnimplementedAuthenticateServer
// for forward compatibility
type AuthenticateServer interface {
	PreLogin(context.Context, *User) (*ActionReply, error)
	Login(context.Context, *User) (*LoginReply, error)
	Logout(context.Context, *User) (*ActionReply, error)
	mustEmbedUnimplementedAuthenticateServer()
}

// UnimplementedAuthenticateServer must be embedded to have forward compatible implementations.
type UnimplementedAuthenticateServer struct {
}

func (UnimplementedAuthenticateServer) PreLogin(context.Context, *User) (*ActionReply, error) {
	return nil, status.Errorf(codes.Unimplemented, "method PreLogin not implemented")
}
func (UnimplementedAuthenticateServer) Login(context.Context, *User) (*LoginReply, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Login not implemented")
}
func (UnimplementedAuthenticateServer) Logout(context.Context, *User) (*ActionReply, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Logout not implemented")
}
func (UnimplementedAuthenticateServer) mustEmbedUnimplementedAuthenticateServer() {}

// UnsafeAuthenticateServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to AuthenticateServer will
// result in compilation errors.
type UnsafeAuthenticateServer interface {
	mustEmbedUnimplementedAuthenticateServer()
}

func RegisterAuthenticateServer(s grpc.ServiceRegistrar, srv AuthenticateServer) {
	s.RegisterService(&Authenticate_ServiceDesc, srv)
}

func _Authenticate_PreLogin_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(User)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(AuthenticateServer).PreLogin(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Authenticate_PreLogin_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(AuthenticateServer).PreLogin(ctx, req.(*User))
	}
	return interceptor(ctx, in, info, handler)
}

func _Authenticate_Login_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(User)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(AuthenticateServer).Login(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Authenticate_Login_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(AuthenticateServer).Login(ctx, req.(*User))
	}
	return interceptor(ctx, in, info, handler)
}

func _Authenticate_Logout_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(User)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(AuthenticateServer).Logout(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Authenticate_Logout_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(AuthenticateServer).Logout(ctx, req.(*User))
	}
	return interceptor(ctx, in, info, handler)
}

// Authenticate_ServiceDesc is the grpc.ServiceDesc for Authenticate service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var Authenticate_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "auth.Authenticate",
	HandlerType: (*AuthenticateServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "preLogin",
			Handler:    _Authenticate_PreLogin_Handler,
		},
		{
			MethodName: "login",
			Handler:    _Authenticate_Login_Handler,
		},
		{
			MethodName: "logout",
			Handler:    _Authenticate_Logout_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "v1/auth.proto",
}
