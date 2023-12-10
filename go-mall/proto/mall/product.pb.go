// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.30.0
// 	protoc        v3.21.12
// source: portal/product.proto

package mall

import (
	_ "github.com/envoyproxy/protoc-gen-validate/validate"
	_ "google.golang.org/genproto/googleapis/api/annotations"
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
	wrapperspb "google.golang.org/protobuf/types/known/wrapperspb"
	reflect "reflect"
	sync "sync"
)

const (
	// Verify that this generated code is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)
	// Verify that runtime/protoimpl is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)

type ProductItem struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// 基本信息
	Id                uint64 `protobuf:"varint,1,opt,name=id,proto3" json:"id,omitempty"`                                                          // 主键
	ProductCategoryId uint64 `protobuf:"varint,2,opt,name=product_category_id,json=productCategoryId,proto3" json:"product_category_id,omitempty"` // 商品分类id
	Name              string `protobuf:"bytes,3,opt,name=name,proto3" json:"name,omitempty"`                                                       // 商品名称
}

func (x *ProductItem) Reset() {
	*x = ProductItem{}
	if protoimpl.UnsafeEnabled {
		mi := &file_portal_product_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ProductItem) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ProductItem) ProtoMessage() {}

func (x *ProductItem) ProtoReflect() protoreflect.Message {
	mi := &file_portal_product_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ProductItem.ProtoReflect.Descriptor instead.
func (*ProductItem) Descriptor() ([]byte, []int) {
	return file_portal_product_proto_rawDescGZIP(), []int{0}
}

func (x *ProductItem) GetId() uint64 {
	if x != nil {
		return x.Id
	}
	return 0
}

func (x *ProductItem) GetProductCategoryId() uint64 {
	if x != nil {
		return x.ProductCategoryId
	}
	return 0
}

func (x *ProductItem) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

type SearchProductReq struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// 条件
	Keyword           string                  `protobuf:"bytes,1,opt,name=keyword,proto3" json:"keyword,omitempty"`                     // 关键词
	BrandId           *wrapperspb.UInt64Value `protobuf:"bytes,2,opt,name=brandId,proto3" json:"brandId,omitempty"`                     // 品牌ID
	ProductCategoryId *wrapperspb.UInt64Value `protobuf:"bytes,3,opt,name=productCategoryId,proto3" json:"productCategoryId,omitempty"` // 商品分类ID
	// 分页排序
	PageNum  uint64 `protobuf:"varint,10,opt,name=pageNum,proto3" json:"pageNum,omitempty"`   //
	PageSize uint64 `protobuf:"varint,11,opt,name=pageSize,proto3" json:"pageSize,omitempty"` //
	Sort     uint32 `protobuf:"varint,12,opt,name=sort,proto3" json:"sort,omitempty"`         // 排序字段:0->按相关度；1->按新品；2->按销量；3->价格从低到高；4->价格从高到低
}

func (x *SearchProductReq) Reset() {
	*x = SearchProductReq{}
	if protoimpl.UnsafeEnabled {
		mi := &file_portal_product_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *SearchProductReq) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*SearchProductReq) ProtoMessage() {}

func (x *SearchProductReq) ProtoReflect() protoreflect.Message {
	mi := &file_portal_product_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use SearchProductReq.ProtoReflect.Descriptor instead.
func (*SearchProductReq) Descriptor() ([]byte, []int) {
	return file_portal_product_proto_rawDescGZIP(), []int{1}
}

func (x *SearchProductReq) GetKeyword() string {
	if x != nil {
		return x.Keyword
	}
	return ""
}

func (x *SearchProductReq) GetBrandId() *wrapperspb.UInt64Value {
	if x != nil {
		return x.BrandId
	}
	return nil
}

func (x *SearchProductReq) GetProductCategoryId() *wrapperspb.UInt64Value {
	if x != nil {
		return x.ProductCategoryId
	}
	return nil
}

func (x *SearchProductReq) GetPageNum() uint64 {
	if x != nil {
		return x.PageNum
	}
	return 0
}

func (x *SearchProductReq) GetPageSize() uint64 {
	if x != nil {
		return x.PageSize
	}
	return 0
}

func (x *SearchProductReq) GetSort() uint32 {
	if x != nil {
		return x.Sort
	}
	return 0
}

type ProductRspItem struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Data      []*ProductItem `protobuf:"bytes,1,rep,name=data,proto3" json:"data,omitempty"`                              // 商品
	Total     uint32         `protobuf:"varint,10,opt,name=total,proto3" json:"total,omitempty"`                          // 总数量
	PageTotal uint32         `protobuf:"varint,11,opt,name=page_total,json=pageTotal,proto3" json:"page_total,omitempty"` // 总页数
	PageSize  uint32         `protobuf:"varint,12,opt,name=page_size,json=pageSize,proto3" json:"page_size,omitempty"`    // 每页的数量
	PageNum   uint32         `protobuf:"varint,13,opt,name=page_num,json=pageNum,proto3" json:"page_num,omitempty"`       // 当前页数
}

func (x *ProductRspItem) Reset() {
	*x = ProductRspItem{}
	if protoimpl.UnsafeEnabled {
		mi := &file_portal_product_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ProductRspItem) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ProductRspItem) ProtoMessage() {}

func (x *ProductRspItem) ProtoReflect() protoreflect.Message {
	mi := &file_portal_product_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ProductRspItem.ProtoReflect.Descriptor instead.
func (*ProductRspItem) Descriptor() ([]byte, []int) {
	return file_portal_product_proto_rawDescGZIP(), []int{2}
}

func (x *ProductRspItem) GetData() []*ProductItem {
	if x != nil {
		return x.Data
	}
	return nil
}

func (x *ProductRspItem) GetTotal() uint32 {
	if x != nil {
		return x.Total
	}
	return 0
}

func (x *ProductRspItem) GetPageTotal() uint32 {
	if x != nil {
		return x.PageTotal
	}
	return 0
}

func (x *ProductRspItem) GetPageSize() uint32 {
	if x != nil {
		return x.PageSize
	}
	return 0
}

func (x *ProductRspItem) GetPageNum() uint32 {
	if x != nil {
		return x.PageNum
	}
	return 0
}

type SearchProductRsp struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Code    uint32          `protobuf:"varint,1,opt,name=code,proto3" json:"code,omitempty"`      // 状态码
	Message string          `protobuf:"bytes,2,opt,name=message,proto3" json:"message,omitempty"` // 提示信息
	Data    *ProductRspItem `protobuf:"bytes,3,opt,name=data,proto3" json:"data,omitempty"`
}

func (x *SearchProductRsp) Reset() {
	*x = SearchProductRsp{}
	if protoimpl.UnsafeEnabled {
		mi := &file_portal_product_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *SearchProductRsp) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*SearchProductRsp) ProtoMessage() {}

func (x *SearchProductRsp) ProtoReflect() protoreflect.Message {
	mi := &file_portal_product_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use SearchProductRsp.ProtoReflect.Descriptor instead.
func (*SearchProductRsp) Descriptor() ([]byte, []int) {
	return file_portal_product_proto_rawDescGZIP(), []int{3}
}

func (x *SearchProductRsp) GetCode() uint32 {
	if x != nil {
		return x.Code
	}
	return 0
}

func (x *SearchProductRsp) GetMessage() string {
	if x != nil {
		return x.Message
	}
	return ""
}

func (x *SearchProductRsp) GetData() *ProductRspItem {
	if x != nil {
		return x.Data
	}
	return nil
}

var File_portal_product_proto protoreflect.FileDescriptor

var file_portal_product_proto_rawDesc = []byte{
	0x0a, 0x14, 0x70, 0x6f, 0x72, 0x74, 0x61, 0x6c, 0x2f, 0x70, 0x72, 0x6f, 0x64, 0x75, 0x63, 0x74,
	0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x05, 0x61, 0x64, 0x6d, 0x69, 0x6e, 0x1a, 0x13, 0x70,
	0x6f, 0x72, 0x74, 0x61, 0x6c, 0x2f, 0x70, 0x6f, 0x72, 0x74, 0x61, 0x6c, 0x2e, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x1a, 0x17, 0x76, 0x61, 0x6c, 0x69, 0x64, 0x61, 0x74, 0x65, 0x2f, 0x76, 0x61, 0x6c,
	0x69, 0x64, 0x61, 0x74, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x1c, 0x67, 0x6f, 0x6f,
	0x67, 0x6c, 0x65, 0x2f, 0x61, 0x70, 0x69, 0x2f, 0x61, 0x6e, 0x6e, 0x6f, 0x74, 0x61, 0x74, 0x69,
	0x6f, 0x6e, 0x73, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x1e, 0x67, 0x6f, 0x6f, 0x67, 0x6c,
	0x65, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2f, 0x77, 0x72, 0x61, 0x70, 0x70,
	0x65, 0x72, 0x73, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0x61, 0x0a, 0x0b, 0x50, 0x72, 0x6f,
	0x64, 0x75, 0x63, 0x74, 0x49, 0x74, 0x65, 0x6d, 0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x01,
	0x20, 0x01, 0x28, 0x04, 0x52, 0x02, 0x69, 0x64, 0x12, 0x2e, 0x0a, 0x13, 0x70, 0x72, 0x6f, 0x64,
	0x75, 0x63, 0x74, 0x5f, 0x63, 0x61, 0x74, 0x65, 0x67, 0x6f, 0x72, 0x79, 0x5f, 0x69, 0x64, 0x18,
	0x02, 0x20, 0x01, 0x28, 0x04, 0x52, 0x11, 0x70, 0x72, 0x6f, 0x64, 0x75, 0x63, 0x74, 0x43, 0x61,
	0x74, 0x65, 0x67, 0x6f, 0x72, 0x79, 0x49, 0x64, 0x12, 0x12, 0x0a, 0x04, 0x6e, 0x61, 0x6d, 0x65,
	0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x22, 0x8b, 0x02, 0x0a,
	0x10, 0x53, 0x65, 0x61, 0x72, 0x63, 0x68, 0x50, 0x72, 0x6f, 0x64, 0x75, 0x63, 0x74, 0x52, 0x65,
	0x71, 0x12, 0x18, 0x0a, 0x07, 0x6b, 0x65, 0x79, 0x77, 0x6f, 0x72, 0x64, 0x18, 0x01, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x07, 0x6b, 0x65, 0x79, 0x77, 0x6f, 0x72, 0x64, 0x12, 0x36, 0x0a, 0x07, 0x62,
	0x72, 0x61, 0x6e, 0x64, 0x49, 0x64, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1c, 0x2e, 0x67,
	0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x55,
	0x49, 0x6e, 0x74, 0x36, 0x34, 0x56, 0x61, 0x6c, 0x75, 0x65, 0x52, 0x07, 0x62, 0x72, 0x61, 0x6e,
	0x64, 0x49, 0x64, 0x12, 0x4a, 0x0a, 0x11, 0x70, 0x72, 0x6f, 0x64, 0x75, 0x63, 0x74, 0x43, 0x61,
	0x74, 0x65, 0x67, 0x6f, 0x72, 0x79, 0x49, 0x64, 0x18, 0x03, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1c,
	0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66,
	0x2e, 0x55, 0x49, 0x6e, 0x74, 0x36, 0x34, 0x56, 0x61, 0x6c, 0x75, 0x65, 0x52, 0x11, 0x70, 0x72,
	0x6f, 0x64, 0x75, 0x63, 0x74, 0x43, 0x61, 0x74, 0x65, 0x67, 0x6f, 0x72, 0x79, 0x49, 0x64, 0x12,
	0x18, 0x0a, 0x07, 0x70, 0x61, 0x67, 0x65, 0x4e, 0x75, 0x6d, 0x18, 0x0a, 0x20, 0x01, 0x28, 0x04,
	0x52, 0x07, 0x70, 0x61, 0x67, 0x65, 0x4e, 0x75, 0x6d, 0x12, 0x1a, 0x0a, 0x08, 0x70, 0x61, 0x67,
	0x65, 0x53, 0x69, 0x7a, 0x65, 0x18, 0x0b, 0x20, 0x01, 0x28, 0x04, 0x52, 0x08, 0x70, 0x61, 0x67,
	0x65, 0x53, 0x69, 0x7a, 0x65, 0x12, 0x23, 0x0a, 0x04, 0x73, 0x6f, 0x72, 0x74, 0x18, 0x0c, 0x20,
	0x01, 0x28, 0x0d, 0x42, 0x0f, 0xfa, 0x42, 0x0c, 0x2a, 0x0a, 0x30, 0x00, 0x30, 0x01, 0x30, 0x02,
	0x30, 0x03, 0x30, 0x04, 0x52, 0x04, 0x73, 0x6f, 0x72, 0x74, 0x22, 0xa5, 0x01, 0x0a, 0x0e, 0x50,
	0x72, 0x6f, 0x64, 0x75, 0x63, 0x74, 0x52, 0x73, 0x70, 0x49, 0x74, 0x65, 0x6d, 0x12, 0x26, 0x0a,
	0x04, 0x64, 0x61, 0x74, 0x61, 0x18, 0x01, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x12, 0x2e, 0x61, 0x64,
	0x6d, 0x69, 0x6e, 0x2e, 0x50, 0x72, 0x6f, 0x64, 0x75, 0x63, 0x74, 0x49, 0x74, 0x65, 0x6d, 0x52,
	0x04, 0x64, 0x61, 0x74, 0x61, 0x12, 0x14, 0x0a, 0x05, 0x74, 0x6f, 0x74, 0x61, 0x6c, 0x18, 0x0a,
	0x20, 0x01, 0x28, 0x0d, 0x52, 0x05, 0x74, 0x6f, 0x74, 0x61, 0x6c, 0x12, 0x1d, 0x0a, 0x0a, 0x70,
	0x61, 0x67, 0x65, 0x5f, 0x74, 0x6f, 0x74, 0x61, 0x6c, 0x18, 0x0b, 0x20, 0x01, 0x28, 0x0d, 0x52,
	0x09, 0x70, 0x61, 0x67, 0x65, 0x54, 0x6f, 0x74, 0x61, 0x6c, 0x12, 0x1b, 0x0a, 0x09, 0x70, 0x61,
	0x67, 0x65, 0x5f, 0x73, 0x69, 0x7a, 0x65, 0x18, 0x0c, 0x20, 0x01, 0x28, 0x0d, 0x52, 0x08, 0x70,
	0x61, 0x67, 0x65, 0x53, 0x69, 0x7a, 0x65, 0x12, 0x19, 0x0a, 0x08, 0x70, 0x61, 0x67, 0x65, 0x5f,
	0x6e, 0x75, 0x6d, 0x18, 0x0d, 0x20, 0x01, 0x28, 0x0d, 0x52, 0x07, 0x70, 0x61, 0x67, 0x65, 0x4e,
	0x75, 0x6d, 0x22, 0x6b, 0x0a, 0x10, 0x53, 0x65, 0x61, 0x72, 0x63, 0x68, 0x50, 0x72, 0x6f, 0x64,
	0x75, 0x63, 0x74, 0x52, 0x73, 0x70, 0x12, 0x12, 0x0a, 0x04, 0x63, 0x6f, 0x64, 0x65, 0x18, 0x01,
	0x20, 0x01, 0x28, 0x0d, 0x52, 0x04, 0x63, 0x6f, 0x64, 0x65, 0x12, 0x18, 0x0a, 0x07, 0x6d, 0x65,
	0x73, 0x73, 0x61, 0x67, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x6d, 0x65, 0x73,
	0x73, 0x61, 0x67, 0x65, 0x12, 0x29, 0x0a, 0x04, 0x64, 0x61, 0x74, 0x61, 0x18, 0x03, 0x20, 0x01,
	0x28, 0x0b, 0x32, 0x15, 0x2e, 0x61, 0x64, 0x6d, 0x69, 0x6e, 0x2e, 0x50, 0x72, 0x6f, 0x64, 0x75,
	0x63, 0x74, 0x52, 0x73, 0x70, 0x49, 0x74, 0x65, 0x6d, 0x52, 0x04, 0x64, 0x61, 0x74, 0x61, 0x32,
	0x6e, 0x0a, 0x10, 0x50, 0x6f, 0x72, 0x74, 0x61, 0x6c, 0x50, 0x72, 0x6f, 0x64, 0x75, 0x63, 0x74,
	0x41, 0x70, 0x69, 0x12, 0x5a, 0x0a, 0x0d, 0x53, 0x65, 0x61, 0x72, 0x63, 0x68, 0x50, 0x72, 0x6f,
	0x64, 0x75, 0x63, 0x74, 0x12, 0x17, 0x2e, 0x61, 0x64, 0x6d, 0x69, 0x6e, 0x2e, 0x53, 0x65, 0x61,
	0x72, 0x63, 0x68, 0x50, 0x72, 0x6f, 0x64, 0x75, 0x63, 0x74, 0x52, 0x65, 0x71, 0x1a, 0x17, 0x2e,
	0x61, 0x64, 0x6d, 0x69, 0x6e, 0x2e, 0x53, 0x65, 0x61, 0x72, 0x63, 0x68, 0x50, 0x72, 0x6f, 0x64,
	0x75, 0x63, 0x74, 0x52, 0x73, 0x70, 0x22, 0x17, 0x82, 0xd3, 0xe4, 0x93, 0x02, 0x11, 0x12, 0x0f,
	0x2f, 0x70, 0x72, 0x6f, 0x64, 0x75, 0x63, 0x74, 0x2f, 0x73, 0x65, 0x61, 0x72, 0x63, 0x68, 0x42,
	0x08, 0x5a, 0x06, 0x2e, 0x3b, 0x6d, 0x61, 0x6c, 0x6c, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x33,
}

var (
	file_portal_product_proto_rawDescOnce sync.Once
	file_portal_product_proto_rawDescData = file_portal_product_proto_rawDesc
)

func file_portal_product_proto_rawDescGZIP() []byte {
	file_portal_product_proto_rawDescOnce.Do(func() {
		file_portal_product_proto_rawDescData = protoimpl.X.CompressGZIP(file_portal_product_proto_rawDescData)
	})
	return file_portal_product_proto_rawDescData
}

var file_portal_product_proto_msgTypes = make([]protoimpl.MessageInfo, 4)
var file_portal_product_proto_goTypes = []interface{}{
	(*ProductItem)(nil),            // 0: admin.ProductItem
	(*SearchProductReq)(nil),       // 1: admin.SearchProductReq
	(*ProductRspItem)(nil),         // 2: admin.ProductRspItem
	(*SearchProductRsp)(nil),       // 3: admin.SearchProductRsp
	(*wrapperspb.UInt64Value)(nil), // 4: google.protobuf.UInt64Value
}
var file_portal_product_proto_depIdxs = []int32{
	4, // 0: admin.SearchProductReq.brandId:type_name -> google.protobuf.UInt64Value
	4, // 1: admin.SearchProductReq.productCategoryId:type_name -> google.protobuf.UInt64Value
	0, // 2: admin.ProductRspItem.data:type_name -> admin.ProductItem
	2, // 3: admin.SearchProductRsp.data:type_name -> admin.ProductRspItem
	1, // 4: admin.PortalProductApi.SearchProduct:input_type -> admin.SearchProductReq
	3, // 5: admin.PortalProductApi.SearchProduct:output_type -> admin.SearchProductRsp
	5, // [5:6] is the sub-list for method output_type
	4, // [4:5] is the sub-list for method input_type
	4, // [4:4] is the sub-list for extension type_name
	4, // [4:4] is the sub-list for extension extendee
	0, // [0:4] is the sub-list for field type_name
}

func init() { file_portal_product_proto_init() }
func file_portal_product_proto_init() {
	if File_portal_product_proto != nil {
		return
	}
	file_portal_portal_proto_init()
	if !protoimpl.UnsafeEnabled {
		file_portal_product_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ProductItem); i {
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
		file_portal_product_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*SearchProductReq); i {
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
		file_portal_product_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ProductRspItem); i {
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
		file_portal_product_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*SearchProductRsp); i {
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
			RawDescriptor: file_portal_product_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   4,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_portal_product_proto_goTypes,
		DependencyIndexes: file_portal_product_proto_depIdxs,
		MessageInfos:      file_portal_product_proto_msgTypes,
	}.Build()
	File_portal_product_proto = out.File
	file_portal_product_proto_rawDesc = nil
	file_portal_product_proto_goTypes = nil
	file_portal_product_proto_depIdxs = nil
}
