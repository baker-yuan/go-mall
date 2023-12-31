// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.3.0
// - protoc             v3.21.12
// source: admin/oms_admin.proto

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
	OmsAdminApi_CreateOrderReturnReason_FullMethodName = "/admin.OmsAdminApi/CreateOrderReturnReason"
	OmsAdminApi_UpdateOrderReturnReason_FullMethodName = "/admin.OmsAdminApi/UpdateOrderReturnReason"
	OmsAdminApi_GetOrderReturnReasons_FullMethodName   = "/admin.OmsAdminApi/GetOrderReturnReasons"
	OmsAdminApi_GetOrderReturnReason_FullMethodName    = "/admin.OmsAdminApi/GetOrderReturnReason"
	OmsAdminApi_DeleteOrderReturnReason_FullMethodName = "/admin.OmsAdminApi/DeleteOrderReturnReason"
	OmsAdminApi_GetOrders_FullMethodName               = "/admin.OmsAdminApi/GetOrders"
	OmsAdminApi_GetOrder_FullMethodName                = "/admin.OmsAdminApi/GetOrder"
	OmsAdminApi_DeleteOrder_FullMethodName             = "/admin.OmsAdminApi/DeleteOrder"
	OmsAdminApi_GetOrderReturnApplies_FullMethodName   = "/admin.OmsAdminApi/GetOrderReturnApplies"
	OmsAdminApi_GetOrderReturnApply_FullMethodName     = "/admin.OmsAdminApi/GetOrderReturnApply"
	OmsAdminApi_CreateCompanyAddress_FullMethodName    = "/admin.OmsAdminApi/CreateCompanyAddress"
	OmsAdminApi_UpdateCompanyAddress_FullMethodName    = "/admin.OmsAdminApi/UpdateCompanyAddress"
	OmsAdminApi_GetCompanyAddresses_FullMethodName     = "/admin.OmsAdminApi/GetCompanyAddresses"
	OmsAdminApi_GetCompanyAddress_FullMethodName       = "/admin.OmsAdminApi/GetCompanyAddress"
	OmsAdminApi_DeleteCompanyAddress_FullMethodName    = "/admin.OmsAdminApi/DeleteCompanyAddress"
)

// OmsAdminApiClient is the client API for OmsAdminApi service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type OmsAdminApiClient interface {
	// START ======================================= 退货原因 ======================================= START
	// 添加退货原因
	CreateOrderReturnReason(ctx context.Context, in *AddOrUpdateOrderReturnReasonParam, opts ...grpc.CallOption) (*CommonRsp, error)
	// 修改退货原因
	UpdateOrderReturnReason(ctx context.Context, in *AddOrUpdateOrderReturnReasonParam, opts ...grpc.CallOption) (*CommonRsp, error)
	// 分页查询退货原因
	GetOrderReturnReasons(ctx context.Context, in *GetOrderReturnReasonsParam, opts ...grpc.CallOption) (*GetOrderReturnReasonsRsp, error)
	// 根据id获取退货原因
	GetOrderReturnReason(ctx context.Context, in *GetOrderReturnReasonReq, opts ...grpc.CallOption) (*GetOrderReturnReasonRsp, error)
	// 删除退货原因
	DeleteOrderReturnReason(ctx context.Context, in *DeleteOrderReturnReasonReq, opts ...grpc.CallOption) (*CommonRsp, error)
	// START ======================================= 订单 ======================================= START
	// 分页查询订单
	GetOrders(ctx context.Context, in *GetOrdersParam, opts ...grpc.CallOption) (*GetOrdersRsp, error)
	// 根据id获取订单
	GetOrder(ctx context.Context, in *GetOrderReq, opts ...grpc.CallOption) (*GetOrderRsp, error)
	// 删除订单
	DeleteOrder(ctx context.Context, in *DeleteOrderReq, opts ...grpc.CallOption) (*CommonRsp, error)
	// START ======================================= 订单退货 ======================================= START
	// 分页查询订单退货申请
	GetOrderReturnApplies(ctx context.Context, in *GetOrderReturnAppliesParam, opts ...grpc.CallOption) (*GetOrderReturnAppliesRsp, error)
	// 根据id获取订单退货申请
	GetOrderReturnApply(ctx context.Context, in *GetOrderReturnApplyReq, opts ...grpc.CallOption) (*GetOrderReturnApplyRsp, error)
	// START ======================================= 公司收发货地址 ======================================= START
	// 添加公司收发货地址
	CreateCompanyAddress(ctx context.Context, in *AddOrUpdateCompanyAddressParam, opts ...grpc.CallOption) (*CommonRsp, error)
	// 修改公司收发货地址
	UpdateCompanyAddress(ctx context.Context, in *AddOrUpdateCompanyAddressParam, opts ...grpc.CallOption) (*CommonRsp, error)
	// 分页查询公司收发货地址
	GetCompanyAddresses(ctx context.Context, in *GetCompanyAddressesParam, opts ...grpc.CallOption) (*GetCompanyAddressesRsp, error)
	// 根据id获取公司收发货地址
	GetCompanyAddress(ctx context.Context, in *GetCompanyAddressReq, opts ...grpc.CallOption) (*GetCompanyAddressRsp, error)
	// 删除公司收发货地址
	DeleteCompanyAddress(ctx context.Context, in *DeleteCompanyAddressReq, opts ...grpc.CallOption) (*CommonRsp, error)
}

type omsAdminApiClient struct {
	cc grpc.ClientConnInterface
}

func NewOmsAdminApiClient(cc grpc.ClientConnInterface) OmsAdminApiClient {
	return &omsAdminApiClient{cc}
}

func (c *omsAdminApiClient) CreateOrderReturnReason(ctx context.Context, in *AddOrUpdateOrderReturnReasonParam, opts ...grpc.CallOption) (*CommonRsp, error) {
	out := new(CommonRsp)
	err := c.cc.Invoke(ctx, OmsAdminApi_CreateOrderReturnReason_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *omsAdminApiClient) UpdateOrderReturnReason(ctx context.Context, in *AddOrUpdateOrderReturnReasonParam, opts ...grpc.CallOption) (*CommonRsp, error) {
	out := new(CommonRsp)
	err := c.cc.Invoke(ctx, OmsAdminApi_UpdateOrderReturnReason_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *omsAdminApiClient) GetOrderReturnReasons(ctx context.Context, in *GetOrderReturnReasonsParam, opts ...grpc.CallOption) (*GetOrderReturnReasonsRsp, error) {
	out := new(GetOrderReturnReasonsRsp)
	err := c.cc.Invoke(ctx, OmsAdminApi_GetOrderReturnReasons_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *omsAdminApiClient) GetOrderReturnReason(ctx context.Context, in *GetOrderReturnReasonReq, opts ...grpc.CallOption) (*GetOrderReturnReasonRsp, error) {
	out := new(GetOrderReturnReasonRsp)
	err := c.cc.Invoke(ctx, OmsAdminApi_GetOrderReturnReason_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *omsAdminApiClient) DeleteOrderReturnReason(ctx context.Context, in *DeleteOrderReturnReasonReq, opts ...grpc.CallOption) (*CommonRsp, error) {
	out := new(CommonRsp)
	err := c.cc.Invoke(ctx, OmsAdminApi_DeleteOrderReturnReason_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *omsAdminApiClient) GetOrders(ctx context.Context, in *GetOrdersParam, opts ...grpc.CallOption) (*GetOrdersRsp, error) {
	out := new(GetOrdersRsp)
	err := c.cc.Invoke(ctx, OmsAdminApi_GetOrders_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *omsAdminApiClient) GetOrder(ctx context.Context, in *GetOrderReq, opts ...grpc.CallOption) (*GetOrderRsp, error) {
	out := new(GetOrderRsp)
	err := c.cc.Invoke(ctx, OmsAdminApi_GetOrder_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *omsAdminApiClient) DeleteOrder(ctx context.Context, in *DeleteOrderReq, opts ...grpc.CallOption) (*CommonRsp, error) {
	out := new(CommonRsp)
	err := c.cc.Invoke(ctx, OmsAdminApi_DeleteOrder_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *omsAdminApiClient) GetOrderReturnApplies(ctx context.Context, in *GetOrderReturnAppliesParam, opts ...grpc.CallOption) (*GetOrderReturnAppliesRsp, error) {
	out := new(GetOrderReturnAppliesRsp)
	err := c.cc.Invoke(ctx, OmsAdminApi_GetOrderReturnApplies_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *omsAdminApiClient) GetOrderReturnApply(ctx context.Context, in *GetOrderReturnApplyReq, opts ...grpc.CallOption) (*GetOrderReturnApplyRsp, error) {
	out := new(GetOrderReturnApplyRsp)
	err := c.cc.Invoke(ctx, OmsAdminApi_GetOrderReturnApply_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *omsAdminApiClient) CreateCompanyAddress(ctx context.Context, in *AddOrUpdateCompanyAddressParam, opts ...grpc.CallOption) (*CommonRsp, error) {
	out := new(CommonRsp)
	err := c.cc.Invoke(ctx, OmsAdminApi_CreateCompanyAddress_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *omsAdminApiClient) UpdateCompanyAddress(ctx context.Context, in *AddOrUpdateCompanyAddressParam, opts ...grpc.CallOption) (*CommonRsp, error) {
	out := new(CommonRsp)
	err := c.cc.Invoke(ctx, OmsAdminApi_UpdateCompanyAddress_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *omsAdminApiClient) GetCompanyAddresses(ctx context.Context, in *GetCompanyAddressesParam, opts ...grpc.CallOption) (*GetCompanyAddressesRsp, error) {
	out := new(GetCompanyAddressesRsp)
	err := c.cc.Invoke(ctx, OmsAdminApi_GetCompanyAddresses_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *omsAdminApiClient) GetCompanyAddress(ctx context.Context, in *GetCompanyAddressReq, opts ...grpc.CallOption) (*GetCompanyAddressRsp, error) {
	out := new(GetCompanyAddressRsp)
	err := c.cc.Invoke(ctx, OmsAdminApi_GetCompanyAddress_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *omsAdminApiClient) DeleteCompanyAddress(ctx context.Context, in *DeleteCompanyAddressReq, opts ...grpc.CallOption) (*CommonRsp, error) {
	out := new(CommonRsp)
	err := c.cc.Invoke(ctx, OmsAdminApi_DeleteCompanyAddress_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// OmsAdminApiServer is the server API for OmsAdminApi service.
// All implementations must embed UnimplementedOmsAdminApiServer
// for forward compatibility
type OmsAdminApiServer interface {
	// START ======================================= 退货原因 ======================================= START
	// 添加退货原因
	CreateOrderReturnReason(context.Context, *AddOrUpdateOrderReturnReasonParam) (*CommonRsp, error)
	// 修改退货原因
	UpdateOrderReturnReason(context.Context, *AddOrUpdateOrderReturnReasonParam) (*CommonRsp, error)
	// 分页查询退货原因
	GetOrderReturnReasons(context.Context, *GetOrderReturnReasonsParam) (*GetOrderReturnReasonsRsp, error)
	// 根据id获取退货原因
	GetOrderReturnReason(context.Context, *GetOrderReturnReasonReq) (*GetOrderReturnReasonRsp, error)
	// 删除退货原因
	DeleteOrderReturnReason(context.Context, *DeleteOrderReturnReasonReq) (*CommonRsp, error)
	// START ======================================= 订单 ======================================= START
	// 分页查询订单
	GetOrders(context.Context, *GetOrdersParam) (*GetOrdersRsp, error)
	// 根据id获取订单
	GetOrder(context.Context, *GetOrderReq) (*GetOrderRsp, error)
	// 删除订单
	DeleteOrder(context.Context, *DeleteOrderReq) (*CommonRsp, error)
	// START ======================================= 订单退货 ======================================= START
	// 分页查询订单退货申请
	GetOrderReturnApplies(context.Context, *GetOrderReturnAppliesParam) (*GetOrderReturnAppliesRsp, error)
	// 根据id获取订单退货申请
	GetOrderReturnApply(context.Context, *GetOrderReturnApplyReq) (*GetOrderReturnApplyRsp, error)
	// START ======================================= 公司收发货地址 ======================================= START
	// 添加公司收发货地址
	CreateCompanyAddress(context.Context, *AddOrUpdateCompanyAddressParam) (*CommonRsp, error)
	// 修改公司收发货地址
	UpdateCompanyAddress(context.Context, *AddOrUpdateCompanyAddressParam) (*CommonRsp, error)
	// 分页查询公司收发货地址
	GetCompanyAddresses(context.Context, *GetCompanyAddressesParam) (*GetCompanyAddressesRsp, error)
	// 根据id获取公司收发货地址
	GetCompanyAddress(context.Context, *GetCompanyAddressReq) (*GetCompanyAddressRsp, error)
	// 删除公司收发货地址
	DeleteCompanyAddress(context.Context, *DeleteCompanyAddressReq) (*CommonRsp, error)
	mustEmbedUnimplementedOmsAdminApiServer()
}

// UnimplementedOmsAdminApiServer must be embedded to have forward compatible implementations.
type UnimplementedOmsAdminApiServer struct {
}

func (UnimplementedOmsAdminApiServer) CreateOrderReturnReason(context.Context, *AddOrUpdateOrderReturnReasonParam) (*CommonRsp, error) {
	return nil, status.Errorf(codes.Unimplemented, "method CreateOrderReturnReason not implemented")
}
func (UnimplementedOmsAdminApiServer) UpdateOrderReturnReason(context.Context, *AddOrUpdateOrderReturnReasonParam) (*CommonRsp, error) {
	return nil, status.Errorf(codes.Unimplemented, "method UpdateOrderReturnReason not implemented")
}
func (UnimplementedOmsAdminApiServer) GetOrderReturnReasons(context.Context, *GetOrderReturnReasonsParam) (*GetOrderReturnReasonsRsp, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetOrderReturnReasons not implemented")
}
func (UnimplementedOmsAdminApiServer) GetOrderReturnReason(context.Context, *GetOrderReturnReasonReq) (*GetOrderReturnReasonRsp, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetOrderReturnReason not implemented")
}
func (UnimplementedOmsAdminApiServer) DeleteOrderReturnReason(context.Context, *DeleteOrderReturnReasonReq) (*CommonRsp, error) {
	return nil, status.Errorf(codes.Unimplemented, "method DeleteOrderReturnReason not implemented")
}
func (UnimplementedOmsAdminApiServer) GetOrders(context.Context, *GetOrdersParam) (*GetOrdersRsp, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetOrders not implemented")
}
func (UnimplementedOmsAdminApiServer) GetOrder(context.Context, *GetOrderReq) (*GetOrderRsp, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetOrder not implemented")
}
func (UnimplementedOmsAdminApiServer) DeleteOrder(context.Context, *DeleteOrderReq) (*CommonRsp, error) {
	return nil, status.Errorf(codes.Unimplemented, "method DeleteOrder not implemented")
}
func (UnimplementedOmsAdminApiServer) GetOrderReturnApplies(context.Context, *GetOrderReturnAppliesParam) (*GetOrderReturnAppliesRsp, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetOrderReturnApplies not implemented")
}
func (UnimplementedOmsAdminApiServer) GetOrderReturnApply(context.Context, *GetOrderReturnApplyReq) (*GetOrderReturnApplyRsp, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetOrderReturnApply not implemented")
}
func (UnimplementedOmsAdminApiServer) CreateCompanyAddress(context.Context, *AddOrUpdateCompanyAddressParam) (*CommonRsp, error) {
	return nil, status.Errorf(codes.Unimplemented, "method CreateCompanyAddress not implemented")
}
func (UnimplementedOmsAdminApiServer) UpdateCompanyAddress(context.Context, *AddOrUpdateCompanyAddressParam) (*CommonRsp, error) {
	return nil, status.Errorf(codes.Unimplemented, "method UpdateCompanyAddress not implemented")
}
func (UnimplementedOmsAdminApiServer) GetCompanyAddresses(context.Context, *GetCompanyAddressesParam) (*GetCompanyAddressesRsp, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetCompanyAddresses not implemented")
}
func (UnimplementedOmsAdminApiServer) GetCompanyAddress(context.Context, *GetCompanyAddressReq) (*GetCompanyAddressRsp, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetCompanyAddress not implemented")
}
func (UnimplementedOmsAdminApiServer) DeleteCompanyAddress(context.Context, *DeleteCompanyAddressReq) (*CommonRsp, error) {
	return nil, status.Errorf(codes.Unimplemented, "method DeleteCompanyAddress not implemented")
}
func (UnimplementedOmsAdminApiServer) mustEmbedUnimplementedOmsAdminApiServer() {}

// UnsafeOmsAdminApiServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to OmsAdminApiServer will
// result in compilation errors.
type UnsafeOmsAdminApiServer interface {
	mustEmbedUnimplementedOmsAdminApiServer()
}

func RegisterOmsAdminApiServer(s grpc.ServiceRegistrar, srv OmsAdminApiServer) {
	s.RegisterService(&OmsAdminApi_ServiceDesc, srv)
}

func _OmsAdminApi_CreateOrderReturnReason_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(AddOrUpdateOrderReturnReasonParam)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(OmsAdminApiServer).CreateOrderReturnReason(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: OmsAdminApi_CreateOrderReturnReason_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(OmsAdminApiServer).CreateOrderReturnReason(ctx, req.(*AddOrUpdateOrderReturnReasonParam))
	}
	return interceptor(ctx, in, info, handler)
}

func _OmsAdminApi_UpdateOrderReturnReason_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(AddOrUpdateOrderReturnReasonParam)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(OmsAdminApiServer).UpdateOrderReturnReason(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: OmsAdminApi_UpdateOrderReturnReason_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(OmsAdminApiServer).UpdateOrderReturnReason(ctx, req.(*AddOrUpdateOrderReturnReasonParam))
	}
	return interceptor(ctx, in, info, handler)
}

func _OmsAdminApi_GetOrderReturnReasons_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetOrderReturnReasonsParam)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(OmsAdminApiServer).GetOrderReturnReasons(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: OmsAdminApi_GetOrderReturnReasons_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(OmsAdminApiServer).GetOrderReturnReasons(ctx, req.(*GetOrderReturnReasonsParam))
	}
	return interceptor(ctx, in, info, handler)
}

func _OmsAdminApi_GetOrderReturnReason_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetOrderReturnReasonReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(OmsAdminApiServer).GetOrderReturnReason(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: OmsAdminApi_GetOrderReturnReason_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(OmsAdminApiServer).GetOrderReturnReason(ctx, req.(*GetOrderReturnReasonReq))
	}
	return interceptor(ctx, in, info, handler)
}

func _OmsAdminApi_DeleteOrderReturnReason_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(DeleteOrderReturnReasonReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(OmsAdminApiServer).DeleteOrderReturnReason(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: OmsAdminApi_DeleteOrderReturnReason_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(OmsAdminApiServer).DeleteOrderReturnReason(ctx, req.(*DeleteOrderReturnReasonReq))
	}
	return interceptor(ctx, in, info, handler)
}

func _OmsAdminApi_GetOrders_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetOrdersParam)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(OmsAdminApiServer).GetOrders(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: OmsAdminApi_GetOrders_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(OmsAdminApiServer).GetOrders(ctx, req.(*GetOrdersParam))
	}
	return interceptor(ctx, in, info, handler)
}

func _OmsAdminApi_GetOrder_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetOrderReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(OmsAdminApiServer).GetOrder(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: OmsAdminApi_GetOrder_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(OmsAdminApiServer).GetOrder(ctx, req.(*GetOrderReq))
	}
	return interceptor(ctx, in, info, handler)
}

func _OmsAdminApi_DeleteOrder_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(DeleteOrderReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(OmsAdminApiServer).DeleteOrder(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: OmsAdminApi_DeleteOrder_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(OmsAdminApiServer).DeleteOrder(ctx, req.(*DeleteOrderReq))
	}
	return interceptor(ctx, in, info, handler)
}

func _OmsAdminApi_GetOrderReturnApplies_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetOrderReturnAppliesParam)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(OmsAdminApiServer).GetOrderReturnApplies(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: OmsAdminApi_GetOrderReturnApplies_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(OmsAdminApiServer).GetOrderReturnApplies(ctx, req.(*GetOrderReturnAppliesParam))
	}
	return interceptor(ctx, in, info, handler)
}

func _OmsAdminApi_GetOrderReturnApply_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetOrderReturnApplyReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(OmsAdminApiServer).GetOrderReturnApply(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: OmsAdminApi_GetOrderReturnApply_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(OmsAdminApiServer).GetOrderReturnApply(ctx, req.(*GetOrderReturnApplyReq))
	}
	return interceptor(ctx, in, info, handler)
}

func _OmsAdminApi_CreateCompanyAddress_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(AddOrUpdateCompanyAddressParam)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(OmsAdminApiServer).CreateCompanyAddress(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: OmsAdminApi_CreateCompanyAddress_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(OmsAdminApiServer).CreateCompanyAddress(ctx, req.(*AddOrUpdateCompanyAddressParam))
	}
	return interceptor(ctx, in, info, handler)
}

func _OmsAdminApi_UpdateCompanyAddress_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(AddOrUpdateCompanyAddressParam)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(OmsAdminApiServer).UpdateCompanyAddress(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: OmsAdminApi_UpdateCompanyAddress_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(OmsAdminApiServer).UpdateCompanyAddress(ctx, req.(*AddOrUpdateCompanyAddressParam))
	}
	return interceptor(ctx, in, info, handler)
}

func _OmsAdminApi_GetCompanyAddresses_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetCompanyAddressesParam)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(OmsAdminApiServer).GetCompanyAddresses(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: OmsAdminApi_GetCompanyAddresses_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(OmsAdminApiServer).GetCompanyAddresses(ctx, req.(*GetCompanyAddressesParam))
	}
	return interceptor(ctx, in, info, handler)
}

func _OmsAdminApi_GetCompanyAddress_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetCompanyAddressReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(OmsAdminApiServer).GetCompanyAddress(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: OmsAdminApi_GetCompanyAddress_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(OmsAdminApiServer).GetCompanyAddress(ctx, req.(*GetCompanyAddressReq))
	}
	return interceptor(ctx, in, info, handler)
}

func _OmsAdminApi_DeleteCompanyAddress_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(DeleteCompanyAddressReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(OmsAdminApiServer).DeleteCompanyAddress(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: OmsAdminApi_DeleteCompanyAddress_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(OmsAdminApiServer).DeleteCompanyAddress(ctx, req.(*DeleteCompanyAddressReq))
	}
	return interceptor(ctx, in, info, handler)
}

// OmsAdminApi_ServiceDesc is the grpc.ServiceDesc for OmsAdminApi service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var OmsAdminApi_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "admin.OmsAdminApi",
	HandlerType: (*OmsAdminApiServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "CreateOrderReturnReason",
			Handler:    _OmsAdminApi_CreateOrderReturnReason_Handler,
		},
		{
			MethodName: "UpdateOrderReturnReason",
			Handler:    _OmsAdminApi_UpdateOrderReturnReason_Handler,
		},
		{
			MethodName: "GetOrderReturnReasons",
			Handler:    _OmsAdminApi_GetOrderReturnReasons_Handler,
		},
		{
			MethodName: "GetOrderReturnReason",
			Handler:    _OmsAdminApi_GetOrderReturnReason_Handler,
		},
		{
			MethodName: "DeleteOrderReturnReason",
			Handler:    _OmsAdminApi_DeleteOrderReturnReason_Handler,
		},
		{
			MethodName: "GetOrders",
			Handler:    _OmsAdminApi_GetOrders_Handler,
		},
		{
			MethodName: "GetOrder",
			Handler:    _OmsAdminApi_GetOrder_Handler,
		},
		{
			MethodName: "DeleteOrder",
			Handler:    _OmsAdminApi_DeleteOrder_Handler,
		},
		{
			MethodName: "GetOrderReturnApplies",
			Handler:    _OmsAdminApi_GetOrderReturnApplies_Handler,
		},
		{
			MethodName: "GetOrderReturnApply",
			Handler:    _OmsAdminApi_GetOrderReturnApply_Handler,
		},
		{
			MethodName: "CreateCompanyAddress",
			Handler:    _OmsAdminApi_CreateCompanyAddress_Handler,
		},
		{
			MethodName: "UpdateCompanyAddress",
			Handler:    _OmsAdminApi_UpdateCompanyAddress_Handler,
		},
		{
			MethodName: "GetCompanyAddresses",
			Handler:    _OmsAdminApi_GetCompanyAddresses_Handler,
		},
		{
			MethodName: "GetCompanyAddress",
			Handler:    _OmsAdminApi_GetCompanyAddress_Handler,
		},
		{
			MethodName: "DeleteCompanyAddress",
			Handler:    _OmsAdminApi_DeleteCompanyAddress_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "admin/oms_admin.proto",
}
