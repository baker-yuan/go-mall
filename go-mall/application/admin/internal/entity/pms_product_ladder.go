package entity

// ProductLadder 商品阶梯价格表
// 商品优惠相关表，购买同商品满足一定数量后，可以使用打折价格进行购买。如：买两件商品可以打八折。
type ProductLadder struct {
	ID        uint64  `gorm:"primary_key;auto_increment;comment:主键"`
	ProductID uint64  `gorm:"type:bigint;unsigned;not null;default:0;comment:商品id"`
	Count     uint32  `gorm:"type:int(10);unsigned;not null;default:0;comment:满足的商品数量"`
	Discount  float64 `gorm:"type:decimal(10,2);not null;default:0.00;comment:折扣"`
	Price     float64 `gorm:"type:decimal(10,2);not null;default:0.00;comment:折后价格"`
}

func (c ProductLadder) TableName() string {
	return "pms_product_ladder"
}
