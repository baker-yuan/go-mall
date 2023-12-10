package usecase

import (
	"context"

	"github.com/baker-yuan/go-mall/application/portal/internal/usecase/assembler"
	pb "github.com/baker-yuan/go-mall/proto/mall"
)

// ProductUseCase 商品管理Service实现类
type ProductUseCase struct {
	productRepo IProductRepo // 操作商品
}

// NewProduct 创建商品管理Service实现类
func NewProduct(
	productRepo IProductRepo,
) *ProductUseCase {
	return &ProductUseCase{
		productRepo: productRepo,
	}
}

// SearchProduct 综合搜索商品
func (c ProductUseCase) SearchProduct(ctx context.Context, req *pb.SearchProductReq) ([]*pb.ProductItem, error) {
	var (
		res = make([]*pb.ProductItem, 0)
	)
	products, err := c.productRepo.SearchProduct(ctx, req)
	if err != nil {
		return nil, err
	}

	for _, product := range products {
		res = append(res, assembler.ProductEntityToModel(product))
	}
	return res, nil
}
