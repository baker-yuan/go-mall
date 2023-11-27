package repo

import (
	"context"
	"errors"

	"github.com/baker-yuan/go-mall/application/admin/internal/entity"
	"github.com/baker-yuan/go-mall/application/admin/pkg/db"
	"github.com/baker-yuan/go-mall/application/admin/pkg/util"
	"gorm.io/gorm"
)

// ProductAttributeValueRepo 产品参数信息
type ProductAttributeValueRepo struct {
	*db.GenericDao[entity.ProductAttributeValue, uint64]
}

// NewProductAttributeValueRepo 创建
func NewProductAttributeValueRepo(conn *gorm.DB) *ProductAttributeValueRepo {
	return &ProductAttributeValueRepo{
		GenericDao: &db.GenericDao[entity.ProductAttributeValue, uint64]{
			DB: conn,
		},
	}
}

func init() {
	registerInitField(initProductAttributeValueField)
}

var (
	// 全字段修改ProductAttributeValue那些字段不修改
	notUpdateProductAttributeValueField = []string{
		"created_at",
	}
	updateProductAttributeValueField []string
)

// InitProductAttributeValueField 全字段修改
func initProductAttributeValueField(db *gorm.DB) error {
	columnTypes, err := db.Migrator().ColumnTypes(&entity.ProductAttributeValue{})
	if err != nil {
		return err
	}
	columns := make([]string, 0)
	for _, v := range columnTypes {
		columns = append(columns, v.Name())
	}
	updateProductAttributeValueField = util.SliceRemove[string](columns, notUpdateProductAttributeValueField...)
	return nil
}

// Create 创建产品参数信息
func (r ProductAttributeValueRepo) Create(ctx context.Context, productAttributeValue *entity.ProductAttributeValue) error {
	if productAttributeValue.ID > 0 {
		return errors.New("illegal argument productAttributeValue id exist")
	}
	return r.GenericDao.Create(ctx, productAttributeValue)
}

// DeleteByID 根据主键ID删除产品参数信息
func (r ProductAttributeValueRepo) DeleteByID(ctx context.Context, id uint64) error {
	return r.GenericDao.DeleteByID(ctx, id)
}

// Update 修改产品参数信息
func (r ProductAttributeValueRepo) Update(ctx context.Context, productAttributeValue *entity.ProductAttributeValue) error {
	if productAttributeValue.ID == 0 {
		return errors.New("illegal argument productAttributeValue exist")
	}
	return r.GenericDao.DB.WithContext(ctx).Select(updateProductAttributeValueField).Updates(productAttributeValue).Error
}

// GetByID 根据主键ID查询产品参数信息
func (r ProductAttributeValueRepo) GetByID(ctx context.Context, id uint64) (*entity.ProductAttributeValue, error) {
	return r.GenericDao.GetByID(ctx, id)
}

// GetByDBOption 根据动态条件查询产品参数信息
func (r ProductAttributeValueRepo) GetByDBOption(ctx context.Context, pageNum uint32, pageSize uint32, opts ...db.DBOption) ([]*entity.ProductAttributeValue, uint32, error) {
	var (
		res       = make([]*entity.ProductAttributeValue, 0)
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

// BatchCreateWithTX 创建产品参数信息
func (r ProductAttributeValueRepo) BatchCreateWithTX(ctx context.Context, productID uint64, productAttributeValues []*entity.ProductAttributeValue) error {
	for _, productAttributeValue := range productAttributeValues {
		productAttributeValue.ProductID = productID
	}
	tdb, err := db.GetTransactionDB(ctx)
	if err != nil {
		return err
	}
	return tdb.WithContext(ctx).Create(productAttributeValues).Error
}

// DeleteByProductIDWithTX 根据商品ID删除记录
func (r ProductAttributeValueRepo) DeleteByProductIDWithTX(ctx context.Context, productID uint64) error {
	tdb, err := db.GetTransactionDB(ctx)
	if err != nil {
		return err
	}
	return tdb.WithContext(ctx).Where("product_id = ?", productID).Delete(&entity.ProductAttributeValue{}).Error
}
