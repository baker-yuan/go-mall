package entity

// SkuStock 商品SKU表
// SKU（Stock Keeping Unit）是指库存量单位，SPU（Standard Product Unit）是指标准产品单位。举个例子：iphone xs是一个SPU，而iphone xs 公开版 64G 银色是一个SKU。
type SkuStock struct {
	ID             uint64  `gorm:"primary_key;auto_increment;comment:主键"`
	ProductID      uint64  `gorm:"type:bigint;unsigned;not null;default:0;comment:产品ID"`
	SkuCode        string  `gorm:"type:varchar(64);not null;default:'';comment:sku编码"`
	Price          float64 `gorm:"type:decimal(10,2);not null;default:0.00;comment:价格"`
	Stock          uint32  `gorm:"type:int(10);unsigned;not null;default:0;comment:库存"`
	LowStock       uint32  `gorm:"type:int(10);unsigned;not null;default:0;comment:预警库存"`
	Pic            string  `gorm:"type:varchar(255);not null;default:'';comment:展示图片"`
	Sale           uint32  `gorm:"type:int(10);unsigned;not null;default:0;comment:销量"`
	PromotionPrice float64 `gorm:"type:decimal(10,2);not null;default:0.00;comment:单品促销价格"`
	LockStock      uint32  `gorm:"type:int(10);unsigned;not null;default:0;comment:锁定库存"`
	SpData         string  `gorm:"type:varchar(500);not null;default:'';comment:商品销售属性，json格式"`
}

func (c SkuStock) TableName() string {
	return "pms_sku_stock"
}
