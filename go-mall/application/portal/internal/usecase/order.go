package usecase

import (
	"context"

	"github.com/baker-yuan/go-mall/application/portal/internal/usecase/assembler"
	pb "github.com/baker-yuan/go-mall/proto/mall"
)

// OrderUseCase 订单表管理Service实现类
type OrderUseCase struct {
	orderRepo                IOrderRepo                // 操作订单表
	memberRepo               IMemberRepo               // 操作会员表
	memberReceiveAddressRepo IMemberReceiveAddressRepo // 操作会员收货地址表
}

// NewOrder 创建订单表管理Service实现类
func NewOrder(orderRepo IOrderRepo, memberRepo IMemberRepo, memberReceiveAddressRepo IMemberReceiveAddressRepo) *OrderUseCase {
	return &OrderUseCase{
		orderRepo:                orderRepo,
		memberRepo:               memberRepo,
		memberReceiveAddressRepo: memberReceiveAddressRepo,
	}
}

// GenerateConfirmOrder 根据用户购物车信息生成确认单信息
func (c OrderUseCase) GenerateConfirmOrder(ctx context.Context, memberID uint64, req *pb.GenerateConfirmOrderReq) (*pb.GenerateConfirmOrderRsp, error) {
	var (
		res = &pb.GenerateConfirmOrderRsp{}
	)

	member, err := c.memberRepo.GetByID(ctx, memberID)
	if err != nil {
		return nil, err
	}
	// 获取购物车信息

	// 获取用户收货地址列表
	memberReceiveAddress, err := c.memberReceiveAddressRepo.GetByMemberID(ctx, memberID)
	if err != nil {
		return nil, err
	}
	// 获取用户可用优惠券列表

	// 获取用户积分
	memberIntegration := member.Integration

	// 获取积分使用规则

	// 计算总金额、活动优惠、应付金额

	res.MemberIntegration = memberIntegration
	res.MemberReceiveAddress = assembler.MemberReceiveAddressEntityToDetail(memberReceiveAddress)
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
func (c OrderUseCase) DeleteOrder(context.Context, *pb.DeleteOrderReq) (*pb.DeleteOrderRsp, error) {
	return nil, nil
}
