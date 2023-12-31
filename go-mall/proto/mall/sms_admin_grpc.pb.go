// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.3.0
// - protoc             v3.21.12
// source: admin/sms_admin.proto

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
	SmsAdminApi_CreateHomeAdvertise_FullMethodName = "/admin.SmsAdminApi/CreateHomeAdvertise"
	SmsAdminApi_UpdateHomeAdvertise_FullMethodName = "/admin.SmsAdminApi/UpdateHomeAdvertise"
	SmsAdminApi_GetHomeAdvertises_FullMethodName   = "/admin.SmsAdminApi/GetHomeAdvertises"
	SmsAdminApi_GetHomeAdvertise_FullMethodName    = "/admin.SmsAdminApi/GetHomeAdvertise"
	SmsAdminApi_DeleteHomeAdvertise_FullMethodName = "/admin.SmsAdminApi/DeleteHomeAdvertise"
)

// SmsAdminApiClient is the client API for SmsAdminApi service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type SmsAdminApiClient interface {
	// 添加首页轮播广告表
	CreateHomeAdvertise(ctx context.Context, in *AddOrUpdateHomeAdvertiseParam, opts ...grpc.CallOption) (*CommonRsp, error)
	// 修改首页轮播广告表
	UpdateHomeAdvertise(ctx context.Context, in *AddOrUpdateHomeAdvertiseParam, opts ...grpc.CallOption) (*CommonRsp, error)
	// 分页查询首页轮播广告表
	GetHomeAdvertises(ctx context.Context, in *GetHomeAdvertisesParam, opts ...grpc.CallOption) (*GetHomeAdvertisesRsp, error)
	// 根据id获取首页轮播广告表
	GetHomeAdvertise(ctx context.Context, in *GetHomeAdvertiseReq, opts ...grpc.CallOption) (*GetHomeAdvertiseRsp, error)
	// 删除首页轮播广告表
	DeleteHomeAdvertise(ctx context.Context, in *DeleteHomeAdvertiseReq, opts ...grpc.CallOption) (*CommonRsp, error)
}

type smsAdminApiClient struct {
	cc grpc.ClientConnInterface
}

func NewSmsAdminApiClient(cc grpc.ClientConnInterface) SmsAdminApiClient {
	return &smsAdminApiClient{cc}
}

func (c *smsAdminApiClient) CreateHomeAdvertise(ctx context.Context, in *AddOrUpdateHomeAdvertiseParam, opts ...grpc.CallOption) (*CommonRsp, error) {
	out := new(CommonRsp)
	err := c.cc.Invoke(ctx, SmsAdminApi_CreateHomeAdvertise_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *smsAdminApiClient) UpdateHomeAdvertise(ctx context.Context, in *AddOrUpdateHomeAdvertiseParam, opts ...grpc.CallOption) (*CommonRsp, error) {
	out := new(CommonRsp)
	err := c.cc.Invoke(ctx, SmsAdminApi_UpdateHomeAdvertise_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *smsAdminApiClient) GetHomeAdvertises(ctx context.Context, in *GetHomeAdvertisesParam, opts ...grpc.CallOption) (*GetHomeAdvertisesRsp, error) {
	out := new(GetHomeAdvertisesRsp)
	err := c.cc.Invoke(ctx, SmsAdminApi_GetHomeAdvertises_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *smsAdminApiClient) GetHomeAdvertise(ctx context.Context, in *GetHomeAdvertiseReq, opts ...grpc.CallOption) (*GetHomeAdvertiseRsp, error) {
	out := new(GetHomeAdvertiseRsp)
	err := c.cc.Invoke(ctx, SmsAdminApi_GetHomeAdvertise_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *smsAdminApiClient) DeleteHomeAdvertise(ctx context.Context, in *DeleteHomeAdvertiseReq, opts ...grpc.CallOption) (*CommonRsp, error) {
	out := new(CommonRsp)
	err := c.cc.Invoke(ctx, SmsAdminApi_DeleteHomeAdvertise_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// SmsAdminApiServer is the server API for SmsAdminApi service.
// All implementations must embed UnimplementedSmsAdminApiServer
// for forward compatibility
type SmsAdminApiServer interface {
	// 添加首页轮播广告表
	CreateHomeAdvertise(context.Context, *AddOrUpdateHomeAdvertiseParam) (*CommonRsp, error)
	// 修改首页轮播广告表
	UpdateHomeAdvertise(context.Context, *AddOrUpdateHomeAdvertiseParam) (*CommonRsp, error)
	// 分页查询首页轮播广告表
	GetHomeAdvertises(context.Context, *GetHomeAdvertisesParam) (*GetHomeAdvertisesRsp, error)
	// 根据id获取首页轮播广告表
	GetHomeAdvertise(context.Context, *GetHomeAdvertiseReq) (*GetHomeAdvertiseRsp, error)
	// 删除首页轮播广告表
	DeleteHomeAdvertise(context.Context, *DeleteHomeAdvertiseReq) (*CommonRsp, error)
	mustEmbedUnimplementedSmsAdminApiServer()
}

// UnimplementedSmsAdminApiServer must be embedded to have forward compatible implementations.
type UnimplementedSmsAdminApiServer struct {
}

func (UnimplementedSmsAdminApiServer) CreateHomeAdvertise(context.Context, *AddOrUpdateHomeAdvertiseParam) (*CommonRsp, error) {
	return nil, status.Errorf(codes.Unimplemented, "method CreateHomeAdvertise not implemented")
}
func (UnimplementedSmsAdminApiServer) UpdateHomeAdvertise(context.Context, *AddOrUpdateHomeAdvertiseParam) (*CommonRsp, error) {
	return nil, status.Errorf(codes.Unimplemented, "method UpdateHomeAdvertise not implemented")
}
func (UnimplementedSmsAdminApiServer) GetHomeAdvertises(context.Context, *GetHomeAdvertisesParam) (*GetHomeAdvertisesRsp, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetHomeAdvertises not implemented")
}
func (UnimplementedSmsAdminApiServer) GetHomeAdvertise(context.Context, *GetHomeAdvertiseReq) (*GetHomeAdvertiseRsp, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetHomeAdvertise not implemented")
}
func (UnimplementedSmsAdminApiServer) DeleteHomeAdvertise(context.Context, *DeleteHomeAdvertiseReq) (*CommonRsp, error) {
	return nil, status.Errorf(codes.Unimplemented, "method DeleteHomeAdvertise not implemented")
}
func (UnimplementedSmsAdminApiServer) mustEmbedUnimplementedSmsAdminApiServer() {}

// UnsafeSmsAdminApiServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to SmsAdminApiServer will
// result in compilation errors.
type UnsafeSmsAdminApiServer interface {
	mustEmbedUnimplementedSmsAdminApiServer()
}

func RegisterSmsAdminApiServer(s grpc.ServiceRegistrar, srv SmsAdminApiServer) {
	s.RegisterService(&SmsAdminApi_ServiceDesc, srv)
}

func _SmsAdminApi_CreateHomeAdvertise_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(AddOrUpdateHomeAdvertiseParam)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(SmsAdminApiServer).CreateHomeAdvertise(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: SmsAdminApi_CreateHomeAdvertise_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(SmsAdminApiServer).CreateHomeAdvertise(ctx, req.(*AddOrUpdateHomeAdvertiseParam))
	}
	return interceptor(ctx, in, info, handler)
}

func _SmsAdminApi_UpdateHomeAdvertise_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(AddOrUpdateHomeAdvertiseParam)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(SmsAdminApiServer).UpdateHomeAdvertise(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: SmsAdminApi_UpdateHomeAdvertise_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(SmsAdminApiServer).UpdateHomeAdvertise(ctx, req.(*AddOrUpdateHomeAdvertiseParam))
	}
	return interceptor(ctx, in, info, handler)
}

func _SmsAdminApi_GetHomeAdvertises_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetHomeAdvertisesParam)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(SmsAdminApiServer).GetHomeAdvertises(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: SmsAdminApi_GetHomeAdvertises_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(SmsAdminApiServer).GetHomeAdvertises(ctx, req.(*GetHomeAdvertisesParam))
	}
	return interceptor(ctx, in, info, handler)
}

func _SmsAdminApi_GetHomeAdvertise_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetHomeAdvertiseReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(SmsAdminApiServer).GetHomeAdvertise(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: SmsAdminApi_GetHomeAdvertise_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(SmsAdminApiServer).GetHomeAdvertise(ctx, req.(*GetHomeAdvertiseReq))
	}
	return interceptor(ctx, in, info, handler)
}

func _SmsAdminApi_DeleteHomeAdvertise_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(DeleteHomeAdvertiseReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(SmsAdminApiServer).DeleteHomeAdvertise(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: SmsAdminApi_DeleteHomeAdvertise_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(SmsAdminApiServer).DeleteHomeAdvertise(ctx, req.(*DeleteHomeAdvertiseReq))
	}
	return interceptor(ctx, in, info, handler)
}

// SmsAdminApi_ServiceDesc is the grpc.ServiceDesc for SmsAdminApi service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var SmsAdminApi_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "admin.SmsAdminApi",
	HandlerType: (*SmsAdminApiServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "CreateHomeAdvertise",
			Handler:    _SmsAdminApi_CreateHomeAdvertise_Handler,
		},
		{
			MethodName: "UpdateHomeAdvertise",
			Handler:    _SmsAdminApi_UpdateHomeAdvertise_Handler,
		},
		{
			MethodName: "GetHomeAdvertises",
			Handler:    _SmsAdminApi_GetHomeAdvertises_Handler,
		},
		{
			MethodName: "GetHomeAdvertise",
			Handler:    _SmsAdminApi_GetHomeAdvertise_Handler,
		},
		{
			MethodName: "DeleteHomeAdvertise",
			Handler:    _SmsAdminApi_DeleteHomeAdvertise_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "admin/sms_admin.proto",
}
