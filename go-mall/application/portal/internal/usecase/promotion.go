package usecase

import (
	"context"
	"fmt"
	"math/big"
	"sort"
	"strings"

	portal_entity "github.com/baker-yuan/go-mall/application/portal/internal/entity"
	"github.com/baker-yuan/go-mall/common/entity"
	"github.com/baker-yuan/go-mall/common/util"
	"github.com/baker-yuan/go-mall/common/util/decimal"
	pb "github.com/baker-yuan/go-mall/proto/mall"
)

// PromotionUseCase 促销管理Service实现类
type PromotionUseCase struct {
	productRepo              IProductRepo              // 操作商品
	skuStockRepo             ISkuStockRepo             // 操作sku库存
	productLadderRepo        IProductLadderRepo        // 操作商品阶梯价格
	productFullReductionRepo IProductFullReductionRepo // 操作商品满减
}

// NewPromotion 创建促销管理管理Service实现类
func NewPromotion(
	productRepo IProductRepo,
	skuStockRepo ISkuStockRepo,
	productLadderRepo IProductLadderRepo,
	productFullReductionRepo IProductFullReductionRepo,
) *PromotionUseCase {
	return &PromotionUseCase{
		productRepo:              productRepo,
		skuStockRepo:             skuStockRepo,
		productLadderRepo:        productLadderRepo,
		productFullReductionRepo: productFullReductionRepo,
	}
}

// CalcCartPromotion 计算购物车中的促销活动信息
// cartItems 购物车
func (c PromotionUseCase) CalcCartPromotion(ctx context.Context, cartItems entity.CartItems) (portal_entity.CartPromotionItems, error) {
	cartPromotionItems := make([]*portal_entity.CartPromotionItem, 0)

	decimalUtils := util.DecimalUtils

	// 1.先根据productId对CartItem进行分组，以spu为单位进行计算优惠
	// key=商品id value=购物车集合
	productCartMap := cartItems.GroupCartItemBySpu()

	// 2.查询所有商品的优惠相关信息
	promotionProducts, err := c.getPromotionProductList(ctx, cartItems)
	if err != nil {
		return nil, err
	}

	// 3.根据商品促销类型计算商品促销优惠价格
	for productID, cartItemList := range productCartMap {
		promotionProduct := promotionProducts.GetByProductID(productID)
		if promotionProduct == nil {
			continue
		}
		// 促销类型：0->没有促销使用原价;1->使用促销价；2->使用会员价；3->使用阶梯价格；4->使用满减价格；5->限时购
		promotionType := promotionProduct.PromotionType
		if promotionType == pb.PromotionType_PROMOTION_TYPE_PROMOTIONAL_PRICE {
			// 单品促销
			for _, cartItem := range cartItemList {
				cartPromotionItem, _ := util.CopyProperties[*portal_entity.CartPromotionItem](cartItem)
				cartPromotionItem.PromotionMessage = "单品促销"
				// 商品原价-促销价
				skuStock := c.getOriginalPrice(promotionProduct, cartItem.ProductSkuID)
				originalPrice := skuStock.Price
				// 单品促销使用原价
				cartPromotionItem.Price = originalPrice
				reduceAmount, _ := util.DecimalUtils.NewBigDecimal(originalPrice).
					Subtract(util.DecimalUtils.NewBigDecimal(skuStock.PromotionPrice)).
					ToString()
				cartPromotionItem.ReduceAmount = reduceAmount
				cartPromotionItem.RealStock = skuStock.Stock - skuStock.LockStock
				cartPromotionItem.Integration = promotionProduct.GiftPoint
				cartPromotionItem.Growth = promotionProduct.GiftGrowth
				cartPromotionItems = append(cartPromotionItems, cartPromotionItem)
			}
		} else if promotionType == pb.PromotionType_PROMOTION_TYPE_TIERED_PRICE {
			// 打折优惠
			count := cartItemList.GetCartItemCount()
			ladder := c.getProductLadder(count, promotionProduct.ProductLadders)
			if ladder != nil {
				for _, cartItem := range cartItemList {
					cartPromotionItem, _ := util.CopyProperties[*portal_entity.CartPromotionItem](cartItem)
					message := c.getLadderPromotionMessage(ladder)
					cartPromotionItem.PromotionMessage = message
					// 商品原价-折扣*商品原价
					skuStock := c.getOriginalPrice(promotionProduct, cartItem.ProductSkuID)
					cartPromotionItem.Price = skuStock.Price
					originalPrice := util.DecimalUtils.NewBigDecimal(skuStock.Price)
					reduceAmount, _ := originalPrice.Subtract(util.DecimalUtils.NewBigDecimal(ladder.Discount).Multiply(originalPrice)).ToString()
					cartPromotionItem.ReduceAmount = reduceAmount
					cartPromotionItem.RealStock = skuStock.Stock - skuStock.LockStock
					cartPromotionItem.Integration = promotionProduct.GiftPoint
					cartPromotionItem.Growth = promotionProduct.GiftGrowth
					cartPromotionItems = append(cartPromotionItems, cartPromotionItem)
				}
			} else {
				cartPromotionItems = append(cartPromotionItems, c.handleNoReduce(cartItemList, promotionProduct)...)
			}
		} else if promotionType == pb.PromotionType_PROMOTION_TYPE_FULL_REDUCTION_PRICE {
			// 满减
			totalAmount, _ := c.getCartItemAmount(cartItemList, promotionProducts)
			fullReduction, _ := c.getProductFullReduction(totalAmount, promotionProduct.ProductFullReductions)
			if fullReduction != nil {
				for _, cartItem := range cartItemList {
					cartPromotionItem, _ := util.CopyProperties[*portal_entity.CartPromotionItem](cartItem)
					message := c.getFullReductionPromotionMessage(fullReduction)
					cartPromotionItem.PromotionMessage = message
					// (商品原价/总价)*满减金额
					skuStock := c.getOriginalPrice(promotionProduct, cartItem.ProductSkuID)
					originalPrice := decimalUtils.NewBigDecimal(skuStock.Price)
					reduceAmount, _ := originalPrice.Divide(totalAmount).Multiply(decimalUtils.NewBigDecimal(fullReduction.ReducePrice)).ToString()
					cartPromotionItem.ReduceAmount = reduceAmount
					cartPromotionItem.RealStock = skuStock.Stock - skuStock.LockStock
					cartPromotionItem.Integration = promotionProduct.GiftPoint
					cartPromotionItem.Growth = promotionProduct.GiftGrowth
					cartPromotionItems = append(cartPromotionItems, cartPromotionItem)
				}
			} else {
				cartPromotionItems = append(cartPromotionItems, c.handleNoReduce(cartItemList, promotionProduct)...)
			}
		} else {
			// 无优惠
			cartPromotionItems = append(cartPromotionItems, c.handleNoReduce(cartItemList, promotionProduct)...)
		}
	}

	return cartPromotionItems, nil
}

// handleNoReduce 对没满足优惠条件的商品进行处理
func (c PromotionUseCase) handleNoReduce(cartItems entity.CartItems, promotionProduct *portal_entity.PromotionProduct) []*portal_entity.CartPromotionItem {
	cartPromotionItems := make([]*portal_entity.CartPromotionItem, 0)
	for _, cartItem := range cartItems {
		cartPromotionItem, _ := util.CopyProperties[*portal_entity.CartPromotionItem](cartItem)
		cartPromotionItem.PromotionMessage = "无优惠"
		cartPromotionItem.ReduceAmount = "0.00"
		skuStock := c.getOriginalPrice(promotionProduct, cartItem.ProductSkuID)
		if skuStock != nil {
			cartPromotionItem.RealStock = skuStock.Stock - skuStock.LockStock
		}
		cartPromotionItem.Integration = promotionProduct.GiftPoint
		cartPromotionItem.Growth = promotionProduct.GiftGrowth
		cartPromotionItems = append(cartPromotionItems, cartPromotionItem)
	}
	return cartPromotionItems
}

// getProductLadder 根据购买商品数量获取满足条件的打折优惠策略
func (c PromotionUseCase) getProductLadder(count uint32, productLadders []*entity.ProductLadder) *entity.ProductLadder {
	// 按数量从大到小排序（数量越多优惠越大，找到数量最多的那个配置）
	sort.Slice(productLadders, func(i, j int) bool {
		return productLadders[j].Count < productLadders[i].Count
	})
	// 遍历排序后的列表，找到满足条件的优惠策略
	for _, productLadder := range productLadders {
		if count >= productLadder.Count {
			return productLadder
		}
	}
	return nil
}

// 查询所有商品的优惠相关信息
func (c PromotionUseCase) getPromotionProductList(ctx context.Context, cartItems entity.CartItems) (portal_entity.PromotionProducts, error) {
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
	temp := make(map[uint64]*portal_entity.PromotionProduct)
	// 产品
	for _, product := range products {
		temp[product.ID] = &portal_entity.PromotionProduct{Product: product}
	}
	// 将SKU库存信息添加到对应的产品中
	for _, stock := range skuStocks {
		if product, ok := temp[stock.ProductID]; ok {
			product.SkuStocks = append(product.SkuStocks, stock)
		}
	}
	// 将产品阶梯价格信息添加到对应的产品中
	for _, ladder := range productLadders {
		if product, ok := temp[ladder.ProductID]; ok {
			product.ProductLadders = append(product.ProductLadders, ladder)
		}
	}

	// 将产品满减信息添加到对应的产品中
	for _, reduction := range fullReductions {
		if product, ok := temp[reduction.ProductID]; ok {
			product.ProductFullReductions = append(product.ProductFullReductions, reduction)
		}
	}

	for _, v := range temp {
		res = append(res, v)
	}
	return res, nil
}

// 获取商品的原价
func (c PromotionUseCase) getOriginalPrice(promotionProduct *portal_entity.PromotionProduct, productSkuID uint64) *entity.SkuStock {
	for _, skuStock := range promotionProduct.SkuStocks {
		if skuStock.ID == productSkuID {
			return skuStock
		}
	}
	return nil
}

// GetLadderPromotionMessage 获取打折优惠的促销信息
func (c PromotionUseCase) getLadderPromotionMessage(ladder *entity.ProductLadder) string {
	// 将字符串表示的折扣比例转换为折扣数（例如："0.8" -> "8折"）
	discount, ok := new(big.Float).SetString(ladder.Discount)
	if !ok {
		return ""
	}
	// 折扣乘以10得到折扣百分比
	discount.Mul(discount, big.NewFloat(10))
	// 将big.Float格式化为字符串，保留一位小数
	discountStr := discount.Text('f', 1)
	// 如果小数部分为0，则去除小数点和尾随的零
	if strings.HasSuffix(discountStr, ".0") {
		discountStr = strings.TrimSuffix(discountStr, ".0")
	}
	// 构建并返回促销信息字符串
	return fmt.Sprintf("打折优惠：满%d件，打%s折", ladder.Count, discountStr)
}

// getCartItemAmount 获取购物车中指定商品的总价
func (c PromotionUseCase) getCartItemAmount(cartItems entity.CartItems, promotionProducts portal_entity.PromotionProducts) (*decimal.BigDecimal, error) {
	decimalUtils := util.DecimalUtils
	amount := decimalUtils.NewBigDecimal("0.00")
	for _, cartItem := range cartItems {
		// 计算出商品原价
		promotionProduct := promotionProducts.GetByProductID(cartItem.ProductID)
		if promotionProduct == nil {
			continue // 如果找不到促销产品，跳过当前商品项
		}
		skuStock := c.getOriginalPrice(promotionProduct, cartItem.ProductSkuID)
		if skuStock == nil {
			continue // 如果找不到原始价格，跳过当前商品项
		}
		// 计算商品项总价
		itemTotal := decimalUtils.NewBigDecimal(skuStock.Price).Multiply(decimalUtils.NewBigDecimal(fmt.Sprintf("%d", cartItem.Quantity)))
		// 累加到总金额中
		amount = amount.Add(itemTotal)
	}
	return amount, nil
}

func (c PromotionUseCase) getProductFullReduction(totalAmount *decimal.BigDecimal, productFullReductions []*entity.ProductFullReduction) (*entity.ProductFullReduction, error) {
	decimalUtils := util.DecimalUtils
	// 按条件从高到低排序
	sort.Slice(productFullReductions, func(i, j int) bool {
		result, _ := decimalUtils.CompareDecimal(productFullReductions[j].FullPrice, productFullReductions[i].FullPrice)
		return result < 0
	})
	// 遍历排序后的列表，找到满足条件的满减优惠策略
	for _, fullReduction := range productFullReductions {
		result, err := totalAmount.Compare(decimalUtils.NewBigDecimal(fullReduction.FullPrice))
		if err != nil {
			return nil, err
		}
		if result >= 0 {
			return fullReduction, nil
		}
	}
	return nil, nil
}

// getFullReductionPromotionMessage 获取满减促销消息
func (c PromotionUseCase) getFullReductionPromotionMessage(fullReduction *entity.ProductFullReduction) string {
	return fmt.Sprintf("满减优惠：满%s元，减%s元",
		util.DecimalUtils.TrimTrailingZeros(fullReduction.FullPrice),
		util.DecimalUtils.TrimTrailingZeros(fullReduction.ReducePrice))
}
