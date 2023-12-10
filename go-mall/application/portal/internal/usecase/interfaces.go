package usecase

import (
	"context"

	"github.com/baker-yuan/go-mall/common/entity"
	pb "github.com/baker-yuan/go-mall/proto/mall"
)

// ProductCategory 商品分类表
type (
	// IProductCategoryRepo 数据存储操作
	IProductCategoryRepo interface {
		// GetByParentID 根据上级分类的编号查询商品分类
		GetByParentID(ctx context.Context, parentID uint64) (entity.ProductCategories, error)
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
