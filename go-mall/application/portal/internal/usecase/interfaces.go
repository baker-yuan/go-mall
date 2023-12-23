package usecase

import (
	"context"

	portal_entity "github.com/baker-yuan/go-mall/application/portal/internal/entity"
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
		// GetByIDs 根据主键ID查询商品
		GetByIDs(ctx context.Context, ids []uint64) (entity.Products, error)

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
		// GetByProductIDs 根据商品ID查询sku的库存
		GetByProductIDs(ctx context.Context, productIDs []uint64) (entity.SkuStocks, error)
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
		CartItemListPromotion(ctx context.Context, memberID uint64, cartIDs []uint64) ([]*pb.CartItemListPromotion, error)
		// CartItemUpdateQuantity 修改购物车中指定商品的数量
		CartItemUpdateQuantity(ctx context.Context, memberID uint64, req *pb.CartItemUpdateQuantityReq) error
		// CartItemGetCartProduct 获取购物车中指定商品的规格,用于重选规格
		CartItemGetCartProduct(ctx context.Context, memberID uint64, req *pb.CartItemGetCartProductReq) (*pb.CartItemGetCartProductRsp, error)
		// CartItemUpdateAttr 修改购物车中商品的规格
		CartItemUpdateAttr(ctx context.Context, memberID uint64, req *pb.CartItemUpdateAttrReq) (*pb.CartItemUpdateAttrRsp, error)
		// CartItemDelete 删除购物车中的指定商品
		CartItemDelete(ctx context.Context, memberID uint64, req *pb.CartItemDeleteReq) error
		// CartItemClear 清空当前会员的购物车
		CartItemClear(ctx context.Context, memberID uint64) error
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

		// SecurityGetByIDS 根据会员id和主键ID查询购物车表
		SecurityGetByIDS(ctx context.Context, memberID uint64, cartIDs []uint64) (entity.CartItems, error)
		// GetEffectCartItemByMemberID 根据会员id查询购物车
		GetEffectCartItemByMemberID(ctx context.Context, memberID uint64) (entity.CartItems, error)

		// GetCartItem 根据会员id，商品id和规格获取购物车中商品
		GetCartItem(ctx context.Context, memberID uint64, productId uint64, productSkuID uint64) (*entity.CartItem, error)
		// CartItemClear 清空当前会员的购物车
		CartItemClear(ctx context.Context, memberID uint64) error
		// CartItemDelete 批量删除购物车中的商品
		CartItemDelete(ctx context.Context, memberID uint64, ids []uint64) error
		//CartItemUpdateQuantity 修改购物车中指定商品的数量
		CartItemUpdateQuantity(ctx context.Context, memberID uint64, id uint64, quantity uint32) error
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

// Order 订单表
type (
	// IOrderUseCase 业务逻辑
	IOrderUseCase interface {
		// GenerateConfirmOrder 根据用户购物车信息生成确认单信息
		GenerateConfirmOrder(ctx context.Context, memberID uint64, req *pb.GenerateConfirmOrderReq) (*pb.GenerateConfirmOrderRsp, error)
		// GenerateOrder 根据提交信息生成订单
		GenerateOrder(ctx context.Context, req *pb.GenerateOrderReq) (*pb.GenerateOrderRsp, error)
		// PaySuccess 支付成功后的回调
		PaySuccess(ctx context.Context, req *pb.PaySuccessReq) (*pb.PaySuccessRsp, error)
		// CancelTimeOutOrder PaySuccess 自动取消超时订单
		CancelTimeOutOrder(ctx context.Context, req *pb.CancelTimeOutOrderReq) (*pb.CancelTimeOutOrderRsp, error)
		// CancelOrder 取消单个超时订单
		CancelOrder(ctx context.Context, req *pb.CancelOrderReq) (*pb.CancelOrderRsp, error)
		// OrderList 按状态分页获取用户订单列表
		OrderList(ctx context.Context, req *pb.OrderListReq) (*pb.OrderListRsp, error)
		// OrderDetail 根据ID获取订单详情
		OrderDetail(ctx context.Context, req *pb.OrderDetailReq) (*pb.OrderDetailRsp, error)
		// CancelUserOrder 用户取消订单
		CancelUserOrder(ctx context.Context, req *pb.CancelUserOrderReq) (*pb.CancelUserOrderRsp, error)
		// ConfirmReceiveOrder 用户确认收货
		ConfirmReceiveOrder(ctx context.Context, req *pb.ConfirmReceiveOrderReq) (*pb.ConfirmReceiveOrderRsp, error)
		// DeleteOrder 用户删除订单
		DeleteOrder(ctx context.Context, req *pb.PortalDeleteOrderReq) (*pb.PortalDeleteOrderRsp, error)
	}

	// IOrderRepo 数据存储操作
	IOrderRepo interface {
		// Create 创建订单表
		Create(ctx context.Context, order *entity.Order) error
		// DeleteByID 根据主键ID删除订单表
		DeleteByID(ctx context.Context, id uint64) error
		// Update 修改订单表
		Update(ctx context.Context, order *entity.Order) error
		// GetByID 根据主键ID查询订单表
		GetByID(ctx context.Context, id uint64) (*entity.Order, error)
		// GetByDBOption 根据动态条件查询订单表
		GetByDBOption(ctx context.Context, pageNum uint32, pageSize uint32, opts ...db.DBOption) ([]*entity.Order, uint32, error)
	}
)

// MemberReceiveAddress 会员收货地址表
type (
	// IMemberReceiveAddressUseCase 业务逻辑
	IMemberReceiveAddressUseCase interface {
	}

	// IMemberReceiveAddressRepo 数据存储操作
	IMemberReceiveAddressRepo interface {
		// Create 创建会员收货地址表
		Create(ctx context.Context, memberReceiveAddress *entity.MemberReceiveAddress) error
		// DeleteByID 根据主键ID删除会员收货地址表
		DeleteByID(ctx context.Context, id uint64) error
		// Update 修改会员收货地址表
		Update(ctx context.Context, memberReceiveAddress *entity.MemberReceiveAddress) error
		// GetByID 根据主键ID查询会员收货地址表
		GetByID(ctx context.Context, id uint64) (*entity.MemberReceiveAddress, error)
		// GetByDBOption 根据动态条件查询会员收货地址表
		GetByDBOption(ctx context.Context, pageNum uint32, pageSize uint32, opts ...db.DBOption) ([]*entity.MemberReceiveAddress, uint32, error)

		// GetByMemberID 根据会员ID查找
		GetByMemberID(ctx context.Context, memberID uint64) (entity.MemberReceiveAddresses, error)
	}
)

// MemberReceiveAddress 促销管理
type (
	// IPromotionUseCase 业务逻辑
	IPromotionUseCase interface {
		// CalcCartPromotion 计算购物车中的促销活动信息
		// cartItems 购物车
		CalcCartPromotion(ctx context.Context, cartItems entity.CartItems) (portal_entity.CartPromotionItems, error)
	}

	// IPromotionRepo 数据存储操作
	IPromotionRepo interface {
	}
)

// ProductLadder 商品阶梯价格表
type (
	// IProductLadderUseCase 业务逻辑
	IProductLadderUseCase interface {
	}

	// IProductLadderRepo 数据存储操作
	IProductLadderRepo interface {
		// Create 创建商品阶梯价格表
		Create(ctx context.Context, productLadder *entity.ProductLadder) error
		// DeleteByID 根据主键ID删除商品阶梯价格表
		DeleteByID(ctx context.Context, id uint64) error
		// Update 修改商品阶梯价格表
		Update(ctx context.Context, productLadder *entity.ProductLadder) error
		// GetByID 根据主键ID查询商品阶梯价格表
		GetByID(ctx context.Context, id uint64) (*entity.ProductLadder, error)
		// GetByDBOption 根据动态条件查询商品阶梯价格表
		GetByDBOption(ctx context.Context, pageNum uint32, pageSize uint32, opts ...db.DBOption) ([]*entity.ProductLadder, uint32, error)

		// GetByProductIDs 根据商品ID查询商品阶梯价格
		GetByProductIDs(ctx context.Context, productIDs []uint64) (entity.ProductLadders, error)
	}
)

// ProductFullReduction 商品满减表
type (
	// IProductFullReductionUseCase 业务逻辑
	IProductFullReductionUseCase interface {
	}

	// IProductFullReductionRepo 数据存储操作
	IProductFullReductionRepo interface {
		// Create 创建商品满减表
		Create(ctx context.Context, productFullReduction *entity.ProductFullReduction) error
		// DeleteByID 根据主键ID删除商品满减表
		DeleteByID(ctx context.Context, id uint64) error
		// Update 修改商品满减表
		Update(ctx context.Context, productFullReduction *entity.ProductFullReduction) error
		// GetByID 根据主键ID查询商品满减表
		GetByID(ctx context.Context, id uint64) (*entity.ProductFullReduction, error)
		// GetByDBOption 根据动态条件查询商品满减表
		GetByDBOption(ctx context.Context, pageNum uint32, pageSize uint32, opts ...db.DBOption) ([]*entity.ProductFullReduction, uint32, error)

		// GetByProductIDs 根据商品ID查询商品满减表
		GetByProductIDs(ctx context.Context, productIDs []uint64) (entity.ProductFullReductions, error)
	}
)

// JsonDynamicConfig JSON动态配置
type (
	// IJsonDynamicConfigUseCase 业务逻辑
	IJsonDynamicConfigUseCase interface {
	}

	// IJsonDynamicConfigRepo 数据存储操作
	IJsonDynamicConfigRepo interface {
		// Create 创建JSON动态配置
		Create(ctx context.Context, jsonDynamicConfig *entity.JsonDynamicConfig) error
		// DeleteByID 根据主键ID删除JSON动态配置
		DeleteByID(ctx context.Context, id uint64) error
		// Update 修改JSON动态配置
		Update(ctx context.Context, jsonDynamicConfig *entity.JsonDynamicConfig) error
		// GetByID 根据主键ID查询JSON动态配置
		GetByID(ctx context.Context, id uint64) (*entity.JsonDynamicConfig, error)
		// GetByDBOption 根据动态条件查询JSON动态配置
		GetByDBOption(ctx context.Context, pageNum uint32, pageSize uint32, opts ...db.DBOption) ([]*entity.JsonDynamicConfig, uint32, error)
	}
)
