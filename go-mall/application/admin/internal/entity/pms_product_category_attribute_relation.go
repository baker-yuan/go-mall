package entity

// ProductCategoryAttributeRelation 产品的分类和属性的关系表，用于设置分类筛选条件（只支持一级分类）
type ProductCategoryAttributeRelation struct {
	ID                 uint64 `gorm:"primary_key;auto_increment;comment:主键"`                     //
	ProductCategoryID  uint64 `gorm:"type:bigint;unsigned;not null;default:0;comment:产品分类表ID"`   // pms_product_category#id
	ProductAttributeID uint64 `gorm:"type:bigint;unsigned;not null;default:0;comment:商品属性参数表ID"` // pms_product_attribute#id
}

func (c ProductCategoryAttributeRelation) TableName() string {
	return "pms_product_category_attribute_relation"
}
