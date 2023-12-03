package usecase

import (
	"context"

	"github.com/baker-yuan/go-mall/application/admin/internal/usecase/assembler"
	"github.com/baker-yuan/go-mall/application/admin/pkg/db"
	pb "github.com/baker-yuan/go-mall/proto/mall"
)

// OrderUseCase 订单管理Service实现类
type OrderUseCase struct {
	orderRepo IOrderRepo // 操作订单
}

// NewOrder 创建订单管理Service实现类
func NewOrder(orderRepo IOrderRepo) *OrderUseCase {
	return &OrderUseCase{
		orderRepo: orderRepo,
	}
}

// GetOrders 分页查询订单
func (c OrderUseCase) GetOrders(ctx context.Context, param *pb.GetOrdersParam) ([]*pb.Order, uint32, error) {
	opts := make([]db.DBOption, 0)

	orders, pageTotal, err := c.orderRepo.GetByDBOption(ctx, param.GetPageNum(), param.GetPageSize(), opts...)
	if err != nil {
		return nil, 0, err
	}

	results := make([]*pb.Order, 0)
	for _, order := range orders {
		results = append(results, assembler.OrderEntityToModel(order))
	}
	return results, pageTotal, nil
}

// GetOrder 根据id获取订单
func (c OrderUseCase) GetOrder(ctx context.Context, id uint64) (*pb.Order, error) {
	order, err := c.orderRepo.GetByID(ctx, id)
	if err != nil {
		return nil, err
	}
	return assembler.OrderEntityToModel(order), nil
}

// DeleteOrder 删除订单
func (c OrderUseCase) DeleteOrder(ctx context.Context, id uint64) error {
	return c.orderRepo.DeleteByID(ctx, id)
}
