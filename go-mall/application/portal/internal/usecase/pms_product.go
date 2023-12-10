package usecase

import (
	"context"

	"github.com/baker-yuan/go-mall/application/portal/internal/usecase/assembler"
	pb "github.com/baker-yuan/go-mall/proto/mall"
)

// ProductUseCase 前台商品管理Service
type ProductUseCase struct {
	productRepo IProductRepo // 操作商品
	brandRepo   IBrandRepo   // 操作商品品牌
}

// NewProduct 创建前台商品管理Service实现类
func NewProduct(
	productRepo IProductRepo,
	brandRepo IBrandRepo,
) *ProductUseCase {
	return &ProductUseCase{
		productRepo: productRepo,
		brandRepo:   brandRepo,
	}
}

// SearchProduct 综合搜索商品
func (c ProductUseCase) SearchProduct(ctx context.Context, req *pb.SearchProductReq) ([]*pb.ProductList_Product, error) {
	var (
		res = make([]*pb.ProductList_Product, 0)
	)
	products, err := c.productRepo.SearchProduct(ctx, req)
	if err != nil {
		return nil, err
	}

	for _, product := range products {
		res = append(res, assembler.ProductEntityToProductListItem(product))
	}
	return res, nil
}

// ProductDetail 获取前台商品详情
func (c ProductUseCase) ProductDetail(ctx context.Context, req *pb.ProductDetailReq) (*pb.ProductAggregation, error) {
	var (
		res = &pb.ProductAggregation{}
	)
	productID := req.GetId()

	// 获取商品信息
	product, err := c.productRepo.GetByID(ctx, productID)
	if err != nil {
		return nil, err
	}
	// 获取品牌信息

	brand, err := c.brandRepo.GetByID(ctx, product.BrandID)
	if err != nil {
		return nil, err
	}

	res.Product = assembler.ProductEntityToDetail(product)
	res.Brand = assembler.BrandEntityToDetail(brand)
	return res, nil
}
