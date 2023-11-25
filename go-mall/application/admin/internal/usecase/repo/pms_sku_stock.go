package repo

import (
	"context"
	"errors"

	"github.com/baker-yuan/go-mall/application/admin/internal/entity"
	"github.com/baker-yuan/go-mall/application/admin/pkg/db"
	"github.com/baker-yuan/go-mall/application/admin/pkg/util"
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

func init() {
	registerInitField(initSkuStockField)
}

var (
	// 全字段修改SkuStock那些字段不修改
	notUpdateSkuStockField = []string{
		"created_at",
	}
	updateSkuStockField []string
)

// InitSkuStockField 全字段修改
func initSkuStockField(db *gorm.DB) error {
	columnTypes, err := db.Migrator().ColumnTypes(&entity.SkuStock{})
	if err != nil {
		return err
	}
	columns := make([]string, 0)
	for _, v := range columnTypes {
		columns = append(columns, v.Name())
	}
	updateSkuStockField = util.SliceRemove[string](columns, notUpdateSkuStockField...)
	return nil
}

// Create 创建sku的库存
func (r SkuStockRepo) Create(ctx context.Context, skuStock *entity.SkuStock) error {
	if skuStock.ID > 0 {
		return errors.New("illegal argument skuStock id exist")
	}
	return r.GenericDao.Create(ctx, skuStock)
}

// DeleteByID 根据主键ID删除sku的库存
func (r SkuStockRepo) DeleteByID(ctx context.Context, id uint64) error {
	return r.GenericDao.DeleteByID(ctx, id)
}

// Update 修改sku的库存
func (r SkuStockRepo) Update(ctx context.Context, skuStock *entity.SkuStock) error {
	if skuStock.ID == 0 {
		return errors.New("illegal argument skuStock exist")
	}
	return r.GenericDao.DB.WithContext(ctx).Select(updateSkuStockField).Updates(skuStock).Error
}

// GetByID 根据主键ID查询sku的库存
func (r SkuStockRepo) GetByID(ctx context.Context, id uint64) (*entity.SkuStock, error) {
	return r.GenericDao.GetByID(ctx, id)
}

// GetByDBOption 根据动态条件查询sku的库存
func (r SkuStockRepo) GetByDBOption(ctx context.Context, pageNum uint32, pageSize uint32, opts ...db.DBOption) ([]*entity.SkuStock, uint32, error) {
	var (
		res       = make([]*entity.SkuStock, 0)
		pageTotal = int64(0)
		offset    = (pageNum - 1) * pageSize
	)

	session := r.GenericDao.DB.WithContext(ctx)
	for _, opt := range opts {
		session = opt(session)
	}

	session = session.Offset(int(offset)).Limit(int(pageSize)).Order("id desc").Find(&res).
		Offset(-1).Limit(-1).Count(&pageTotal)

	if err := session.Error; err != nil {
		return nil, 0, err
	}
	return res, uint32(pageTotal), nil
}

// BatchCreateWithTX 创建sku库存
func (r SkuStockRepo) BatchCreateWithTX(ctx context.Context, productID uint64, skuStocks []*entity.SkuStock) error {
	for _, skuStock := range skuStocks {
		skuStock.ProductID = productID
	}
	db, err := db.GetDbToCtx(ctx)
	if err != nil {
		return err
	}
	return db.WithContext(ctx).Create(skuStocks).Error
}
