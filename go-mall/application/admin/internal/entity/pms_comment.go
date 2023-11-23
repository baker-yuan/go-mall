package entity

// Comment 商品评论表
type Comment struct {
	ID               uint64 `gorm:"type:bigint;primary_key;auto_increment;comment:主键"`
	ProductID        uint64 `gorm:"type:bigint;not null;default:0;comment:商品id"` // pms_product#id
	MemberNickName   string `gorm:"type:varchar(255);not null;default:'';comment:会员昵称"`
	ProductName      string `gorm:"type:varchar(255);not null;default:'';comment:商品名称"`
	Star             uint8  `gorm:"type:tinyint(4);unsigned;not null;default:0;comment:评价星数：0->5"`
	MemberIP         string `gorm:"type:varchar(64);not null;default:'';comment:评价的ip"`
	CreateTime       uint32 `gorm:"type:int(10);unsigned;not null;default:0;comment:创建时间"`
	ShowStatus       uint8  `gorm:"type:tinyint(4);unsigned;not null;default:0;comment:是否显示"`
	ProductAttribute string `gorm:"type:varchar(255);not null;default:'';comment:购买时的商品属性"`
	CollectCount     uint32 `gorm:"type:int(10);unsigned;not null;default:0;comment:收藏数"`
	ReadCount        uint32 `gorm:"type:int(10);unsigned;not null;default:0;comment:阅读数"`
	Content          string `gorm:"type:text;not null;comment:内容"`
	Pics             string `gorm:"type:varchar(1000);not null;default:'';comment:上传图片地址，以逗号隔开"`
	MemberIcon       string `gorm:"type:varchar(255);not null;default:'';comment:评论用户头像"`
	ReplayCount      uint32 `gorm:"type:int(10);unsigned;not null;default:0;comment:回复数"`
	// 公共字段
	BaseTime
}

func (c Comment) TableName() string {
	return "pms_comment"
}
