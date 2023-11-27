package repo

import (
	"context"
	"errors"

	"github.com/baker-yuan/go-mall/application/admin/internal/entity"
	"github.com/baker-yuan/go-mall/application/admin/pkg/db"
	"github.com/baker-yuan/go-mall/application/admin/pkg/util"
	"gorm.io/gorm"
)

// PrefrenceAreaProductRelationRepo 优选专区和产品关系表
type PrefrenceAreaProductRelationRepo struct {
	*db.GenericDao[entity.PrefrenceAreaProductRelation, uint64]
}

// NewPrefrenceAreaProductRelationRepo 创建
func NewPrefrenceAreaProductRelationRepo(conn *gorm.DB) *PrefrenceAreaProductRelationRepo {
	return &PrefrenceAreaProductRelationRepo{
		GenericDao: &db.GenericDao[entity.PrefrenceAreaProductRelation, uint64]{
			DB: conn,
		},
	}
}

func init() {
	registerInitField(initPrefrenceAreaProductRelationField)
}

var (
	// 全字段修改PrefrenceAreaProductRelation那些字段不修改
	notUpdatePrefrenceAreaProductRelationField = []string{
		"created_at",
	}
	updatePrefrenceAreaProductRelationField []string
)

// InitPrefrenceAreaProductRelationField 全字段修改
func initPrefrenceAreaProductRelationField(db *gorm.DB) error {
	columnTypes, err := db.Migrator().ColumnTypes(&entity.PrefrenceAreaProductRelation{})
	if err != nil {
		return err
	}
	columns := make([]string, 0)
	for _, v := range columnTypes {
		columns = append(columns, v.Name())
	}
	updatePrefrenceAreaProductRelationField = util.SliceRemove[string](columns, notUpdatePrefrenceAreaProductRelationField...)
	return nil
}

// Create 创建优选专区和产品关系表
func (r PrefrenceAreaProductRelationRepo) Create(ctx context.Context, prefrenceAreaProductRelation *entity.PrefrenceAreaProductRelation) error {
	if prefrenceAreaProductRelation.ID > 0 {
		return errors.New("illegal argument prefrenceAreaProductRelation id exist")
	}
	return r.GenericDao.Create(ctx, prefrenceAreaProductRelation)
}

// DeleteByID 根据主键ID删除优选专区和产品关系表
func (r PrefrenceAreaProductRelationRepo) DeleteByID(ctx context.Context, id uint64) error {
	return r.GenericDao.DeleteByID(ctx, id)
}

// Update 修改优选专区和产品关系表
func (r PrefrenceAreaProductRelationRepo) Update(ctx context.Context, prefrenceAreaProductRelation *entity.PrefrenceAreaProductRelation) error {
	if prefrenceAreaProductRelation.ID == 0 {
		return errors.New("illegal argument prefrenceAreaProductRelation exist")
	}
	return r.GenericDao.DB.WithContext(ctx).Select(updatePrefrenceAreaProductRelationField).Updates(prefrenceAreaProductRelation).Error
}

// GetByID 根据主键ID查询优选专区和产品关系表
func (r PrefrenceAreaProductRelationRepo) GetByID(ctx context.Context, id uint64) (*entity.PrefrenceAreaProductRelation, error) {
	return r.GenericDao.GetByID(ctx, id)
}

// GetByDBOption 根据动态条件查询优选专区和产品关系
func (r PrefrenceAreaProductRelationRepo) GetByDBOption(ctx context.Context, pageNum uint32, pageSize uint32, opts ...db.DBOption) ([]*entity.PrefrenceAreaProductRelation, uint32, error) {
	var (
		res       = make([]*entity.PrefrenceAreaProductRelation, 0)
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

// BatchCreateWithTX 创建优选专区和产品关系
func (r PrefrenceAreaProductRelationRepo) BatchCreateWithTX(ctx context.Context, productID uint64, prefrenceAreaProductRelations []*entity.PrefrenceAreaProductRelation) error {
	for _, prefrenceAreaProductRelation := range prefrenceAreaProductRelations {
		prefrenceAreaProductRelation.ProductID = productID
	}
	tdb, err := db.GetTransactionDB(ctx)
	if err != nil {
		return err
	}
	return tdb.WithContext(ctx).Create(prefrenceAreaProductRelations).Error
}

// DeleteByProductIDWithTX 根据商品ID删除记录
func (r PrefrenceAreaProductRelationRepo) DeleteByProductIDWithTX(ctx context.Context, productID uint64) error {
	tdb, err := db.GetTransactionDB(ctx)
	if err != nil {
		return err
	}
	return tdb.WithContext(ctx).Where("product_id = ?", productID).Delete(&entity.PrefrenceAreaProductRelation{}).Error
}
