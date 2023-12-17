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
		// HomeContent 获取首页内容
		HomeContent(ctx context.Context, req *pb.HomeContentReq) (*pb.HomeContentRsp, error)
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
		MemberInfo(ctx context.Context, memberID uint64) (*pb.MemberInfoRsp, error)
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
		// GetByIDs 根据主键ID集合查询会员表
		GetByIDs(ctx context.Context, ids []uint64) (entity.Members, error)
		// GetByDBOption 根据动态条件查询会员表
		GetByDBOption(ctx context.Context, pageNum uint32, pageSize uint32, opts ...db.DBOption) ([]*entity.Member, uint32, error)

		// GetByUsername 根据用户名查询会员表
		GetByUsername(ctx context.Context, username string) (*entity.Member, error)
		// GetByMemberID 根据用户id查询会员表
		GetByMemberID(ctx context.Context, memberID uint64) (*entity.Member, error)
	}
)

// CartItem 购物车
type (
	// ICartItemUseCase 业务逻辑
	ICartItemUseCase interface {
		// CartItemAdd 查询购物车中是否包含该商品，有增加数量，无添加到购物车
		CartItemAdd(ctx context.Context, memberID uint64, req *pb.CartItemAddReq) error
		// CartItemList 获取当前会员的购物车列表
		CartItemList(ctx context.Context, memberID uint64) ([]*pb.CartItem, error)
		// CartItemListPromotion 获取当前会员的购物车列表,包括促销信息
		CartItemListPromotion(context.Context, *pb.CartItemListPromotionReq) (*pb.CartItemListPromotionRsp, error)
		// CartItemUpdateQuantity 修改购物车中指定商品的数量
		CartItemUpdateQuantity(context.Context, *pb.CartItemUpdateQuantityReq) (*pb.CartItemUpdateQuantityRsp, error)
		// CartItemGetCartProduct 获取购物车中指定商品的规格,用于重选规格
		CartItemGetCartProduct(context.Context, *pb.CartItemGetCartProductReq) (*pb.CartItemGetCartProductRsp, error)
		// CartItemUpdateAttr 修改购物车中商品的规格
		CartItemUpdateAttr(context.Context, *pb.CartItemUpdateAttrReq) (*pb.CartItemUpdateAttrRsp, error)
		// CartItemDelete 删除购物车中的指定商品
		CartItemDelete(context.Context, *pb.CartItemDeleteReq) (*pb.CartItemDeleteRsp, error)
		// CartItemClear 清空当前会员的购物车
		CartItemClear(context.Context, *pb.CartItemClearReq) (*pb.CartItemClearRsp, error)
	}

	// ICartItemRepo 数据存储操作
	ICartItemRepo interface {
		// Create 创建购物车表
		Create(ctx context.Context, cartItem *entity.CartItem) error
		// DeleteByID 根据主键ID删除购物车表
		DeleteByID(ctx context.Context, id uint64) error
		// Update 修改购物车表
		Update(ctx context.Context, cartItem *entity.CartItem) error
		// GetByID 根据主键ID查询购物车表
		GetByID(ctx context.Context, id uint64) (*entity.CartItem, error)
		// GetByDBOption 根据动态条件查询购物车表
		GetByDBOption(ctx context.Context, pageNum uint32, pageSize uint32, opts ...db.DBOption) ([]*entity.CartItem, uint32, error)
		// GetEffectCartItemByMemberID 根据会员id查询购物车
		GetEffectCartItemByMemberID(ctx context.Context, memberID uint64) (entity.CartItems, error)

		// GetCartItem 根据会员id，商品id和规格获取购物车中商品
		GetCartItem(ctx context.Context, memberID uint64, productId uint64, productSkuID uint64) (*entity.CartItem, error)
	}
)

// HomeAdvertise 首页轮播广告表
type (
	// IHomeAdvertiseUseCase 业务逻辑
	IHomeAdvertiseUseCase interface {
	}

	// IHomeAdvertiseRepo 数据存储操作
	IHomeAdvertiseRepo interface {
		// GetHomeAdvertises 获取首页广告
		GetHomeAdvertises(ctx context.Context) ([]*entity.HomeAdvertise, error)
	}
)

// GetRecommendBrandList 获取推荐品牌
//GetRecommendBrandList(ctx context.Context, offset uint32, limit uint32) ([]*entity.Brand, error)
