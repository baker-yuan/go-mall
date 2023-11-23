package entity

// Product 商品信息表
// 商品信息主要包括四部分：商品的基本信息、商品的促销信息、商品的属性信息、商品的关联，商品表是整个商品的基本信息部分。
type Product struct {
	ID                         uint64  `gorm:"type:bigint;primary_key;auto_increment;comment:主键"`
	BrandID                    uint64  `gorm:"type:bigint;unsigned;not null;default:0;comment:品牌id"`   // pms_brand#id
	ProductCategoryID          uint64  `gorm:"type:bigint;unsigned;not null;default:0;comment:品牌分类id"` // pms_product_category#id
	FeightTemplateID           uint64  `gorm:"type:bigint;unsigned;not null;default:0;comment:运费模版id"`
	ProductAttributeCategoryID uint64  `gorm:"type:bigint;unsigned;not null;default:0;comment:品牌属性分类id"`
	Name                       string  `gorm:"type:varchar(64);not null;default:'';comment:商品名称"`
	Pic                        string  `gorm:"type:varchar(255);not null;default:'';comment:图片"`
	ProductSN                  string  `gorm:"type:varchar(64);not null;default:'';comment:货号"`
	DeleteStatus               uint8   `gorm:"type:tinyint(4);unsigned;not null;default:0;comment:删除状态：0->未删除；1->已删除"`
	PublishStatus              uint8   `gorm:"type:tinyint(4);unsigned;not null;default:0;comment:上架状态：0->下架；1->上架"`
	NewStatus                  uint8   `gorm:"type:tinyint(4);unsigned;not null;default:0;comment:新品状态:0->不是新品；1->新品"`
	RecommandStatus            uint8   `gorm:"type:tinyint(4);unsigned;not null;default:0;comment:推荐状态；0->不推荐；1->推荐"`
	VerifyStatus               uint8   `gorm:"type:tinyint(4);unsigned;not null;default:0;comment:审核状态：0->未审核；1->审核通过"`
	Sort                       uint32  `gorm:"type:int(10);unsigned;not null;default:0;comment:排序"`
	Sale                       uint32  `gorm:"type:int(10);unsigned;not null;default:0;comment:销量"`
	Price                      float64 `gorm:"type:decimal(10,2);not null;default:0.00;comment:价格"`
	PromotionPrice             float64 `gorm:"type:decimal(10,2);not null;default:0.00;comment:促销价格"`
	GiftGrowth                 uint32  `gorm:"type:int(10);unsigned;not null;default:0;comment:赠送的成长值"`
	GiftPoint                  uint32  `gorm:"type:int(10);unsigned;not null;default:0;comment:赠送的积分"`
	UsePointLimit              uint32  `gorm:"type:int(10);unsigned;not null;default:0;comment:限制使用的积分数"`
	SubTitle                   string  `gorm:"type:varchar(255);not null;default:'';comment:副标题"`
	Description                string  `gorm:"type:text;not null;comment:商品描述"`
	OriginalPrice              float64 `gorm:"type:decimal(10,2);not null;default:0.00;comment:市场价"`
	Stock                      uint32  `gorm:"type:int(10);unsigned;not null;default:0;comment:库存"`
	LowStock                   uint32  `gorm:"type:int(10);unsigned;not null;default:0;comment:库存预警值"`
	Unit                       string  `gorm:"type:varchar(16);not null;default:'';comment:单位"`
	Weight                     float64 `gorm:"type:decimal(10,2);not null;default:0.00;comment:商品重量，默认为克"`
	PreviewStatus              uint8   `gorm:"type:tinyint(4);unsigned;not null;default:0;comment:是否为预告商品：0->不是；1->是"`
	ServiceIDs                 string  `gorm:"type:varchar(64);not null;default:'';comment:以逗号分割的产品服务：1->无忧退货；2->快速退款；3->免费包邮"`
	Keywords                   string  `gorm:"type:varchar(255);not null;default:'';comment:关键字"`
	Note                       string  `gorm:"type:varchar(255);not null;default:'';comment:备注"`
	AlbumPics                  string  `gorm:"type:varchar(255);not null;default:'';comment:画册图片，连产品图片限制为5张，以逗号分割"`
	DetailTitle                string  `gorm:"type:varchar(255);not null;default:'';comment:详情标题"`
	DetailDesc                 string  `gorm:"type:text;not null;comment:详情描述"`
	DetailHTML                 string  `gorm:"type:text;not null;comment:产品详情网页内容"`
	DetailMobileHTML           string  `gorm:"type:text;not null;comment:移动端网页详情"`
	PromotionStartTime         uint32  `gorm:"type:int(10);unsigned;not null;default:0;comment:促销开始时间"`
	PromotionEndTime           uint32  `gorm:"type:int(10);unsigned;not null;default:0;comment:促销结束时间"`
	PromotionPerLimit          uint32  `gorm:"type:int(10);unsigned;not null;default:0;comment:活动限购数量"`
	PromotionType              uint8   `gorm:"type:tinyint(4);unsigned;not null;default:0;comment:促销类型：0->没有促销使用原价;1->使用促销价；2->使用会员价；3->使用阶梯价格；4->使用满减价格；5->限时购"`

	// 冗余字段
	BrandName           string `gorm:"type:varchar(255);comment:品牌名称" json:"brand_name"`
	ProductCategoryName string `gorm:"type:varchar(255);comment:商品分类名称" json:"product_category_name"`
	// 公共字段
	BaseTime
}

func (c Product) TableName() string {
	return "pms_product"
}
