package usecase

import (
	"context"

	"github.com/baker-yuan/go-mall/application/admin/internal/usecase/assembler"
	"github.com/baker-yuan/go-mall/application/admin/pkg/db"
	pb "github.com/baker-yuan/go-mall/proto/mall"
)

// OrderReturnApplyUseCase 订单退货申请管理Service实现类
type OrderReturnApplyUseCase struct {
	orderReturnApplyRepo IOrderReturnApplyRepo // 操作订单退货申请
}

// NewOrderReturnApply 创建订单退货申请管理Service实现类
func NewOrderReturnApply(orderReturnApplyRepo IOrderReturnApplyRepo) *OrderReturnApplyUseCase {
	return &OrderReturnApplyUseCase{
		orderReturnApplyRepo: orderReturnApplyRepo,
	}
}

// GetOrderReturnApplies 分页查询订单退货申请
func (c OrderReturnApplyUseCase) GetOrderReturnApplies(ctx context.Context, param *pb.GetOrderReturnAppliesParam) ([]*pb.OrderReturnApply, uint32, error) {
	opts := make([]db.DBOption, 0)

	orderReturnApplies, pageTotal, err := c.orderReturnApplyRepo.GetByDBOption(ctx, param.GetPageNum(), param.GetPageSize(), opts...)
	if err != nil {
		return nil, 0, err
	}

	results := make([]*pb.OrderReturnApply, 0)
	for _, orderReturnApply := range orderReturnApplies {
		results = append(results, assembler.OrderReturnApplyEntityToModel(orderReturnApply))
	}
	return results, pageTotal, nil
}

// GetOrderReturnApply 根据id获取订单退货申请
func (c OrderReturnApplyUseCase) GetOrderReturnApply(ctx context.Context, id uint64) (*pb.OrderReturnApply, error) {
	orderReturnApply, err := c.orderReturnApplyRepo.GetByID(ctx, id)
	if err != nil {
		return nil, err
	}
	return assembler.OrderReturnApplyEntityToModel(orderReturnApply), nil
}

// DeleteOrderReturnApply 删除订单退货申请
func (c OrderReturnApplyUseCase) DeleteOrderReturnApply(ctx context.Context, id uint64) error {
	return c.orderReturnApplyRepo.DeleteByID(ctx, id)
}
