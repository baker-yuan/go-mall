package assembler

import (
	"github.com/baker-yuan/go-mall/common/entity"
	"github.com/baker-yuan/go-mall/common/util"
	pb "github.com/baker-yuan/go-mall/proto/mall"
)

// CartItemModelToEntity pb转entity
func CartItemModelToEntity(cartItemPb *pb.CartItemAddReq) *entity.CartItem {
	if cartItemPb == nil {
		return nil
	}
	return &entity.CartItem{
		// 商品信息
		ProductID: cartItemPb.ProductId,
		// 商品属性
		ProductAttr:    cartItemPb.ProductAttr,
		ProductSkuID:   cartItemPb.ProductSkuId,
		ProductSkuCode: cartItemPb.ProductSkuCode,
		// 价格数量
		Quantity: cartItemPb.Quantity,
	}
}

// CartItemEntityToModel entity转pb
func CartItemEntityToModel(cartItem *entity.CartItem, memberMap map[uint64]*entity.Member) *pb.CartItem {
	if cartItem == nil {
		return nil
	}
	res := &pb.CartItem{
		// 主键
		Id: cartItem.ID,
		// 用户信息
		MemberId: cartItem.MemberID,
		// 商品信息
		ProductId:         cartItem.ProductID,
		ProductName:       cartItem.ProductName,
		ProductPic:        util.ImgUtils.GetFullUrl(cartItem.ProductPic),
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
	}
	if member, exist := memberMap[cartItem.MemberID]; exist {
		res.MemberNickname = member.Nickname
	}
	return res
}
