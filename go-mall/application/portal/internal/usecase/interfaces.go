package usecase

import (
	"context"

	"github.com/baker-yuan/go-mall/common/db"
	"github.com/baker-yuan/go-mall/common/entity"
	pb "github.com/baker-yuan/go-mall/proto/mall"
)

// ProductCategory 商品分类表
type (
	// IProductCategoryRepo 数据存储操作
	IProductCategoryRepo interface {
		WithByParentID(parentID uint64) db.DBOption
		WithByShowStatus(showStatus uint8) db.DBOption

		// GetShowProductCategory 根据上级分类的编号查询商品分类
		GetShowProductCategory(ctx context.Context, parentID uint64) (entity.ProductCategories, error)
	}
)

// Brand 商品品牌
type (
	// IBrandUseCase 业务逻辑
	IBrandUseCase interface {
	}

	// IBrandRepo 数据存储操作
	IBrandRepo interface {
		// GetByID 根据主键ID查询商品品牌表
		GetByID(ctx context.Context, id uint64) (*entity.Brand, error)
	}
)

// IHomeUseCase 首页
type (
	// IHomeUseCase 业务逻辑
	IHomeUseCase interface {
		// ProductCategoryList 分页查询订单
		ProductCategoryList(context.Context, *pb.ProductCategoryListReq) ([]*pb.ProductCategoryItem, error)
	}
)

// PmsProduct 商品信息
type (
	// IProductUseCase 业务逻辑
	IProductUseCase interface {
		// SearchProduct 综合搜索商品
		SearchProduct(ctx context.Context, req *pb.SearchProductReq) ([]*pb.SearchProductRsp_Product, error)
		// ProductDetail 获取前台商品详情
		ProductDetail(ctx context.Context, req *pb.ProductDetailReq) (*pb.ProductDetailRsp, error)
	}

	// IProductRepo 数据存储操作
	IProductRepo interface {
		// GetByID 根据主键ID查询商品
		GetByID(ctx context.Context, id uint64) (*entity.Product, error)

		// SearchProduct 综合搜索商品
		SearchProduct(ctx context.Context, req *pb.SearchProductReq) (entity.Products, error)
	}
)

// ProductAttribute 商品属性参数
type (
	// IProductAttributeUseCase 业务逻辑
	IProductAttributeUseCase interface {
	}

	// IProductAttributeRepo 数据存储操作
	IProductAttributeRepo interface {
		// GetProductAttributeByCategoryID 根据产品属性分类表ID获取商品属性参数表
		GetProductAttributeByCategoryID(ctx context.Context, productAttributeCategoryID uint64) (entity.ProductAttributes, error)
	}
)

// ProductAttributeValue 产品参数信息
type (
	// IProductAttributeValueUseCase 业务逻辑
	IProductAttributeValueUseCase interface {
	}

	// IProductAttributeValueRepo 数据存储操作
	IProductAttributeValueRepo interface {
		// GetByProductAttributeID 根据productAttributeID查询产品参数信息
		GetByProductAttributeID(ctx context.Context, productID uint64, productAttributeIDs []uint64) (entity.ProductAttributeValues, error)
	}
)

// SkuStock sku的库存
type (
	// ISkuStockUseCase 业务逻辑
	ISkuStockUseCase interface {
	}

	// ISkuStockRepo 数据存储操作
	ISkuStockRepo interface {
		// GetByProductID 根据商品ID查询sku的库存
		GetByProductID(ctx context.Context, productID uint64) (entity.SkuStocks, error)
	}
)

// Member 会员
type (
	// IMemberUseCase 业务逻辑
	IMemberUseCase interface {
		// MemberRegister 会员注册
		MemberRegister(ctx context.Context, req *pb.MemberRegisterReq) (*pb.EmptyRsp, error)
		// MemberLogin 会员登录
		MemberLogin(ctx context.Context, req *pb.MemberLoginReq) (*pb.MemberLoginRsp, error)
		// MemberInfo 获取会员信息
		MemberInfo(ctx context.Context, username string) (*pb.MemberInfoRsp, error)
		// MemberGetAuthCode 获取验证码
		MemberGetAuthCode(ctx context.Context, req *pb.MemberGetAuthCodeReq) (*pb.MemberGetAuthCodeRsp, error)
		// MemberUpdatePassword 修改密码
		MemberUpdatePassword(ctx context.Context, req *pb.MemberUpdatePasswordReq) (*pb.MemberUpdatePasswordRsp, error)
		// MemberRefreshToken 刷新token
		MemberRefreshToken(ctx context.Context, req *pb.MemberRefreshTokenReq) (*pb.MemberRefreshTokenRsp, error)
	}

	// IMemberRepo 数据存储操作
	IMemberRepo interface {
		// Create 创建会员表
		Create(ctx context.Context, member *entity.Member) error
		// DeleteByID 根据主键ID删除会员表
		DeleteByID(ctx context.Context, id uint64) error
		// Update 修改会员表
		Update(ctx context.Context, member *entity.Member) error
		// GetByID 根据主键ID查询会员表
		GetByID(ctx context.Context, id uint64) (*entity.Member, error)
		// GetByDBOption 根据动态条件查询会员表
		GetByDBOption(ctx context.Context, pageNum uint32, pageSize uint32, opts ...db.DBOption) ([]*entity.Member, uint32, error)

		// GetByUsername 根据用户名查询会员表
		GetByUsername(ctx context.Context, username string) (*entity.Member, error)
	}
)
