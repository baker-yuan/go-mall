package entity

// ProductCategory 商品分类表
type ProductCategory struct {
	ID          uint64 `gorm:"type:bigint;primary_key;auto_increment;comment:分类编号"`
	ParentID    uint64 `gorm:"type:bigint;unsigned;not null;default:0;comment:上级分类的编号：0表示一级分类"`
	Name        string `gorm:"type:varchar(64);not null;default:'';comment:分类名称"`
	ProductUnit string `gorm:"type:varchar(64);not null;default:'';comment:商品单位"`
	NavStatus   uint8  `gorm:"type:tinyint(4);unsigned;not null;default:0;comment:是否显示在导航栏：0->不显示；1->显示"`
	ShowStatus  uint8  `gorm:"type:tinyint(4);unsigned;not null;default:0;comment:显示状态：0->不显示；1->显示"`
	Sort        uint32 `gorm:"type:int(10);unsigned;not null;default:0;comment:排序"`
	Icon        string `gorm:"type:varchar(255);not null;default:'';comment:图标"`
	Keywords    string `gorm:"type:varchar(255);not null;default:'';comment:关键字"`
	Description string `gorm:"type:text;not null;comment:描述"`
	// 计算得出
	Level uint8 `gorm:"type:tinyint(4);not null;default:0;comment:分类级别：0->1级；1->2级"`
	// 冗余字段
	ProductCount uint32 `gorm:"type:int(10);unsigned;not null;default:0;comment:商品数量"`
	// 公共字段
	BaseTime
}

func (c ProductCategory) TableName() string {
	return "pms_product_category"
}
