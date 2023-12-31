// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.30.0
// 	protoc        v3.21.12
// source: model/sms_home_advertise.proto

package mall

import (
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
	reflect "reflect"
	sync "sync"
)

const (
	// Verify that this generated code is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)
	// Verify that runtime/protoimpl is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)

// 首页轮播广告表
type HomeAdvertise struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id   uint64 `protobuf:"varint,1,opt,name=id,proto3" json:"id,omitempty"`     // 主键
	Name string `protobuf:"bytes,2,opt,name=name,proto3" json:"name,omitempty"`  // 名称
	Pic  string `protobuf:"bytes,3,opt,name=pic,proto3" json:"pic,omitempty"`    // 图片地址
	Url  string `protobuf:"bytes,4,opt,name=url,proto3" json:"url,omitempty"`    // 链接地址
	Sort uint32 `protobuf:"varint,5,opt,name=sort,proto3" json:"sort,omitempty"` // 排序
	Note string `protobuf:"bytes,6,opt,name=note,proto3" json:"note,omitempty"`  // 备注
	// 类型
	Type uint32 `protobuf:"varint,7,opt,name=type,proto3" json:"type,omitempty"` // 轮播位置：0->PC首页轮播；1->app首页轮播，注意：在proto中使用uint32代替uint8
	// 时间
	StartTime uint32 `protobuf:"varint,8,opt,name=startTime,proto3" json:"startTime,omitempty"` // 开始时间
	EndTime   uint32 `protobuf:"varint,9,opt,name=endTime,proto3" json:"endTime,omitempty"`     // 结束时间
	// 状态
	Status uint32 `protobuf:"varint,10,opt,name=status,proto3" json:"status,omitempty"` // 上下线状态：0->下线；1->上线，注意：在proto中使用uint32代替uint8
	// 统计
	ClickCount uint32 `protobuf:"varint,11,opt,name=clickCount,proto3" json:"clickCount,omitempty"` // 点击数
	OrderCount uint32 `protobuf:"varint,12,opt,name=orderCount,proto3" json:"orderCount,omitempty"` // 下单数
}

func (x *HomeAdvertise) Reset() {
	*x = HomeAdvertise{}
	if protoimpl.UnsafeEnabled {
		mi := &file_model_sms_home_advertise_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *HomeAdvertise) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*HomeAdvertise) ProtoMessage() {}

func (x *HomeAdvertise) ProtoReflect() protoreflect.Message {
	mi := &file_model_sms_home_advertise_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use HomeAdvertise.ProtoReflect.Descriptor instead.
func (*HomeAdvertise) Descriptor() ([]byte, []int) {
	return file_model_sms_home_advertise_proto_rawDescGZIP(), []int{0}
}

func (x *HomeAdvertise) GetId() uint64 {
	if x != nil {
		return x.Id
	}
	return 0
}

func (x *HomeAdvertise) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *HomeAdvertise) GetPic() string {
	if x != nil {
		return x.Pic
	}
	return ""
}

func (x *HomeAdvertise) GetUrl() string {
	if x != nil {
		return x.Url
	}
	return ""
}

func (x *HomeAdvertise) GetSort() uint32 {
	if x != nil {
		return x.Sort
	}
	return 0
}

func (x *HomeAdvertise) GetNote() string {
	if x != nil {
		return x.Note
	}
	return ""
}

func (x *HomeAdvertise) GetType() uint32 {
	if x != nil {
		return x.Type
	}
	return 0
}

func (x *HomeAdvertise) GetStartTime() uint32 {
	if x != nil {
		return x.StartTime
	}
	return 0
}

func (x *HomeAdvertise) GetEndTime() uint32 {
	if x != nil {
		return x.EndTime
	}
	return 0
}

func (x *HomeAdvertise) GetStatus() uint32 {
	if x != nil {
		return x.Status
	}
	return 0
}

func (x *HomeAdvertise) GetClickCount() uint32 {
	if x != nil {
		return x.ClickCount
	}
	return 0
}

func (x *HomeAdvertise) GetOrderCount() uint32 {
	if x != nil {
		return x.OrderCount
	}
	return 0
}

var File_model_sms_home_advertise_proto protoreflect.FileDescriptor

var file_model_sms_home_advertise_proto_rawDesc = []byte{
	0x0a, 0x1e, 0x6d, 0x6f, 0x64, 0x65, 0x6c, 0x2f, 0x73, 0x6d, 0x73, 0x5f, 0x68, 0x6f, 0x6d, 0x65,
	0x5f, 0x61, 0x64, 0x76, 0x65, 0x72, 0x74, 0x69, 0x73, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x12, 0x05, 0x6d, 0x6f, 0x64, 0x65, 0x6c, 0x22, 0xa3, 0x02, 0x0a, 0x0d, 0x48, 0x6f, 0x6d, 0x65,
	0x41, 0x64, 0x76, 0x65, 0x72, 0x74, 0x69, 0x73, 0x65, 0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18,
	0x01, 0x20, 0x01, 0x28, 0x04, 0x52, 0x02, 0x69, 0x64, 0x12, 0x12, 0x0a, 0x04, 0x6e, 0x61, 0x6d,
	0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x12, 0x10, 0x0a,
	0x03, 0x70, 0x69, 0x63, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x03, 0x70, 0x69, 0x63, 0x12,
	0x10, 0x0a, 0x03, 0x75, 0x72, 0x6c, 0x18, 0x04, 0x20, 0x01, 0x28, 0x09, 0x52, 0x03, 0x75, 0x72,
	0x6c, 0x12, 0x12, 0x0a, 0x04, 0x73, 0x6f, 0x72, 0x74, 0x18, 0x05, 0x20, 0x01, 0x28, 0x0d, 0x52,
	0x04, 0x73, 0x6f, 0x72, 0x74, 0x12, 0x12, 0x0a, 0x04, 0x6e, 0x6f, 0x74, 0x65, 0x18, 0x06, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x04, 0x6e, 0x6f, 0x74, 0x65, 0x12, 0x12, 0x0a, 0x04, 0x74, 0x79, 0x70,
	0x65, 0x18, 0x07, 0x20, 0x01, 0x28, 0x0d, 0x52, 0x04, 0x74, 0x79, 0x70, 0x65, 0x12, 0x1c, 0x0a,
	0x09, 0x73, 0x74, 0x61, 0x72, 0x74, 0x54, 0x69, 0x6d, 0x65, 0x18, 0x08, 0x20, 0x01, 0x28, 0x0d,
	0x52, 0x09, 0x73, 0x74, 0x61, 0x72, 0x74, 0x54, 0x69, 0x6d, 0x65, 0x12, 0x18, 0x0a, 0x07, 0x65,
	0x6e, 0x64, 0x54, 0x69, 0x6d, 0x65, 0x18, 0x09, 0x20, 0x01, 0x28, 0x0d, 0x52, 0x07, 0x65, 0x6e,
	0x64, 0x54, 0x69, 0x6d, 0x65, 0x12, 0x16, 0x0a, 0x06, 0x73, 0x74, 0x61, 0x74, 0x75, 0x73, 0x18,
	0x0a, 0x20, 0x01, 0x28, 0x0d, 0x52, 0x06, 0x73, 0x74, 0x61, 0x74, 0x75, 0x73, 0x12, 0x1e, 0x0a,
	0x0a, 0x63, 0x6c, 0x69, 0x63, 0x6b, 0x43, 0x6f, 0x75, 0x6e, 0x74, 0x18, 0x0b, 0x20, 0x01, 0x28,
	0x0d, 0x52, 0x0a, 0x63, 0x6c, 0x69, 0x63, 0x6b, 0x43, 0x6f, 0x75, 0x6e, 0x74, 0x12, 0x1e, 0x0a,
	0x0a, 0x6f, 0x72, 0x64, 0x65, 0x72, 0x43, 0x6f, 0x75, 0x6e, 0x74, 0x18, 0x0c, 0x20, 0x01, 0x28,
	0x0d, 0x52, 0x0a, 0x6f, 0x72, 0x64, 0x65, 0x72, 0x43, 0x6f, 0x75, 0x6e, 0x74, 0x42, 0x08, 0x5a,
	0x06, 0x2e, 0x3b, 0x6d, 0x61, 0x6c, 0x6c, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_model_sms_home_advertise_proto_rawDescOnce sync.Once
	file_model_sms_home_advertise_proto_rawDescData = file_model_sms_home_advertise_proto_rawDesc
)

func file_model_sms_home_advertise_proto_rawDescGZIP() []byte {
	file_model_sms_home_advertise_proto_rawDescOnce.Do(func() {
		file_model_sms_home_advertise_proto_rawDescData = protoimpl.X.CompressGZIP(file_model_sms_home_advertise_proto_rawDescData)
	})
	return file_model_sms_home_advertise_proto_rawDescData
}

var file_model_sms_home_advertise_proto_msgTypes = make([]protoimpl.MessageInfo, 1)
var file_model_sms_home_advertise_proto_goTypes = []interface{}{
	(*HomeAdvertise)(nil), // 0: model.HomeAdvertise
}
var file_model_sms_home_advertise_proto_depIdxs = []int32{
	0, // [0:0] is the sub-list for method output_type
	0, // [0:0] is the sub-list for method input_type
	0, // [0:0] is the sub-list for extension type_name
	0, // [0:0] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_model_sms_home_advertise_proto_init() }
func file_model_sms_home_advertise_proto_init() {
	if File_model_sms_home_advertise_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_model_sms_home_advertise_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*HomeAdvertise); i {
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
			RawDescriptor: file_model_sms_home_advertise_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   1,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_model_sms_home_advertise_proto_goTypes,
		DependencyIndexes: file_model_sms_home_advertise_proto_depIdxs,
		MessageInfos:      file_model_sms_home_advertise_proto_msgTypes,
	}.Build()
	File_model_sms_home_advertise_proto = out.File
	file_model_sms_home_advertise_proto_rawDesc = nil
	file_model_sms_home_advertise_proto_goTypes = nil
	file_model_sms_home_advertise_proto_depIdxs = nil
}
