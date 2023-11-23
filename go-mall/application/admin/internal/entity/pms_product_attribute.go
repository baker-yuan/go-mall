package entity

// ProductAttribute 商品属性参数表
// type字段用于控制其是规格还是参数。
type ProductAttribute struct {
	ID                         uint64 `gorm:"column:id;type:bigint;primary_key;auto_increment;comment:编号"`
	ProductAttributeCategoryID uint64 `gorm:"column:product_attribute_category_id;type:bigint;unsigned;not null;default:0;comment:产品属性分类表ID"` // pms_product_attribute_category#id
	Name                       string `gorm:"column:name;type:varchar(64);not null;default:'';comment:属性名称"`
	SelectType                 uint8  `gorm:"column:select_type;type:tinyint(4);unsigned;not null;default:0;comment:属性选择类型：0->唯一；1->单选；2->多选"`
	InputType                  uint8  `gorm:"column:input_type;type:tinyint(4);unsigned;not null;default:0;comment:属性录入方式：0->手工录入；1->从列表中选取"`
	InputList                  string `gorm:"column:input_list;type:varchar(255);not null;default:'';comment:可选值列表，以逗号隔开"`
	Sort                       int    `gorm:"column:sort;type:int(10);unsigned;not null;default:0;comment:排序字段"`
	FilterType                 uint8  `gorm:"column:filter_type;type:tinyint(4);unsigned;not null;default:0;comment:分类筛选样式：1->普通；1->颜色"`
	SearchType                 uint8  `gorm:"column:search_type;type:tinyint(4);unsigned;not null;default:0;comment:检索类型；0->不需要进行检索；1->关键字检索；2->范围检索"`
	RelatedStatus              uint8  `gorm:"column:related_status;type:tinyint(4);unsigned;not null;default:0;comment:相同属性产品是否关联；0->不关联；1->关联"`
	HandAddStatus              uint8  `gorm:"column:hand_add_status;type:tinyint(4);unsigned;not null;default:0;comment:是否支持手动新增；0->不支持；1->支持"`
	Type                       uint8  `gorm:"column:type;type:tinyint(4);unsigned;not null;default:0;comment:属性的类型；0->规格；1->参数"`
	// 公共字段
	BaseTime
}

func (c ProductAttribute) TableName() string {
	return "pms_product_attribute"
}
