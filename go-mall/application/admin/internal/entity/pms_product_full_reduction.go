package entity

// ProductFullReduction 商品满减表
// 商品优惠相关表，购买同商品满足一定金额后，可以减免一定金额。如：买满1000减100元。
type ProductFullReduction struct {
	ID          uint64  `gorm:"primary_key;auto_increment;comment:主键"`
	ProductID   uint64  `gorm:"type:bigint;unsigned;not null;default:0;comment:商品id"`
	FullPrice   float64 `gorm:"type:decimal(10,2);not null;default:0.00;comment:商品满足金额"`
	ReducePrice float64 `gorm:"type:decimal(10,2);not null;default:0.00;comment:商品减少金额"`
}

func (c ProductFullReduction) TableName() string {
	return "pms_product_full_reduction"
}
