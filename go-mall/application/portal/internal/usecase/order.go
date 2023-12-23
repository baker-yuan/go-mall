package usecase

import (
	"context"
	"fmt"

	"github.com/baker-yuan/go-mall/application/portal/internal/usecase/assembler"
	"github.com/baker-yuan/go-mall/common/entity"
	"github.com/baker-yuan/go-mall/common/util"
	pb "github.com/baker-yuan/go-mall/proto/mall"
)

// OrderUseCase 订单表管理Service实现类
type OrderUseCase struct {
	orderRepo                IOrderRepo                // 操作订单表
	memberRepo               IMemberRepo               // 操作会员表
	memberReceiveAddressRepo IMemberReceiveAddressRepo // 操作会员收货地址表
	jsonDynamicConfigRepo    IJsonDynamicConfigRepo    // 操作JSON动态配置
	//
	cartItemUseCase ICartItemUseCase // 购物车
	couponUseCase   ICouponUseCase   // 优惠券
}

// NewOrder 创建订单表管理Service实现类
func NewOrder(
	orderRepo IOrderRepo,
	memberRepo IMemberRepo,
	memberReceiveAddressRepo IMemberReceiveAddressRepo,
	jsonDynamicConfigRepo IJsonDynamicConfigRepo,
	//
	cartItemUseCase ICartItemUseCase,
	couponUseCase ICouponUseCase,
) *OrderUseCase {
	return &OrderUseCase{
		orderRepo:                orderRepo,
		memberRepo:               memberRepo,
		memberReceiveAddressRepo: memberReceiveAddressRepo,
		jsonDynamicConfigRepo:    jsonDynamicConfigRepo,
		//
		cartItemUseCase: cartItemUseCase,
		couponUseCase:   couponUseCase,
	}
}

// GenerateConfirmOrder 根据用户购物车信息生成确认单信息
func (c OrderUseCase) GenerateConfirmOrder(ctx context.Context, memberID uint64, req *pb.GenerateConfirmOrderReq) (*pb.GenerateConfirmOrderRsp, error) {
	var (
		res = &pb.GenerateConfirmOrderRsp{}
	)
	// 获取用户信息
	member, err := c.memberRepo.GetByID(ctx, memberID)
	if err != nil {
		return nil, err
	}

	// 获取购物车信息
	cartPromotionItems, err := c.cartItemUseCase.CartItemListPromotion(ctx, memberID, req.GetCartIds())
	if err != nil {
		return nil, err
	}
	res.CartPromotionItems = cartPromotionItems

	// 获取用户收货地址列表
	memberReceiveAddresses, err := c.memberReceiveAddressRepo.GetByMemberID(ctx, memberID)
	if err != nil {
		return nil, err
	}
	res.MemberReceiveAddresses = assembler.MemberReceiveAddressEntityToDetail(memberReceiveAddresses)

	// 获取用户可用优惠券列表
	couponHistoryDetails, err := c.couponUseCase.CouponListCart(ctx, memberID, cartPromotionItems, true)
	if err != nil {
		return nil, err
	}
	res.CouponHistoryDetails = assembler.CouponHistoryDetailToModel(couponHistoryDetails)

	// 获取用户积分
	memberIntegration := member.Integration
	res.MemberIntegration = memberIntegration

	// 获取积分使用规则
	cfg, _ := c.jsonDynamicConfigRepo.GetByBizType(ctx, entity.IntegrationConsumeSetting)
	integrationConsumeSetting, err := util.Unmarshal[entity.UmsIntegrationConsumeSetting](cfg)
	res.IntegrationConsumeSetting = assembler.IntegrationConsumeSettingEntityToDetail(integrationConsumeSetting)

	// 计算总金额、活动优惠、应付金额
	calcAmount, err := c.calcCartAmount(cartPromotionItems)
	res.CalcAmount = calcAmount

	return res, nil
}

// GenerateOrder 根据提交信息生成订单
func (c OrderUseCase) GenerateOrder(context.Context, *pb.GenerateOrderReq) (*pb.GenerateOrderRsp, error) {
	return nil, nil
}

// PaySuccess 用户支付成功的回调
func (c OrderUseCase) PaySuccess(context.Context, *pb.PaySuccessReq) (*pb.PaySuccessRsp, error) {
	return nil, nil
}

// CancelTimeOutOrder PaySuccess 自动取消超时订单
func (c OrderUseCase) CancelTimeOutOrder(context.Context, *pb.CancelTimeOutOrderReq) (*pb.CancelTimeOutOrderRsp, error) {
	return nil, nil
}

// CancelOrder 取消单个超时订单
func (c OrderUseCase) CancelOrder(context.Context, *pb.CancelOrderReq) (*pb.CancelOrderRsp, error) {
	return nil, nil
}

// OrderList 按状态分页获取用户订单列表
func (c OrderUseCase) OrderList(context.Context, *pb.OrderListReq) (*pb.OrderListRsp, error) {
	return nil, nil
}

// OrderDetail 根据ID获取订单详情
func (c OrderUseCase) OrderDetail(context.Context, *pb.OrderDetailReq) (*pb.OrderDetailRsp, error) {
	return nil, nil
}

// CancelUserOrder 用户取消订单
func (c OrderUseCase) CancelUserOrder(context.Context, *pb.CancelUserOrderReq) (*pb.CancelUserOrderRsp, error) {
	return nil, nil
}

// ConfirmReceiveOrder 用户确认收货
func (c OrderUseCase) ConfirmReceiveOrder(context.Context, *pb.ConfirmReceiveOrderReq) (*pb.ConfirmReceiveOrderRsp, error) {
	return nil, nil
}

// DeleteOrder 用户删除订单
func (c OrderUseCase) DeleteOrder(context.Context, *pb.PortalDeleteOrderReq) (*pb.PortalDeleteOrderRsp, error) {
	return nil, nil
}

// calcCartAmount 计算购物车中商品的价格
func (c OrderUseCase) calcCartAmount(cartItemListPromotions []*pb.CartItemListPromotion) (*pb.GenerateConfirmOrderRsp_CalcAmount, error) {
	decimalUtils := util.DecimalUtils

	calcAmount := &pb.GenerateConfirmOrderRsp_CalcAmount{
		FreightAmount: "0.00",
	}

	// 初始化总金额和促销金额
	totalAmount := "0.00"
	promotionAmount := "0.00"

	// 遍历购物车促销商品列表
	for _, item := range cartItemListPromotions {
		quantity := fmt.Sprintf("%d", item.Quantity)
		// 计算商品总价
		totalPrice, _ := decimalUtils.MultiplyDecimal(item.Price, quantity)
		totalAmount, _ = decimalUtils.AddDecimal(totalAmount, totalPrice)

		// 计算促销金额
		totalReduce, _ := decimalUtils.MultiplyDecimal(item.ReduceAmount, quantity)
		promotionAmount, _ = decimalUtils.AddDecimal(promotionAmount, totalReduce)
	}

	// 设置计算结果
	calcAmount.TotalAmount = util.DecimalUtils.TrimTrailingZeros(totalAmount)
	calcAmount.PromotionAmount = util.DecimalUtils.TrimTrailingZeros(promotionAmount)
	payAmount, _ := decimalUtils.SubtractDecimal(totalAmount, promotionAmount)
	calcAmount.PayAmount = util.DecimalUtils.TrimTrailingZeros(payAmount)

	return calcAmount, nil
}
