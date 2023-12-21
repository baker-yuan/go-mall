package repo

import (
	"context"

	db "github.com/baker-yuan/go-mall/common/db"
	"github.com/baker-yuan/go-mall/common/entity"
	"gorm.io/gorm"
)

// SkuStockRepo sku的库存
type SkuStockRepo struct {
	*db.GenericDao[entity.SkuStock, uint64]
}

// NewSkuStockRepo 创建
func NewSkuStockRepo(conn *gorm.DB) *SkuStockRepo {
	return &SkuStockRepo{
		GenericDao: &db.GenericDao[entity.SkuStock, uint64]{
			DB: conn,
		},
	}
}

// GetByProductID 根据商品ID查询sku的库存
func (r SkuStockRepo) GetByProductID(ctx context.Context, productID uint64) (entity.SkuStocks, error) {
	var (
		res = make([]*entity.SkuStock, 0)
	)
	if err := r.GenericDao.DB.WithContext(ctx).
		Where("product_id = ?", productID).
		Order("id desc").Find(&res).Error; err != nil {
		return nil, err
	}
	return res, nil
}

// GetByProductIDs 根据商品ID查询sku的库存
func (r SkuStockRepo) GetByProductIDs(ctx context.Context, productIDs []uint64) (entity.SkuStocks, error) {
	var (
		res = make([]*entity.SkuStock, 0)
	)
	if err := r.GenericDao.DB.WithContext(ctx).
		Where("product_id in ?", productIDs).
		Order("id desc").Find(&res).Error; err != nil {
		return nil, err
	}
	return res, nil
}
