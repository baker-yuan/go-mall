// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.3.0
// - protoc             v3.21.12
// source: portal/member.proto

package mall

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
	PortalMemberApi_MemberRegister_FullMethodName       = "/admin.PortalMemberApi/MemberRegister"
	PortalMemberApi_MemberLogin_FullMethodName          = "/admin.PortalMemberApi/MemberLogin"
	PortalMemberApi_MemberInfo_FullMethodName           = "/admin.PortalMemberApi/MemberInfo"
	PortalMemberApi_MemberGetAuthCode_FullMethodName    = "/admin.PortalMemberApi/MemberGetAuthCode"
	PortalMemberApi_MemberUpdatePassword_FullMethodName = "/admin.PortalMemberApi/MemberUpdatePassword"
	PortalMemberApi_MemberRefreshToken_FullMethodName   = "/admin.PortalMemberApi/MemberRefreshToken"
)

// PortalMemberApiClient is the client API for PortalMemberApi service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type PortalMemberApiClient interface {
	// 会员注册
	MemberRegister(ctx context.Context, in *MemberRegisterReq, opts ...grpc.CallOption) (*EmptyRsp, error)
	// 会员登录
	MemberLogin(ctx context.Context, in *MemberLoginReq, opts ...grpc.CallOption) (*MemberLoginRsp, error)
	// 获取会员信息
	MemberInfo(ctx context.Context, in *MemberInfoReq, opts ...grpc.CallOption) (*MemberInfoRsp, error)
	// 获取验证码
	MemberGetAuthCode(ctx context.Context, in *MemberGetAuthCodeReq, opts ...grpc.CallOption) (*MemberGetAuthCodeRsp, error)
	// 修改密码
	MemberUpdatePassword(ctx context.Context, in *MemberUpdatePasswordReq, opts ...grpc.CallOption) (*MemberUpdatePasswordRsp, error)
	// 刷新token
	MemberRefreshToken(ctx context.Context, in *MemberRefreshTokenReq, opts ...grpc.CallOption) (*MemberRefreshTokenRsp, error)
}

type portalMemberApiClient struct {
	cc grpc.ClientConnInterface
}

func NewPortalMemberApiClient(cc grpc.ClientConnInterface) PortalMemberApiClient {
	return &portalMemberApiClient{cc}
}

func (c *portalMemberApiClient) MemberRegister(ctx context.Context, in *MemberRegisterReq, opts ...grpc.CallOption) (*EmptyRsp, error) {
	out := new(EmptyRsp)
	err := c.cc.Invoke(ctx, PortalMemberApi_MemberRegister_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *portalMemberApiClient) MemberLogin(ctx context.Context, in *MemberLoginReq, opts ...grpc.CallOption) (*MemberLoginRsp, error) {
	out := new(MemberLoginRsp)
	err := c.cc.Invoke(ctx, PortalMemberApi_MemberLogin_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *portalMemberApiClient) MemberInfo(ctx context.Context, in *MemberInfoReq, opts ...grpc.CallOption) (*MemberInfoRsp, error) {
	out := new(MemberInfoRsp)
	err := c.cc.Invoke(ctx, PortalMemberApi_MemberInfo_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *portalMemberApiClient) MemberGetAuthCode(ctx context.Context, in *MemberGetAuthCodeReq, opts ...grpc.CallOption) (*MemberGetAuthCodeRsp, error) {
	out := new(MemberGetAuthCodeRsp)
	err := c.cc.Invoke(ctx, PortalMemberApi_MemberGetAuthCode_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *portalMemberApiClient) MemberUpdatePassword(ctx context.Context, in *MemberUpdatePasswordReq, opts ...grpc.CallOption) (*MemberUpdatePasswordRsp, error) {
	out := new(MemberUpdatePasswordRsp)
	err := c.cc.Invoke(ctx, PortalMemberApi_MemberUpdatePassword_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *portalMemberApiClient) MemberRefreshToken(ctx context.Context, in *MemberRefreshTokenReq, opts ...grpc.CallOption) (*MemberRefreshTokenRsp, error) {
	out := new(MemberRefreshTokenRsp)
	err := c.cc.Invoke(ctx, PortalMemberApi_MemberRefreshToken_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// PortalMemberApiServer is the server API for PortalMemberApi service.
// All implementations must embed UnimplementedPortalMemberApiServer
// for forward compatibility
type PortalMemberApiServer interface {
	// 会员注册
	MemberRegister(context.Context, *MemberRegisterReq) (*EmptyRsp, error)
	// 会员登录
	MemberLogin(context.Context, *MemberLoginReq) (*MemberLoginRsp, error)
	// 获取会员信息
	MemberInfo(context.Context, *MemberInfoReq) (*MemberInfoRsp, error)
	// 获取验证码
	MemberGetAuthCode(context.Context, *MemberGetAuthCodeReq) (*MemberGetAuthCodeRsp, error)
	// 修改密码
	MemberUpdatePassword(context.Context, *MemberUpdatePasswordReq) (*MemberUpdatePasswordRsp, error)
	// 刷新token
	MemberRefreshToken(context.Context, *MemberRefreshTokenReq) (*MemberRefreshTokenRsp, error)
	mustEmbedUnimplementedPortalMemberApiServer()
}

// UnimplementedPortalMemberApiServer must be embedded to have forward compatible implementations.
type UnimplementedPortalMemberApiServer struct {
}

func (UnimplementedPortalMemberApiServer) MemberRegister(context.Context, *MemberRegisterReq) (*EmptyRsp, error) {
	return nil, status.Errorf(codes.Unimplemented, "method MemberRegister not implemented")
}
func (UnimplementedPortalMemberApiServer) MemberLogin(context.Context, *MemberLoginReq) (*MemberLoginRsp, error) {
	return nil, status.Errorf(codes.Unimplemented, "method MemberLogin not implemented")
}
func (UnimplementedPortalMemberApiServer) MemberInfo(context.Context, *MemberInfoReq) (*MemberInfoRsp, error) {
	return nil, status.Errorf(codes.Unimplemented, "method MemberInfo not implemented")
}
func (UnimplementedPortalMemberApiServer) MemberGetAuthCode(context.Context, *MemberGetAuthCodeReq) (*MemberGetAuthCodeRsp, error) {
	return nil, status.Errorf(codes.Unimplemented, "method MemberGetAuthCode not implemented")
}
func (UnimplementedPortalMemberApiServer) MemberUpdatePassword(context.Context, *MemberUpdatePasswordReq) (*MemberUpdatePasswordRsp, error) {
	return nil, status.Errorf(codes.Unimplemented, "method MemberUpdatePassword not implemented")
}
func (UnimplementedPortalMemberApiServer) MemberRefreshToken(context.Context, *MemberRefreshTokenReq) (*MemberRefreshTokenRsp, error) {
	return nil, status.Errorf(codes.Unimplemented, "method MemberRefreshToken not implemented")
}
func (UnimplementedPortalMemberApiServer) mustEmbedUnimplementedPortalMemberApiServer() {}

// UnsafePortalMemberApiServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to PortalMemberApiServer will
// result in compilation errors.
type UnsafePortalMemberApiServer interface {
	mustEmbedUnimplementedPortalMemberApiServer()
}

func RegisterPortalMemberApiServer(s grpc.ServiceRegistrar, srv PortalMemberApiServer) {
	s.RegisterService(&PortalMemberApi_ServiceDesc, srv)
}

func _PortalMemberApi_MemberRegister_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(MemberRegisterReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(PortalMemberApiServer).MemberRegister(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: PortalMemberApi_MemberRegister_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(PortalMemberApiServer).MemberRegister(ctx, req.(*MemberRegisterReq))
	}
	return interceptor(ctx, in, info, handler)
}

func _PortalMemberApi_MemberLogin_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(MemberLoginReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(PortalMemberApiServer).MemberLogin(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: PortalMemberApi_MemberLogin_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(PortalMemberApiServer).MemberLogin(ctx, req.(*MemberLoginReq))
	}
	return interceptor(ctx, in, info, handler)
}

func _PortalMemberApi_MemberInfo_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(MemberInfoReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(PortalMemberApiServer).MemberInfo(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: PortalMemberApi_MemberInfo_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(PortalMemberApiServer).MemberInfo(ctx, req.(*MemberInfoReq))
	}
	return interceptor(ctx, in, info, handler)
}

func _PortalMemberApi_MemberGetAuthCode_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(MemberGetAuthCodeReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(PortalMemberApiServer).MemberGetAuthCode(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: PortalMemberApi_MemberGetAuthCode_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(PortalMemberApiServer).MemberGetAuthCode(ctx, req.(*MemberGetAuthCodeReq))
	}
	return interceptor(ctx, in, info, handler)
}

func _PortalMemberApi_MemberUpdatePassword_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(MemberUpdatePasswordReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(PortalMemberApiServer).MemberUpdatePassword(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: PortalMemberApi_MemberUpdatePassword_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(PortalMemberApiServer).MemberUpdatePassword(ctx, req.(*MemberUpdatePasswordReq))
	}
	return interceptor(ctx, in, info, handler)
}

func _PortalMemberApi_MemberRefreshToken_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(MemberRefreshTokenReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(PortalMemberApiServer).MemberRefreshToken(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: PortalMemberApi_MemberRefreshToken_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(PortalMemberApiServer).MemberRefreshToken(ctx, req.(*MemberRefreshTokenReq))
	}
	return interceptor(ctx, in, info, handler)
}

// PortalMemberApi_ServiceDesc is the grpc.ServiceDesc for PortalMemberApi service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var PortalMemberApi_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "admin.PortalMemberApi",
	HandlerType: (*PortalMemberApiServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "MemberRegister",
			Handler:    _PortalMemberApi_MemberRegister_Handler,
		},
		{
			MethodName: "MemberLogin",
			Handler:    _PortalMemberApi_MemberLogin_Handler,
		},
		{
			MethodName: "MemberInfo",
			Handler:    _PortalMemberApi_MemberInfo_Handler,
		},
		{
			MethodName: "MemberGetAuthCode",
			Handler:    _PortalMemberApi_MemberGetAuthCode_Handler,
		},
		{
			MethodName: "MemberUpdatePassword",
			Handler:    _PortalMemberApi_MemberUpdatePassword_Handler,
		},
		{
			MethodName: "MemberRefreshToken",
			Handler:    _PortalMemberApi_MemberRefreshToken_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "portal/member.proto",
}