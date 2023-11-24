package usecase

import (
	"context"

	"github.com/baker-yuan/go-mall/application/admin/internal/usecase/assembler"
	"github.com/baker-yuan/go-mall/application/admin/pkg/db"
	pb "github.com/baker-yuan/go-mall/proto/mall"
)

// ProductUseCase 商品管理Service实现类
type ProductUseCase struct {
	productRepo         IProductRepo         // 操作商品
	brandRepo           IBrandRepo           // 操作商品品牌
	productCategoryRepo IProductCategoryRepo // 操作商品分类
}

// NewProduct 创建商品管理Service实现类
func NewProduct(productRepo IProductRepo,
	brandRepo IBrandRepo,
	productCategoryRepo IProductCategoryRepo) *ProductUseCase {
	return &ProductUseCase{
		productRepo:         productRepo,
		brandRepo:           brandRepo,
		productCategoryRepo: productCategoryRepo,
	}
}

// CreateProduct 添加商品
func (c ProductUseCase) CreateProduct(ctx context.Context, param *pb.AddOrUpdateProductParam) error {
	// 数据转换
	product := assembler.AddOrUpdateProductParamToEntity(param)

	// 保存
	if err := c.productRepo.Create(ctx, product); err != nil {
		return err
	}

	return nil
}

// UpdateProduct 修改商品
func (c ProductUseCase) UpdateProduct(ctx context.Context, param *pb.AddOrUpdateProductParam) error {
	oldProduct, err := c.productRepo.GetByID(ctx, param.GetId())
	if err != nil {
		return err
	}

	// 数据转换
	product := assembler.AddOrUpdateProductParamToEntity(param)
	product.ID = param.Id
	product.CreatedAt = oldProduct.CreatedAt

	// 更新商品
	return c.productRepo.Update(ctx, product)
}

// GetProducts 分页查询商品
func (c ProductUseCase) GetProducts(ctx context.Context, param *pb.GetProductsParam) ([]*pb.Product, uint32, error) {
	opts := make([]db.DBOption, 0)
	products, pageTotal, err := c.productRepo.GetByDBOption(ctx, param.GetPageNum(), param.GetPageSize(), opts...)
	if err != nil {
		return nil, 0, err
	}

	categories, err := c.productCategoryRepo.GetByIDs(ctx, products.CategoryIDs())
	if err != nil {
		return nil, 0, err
	}

	brands, err := c.brandRepo.GetByIDs(ctx, products.BrandIDs())
	if err != nil {
		return nil, 0, err
	}

	models := make([]*pb.Product, 0)
	for _, product := range products {
		models = append(models, assembler.ProductEntityToModel(product, categories.NameMap(), brands.NameMap()))

	}
	return models, pageTotal, nil
}

// GetProduct 根据id获取商品
func (c ProductUseCase) GetProduct(ctx context.Context, id uint64) (*pb.Product, error) {
	product, err := c.productRepo.GetByID(ctx, id)
	if err != nil {
		return nil, err
	}

	categories, err := c.productCategoryRepo.GetByIDs(ctx, []uint64{product.ProductCategoryID})
	if err != nil {
		return nil, err
	}

	brands, err := c.brandRepo.GetByIDs(ctx, []uint64{product.BrandID})
	if err != nil {
		return nil, err
	}

	return assembler.ProductEntityToModel(product, categories.NameMap(), brands.NameMap()), nil
}

// DeleteProduct 删除商品
func (c ProductUseCase) DeleteProduct(ctx context.Context, id uint64) error {
	return c.productRepo.DeleteByID(ctx, id)
}
