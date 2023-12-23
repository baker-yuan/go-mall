package assembler

import (
	"github.com/baker-yuan/go-mall/application/portal/internal/entity"
	"github.com/baker-yuan/go-mall/common/util"
	pb "github.com/baker-yuan/go-mall/proto/mall"
)

func CartPromotionItemToModel(cartPromotionItems []*entity.CartPromotionItem) []*pb.CartItemListPromotion {
	res := make([]*pb.CartItemListPromotion, 0)

	for _, cartPromotionItem := range cartPromotionItems {
		item := &pb.CartItemListPromotion{
			// 一、订单信息
			// 用户信息
			MemberId: cartPromotionItem.MemberID, // 会员id
			// 商品信息
			ProductId:         cartPromotionItem.ProductID,
			ProductName:       cartPromotionItem.ProductName,
			ProductPic:        util.ImgUtils.GetFullUrl(cartPromotionItem.ProductPic),
			ProductSubTitle:   cartPromotionItem.ProductSubTitle,
			ProductSn:         cartPromotionItem.ProductSN,
			ProductBrand:      cartPromotionItem.ProductBrand,
			ProductCategoryId: cartPromotionItem.ProductCategoryID,
			// 商品sku
			ProductSkuId:   cartPromotionItem.ProductSkuID,
			ProductSkuCode: cartPromotionItem.ProductSkuCode,
			ProductAttr:    cartPromotionItem.ProductAttr,
			// 价格数量
			Price:    util.DecimalUtils.TrimTrailingZeros(cartPromotionItem.Price),
			Quantity: cartPromotionItem.Quantity,
			// 冗余字段
			MemberNickname: cartPromotionItem.MemberNickname,

			// 二、扩展字段
			PromotionMessage: cartPromotionItem.PromotionMessage,
			ReduceAmount:     util.DecimalUtils.TrimTrailingZeros(cartPromotionItem.ReduceAmount),
			RealStock:        cartPromotionItem.RealStock,
			Integration:      cartPromotionItem.Integration,
			Growth:           cartPromotionItem.Growth,
		}

		res = append(res, item)
	}

	return res
}
