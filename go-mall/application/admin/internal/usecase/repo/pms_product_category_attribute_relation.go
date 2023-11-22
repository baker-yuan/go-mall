package repo

import (
	"context"

	"github.com/baker-yuan/go-mall/application/admin/internal/entity"
	"github.com/baker-yuan/go-mall/application/admin/pkg/db"
	"gorm.io/gorm"
)

// PmsProductCategoryAttributeRelationRepo 产品的分类和属性的关系表，用于设置分类筛选条件（只支持一级分类）
type PmsProductCategoryAttributeRelationRepo struct {
	*db.GenericDao[entity.ProductCategoryAttributeRelation, uint64]
}

// NewProductCategoryAttributeRelationRepo 创建
func NewProductCategoryAttributeRelationRepo(conn *gorm.DB) *PmsProductCategoryAttributeRelationRepo {
	return &PmsProductCategoryAttributeRelationRepo{
		GenericDao: &db.GenericDao[entity.ProductCategoryAttributeRelation, uint64]{
			DB: conn,
		},
	}
}

// Create 创建
func (p PmsProductCategoryAttributeRelationRepo) Create(ctx context.Context, relation *entity.ProductCategoryAttributeRelation) error {
	return p.GenericDao.Create(ctx, relation)
}

// DeleteByID 根据主键ID删除
func (p PmsProductCategoryAttributeRelationRepo) DeleteByID(ctx context.Context, id uint64) error {
	return p.GenericDao.DeleteByID(ctx, id)
}

// Update 修改商品
func (p PmsProductCategoryAttributeRelationRepo) Update(ctx context.Context, relation *entity.ProductCategoryAttributeRelation) error {
	return p.GenericDao.UpdateByID(ctx, relation)
}

// GetByID 根据主键ID查询
func (p PmsProductCategoryAttributeRelationRepo) GetByID(ctx context.Context, id uint64) (*entity.ProductCategoryAttributeRelation, error) {
	return p.GenericDao.GetByID(ctx, id)
}

// CreateWithTX 创建
func (p PmsProductCategoryAttributeRelationRepo) CreateWithTX(ctx context.Context, relation *entity.ProductCategoryAttributeRelation) error {
	db, err := db.GetDbToCtx(ctx)
	if err != nil {
		return err
	}
	return db.WithContext(ctx).Create(relation).Error
}

// BatchCreateWithTX 批量插入商品分类与筛选属性关系表
// attributeIds 商品分类id
// productAttributeIDList 相关商品筛选属性id集合
func (p PmsProductCategoryAttributeRelationRepo) BatchCreateWithTX(ctx context.Context, productCategoryID uint64, attributeIds []uint64) error {
	relations := make([]*entity.ProductCategoryAttributeRelation, 0)
	for _, productAttrID := range attributeIds {
		relations = append(relations, &entity.ProductCategoryAttributeRelation{
			ProductCategoryID:  productCategoryID,
			ProductAttributeID: productAttrID,
		})
	}

	db, err := db.GetDbToCtx(ctx)
	if err != nil {
		return err
	}
	return db.WithContext(ctx).Create(relations).Error
}

func (p PmsProductCategoryAttributeRelationRepo) DeleteByProductCategoryIDWithTX(ctx context.Context, productCategoryID uint64) error {
	db, err := db.GetDbToCtx(ctx)
	if err != nil {
		return err
	}
	return db.WithContext(ctx).
		Where("product_category_id = ?", productCategoryID).
		Delete(&entity.ProductCategoryAttributeRelation{}).Error
}
