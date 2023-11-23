package entity

// OrderSetting 订单设置表
// 用于对订单的一些超时操作进行设置。
type OrderSetting struct {
	ID                  uint64 `gorm:"column:id;type:bigint;primary_key;auto_increment"`
	FlashOrderOvertime  uint32 `gorm:"column:flash_order_overtime;type:int(10);unsigned;not null;default:0;comment:秒杀订单超时关闭时间(分)"`
	NormalOrderOvertime uint32 `gorm:"column:normal_order_overtime;type:int(10);unsigned;not null;default:0;comment:正常订单超时时间(分)"`
	ConfirmOvertime     uint32 `gorm:"column:confirm_overtime;type:int(10);unsigned;not null;default:0;comment:发货后自动确认收货时间（天）"`
	FinishOvertime      uint32 `gorm:"column:finish_overtime;type:int(10);unsigned;not null;default:0;comment:自动完成交易时间，不能申请售后（天）"`
	CommentOvertime     uint32 `gorm:"column:comment_overtime;type:int(10);unsigned;not null;default:0;comment:订单完成后自动好评时间（天）"`
}

func (o OrderSetting) TableName() string {
	return "oms_order_setting"
}
