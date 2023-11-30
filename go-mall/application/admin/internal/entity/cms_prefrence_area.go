package entity

// PrefrenceArea 优选专区
type PrefrenceArea struct {
	ID         uint64 `gorm:"column:id;type:bigint;primary_key;auto_increment;comment:主键"`
	Name       string `gorm:"column:name;type:varchar(255);not null;default:'';comment:名称"`
	SubTitle   string `gorm:"column:sub_title;type:varchar(255);not null;default:'';comment:子标题"`
	Pic        []byte `gorm:"column:pic;type:varbinary(500);not null;comment:展示图片"`
	Sort       uint32 `gorm:"column:sort;type:int(10);unsigned;not null;default:0;comment:排序"`
	ShowStatus uint8  `gorm:"column:show_status;type:tinyint(4);unsigned;not null;default:0;comment:显示状态"`
}

func (c PrefrenceArea) TableName() string {
	return "cms_prefrence_area"
}
