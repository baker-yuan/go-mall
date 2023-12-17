package usecase

import (
	"context"
	"time"

	"github.com/baker-yuan/go-mall/application/portal/internal/usecase/assembler"
	pb "github.com/baker-yuan/go-mall/proto/mall"
)

// CartItemUseCase 购物车表管理Service实现类
type CartItemUseCase struct {
	cartItemRepo ICartItemRepo // 操作购物车表
	memberRepo   IMemberRepo   // 操作会员
	productRepo  IProductRepo  // 操作商品
	brandRepo    IBrandRepo    // 操作品牌
}

// NewCartItem 创建购物车表管理Service实现类
func NewCartItem(cartItemRepo ICartItemRepo, memberRepo IMemberRepo, productRepo IProductRepo, brandRepo IBrandRepo) *CartItemUseCase {
	return &CartItemUseCase{
		cartItemRepo: cartItemRepo,
		memberRepo:   memberRepo,
		productRepo:  productRepo,
		brandRepo:    brandRepo,
	}
}

// CartItemAdd 查询购物车中是否包含该商品，有增加数量，无添加到购物车
func (c CartItemUseCase) CartItemAdd(ctx context.Context, memberID uint64, req *pb.CartItemAddReq) error {
	// 类型转换
	cartItem := assembler.CartItemModelToEntity(req)
	cartItem.MemberID = memberID
	cartItem.DeleteStatus = 0

	// 获取商品信息
	product, err := c.productRepo.GetByID(ctx, req.GetProductId())
	if err != nil {
		return err
	}
	// 获取品牌信息
	brand, err := c.brandRepo.GetByID(ctx, product.BrandID)
	if err != nil {
		return err
	}

	cartItem.ProductName = product.Name
	cartItem.ProductPic = product.Pic
	cartItem.ProductSubTitle = product.SubTitle
	cartItem.ProductSN = product.ProductSN
	cartItem.ProductCategoryID = product.ProductCategoryID
	cartItem.ProductBrand = brand.Name
	cartItem.Price = product.Price

	// 查询db是否存在
	existCartItem, err := c.cartItemRepo.GetCartItem(ctx, memberID, req.GetProductId(), req.GetProductSkuId())
	if err != nil {
		return err
	}
	if existCartItem == nil {
		// 添加购物车
		cartItem.CreateDate = uint32(time.Now().Unix())
		cartItem.ModifyDate = uint32(time.Now().Unix())
		return c.cartItemRepo.Create(ctx, cartItem)
	} else {
		existCartItem.DeleteStatus = 0
		existCartItem.ModifyDate = uint32(time.Now().Unix())
		// 数量++
		existCartItem.Quantity = existCartItem.Quantity + cartItem.Quantity
		return c.cartItemRepo.Update(ctx, existCartItem)
	}
}

// CartItemList 获取当前会员的购物车列表
func (c CartItemUseCase) CartItemList(ctx context.Context, memberID uint64) ([]*pb.CartItem, error) {
	cartItems, err := c.cartItemRepo.GetEffectCartItemByMemberID(ctx, memberID)
	if err != nil {
		return nil, err
	}
	// 补充会员名称
	members, err := c.memberRepo.GetByIDs(ctx, cartItems.GetMemberIDs())
	if err != nil {
		return nil, err
	}
	// 组装数据
	res := make([]*pb.CartItem, 0)
	for _, cartItem := range cartItems {
		res = append(res, assembler.CartItemEntityToModel(cartItem, members.GetMap()))
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
