package entity

// Brand 商品品牌表
type Brand struct {
	ID            uint64 `gorm:"type:bigint;primary_key;auto_increment;comment:主键"`
	Name          string `gorm:"type:varchar(64);not null;default:'';comment:名称"`
	FirstLetter   string `gorm:"type:varchar(8);not null;default:'';comment:首字母"`
	Sort          uint32 `gorm:"type:int(10);unsigned;not null;default:0;comment:排序"`
	FactoryStatus uint8  `gorm:"type:tinyint(4);unsigned;not null;default:0;comment:是否为品牌制造商：0->不是；1->是"`
	ShowStatus    uint8  `gorm:"type:tinyint(4);unsigned;not null;default:0;comment:是否显示"`
	Logo          string `gorm:"type:varchar(255);not null;default:'';comment:品牌logo"`
	BigPic        string `gorm:"type:varchar(255);not null;default:'';comment:专区大图"`
	BrandStory    string `gorm:"type:text;not null;comment:品牌故事"`
	// 冗余字段
	ProductCount        uint32 `gorm:"type:int(10);unsigned;not null;default:0;comment:产品数量"`
	ProductCommentCount uint32 `gorm:"type:int(10);unsigned;not null;default:0;comment:产品评论数量"`
	// 公共字段
	BaseTime
}

func (c Brand) TableName() string {
	return "pms_brand"
}
