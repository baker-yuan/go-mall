package assembler

import (
	"github.com/baker-yuan/go-mall/common/entity"
	"github.com/baker-yuan/go-mall/common/util"
	pb "github.com/baker-yuan/go-mall/proto/mall"
)

// CartItemEntityToModel entity转pb
func CartItemEntityToModel(cartItem *entity.CartItem) *pb.CartItem {
	if cartItem == nil {
		return nil
	}
	return &pb.CartItem{
		// 主键
		Id: cartItem.ID,
		// 用户信息
		MemberId: cartItem.MemberID,
		// 商品信息
		ProductId:         cartItem.ProductID,
		ProductName:       cartItem.ProductName,
		ProductPic:        util.GetFullUrl(cartItem.ProductPic),
		ProductSubTitle:   cartItem.ProductSubTitle,
		ProductSn:         cartItem.ProductSN,
		ProductBrand:      cartItem.ProductBrand,
		ProductCategoryId: cartItem.ProductCategoryID,
		// 商品属性
		ProductAttr:    cartItem.ProductAttr,
		ProductSkuId:   cartItem.ProductSkuID,
		ProductSkuCode: cartItem.ProductSkuCode,
		// 价格数量
		Price:    cartItem.Price,
		Quantity: cartItem.Quantity,
		// 状态
		CreateDate:   cartItem.CreateDate,
		ModifyDate:   cartItem.ModifyDate,
		DeleteStatus: uint32(cartItem.DeleteStatus),
		// 冗余字段
		MemberNickname: cartItem.MemberNickname,
	}
}
