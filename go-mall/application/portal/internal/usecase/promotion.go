package usecase

import (
	"context"

	portal_entity "github.com/baker-yuan/go-mall/application/portal/internal/entity"
	"github.com/baker-yuan/go-mall/common/entity"
)

// PromotionUseCase 促销管理Service实现类
type PromotionUseCase struct {
}

// NewPromotion 创建促销管理管理Service实现类
func NewPromotion() *OrderUseCase {
	return &OrderUseCase{}
}

// CalcCartPromotion 计算购物车中的促销活动信息
// cartItems 购物车
func (c OrderUseCase) CalcCartPromotion(ctx context.Context, cartItems entity.CartItems) (portal_entity.CartPromotionItems, error) {
	return nil, nil
}
