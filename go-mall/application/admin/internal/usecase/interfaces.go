// Package usecase implements application business logic. Each logic group in own file.
package usecase

import (
	"context"

	"github.com/baker-yuan/go-mall/application/admin/internal/entity"
	"github.com/baker-yuan/go-mall/application/admin/pkg/db"
	pb "github.com/baker-yuan/go-mall/proto/mall"
)

//go:generate mockgen -source=interfaces.go -destination=./mocks_test.go -package=usecase_test

// ProductCategory 商品分类表
type (
	// IProductCategoryUseCase 业务逻辑
	IProductCategoryUseCase interface {
		// CreateProductCategory 添加商品分类
		CreateProductCategory(ctx context.Context, productCategoryParam *pb.AddOrUpdateProductCategoryParam) error
		// UpdateProductCategory 修改商品分类
		UpdateProductCategory(ctx context.Context, productCategoryParam *pb.AddOrUpdateProductCategoryParam) error

		// GetProductCategories 分页查询商品分类
		GetProductCategories(ctx context.Context, param *pb.GetProductCategoriesParam) ([]*pb.ProductCategory, uint32, error)
		// GetProductCategory 根据id获取商品分类
		GetProductCategory(ctx context.Context, categoryID uint64) (*pb.ProductCategory, error)
		// DeleteProductCategory 删除商品分类
		DeleteProductCategory(ctx context.Context, categoryID uint64) error
		// UpdateProductCategoryNavStatus 修改导航栏显示状态
		UpdateProductCategoryNavStatus(ctx context.Context, categoryIDs []uint64, navStatus uint32) error
		// UpdateProductCategoryShowStatus 修改显示状态
		UpdateProductCategoryShowStatus(ctx context.Context, categoryIDs []uint64, showStatus uint32) error
		// GetProductCategoriesWithChildren 查询所有一级分类及子分类
		GetProductCategoriesWithChildren(ctx context.Context) ([]*pb.ProductCategoryTreeItem, error)
	}

	// IProductCategoryRepo 数据存储操作
	IProductCategoryRepo interface {
		WithByParentID(parentID uint64) db.DBOption
		WithByID(value uint64) db.DBOption
		WithByName(name string) db.DBOption

		// Create 创建商品分类
		Create(ctx context.Context, productCategory *entity.ProductCategory) error
		// DeleteByID 根据主键ID删除商品分类
		DeleteByID(ctx context.Context, categoryID uint64) error
		// Update 修改商品分类
		Update(ctx context.Context, productCategory *entity.ProductCategory) error
		// GetByID 根据主键ID查询商品分类
		GetByID(ctx context.Context, id uint64) (*entity.ProductCategory, error)

		// GetByDBOption 根据动态条件查询商品分类
		GetByDBOption(ctx context.Context, pageNum uint32, pageSize uint32, opts ...db.DBOption) ([]*entity.ProductCategory, uint32, error)

		// CreateWithTX 创建商品分类
		CreateWithTX(ctx context.Context, productCategory *entity.ProductCategory) error
		// UpdateWithTX 修改商品分类
		UpdateWithTX(ctx context.Context, productCategory *entity.ProductCategory) error

		// UpdateFieldByID 根据ID修改
		UpdateFieldByID(ctx context.Context, category *entity.ProductCategory, fields ...string) error
		// UpdateProductCategoryNavStatus 修改导航栏显示状态
		UpdateProductCategoryNavStatus(ctx context.Context, categoryIDs []uint64, navStatus uint32) error
		// UpdateProductCategoryShowStatus 修改显示状态
		UpdateProductCategoryShowStatus(ctx context.Context, categoryIDs []uint64, showStatus uint32) error
	}
)

// ProductCategory 产品的分类和属性的关系表，用于设置分类筛选条件（只支持一级分类）
type (
	// IProductCategoryAttributeRelationRepo 数据存储操作
	IProductCategoryAttributeRelationRepo interface {
		// Create 创建
		Create(ctx context.Context, relation *entity.ProductCategoryAttributeRelation) error
		// DeleteByID 根据主键ID删除
		DeleteByID(ctx context.Context, id uint64) error
		// Update 修改
		Update(ctx context.Context, relation *entity.ProductCategoryAttributeRelation) error
		// GetByID 根据主键ID查询
		GetByID(ctx context.Context, id uint64) (*entity.ProductCategoryAttributeRelation, error)

		// BatchCreateWithTX 创建
		BatchCreateWithTX(ctx context.Context, productCategoryID uint64, attributeIds []uint64) error
		// CreateWithTX 创建
		CreateWithTX(ctx context.Context, relation *entity.ProductCategoryAttributeRelation) error
		// DeleteByProductCategoryIDWithTX 根据productCategoryID删除
		DeleteByProductCategoryIDWithTX(ctx context.Context, productCategoryID uint64) error
	}
)

// PmsProduct 商品信息
type (
	// IProductUseCase 业务逻辑
	IProductUseCase interface {
	}

	// IProductRepo 数据存储操作
	IProductRepo interface {
		// UpdateProductCategoryNameByProductCategoryIDWithTX 新商品中的名称
		UpdateProductCategoryNameByProductCategoryIDWithTX(ctx context.Context, productCategoryID uint64, productCategoryName string) error
	}
)

// Brand 商品品牌表
type (
	// IBrandUseCase 业务逻辑
	IBrandUseCase interface {
		// CreateBrand 添加商品品牌表
		CreateBrand(ctx context.Context, param *pb.AddOrUpdateBrandParam) error
		// UpdateBrand 修改商品品牌表
		UpdateBrand(ctx context.Context, param *pb.AddOrUpdateBrandParam) error
		// GetBrands 分页查询商品品牌表
		GetBrands(ctx context.Context, param *pb.GetBrandsParam) ([]*pb.Brand, uint32, error)
		// GetBrand 根据id获取商品品牌表
		GetBrand(ctx context.Context, id uint64) (*pb.Brand, error)
		// DeleteBrand 删除商品品牌表
		DeleteBrand(ctx context.Context, id uint64) error
	}

	// IBrandRepo 数据存储操作
	IBrandRepo interface {
		WithByName(name string) db.DBOption
		WithByShowStatus(showStatus uint8) db.DBOption

		// Create 创建商品品牌表
		Create(ctx context.Context, brand *entity.Brand) error
		// DeleteByID 根据主键ID删除商品品牌表
		DeleteByID(ctx context.Context, id uint64) error
		// Update 修改商品品牌表
		Update(ctx context.Context, brand *entity.Brand) error
		// GetByID 根据主键ID查询商品品牌表
		GetByID(ctx context.Context, id uint64) (*entity.Brand, error)
		// GetByDBOption 根据动态条件查询商品品牌表
		GetByDBOption(ctx context.Context, pageNum uint32, pageSize uint32, opts ...db.DBOption) ([]*entity.Brand, uint32, error)
	}
)

// ProductAttributeCategory 产品属性分类表
type (
	// IProductAttributeCategoryUseCase 业务逻辑
	IProductAttributeCategoryUseCase interface {
		// CreateProductAttributeCategory 添加产品属性分类表
		CreateProductAttributeCategory(ctx context.Context, param *pb.AddOrUpdateProductAttributeCategoryParam) error
		// UpdateProductAttributeCategory 修改产品属性分类表
		UpdateProductAttributeCategory(ctx context.Context, param *pb.AddOrUpdateProductAttributeCategoryParam) error
		// GetProductAttributeCategories 分页查询产品属性分类表
		GetProductAttributeCategories(ctx context.Context, param *pb.GetProductAttributeCategoriesParam) ([]*pb.ProductAttributeCategory, uint32, error)
		// GetProductAttributeCategory 根据id获取产品属性分类表
		GetProductAttributeCategory(ctx context.Context, id uint64) (*pb.ProductAttributeCategory, error)
		// DeleteProductAttributeCategory 删除产品属性分类表
		DeleteProductAttributeCategory(ctx context.Context, id uint64) error
	}

	// IProductAttributeCategoryRepo 数据存储操作
	IProductAttributeCategoryRepo interface {
		WithByName(name string) db.DBOption

		// Create 创建产品属性分类表
		Create(ctx context.Context, productAttributeCategory *entity.ProductAttributeCategory) error
		// DeleteByID 根据主键ID删除产品属性分类表
		DeleteByID(ctx context.Context, id uint64) error
		// Update 修改产品属性分类表
		Update(ctx context.Context, productAttributeCategory *entity.ProductAttributeCategory) error
		// GetByID 根据主键ID查询产品属性分类表
		GetByID(ctx context.Context, id uint64) (*entity.ProductAttributeCategory, error)
		// GetByDBOption 根据动态条件查询产品属性分类表
		GetByDBOption(ctx context.Context, pageNum uint32, pageSize uint32, opts ...db.DBOption) ([]*entity.ProductAttributeCategory, uint32, error)
	}
)

// ProductAttribute 商品属性参数表
type (
	// IProductAttributeUseCase 业务逻辑
	IProductAttributeUseCase interface {
		// CreateProductAttribute 添加商品属性参数表
		CreateProductAttribute(ctx context.Context, param *pb.AddOrUpdateProductAttributeParam) error
		// UpdateProductAttribute 修改商品属性参数表
		UpdateProductAttribute(ctx context.Context, param *pb.AddOrUpdateProductAttributeParam) error
		// GetProductAttributes 分页查询商品属性参数表
		GetProductAttributes(ctx context.Context, param *pb.GetProductAttributesParam) ([]*pb.ProductAttribute, uint32, error)
		// GetProductAttribute 根据id获取商品属性参数表
		GetProductAttribute(ctx context.Context, id uint64) (*pb.ProductAttribute, error)
		// DeleteProductAttribute 删除商品属性参数表
		DeleteProductAttribute(ctx context.Context, id uint64) error
	}

	// IProductAttributeRepo 数据存储操作
	IProductAttributeRepo interface {
		WithByName(name string) db.DBOption
		WithByProductAttributeCategoryID(productAttributeCategoryID uint32) db.DBOption
		WithByType(tpe uint32) db.DBOption

		// Create 创建商品属性参数表
		Create(ctx context.Context, productAttribute *entity.ProductAttribute) error
		// DeleteByID 根据主键ID删除商品属性参数表
		DeleteByID(ctx context.Context, id uint64) error
		// Update 修改商品属性参数表
		Update(ctx context.Context, productAttribute *entity.ProductAttribute) error
		// GetByID 根据主键ID查询商品属性参数表
		GetByID(ctx context.Context, id uint64) (*entity.ProductAttribute, error)
		// GetByDBOption 根据动态条件查询商品属性参数表
		GetByDBOption(ctx context.Context, pageNum uint32, pageSize uint32, opts ...db.DBOption) ([]*entity.ProductAttribute, uint32, error)
	}
)
