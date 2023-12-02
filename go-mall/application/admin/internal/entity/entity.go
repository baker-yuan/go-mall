package entity

import (
	"fmt"
	"time"

	"gorm.io/gorm"
)

type BaseModel struct {
	ID        uint64 `gorm:"primarykey"`
	CreatedAt time.Time
	UpdatedAt time.Time
	DeletedAt gorm.DeletedAt `gorm:"index"`
}

type BaseTime struct {
	CreatedAt uint32 `gorm:"type:int(10);unsigned;not null;default:0;comment:创建时间"` // 使用时间戳秒数填充创建时间
	UpdatedAt uint32 `gorm:"type:int(10);unsigned;not null;default:0;comment:修改时间"` // 使用时间戳秒数填充修改时间
}

func (bt *BaseTime) BeforeCreate(tx *gorm.DB) (err error) {
	timestamp := uint32(time.Now().Unix())
	bt.CreatedAt = timestamp
	bt.UpdatedAt = timestamp
	return
}

func (bt *BaseTime) BeforeUpdate(tx *gorm.DB) (err error) {
	bt.UpdatedAt = uint32(time.Now().Unix())
	return
}

func Init(db *gorm.DB) error {
	schemas := []tableSchema{
		// cms
		{
			TableName: "商品分类表",
			StructPtr: &ProductCategory{},
		},
		{
			TableName: "商品分类和属性的关系表",
			StructPtr: &ProductCategoryAttributeRelation{},
		},
		{
			TableName: "商品品牌表",
			StructPtr: &Brand{},
		},
		{
			TableName: "商品属性分类表",
			StructPtr: &ProductAttributeCategory{},
		},
		{
			TableName: "商品属性表",
			StructPtr: &ProductAttribute{},
		},
		{
			TableName: "商品属性值表",
			StructPtr: &ProductAttributeValue{},
		},
		{
			TableName: "商品表",
			StructPtr: &Product{},
		},
		{
			TableName: "商品SKU表",
			StructPtr: &SkuStock{},
		},
		{
			TableName: "商品阶梯价格表",
			StructPtr: &ProductLadder{},
		},
		{
			TableName: "商品满减表",
			StructPtr: &ProductFullReduction{},
		},
		{
			TableName: "商品会员价格表",
			StructPtr: &MemberPrice{},
		},
		{
			TableName: "优选专区",
			StructPtr: &PrefrenceArea{},
		},
		{
			TableName: "优选专区和产品关系表",
			StructPtr: &PrefrenceAreaProductRelation{},
		},
		{
			TableName: "专题表",
			StructPtr: &Subject{},
		},
		{
			TableName: "专题商品关系表",
			StructPtr: &SubjectProductRelation{},
		},
		{
			TableName: "商品评价表",
			StructPtr: &Comment{},
		},
		{
			TableName: "商品评价回复表",
			StructPtr: &CommentReplay{},
		},
		{
			TableName: "商品审核记录表",
			StructPtr: &ProductVertifyRecord{},
		},
		{
			TableName: "商品操作记录表",
			StructPtr: &ProductOperateLog{},
		},

		// oms
		{
			TableName: "订单表",
			StructPtr: &Order{},
		},
		{
			TableName: "订单商品信息表",
			StructPtr: &OrderItem{},
		},
		{
			TableName: "订单操作历史记录表",
			StructPtr: &OrderOperateHistory{},
		},
		{
			TableName: "订单设置表",
			StructPtr: &OrderSetting{},
		},
		{
			TableName: "退货原因表",
			StructPtr: &OrderReturnReason{},
		},

		// sms
		//{
		//	TableName: "限时购表",
		//	StructPtr: &FlashPromotion{},
		//},
		//{
		//	TableName: "限时购场次表",
		//	StructPtr: &FlashPromotionSession{},
		//},
		//{
		//	TableName: "限时购与商品关系表",
		//	StructPtr: &FlashPromotionProductRelation{},
		//},
		//{
		//	TableName: "限时购通知记录表",
		//	StructPtr: &FlashPromotionLog{},
		//},
		//{
		//	TableName: "优惠券表",
		//	StructPtr: &Coupon{},
		//},
		//{
		//	TableName: "优惠券使用历史表",
		//	StructPtr: &CouponHistory{},
		//},
		//{
		//	TableName: "优惠券和商品的关系表",
		//	StructPtr: &SmsCouponProductRelation{},
		//},
		//{
		//	TableName: "优惠券和商品分类关系表",
		//	StructPtr: &CouponProductCategoryRelation{},
		//},
		//{
		//	TableName: "首页品牌推荐表",
		//	StructPtr: &HomeBrand{},
		//},
		//{
		//	TableName: "新品推荐商品表",
		//	StructPtr: &HomeNewProduct{},
		//},
		//{
		//	TableName: "人气推荐商品表",
		//	StructPtr: &HomeRecommendProduct{},
		//},
		//{
		//	TableName: "首页专题推荐表",
		//	StructPtr: &HomeRecommendSubject{},
		//},
		//{
		//	TableName: "首页轮播广告表",
		//	StructPtr: &HomeAdvertise{},
		//},
		//{
		//	TableName: "会员表",
		//	StructPtr: &Member{},
		//},
		//{
		//	TableName: "会员收货地址表",
		//	StructPtr: &MemberReceiveAddress{},
		//},
	}
	if err := autoMigrate(db, schemas); err != nil {
		return err
	}

	return nil
}

type tableSchema struct {
	TableName string      // 表名
	StructPtr interface{} // 结构体指针
}

func autoMigrate(db *gorm.DB, schemas []tableSchema) error {
	for _, schema := range schemas {
		if err := db.Set("gorm:table_options", fmt.Sprintf("COMMENT='%s'", schema.TableName)).
			AutoMigrate(schema.StructPtr); err != nil {
			return err
		}
	}
	return nil
}
