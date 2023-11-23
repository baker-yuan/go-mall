package repo

import (
	"context"
	"errors"

	"github.com/baker-yuan/go-mall/application/admin/internal/entity"
	"github.com/baker-yuan/go-mall/application/admin/pkg/db"
	"github.com/baker-yuan/go-mall/application/admin/pkg/util"
	"gorm.io/gorm"
)

// ProductRepo 商品信息表
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

func init() {
	registerInitField(initProductField)
}

var (
	// 全字段修改Product那些字段不修改
	notUpdateProductField = []string{
		"created_at",
	}
	updateProductField []string
)

// InitProductField 全字段修改
func initProductField(db *gorm.DB) error {
	columnTypes, err := db.Migrator().ColumnTypes(&entity.Product{})
	if err != nil {
		return err
	}
	columns := make([]string, 0)
	for _, v := range columnTypes {
		columns = append(columns, v.Name())
	}
	updateProductField = util.SliceRemove[string](columns, notUpdateProductField...)
	return nil
}

// Create 创建商品
func (r ProductRepo) Create(ctx context.Context, product *entity.Product) error {
	if product.ID > 0 {
		return errors.New("illegal argument product id exist")
	}
	return r.GenericDao.Create(ctx, product)
}

// DeleteByID 根据主键ID删除商品
func (r ProductRepo) DeleteByID(ctx context.Context, id uint64) error {
	return r.GenericDao.DeleteByID(ctx, id)
}

// Update 修改商品
func (r ProductRepo) Update(ctx context.Context, product *entity.Product) error {
	if product.ID == 0 {
		return errors.New("illegal argument product exist")
	}
	return r.GenericDao.DB.WithContext(ctx).Select(updateProductField).Updates(product).Error
}

// GetByID 根据主键ID查询商品
func (r ProductRepo) GetByID(ctx context.Context, id uint64) (*entity.Product, error) {
	return r.GenericDao.GetByID(ctx, id)
}

// GetByDBOption 根据动态条件查询商品
func (r ProductRepo) GetByDBOption(ctx context.Context, pageNum uint32, pageSize uint32, opts ...db.DBOption) ([]*entity.Product, uint32, error) {
	var (
		res       = make([]*entity.Product, 0)
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

func (r ProductRepo) UpdateProductCategoryNameByProductCategoryIDWithTX(ctx context.Context, productCategoryID uint64, productCategoryName string) error {
	db, err := db.GetDbToCtx(ctx)
	if err != nil {
		return err
	}

	return db.WithContext(ctx).Model(&entity.Product{}).
		Where("product_category_id = ?", productCategoryID).
		Update("product_category_name", productCategoryName).Error
}
