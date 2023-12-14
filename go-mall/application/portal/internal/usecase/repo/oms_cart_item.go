package repo

import (
	"context"
	"errors"

	db "github.com/baker-yuan/go-mall/common/db"
	"github.com/baker-yuan/go-mall/common/entity"
	"github.com/baker-yuan/go-mall/common/util"
	"gorm.io/gorm"
)

// CartItemRepo 购物车表
type CartItemRepo struct {
	*db.GenericDao[entity.CartItem, uint64]
}

// NewCartItemRepo 创建
func NewCartItemRepo(conn *gorm.DB) *CartItemRepo {
	return &CartItemRepo{
		GenericDao: &db.GenericDao[entity.CartItem, uint64]{
			DB: conn,
		},
	}
}

func init() {
	registerInitField(initCartItemField)
}

var (
	// 全字段修改CartItem那些字段不修改
	notUpdateCartItemField = []string{
		"created_at",
	}
	updateCartItemField []string
)

// InitCartItemField 全字段修改
func initCartItemField(db *gorm.DB) error {
	columnTypes, err := db.Migrator().ColumnTypes(&entity.CartItem{})
	if err != nil {
		return err
	}
	columns := make([]string, 0)
	for _, v := range columnTypes {
		columns = append(columns, v.Name())
	}
	updateCartItemField = util.SliceRemove[string](columns, notUpdateCartItemField...)
	return nil
}

// Create 创建购物车表
func (r CartItemRepo) Create(ctx context.Context, cartItem *entity.CartItem) error {
	if cartItem.ID > 0 {
		return errors.New("illegal argument cartItem id exist")
	}
	return r.GenericDao.Create(ctx, cartItem)
}

// DeleteByID 根据主键ID删除购物车表
func (r CartItemRepo) DeleteByID(ctx context.Context, id uint64) error {
	return r.GenericDao.DeleteByID(ctx, id)
}

// Update 修改购物车表
func (r CartItemRepo) Update(ctx context.Context, cartItem *entity.CartItem) error {
	if cartItem.ID == 0 {
		return errors.New("illegal argument cartItem exist")
	}
	return r.GenericDao.DB.WithContext(ctx).Select(updateCartItemField).Updates(cartItem).Error
}

// GetByID 根据主键ID查询购物车表
func (r CartItemRepo) GetByID(ctx context.Context, id uint64) (*entity.CartItem, error) {
	return r.GenericDao.GetByID(ctx, id)
}

// GetByDBOption 根据动态条件查询购物车表
func (r CartItemRepo) GetByDBOption(ctx context.Context, pageNum uint32, pageSize uint32, opts ...db.DBOption) ([]*entity.CartItem, uint32, error) {
	var (
		res       = make([]*entity.CartItem, 0)
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

// GetEffectCartItemByMemberID 根据会员id查询购物车
func (r CartItemRepo) GetEffectCartItemByMemberID(ctx context.Context, memberID uint64) ([]*entity.CartItem, error) {
	res := make([]*entity.CartItem, 0)
	if err := r.GenericDao.DB.WithContext(ctx).
		Where("delete_status = 0").
		Where("member_id = ?", memberID).
		Find(&res).Error; err != nil {
		return nil, err
	}
	return res, nil
}
