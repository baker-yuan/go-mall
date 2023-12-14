package usecase

import (
	"context"

	"github.com/baker-yuan/go-mall/application/portal/internal/usecase/assembler"
	pb "github.com/baker-yuan/go-mall/proto/mall"
)

// CartItemUseCase 购物车表管理Service实现类
type CartItemUseCase struct {
	cartItemRepo ICartItemRepo // 操作购物车表
}

// NewCartItem 创建购物车表管理Service实现类
func NewCartItem(cartItemRepo ICartItemRepo) *CartItemUseCase {
	return &CartItemUseCase{
		cartItemRepo: cartItemRepo,
	}
}

// CartItemAdd 添加商品到购物车
func (c CartItemUseCase) CartItemAdd(ctx context.Context, req *pb.CartItemAddReq) (*pb.CartItemAddRsp, error) {
	//TODO implement me
	panic("implement me")
}

// CartItemList 获取当前会员的购物车列表
func (c CartItemUseCase) CartItemList(ctx context.Context, memberID uint64) ([]*pb.CartItem, error) {
	cartItems, err := c.cartItemRepo.GetEffectCartItemByMemberID(ctx, memberID)
	if err != nil {
		return nil, err
	}
	res := make([]*pb.CartItem, 0)
	for _, cartItem := range cartItems {
		res = append(res, assembler.CartItemEntityToModel(cartItem))
	}
	return res, nil
}

// CartItemListPromotion 获取当前会员的购物车列表,包括促销信息
func (c CartItemUseCase) CartItemListPromotion(ctx context.Context, req *pb.CartItemListPromotionReq) (*pb.CartItemListPromotionRsp, error) {
	//TODO implement me
	panic("implement me")
}

// CartItemUpdateQuantity 修改购物车中指定商品的数量
func (c CartItemUseCase) CartItemUpdateQuantity(ctx context.Context, req *pb.CartItemUpdateQuantityReq) (*pb.CartItemUpdateQuantityRsp, error) {
	//TODO implement me
	panic("implement me")
}

// CartItemGetCartProduct 获取购物车中指定商品的规格,用于重选规格
func (c CartItemUseCase) CartItemGetCartProduct(ctx context.Context, req *pb.CartItemGetCartProductReq) (*pb.CartItemGetCartProductRsp, error) {
	//TODO implement me
	panic("implement me")
}

// CartItemUpdateAttr 修改购物车中商品的规格
func (c CartItemUseCase) CartItemUpdateAttr(ctx context.Context, req *pb.CartItemUpdateAttrReq) (*pb.CartItemUpdateAttrRsp, error) {
	//TODO implement me
	panic("implement me")
}

// CartItemDelete 删除购物车中的指定商品
func (c CartItemUseCase) CartItemDelete(ctx context.Context, req *pb.CartItemDeleteReq) (*pb.CartItemDeleteRsp, error) {
	//TODO implement me
	panic("implement me")
}

// CartItemClear 清空当前会员的购物车
func (c CartItemUseCase) CartItemClear(ctx context.Context, req *pb.CartItemClearReq) (*pb.CartItemClearRsp, error) {
	//TODO implement me
	panic("implement me")
}
