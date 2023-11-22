package entity

// ProductAttributeCategory 产品属性分类表
type ProductAttributeCategory struct {
	ID   uint64 `gorm:"primary_key;auto_increment;comment:编号"`
	Name string `gorm:"type:varchar(64);not null;default:'';comment:类型名称"`
	BaseTime
	// 冗余字段
	AttributeCount int `gorm:"type:int(10);unsigned;not null;default:0;comment:属性数量"` // pms_product_attribute type=0
	ParamCount     int `gorm:"type:int(10);unsigned;not null;default:0;comment:参数数量"` // pms_product_attribute type=1
}

func (c ProductAttributeCategory) TableName() string {
	return "pms_product_attribute_category"
}
