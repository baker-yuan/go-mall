package entity

// MemberPrice 商品会员价格表
// 根据不同会员等级，可以以不同的会员价格购买。此处设计有缺陷，可以做成不同会员等级可以减免多少元或者按多少折扣进行购买。
type MemberPrice struct {
	ID              uint64  `gorm:"type:bigint;primary_key;auto_increment;comment:主键"`
	ProductID       uint64  `gorm:"type:bigint;unsigned;not null;default:0;comment:商品id"` // pms_product#id
	MemberLevelID   uint64  `gorm:"type:bigint;unsigned;not null;default:0;comment:会员等级id"`
	MemberPrice     float64 `gorm:"type:decimal(10,2);not null;default:0.00;comment:会员价格"`
	MemberLevelName string  `gorm:"size:100;not null;default:0.00;comment:会员等级名称"`
	// 公共字段
	BaseTime
}

func (c MemberPrice) TableName() string {
	return "pms_member_price"
}
