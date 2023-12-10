package usecase

import (
	"context"

	"github.com/baker-yuan/go-mall/application/portal/internal/usecase/assembler"
	pb "github.com/baker-yuan/go-mall/proto/mall"
)

// HomeUseCase 首页
type HomeUseCase struct {
	productCategoryRepo IProductCategoryRepo // 操作商品分类
}

// NewHome 创建首页Service实现类
func NewHome(productCategoryRepo IProductCategoryRepo) *HomeUseCase {
	return &HomeUseCase{
		productCategoryRepo: productCategoryRepo,
	}
}

// ProductCategoryList 分页查询订单
func (h HomeUseCase) ProductCategoryList(ctx context.Context, req *pb.ProductCategoryListReq) ([]*pb.ProductCategoryItem, error) {
	var (
		res = make([]*pb.ProductCategoryItem, 0)
	)
	categories, err := h.productCategoryRepo.GetShowProductCategory(ctx, req.GetParentId())
	if err != nil {
		return nil, err
	}
	for _, category := range categories {
		res = append(res, assembler.ProductCategoryEntityToModel(category))
	}
	return res, nil
}
