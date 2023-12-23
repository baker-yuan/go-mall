package usecase

import (
	"context"
	"fmt"
	"time"

	portal_entity "github.com/baker-yuan/go-mall/application/portal/internal/entity"
	"github.com/baker-yuan/go-mall/common/util"
	pb "github.com/baker-yuan/go-mall/proto/mall"
)

// CouponUseCase 优惠券表管理Service实现类
type CouponUseCase struct {
	couponRepo                        ICouponRepo                        // 优惠券
	couponHistoryRepo                 ICouponHistoryRepo                 // 优惠券使用、领取历史
	couponProductCategoryRelationRepo ICouponProductCategoryRelationRepo // 优惠券和商品分类关系
	couponProductRelationRepo         ICouponProductRelationRepo         // 优惠券和商品的关系
}

// NewCoupon 创建优惠券表管理Service实现类
func NewCoupon(
	couponRepo ICouponRepo,
	couponHistoryRepo ICouponHistoryRepo,
	couponProductCategoryRelationRepo ICouponProductCategoryRelationRepo,
	couponProductRelationRepo ICouponProductRelationRepo,
) *CouponUseCase {
	return &CouponUseCase{
		couponRepo:                        couponRepo,
		couponHistoryRepo:                 couponHistoryRepo,
		couponProductCategoryRelationRepo: couponProductCategoryRelationRepo,
		couponProductRelationRepo:         couponProductRelationRepo,
	}
}

// CouponAdd 领取指定优惠券
func (s CouponUseCase) CouponAdd(ctx context.Context, req *pb.CouponAddReq) (*pb.CouponAddRsp, error) {
	//TODO implement me
	panic("implement me")
}

// CouponListHistory 获取会员优惠券历史列表
func (s CouponUseCase) CouponListHistory(ctx context.Context, req *pb.CouponListHistoryReq) (*pb.CouponListHistoryRsp, error) {
	//TODO implement me
	panic("implement me")
}

// CouponList 获取会员优惠券列表
func (s CouponUseCase) CouponList(ctx context.Context, req *pb.CouponListReq) (*pb.CouponListRsp, error) {
	//TODO implement me
	panic("implement me")
}

// CouponListCart 根据购物车信息获取可用优惠券
func (s CouponUseCase) CouponListCart(ctx context.Context, memberID uint64, cartPromotionItems []*pb.CartItemListPromotion, canUse bool) ([]*portal_entity.CouponHistoryDetail, error) {
	// 获取该用户所有优惠券
	allList, err := s.getDetailList(ctx, memberID)
	if err != nil {
		return nil, err
	}

	// 根据优惠券使用类型来判断优惠券是否可用
	enableList := make([]*portal_entity.CouponHistoryDetail, 0)
	disableList := make([]*portal_entity.CouponHistoryDetail, 0)

	now := time.Now()
	decimalUtils := util.DecimalUtils

	for _, couponHistoryDetail := range allList {
		// 条件
		useType := couponHistoryDetail.Coupon.UseType                      // 使用类型：0->全场通用；1->指定分类；2->指定商品
		minPoint := couponHistoryDetail.Coupon.MinPoint                    // 使用门槛；0表示无门槛
		endTime := time.Unix(int64(couponHistoryDetail.Coupon.EndTime), 0) // 结束使用时间

		switch useType {
		case pb.CouponUseType_COUPON_USE_TYPE_GENERAL:
			// 0->全场通用
			// 判断是否满足优惠起点
			// 计算购物车商品的总价
			totalAmount, _ := s.calcTotalAmount(cartPromotionItems)
			compare, _ := decimalUtils.CompareDecimal(totalAmount, minPoint)
			if now.Before(endTime) && compare >= 0 {
				enableList = append(enableList, couponHistoryDetail)
			} else {
				disableList = append(disableList, couponHistoryDetail)
			}
		case pb.CouponUseType_COUPON_USE_TYPE_SPECIFIC_CATEGORY:
			// 1->指定分类
			// 计算指定分类商品的总价
			productCategoryIDs := couponHistoryDetail.CategoryRelations.GetProductCategoryIDs()
			totalAmount, _ := s.calcTotalAmountByProductCategoryID(cartPromotionItems, productCategoryIDs)
			compare1, _ := decimalUtils.CompareDecimal(totalAmount, "0.00")
			compare2, _ := decimalUtils.CompareDecimal(totalAmount, minPoint)
			if now.Before(endTime) && compare1 > 0 && compare2 >= 0 {
				enableList = append(enableList, couponHistoryDetail)
			} else {
				disableList = append(disableList, couponHistoryDetail)
			}
		case pb.CouponUseType_COUPON_USE_TYPE_SPECIFIC_PRODUCT:
			// 2->指定商品
			// 计算指定商品的总价
			productIDs := couponHistoryDetail.ProductRelations.GetProductIDs()
			totalAmount, _ := s.calcTotalAmountByProductID(cartPromotionItems, productIDs)
			compare1, _ := decimalUtils.CompareDecimal(totalAmount, "0.00")
			compare2, _ := decimalUtils.CompareDecimal(totalAmount, minPoint)
			if now.Before(endTime) && compare1 > 0 && compare2 >= 0 {
				enableList = append(enableList, couponHistoryDetail)
			} else {
				disableList = append(disableList, couponHistoryDetail)
			}
		}
	}

	if canUse {
		return enableList, nil
	} else {
		return disableList, nil
	}
}

// CouponListByProduct 获取当前商品相关优惠券
func (s CouponUseCase) CouponListByProduct(ctx context.Context, req *pb.CouponListByProductReq) (*pb.CouponListByProductRsp, error) {
	//TODO implement me
	panic("implement me")
}

// 获取优惠券历史详情
func (s CouponUseCase) getDetailList(ctx context.Context, memberID uint64) ([]*portal_entity.CouponHistoryDetail, error) {
	res := make([]*portal_entity.CouponHistoryDetail, 0)

	// 查询符合条件的优惠券历史记录
	couponHistories, err := s.couponHistoryRepo.GetNoUseCouponHistory(ctx, memberID)
	if err != nil {
		return nil, err
	}

	// 优惠券id
	couponIDs := couponHistories.GetCouponIDs()

	// 查询所有优惠券
	coupons, err := s.couponRepo.GetByIDs(ctx, couponIDs)
	if err != nil {
		return nil, err
	}

	// 查询优惠券和商品的关系
	productRelations, err := s.couponProductRelationRepo.GetByCouponID(ctx, couponIDs)
	if err != nil {
		return nil, err
	}

	// 查询优惠券和商品分类关系
	productCategoryRelations, err := s.couponProductCategoryRelationRepo.GetByCouponID(ctx, couponIDs)
	if err != nil {
		return nil, err
	}

	temp := make(map[uint64]*portal_entity.CouponHistoryDetail)

	// 填充优惠券
	for _, coupon := range coupons {
		temp[coupon.ID] = &portal_entity.CouponHistoryDetail{Coupon: coupon}
	}
	// 填充优惠券历史
	for _, couponHistory := range couponHistories {
		if detail, ok := temp[couponHistory.CouponID]; ok {
			detail.CouponHistory = couponHistory
		}
	}
	// 填充优惠券和商品的关系
	for _, productRelation := range productRelations {
		if detail, ok := temp[productRelation.CouponID]; ok {
			detail.ProductRelations = append(detail.ProductRelations, productRelation)
		}
	}
	// 填充优惠券和商品分类关系
	for _, productCategoryRelation := range productCategoryRelations {
		if detail, ok := temp[productCategoryRelation.CouponID]; ok {
			detail.CategoryRelations = append(detail.CategoryRelations, productCategoryRelation)
		}
	}

	for _, v := range temp {
		res = append(res, v)
	}
	return res, nil
}

func (s CouponUseCase) calcTotalAmount(cartItemListPromotions []*pb.CartItemListPromotion) (string, error) {
	decimalUtils := util.DecimalUtils
	total := "0.00"
	for _, item := range cartItemListPromotions {
		// 计算实际价格
		realPrice, err := decimalUtils.SubtractDecimal(item.Price, item.ReduceAmount)
		if err != nil {
			return "", err
		}
		// 计算总价
		totalForItem, err := decimalUtils.MultiplyDecimal(realPrice, fmt.Sprintf("%d", item.Quantity))
		if err != nil {
			return "", err
		}
		total, err = decimalUtils.AddDecimal(total, totalForItem)
		if err != nil {
			return "", err
		}
	}
	return total, nil
}

func (s CouponUseCase) calcTotalAmountByProductCategoryID(cartItemListPromotions []*pb.CartItemListPromotion, productCategoryIDs []uint64) (string, error) {
	decimalUtils := util.DecimalUtils
	total := "0.00"
	for _, item := range cartItemListPromotions {
		// 检查商品是否属于指定分类
		if util.SliceExist[uint64](productCategoryIDs, item.ProductCategoryId) {
			// 计算实际价格
			realPrice, err := decimalUtils.SubtractDecimal(item.Price, item.ReduceAmount)
			if err != nil {
				return "", err
			}
			// 计算总价
			totalForItem, err := decimalUtils.MultiplyDecimal(realPrice, fmt.Sprintf("%d", item.Quantity))
			if err != nil {
				return "", err
			}
			total, err = decimalUtils.AddDecimal(total, totalForItem)
			if err != nil {
				return "", err
			}
		}
	}
	return total, nil
}

// calcTotalAmountByProductID 计算指定商品的总价
func (s CouponUseCase) calcTotalAmountByProductID(cartItemListPromotions []*pb.CartItemListPromotion, productIDs []uint64) (string, error) {
	decimalUtils := util.DecimalUtils
	total := "0"
	for _, item := range cartItemListPromotions {
		// 检查商品是否属于指定的产品 ID
		if util.SliceExist[uint64](productIDs, item.ProductId) {
			// 计算实际价格
			realPrice, err := decimalUtils.SubtractDecimal(item.Price, item.ReduceAmount)
			if err != nil {
				return "", err
			}
			// 计算总价
			totalForItem, err := decimalUtils.MultiplyDecimal(realPrice, fmt.Sprintf("%d", item.Quantity))
			if err != nil {
				return "", err
			}
			total, err = decimalUtils.AddDecimal(total, totalForItem)
			if err != nil {
				return "", err
			}
		}
	}
	return total, nil
}
