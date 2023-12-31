// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.30.0
// 	protoc        v3.21.12
// source: model/pms_sku_stock.proto

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

// SkuStock 商品SKU表
// SKU（Stock Keeping Unit）是指库存量单位，SPU（Standard Product Unit）是指标准产品单位。举个例子：iphone xs是一个SPU，而iphone xs 公开版 64G 银色是一个SKU。
type SkuStock struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id        uint64 `protobuf:"varint,1,opt,name=id,proto3" json:"id,omitempty"`                                // 主键
	SkuCode   string `protobuf:"bytes,3,opt,name=sku_code,json=skuCode,proto3" json:"sku_code,omitempty"`        // sku编码
	Pic       string `protobuf:"bytes,7,opt,name=pic,proto3" json:"pic,omitempty"`                               // 展示图片
	Sale      uint32 `protobuf:"varint,8,opt,name=sale,proto3" json:"sale,omitempty"`                            // 销量
	SpData    string `protobuf:"bytes,11,opt,name=sp_data,json=spData,proto3" json:"sp_data,omitempty"`          // 商品销售属性，json格式
	ProductId uint64 `protobuf:"varint,2,opt,name=product_id,json=productId,proto3" json:"product_id,omitempty"` // 产品ID
	// 价格
	Price          string `protobuf:"bytes,4,opt,name=price,proto3" json:"price,omitempty"`                                         // 价格
	PromotionPrice string `protobuf:"bytes,9,opt,name=promotion_price,json=promotionPrice,proto3" json:"promotion_price,omitempty"` // 单品促销价格
	// 库存
	Stock     uint32 `protobuf:"varint,5,opt,name=stock,proto3" json:"stock,omitempty"`                           // 库存
	LowStock  uint32 `protobuf:"varint,6,opt,name=low_stock,json=lowStock,proto3" json:"low_stock,omitempty"`     // 预警库存
	LockStock uint32 `protobuf:"varint,10,opt,name=lock_stock,json=lockStock,proto3" json:"lock_stock,omitempty"` // 锁定库存
}

func (x *SkuStock) Reset() {
	*x = SkuStock{}
	if protoimpl.UnsafeEnabled {
		mi := &file_model_pms_sku_stock_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *SkuStock) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*SkuStock) ProtoMessage() {}

func (x *SkuStock) ProtoReflect() protoreflect.Message {
	mi := &file_model_pms_sku_stock_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use SkuStock.ProtoReflect.Descriptor instead.
func (*SkuStock) Descriptor() ([]byte, []int) {
	return file_model_pms_sku_stock_proto_rawDescGZIP(), []int{0}
}

func (x *SkuStock) GetId() uint64 {
	if x != nil {
		return x.Id
	}
	return 0
}

func (x *SkuStock) GetSkuCode() string {
	if x != nil {
		return x.SkuCode
	}
	return ""
}

func (x *SkuStock) GetPic() string {
	if x != nil {
		return x.Pic
	}
	return ""
}

func (x *SkuStock) GetSale() uint32 {
	if x != nil {
		return x.Sale
	}
	return 0
}

func (x *SkuStock) GetSpData() string {
	if x != nil {
		return x.SpData
	}
	return ""
}

func (x *SkuStock) GetProductId() uint64 {
	if x != nil {
		return x.ProductId
	}
	return 0
}

func (x *SkuStock) GetPrice() string {
	if x != nil {
		return x.Price
	}
	return ""
}

func (x *SkuStock) GetPromotionPrice() string {
	if x != nil {
		return x.PromotionPrice
	}
	return ""
}

func (x *SkuStock) GetStock() uint32 {
	if x != nil {
		return x.Stock
	}
	return 0
}

func (x *SkuStock) GetLowStock() uint32 {
	if x != nil {
		return x.LowStock
	}
	return 0
}

func (x *SkuStock) GetLockStock() uint32 {
	if x != nil {
		return x.LockStock
	}
	return 0
}

var File_model_pms_sku_stock_proto protoreflect.FileDescriptor

var file_model_pms_sku_stock_proto_rawDesc = []byte{
	0x0a, 0x19, 0x6d, 0x6f, 0x64, 0x65, 0x6c, 0x2f, 0x70, 0x6d, 0x73, 0x5f, 0x73, 0x6b, 0x75, 0x5f,
	0x73, 0x74, 0x6f, 0x63, 0x6b, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x05, 0x6d, 0x6f, 0x64,
	0x65, 0x6c, 0x22, 0xa4, 0x02, 0x0a, 0x08, 0x53, 0x6b, 0x75, 0x53, 0x74, 0x6f, 0x63, 0x6b, 0x12,
	0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x04, 0x52, 0x02, 0x69, 0x64, 0x12,
	0x19, 0x0a, 0x08, 0x73, 0x6b, 0x75, 0x5f, 0x63, 0x6f, 0x64, 0x65, 0x18, 0x03, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x07, 0x73, 0x6b, 0x75, 0x43, 0x6f, 0x64, 0x65, 0x12, 0x10, 0x0a, 0x03, 0x70, 0x69,
	0x63, 0x18, 0x07, 0x20, 0x01, 0x28, 0x09, 0x52, 0x03, 0x70, 0x69, 0x63, 0x12, 0x12, 0x0a, 0x04,
	0x73, 0x61, 0x6c, 0x65, 0x18, 0x08, 0x20, 0x01, 0x28, 0x0d, 0x52, 0x04, 0x73, 0x61, 0x6c, 0x65,
	0x12, 0x17, 0x0a, 0x07, 0x73, 0x70, 0x5f, 0x64, 0x61, 0x74, 0x61, 0x18, 0x0b, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x06, 0x73, 0x70, 0x44, 0x61, 0x74, 0x61, 0x12, 0x1d, 0x0a, 0x0a, 0x70, 0x72, 0x6f,
	0x64, 0x75, 0x63, 0x74, 0x5f, 0x69, 0x64, 0x18, 0x02, 0x20, 0x01, 0x28, 0x04, 0x52, 0x09, 0x70,
	0x72, 0x6f, 0x64, 0x75, 0x63, 0x74, 0x49, 0x64, 0x12, 0x14, 0x0a, 0x05, 0x70, 0x72, 0x69, 0x63,
	0x65, 0x18, 0x04, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x70, 0x72, 0x69, 0x63, 0x65, 0x12, 0x27,
	0x0a, 0x0f, 0x70, 0x72, 0x6f, 0x6d, 0x6f, 0x74, 0x69, 0x6f, 0x6e, 0x5f, 0x70, 0x72, 0x69, 0x63,
	0x65, 0x18, 0x09, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0e, 0x70, 0x72, 0x6f, 0x6d, 0x6f, 0x74, 0x69,
	0x6f, 0x6e, 0x50, 0x72, 0x69, 0x63, 0x65, 0x12, 0x14, 0x0a, 0x05, 0x73, 0x74, 0x6f, 0x63, 0x6b,
	0x18, 0x05, 0x20, 0x01, 0x28, 0x0d, 0x52, 0x05, 0x73, 0x74, 0x6f, 0x63, 0x6b, 0x12, 0x1b, 0x0a,
	0x09, 0x6c, 0x6f, 0x77, 0x5f, 0x73, 0x74, 0x6f, 0x63, 0x6b, 0x18, 0x06, 0x20, 0x01, 0x28, 0x0d,
	0x52, 0x08, 0x6c, 0x6f, 0x77, 0x53, 0x74, 0x6f, 0x63, 0x6b, 0x12, 0x1d, 0x0a, 0x0a, 0x6c, 0x6f,
	0x63, 0x6b, 0x5f, 0x73, 0x74, 0x6f, 0x63, 0x6b, 0x18, 0x0a, 0x20, 0x01, 0x28, 0x0d, 0x52, 0x09,
	0x6c, 0x6f, 0x63, 0x6b, 0x53, 0x74, 0x6f, 0x63, 0x6b, 0x42, 0x08, 0x5a, 0x06, 0x2e, 0x3b, 0x6d,
	0x61, 0x6c, 0x6c, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_model_pms_sku_stock_proto_rawDescOnce sync.Once
	file_model_pms_sku_stock_proto_rawDescData = file_model_pms_sku_stock_proto_rawDesc
)

func file_model_pms_sku_stock_proto_rawDescGZIP() []byte {
	file_model_pms_sku_stock_proto_rawDescOnce.Do(func() {
		file_model_pms_sku_stock_proto_rawDescData = protoimpl.X.CompressGZIP(file_model_pms_sku_stock_proto_rawDescData)
	})
	return file_model_pms_sku_stock_proto_rawDescData
}

var file_model_pms_sku_stock_proto_msgTypes = make([]protoimpl.MessageInfo, 1)
var file_model_pms_sku_stock_proto_goTypes = []interface{}{
	(*SkuStock)(nil), // 0: model.SkuStock
}
var file_model_pms_sku_stock_proto_depIdxs = []int32{
	0, // [0:0] is the sub-list for method output_type
	0, // [0:0] is the sub-list for method input_type
	0, // [0:0] is the sub-list for extension type_name
	0, // [0:0] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_model_pms_sku_stock_proto_init() }
func file_model_pms_sku_stock_proto_init() {
	if File_model_pms_sku_stock_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_model_pms_sku_stock_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*SkuStock); i {
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
			RawDescriptor: file_model_pms_sku_stock_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   1,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_model_pms_sku_stock_proto_goTypes,
		DependencyIndexes: file_model_pms_sku_stock_proto_depIdxs,
		MessageInfos:      file_model_pms_sku_stock_proto_msgTypes,
	}.Build()
	File_model_pms_sku_stock_proto = out.File
	file_model_pms_sku_stock_proto_rawDesc = nil
	file_model_pms_sku_stock_proto_goTypes = nil
	file_model_pms_sku_stock_proto_depIdxs = nil
}
