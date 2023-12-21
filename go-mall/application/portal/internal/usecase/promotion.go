package usecase

import (
	"context"

	portal_entity "github.com/baker-yuan/go-mall/application/portal/internal/entity"
	"github.com/baker-yuan/go-mall/common/entity"
)

// PromotionUseCase 促销管理Service实现类
type PromotionUseCase struct {
	productRepo              IProductRepo              // 操作商品
	skuStockRepo             ISkuStockRepo             // 操作sku库存
	productLadderRepo        IProductLadderRepo        // 操作商品阶梯价格
	productFullReductionRepo IProductFullReductionRepo // 操作商品满减
}

// NewPromotion 创建促销管理管理Service实现类
func NewPromotion() *PromotionUseCase {
	return &PromotionUseCase{}
}

// CalcCartPromotion 计算购物车中的促销活动信息
// cartItems 购物车
func (c PromotionUseCase) CalcCartPromotion(ctx context.Context, cartItems entity.CartItems) (portal_entity.CartPromotionItems, error) {
	// 1.先根据productId对CartItem进行分组，以spu为单位进行计算优惠
	// key=商品id value=购物车集合
	productCartMap := cartItems.GroupCartItemBySpu()

	// 2.查询所有商品的优惠相关信息
	promotionProducts, err := c.getPromotionProductList(ctx, cartItems)
	if err != nil {
		return nil, err
	}
	// 3.根据商品促销类型计算商品促销优惠价格
	//for kv := range productCartMap.Range() {
	//	productID := kv.Key
	//	v := kv.Value
	//}

	return nil, nil
}

// 查询所有商品的优惠相关信息
func (c PromotionUseCase) getPromotionProductList(ctx context.Context, cartItems entity.CartItems) ([]*portal_entity.PromotionProduct, error) {
	productIDs := cartItems.GetProductIDs()
	// 查询产品信息
	products, err := c.productRepo.GetByIDs(ctx, productIDs)
	if err != nil {
		return nil, err
	}
	// 查询SKU库存信息
	skuStocks, err := c.skuStockRepo.GetByProductIDs(ctx, productIDs)
	if err != nil {
		return nil, err
	}

	// 查询产品阶梯价格信息
	productLadders, err := c.productLadderRepo.GetByProductIDs(ctx, productIDs)
	if err != nil {
		return nil, err
	}

	// 查询产品满减信息
	fullReductions, err := c.productFullReductionRepo.GetByProductIDs(ctx, productIDs)
	if err != nil {
		return nil, err
	}

	res := make([]*portal_entity.PromotionProduct, 0)

	// 创建一个以 product_id 为键的 map
	promotionProducts := make(map[uint64]*portal_entity.PromotionProduct)
	// 产品
	for _, product := range products {
		promotionProducts[product.ID] = &portal_entity.PromotionProduct{Product: product}
	}
	// 将SKU库存信息添加到对应的产品中
	for _, stock := range skuStocks {
		if product, ok := promotionProducts[stock.ProductID]; ok {
			product.SkuStocks = append(product.SkuStocks, stock)
			promotionProducts[stock.ProductID] = product
		}
	}
	// 将产品阶梯价格信息添加到对应的产品中
	for _, ladder := range productLadders {
		if product, ok := promotionProducts[ladder.ProductID]; ok {
			product.ProductLadders = append(product.ProductLadders, ladder)
			promotionProducts[ladder.ProductID] = product
		}
	}

	// 将产品满减信息添加到对应的产品中
	for _, reduction := range fullReductions {
		if product, ok := promotionProducts[reduction.ProductID]; ok {
			product.ProductFullReductions = append(product.ProductFullReductions, reduction)
			promotionProducts[reduction.ProductID] = product
		}
	}

	for _, v := range promotionProducts {
		res = append(res, v)
	}
	return res, nil
}
