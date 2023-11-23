package entity

// ProductOperateLog 商品操作记录表
// 用于记录商品操作记录
type ProductOperateLog struct {
	ID               uint64  `gorm:"type:bigint;primary_key;auto_increment;comment:主键"`
	ProductID        uint64  `gorm:"type:bigint;not null;default:0;comment:商品id"`
	PriceOld         float64 `gorm:"type:decimal(10,2);not null;default:0.00;comment:改变前价格"`
	PriceNew         float64 `gorm:"type:decimal(10,2);not null;default:0.00;comment:改变后价格"`
	SalePriceOld     float64 `gorm:"type:decimal(10,2);not null;default:0.00;comment:改变前优惠价"`
	SalePriceNew     float64 `gorm:"type:decimal(10,2);not null;default:0.00;comment:改变后优惠价"`
	GiftPointOld     uint32  `gorm:"type:int(10);unsigned;not null;default:0;comment:改变前积分"`
	GiftPointNew     uint32  `gorm:"type:int(10);unsigned;not null;default:0;comment:改变后积分"`
	UsePointLimitOld uint32  `gorm:"type:int(10);unsigned;not null;default:0;comment:改变前积分使用限制"`
	UsePointLimitNew uint32  `gorm:"type:int(10);unsigned;not null;default:0;comment:改变后积分使用限制"`
	OperateMan       string  `gorm:"type:varchar(64);not null;default:'';comment:操作人"`
	CreateTime       uint32  `gorm:"type:int(10);unsigned;not null;default:0;comment:创建时间"`
	// 公共字段
	BaseTime
}

func (c ProductOperateLog) TableName() string {
	return "pms_product_operate_log"
}
