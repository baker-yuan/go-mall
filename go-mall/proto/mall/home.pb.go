// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.30.0
// 	protoc        v3.21.12
// source: portal/home.proto

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

// 首页内容信息展示
type HomeContentReq struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *HomeContentReq) Reset() {
	*x = HomeContentReq{}
	if protoimpl.UnsafeEnabled {
		mi := &file_portal_home_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *HomeContentReq) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*HomeContentReq) ProtoMessage() {}

func (x *HomeContentReq) ProtoReflect() protoreflect.Message {
	mi := &file_portal_home_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use HomeContentReq.ProtoReflect.Descriptor instead.
func (*HomeContentReq) Descriptor() ([]byte, []int) {
	return file_portal_home_proto_rawDescGZIP(), []int{0}
}

type HomeContentRsp struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Advertises *HomeContentRsp_HomeAdvertise `protobuf:"bytes,1,opt,name=advertises,proto3" json:"advertises,omitempty"` // 轮播广告
}

func (x *HomeContentRsp) Reset() {
	*x = HomeContentRsp{}
	if protoimpl.UnsafeEnabled {
		mi := &file_portal_home_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *HomeContentRsp) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*HomeContentRsp) ProtoMessage() {}

func (x *HomeContentRsp) ProtoReflect() protoreflect.Message {
	mi := &file_portal_home_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use HomeContentRsp.ProtoReflect.Descriptor instead.
func (*HomeContentRsp) Descriptor() ([]byte, []int) {
	return file_portal_home_proto_rawDescGZIP(), []int{1}
}

func (x *HomeContentRsp) GetAdvertises() *HomeContentRsp_HomeAdvertise {
	if x != nil {
		return x.Advertises
	}
	return nil
}

// 商品分类
type ProductCategoryItem struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id       uint64 `protobuf:"varint,1,opt,name=id,proto3" json:"id,omitempty"`                             //
	ParentId uint64 `protobuf:"varint,2,opt,name=parent_id,json=parentId,proto3" json:"parent_id,omitempty"` // 上级分类的编号：0表示一级分类
	Name     string `protobuf:"bytes,3,opt,name=name,proto3" json:"name,omitempty"`                          // 名称
	Icon     string `protobuf:"bytes,10,opt,name=icon,proto3" json:"icon,omitempty"`                         // 图标
}

func (x *ProductCategoryItem) Reset() {
	*x = ProductCategoryItem{}
	if protoimpl.UnsafeEnabled {
		mi := &file_portal_home_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ProductCategoryItem) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ProductCategoryItem) ProtoMessage() {}

func (x *ProductCategoryItem) ProtoReflect() protoreflect.Message {
	mi := &file_portal_home_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ProductCategoryItem.ProtoReflect.Descriptor instead.
func (*ProductCategoryItem) Descriptor() ([]byte, []int) {
	return file_portal_home_proto_rawDescGZIP(), []int{2}
}

func (x *ProductCategoryItem) GetId() uint64 {
	if x != nil {
		return x.Id
	}
	return 0
}

func (x *ProductCategoryItem) GetParentId() uint64 {
	if x != nil {
		return x.ParentId
	}
	return 0
}

func (x *ProductCategoryItem) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *ProductCategoryItem) GetIcon() string {
	if x != nil {
		return x.Icon
	}
	return ""
}

// 获取首页商品分类
type ProductCategoryListReq struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	ParentId uint64 `protobuf:"varint,1,opt,name=parentId,proto3" json:"parentId,omitempty"` // 上级分类的编号：0表示一级分类
}

func (x *ProductCategoryListReq) Reset() {
	*x = ProductCategoryListReq{}
	if protoimpl.UnsafeEnabled {
		mi := &file_portal_home_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ProductCategoryListReq) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ProductCategoryListReq) ProtoMessage() {}

func (x *ProductCategoryListReq) ProtoReflect() protoreflect.Message {
	mi := &file_portal_home_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ProductCategoryListReq.ProtoReflect.Descriptor instead.
func (*ProductCategoryListReq) Descriptor() ([]byte, []int) {
	return file_portal_home_proto_rawDescGZIP(), []int{3}
}

func (x *ProductCategoryListReq) GetParentId() uint64 {
	if x != nil {
		return x.ParentId
	}
	return 0
}

type ProductCategoryListRsp struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Code    uint32                 `protobuf:"varint,1,opt,name=code,proto3" json:"code,omitempty"`      // 状态码
	Message string                 `protobuf:"bytes,2,opt,name=message,proto3" json:"message,omitempty"` // 提示信息
	Data    []*ProductCategoryItem `protobuf:"bytes,3,rep,name=data,proto3" json:"data,omitempty"`       // 商品分类
}

func (x *ProductCategoryListRsp) Reset() {
	*x = ProductCategoryListRsp{}
	if protoimpl.UnsafeEnabled {
		mi := &file_portal_home_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ProductCategoryListRsp) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ProductCategoryListRsp) ProtoMessage() {}

func (x *ProductCategoryListRsp) ProtoReflect() protoreflect.Message {
	mi := &file_portal_home_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ProductCategoryListRsp.ProtoReflect.Descriptor instead.
func (*ProductCategoryListRsp) Descriptor() ([]byte, []int) {
	return file_portal_home_proto_rawDescGZIP(), []int{4}
}

func (x *ProductCategoryListRsp) GetCode() uint32 {
	if x != nil {
		return x.Code
	}
	return 0
}

func (x *ProductCategoryListRsp) GetMessage() string {
	if x != nil {
		return x.Message
	}
	return ""
}

func (x *ProductCategoryListRsp) GetData() []*ProductCategoryItem {
	if x != nil {
		return x.Data
	}
	return nil
}

type HomeContentRsp_HomeAdvertise struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id   uint64 `protobuf:"varint,1,opt,name=id,proto3" json:"id,omitempty"`     // 主键
	Name string `protobuf:"bytes,2,opt,name=name,proto3" json:"name,omitempty"`  // 名称
	Pic  string `protobuf:"bytes,3,opt,name=pic,proto3" json:"pic,omitempty"`    // 图片地址
	Url  string `protobuf:"bytes,4,opt,name=url,proto3" json:"url,omitempty"`    // 链接地址
	Sort uint32 `protobuf:"varint,5,opt,name=sort,proto3" json:"sort,omitempty"` // 排序
	Note string `protobuf:"bytes,6,opt,name=note,proto3" json:"note,omitempty"`  // 备注
}

func (x *HomeContentRsp_HomeAdvertise) Reset() {
	*x = HomeContentRsp_HomeAdvertise{}
	if protoimpl.UnsafeEnabled {
		mi := &file_portal_home_proto_msgTypes[5]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *HomeContentRsp_HomeAdvertise) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*HomeContentRsp_HomeAdvertise) ProtoMessage() {}

func (x *HomeContentRsp_HomeAdvertise) ProtoReflect() protoreflect.Message {
	mi := &file_portal_home_proto_msgTypes[5]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use HomeContentRsp_HomeAdvertise.ProtoReflect.Descriptor instead.
func (*HomeContentRsp_HomeAdvertise) Descriptor() ([]byte, []int) {
	return file_portal_home_proto_rawDescGZIP(), []int{1, 0}
}

func (x *HomeContentRsp_HomeAdvertise) GetId() uint64 {
	if x != nil {
		return x.Id
	}
	return 0
}

func (x *HomeContentRsp_HomeAdvertise) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *HomeContentRsp_HomeAdvertise) GetPic() string {
	if x != nil {
		return x.Pic
	}
	return ""
}

func (x *HomeContentRsp_HomeAdvertise) GetUrl() string {
	if x != nil {
		return x.Url
	}
	return ""
}

func (x *HomeContentRsp_HomeAdvertise) GetSort() uint32 {
	if x != nil {
		return x.Sort
	}
	return 0
}

func (x *HomeContentRsp_HomeAdvertise) GetNote() string {
	if x != nil {
		return x.Note
	}
	return ""
}

var File_portal_home_proto protoreflect.FileDescriptor

var file_portal_home_proto_rawDesc = []byte{
	0x0a, 0x11, 0x70, 0x6f, 0x72, 0x74, 0x61, 0x6c, 0x2f, 0x68, 0x6f, 0x6d, 0x65, 0x2e, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x12, 0x05, 0x61, 0x64, 0x6d, 0x69, 0x6e, 0x1a, 0x13, 0x70, 0x6f, 0x72, 0x74,
	0x61, 0x6c, 0x2f, 0x70, 0x6f, 0x72, 0x74, 0x61, 0x6c, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a,
	0x17, 0x76, 0x61, 0x6c, 0x69, 0x64, 0x61, 0x74, 0x65, 0x2f, 0x76, 0x61, 0x6c, 0x69, 0x64, 0x61,
	0x74, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x1c, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65,
	0x2f, 0x61, 0x70, 0x69, 0x2f, 0x61, 0x6e, 0x6e, 0x6f, 0x74, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x73,
	0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x1e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2f, 0x77, 0x72, 0x61, 0x70, 0x70, 0x65, 0x72, 0x73,
	0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0x10, 0x0a, 0x0e, 0x48, 0x6f, 0x6d, 0x65, 0x43, 0x6f,
	0x6e, 0x74, 0x65, 0x6e, 0x74, 0x52, 0x65, 0x71, 0x22, 0xd6, 0x01, 0x0a, 0x0e, 0x48, 0x6f, 0x6d,
	0x65, 0x43, 0x6f, 0x6e, 0x74, 0x65, 0x6e, 0x74, 0x52, 0x73, 0x70, 0x12, 0x43, 0x0a, 0x0a, 0x61,
	0x64, 0x76, 0x65, 0x72, 0x74, 0x69, 0x73, 0x65, 0x73, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32,
	0x23, 0x2e, 0x61, 0x64, 0x6d, 0x69, 0x6e, 0x2e, 0x48, 0x6f, 0x6d, 0x65, 0x43, 0x6f, 0x6e, 0x74,
	0x65, 0x6e, 0x74, 0x52, 0x73, 0x70, 0x2e, 0x48, 0x6f, 0x6d, 0x65, 0x41, 0x64, 0x76, 0x65, 0x72,
	0x74, 0x69, 0x73, 0x65, 0x52, 0x0a, 0x61, 0x64, 0x76, 0x65, 0x72, 0x74, 0x69, 0x73, 0x65, 0x73,
	0x1a, 0x7f, 0x0a, 0x0d, 0x48, 0x6f, 0x6d, 0x65, 0x41, 0x64, 0x76, 0x65, 0x72, 0x74, 0x69, 0x73,
	0x65, 0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x04, 0x52, 0x02, 0x69,
	0x64, 0x12, 0x12, 0x0a, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x04, 0x6e, 0x61, 0x6d, 0x65, 0x12, 0x10, 0x0a, 0x03, 0x70, 0x69, 0x63, 0x18, 0x03, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x03, 0x70, 0x69, 0x63, 0x12, 0x10, 0x0a, 0x03, 0x75, 0x72, 0x6c, 0x18, 0x04,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x03, 0x75, 0x72, 0x6c, 0x12, 0x12, 0x0a, 0x04, 0x73, 0x6f, 0x72,
	0x74, 0x18, 0x05, 0x20, 0x01, 0x28, 0x0d, 0x52, 0x04, 0x73, 0x6f, 0x72, 0x74, 0x12, 0x12, 0x0a,
	0x04, 0x6e, 0x6f, 0x74, 0x65, 0x18, 0x06, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x6e, 0x6f, 0x74,
	0x65, 0x22, 0x6a, 0x0a, 0x13, 0x50, 0x72, 0x6f, 0x64, 0x75, 0x63, 0x74, 0x43, 0x61, 0x74, 0x65,
	0x67, 0x6f, 0x72, 0x79, 0x49, 0x74, 0x65, 0x6d, 0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x01,
	0x20, 0x01, 0x28, 0x04, 0x52, 0x02, 0x69, 0x64, 0x12, 0x1b, 0x0a, 0x09, 0x70, 0x61, 0x72, 0x65,
	0x6e, 0x74, 0x5f, 0x69, 0x64, 0x18, 0x02, 0x20, 0x01, 0x28, 0x04, 0x52, 0x08, 0x70, 0x61, 0x72,
	0x65, 0x6e, 0x74, 0x49, 0x64, 0x12, 0x12, 0x0a, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x03, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x12, 0x12, 0x0a, 0x04, 0x69, 0x63, 0x6f,
	0x6e, 0x18, 0x0a, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x69, 0x63, 0x6f, 0x6e, 0x22, 0x3d, 0x0a,
	0x16, 0x50, 0x72, 0x6f, 0x64, 0x75, 0x63, 0x74, 0x43, 0x61, 0x74, 0x65, 0x67, 0x6f, 0x72, 0x79,
	0x4c, 0x69, 0x73, 0x74, 0x52, 0x65, 0x71, 0x12, 0x23, 0x0a, 0x08, 0x70, 0x61, 0x72, 0x65, 0x6e,
	0x74, 0x49, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x04, 0x42, 0x07, 0xfa, 0x42, 0x04, 0x32, 0x02,
	0x28, 0x00, 0x52, 0x08, 0x70, 0x61, 0x72, 0x65, 0x6e, 0x74, 0x49, 0x64, 0x22, 0x76, 0x0a, 0x16,
	0x50, 0x72, 0x6f, 0x64, 0x75, 0x63, 0x74, 0x43, 0x61, 0x74, 0x65, 0x67, 0x6f, 0x72, 0x79, 0x4c,
	0x69, 0x73, 0x74, 0x52, 0x73, 0x70, 0x12, 0x12, 0x0a, 0x04, 0x63, 0x6f, 0x64, 0x65, 0x18, 0x01,
	0x20, 0x01, 0x28, 0x0d, 0x52, 0x04, 0x63, 0x6f, 0x64, 0x65, 0x12, 0x18, 0x0a, 0x07, 0x6d, 0x65,
	0x73, 0x73, 0x61, 0x67, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x6d, 0x65, 0x73,
	0x73, 0x61, 0x67, 0x65, 0x12, 0x2e, 0x0a, 0x04, 0x64, 0x61, 0x74, 0x61, 0x18, 0x03, 0x20, 0x03,
	0x28, 0x0b, 0x32, 0x1a, 0x2e, 0x61, 0x64, 0x6d, 0x69, 0x6e, 0x2e, 0x50, 0x72, 0x6f, 0x64, 0x75,
	0x63, 0x74, 0x43, 0x61, 0x74, 0x65, 0x67, 0x6f, 0x72, 0x79, 0x49, 0x74, 0x65, 0x6d, 0x52, 0x04,
	0x64, 0x61, 0x74, 0x61, 0x32, 0xe2, 0x01, 0x0a, 0x0d, 0x50, 0x6f, 0x72, 0x74, 0x61, 0x6c, 0x48,
	0x6f, 0x6d, 0x65, 0x41, 0x70, 0x69, 0x12, 0x52, 0x0a, 0x0b, 0x48, 0x6f, 0x6d, 0x65, 0x43, 0x6f,
	0x6e, 0x74, 0x65, 0x6e, 0x74, 0x12, 0x15, 0x2e, 0x61, 0x64, 0x6d, 0x69, 0x6e, 0x2e, 0x48, 0x6f,
	0x6d, 0x65, 0x43, 0x6f, 0x6e, 0x74, 0x65, 0x6e, 0x74, 0x52, 0x65, 0x71, 0x1a, 0x15, 0x2e, 0x61,
	0x64, 0x6d, 0x69, 0x6e, 0x2e, 0x48, 0x6f, 0x6d, 0x65, 0x43, 0x6f, 0x6e, 0x74, 0x65, 0x6e, 0x74,
	0x52, 0x73, 0x70, 0x22, 0x15, 0x82, 0xd3, 0xe4, 0x93, 0x02, 0x0f, 0x12, 0x0d, 0x2f, 0x68, 0x6f,
	0x6d, 0x65, 0x2f, 0x63, 0x6f, 0x6e, 0x74, 0x65, 0x6e, 0x74, 0x12, 0x7d, 0x0a, 0x13, 0x50, 0x72,
	0x6f, 0x64, 0x75, 0x63, 0x74, 0x43, 0x61, 0x74, 0x65, 0x67, 0x6f, 0x72, 0x79, 0x4c, 0x69, 0x73,
	0x74, 0x12, 0x1d, 0x2e, 0x61, 0x64, 0x6d, 0x69, 0x6e, 0x2e, 0x50, 0x72, 0x6f, 0x64, 0x75, 0x63,
	0x74, 0x43, 0x61, 0x74, 0x65, 0x67, 0x6f, 0x72, 0x79, 0x4c, 0x69, 0x73, 0x74, 0x52, 0x65, 0x71,
	0x1a, 0x1d, 0x2e, 0x61, 0x64, 0x6d, 0x69, 0x6e, 0x2e, 0x50, 0x72, 0x6f, 0x64, 0x75, 0x63, 0x74,
	0x43, 0x61, 0x74, 0x65, 0x67, 0x6f, 0x72, 0x79, 0x4c, 0x69, 0x73, 0x74, 0x52, 0x73, 0x70, 0x22,
	0x28, 0x82, 0xd3, 0xe4, 0x93, 0x02, 0x22, 0x12, 0x20, 0x2f, 0x68, 0x6f, 0x6d, 0x65, 0x2f, 0x70,
	0x72, 0x6f, 0x64, 0x75, 0x63, 0x74, 0x43, 0x61, 0x74, 0x65, 0x4c, 0x69, 0x73, 0x74, 0x2f, 0x7b,
	0x70, 0x61, 0x72, 0x65, 0x6e, 0x74, 0x49, 0x64, 0x7d, 0x42, 0x08, 0x5a, 0x06, 0x2e, 0x3b, 0x6d,
	0x61, 0x6c, 0x6c, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_portal_home_proto_rawDescOnce sync.Once
	file_portal_home_proto_rawDescData = file_portal_home_proto_rawDesc
)

func file_portal_home_proto_rawDescGZIP() []byte {
	file_portal_home_proto_rawDescOnce.Do(func() {
		file_portal_home_proto_rawDescData = protoimpl.X.CompressGZIP(file_portal_home_proto_rawDescData)
	})
	return file_portal_home_proto_rawDescData
}

var file_portal_home_proto_msgTypes = make([]protoimpl.MessageInfo, 6)
var file_portal_home_proto_goTypes = []interface{}{
	(*HomeContentReq)(nil),               // 0: admin.HomeContentReq
	(*HomeContentRsp)(nil),               // 1: admin.HomeContentRsp
	(*ProductCategoryItem)(nil),          // 2: admin.ProductCategoryItem
	(*ProductCategoryListReq)(nil),       // 3: admin.ProductCategoryListReq
	(*ProductCategoryListRsp)(nil),       // 4: admin.ProductCategoryListRsp
	(*HomeContentRsp_HomeAdvertise)(nil), // 5: admin.HomeContentRsp.HomeAdvertise
}
var file_portal_home_proto_depIdxs = []int32{
	5, // 0: admin.HomeContentRsp.advertises:type_name -> admin.HomeContentRsp.HomeAdvertise
	2, // 1: admin.ProductCategoryListRsp.data:type_name -> admin.ProductCategoryItem
	0, // 2: admin.PortalHomeApi.HomeContent:input_type -> admin.HomeContentReq
	3, // 3: admin.PortalHomeApi.ProductCategoryList:input_type -> admin.ProductCategoryListReq
	1, // 4: admin.PortalHomeApi.HomeContent:output_type -> admin.HomeContentRsp
	4, // 5: admin.PortalHomeApi.ProductCategoryList:output_type -> admin.ProductCategoryListRsp
	4, // [4:6] is the sub-list for method output_type
	2, // [2:4] is the sub-list for method input_type
	2, // [2:2] is the sub-list for extension type_name
	2, // [2:2] is the sub-list for extension extendee
	0, // [0:2] is the sub-list for field type_name
}

func init() { file_portal_home_proto_init() }
func file_portal_home_proto_init() {
	if File_portal_home_proto != nil {
		return
	}
	file_portal_portal_proto_init()
	if !protoimpl.UnsafeEnabled {
		file_portal_home_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*HomeContentReq); i {
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
		file_portal_home_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*HomeContentRsp); i {
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
		file_portal_home_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ProductCategoryItem); i {
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
		file_portal_home_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ProductCategoryListReq); i {
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
		file_portal_home_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ProductCategoryListRsp); i {
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
		file_portal_home_proto_msgTypes[5].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*HomeContentRsp_HomeAdvertise); i {
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
			RawDescriptor: file_portal_home_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   6,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_portal_home_proto_goTypes,
		DependencyIndexes: file_portal_home_proto_depIdxs,
		MessageInfos:      file_portal_home_proto_msgTypes,
	}.Build()
	File_portal_home_proto = out.File
	file_portal_home_proto_rawDesc = nil
	file_portal_home_proto_goTypes = nil
	file_portal_home_proto_depIdxs = nil
}
