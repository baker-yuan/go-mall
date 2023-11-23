package entity

// FlashPromotionProductRelation 限时购和商品关系表
type FlashPromotionProductRelation struct {
	ID                      uint64  `gorm:"column:id;type:bigint;primary_key;auto_increment"`
	FlashPromotionID        uint64  `gorm:"column:flash_promotion_id;type:bigint;unsigned;not null;default:0;comment:限时购id"`
	FlashPromotionSessionID uint64  `gorm:"column:flash_promotion_session_id;type:bigint;unsigned;not null;default:0;comment:限时购场次id"`
	ProductID               uint64  `gorm:"column:product_id;type:bigint;unsigned;not null;default:0;comment:商品id"`
	FlashPromotionPrice     float64 `gorm:"column:flash_promotion_price;type:decimal(10,2);not null;default:0.00;comment:限时购价格"`
	FlashPromotionCount     uint32  `gorm:"column:flash_promotion_count;type:int(10);unsigned;not null;default:0;comment:限时购数量"`
	FlashPromotionLimit     uint32  `gorm:"column:flash_promotion_limit;type:int(10);unsigned;not null;default:0;comment:每人限购数量"`
	Sort                    uint32  `gorm:"column:sort;type:int(10);unsigned;not null;default:0;comment:排序"`
}

func (s FlashPromotionProductRelation) TableName() string {
	return "sms_flash_promotion_product_relation"
}
