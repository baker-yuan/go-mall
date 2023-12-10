package repo

import (
	"context"

	db "github.com/baker-yuan/go-mall/common/db"
	"github.com/baker-yuan/go-mall/common/entity"
	"gorm.io/gorm"
)

// ProductCategoryRepo 商品分类表
type ProductCategoryRepo struct {
	*db.GenericDao[entity.ProductCategory, uint64]
}

// NewProductCategoryRepo 创建
func NewProductCategoryRepo(conn *gorm.DB) *ProductCategoryRepo {
	return &ProductCategoryRepo{
		GenericDao: &db.GenericDao[entity.ProductCategory, uint64]{
			DB: conn,
		},
	}
}

// GetByParentID 根据上级分类的编号查询商品分类
func (p ProductCategoryRepo) GetByParentID(ctx context.Context, parentID uint64) (entity.ProductCategories, error) {
	res := make([]*entity.ProductCategory, 0)
	if err := p.GenericDao.DB.WithContext(ctx).Where("parent_id = ?", parentID).Find(&res).Error; err != nil {
		return nil, err
	}
	return res, nil
}
