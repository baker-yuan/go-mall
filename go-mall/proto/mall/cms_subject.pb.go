// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.30.0
// 	protoc        v3.21.12
// source: model/cms_subject.proto

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

// Subject 专题
type Subject struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id              uint64 `protobuf:"varint,1,opt,name=id,proto3" json:"id,omitempty"`                                                  // 主键
	CategoryId      uint64 `protobuf:"varint,2,opt,name=category_id,json=categoryId,proto3" json:"category_id,omitempty"`                // 分类id
	Title           string `protobuf:"bytes,3,opt,name=title,proto3" json:"title,omitempty"`                                             // 标题
	Pic             string `protobuf:"bytes,4,opt,name=pic,proto3" json:"pic,omitempty"`                                                 // 专题主图
	AlbumPics       string `protobuf:"bytes,5,opt,name=album_pics,json=albumPics,proto3" json:"album_pics,omitempty"`                    // 画册图片，用逗号分割
	Description     string `protobuf:"bytes,6,opt,name=description,proto3" json:"description,omitempty"`                                 // 描述
	Content         string `protobuf:"bytes,7,opt,name=content,proto3" json:"content,omitempty"`                                         // 内容
	ShowStatus      uint32 `protobuf:"varint,8,opt,name=show_status,json=showStatus,proto3" json:"show_status,omitempty"`                // 显示状态：0->不显示；1->显示
	RecommendStatus uint32 `protobuf:"varint,9,opt,name=recommend_status,json=recommendStatus,proto3" json:"recommend_status,omitempty"` // 推荐状态
	CreateTime      uint32 `protobuf:"varint,10,opt,name=create_time,json=createTime,proto3" json:"create_time,omitempty"`               // 创建时间
	// 冗余字段
	ForwardCount uint32 `protobuf:"varint,11,opt,name=forward_count,json=forwardCount,proto3" json:"forward_count,omitempty"` // 转发数
	CollectCount uint32 `protobuf:"varint,12,opt,name=collect_count,json=collectCount,proto3" json:"collect_count,omitempty"` // 收藏数量
	ReadCount    uint32 `protobuf:"varint,13,opt,name=read_count,json=readCount,proto3" json:"read_count,omitempty"`          // 阅读数量
	CommentCount uint32 `protobuf:"varint,14,opt,name=comment_count,json=commentCount,proto3" json:"comment_count,omitempty"` // 评论数量
	ProductCount uint32 `protobuf:"varint,15,opt,name=product_count,json=productCount,proto3" json:"product_count,omitempty"` // 关联产品数量
	CategoryName string `protobuf:"bytes,16,opt,name=category_name,json=categoryName,proto3" json:"category_name,omitempty"`  // 专题分类名称
}

func (x *Subject) Reset() {
	*x = Subject{}
	if protoimpl.UnsafeEnabled {
		mi := &file_model_cms_subject_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Subject) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Subject) ProtoMessage() {}

func (x *Subject) ProtoReflect() protoreflect.Message {
	mi := &file_model_cms_subject_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Subject.ProtoReflect.Descriptor instead.
func (*Subject) Descriptor() ([]byte, []int) {
	return file_model_cms_subject_proto_rawDescGZIP(), []int{0}
}

func (x *Subject) GetId() uint64 {
	if x != nil {
		return x.Id
	}
	return 0
}

func (x *Subject) GetCategoryId() uint64 {
	if x != nil {
		return x.CategoryId
	}
	return 0
}

func (x *Subject) GetTitle() string {
	if x != nil {
		return x.Title
	}
	return ""
}

func (x *Subject) GetPic() string {
	if x != nil {
		return x.Pic
	}
	return ""
}

func (x *Subject) GetAlbumPics() string {
	if x != nil {
		return x.AlbumPics
	}
	return ""
}

func (x *Subject) GetDescription() string {
	if x != nil {
		return x.Description
	}
	return ""
}

func (x *Subject) GetContent() string {
	if x != nil {
		return x.Content
	}
	return ""
}

func (x *Subject) GetShowStatus() uint32 {
	if x != nil {
		return x.ShowStatus
	}
	return 0
}

func (x *Subject) GetRecommendStatus() uint32 {
	if x != nil {
		return x.RecommendStatus
	}
	return 0
}

func (x *Subject) GetCreateTime() uint32 {
	if x != nil {
		return x.CreateTime
	}
	return 0
}

func (x *Subject) GetForwardCount() uint32 {
	if x != nil {
		return x.ForwardCount
	}
	return 0
}

func (x *Subject) GetCollectCount() uint32 {
	if x != nil {
		return x.CollectCount
	}
	return 0
}

func (x *Subject) GetReadCount() uint32 {
	if x != nil {
		return x.ReadCount
	}
	return 0
}

func (x *Subject) GetCommentCount() uint32 {
	if x != nil {
		return x.CommentCount
	}
	return 0
}

func (x *Subject) GetProductCount() uint32 {
	if x != nil {
		return x.ProductCount
	}
	return 0
}

func (x *Subject) GetCategoryName() string {
	if x != nil {
		return x.CategoryName
	}
	return ""
}

var File_model_cms_subject_proto protoreflect.FileDescriptor

var file_model_cms_subject_proto_rawDesc = []byte{
	0x0a, 0x17, 0x6d, 0x6f, 0x64, 0x65, 0x6c, 0x2f, 0x63, 0x6d, 0x73, 0x5f, 0x73, 0x75, 0x62, 0x6a,
	0x65, 0x63, 0x74, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x05, 0x6d, 0x6f, 0x64, 0x65, 0x6c,
	0x22, 0x82, 0x04, 0x0a, 0x07, 0x53, 0x75, 0x62, 0x6a, 0x65, 0x63, 0x74, 0x12, 0x0e, 0x0a, 0x02,
	0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x04, 0x52, 0x02, 0x69, 0x64, 0x12, 0x1f, 0x0a, 0x0b,
	0x63, 0x61, 0x74, 0x65, 0x67, 0x6f, 0x72, 0x79, 0x5f, 0x69, 0x64, 0x18, 0x02, 0x20, 0x01, 0x28,
	0x04, 0x52, 0x0a, 0x63, 0x61, 0x74, 0x65, 0x67, 0x6f, 0x72, 0x79, 0x49, 0x64, 0x12, 0x14, 0x0a,
	0x05, 0x74, 0x69, 0x74, 0x6c, 0x65, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x74, 0x69,
	0x74, 0x6c, 0x65, 0x12, 0x10, 0x0a, 0x03, 0x70, 0x69, 0x63, 0x18, 0x04, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x03, 0x70, 0x69, 0x63, 0x12, 0x1d, 0x0a, 0x0a, 0x61, 0x6c, 0x62, 0x75, 0x6d, 0x5f, 0x70,
	0x69, 0x63, 0x73, 0x18, 0x05, 0x20, 0x01, 0x28, 0x09, 0x52, 0x09, 0x61, 0x6c, 0x62, 0x75, 0x6d,
	0x50, 0x69, 0x63, 0x73, 0x12, 0x20, 0x0a, 0x0b, 0x64, 0x65, 0x73, 0x63, 0x72, 0x69, 0x70, 0x74,
	0x69, 0x6f, 0x6e, 0x18, 0x06, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0b, 0x64, 0x65, 0x73, 0x63, 0x72,
	0x69, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x12, 0x18, 0x0a, 0x07, 0x63, 0x6f, 0x6e, 0x74, 0x65, 0x6e,
	0x74, 0x18, 0x07, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x63, 0x6f, 0x6e, 0x74, 0x65, 0x6e, 0x74,
	0x12, 0x1f, 0x0a, 0x0b, 0x73, 0x68, 0x6f, 0x77, 0x5f, 0x73, 0x74, 0x61, 0x74, 0x75, 0x73, 0x18,
	0x08, 0x20, 0x01, 0x28, 0x0d, 0x52, 0x0a, 0x73, 0x68, 0x6f, 0x77, 0x53, 0x74, 0x61, 0x74, 0x75,
	0x73, 0x12, 0x29, 0x0a, 0x10, 0x72, 0x65, 0x63, 0x6f, 0x6d, 0x6d, 0x65, 0x6e, 0x64, 0x5f, 0x73,
	0x74, 0x61, 0x74, 0x75, 0x73, 0x18, 0x09, 0x20, 0x01, 0x28, 0x0d, 0x52, 0x0f, 0x72, 0x65, 0x63,
	0x6f, 0x6d, 0x6d, 0x65, 0x6e, 0x64, 0x53, 0x74, 0x61, 0x74, 0x75, 0x73, 0x12, 0x1f, 0x0a, 0x0b,
	0x63, 0x72, 0x65, 0x61, 0x74, 0x65, 0x5f, 0x74, 0x69, 0x6d, 0x65, 0x18, 0x0a, 0x20, 0x01, 0x28,
	0x0d, 0x52, 0x0a, 0x63, 0x72, 0x65, 0x61, 0x74, 0x65, 0x54, 0x69, 0x6d, 0x65, 0x12, 0x23, 0x0a,
	0x0d, 0x66, 0x6f, 0x72, 0x77, 0x61, 0x72, 0x64, 0x5f, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x18, 0x0b,
	0x20, 0x01, 0x28, 0x0d, 0x52, 0x0c, 0x66, 0x6f, 0x72, 0x77, 0x61, 0x72, 0x64, 0x43, 0x6f, 0x75,
	0x6e, 0x74, 0x12, 0x23, 0x0a, 0x0d, 0x63, 0x6f, 0x6c, 0x6c, 0x65, 0x63, 0x74, 0x5f, 0x63, 0x6f,
	0x75, 0x6e, 0x74, 0x18, 0x0c, 0x20, 0x01, 0x28, 0x0d, 0x52, 0x0c, 0x63, 0x6f, 0x6c, 0x6c, 0x65,
	0x63, 0x74, 0x43, 0x6f, 0x75, 0x6e, 0x74, 0x12, 0x1d, 0x0a, 0x0a, 0x72, 0x65, 0x61, 0x64, 0x5f,
	0x63, 0x6f, 0x75, 0x6e, 0x74, 0x18, 0x0d, 0x20, 0x01, 0x28, 0x0d, 0x52, 0x09, 0x72, 0x65, 0x61,
	0x64, 0x43, 0x6f, 0x75, 0x6e, 0x74, 0x12, 0x23, 0x0a, 0x0d, 0x63, 0x6f, 0x6d, 0x6d, 0x65, 0x6e,
	0x74, 0x5f, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x18, 0x0e, 0x20, 0x01, 0x28, 0x0d, 0x52, 0x0c, 0x63,
	0x6f, 0x6d, 0x6d, 0x65, 0x6e, 0x74, 0x43, 0x6f, 0x75, 0x6e, 0x74, 0x12, 0x23, 0x0a, 0x0d, 0x70,
	0x72, 0x6f, 0x64, 0x75, 0x63, 0x74, 0x5f, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x18, 0x0f, 0x20, 0x01,
	0x28, 0x0d, 0x52, 0x0c, 0x70, 0x72, 0x6f, 0x64, 0x75, 0x63, 0x74, 0x43, 0x6f, 0x75, 0x6e, 0x74,
	0x12, 0x23, 0x0a, 0x0d, 0x63, 0x61, 0x74, 0x65, 0x67, 0x6f, 0x72, 0x79, 0x5f, 0x6e, 0x61, 0x6d,
	0x65, 0x18, 0x10, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0c, 0x63, 0x61, 0x74, 0x65, 0x67, 0x6f, 0x72,
	0x79, 0x4e, 0x61, 0x6d, 0x65, 0x42, 0x08, 0x5a, 0x06, 0x2e, 0x3b, 0x6d, 0x61, 0x6c, 0x6c, 0x62,
	0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_model_cms_subject_proto_rawDescOnce sync.Once
	file_model_cms_subject_proto_rawDescData = file_model_cms_subject_proto_rawDesc
)

func file_model_cms_subject_proto_rawDescGZIP() []byte {
	file_model_cms_subject_proto_rawDescOnce.Do(func() {
		file_model_cms_subject_proto_rawDescData = protoimpl.X.CompressGZIP(file_model_cms_subject_proto_rawDescData)
	})
	return file_model_cms_subject_proto_rawDescData
}

var file_model_cms_subject_proto_msgTypes = make([]protoimpl.MessageInfo, 1)
var file_model_cms_subject_proto_goTypes = []interface{}{
	(*Subject)(nil), // 0: model.Subject
}
var file_model_cms_subject_proto_depIdxs = []int32{
	0, // [0:0] is the sub-list for method output_type
	0, // [0:0] is the sub-list for method input_type
	0, // [0:0] is the sub-list for extension type_name
	0, // [0:0] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_model_cms_subject_proto_init() }
func file_model_cms_subject_proto_init() {
	if File_model_cms_subject_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_model_cms_subject_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Subject); i {
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
			RawDescriptor: file_model_cms_subject_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   1,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_model_cms_subject_proto_goTypes,
		DependencyIndexes: file_model_cms_subject_proto_depIdxs,
		MessageInfos:      file_model_cms_subject_proto_msgTypes,
	}.Build()
	File_model_cms_subject_proto = out.File
	file_model_cms_subject_proto_rawDesc = nil
	file_model_cms_subject_proto_goTypes = nil
	file_model_cms_subject_proto_depIdxs = nil
}