package usecase

import (
	"context"

	"github.com/baker-yuan/go-mall/application/portal/internal/usecase/assembler"
	pb "github.com/baker-yuan/go-mall/proto/mall"
)

// ProductUseCase 前台商品管理Service
type ProductUseCase struct {
	productRepo               IProductRepo               // 操作商品
	brandRepo                 IBrandRepo                 // 操作商品品牌
	productAttributeRepo      IProductAttributeRepo      // 操作商品属性参数
	productAttributeValueRepo IProductAttributeValueRepo // 操作产品参数信息
}

// NewProduct 创建前台商品管理Service实现类
func NewProduct(
	productRepo IProductRepo,
	brandRepo IBrandRepo,
	productAttributeRepo IProductAttributeRepo,
	productAttributeValueRepo IProductAttributeValueRepo,
) *ProductUseCase {
	return &ProductUseCase{
		productRepo:               productRepo,
		brandRepo:                 brandRepo,
		productAttributeRepo:      productAttributeRepo,
		productAttributeValueRepo: productAttributeValueRepo,
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
	// 获取商品属性信息
	productAttributes, err := c.productAttributeRepo.GetProductAttributeByCategoryID(ctx, product.ProductAttributeCategoryID)
	if err != nil {
		return nil, err
	}

	// 获取商品属性值信息
	attributeValues, err := c.productAttributeValueRepo.GetByProductAttributeID(ctx, productID, productAttributes.GetIDs())
	if err != nil {
		return nil, err
	}

	res.Product = assembler.ProductEntityToDetail(product)
	res.Brand = assembler.BrandEntityToDetail(brand)
	res.ProductAttributes = assembler.ProductAttributesToDetail(productAttributes)
	res.ProductAttributeValues = assembler.ProductAttributeValuesToDetail(attributeValues)
	return res, nil
}
