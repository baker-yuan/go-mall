// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.3.0
// - protoc             v3.21.12
// source: portal/cart_item.proto

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
	PortalCartItemApi_CartItemAdd_FullMethodName            = "/admin.PortalCartItemApi/CartItemAdd"
	PortalCartItemApi_CartItemList_FullMethodName           = "/admin.PortalCartItemApi/CartItemList"
	PortalCartItemApi_CartItemListPromotion_FullMethodName  = "/admin.PortalCartItemApi/CartItemListPromotion"
	PortalCartItemApi_CartItemUpdateQuantity_FullMethodName = "/admin.PortalCartItemApi/CartItemUpdateQuantity"
	PortalCartItemApi_CartItemGetCartProduct_FullMethodName = "/admin.PortalCartItemApi/CartItemGetCartProduct"
	PortalCartItemApi_CartItemUpdateAttr_FullMethodName     = "/admin.PortalCartItemApi/CartItemUpdateAttr"
	PortalCartItemApi_CartItemDelete_FullMethodName         = "/admin.PortalCartItemApi/CartItemDelete"
	PortalCartItemApi_CartItemClear_FullMethodName          = "/admin.PortalCartItemApi/CartItemClear"
)

// PortalCartItemApiClient is the client API for PortalCartItemApi service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type PortalCartItemApiClient interface {
	// 添加商品到购物车
	CartItemAdd(ctx context.Context, in *CartItemAddReq, opts ...grpc.CallOption) (*CartItemAddRsp, error)
	// 获取当前会员的购物车列表
	CartItemList(ctx context.Context, in *CartItemListReq, opts ...grpc.CallOption) (*CartItemListRsp, error)
	// 获取当前会员的购物车列表,包括促销信息
	CartItemListPromotion(ctx context.Context, in *CartItemListPromotionReq, opts ...grpc.CallOption) (*CartItemListPromotionRsp, error)
	// 修改购物车中指定商品的数量
	CartItemUpdateQuantity(ctx context.Context, in *CartItemUpdateQuantityReq, opts ...grpc.CallOption) (*CartItemUpdateQuantityRsp, error)
	// 获取购物车中指定商品的规格,用于重选规格
	CartItemGetCartProduct(ctx context.Context, in *CartItemGetCartProductReq, opts ...grpc.CallOption) (*CartItemGetCartProductRsp, error)
	// 修改购物车中商品的规格
	CartItemUpdateAttr(ctx context.Context, in *CartItemUpdateAttrReq, opts ...grpc.CallOption) (*CartItemUpdateAttrRsp, error)
	// 删除购物车中的指定商品
	CartItemDelete(ctx context.Context, in *CartItemDeleteReq, opts ...grpc.CallOption) (*CartItemDeleteRsp, error)
	// 清空当前会员的购物车
	CartItemClear(ctx context.Context, in *CartItemClearReq, opts ...grpc.CallOption) (*CartItemClearRsp, error)
}

type portalCartItemApiClient struct {
	cc grpc.ClientConnInterface
}

func NewPortalCartItemApiClient(cc grpc.ClientConnInterface) PortalCartItemApiClient {
	return &portalCartItemApiClient{cc}
}

func (c *portalCartItemApiClient) CartItemAdd(ctx context.Context, in *CartItemAddReq, opts ...grpc.CallOption) (*CartItemAddRsp, error) {
	out := new(CartItemAddRsp)
	err := c.cc.Invoke(ctx, PortalCartItemApi_CartItemAdd_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *portalCartItemApiClient) CartItemList(ctx context.Context, in *CartItemListReq, opts ...grpc.CallOption) (*CartItemListRsp, error) {
	out := new(CartItemListRsp)
	err := c.cc.Invoke(ctx, PortalCartItemApi_CartItemList_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *portalCartItemApiClient) CartItemListPromotion(ctx context.Context, in *CartItemListPromotionReq, opts ...grpc.CallOption) (*CartItemListPromotionRsp, error) {
	out := new(CartItemListPromotionRsp)
	err := c.cc.Invoke(ctx, PortalCartItemApi_CartItemListPromotion_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *portalCartItemApiClient) CartItemUpdateQuantity(ctx context.Context, in *CartItemUpdateQuantityReq, opts ...grpc.CallOption) (*CartItemUpdateQuantityRsp, error) {
	out := new(CartItemUpdateQuantityRsp)
	err := c.cc.Invoke(ctx, PortalCartItemApi_CartItemUpdateQuantity_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *portalCartItemApiClient) CartItemGetCartProduct(ctx context.Context, in *CartItemGetCartProductReq, opts ...grpc.CallOption) (*CartItemGetCartProductRsp, error) {
	out := new(CartItemGetCartProductRsp)
	err := c.cc.Invoke(ctx, PortalCartItemApi_CartItemGetCartProduct_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *portalCartItemApiClient) CartItemUpdateAttr(ctx context.Context, in *CartItemUpdateAttrReq, opts ...grpc.CallOption) (*CartItemUpdateAttrRsp, error) {
	out := new(CartItemUpdateAttrRsp)
	err := c.cc.Invoke(ctx, PortalCartItemApi_CartItemUpdateAttr_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *portalCartItemApiClient) CartItemDelete(ctx context.Context, in *CartItemDeleteReq, opts ...grpc.CallOption) (*CartItemDeleteRsp, error) {
	out := new(CartItemDeleteRsp)
	err := c.cc.Invoke(ctx, PortalCartItemApi_CartItemDelete_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *portalCartItemApiClient) CartItemClear(ctx context.Context, in *CartItemClearReq, opts ...grpc.CallOption) (*CartItemClearRsp, error) {
	out := new(CartItemClearRsp)
	err := c.cc.Invoke(ctx, PortalCartItemApi_CartItemClear_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// PortalCartItemApiServer is the server API for PortalCartItemApi service.
// All implementations must embed UnimplementedPortalCartItemApiServer
// for forward compatibility
type PortalCartItemApiServer interface {
	// 添加商品到购物车
	CartItemAdd(context.Context, *CartItemAddReq) (*CartItemAddRsp, error)
	// 获取当前会员的购物车列表
	CartItemList(context.Context, *CartItemListReq) (*CartItemListRsp, error)
	// 获取当前会员的购物车列表,包括促销信息
	CartItemListPromotion(context.Context, *CartItemListPromotionReq) (*CartItemListPromotionRsp, error)
	// 修改购物车中指定商品的数量
	CartItemUpdateQuantity(context.Context, *CartItemUpdateQuantityReq) (*CartItemUpdateQuantityRsp, error)
	// 获取购物车中指定商品的规格,用于重选规格
	CartItemGetCartProduct(context.Context, *CartItemGetCartProductReq) (*CartItemGetCartProductRsp, error)
	// 修改购物车中商品的规格
	CartItemUpdateAttr(context.Context, *CartItemUpdateAttrReq) (*CartItemUpdateAttrRsp, error)
	// 删除购物车中的指定商品
	CartItemDelete(context.Context, *CartItemDeleteReq) (*CartItemDeleteRsp, error)
	// 清空当前会员的购物车
	CartItemClear(context.Context, *CartItemClearReq) (*CartItemClearRsp, error)
	mustEmbedUnimplementedPortalCartItemApiServer()
}

// UnimplementedPortalCartItemApiServer must be embedded to have forward compatible implementations.
type UnimplementedPortalCartItemApiServer struct {
}

func (UnimplementedPortalCartItemApiServer) CartItemAdd(context.Context, *CartItemAddReq) (*CartItemAddRsp, error) {
	return nil, status.Errorf(codes.Unimplemented, "method CartItemAdd not implemented")
}
func (UnimplementedPortalCartItemApiServer) CartItemList(context.Context, *CartItemListReq) (*CartItemListRsp, error) {
	return nil, status.Errorf(codes.Unimplemented, "method CartItemList not implemented")
}
func (UnimplementedPortalCartItemApiServer) CartItemListPromotion(context.Context, *CartItemListPromotionReq) (*CartItemListPromotionRsp, error) {
	return nil, status.Errorf(codes.Unimplemented, "method CartItemListPromotion not implemented")
}
func (UnimplementedPortalCartItemApiServer) CartItemUpdateQuantity(context.Context, *CartItemUpdateQuantityReq) (*CartItemUpdateQuantityRsp, error) {
	return nil, status.Errorf(codes.Unimplemented, "method CartItemUpdateQuantity not implemented")
}
func (UnimplementedPortalCartItemApiServer) CartItemGetCartProduct(context.Context, *CartItemGetCartProductReq) (*CartItemGetCartProductRsp, error) {
	return nil, status.Errorf(codes.Unimplemented, "method CartItemGetCartProduct not implemented")
}
func (UnimplementedPortalCartItemApiServer) CartItemUpdateAttr(context.Context, *CartItemUpdateAttrReq) (*CartItemUpdateAttrRsp, error) {
	return nil, status.Errorf(codes.Unimplemented, "method CartItemUpdateAttr not implemented")
}
func (UnimplementedPortalCartItemApiServer) CartItemDelete(context.Context, *CartItemDeleteReq) (*CartItemDeleteRsp, error) {
	return nil, status.Errorf(codes.Unimplemented, "method CartItemDelete not implemented")
}
func (UnimplementedPortalCartItemApiServer) CartItemClear(context.Context, *CartItemClearReq) (*CartItemClearRsp, error) {
	return nil, status.Errorf(codes.Unimplemented, "method CartItemClear not implemented")
}
func (UnimplementedPortalCartItemApiServer) mustEmbedUnimplementedPortalCartItemApiServer() {}

// UnsafePortalCartItemApiServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to PortalCartItemApiServer will
// result in compilation errors.
type UnsafePortalCartItemApiServer interface {
	mustEmbedUnimplementedPortalCartItemApiServer()
}

func RegisterPortalCartItemApiServer(s grpc.ServiceRegistrar, srv PortalCartItemApiServer) {
	s.RegisterService(&PortalCartItemApi_ServiceDesc, srv)
}

func _PortalCartItemApi_CartItemAdd_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CartItemAddReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(PortalCartItemApiServer).CartItemAdd(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: PortalCartItemApi_CartItemAdd_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(PortalCartItemApiServer).CartItemAdd(ctx, req.(*CartItemAddReq))
	}
	return interceptor(ctx, in, info, handler)
}

func _PortalCartItemApi_CartItemList_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CartItemListReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(PortalCartItemApiServer).CartItemList(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: PortalCartItemApi_CartItemList_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(PortalCartItemApiServer).CartItemList(ctx, req.(*CartItemListReq))
	}
	return interceptor(ctx, in, info, handler)
}

func _PortalCartItemApi_CartItemListPromotion_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CartItemListPromotionReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(PortalCartItemApiServer).CartItemListPromotion(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: PortalCartItemApi_CartItemListPromotion_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(PortalCartItemApiServer).CartItemListPromotion(ctx, req.(*CartItemListPromotionReq))
	}
	return interceptor(ctx, in, info, handler)
}

func _PortalCartItemApi_CartItemUpdateQuantity_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CartItemUpdateQuantityReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(PortalCartItemApiServer).CartItemUpdateQuantity(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: PortalCartItemApi_CartItemUpdateQuantity_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(PortalCartItemApiServer).CartItemUpdateQuantity(ctx, req.(*CartItemUpdateQuantityReq))
	}
	return interceptor(ctx, in, info, handler)
}

func _PortalCartItemApi_CartItemGetCartProduct_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CartItemGetCartProductReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(PortalCartItemApiServer).CartItemGetCartProduct(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: PortalCartItemApi_CartItemGetCartProduct_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(PortalCartItemApiServer).CartItemGetCartProduct(ctx, req.(*CartItemGetCartProductReq))
	}
	return interceptor(ctx, in, info, handler)
}

func _PortalCartItemApi_CartItemUpdateAttr_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CartItemUpdateAttrReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(PortalCartItemApiServer).CartItemUpdateAttr(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: PortalCartItemApi_CartItemUpdateAttr_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(PortalCartItemApiServer).CartItemUpdateAttr(ctx, req.(*CartItemUpdateAttrReq))
	}
	return interceptor(ctx, in, info, handler)
}

func _PortalCartItemApi_CartItemDelete_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CartItemDeleteReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(PortalCartItemApiServer).CartItemDelete(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: PortalCartItemApi_CartItemDelete_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(PortalCartItemApiServer).CartItemDelete(ctx, req.(*CartItemDeleteReq))
	}
	return interceptor(ctx, in, info, handler)
}

func _PortalCartItemApi_CartItemClear_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CartItemClearReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(PortalCartItemApiServer).CartItemClear(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: PortalCartItemApi_CartItemClear_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(PortalCartItemApiServer).CartItemClear(ctx, req.(*CartItemClearReq))
	}
	return interceptor(ctx, in, info, handler)
}

// PortalCartItemApi_ServiceDesc is the grpc.ServiceDesc for PortalCartItemApi service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var PortalCartItemApi_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "admin.PortalCartItemApi",
	HandlerType: (*PortalCartItemApiServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "CartItemAdd",
			Handler:    _PortalCartItemApi_CartItemAdd_Handler,
		},
		{
			MethodName: "CartItemList",
			Handler:    _PortalCartItemApi_CartItemList_Handler,
		},
		{
			MethodName: "CartItemListPromotion",
			Handler:    _PortalCartItemApi_CartItemListPromotion_Handler,
		},
		{
			MethodName: "CartItemUpdateQuantity",
			Handler:    _PortalCartItemApi_CartItemUpdateQuantity_Handler,
		},
		{
			MethodName: "CartItemGetCartProduct",
			Handler:    _PortalCartItemApi_CartItemGetCartProduct_Handler,
		},
		{
			MethodName: "CartItemUpdateAttr",
			Handler:    _PortalCartItemApi_CartItemUpdateAttr_Handler,
		},
		{
			MethodName: "CartItemDelete",
			Handler:    _PortalCartItemApi_CartItemDelete_Handler,
		},
		{
			MethodName: "CartItemClear",
			Handler:    _PortalCartItemApi_CartItemClear_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "portal/cart_item.proto",
}
