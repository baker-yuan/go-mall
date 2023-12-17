// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.30.0
// 	protoc        v3.21.12
// source: portal/cart_item.proto

package mall

import (
	_ "github.com/envoyproxy/protoc-gen-validate/validate"
	_ "google.golang.org/genproto/googleapis/api/annotations"
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
	_ "google.golang.org/protobuf/types/known/wrapperspb"
	reflect "reflect"
	sync "sync"
)

const (
	// Verify that this generated code is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)
	// Verify that runtime/protoimpl is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)

// 添加商品到购物车
type CartItemAddReq struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// 商品信息
	ProductId uint64 `protobuf:"varint,3,opt,name=product_id,json=productId,proto3" json:"product_id,omitempty"` // 商品id
	// 商品属性
	ProductAttr string `protobuf:"bytes,10,opt,name=product_attr,json=productAttr,proto3" json:"product_attr,omitempty"` // 商品销售属性，JSON 字符串
	// 商品sku
	ProductSkuId   uint64 `protobuf:"varint,11,opt,name=product_sku_id,json=productSkuId,proto3" json:"product_sku_id,omitempty"`      // 商品sku id
	ProductSkuCode string `protobuf:"bytes,12,opt,name=product_sku_code,json=productSkuCode,proto3" json:"product_sku_code,omitempty"` // 商品sku条码
	// 价格数量
	Quantity uint32 `protobuf:"varint,14,opt,name=quantity,proto3" json:"quantity,omitempty"` // 购买数量
}

func (x *CartItemAddReq) Reset() {
	*x = CartItemAddReq{}
	if protoimpl.UnsafeEnabled {
		mi := &file_portal_cart_item_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CartItemAddReq) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CartItemAddReq) ProtoMessage() {}

func (x *CartItemAddReq) ProtoReflect() protoreflect.Message {
	mi := &file_portal_cart_item_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CartItemAddReq.ProtoReflect.Descriptor instead.
func (*CartItemAddReq) Descriptor() ([]byte, []int) {
	return file_portal_cart_item_proto_rawDescGZIP(), []int{0}
}

func (x *CartItemAddReq) GetProductId() uint64 {
	if x != nil {
		return x.ProductId
	}
	return 0
}

func (x *CartItemAddReq) GetProductAttr() string {
	if x != nil {
		return x.ProductAttr
	}
	return ""
}

func (x *CartItemAddReq) GetProductSkuId() uint64 {
	if x != nil {
		return x.ProductSkuId
	}
	return 0
}

func (x *CartItemAddReq) GetProductSkuCode() string {
	if x != nil {
		return x.ProductSkuCode
	}
	return ""
}

func (x *CartItemAddReq) GetQuantity() uint32 {
	if x != nil {
		return x.Quantity
	}
	return 0
}

type CartItemAddRsp struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *CartItemAddRsp) Reset() {
	*x = CartItemAddRsp{}
	if protoimpl.UnsafeEnabled {
		mi := &file_portal_cart_item_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CartItemAddRsp) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CartItemAddRsp) ProtoMessage() {}

func (x *CartItemAddRsp) ProtoReflect() protoreflect.Message {
	mi := &file_portal_cart_item_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CartItemAddRsp.ProtoReflect.Descriptor instead.
func (*CartItemAddRsp) Descriptor() ([]byte, []int) {
	return file_portal_cart_item_proto_rawDescGZIP(), []int{1}
}

// 获取当前会员的购物车列表
type CartItemListReq struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *CartItemListReq) Reset() {
	*x = CartItemListReq{}
	if protoimpl.UnsafeEnabled {
		mi := &file_portal_cart_item_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CartItemListReq) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CartItemListReq) ProtoMessage() {}

func (x *CartItemListReq) ProtoReflect() protoreflect.Message {
	mi := &file_portal_cart_item_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CartItemListReq.ProtoReflect.Descriptor instead.
func (*CartItemListReq) Descriptor() ([]byte, []int) {
	return file_portal_cart_item_proto_rawDescGZIP(), []int{2}
}

type CartItemListRsp struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Code    uint32      `protobuf:"varint,1,opt,name=code,proto3" json:"code,omitempty"`      // 状态码
	Message string      `protobuf:"bytes,2,opt,name=message,proto3" json:"message,omitempty"` // 提示信息
	Data    []*CartItem `protobuf:"bytes,3,rep,name=data,proto3" json:"data,omitempty"`
}

func (x *CartItemListRsp) Reset() {
	*x = CartItemListRsp{}
	if protoimpl.UnsafeEnabled {
		mi := &file_portal_cart_item_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CartItemListRsp) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CartItemListRsp) ProtoMessage() {}

func (x *CartItemListRsp) ProtoReflect() protoreflect.Message {
	mi := &file_portal_cart_item_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CartItemListRsp.ProtoReflect.Descriptor instead.
func (*CartItemListRsp) Descriptor() ([]byte, []int) {
	return file_portal_cart_item_proto_rawDescGZIP(), []int{3}
}

func (x *CartItemListRsp) GetCode() uint32 {
	if x != nil {
		return x.Code
	}
	return 0
}

func (x *CartItemListRsp) GetMessage() string {
	if x != nil {
		return x.Message
	}
	return ""
}

func (x *CartItemListRsp) GetData() []*CartItem {
	if x != nil {
		return x.Data
	}
	return nil
}

// 获取当前会员的购物车列表,包括促销信息
type CartItemListPromotionReq struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *CartItemListPromotionReq) Reset() {
	*x = CartItemListPromotionReq{}
	if protoimpl.UnsafeEnabled {
		mi := &file_portal_cart_item_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CartItemListPromotionReq) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CartItemListPromotionReq) ProtoMessage() {}

func (x *CartItemListPromotionReq) ProtoReflect() protoreflect.Message {
	mi := &file_portal_cart_item_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CartItemListPromotionReq.ProtoReflect.Descriptor instead.
func (*CartItemListPromotionReq) Descriptor() ([]byte, []int) {
	return file_portal_cart_item_proto_rawDescGZIP(), []int{4}
}

type CartItemListPromotionRsp struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *CartItemListPromotionRsp) Reset() {
	*x = CartItemListPromotionRsp{}
	if protoimpl.UnsafeEnabled {
		mi := &file_portal_cart_item_proto_msgTypes[5]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CartItemListPromotionRsp) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CartItemListPromotionRsp) ProtoMessage() {}

func (x *CartItemListPromotionRsp) ProtoReflect() protoreflect.Message {
	mi := &file_portal_cart_item_proto_msgTypes[5]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CartItemListPromotionRsp.ProtoReflect.Descriptor instead.
func (*CartItemListPromotionRsp) Descriptor() ([]byte, []int) {
	return file_portal_cart_item_proto_rawDescGZIP(), []int{5}
}

// 修改购物车中指定商品的数量
type CartItemUpdateQuantityReq struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *CartItemUpdateQuantityReq) Reset() {
	*x = CartItemUpdateQuantityReq{}
	if protoimpl.UnsafeEnabled {
		mi := &file_portal_cart_item_proto_msgTypes[6]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CartItemUpdateQuantityReq) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CartItemUpdateQuantityReq) ProtoMessage() {}

func (x *CartItemUpdateQuantityReq) ProtoReflect() protoreflect.Message {
	mi := &file_portal_cart_item_proto_msgTypes[6]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CartItemUpdateQuantityReq.ProtoReflect.Descriptor instead.
func (*CartItemUpdateQuantityReq) Descriptor() ([]byte, []int) {
	return file_portal_cart_item_proto_rawDescGZIP(), []int{6}
}

type CartItemUpdateQuantityRsp struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *CartItemUpdateQuantityRsp) Reset() {
	*x = CartItemUpdateQuantityRsp{}
	if protoimpl.UnsafeEnabled {
		mi := &file_portal_cart_item_proto_msgTypes[7]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CartItemUpdateQuantityRsp) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CartItemUpdateQuantityRsp) ProtoMessage() {}

func (x *CartItemUpdateQuantityRsp) ProtoReflect() protoreflect.Message {
	mi := &file_portal_cart_item_proto_msgTypes[7]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CartItemUpdateQuantityRsp.ProtoReflect.Descriptor instead.
func (*CartItemUpdateQuantityRsp) Descriptor() ([]byte, []int) {
	return file_portal_cart_item_proto_rawDescGZIP(), []int{7}
}

// 获取购物车中指定商品的规格,用于重选规格
type CartItemGetCartProductReq struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	ProductId uint64 `protobuf:"varint,1,opt,name=product_id,json=productId,proto3" json:"product_id,omitempty"` // 商品ID
}

func (x *CartItemGetCartProductReq) Reset() {
	*x = CartItemGetCartProductReq{}
	if protoimpl.UnsafeEnabled {
		mi := &file_portal_cart_item_proto_msgTypes[8]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CartItemGetCartProductReq) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CartItemGetCartProductReq) ProtoMessage() {}

func (x *CartItemGetCartProductReq) ProtoReflect() protoreflect.Message {
	mi := &file_portal_cart_item_proto_msgTypes[8]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CartItemGetCartProductReq.ProtoReflect.Descriptor instead.
func (*CartItemGetCartProductReq) Descriptor() ([]byte, []int) {
	return file_portal_cart_item_proto_rawDescGZIP(), []int{8}
}

func (x *CartItemGetCartProductReq) GetProductId() uint64 {
	if x != nil {
		return x.ProductId
	}
	return 0
}

type CartItemGetCartProductRsp struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *CartItemGetCartProductRsp) Reset() {
	*x = CartItemGetCartProductRsp{}
	if protoimpl.UnsafeEnabled {
		mi := &file_portal_cart_item_proto_msgTypes[9]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CartItemGetCartProductRsp) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CartItemGetCartProductRsp) ProtoMessage() {}

func (x *CartItemGetCartProductRsp) ProtoReflect() protoreflect.Message {
	mi := &file_portal_cart_item_proto_msgTypes[9]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CartItemGetCartProductRsp.ProtoReflect.Descriptor instead.
func (*CartItemGetCartProductRsp) Descriptor() ([]byte, []int) {
	return file_portal_cart_item_proto_rawDescGZIP(), []int{9}
}

// 修改购物车中商品的规格
type CartItemUpdateAttrReq struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *CartItemUpdateAttrReq) Reset() {
	*x = CartItemUpdateAttrReq{}
	if protoimpl.UnsafeEnabled {
		mi := &file_portal_cart_item_proto_msgTypes[10]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CartItemUpdateAttrReq) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CartItemUpdateAttrReq) ProtoMessage() {}

func (x *CartItemUpdateAttrReq) ProtoReflect() protoreflect.Message {
	mi := &file_portal_cart_item_proto_msgTypes[10]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CartItemUpdateAttrReq.ProtoReflect.Descriptor instead.
func (*CartItemUpdateAttrReq) Descriptor() ([]byte, []int) {
	return file_portal_cart_item_proto_rawDescGZIP(), []int{10}
}

type CartItemUpdateAttrRsp struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *CartItemUpdateAttrRsp) Reset() {
	*x = CartItemUpdateAttrRsp{}
	if protoimpl.UnsafeEnabled {
		mi := &file_portal_cart_item_proto_msgTypes[11]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CartItemUpdateAttrRsp) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CartItemUpdateAttrRsp) ProtoMessage() {}

func (x *CartItemUpdateAttrRsp) ProtoReflect() protoreflect.Message {
	mi := &file_portal_cart_item_proto_msgTypes[11]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CartItemUpdateAttrRsp.ProtoReflect.Descriptor instead.
func (*CartItemUpdateAttrRsp) Descriptor() ([]byte, []int) {
	return file_portal_cart_item_proto_rawDescGZIP(), []int{11}
}

// 删除购物车中的指定商品
type CartItemDeleteReq struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Ids []uint64 `protobuf:"varint,1,rep,packed,name=ids,proto3" json:"ids,omitempty"` // 商品ID集合
}

func (x *CartItemDeleteReq) Reset() {
	*x = CartItemDeleteReq{}
	if protoimpl.UnsafeEnabled {
		mi := &file_portal_cart_item_proto_msgTypes[12]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CartItemDeleteReq) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CartItemDeleteReq) ProtoMessage() {}

func (x *CartItemDeleteReq) ProtoReflect() protoreflect.Message {
	mi := &file_portal_cart_item_proto_msgTypes[12]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CartItemDeleteReq.ProtoReflect.Descriptor instead.
func (*CartItemDeleteReq) Descriptor() ([]byte, []int) {
	return file_portal_cart_item_proto_rawDescGZIP(), []int{12}
}

func (x *CartItemDeleteReq) GetIds() []uint64 {
	if x != nil {
		return x.Ids
	}
	return nil
}

type CartItemDeleteRsp struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *CartItemDeleteRsp) Reset() {
	*x = CartItemDeleteRsp{}
	if protoimpl.UnsafeEnabled {
		mi := &file_portal_cart_item_proto_msgTypes[13]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CartItemDeleteRsp) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CartItemDeleteRsp) ProtoMessage() {}

func (x *CartItemDeleteRsp) ProtoReflect() protoreflect.Message {
	mi := &file_portal_cart_item_proto_msgTypes[13]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CartItemDeleteRsp.ProtoReflect.Descriptor instead.
func (*CartItemDeleteRsp) Descriptor() ([]byte, []int) {
	return file_portal_cart_item_proto_rawDescGZIP(), []int{13}
}

// 清空当前会员的购物车
type CartItemClearReq struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *CartItemClearReq) Reset() {
	*x = CartItemClearReq{}
	if protoimpl.UnsafeEnabled {
		mi := &file_portal_cart_item_proto_msgTypes[14]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CartItemClearReq) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CartItemClearReq) ProtoMessage() {}

func (x *CartItemClearReq) ProtoReflect() protoreflect.Message {
	mi := &file_portal_cart_item_proto_msgTypes[14]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CartItemClearReq.ProtoReflect.Descriptor instead.
func (*CartItemClearReq) Descriptor() ([]byte, []int) {
	return file_portal_cart_item_proto_rawDescGZIP(), []int{14}
}

type CartItemClearRsp struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *CartItemClearRsp) Reset() {
	*x = CartItemClearRsp{}
	if protoimpl.UnsafeEnabled {
		mi := &file_portal_cart_item_proto_msgTypes[15]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CartItemClearRsp) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CartItemClearRsp) ProtoMessage() {}

func (x *CartItemClearRsp) ProtoReflect() protoreflect.Message {
	mi := &file_portal_cart_item_proto_msgTypes[15]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CartItemClearRsp.ProtoReflect.Descriptor instead.
func (*CartItemClearRsp) Descriptor() ([]byte, []int) {
	return file_portal_cart_item_proto_rawDescGZIP(), []int{15}
}

var File_portal_cart_item_proto protoreflect.FileDescriptor

var file_portal_cart_item_proto_rawDesc = []byte{
	0x0a, 0x16, 0x70, 0x6f, 0x72, 0x74, 0x61, 0x6c, 0x2f, 0x63, 0x61, 0x72, 0x74, 0x5f, 0x69, 0x74,
	0x65, 0x6d, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x05, 0x61, 0x64, 0x6d, 0x69, 0x6e, 0x1a,
	0x13, 0x70, 0x6f, 0x72, 0x74, 0x61, 0x6c, 0x2f, 0x70, 0x6f, 0x72, 0x74, 0x61, 0x6c, 0x2e, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x17, 0x76, 0x61, 0x6c, 0x69, 0x64, 0x61, 0x74, 0x65, 0x2f, 0x76,
	0x61, 0x6c, 0x69, 0x64, 0x61, 0x74, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x1c, 0x67,
	0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x61, 0x70, 0x69, 0x2f, 0x61, 0x6e, 0x6e, 0x6f, 0x74, 0x61,
	0x74, 0x69, 0x6f, 0x6e, 0x73, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x1e, 0x67, 0x6f, 0x6f,
	0x67, 0x6c, 0x65, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2f, 0x77, 0x72, 0x61,
	0x70, 0x70, 0x65, 0x72, 0x73, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x19, 0x6d, 0x6f, 0x64,
	0x65, 0x6c, 0x2f, 0x6f, 0x6d, 0x73, 0x5f, 0x63, 0x61, 0x72, 0x74, 0x5f, 0x69, 0x74, 0x65, 0x6d,
	0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0xc7, 0x01, 0x0a, 0x0e, 0x43, 0x61, 0x72, 0x74, 0x49,
	0x74, 0x65, 0x6d, 0x41, 0x64, 0x64, 0x52, 0x65, 0x71, 0x12, 0x26, 0x0a, 0x0a, 0x70, 0x72, 0x6f,
	0x64, 0x75, 0x63, 0x74, 0x5f, 0x69, 0x64, 0x18, 0x03, 0x20, 0x01, 0x28, 0x04, 0x42, 0x07, 0xfa,
	0x42, 0x04, 0x32, 0x02, 0x28, 0x00, 0x52, 0x09, 0x70, 0x72, 0x6f, 0x64, 0x75, 0x63, 0x74, 0x49,
	0x64, 0x12, 0x21, 0x0a, 0x0c, 0x70, 0x72, 0x6f, 0x64, 0x75, 0x63, 0x74, 0x5f, 0x61, 0x74, 0x74,
	0x72, 0x18, 0x0a, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0b, 0x70, 0x72, 0x6f, 0x64, 0x75, 0x63, 0x74,
	0x41, 0x74, 0x74, 0x72, 0x12, 0x24, 0x0a, 0x0e, 0x70, 0x72, 0x6f, 0x64, 0x75, 0x63, 0x74, 0x5f,
	0x73, 0x6b, 0x75, 0x5f, 0x69, 0x64, 0x18, 0x0b, 0x20, 0x01, 0x28, 0x04, 0x52, 0x0c, 0x70, 0x72,
	0x6f, 0x64, 0x75, 0x63, 0x74, 0x53, 0x6b, 0x75, 0x49, 0x64, 0x12, 0x28, 0x0a, 0x10, 0x70, 0x72,
	0x6f, 0x64, 0x75, 0x63, 0x74, 0x5f, 0x73, 0x6b, 0x75, 0x5f, 0x63, 0x6f, 0x64, 0x65, 0x18, 0x0c,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x0e, 0x70, 0x72, 0x6f, 0x64, 0x75, 0x63, 0x74, 0x53, 0x6b, 0x75,
	0x43, 0x6f, 0x64, 0x65, 0x12, 0x1a, 0x0a, 0x08, 0x71, 0x75, 0x61, 0x6e, 0x74, 0x69, 0x74, 0x79,
	0x18, 0x0e, 0x20, 0x01, 0x28, 0x0d, 0x52, 0x08, 0x71, 0x75, 0x61, 0x6e, 0x74, 0x69, 0x74, 0x79,
	0x22, 0x10, 0x0a, 0x0e, 0x43, 0x61, 0x72, 0x74, 0x49, 0x74, 0x65, 0x6d, 0x41, 0x64, 0x64, 0x52,
	0x73, 0x70, 0x22, 0x11, 0x0a, 0x0f, 0x43, 0x61, 0x72, 0x74, 0x49, 0x74, 0x65, 0x6d, 0x4c, 0x69,
	0x73, 0x74, 0x52, 0x65, 0x71, 0x22, 0x64, 0x0a, 0x0f, 0x43, 0x61, 0x72, 0x74, 0x49, 0x74, 0x65,
	0x6d, 0x4c, 0x69, 0x73, 0x74, 0x52, 0x73, 0x70, 0x12, 0x12, 0x0a, 0x04, 0x63, 0x6f, 0x64, 0x65,
	0x18, 0x01, 0x20, 0x01, 0x28, 0x0d, 0x52, 0x04, 0x63, 0x6f, 0x64, 0x65, 0x12, 0x18, 0x0a, 0x07,
	0x6d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x6d,
	0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x12, 0x23, 0x0a, 0x04, 0x64, 0x61, 0x74, 0x61, 0x18, 0x03,
	0x20, 0x03, 0x28, 0x0b, 0x32, 0x0f, 0x2e, 0x6d, 0x6f, 0x64, 0x65, 0x6c, 0x2e, 0x43, 0x61, 0x72,
	0x74, 0x49, 0x74, 0x65, 0x6d, 0x52, 0x04, 0x64, 0x61, 0x74, 0x61, 0x22, 0x1a, 0x0a, 0x18, 0x43,
	0x61, 0x72, 0x74, 0x49, 0x74, 0x65, 0x6d, 0x4c, 0x69, 0x73, 0x74, 0x50, 0x72, 0x6f, 0x6d, 0x6f,
	0x74, 0x69, 0x6f, 0x6e, 0x52, 0x65, 0x71, 0x22, 0x1a, 0x0a, 0x18, 0x43, 0x61, 0x72, 0x74, 0x49,
	0x74, 0x65, 0x6d, 0x4c, 0x69, 0x73, 0x74, 0x50, 0x72, 0x6f, 0x6d, 0x6f, 0x74, 0x69, 0x6f, 0x6e,
	0x52, 0x73, 0x70, 0x22, 0x1b, 0x0a, 0x19, 0x43, 0x61, 0x72, 0x74, 0x49, 0x74, 0x65, 0x6d, 0x55,
	0x70, 0x64, 0x61, 0x74, 0x65, 0x51, 0x75, 0x61, 0x6e, 0x74, 0x69, 0x74, 0x79, 0x52, 0x65, 0x71,
	0x22, 0x1b, 0x0a, 0x19, 0x43, 0x61, 0x72, 0x74, 0x49, 0x74, 0x65, 0x6d, 0x55, 0x70, 0x64, 0x61,
	0x74, 0x65, 0x51, 0x75, 0x61, 0x6e, 0x74, 0x69, 0x74, 0x79, 0x52, 0x73, 0x70, 0x22, 0x3a, 0x0a,
	0x19, 0x43, 0x61, 0x72, 0x74, 0x49, 0x74, 0x65, 0x6d, 0x47, 0x65, 0x74, 0x43, 0x61, 0x72, 0x74,
	0x50, 0x72, 0x6f, 0x64, 0x75, 0x63, 0x74, 0x52, 0x65, 0x71, 0x12, 0x1d, 0x0a, 0x0a, 0x70, 0x72,
	0x6f, 0x64, 0x75, 0x63, 0x74, 0x5f, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x04, 0x52, 0x09,
	0x70, 0x72, 0x6f, 0x64, 0x75, 0x63, 0x74, 0x49, 0x64, 0x22, 0x1b, 0x0a, 0x19, 0x43, 0x61, 0x72,
	0x74, 0x49, 0x74, 0x65, 0x6d, 0x47, 0x65, 0x74, 0x43, 0x61, 0x72, 0x74, 0x50, 0x72, 0x6f, 0x64,
	0x75, 0x63, 0x74, 0x52, 0x73, 0x70, 0x22, 0x17, 0x0a, 0x15, 0x43, 0x61, 0x72, 0x74, 0x49, 0x74,
	0x65, 0x6d, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x41, 0x74, 0x74, 0x72, 0x52, 0x65, 0x71, 0x22,
	0x17, 0x0a, 0x15, 0x43, 0x61, 0x72, 0x74, 0x49, 0x74, 0x65, 0x6d, 0x55, 0x70, 0x64, 0x61, 0x74,
	0x65, 0x41, 0x74, 0x74, 0x72, 0x52, 0x73, 0x70, 0x22, 0x25, 0x0a, 0x11, 0x43, 0x61, 0x72, 0x74,
	0x49, 0x74, 0x65, 0x6d, 0x44, 0x65, 0x6c, 0x65, 0x74, 0x65, 0x52, 0x65, 0x71, 0x12, 0x10, 0x0a,
	0x03, 0x69, 0x64, 0x73, 0x18, 0x01, 0x20, 0x03, 0x28, 0x04, 0x52, 0x03, 0x69, 0x64, 0x73, 0x22,
	0x13, 0x0a, 0x11, 0x43, 0x61, 0x72, 0x74, 0x49, 0x74, 0x65, 0x6d, 0x44, 0x65, 0x6c, 0x65, 0x74,
	0x65, 0x52, 0x73, 0x70, 0x22, 0x12, 0x0a, 0x10, 0x43, 0x61, 0x72, 0x74, 0x49, 0x74, 0x65, 0x6d,
	0x43, 0x6c, 0x65, 0x61, 0x72, 0x52, 0x65, 0x71, 0x22, 0x12, 0x0a, 0x10, 0x43, 0x61, 0x72, 0x74,
	0x49, 0x74, 0x65, 0x6d, 0x43, 0x6c, 0x65, 0x61, 0x72, 0x52, 0x73, 0x70, 0x32, 0xe3, 0x06, 0x0a,
	0x11, 0x50, 0x6f, 0x72, 0x74, 0x61, 0x6c, 0x43, 0x61, 0x72, 0x74, 0x49, 0x74, 0x65, 0x6d, 0x41,
	0x70, 0x69, 0x12, 0x51, 0x0a, 0x0b, 0x43, 0x61, 0x72, 0x74, 0x49, 0x74, 0x65, 0x6d, 0x41, 0x64,
	0x64, 0x12, 0x15, 0x2e, 0x61, 0x64, 0x6d, 0x69, 0x6e, 0x2e, 0x43, 0x61, 0x72, 0x74, 0x49, 0x74,
	0x65, 0x6d, 0x41, 0x64, 0x64, 0x52, 0x65, 0x71, 0x1a, 0x15, 0x2e, 0x61, 0x64, 0x6d, 0x69, 0x6e,
	0x2e, 0x43, 0x61, 0x72, 0x74, 0x49, 0x74, 0x65, 0x6d, 0x41, 0x64, 0x64, 0x52, 0x73, 0x70, 0x22,
	0x14, 0x82, 0xd3, 0xe4, 0x93, 0x02, 0x0e, 0x3a, 0x01, 0x2a, 0x22, 0x09, 0x2f, 0x63, 0x61, 0x72,
	0x74, 0x2f, 0x61, 0x64, 0x64, 0x12, 0x52, 0x0a, 0x0c, 0x43, 0x61, 0x72, 0x74, 0x49, 0x74, 0x65,
	0x6d, 0x4c, 0x69, 0x73, 0x74, 0x12, 0x16, 0x2e, 0x61, 0x64, 0x6d, 0x69, 0x6e, 0x2e, 0x43, 0x61,
	0x72, 0x74, 0x49, 0x74, 0x65, 0x6d, 0x4c, 0x69, 0x73, 0x74, 0x52, 0x65, 0x71, 0x1a, 0x16, 0x2e,
	0x61, 0x64, 0x6d, 0x69, 0x6e, 0x2e, 0x43, 0x61, 0x72, 0x74, 0x49, 0x74, 0x65, 0x6d, 0x4c, 0x69,
	0x73, 0x74, 0x52, 0x73, 0x70, 0x22, 0x12, 0x82, 0xd3, 0xe4, 0x93, 0x02, 0x0c, 0x12, 0x0a, 0x2f,
	0x63, 0x61, 0x72, 0x74, 0x2f, 0x6c, 0x69, 0x73, 0x74, 0x12, 0x77, 0x0a, 0x15, 0x43, 0x61, 0x72,
	0x74, 0x49, 0x74, 0x65, 0x6d, 0x4c, 0x69, 0x73, 0x74, 0x50, 0x72, 0x6f, 0x6d, 0x6f, 0x74, 0x69,
	0x6f, 0x6e, 0x12, 0x1f, 0x2e, 0x61, 0x64, 0x6d, 0x69, 0x6e, 0x2e, 0x43, 0x61, 0x72, 0x74, 0x49,
	0x74, 0x65, 0x6d, 0x4c, 0x69, 0x73, 0x74, 0x50, 0x72, 0x6f, 0x6d, 0x6f, 0x74, 0x69, 0x6f, 0x6e,
	0x52, 0x65, 0x71, 0x1a, 0x1f, 0x2e, 0x61, 0x64, 0x6d, 0x69, 0x6e, 0x2e, 0x43, 0x61, 0x72, 0x74,
	0x49, 0x74, 0x65, 0x6d, 0x4c, 0x69, 0x73, 0x74, 0x50, 0x72, 0x6f, 0x6d, 0x6f, 0x74, 0x69, 0x6f,
	0x6e, 0x52, 0x73, 0x70, 0x22, 0x1c, 0x82, 0xd3, 0xe4, 0x93, 0x02, 0x16, 0x12, 0x14, 0x2f, 0x63,
	0x61, 0x72, 0x74, 0x2f, 0x6c, 0x69, 0x73, 0x74, 0x2f, 0x70, 0x72, 0x6f, 0x6d, 0x6f, 0x74, 0x69,
	0x6f, 0x6e, 0x12, 0x7e, 0x0a, 0x16, 0x43, 0x61, 0x72, 0x74, 0x49, 0x74, 0x65, 0x6d, 0x55, 0x70,
	0x64, 0x61, 0x74, 0x65, 0x51, 0x75, 0x61, 0x6e, 0x74, 0x69, 0x74, 0x79, 0x12, 0x20, 0x2e, 0x61,
	0x64, 0x6d, 0x69, 0x6e, 0x2e, 0x43, 0x61, 0x72, 0x74, 0x49, 0x74, 0x65, 0x6d, 0x55, 0x70, 0x64,
	0x61, 0x74, 0x65, 0x51, 0x75, 0x61, 0x6e, 0x74, 0x69, 0x74, 0x79, 0x52, 0x65, 0x71, 0x1a, 0x20,
	0x2e, 0x61, 0x64, 0x6d, 0x69, 0x6e, 0x2e, 0x43, 0x61, 0x72, 0x74, 0x49, 0x74, 0x65, 0x6d, 0x55,
	0x70, 0x64, 0x61, 0x74, 0x65, 0x51, 0x75, 0x61, 0x6e, 0x74, 0x69, 0x74, 0x79, 0x52, 0x73, 0x70,
	0x22, 0x20, 0x82, 0xd3, 0xe4, 0x93, 0x02, 0x1a, 0x3a, 0x01, 0x2a, 0x22, 0x15, 0x2f, 0x63, 0x61,
	0x72, 0x74, 0x2f, 0x75, 0x70, 0x64, 0x61, 0x74, 0x65, 0x2f, 0x71, 0x75, 0x61, 0x6e, 0x74, 0x69,
	0x74, 0x79, 0x12, 0x83, 0x01, 0x0a, 0x16, 0x43, 0x61, 0x72, 0x74, 0x49, 0x74, 0x65, 0x6d, 0x47,
	0x65, 0x74, 0x43, 0x61, 0x72, 0x74, 0x50, 0x72, 0x6f, 0x64, 0x75, 0x63, 0x74, 0x12, 0x20, 0x2e,
	0x61, 0x64, 0x6d, 0x69, 0x6e, 0x2e, 0x43, 0x61, 0x72, 0x74, 0x49, 0x74, 0x65, 0x6d, 0x47, 0x65,
	0x74, 0x43, 0x61, 0x72, 0x74, 0x50, 0x72, 0x6f, 0x64, 0x75, 0x63, 0x74, 0x52, 0x65, 0x71, 0x1a,
	0x20, 0x2e, 0x61, 0x64, 0x6d, 0x69, 0x6e, 0x2e, 0x43, 0x61, 0x72, 0x74, 0x49, 0x74, 0x65, 0x6d,
	0x47, 0x65, 0x74, 0x43, 0x61, 0x72, 0x74, 0x50, 0x72, 0x6f, 0x64, 0x75, 0x63, 0x74, 0x52, 0x73,
	0x70, 0x22, 0x25, 0x82, 0xd3, 0xe4, 0x93, 0x02, 0x1f, 0x12, 0x1d, 0x2f, 0x63, 0x61, 0x72, 0x74,
	0x2f, 0x67, 0x65, 0x74, 0x50, 0x72, 0x6f, 0x64, 0x75, 0x63, 0x74, 0x2f, 0x7b, 0x70, 0x72, 0x6f,
	0x64, 0x75, 0x63, 0x74, 0x5f, 0x69, 0x64, 0x7d, 0x12, 0x6e, 0x0a, 0x12, 0x43, 0x61, 0x72, 0x74,
	0x49, 0x74, 0x65, 0x6d, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x41, 0x74, 0x74, 0x72, 0x12, 0x1c,
	0x2e, 0x61, 0x64, 0x6d, 0x69, 0x6e, 0x2e, 0x43, 0x61, 0x72, 0x74, 0x49, 0x74, 0x65, 0x6d, 0x55,
	0x70, 0x64, 0x61, 0x74, 0x65, 0x41, 0x74, 0x74, 0x72, 0x52, 0x65, 0x71, 0x1a, 0x1c, 0x2e, 0x61,
	0x64, 0x6d, 0x69, 0x6e, 0x2e, 0x43, 0x61, 0x72, 0x74, 0x49, 0x74, 0x65, 0x6d, 0x55, 0x70, 0x64,
	0x61, 0x74, 0x65, 0x41, 0x74, 0x74, 0x72, 0x52, 0x73, 0x70, 0x22, 0x1c, 0x82, 0xd3, 0xe4, 0x93,
	0x02, 0x16, 0x3a, 0x01, 0x2a, 0x22, 0x11, 0x2f, 0x63, 0x61, 0x72, 0x74, 0x2f, 0x75, 0x70, 0x64,
	0x61, 0x74, 0x65, 0x2f, 0x61, 0x74, 0x74, 0x72, 0x12, 0x5d, 0x0a, 0x0e, 0x43, 0x61, 0x72, 0x74,
	0x49, 0x74, 0x65, 0x6d, 0x44, 0x65, 0x6c, 0x65, 0x74, 0x65, 0x12, 0x18, 0x2e, 0x61, 0x64, 0x6d,
	0x69, 0x6e, 0x2e, 0x43, 0x61, 0x72, 0x74, 0x49, 0x74, 0x65, 0x6d, 0x44, 0x65, 0x6c, 0x65, 0x74,
	0x65, 0x52, 0x65, 0x71, 0x1a, 0x18, 0x2e, 0x61, 0x64, 0x6d, 0x69, 0x6e, 0x2e, 0x43, 0x61, 0x72,
	0x74, 0x49, 0x74, 0x65, 0x6d, 0x44, 0x65, 0x6c, 0x65, 0x74, 0x65, 0x52, 0x73, 0x70, 0x22, 0x17,
	0x82, 0xd3, 0xe4, 0x93, 0x02, 0x11, 0x3a, 0x01, 0x2a, 0x22, 0x0c, 0x2f, 0x63, 0x61, 0x72, 0x74,
	0x2f, 0x64, 0x65, 0x6c, 0x65, 0x74, 0x65, 0x12, 0x59, 0x0a, 0x0d, 0x43, 0x61, 0x72, 0x74, 0x49,
	0x74, 0x65, 0x6d, 0x43, 0x6c, 0x65, 0x61, 0x72, 0x12, 0x17, 0x2e, 0x61, 0x64, 0x6d, 0x69, 0x6e,
	0x2e, 0x43, 0x61, 0x72, 0x74, 0x49, 0x74, 0x65, 0x6d, 0x43, 0x6c, 0x65, 0x61, 0x72, 0x52, 0x65,
	0x71, 0x1a, 0x17, 0x2e, 0x61, 0x64, 0x6d, 0x69, 0x6e, 0x2e, 0x43, 0x61, 0x72, 0x74, 0x49, 0x74,
	0x65, 0x6d, 0x43, 0x6c, 0x65, 0x61, 0x72, 0x52, 0x73, 0x70, 0x22, 0x16, 0x82, 0xd3, 0xe4, 0x93,
	0x02, 0x10, 0x3a, 0x01, 0x2a, 0x22, 0x0b, 0x2f, 0x63, 0x61, 0x72, 0x74, 0x2f, 0x63, 0x6c, 0x65,
	0x61, 0x72, 0x42, 0x08, 0x5a, 0x06, 0x2e, 0x3b, 0x6d, 0x61, 0x6c, 0x6c, 0x62, 0x06, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_portal_cart_item_proto_rawDescOnce sync.Once
	file_portal_cart_item_proto_rawDescData = file_portal_cart_item_proto_rawDesc
)

func file_portal_cart_item_proto_rawDescGZIP() []byte {
	file_portal_cart_item_proto_rawDescOnce.Do(func() {
		file_portal_cart_item_proto_rawDescData = protoimpl.X.CompressGZIP(file_portal_cart_item_proto_rawDescData)
	})
	return file_portal_cart_item_proto_rawDescData
}

var file_portal_cart_item_proto_msgTypes = make([]protoimpl.MessageInfo, 16)
var file_portal_cart_item_proto_goTypes = []interface{}{
	(*CartItemAddReq)(nil),            // 0: admin.CartItemAddReq
	(*CartItemAddRsp)(nil),            // 1: admin.CartItemAddRsp
	(*CartItemListReq)(nil),           // 2: admin.CartItemListReq
	(*CartItemListRsp)(nil),           // 3: admin.CartItemListRsp
	(*CartItemListPromotionReq)(nil),  // 4: admin.CartItemListPromotionReq
	(*CartItemListPromotionRsp)(nil),  // 5: admin.CartItemListPromotionRsp
	(*CartItemUpdateQuantityReq)(nil), // 6: admin.CartItemUpdateQuantityReq
	(*CartItemUpdateQuantityRsp)(nil), // 7: admin.CartItemUpdateQuantityRsp
	(*CartItemGetCartProductReq)(nil), // 8: admin.CartItemGetCartProductReq
	(*CartItemGetCartProductRsp)(nil), // 9: admin.CartItemGetCartProductRsp
	(*CartItemUpdateAttrReq)(nil),     // 10: admin.CartItemUpdateAttrReq
	(*CartItemUpdateAttrRsp)(nil),     // 11: admin.CartItemUpdateAttrRsp
	(*CartItemDeleteReq)(nil),         // 12: admin.CartItemDeleteReq
	(*CartItemDeleteRsp)(nil),         // 13: admin.CartItemDeleteRsp
	(*CartItemClearReq)(nil),          // 14: admin.CartItemClearReq
	(*CartItemClearRsp)(nil),          // 15: admin.CartItemClearRsp
	(*CartItem)(nil),                  // 16: model.CartItem
}
var file_portal_cart_item_proto_depIdxs = []int32{
	16, // 0: admin.CartItemListRsp.data:type_name -> model.CartItem
	0,  // 1: admin.PortalCartItemApi.CartItemAdd:input_type -> admin.CartItemAddReq
	2,  // 2: admin.PortalCartItemApi.CartItemList:input_type -> admin.CartItemListReq
	4,  // 3: admin.PortalCartItemApi.CartItemListPromotion:input_type -> admin.CartItemListPromotionReq
	6,  // 4: admin.PortalCartItemApi.CartItemUpdateQuantity:input_type -> admin.CartItemUpdateQuantityReq
	8,  // 5: admin.PortalCartItemApi.CartItemGetCartProduct:input_type -> admin.CartItemGetCartProductReq
	10, // 6: admin.PortalCartItemApi.CartItemUpdateAttr:input_type -> admin.CartItemUpdateAttrReq
	12, // 7: admin.PortalCartItemApi.CartItemDelete:input_type -> admin.CartItemDeleteReq
	14, // 8: admin.PortalCartItemApi.CartItemClear:input_type -> admin.CartItemClearReq
	1,  // 9: admin.PortalCartItemApi.CartItemAdd:output_type -> admin.CartItemAddRsp
	3,  // 10: admin.PortalCartItemApi.CartItemList:output_type -> admin.CartItemListRsp
	5,  // 11: admin.PortalCartItemApi.CartItemListPromotion:output_type -> admin.CartItemListPromotionRsp
	7,  // 12: admin.PortalCartItemApi.CartItemUpdateQuantity:output_type -> admin.CartItemUpdateQuantityRsp
	9,  // 13: admin.PortalCartItemApi.CartItemGetCartProduct:output_type -> admin.CartItemGetCartProductRsp
	11, // 14: admin.PortalCartItemApi.CartItemUpdateAttr:output_type -> admin.CartItemUpdateAttrRsp
	13, // 15: admin.PortalCartItemApi.CartItemDelete:output_type -> admin.CartItemDeleteRsp
	15, // 16: admin.PortalCartItemApi.CartItemClear:output_type -> admin.CartItemClearRsp
	9,  // [9:17] is the sub-list for method output_type
	1,  // [1:9] is the sub-list for method input_type
	1,  // [1:1] is the sub-list for extension type_name
	1,  // [1:1] is the sub-list for extension extendee
	0,  // [0:1] is the sub-list for field type_name
}

func init() { file_portal_cart_item_proto_init() }
func file_portal_cart_item_proto_init() {
	if File_portal_cart_item_proto != nil {
		return
	}
	file_portal_portal_proto_init()
	file_model_oms_cart_item_proto_init()
	if !protoimpl.UnsafeEnabled {
		file_portal_cart_item_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*CartItemAddReq); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_portal_cart_item_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*CartItemAddRsp); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_portal_cart_item_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*CartItemListReq); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_portal_cart_item_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*CartItemListRsp); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_portal_cart_item_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*CartItemListPromotionReq); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_portal_cart_item_proto_msgTypes[5].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*CartItemListPromotionRsp); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_portal_cart_item_proto_msgTypes[6].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*CartItemUpdateQuantityReq); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_portal_cart_item_proto_msgTypes[7].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*CartItemUpdateQuantityRsp); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_portal_cart_item_proto_msgTypes[8].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*CartItemGetCartProductReq); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_portal_cart_item_proto_msgTypes[9].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*CartItemGetCartProductRsp); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_portal_cart_item_proto_msgTypes[10].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*CartItemUpdateAttrReq); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_portal_cart_item_proto_msgTypes[11].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*CartItemUpdateAttrRsp); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_portal_cart_item_proto_msgTypes[12].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*CartItemDeleteReq); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_portal_cart_item_proto_msgTypes[13].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*CartItemDeleteRsp); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_portal_cart_item_proto_msgTypes[14].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*CartItemClearReq); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_portal_cart_item_proto_msgTypes[15].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*CartItemClearRsp); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_portal_cart_item_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   16,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_portal_cart_item_proto_goTypes,
		DependencyIndexes: file_portal_cart_item_proto_depIdxs,
		MessageInfos:      file_portal_cart_item_proto_msgTypes,
	}.Build()
	File_portal_cart_item_proto = out.File
	file_portal_cart_item_proto_rawDesc = nil
	file_portal_cart_item_proto_goTypes = nil
	file_portal_cart_item_proto_depIdxs = nil
}
