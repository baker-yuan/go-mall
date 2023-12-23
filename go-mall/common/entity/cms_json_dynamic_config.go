package entity

// JsonDynamicConfig 动态配置
type JsonDynamicConfig struct {
	ID         uint64  `gorm:"column:id;type:bigint;primary_key;auto_increment;comment:主键"`
	BizType    BizType `gorm:"uniqueIndex:uk_biz_type;column:biz_type;type:varchar(64);not null;default:'';comment:业务类型"`
	BizDesc    string  `gorm:"column:biz_desc;type:varchar(64);not null;default:'';comment:业务描述"`
	Content    string  `gorm:"column:content;type:text;not null;comment:内容"`
	JsonSchema string  `gorm:"column:json_schema;type:text;not null;comment:json内容约束"`
	// 公共字段
	BaseTime
}

func (c JsonDynamicConfig) TableName() string {
	return "cms_json_dynamic_config"
}

// BizType 业务类型
type BizType string

const (
	IntegrationConsumeSetting BizType = "ums_integration_consume_setting" // 积分消费设置
)

// UmsIntegrationConsumeSetting 积分消费设置
type UmsIntegrationConsumeSetting struct {
	DeductionPerAmount uint32 // 每一元需要抵扣的积分数量
	MaxPercentPerOrder uint32 // 每笔订单最高抵用百分比
	UseUnit            uint32 // 每次使用积分最小单位100
	CouponStatus       uint8  // 是否可以和优惠券同用；0->不可以；1->可以
}
