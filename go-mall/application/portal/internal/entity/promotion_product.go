package entity

import g_entity "github.com/baker-yuan/go-mall/common/entity"

// PromotionProduct 促销商品信息，包括sku、打折优惠、满减优惠
type PromotionProduct struct {
	*g_entity.Product
	SkuStocks             g_entity.SkuStocks             // 商品库存信息
	ProductLadders        g_entity.ProductLadders        // 商品打折信息
	ProductFullReductions g_entity.ProductFullReductions // 商品满减信息
}
