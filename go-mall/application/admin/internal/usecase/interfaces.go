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
		// GetByIDs 根据主键ID批量查询商品分类
		GetByIDs(ctx context.Context, ids []uint64) (entity.ProductCategories, error)

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
		// CreateProduct 添加商品
		CreateProduct(ctx context.Context, param *pb.AddOrUpdateProductParam) error
		// UpdateProduct 修改商品
		UpdateProduct(ctx context.Context, param *pb.AddOrUpdateProductParam) error
		// GetProducts 分页查询商品
		GetProducts(ctx context.Context, param *pb.GetProductsParam) ([]*pb.Product, uint32, error)
		// GetProduct 根据id获取商品
		GetProduct(ctx context.Context, id uint64) (*pb.Product, error)
		// DeleteProduct 删除商品
		DeleteProduct(ctx context.Context, id uint64) error
	}

	// IProductRepo 数据存储操作
	IProductRepo interface {
		WithByID(value uint64) db.DBOption
		WithByName(name string) db.DBOption
		WithByProductSN(productSN string) db.DBOption
		WithByBrandID(brandID uint64) db.DBOption
		WithByPublishStatus(publishStatus uint32) db.DBOption
		WithByVerifyStatus(verifyStatus uint32) db.DBOption
		WithByProductCategoryID(productCategoryID uint64) db.DBOption

		// Create 创建商品
		Create(ctx context.Context, product *entity.Product) error
		// DeleteByID 根据主键ID删除商品
		DeleteByID(ctx context.Context, id uint64) error
		// Update 修改商品
		Update(ctx context.Context, product *entity.Product) error
		// GetByID 根据主键ID查询商品
		GetByID(ctx context.Context, id uint64) (*entity.Product, error)
		// GetByDBOption 根据动态条件查询商品
		GetByDBOption(ctx context.Context, pageNum uint32, pageSize uint32, opts ...db.DBOption) (entity.Products, uint32, error)

		// CreateWithTX 创建商品
		CreateWithTX(ctx context.Context, product *entity.Product) error
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
		// GetByIDs 根据主键ID批量查询商品品牌表
		GetByIDs(ctx context.Context, ids []uint64) (entity.Brands, error)
		// GetByDBOption 根据动态条件查询商品品牌表
		GetByDBOption(ctx context.Context, pageNum uint32, pageSize uint32, opts ...db.DBOption) (entity.Brands, uint32, error)
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

// PmsMemberPrice 商品会员价格
type (
	// IMemberPriceUseCase 业务逻辑
	IMemberPriceUseCase interface {
	}

	// IMemberPriceRepo 数据存储操作
	IMemberPriceRepo interface {
		// Create 创建商品会员价格
		Create(ctx context.Context, pmsMemberPrice *entity.MemberPrice) error
		// DeleteByID 根据主键ID删除商品会员价格
		DeleteByID(ctx context.Context, id uint64) error
		// Update 修改商品会员价格
		Update(ctx context.Context, pmsMemberPrice *entity.MemberPrice) error
		// GetByID 根据主键ID查询商品会员价格
		GetByID(ctx context.Context, id uint64) (*entity.MemberPrice, error)
		// GetByDBOption 根据动态条件查询商品会员价格
		GetByDBOption(ctx context.Context, pageNum uint32, pageSize uint32, opts ...db.DBOption) ([]*entity.MemberPrice, uint32, error)

		// BatchCreateWithTX 创建商品会员价格
		BatchCreateWithTX(ctx context.Context, productID uint64, pmsMemberPrices []*entity.MemberPrice) error
	}
)

// ProductLadder 产品阶梯价格(只针对同商品)
type (
	// IProductLadderUseCase 业务逻辑
	IProductLadderUseCase interface {
	}

	// IProductLadderRepo 数据存储操作
	IProductLadderRepo interface {
		// Create 创建商品阶梯价格
		Create(ctx context.Context, productLadder *entity.ProductLadder) error
		// DeleteByID 根据主键ID删除商品阶梯价格
		DeleteByID(ctx context.Context, id uint64) error
		// Update 修改商品阶梯价格
		Update(ctx context.Context, productLadder *entity.ProductLadder) error
		// GetByID 根据主键ID查询商品阶梯价格
		GetByID(ctx context.Context, id uint64) (*entity.ProductLadder, error)
		// GetByDBOption 根据动态条件查询商品阶梯价格
		GetByDBOption(ctx context.Context, pageNum uint32, pageSize uint32, opts ...db.DBOption) ([]*entity.ProductLadder, uint32, error)

		// BatchCreateWithTX 创建商品阶梯价格
		BatchCreateWithTX(ctx context.Context, productID uint64, productLadders []*entity.ProductLadder) error
	}
)

// ProductFullReduction 产品满减(只针对同商品)
type (
	// IProductFullReductionUseCase 业务逻辑
	IProductFullReductionUseCase interface {
	}

	// IProductFullReductionRepo 数据存储操作
	IProductFullReductionRepo interface {
		// Create 创建产品满减
		Create(ctx context.Context, productFullReduction *entity.ProductFullReduction) error
		// DeleteByID 根据主键ID删除产品满减
		DeleteByID(ctx context.Context, id uint64) error
		// Update 修改产品满减
		Update(ctx context.Context, productFullReduction *entity.ProductFullReduction) error
		// GetByID 根据主键ID查询产品满减
		GetByID(ctx context.Context, id uint64) (*entity.ProductFullReduction, error)
		// GetByDBOption 根据动态条件查询产品满减
		GetByDBOption(ctx context.Context, pageNum uint32, pageSize uint32, opts ...db.DBOption) ([]*entity.ProductFullReduction, uint32, error)

		// BatchCreateWithTX 创建产品满减
		BatchCreateWithTX(ctx context.Context, productID uint64, productFullReductions []*entity.ProductFullReduction) error
	}
)

// SkuStock sku的库存
type (
	// ISkuStockUseCase 业务逻辑
	ISkuStockUseCase interface {
		// BatchUpdateSkuStock 批量修改sku的库存
		BatchUpdateSkuStock(ctx context.Context, param *pb.BatchUpdateSkuStockParam) error
		// GetSkuStocksByProductID 根据商品id分页查询sku的库存
		GetSkuStocksByProductID(ctx context.Context, param *pb.GetSkuStocksByProductIdParam) ([]*pb.SkuStock, error)
	}

	// ISkuStockRepo 数据存储操作
	ISkuStockRepo interface {
		WithByProductID(productId uint64) db.DBOption
		WithBySkuCode(skuCode string) db.DBOption

		// Create 创建sku的库存
		Create(ctx context.Context, skuStock *entity.SkuStock) error
		// DeleteByID 根据主键ID删除sku的库存
		DeleteByID(ctx context.Context, id uint64) error
		// Update 修改sku的库存
		Update(ctx context.Context, skuStock *entity.SkuStock) error
		// GetByID 根据主键ID查询sku的库存
		GetByID(ctx context.Context, id uint64) (*entity.SkuStock, error)
		// GetByDBOption 根据动态条件查询sku的库存
		GetByDBOption(ctx context.Context, pageNum uint32, pageSize uint32, opts ...db.DBOption) ([]*entity.SkuStock, uint32, error)

		// BatchCreateWithTX 创建sku的库存
		BatchCreateWithTX(ctx context.Context, productID uint64, skuStocks []*entity.SkuStock) error
		// BatchUpdateOrInsertSkuStock 批量插入或者更新
		BatchUpdateOrInsertSkuStock(ctx context.Context, stocks []*entity.SkuStock) error
	}
)

// ProductAttributeValue 产品参数信息
type (
	// IProductAttributeValueUseCase 业务逻辑
	IProductAttributeValueUseCase interface {
	}

	// IProductAttributeValueRepo 数据存储操作
	IProductAttributeValueRepo interface {
		// Create 创建产品参数信息
		Create(ctx context.Context, productAttributeValue *entity.ProductAttributeValue) error
		// DeleteByID 根据主键ID删除产品参数信息
		DeleteByID(ctx context.Context, id uint64) error
		// Update 修改产品参数信息
		Update(ctx context.Context, productAttributeValue *entity.ProductAttributeValue) error
		// GetByID 根据主键ID查询产品参数信息
		GetByID(ctx context.Context, id uint64) (*entity.ProductAttributeValue, error)
		// GetByDBOption 根据动态条件查询产品参数信息
		GetByDBOption(ctx context.Context, pageNum uint32, pageSize uint32, opts ...db.DBOption) ([]*entity.ProductAttributeValue, uint32, error)

		// BatchCreateWithTX 创建产品参数信息
		BatchCreateWithTX(ctx context.Context, productID uint64, productAttributeValues []*entity.ProductAttributeValue) error
	}
)

// SubjectProductRelation 专题商品关系
type (
	// ISubjectProductRelationUseCase 业务逻辑
	ISubjectProductRelationUseCase interface {
	}

	// ISubjectProductRelationRepo 数据存储操作
	ISubjectProductRelationRepo interface {
		// Create 创建专题商品关系
		Create(ctx context.Context, subjectProductRelation *entity.SubjectProductRelation) error
		// DeleteByID 根据主键ID删除专题商品关
		DeleteByID(ctx context.Context, id uint64) error
		// Update 修改专题商品关系
		Update(ctx context.Context, subjectProductRelation *entity.SubjectProductRelation) error
		// GetByID 根据主键ID查询专题商品关系
		GetByID(ctx context.Context, id uint64) (*entity.SubjectProductRelation, error)
		// GetByDBOption 根据动态条件查询专题商品关系
		GetByDBOption(ctx context.Context, pageNum uint32, pageSize uint32, opts ...db.DBOption) ([]*entity.SubjectProductRelation, uint32, error)

		// BatchCreateWithTX 创建专题商品关系
		BatchCreateWithTX(ctx context.Context, productID uint64, subjectProductRelations []*entity.SubjectProductRelation) error
	}
)

// PrefrenceAreaProductRelation 优选专区和产品关系
type (
	// IPrefrenceAreaProductRelationUseCase 业务逻辑
	IPrefrenceAreaProductRelationUseCase interface {
	}

	// IPrefrenceAreaProductRelationRepo 数据存储操作
	IPrefrenceAreaProductRelationRepo interface {
		// Create 创建优选专区和产品关系
		Create(ctx context.Context, prefrenceAreaProductRelation *entity.PrefrenceAreaProductRelation) error
		// DeleteByID 根据主键ID删除优选专区和产品关系
		DeleteByID(ctx context.Context, id uint64) error
		// Update 修改优选专区和产品关系
		Update(ctx context.Context, prefrenceAreaProductRelation *entity.PrefrenceAreaProductRelation) error
		// GetByID 根据主键ID查询优选专区和产品关系
		GetByID(ctx context.Context, id uint64) (*entity.PrefrenceAreaProductRelation, error)
		// GetByDBOption 根据动态条件查询优选专区和产品关系
		GetByDBOption(ctx context.Context, pageNum uint32, pageSize uint32, opts ...db.DBOption) ([]*entity.PrefrenceAreaProductRelation, uint32, error)

		// BatchCreateWithTX 创建优选专区和产品关系
		BatchCreateWithTX(ctx context.Context, productID uint64, prefrenceAreaProductRelations []*entity.PrefrenceAreaProductRelation) error
	}
)
