// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.30.0
// 	protoc        v3.21.12
// source: admin/oms_admin.proto

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

// 添加或更新退货原因参数
type AddOrUpdateOrderReturnReasonParam struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id         uint64 `protobuf:"varint,1,opt,name=id,proto3" json:"id,omitempty"`                                   // 主键
	Name       string `protobuf:"bytes,2,opt,name=name,proto3" json:"name,omitempty"`                                // 退货类型
	Sort       uint32 `protobuf:"varint,3,opt,name=sort,proto3" json:"sort,omitempty"`                               // 排序
	Status     uint32 `protobuf:"varint,4,opt,name=status,proto3" json:"status,omitempty"`                           // 状态：0->不启用；1->启用
	CreateTime uint32 `protobuf:"varint,5,opt,name=create_time,json=createTime,proto3" json:"create_time,omitempty"` // 添加时间
}

func (x *AddOrUpdateOrderReturnReasonParam) Reset() {
	*x = AddOrUpdateOrderReturnReasonParam{}
	if protoimpl.UnsafeEnabled {
		mi := &file_admin_oms_admin_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *AddOrUpdateOrderReturnReasonParam) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*AddOrUpdateOrderReturnReasonParam) ProtoMessage() {}

func (x *AddOrUpdateOrderReturnReasonParam) ProtoReflect() protoreflect.Message {
	mi := &file_admin_oms_admin_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use AddOrUpdateOrderReturnReasonParam.ProtoReflect.Descriptor instead.
func (*AddOrUpdateOrderReturnReasonParam) Descriptor() ([]byte, []int) {
	return file_admin_oms_admin_proto_rawDescGZIP(), []int{0}
}

func (x *AddOrUpdateOrderReturnReasonParam) GetId() uint64 {
	if x != nil {
		return x.Id
	}
	return 0
}

func (x *AddOrUpdateOrderReturnReasonParam) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *AddOrUpdateOrderReturnReasonParam) GetSort() uint32 {
	if x != nil {
		return x.Sort
	}
	return 0
}

func (x *AddOrUpdateOrderReturnReasonParam) GetStatus() uint32 {
	if x != nil {
		return x.Status
	}
	return 0
}

func (x *AddOrUpdateOrderReturnReasonParam) GetCreateTime() uint32 {
	if x != nil {
		return x.CreateTime
	}
	return 0
}

// 分页查询退货原因
type GetOrderReturnReasonsParam struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id       *wrapperspb.UInt64Value `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`                               // 主键
	PageNum  uint32                  `protobuf:"varint,10,opt,name=page_num,json=pageNum,proto3" json:"page_num,omitempty"`    // 页面大小
	PageSize uint32                  `protobuf:"varint,11,opt,name=page_size,json=pageSize,proto3" json:"page_size,omitempty"` // 页码
}

func (x *GetOrderReturnReasonsParam) Reset() {
	*x = GetOrderReturnReasonsParam{}
	if protoimpl.UnsafeEnabled {
		mi := &file_admin_oms_admin_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetOrderReturnReasonsParam) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetOrderReturnReasonsParam) ProtoMessage() {}

func (x *GetOrderReturnReasonsParam) ProtoReflect() protoreflect.Message {
	mi := &file_admin_oms_admin_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetOrderReturnReasonsParam.ProtoReflect.Descriptor instead.
func (*GetOrderReturnReasonsParam) Descriptor() ([]byte, []int) {
	return file_admin_oms_admin_proto_rawDescGZIP(), []int{1}
}

func (x *GetOrderReturnReasonsParam) GetId() *wrapperspb.UInt64Value {
	if x != nil {
		return x.Id
	}
	return nil
}

func (x *GetOrderReturnReasonsParam) GetPageNum() uint32 {
	if x != nil {
		return x.PageNum
	}
	return 0
}

func (x *GetOrderReturnReasonsParam) GetPageSize() uint32 {
	if x != nil {
		return x.PageSize
	}
	return 0
}

// 分页查询退货原因
type OrderReturnReasonsData struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Data      []*OrderReturnReason `protobuf:"bytes,1,rep,name=data,proto3" json:"data,omitempty"`                             // 退货原因
	PageTotal uint32               `protobuf:"varint,2,opt,name=page_total,json=pageTotal,proto3" json:"page_total,omitempty"` // 数据总数
	PageSize  uint32               `protobuf:"varint,3,opt,name=page_size,json=pageSize,proto3" json:"page_size,omitempty"`    // 页码
	PageNum   uint32               `protobuf:"varint,4,opt,name=page_num,json=pageNum,proto3" json:"page_num,omitempty"`       // 页面大小
}

func (x *OrderReturnReasonsData) Reset() {
	*x = OrderReturnReasonsData{}
	if protoimpl.UnsafeEnabled {
		mi := &file_admin_oms_admin_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *OrderReturnReasonsData) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*OrderReturnReasonsData) ProtoMessage() {}

func (x *OrderReturnReasonsData) ProtoReflect() protoreflect.Message {
	mi := &file_admin_oms_admin_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use OrderReturnReasonsData.ProtoReflect.Descriptor instead.
func (*OrderReturnReasonsData) Descriptor() ([]byte, []int) {
	return file_admin_oms_admin_proto_rawDescGZIP(), []int{2}
}

func (x *OrderReturnReasonsData) GetData() []*OrderReturnReason {
	if x != nil {
		return x.Data
	}
	return nil
}

func (x *OrderReturnReasonsData) GetPageTotal() uint32 {
	if x != nil {
		return x.PageTotal
	}
	return 0
}

func (x *OrderReturnReasonsData) GetPageSize() uint32 {
	if x != nil {
		return x.PageSize
	}
	return 0
}

func (x *OrderReturnReasonsData) GetPageNum() uint32 {
	if x != nil {
		return x.PageNum
	}
	return 0
}

type GetOrderReturnReasonsRsp struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Code    uint32                  `protobuf:"varint,1,opt,name=code,proto3" json:"code,omitempty"`      // 状态码
	Message string                  `protobuf:"bytes,2,opt,name=message,proto3" json:"message,omitempty"` // 提示信息
	Data    *OrderReturnReasonsData `protobuf:"bytes,3,opt,name=data,proto3" json:"data,omitempty"`       //
}

func (x *GetOrderReturnReasonsRsp) Reset() {
	*x = GetOrderReturnReasonsRsp{}
	if protoimpl.UnsafeEnabled {
		mi := &file_admin_oms_admin_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetOrderReturnReasonsRsp) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetOrderReturnReasonsRsp) ProtoMessage() {}

func (x *GetOrderReturnReasonsRsp) ProtoReflect() protoreflect.Message {
	mi := &file_admin_oms_admin_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetOrderReturnReasonsRsp.ProtoReflect.Descriptor instead.
func (*GetOrderReturnReasonsRsp) Descriptor() ([]byte, []int) {
	return file_admin_oms_admin_proto_rawDescGZIP(), []int{3}
}

func (x *GetOrderReturnReasonsRsp) GetCode() uint32 {
	if x != nil {
		return x.Code
	}
	return 0
}

func (x *GetOrderReturnReasonsRsp) GetMessage() string {
	if x != nil {
		return x.Message
	}
	return ""
}

func (x *GetOrderReturnReasonsRsp) GetData() *OrderReturnReasonsData {
	if x != nil {
		return x.Data
	}
	return nil
}

// 根据id获取退货原因
type GetOrderReturnReasonReq struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id uint64 `protobuf:"varint,1,opt,name=id,proto3" json:"id,omitempty"`
}

func (x *GetOrderReturnReasonReq) Reset() {
	*x = GetOrderReturnReasonReq{}
	if protoimpl.UnsafeEnabled {
		mi := &file_admin_oms_admin_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetOrderReturnReasonReq) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetOrderReturnReasonReq) ProtoMessage() {}

func (x *GetOrderReturnReasonReq) ProtoReflect() protoreflect.Message {
	mi := &file_admin_oms_admin_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetOrderReturnReasonReq.ProtoReflect.Descriptor instead.
func (*GetOrderReturnReasonReq) Descriptor() ([]byte, []int) {
	return file_admin_oms_admin_proto_rawDescGZIP(), []int{4}
}

func (x *GetOrderReturnReasonReq) GetId() uint64 {
	if x != nil {
		return x.Id
	}
	return 0
}

// 根据id获取退货原因
type GetOrderReturnReasonRsp struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Code    uint32             `protobuf:"varint,1,opt,name=code,proto3" json:"code,omitempty"`      // 状态码
	Message string             `protobuf:"bytes,2,opt,name=message,proto3" json:"message,omitempty"` // 提示信息
	Data    *OrderReturnReason `protobuf:"bytes,3,opt,name=data,proto3" json:"data,omitempty"`       // 数据
}

func (x *GetOrderReturnReasonRsp) Reset() {
	*x = GetOrderReturnReasonRsp{}
	if protoimpl.UnsafeEnabled {
		mi := &file_admin_oms_admin_proto_msgTypes[5]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetOrderReturnReasonRsp) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetOrderReturnReasonRsp) ProtoMessage() {}

func (x *GetOrderReturnReasonRsp) ProtoReflect() protoreflect.Message {
	mi := &file_admin_oms_admin_proto_msgTypes[5]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetOrderReturnReasonRsp.ProtoReflect.Descriptor instead.
func (*GetOrderReturnReasonRsp) Descriptor() ([]byte, []int) {
	return file_admin_oms_admin_proto_rawDescGZIP(), []int{5}
}

func (x *GetOrderReturnReasonRsp) GetCode() uint32 {
	if x != nil {
		return x.Code
	}
	return 0
}

func (x *GetOrderReturnReasonRsp) GetMessage() string {
	if x != nil {
		return x.Message
	}
	return ""
}

func (x *GetOrderReturnReasonRsp) GetData() *OrderReturnReason {
	if x != nil {
		return x.Data
	}
	return nil
}

// 删除退货原因
type DeleteOrderReturnReasonReq struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id uint64 `protobuf:"varint,1,opt,name=id,proto3" json:"id,omitempty"`
}

func (x *DeleteOrderReturnReasonReq) Reset() {
	*x = DeleteOrderReturnReasonReq{}
	if protoimpl.UnsafeEnabled {
		mi := &file_admin_oms_admin_proto_msgTypes[6]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *DeleteOrderReturnReasonReq) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*DeleteOrderReturnReasonReq) ProtoMessage() {}

func (x *DeleteOrderReturnReasonReq) ProtoReflect() protoreflect.Message {
	mi := &file_admin_oms_admin_proto_msgTypes[6]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use DeleteOrderReturnReasonReq.ProtoReflect.Descriptor instead.
func (*DeleteOrderReturnReasonReq) Descriptor() ([]byte, []int) {
	return file_admin_oms_admin_proto_rawDescGZIP(), []int{6}
}

func (x *DeleteOrderReturnReasonReq) GetId() uint64 {
	if x != nil {
		return x.Id
	}
	return 0
}

var File_admin_oms_admin_proto protoreflect.FileDescriptor

var file_admin_oms_admin_proto_rawDesc = []byte{
	0x0a, 0x15, 0x61, 0x64, 0x6d, 0x69, 0x6e, 0x2f, 0x6f, 0x6d, 0x73, 0x5f, 0x61, 0x64, 0x6d, 0x69,
	0x6e, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x05, 0x61, 0x64, 0x6d, 0x69, 0x6e, 0x1a, 0x17,
	0x76, 0x61, 0x6c, 0x69, 0x64, 0x61, 0x74, 0x65, 0x2f, 0x76, 0x61, 0x6c, 0x69, 0x64, 0x61, 0x74,
	0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x1c, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f,
	0x61, 0x70, 0x69, 0x2f, 0x61, 0x6e, 0x6e, 0x6f, 0x74, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x2e,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x1e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2f, 0x77, 0x72, 0x61, 0x70, 0x70, 0x65, 0x72, 0x73, 0x2e,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x11, 0x61, 0x64, 0x6d, 0x69, 0x6e, 0x2f, 0x61, 0x64, 0x6d,
	0x69, 0x6e, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x23, 0x6d, 0x6f, 0x64, 0x65, 0x6c, 0x2f,
	0x6f, 0x6d, 0x73, 0x5f, 0x6f, 0x72, 0x64, 0x65, 0x72, 0x5f, 0x72, 0x65, 0x74, 0x75, 0x72, 0x6e,
	0x5f, 0x72, 0x65, 0x61, 0x73, 0x6f, 0x6e, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0xa6, 0x01,
	0x0a, 0x21, 0x41, 0x64, 0x64, 0x4f, 0x72, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x4f, 0x72, 0x64,
	0x65, 0x72, 0x52, 0x65, 0x74, 0x75, 0x72, 0x6e, 0x52, 0x65, 0x61, 0x73, 0x6f, 0x6e, 0x50, 0x61,
	0x72, 0x61, 0x6d, 0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x04, 0x52,
	0x02, 0x69, 0x64, 0x12, 0x1b, 0x0a, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28,
	0x09, 0x42, 0x07, 0xfa, 0x42, 0x04, 0x72, 0x02, 0x10, 0x01, 0x52, 0x04, 0x6e, 0x61, 0x6d, 0x65,
	0x12, 0x1b, 0x0a, 0x04, 0x73, 0x6f, 0x72, 0x74, 0x18, 0x03, 0x20, 0x01, 0x28, 0x0d, 0x42, 0x07,
	0xfa, 0x42, 0x04, 0x2a, 0x02, 0x28, 0x00, 0x52, 0x04, 0x73, 0x6f, 0x72, 0x74, 0x12, 0x16, 0x0a,
	0x06, 0x73, 0x74, 0x61, 0x74, 0x75, 0x73, 0x18, 0x04, 0x20, 0x01, 0x28, 0x0d, 0x52, 0x06, 0x73,
	0x74, 0x61, 0x74, 0x75, 0x73, 0x12, 0x1f, 0x0a, 0x0b, 0x63, 0x72, 0x65, 0x61, 0x74, 0x65, 0x5f,
	0x74, 0x69, 0x6d, 0x65, 0x18, 0x05, 0x20, 0x01, 0x28, 0x0d, 0x52, 0x0a, 0x63, 0x72, 0x65, 0x61,
	0x74, 0x65, 0x54, 0x69, 0x6d, 0x65, 0x22, 0x94, 0x01, 0x0a, 0x1a, 0x47, 0x65, 0x74, 0x4f, 0x72,
	0x64, 0x65, 0x72, 0x52, 0x65, 0x74, 0x75, 0x72, 0x6e, 0x52, 0x65, 0x61, 0x73, 0x6f, 0x6e, 0x73,
	0x50, 0x61, 0x72, 0x61, 0x6d, 0x12, 0x2c, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28,
	0x0b, 0x32, 0x1c, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x62, 0x75, 0x66, 0x2e, 0x55, 0x49, 0x6e, 0x74, 0x36, 0x34, 0x56, 0x61, 0x6c, 0x75, 0x65, 0x52,
	0x02, 0x69, 0x64, 0x12, 0x22, 0x0a, 0x08, 0x70, 0x61, 0x67, 0x65, 0x5f, 0x6e, 0x75, 0x6d, 0x18,
	0x0a, 0x20, 0x01, 0x28, 0x0d, 0x42, 0x07, 0xfa, 0x42, 0x04, 0x2a, 0x02, 0x28, 0x00, 0x52, 0x07,
	0x70, 0x61, 0x67, 0x65, 0x4e, 0x75, 0x6d, 0x12, 0x24, 0x0a, 0x09, 0x70, 0x61, 0x67, 0x65, 0x5f,
	0x73, 0x69, 0x7a, 0x65, 0x18, 0x0b, 0x20, 0x01, 0x28, 0x0d, 0x42, 0x07, 0xfa, 0x42, 0x04, 0x2a,
	0x02, 0x28, 0x00, 0x52, 0x08, 0x70, 0x61, 0x67, 0x65, 0x53, 0x69, 0x7a, 0x65, 0x22, 0x9d, 0x01,
	0x0a, 0x16, 0x4f, 0x72, 0x64, 0x65, 0x72, 0x52, 0x65, 0x74, 0x75, 0x72, 0x6e, 0x52, 0x65, 0x61,
	0x73, 0x6f, 0x6e, 0x73, 0x44, 0x61, 0x74, 0x61, 0x12, 0x2c, 0x0a, 0x04, 0x64, 0x61, 0x74, 0x61,
	0x18, 0x01, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x18, 0x2e, 0x6d, 0x6f, 0x64, 0x65, 0x6c, 0x2e, 0x4f,
	0x72, 0x64, 0x65, 0x72, 0x52, 0x65, 0x74, 0x75, 0x72, 0x6e, 0x52, 0x65, 0x61, 0x73, 0x6f, 0x6e,
	0x52, 0x04, 0x64, 0x61, 0x74, 0x61, 0x12, 0x1d, 0x0a, 0x0a, 0x70, 0x61, 0x67, 0x65, 0x5f, 0x74,
	0x6f, 0x74, 0x61, 0x6c, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0d, 0x52, 0x09, 0x70, 0x61, 0x67, 0x65,
	0x54, 0x6f, 0x74, 0x61, 0x6c, 0x12, 0x1b, 0x0a, 0x09, 0x70, 0x61, 0x67, 0x65, 0x5f, 0x73, 0x69,
	0x7a, 0x65, 0x18, 0x03, 0x20, 0x01, 0x28, 0x0d, 0x52, 0x08, 0x70, 0x61, 0x67, 0x65, 0x53, 0x69,
	0x7a, 0x65, 0x12, 0x19, 0x0a, 0x08, 0x70, 0x61, 0x67, 0x65, 0x5f, 0x6e, 0x75, 0x6d, 0x18, 0x04,
	0x20, 0x01, 0x28, 0x0d, 0x52, 0x07, 0x70, 0x61, 0x67, 0x65, 0x4e, 0x75, 0x6d, 0x22, 0x7b, 0x0a,
	0x18, 0x47, 0x65, 0x74, 0x4f, 0x72, 0x64, 0x65, 0x72, 0x52, 0x65, 0x74, 0x75, 0x72, 0x6e, 0x52,
	0x65, 0x61, 0x73, 0x6f, 0x6e, 0x73, 0x52, 0x73, 0x70, 0x12, 0x12, 0x0a, 0x04, 0x63, 0x6f, 0x64,
	0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0d, 0x52, 0x04, 0x63, 0x6f, 0x64, 0x65, 0x12, 0x18, 0x0a,
	0x07, 0x6d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07,
	0x6d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x12, 0x31, 0x0a, 0x04, 0x64, 0x61, 0x74, 0x61, 0x18,
	0x03, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1d, 0x2e, 0x61, 0x64, 0x6d, 0x69, 0x6e, 0x2e, 0x4f, 0x72,
	0x64, 0x65, 0x72, 0x52, 0x65, 0x74, 0x75, 0x72, 0x6e, 0x52, 0x65, 0x61, 0x73, 0x6f, 0x6e, 0x73,
	0x44, 0x61, 0x74, 0x61, 0x52, 0x04, 0x64, 0x61, 0x74, 0x61, 0x22, 0x29, 0x0a, 0x17, 0x47, 0x65,
	0x74, 0x4f, 0x72, 0x64, 0x65, 0x72, 0x52, 0x65, 0x74, 0x75, 0x72, 0x6e, 0x52, 0x65, 0x61, 0x73,
	0x6f, 0x6e, 0x52, 0x65, 0x71, 0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28,
	0x04, 0x52, 0x02, 0x69, 0x64, 0x22, 0x75, 0x0a, 0x17, 0x47, 0x65, 0x74, 0x4f, 0x72, 0x64, 0x65,
	0x72, 0x52, 0x65, 0x74, 0x75, 0x72, 0x6e, 0x52, 0x65, 0x61, 0x73, 0x6f, 0x6e, 0x52, 0x73, 0x70,
	0x12, 0x12, 0x0a, 0x04, 0x63, 0x6f, 0x64, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0d, 0x52, 0x04,
	0x63, 0x6f, 0x64, 0x65, 0x12, 0x18, 0x0a, 0x07, 0x6d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x18,
	0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x6d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x12, 0x2c,
	0x0a, 0x04, 0x64, 0x61, 0x74, 0x61, 0x18, 0x03, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x18, 0x2e, 0x6d,
	0x6f, 0x64, 0x65, 0x6c, 0x2e, 0x4f, 0x72, 0x64, 0x65, 0x72, 0x52, 0x65, 0x74, 0x75, 0x72, 0x6e,
	0x52, 0x65, 0x61, 0x73, 0x6f, 0x6e, 0x52, 0x04, 0x64, 0x61, 0x74, 0x61, 0x22, 0x2c, 0x0a, 0x1a,
	0x44, 0x65, 0x6c, 0x65, 0x74, 0x65, 0x4f, 0x72, 0x64, 0x65, 0x72, 0x52, 0x65, 0x74, 0x75, 0x72,
	0x6e, 0x52, 0x65, 0x61, 0x73, 0x6f, 0x6e, 0x52, 0x65, 0x71, 0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64,
	0x18, 0x01, 0x20, 0x01, 0x28, 0x04, 0x52, 0x02, 0x69, 0x64, 0x32, 0xe6, 0x04, 0x0a, 0x0b, 0x4f,
	0x6d, 0x73, 0x41, 0x64, 0x6d, 0x69, 0x6e, 0x41, 0x70, 0x69, 0x12, 0x75, 0x0a, 0x17, 0x43, 0x72,
	0x65, 0x61, 0x74, 0x65, 0x4f, 0x72, 0x64, 0x65, 0x72, 0x52, 0x65, 0x74, 0x75, 0x72, 0x6e, 0x52,
	0x65, 0x61, 0x73, 0x6f, 0x6e, 0x12, 0x28, 0x2e, 0x61, 0x64, 0x6d, 0x69, 0x6e, 0x2e, 0x41, 0x64,
	0x64, 0x4f, 0x72, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x4f, 0x72, 0x64, 0x65, 0x72, 0x52, 0x65,
	0x74, 0x75, 0x72, 0x6e, 0x52, 0x65, 0x61, 0x73, 0x6f, 0x6e, 0x50, 0x61, 0x72, 0x61, 0x6d, 0x1a,
	0x10, 0x2e, 0x61, 0x64, 0x6d, 0x69, 0x6e, 0x2e, 0x43, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x52, 0x73,
	0x70, 0x22, 0x1e, 0x82, 0xd3, 0xe4, 0x93, 0x02, 0x18, 0x3a, 0x01, 0x2a, 0x22, 0x13, 0x2f, 0x6f,
	0x72, 0x64, 0x65, 0x72, 0x52, 0x65, 0x74, 0x75, 0x72, 0x6e, 0x52, 0x65, 0x61, 0x73, 0x6f, 0x6e,
	0x73, 0x12, 0x7a, 0x0a, 0x17, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x4f, 0x72, 0x64, 0x65, 0x72,
	0x52, 0x65, 0x74, 0x75, 0x72, 0x6e, 0x52, 0x65, 0x61, 0x73, 0x6f, 0x6e, 0x12, 0x28, 0x2e, 0x61,
	0x64, 0x6d, 0x69, 0x6e, 0x2e, 0x41, 0x64, 0x64, 0x4f, 0x72, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65,
	0x4f, 0x72, 0x64, 0x65, 0x72, 0x52, 0x65, 0x74, 0x75, 0x72, 0x6e, 0x52, 0x65, 0x61, 0x73, 0x6f,
	0x6e, 0x50, 0x61, 0x72, 0x61, 0x6d, 0x1a, 0x10, 0x2e, 0x61, 0x64, 0x6d, 0x69, 0x6e, 0x2e, 0x43,
	0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x52, 0x73, 0x70, 0x22, 0x23, 0x82, 0xd3, 0xe4, 0x93, 0x02, 0x1d,
	0x3a, 0x01, 0x2a, 0x1a, 0x18, 0x2f, 0x6f, 0x72, 0x64, 0x65, 0x72, 0x52, 0x65, 0x74, 0x75, 0x72,
	0x6e, 0x52, 0x65, 0x61, 0x73, 0x6f, 0x6e, 0x73, 0x2f, 0x7b, 0x69, 0x64, 0x7d, 0x12, 0x78, 0x0a,
	0x15, 0x47, 0x65, 0x74, 0x4f, 0x72, 0x64, 0x65, 0x72, 0x52, 0x65, 0x74, 0x75, 0x72, 0x6e, 0x52,
	0x65, 0x61, 0x73, 0x6f, 0x6e, 0x73, 0x12, 0x21, 0x2e, 0x61, 0x64, 0x6d, 0x69, 0x6e, 0x2e, 0x47,
	0x65, 0x74, 0x4f, 0x72, 0x64, 0x65, 0x72, 0x52, 0x65, 0x74, 0x75, 0x72, 0x6e, 0x52, 0x65, 0x61,
	0x73, 0x6f, 0x6e, 0x73, 0x50, 0x61, 0x72, 0x61, 0x6d, 0x1a, 0x1f, 0x2e, 0x61, 0x64, 0x6d, 0x69,
	0x6e, 0x2e, 0x47, 0x65, 0x74, 0x4f, 0x72, 0x64, 0x65, 0x72, 0x52, 0x65, 0x74, 0x75, 0x72, 0x6e,
	0x52, 0x65, 0x61, 0x73, 0x6f, 0x6e, 0x73, 0x52, 0x73, 0x70, 0x22, 0x1b, 0x82, 0xd3, 0xe4, 0x93,
	0x02, 0x15, 0x12, 0x13, 0x2f, 0x6f, 0x72, 0x64, 0x65, 0x72, 0x52, 0x65, 0x74, 0x75, 0x72, 0x6e,
	0x52, 0x65, 0x61, 0x73, 0x6f, 0x6e, 0x73, 0x12, 0x78, 0x0a, 0x14, 0x47, 0x65, 0x74, 0x4f, 0x72,
	0x64, 0x65, 0x72, 0x52, 0x65, 0x74, 0x75, 0x72, 0x6e, 0x52, 0x65, 0x61, 0x73, 0x6f, 0x6e, 0x12,
	0x1e, 0x2e, 0x61, 0x64, 0x6d, 0x69, 0x6e, 0x2e, 0x47, 0x65, 0x74, 0x4f, 0x72, 0x64, 0x65, 0x72,
	0x52, 0x65, 0x74, 0x75, 0x72, 0x6e, 0x52, 0x65, 0x61, 0x73, 0x6f, 0x6e, 0x52, 0x65, 0x71, 0x1a,
	0x1e, 0x2e, 0x61, 0x64, 0x6d, 0x69, 0x6e, 0x2e, 0x47, 0x65, 0x74, 0x4f, 0x72, 0x64, 0x65, 0x72,
	0x52, 0x65, 0x74, 0x75, 0x72, 0x6e, 0x52, 0x65, 0x61, 0x73, 0x6f, 0x6e, 0x52, 0x73, 0x70, 0x22,
	0x20, 0x82, 0xd3, 0xe4, 0x93, 0x02, 0x1a, 0x12, 0x18, 0x2f, 0x6f, 0x72, 0x64, 0x65, 0x72, 0x52,
	0x65, 0x74, 0x75, 0x72, 0x6e, 0x52, 0x65, 0x61, 0x73, 0x6f, 0x6e, 0x73, 0x2f, 0x7b, 0x69, 0x64,
	0x7d, 0x12, 0x70, 0x0a, 0x17, 0x44, 0x65, 0x6c, 0x65, 0x74, 0x65, 0x4f, 0x72, 0x64, 0x65, 0x72,
	0x52, 0x65, 0x74, 0x75, 0x72, 0x6e, 0x52, 0x65, 0x61, 0x73, 0x6f, 0x6e, 0x12, 0x21, 0x2e, 0x61,
	0x64, 0x6d, 0x69, 0x6e, 0x2e, 0x44, 0x65, 0x6c, 0x65, 0x74, 0x65, 0x4f, 0x72, 0x64, 0x65, 0x72,
	0x52, 0x65, 0x74, 0x75, 0x72, 0x6e, 0x52, 0x65, 0x61, 0x73, 0x6f, 0x6e, 0x52, 0x65, 0x71, 0x1a,
	0x10, 0x2e, 0x61, 0x64, 0x6d, 0x69, 0x6e, 0x2e, 0x43, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x52, 0x73,
	0x70, 0x22, 0x20, 0x82, 0xd3, 0xe4, 0x93, 0x02, 0x1a, 0x2a, 0x18, 0x2f, 0x6f, 0x72, 0x64, 0x65,
	0x72, 0x52, 0x65, 0x74, 0x75, 0x72, 0x6e, 0x52, 0x65, 0x61, 0x73, 0x6f, 0x6e, 0x73, 0x2f, 0x7b,
	0x69, 0x64, 0x7d, 0x42, 0x08, 0x5a, 0x06, 0x2e, 0x3b, 0x6d, 0x61, 0x6c, 0x6c, 0x62, 0x06, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_admin_oms_admin_proto_rawDescOnce sync.Once
	file_admin_oms_admin_proto_rawDescData = file_admin_oms_admin_proto_rawDesc
)

func file_admin_oms_admin_proto_rawDescGZIP() []byte {
	file_admin_oms_admin_proto_rawDescOnce.Do(func() {
		file_admin_oms_admin_proto_rawDescData = protoimpl.X.CompressGZIP(file_admin_oms_admin_proto_rawDescData)
	})
	return file_admin_oms_admin_proto_rawDescData
}

var file_admin_oms_admin_proto_msgTypes = make([]protoimpl.MessageInfo, 7)
var file_admin_oms_admin_proto_goTypes = []interface{}{
	(*AddOrUpdateOrderReturnReasonParam)(nil), // 0: admin.AddOrUpdateOrderReturnReasonParam
	(*GetOrderReturnReasonsParam)(nil),        // 1: admin.GetOrderReturnReasonsParam
	(*OrderReturnReasonsData)(nil),            // 2: admin.OrderReturnReasonsData
	(*GetOrderReturnReasonsRsp)(nil),          // 3: admin.GetOrderReturnReasonsRsp
	(*GetOrderReturnReasonReq)(nil),           // 4: admin.GetOrderReturnReasonReq
	(*GetOrderReturnReasonRsp)(nil),           // 5: admin.GetOrderReturnReasonRsp
	(*DeleteOrderReturnReasonReq)(nil),        // 6: admin.DeleteOrderReturnReasonReq
	(*wrapperspb.UInt64Value)(nil),            // 7: google.protobuf.UInt64Value
	(*OrderReturnReason)(nil),                 // 8: model.OrderReturnReason
	(*CommonRsp)(nil),                         // 9: admin.CommonRsp
}
var file_admin_oms_admin_proto_depIdxs = []int32{
	7, // 0: admin.GetOrderReturnReasonsParam.id:type_name -> google.protobuf.UInt64Value
	8, // 1: admin.OrderReturnReasonsData.data:type_name -> model.OrderReturnReason
	2, // 2: admin.GetOrderReturnReasonsRsp.data:type_name -> admin.OrderReturnReasonsData
	8, // 3: admin.GetOrderReturnReasonRsp.data:type_name -> model.OrderReturnReason
	0, // 4: admin.OmsAdminApi.CreateOrderReturnReason:input_type -> admin.AddOrUpdateOrderReturnReasonParam
	0, // 5: admin.OmsAdminApi.UpdateOrderReturnReason:input_type -> admin.AddOrUpdateOrderReturnReasonParam
	1, // 6: admin.OmsAdminApi.GetOrderReturnReasons:input_type -> admin.GetOrderReturnReasonsParam
	4, // 7: admin.OmsAdminApi.GetOrderReturnReason:input_type -> admin.GetOrderReturnReasonReq
	6, // 8: admin.OmsAdminApi.DeleteOrderReturnReason:input_type -> admin.DeleteOrderReturnReasonReq
	9, // 9: admin.OmsAdminApi.CreateOrderReturnReason:output_type -> admin.CommonRsp
	9, // 10: admin.OmsAdminApi.UpdateOrderReturnReason:output_type -> admin.CommonRsp
	3, // 11: admin.OmsAdminApi.GetOrderReturnReasons:output_type -> admin.GetOrderReturnReasonsRsp
	5, // 12: admin.OmsAdminApi.GetOrderReturnReason:output_type -> admin.GetOrderReturnReasonRsp
	9, // 13: admin.OmsAdminApi.DeleteOrderReturnReason:output_type -> admin.CommonRsp
	9, // [9:14] is the sub-list for method output_type
	4, // [4:9] is the sub-list for method input_type
	4, // [4:4] is the sub-list for extension type_name
	4, // [4:4] is the sub-list for extension extendee
	0, // [0:4] is the sub-list for field type_name
}

func init() { file_admin_oms_admin_proto_init() }
func file_admin_oms_admin_proto_init() {
	if File_admin_oms_admin_proto != nil {
		return
	}
	file_admin_admin_proto_init()
	file_model_oms_order_return_reason_proto_init()
	if !protoimpl.UnsafeEnabled {
		file_admin_oms_admin_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*AddOrUpdateOrderReturnReasonParam); i {
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
		file_admin_oms_admin_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetOrderReturnReasonsParam); i {
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
		file_admin_oms_admin_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*OrderReturnReasonsData); i {
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
		file_admin_oms_admin_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetOrderReturnReasonsRsp); i {
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
		file_admin_oms_admin_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetOrderReturnReasonReq); i {
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
		file_admin_oms_admin_proto_msgTypes[5].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetOrderReturnReasonRsp); i {
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
		file_admin_oms_admin_proto_msgTypes[6].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*DeleteOrderReturnReasonReq); i {
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
			RawDescriptor: file_admin_oms_admin_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   7,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_admin_oms_admin_proto_goTypes,
		DependencyIndexes: file_admin_oms_admin_proto_depIdxs,
		MessageInfos:      file_admin_oms_admin_proto_msgTypes,
	}.Build()
	File_admin_oms_admin_proto = out.File
	file_admin_oms_admin_proto_rawDesc = nil
	file_admin_oms_admin_proto_goTypes = nil
	file_admin_oms_admin_proto_depIdxs = nil
}
