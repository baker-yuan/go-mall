package usecase

import (
	"context"
	"fmt"
	"strings"
	"time"

	portal_entity "github.com/baker-yuan/go-mall/application/portal/internal/entity"
	"github.com/baker-yuan/go-mall/application/portal/internal/usecase/assembler"
	"github.com/baker-yuan/go-mall/common/entity"
	"github.com/baker-yuan/go-mall/common/retcode"
	"github.com/baker-yuan/go-mall/common/util"
	"github.com/baker-yuan/go-mall/common/util/decimal"
	pb "github.com/baker-yuan/go-mall/proto/mall"
)

// OrderUseCase 订单表管理Service实现类
type OrderUseCase struct {
	orderRepo                IOrderRepo                // 操作订单表
	orderItemRepo            IOrderItemRepo            // 操作订单商品信息
	cartItemRepo             ICartItemRepo             // 操作购物车
	memberRepo               IMemberRepo               // 操作会员表
	memberReceiveAddressRepo IMemberReceiveAddressRepo // 操作会员收货地址表
	jsonDynamicConfigRepo    IJsonDynamicConfigRepo    // 操作JSON动态配置
	skuStockRepo             ISkuStockRepo             // 操作sku
	couponHistoryRepo        ICouponHistoryRepo        // 操作优惠券使用、领取历史
	//
	cartItemUseCase ICartItemUseCase // 购物车
	couponUseCase   ICouponUseCase   // 优惠券
}

// NewOrder 创建订单表管理Service实现类
func NewOrder(
	orderRepo IOrderRepo,
	orderItemRepo IOrderItemRepo,
	cartItemRepo ICartItemRepo,
	memberRepo IMemberRepo,
	memberReceiveAddressRepo IMemberReceiveAddressRepo,
	jsonDynamicConfigRepo IJsonDynamicConfigRepo,
	skuStockRepo ISkuStockRepo,
	couponHistoryRepo ICouponHistoryRepo,
	//
	cartItemUseCase ICartItemUseCase,
	couponUseCase ICouponUseCase,
) *OrderUseCase {
	return &OrderUseCase{
		orderRepo:                orderRepo,
		orderItemRepo:            orderItemRepo,
		cartItemRepo:             cartItemRepo,
		memberRepo:               memberRepo,
		memberReceiveAddressRepo: memberReceiveAddressRepo,
		jsonDynamicConfigRepo:    jsonDynamicConfigRepo,
		skuStockRepo:             skuStockRepo,
		couponHistoryRepo:        couponHistoryRepo,
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
func (c OrderUseCase) GenerateOrder(ctx context.Context, memberID uint64, orderParam *pb.GenerateOrderReq) (*pb.GenerateOrderRsp, error) {
	// 获取用户信息
	currentMember, err := c.memberRepo.GetByID(ctx, memberID)
	if err != nil {
		return nil, err
	}

	// 获取购物车信息
	cartPromotionItems, err := c.cartItemUseCase.CartItemListPromotion(ctx, memberID, orderParam.GetCartIds())
	if err != nil {
		return nil, err
	}

	// 收货地址id为空
	if orderParam.MemberReceiveAddressId == 0 {
		return nil, retcode.NewError(retcode.RetGenOrderMemberReceiveAddressIDCheckFail)
	}

	var (
		orderItems = make([]*entity.OrderItem, 0)
	)

	for _, cartPromotionItem := range cartPromotionItems {
		// 生成下单商品信息
		orderItem := &entity.OrderItem{
			// 商品信息
			ProductID:         cartPromotionItem.ProductId,
			ProductName:       cartPromotionItem.ProductName,
			ProductPic:        cartPromotionItem.ProductPic,
			ProductAttr:       cartPromotionItem.ProductAttr,
			ProductBrand:      cartPromotionItem.ProductBrand,
			ProductSN:         cartPromotionItem.ProductSn,
			ProductPrice:      cartPromotionItem.Price,
			ProductQuantity:   cartPromotionItem.Quantity,
			ProductSkuID:      cartPromotionItem.ProductSkuId,
			ProductSkuCode:    cartPromotionItem.ProductSkuCode,
			ProductCategoryID: cartPromotionItem.ProductCategoryId,
			PromotionAmount:   cartPromotionItem.ReduceAmount,
			PromotionName:     cartPromotionItem.PromotionMessage,
			GiftIntegration:   cartPromotionItem.Integration,
			GiftGrowth:        cartPromotionItem.Growth,
		}
		orderItems = append(orderItems, orderItem)
	}
	// 判断购物车中商品是否都有库存
	if c.hasStock(cartPromotionItems) {
		return nil, retcode.NewError(retcode.RetGenOrderNoStock)
	}

	// 判断使用使用了优惠券
	if orderParam.CouponId == 0 {
		// 不用优惠券
		for _, orderItem := range orderItems {
			orderItem.CouponAmount = "0.00"
		}
	} else {
		// 使用优惠券
		couponHistoryDetail, _ := c.getUseCoupon(ctx, cartPromotionItems, memberID, orderParam.CouponId)
		if couponHistoryDetail == nil {
			return nil, retcode.NewError(retcode.RetGenOrderCouponNotUse)
		}
		// 对下单商品的优惠券进行处理
		c.handleCouponAmount(orderItems, couponHistoryDetail)
	}

	// 判断是否使用积分
	if orderParam.UseIntegration == 0 {
		// 不使用积分
		for _, orderItem := range orderItems {
			orderItem.IntegrationAmount = "0.00"
		}
	} else {
		// 使用积分
		totalAmount := c.calcTotalAmount(orderItems)
		integrationAmount, _ := c.getUseIntegrationAmount(ctx, orderParam.UseIntegration, totalAmount, currentMember, orderParam.CouponId != 0)
		compare, _ := integrationAmount.Compare(util.DecimalUtils.NewBigDecimal("0.00"))
		if compare == 0 {
			return nil, retcode.NewError(retcode.RetGenOrderIntegrationAmountNotUse)
		} else {
			// 可用情况下分摊到可用商品中
			for _, orderItem := range orderItems {
				perAmount, _ := util.DecimalUtils.NewBigDecimal(orderItem.ProductPrice).
					DivideV2(totalAmount, 3, decimal.RoundHalfEven).
					Multiply(integrationAmount).ToString()
				orderItem.IntegrationAmount = perAmount
			}
		}
	}

	// 计算order_item的实付金额
	c.handleRealAmount(orderItems)
	// 进行库存锁定
	if err := c.lockStock(ctx, cartPromotionItems); err != nil {
		return nil, err
	}
	// 根据商品合计、运费、活动优惠、优惠券、积分计算应付金额

	order := &entity.Order{}

	order.DiscountAmount = "0.00"
	totalAmount, _ := c.calcTotalAmount(orderItems).ToString()
	order.TotalAmount = totalAmount
	order.FreightAmount = "0.00"

	promotionAmount, _ := c.calcPromotionAmount(orderItems)
	order.PromotionAmount = promotionAmount

	order.PromotionInfo = c.getOrderPromotionInfo(orderItems)

	if orderParam.CouponId == 0 {
		order.CouponAmount = "0.00"
	} else {
		order.CouponID = orderParam.CouponId
		couponAmount, _ := c.calcCouponAmount(orderItems)
		order.CouponAmount = couponAmount
	}

	if orderParam.UseIntegration == 0 {
		order.Integration = 0
		order.IntegrationAmount = "0.00"
	} else {
		order.Integration = orderParam.UseIntegration
		integrationAmount, _ := c.calcIntegrationAmount(orderItems)
		order.IntegrationAmount = integrationAmount
	}

	payAmount, _ := c.calcPayAmount(order)
	order.PayAmount = payAmount
	// 转化为订单信息并插入数据库
	order.MemberID = currentMember.ID
	order.MemberUsername = currentMember.Username
	order.CreateTime = uint32(time.Now().Unix())

	// 支付方式：0->未支付；1->支付宝；2->微信
	order.PayType = uint8(orderParam.PayType)
	// 订单来源：0->PC订单；1->app订单
	order.SourceType = 1
	// 订单状态：0->待付款；1->待发货；2->已发货；3->已完成；4->已关闭；5->无效订单
	order.Status = 0
	// 订单类型：0->正常订单；1->秒杀订单
	order.OrderType = 0
	// 收货人信息：姓名、电话、邮编、地址
	address, err := c.memberReceiveAddressRepo.SecurityGetByID(ctx, memberID, orderParam.MemberReceiveAddressId)
	if err != nil {
		return nil, err
	}
	order.ReceiverName = address.Name
	order.ReceiverPhone = address.PhoneNumber
	order.ReceiverPostCode = address.PostCode
	order.ReceiverProvince = address.Province
	order.ReceiverCity = address.City
	order.ReceiverRegion = address.Region
	order.ReceiverDetailAddress = address.DetailAddress
	// 0->未确认；1->已确认
	order.ConfirmStatus = 0
	order.DeleteStatus = 0
	// 计算赠送积分
	order.Integration = c.calcGifIntegration(orderItems)
	// 计算赠送成长值
	order.Growth = c.calcGiftGrowth(orderItems)
	// 生成订单号
	order.OrderSN = c.generateOrderSN(order)
	// 设置自动收货天数
	cfg, err := c.jsonDynamicConfigRepo.GetByBizType(ctx, entity.OrderSetting)
	orderSetting, err := util.Unmarshal[entity.OmsOrderSetting](cfg)
	order.AutoConfirmDay = orderSetting.ConfirmOvertime

	// TODO: 2018/9/3 bill_*,delivery_*
	// 插入order表和order_item表
	if err := c.orderRepo.Create(ctx, order); err != nil {
		return nil, err
	}
	for _, orderItem := range orderItems {
		orderItem.OrderID = order.ID
		orderItem.OrderSN = order.OrderSN
	}

	if err := c.orderItemRepo.Creates(ctx, orderItems); err != nil {
		return nil, err
	}
	//如使用优惠券更新优惠券使用状态
	if orderParam.CouponId != 0 {
		if err := c.updateCouponStatus(ctx, orderParam.CouponId, currentMember.ID, 1); err != nil {
			return nil, err
		}
	}
	// 如使用积分需要扣除积分
	if orderParam.UseIntegration != 0 {
		order.UseIntegration = orderParam.UseIntegration
		if currentMember.Integration == 0 {
			currentMember.Integration = 0
		}
		if err := c.memberRepo.UpdateIntegration(ctx, currentMember.ID, currentMember.Integration-orderParam.UseIntegration); err != nil {
			return nil, err
		}
	}
	// 删除购物车中的下单商品
	if err := c.deleteCartItemList(ctx, cartPromotionItems, memberID); err != nil {
		return nil, err
	}
	// 发送延迟消息取消订单
	//sendDelayMessageCancelOrder(order.getId());
	//Map<String, Object> result = new HashMap<>();
	//result.put("order", order);
	//result.put("orderItemList", orderItemList);
	return nil, err
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
func (c OrderUseCase) calcCartAmount(cartItemListPromotions []*pb.CartPromotionItem) (*pb.GenerateConfirmOrderRsp_CalcAmount, error) {
	decimalUtils := util.DecimalUtils
	calcAmount := &pb.GenerateConfirmOrderRsp_CalcAmount{
		FreightAmount: "0.00",
	}

	// 初始化总金额和促销金额
	totalAmount := decimalUtils.NewBigDecimal("0.00")
	promotionAmount := decimalUtils.NewBigDecimal("0.00")

	// 遍历购物车促销商品列表
	for _, item := range cartItemListPromotions {
		quantity := decimalUtils.NewBigDecimal(fmt.Sprintf("%d", item.Quantity))
		// 计算商品总价
		totalPrice := decimalUtils.NewBigDecimal(item.Price).Multiply(quantity)
		totalAmount = totalAmount.Add(totalPrice)

		// 计算促销金额
		totalReduce := decimalUtils.NewBigDecimal(item.ReduceAmount).Multiply(quantity)
		promotionAmount = promotionAmount.Add(totalReduce)
	}

	// 设置计算结果
	totalAmountStr, _ := totalAmount.ToString()
	calcAmount.TotalAmount = totalAmountStr

	promotionAmountStr, _ := promotionAmount.ToString()
	calcAmount.PromotionAmount = promotionAmountStr

	payAmount := totalAmount.Subtract(promotionAmount)
	payAmountStr, _ := payAmount.ToString()
	calcAmount.PayAmount = payAmountStr

	return calcAmount, nil
}

// hasStock 判断下单商品是否都有库存
func (c OrderUseCase) hasStock(cartItemListPromotions []*pb.CartPromotionItem) bool {
	for _, cartItemListPromotion := range cartItemListPromotions {
		// 判断真实库存是否为空
		// 判断真实库存是否小于0
		// 判断真实库存是否小于下单的数量
		if cartItemListPromotion.RealStock <= 0 || cartItemListPromotion.RealStock < cartItemListPromotion.Quantity {
			return false
		}
	}
	return true
}

// getUseCoupon 获取该用户可以使用的优惠券
// cartItemListPromotions 购物车优惠列表
// couponID 使用优惠券id
func (c OrderUseCase) getUseCoupon(ctx context.Context, cartItemListPromotions []*pb.CartPromotionItem,
	memberID uint64, couponID uint64) (*portal_entity.CouponHistoryDetail, error) {
	// 根据购物车信息获取可用优惠券
	couponHistoryDetails, err := c.couponUseCase.CouponListCart(ctx, memberID, cartItemListPromotions, true)
	if err != nil {
		return nil, err
	}
	for _, couponHistoryDetail := range couponHistoryDetails {
		if couponHistoryDetail.Coupon.ID == couponID {
			return couponHistoryDetail, nil
		}
	}
	return nil, nil
}

// handleCouponAmount 对优惠券优惠进行处理
//
// orderItems order_item列表
// couponHistoryDetail 可用优惠券详情
func (c OrderUseCase) handleCouponAmount(orderItems []*entity.OrderItem, couponHistoryDetail *portal_entity.CouponHistoryDetail) {
	coupon := couponHistoryDetail.Coupon
	switch coupon.UseType {
	case pb.CouponUseType_COUPON_USE_TYPE_GENERAL:
		// 全场通用
		c.calcPerCouponAmount(orderItems, coupon)
	case pb.CouponUseType_COUPON_USE_TYPE_SPECIFIC_CATEGORY:
		// 指定分类
		couponOrderItems, _ := c.getCouponOrderItemByRelation(couponHistoryDetail, orderItems, 0)
		c.calcPerCouponAmount(couponOrderItems, coupon)
	case pb.CouponUseType_COUPON_USE_TYPE_SPECIFIC_PRODUCT:
		// 指定商品
		couponOrderItems, _ := c.getCouponOrderItemByRelation(couponHistoryDetail, orderItems, 1)
		c.calcPerCouponAmount(couponOrderItems, coupon)
	}
}

// 对每个下单商品进行优惠券金额分摊的计算
//
// orderItems 可用优惠券的下单商品商品
func (c OrderUseCase) calcPerCouponAmount(orderItems []*entity.OrderItem, coupon *entity.Coupon) {
	totalAmount := c.calcTotalAmount(orderItems)
	for _, orderItem := range orderItems {
		// (商品价格/可用商品总价)*优惠券面额
		productPrice := util.DecimalUtils.NewBigDecimal(orderItem.ProductPrice)
		couponAmount, _ := productPrice.
			DivideV2(totalAmount, 3, decimal.RoundHalfEven).
			Multiply(util.DecimalUtils.NewBigDecimal(coupon.Amount)).
			ToString()
		orderItem.CouponAmount = couponAmount
	}
}

// getCouponOrderItemByRelation获取与优惠券有关系的下单商品
//
// couponHistoryDetails 优惠券详情
// orderItems 下单商品
// tpe 使用关系类型：0->相关分类；1->指定商品
func (c OrderUseCase) getCouponOrderItemByRelation(couponHistoryDetail *portal_entity.CouponHistoryDetail, orderItems []*entity.OrderItem, tpe int) ([]*entity.OrderItem, error) {
	var result []*entity.OrderItem
	if tpe == 0 {
		categoryIdList := couponHistoryDetail.CategoryRelations.GetProductCategoryIDs()
		for _, orderItem := range orderItems {
			if util.SliceExist[uint64](categoryIdList, orderItem.ProductCategoryID) {
				result = append(result, orderItem)
			} else {
				orderItem.CouponAmount = "0.00"
			}
		}
	} else if tpe == 1 {
		productIdList := couponHistoryDetail.ProductRelations.GetProductIDs()
		for _, orderItem := range orderItems {
			if util.SliceExist[uint64](productIdList, orderItem.ProductID) {
				result = append(result, orderItem)
			} else {
				orderItem.CouponAmount = "0.00"
			}
		}
	}
	return result, nil
}

// calcTotalAmount 计算总金额
func (c OrderUseCase) calcTotalAmount(orderItems []*entity.OrderItem) *decimal.BigDecimal {
	totalAmount := util.DecimalUtils.NewBigDecimal("0.00") // 初始化总金额为0
	for _, item := range orderItems {
		// 对每个订单项的价格和数量进行乘法运算
		productPrice := util.DecimalUtils.NewBigDecimal(item.ProductPrice)
		productQuantity := util.DecimalUtils.NewBigDecimal(fmt.Sprintf("%d", item.ProductQuantity))
		itemAmount := productPrice.Multiply(productQuantity)
		// 将结果累加到总金额中
		totalAmount = totalAmount.Add(itemAmount)
	}
	return totalAmount
}

// getUseIntegrationAmount 获取可用积分抵扣金额
//
// useIntegration 使用的积分数量
// totalAmount 订单总金额
// currentMember 使用的用户
// hasCoupon 是否已经使用优惠券
func (c OrderUseCase) getUseIntegrationAmount(ctx context.Context, useIntegration uint32, totalAmount *decimal.BigDecimal, currentMember *entity.Member, hasCoupon bool) (*decimal.BigDecimal, error) {
	decimalUtils := util.DecimalUtils
	zeroAmount := decimalUtils.NewBigDecimal("0.00")

	// 判断用户是否有这么多积分
	if useIntegration > currentMember.Integration {
		return zeroAmount, nil
	}

	// 根据积分使用规则判断是否可用
	// 是否可与优惠券共用
	cfg, _ := c.jsonDynamicConfigRepo.GetByBizType(ctx, entity.IntegrationConsumeSetting)
	integrationConsumeSetting, _ := util.Unmarshal[entity.UmsIntegrationConsumeSetting](cfg)
	if hasCoupon && integrationConsumeSetting.CouponStatus == 0 {
		// 不可与优惠券共用
		return zeroAmount, nil
	}

	// 是否达到最低使用积分门槛
	if useIntegration >= integrationConsumeSetting.UseUnit {
		return zeroAmount, nil
	}

	// 是否超过订单抵用最高百分比
	integrationAmount := decimalUtils.NewUint32BigDecimal(useIntegration).
		Divide(decimalUtils.NewUint32BigDecimal(integrationConsumeSetting.UseUnit))
	maxPercentBD := decimalUtils.NewUint32BigDecimal(integrationConsumeSetting.MaxPercentPerOrder).
		Divide(decimalUtils.NewBigDecimal("100"))
	maxAmount := totalAmount.Multiply(maxPercentBD)

	compare, _ := integrationAmount.Compare(maxAmount)
	if compare > 0 {
		return zeroAmount, nil
	}

	return integrationAmount, nil
}

// handleRealAmount 处理订单项的实际金额
func (c OrderUseCase) handleRealAmount(orderItems []*entity.OrderItem) {
	decimalUtils := util.DecimalUtils
	for _, orderItem := range orderItems {
		// 原价 - 促销优惠 - 优惠券抵扣 - 积分抵扣
		realAmount, _ := decimalUtils.NewBigDecimal(orderItem.ProductPrice).
			Subtract(decimalUtils.NewBigDecimal(orderItem.PromotionAmount)).
			Subtract(decimalUtils.NewBigDecimal(orderItem.CouponAmount)).
			Subtract(decimalUtils.NewBigDecimal(orderItem.IntegrationAmount)).ToString()
		orderItem.RealAmount = realAmount
	}
}

// lockStock 锁定下单商品的所有库存
func (c OrderUseCase) lockStock(ctx context.Context, cartItemListPromotions []*pb.CartPromotionItem) error {
	for _, cartPromotionItem := range cartItemListPromotions {
		skuStock, err := c.skuStockRepo.GetByID(ctx, cartPromotionItem.ProductSkuId)
		if err != nil {
			return err
		}
		skuStock.LockStock = skuStock.LockStock + cartPromotionItem.Quantity
		return c.skuStockRepo.Update(ctx, skuStock)
	}
	return nil
}

// deleteCartItemList 批量删除购物车中的商品
func (c OrderUseCase) deleteCartItemList(ctx context.Context, cartPromotionItems []*pb.CartPromotionItem, memberID uint64) error {
	ids := make([]uint64, 0)
	for _, cartPromotionItem := range cartPromotionItems {
		ids = append(ids, cartPromotionItem.Id)
	}
	return c.cartItemRepo.CartItemDelete(ctx, memberID, ids)
}

// updateCouponStatus 将优惠券信息更改为指定状态
//
// couponID 优惠券id
// memberID 会员id
// useStatus 0->未使用；1->已使用
func (c OrderUseCase) updateCouponStatus(ctx context.Context, couponID uint64, memberID uint64, useStatus uint8) error {
	if couponID == 0 {
		return nil
	}
	// 查询第一张优惠券
	couponHistory, err := c.couponHistoryRepo.GetNoUseFirstByMemberIDAndCouponID(ctx, memberID, couponID)
	if err != nil {
		return err
	}
	if couponHistory == nil {
		return nil
	}
	couponHistory.UseTime = uint32(time.Now().Unix())
	couponHistory.UseStatus = useStatus
	return c.couponHistoryRepo.Update(ctx, couponHistory)
}

// 计算该订单赠送的积分 calcGifIntegration
func (c OrderUseCase) calcGifIntegration(orderItems []*entity.OrderItem) uint32 {
	sum := uint32(0)
	for _, orderItem := range orderItems {
		sum += orderItem.GiftIntegration * orderItem.ProductQuantity
	}
	return sum
}

// calcGiftGrowth 计算该订单赠送的成长值
func (c OrderUseCase) calcGiftGrowth(orderItems []*entity.OrderItem) uint32 {
	sum := uint32(0)
	for _, orderItem := range orderItems {
		sum += orderItem.GiftGrowth * orderItem.ProductQuantity
	}
	return sum
}

// generateOrderSN 生成18位订单编号:8位日期+2位平台号码+2位支付方式+6位以上自增id
func (c OrderUseCase) generateOrderSN(order *entity.Order) string {
	// todo
	return ""
}

// calcPromotionAmount 计算订单活动优惠
func (c OrderUseCase) calcPromotionAmount(orderItems []*entity.OrderItem) (string, error) {
	decimalUtils := util.DecimalUtils
	promotionAmount := decimalUtils.NewBigDecimal("0.00")
	for _, orderItem := range orderItems {
		if len(orderItem.PromotionAmount) != 0 {
			// 计算单个订单项的总优惠金额
			itemPromotion := decimalUtils.NewBigDecimal(orderItem.PromotionAmount)
			itemQuantity := decimalUtils.NewUint32BigDecimal(orderItem.ProductQuantity)
			totalItemPromotion := itemPromotion.Multiply(itemQuantity)

			// 累加到总优惠金额
			promotionAmount = promotionAmount.Add(totalItemPromotion)
		}
	}

	return promotionAmount.ToString()
}

// getOrderPromotionInfo 获取订单促销信息
func (c OrderUseCase) getOrderPromotionInfo(orderItems []*entity.OrderItem) string {
	var promotionNames []string
	for _, orderItem := range orderItems {
		promotionNames = append(promotionNames, orderItem.PromotionName)
	}
	return strings.Join(promotionNames, ";")
}

// calcCouponAmount 计算订单优惠券金额
func (c OrderUseCase) calcCouponAmount(orderItems []*entity.OrderItem) (string, error) {
	decimalUtils := util.DecimalUtils
	couponAmount := decimalUtils.NewUint32BigDecimal(0)
	for _, orderItem := range orderItems {
		if len(orderItem.CouponAmount) != 0 {
			itemCouponAmount := decimalUtils.NewBigDecimal(orderItem.CouponAmount).
				Multiply(decimalUtils.NewUint32BigDecimal(orderItem.ProductQuantity))
			couponAmount = couponAmount.Add(itemCouponAmount)
		}
	}
	return couponAmount.ToString()
}

// 计算订单优惠券金额
func (c OrderUseCase) calcIntegrationAmount(orderItems []*entity.OrderItem) (string, error) {
	decimalUtils := util.DecimalUtils
	integrationAmount := decimalUtils.NewBigDecimal("0")
	for _, orderItem := range orderItems {
		if len(orderItem.IntegrationAmount) != 0 {
			// 计算单个订单项的总积分金额
			itemIntegrationAmount := decimalUtils.NewBigDecimal(orderItem.IntegrationAmount)
			itemQuantity := decimalUtils.NewUint32BigDecimal(orderItem.ProductQuantity)
			totalItemIntegrationAmount := itemIntegrationAmount.Multiply(itemQuantity)
			// 累加到总积分金额
			integrationAmount = integrationAmount.Add(totalItemIntegrationAmount)
		}
	}
	return integrationAmount.ToString()
}

// calcPayAmount 计算订单应付金额
func (c OrderUseCase) calcPayAmount(order *entity.Order) (string, error) {
	decimalUtils := util.DecimalUtils
	// 总金额 + 运费 - 促销优惠 - 优惠券优惠 - 积分抵扣
	payAmount := decimalUtils.NewBigDecimal(order.TotalAmount).Add(decimalUtils.NewBigDecimal(order.FreightAmount)).
		Subtract(decimalUtils.NewBigDecimal(order.PromotionAmount)).
		Subtract(decimalUtils.NewBigDecimal(order.CouponAmount)).
		Subtract(decimalUtils.NewBigDecimal(order.IntegrationAmount))
	return payAmount.ToString()
}
