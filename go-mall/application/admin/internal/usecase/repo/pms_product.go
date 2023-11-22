package repo

import (
	"context"

	"github.com/baker-yuan/go-mall/application/admin/internal/entity"
	"github.com/baker-yuan/go-mall/application/admin/pkg/db"
	"gorm.io/gorm"
)

// ProductRepo 商品分类表
type ProductRepo struct {
	*db.GenericDao[entity.Product, uint64]
}

// NewProductRepo 创建
func NewProductRepo(conn *gorm.DB) *ProductRepo {
	return &ProductRepo{
		GenericDao: &db.GenericDao[entity.Product, uint64]{
			DB: conn,
		},
	}
}

func (r ProductRepo) UpdateProductCategoryNameByProductCategoryIDWithTX(ctx context.Context, productCategoryID uint64, productCategoryName string) error {
	db, err := db.GetDbToCtx(ctx)
	if err != nil {
		return err
	}

	return db.WithContext(ctx).Model(&entity.Product{}).
		Where("product_category_id = ?", productCategoryID).
		Update("product_category_name", productCategoryName).Error
}
