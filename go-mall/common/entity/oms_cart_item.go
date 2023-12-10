package entity

// CartItem 购物车表
type CartItem struct {
	ID                uint64  `gorm:"column:id;type:bigint;primary_key;auto_increment;comment:主键"`
	MemberID          uint64  `gorm:"column:member_id;type:bigint;unsigned;not null;default:0;comment:会员id"`
	ProductCategoryID uint64  `gorm:"column:product_category_id;type:bigint;unsigned;not null;default:0;comment:商品分类"`
	ProductBrand      string  `gorm:"column:product_brand;type:varchar(200);not null;default:'';comment:商品品牌"`
	ProductID         uint64  `gorm:"column:product_id;type:bigint;unsigned;not null;default:0;comment:商品id"`
	ProductSkuID      uint64  `gorm:"column:product_sku_id;type:bigint;unsigned;not null;default:0;comment:商品sku id"`
	ProductPic        string  `gorm:"column:product_pic;type:varchar(1000);not null;default:'';comment:商品主图"`
	ProductSubTitle   string  `gorm:"column:product_sub_title;type:varchar(500);not null;default:'';comment:商品副标题（卖点）"`
	ProductSkuCode    string  `gorm:"column:product_sku_code;type:varchar(200);not null;default:'';comment:商品sku条码"`
	ProductSN         string  `gorm:"column:product_sn;type:varchar(200);not null;default:'';comment:商品货号"`
	ProductAttr       string  `gorm:"column:product_attr;type:varchar(500);not null;default:'';comment:商品销售属性"`
	Quantity          uint32  `gorm:"column:quantity;type:int(10);unsigned;not null;default:0;comment:购买数量"`
	Price             float64 `gorm:"column:price;type:decimal(10,2);not null;default:0.00;comment:添加到购物车的价格"`
	//
	DeleteStatus uint8  `gorm:"column:delete_status;type:tinyint(4);unsigned;not null;default:0;comment:是否删除"`
	CreateDate   uint32 `gorm:"column:create_date;type:int(10);unsigned;not null;default:0;comment:创建时间"`
	ModifyDate   uint32 `gorm:"column:modify_date;type:int(10);unsigned;not null;default:0;comment:修改时间"`
	// 冗余字段
	ProductName    string `gorm:"column:product_name;type:varchar(500);not null;default:'';comment:商品名称"`
	MemberNickname string `gorm:"column:member_nickname;type:varchar(500);not null;default:'';comment:会员昵称"`
	// 公共字段
	BaseTime
}

func (c CartItem) TableName() string {
	return "oms_cart_item"
}
