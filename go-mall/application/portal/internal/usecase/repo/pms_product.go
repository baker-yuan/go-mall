package repo

import (
	"context"

	db2 "github.com/baker-yuan/go-mall/common/db"
	"github.com/baker-yuan/go-mall/common/entity"
	pb "github.com/baker-yuan/go-mall/proto/mall"
	"gorm.io/gorm"
)

// ProductRepo 商品信息表
type ProductRepo struct {
	*db2.GenericDao[entity.Product, uint64]
}

// NewProductRepo 创建
func NewProductRepo(conn *gorm.DB) *ProductRepo {
	return &ProductRepo{
		GenericDao: &db2.GenericDao[entity.Product, uint64]{
			DB: conn,
		},
	}
}

// SearchProduct 综合搜索商品
func (r ProductRepo) SearchProduct(ctx context.Context, req *pb.SearchProductReq) (entity.Products, error) {
	var (
		res = make([]*entity.Product, 0)
	)
	tx := r.GenericDao.DB.WithContext(ctx)
	tx = tx.Where("delete_status = ?", 0)
	tx = tx.Where("publish_status = ?", 1)
	if len(req.GetKeyword()) != 0 {
		tx = tx.Where("name like ?", "%"+req.GetKeyword()+"%")
	}
	if req.GetBrandId() != nil {
		tx = tx.Where("brand_id = ?", req.GetBrandId().GetValue())
	}
	if req.GetProductCategoryId() != nil {
		tx = tx.Where("product_category_id = ?", req.GetProductCategoryId().GetValue())
	}

	tx = tx.Offset(int((req.GetPageNum() - 1) * req.GetPageSize())).Limit(int(req.GetPageSize()))

	switch req.GetSort() {
	case 1:
		tx = tx.Order("id desc")
	case 2:
		tx = tx.Order("sale desc")
	case 3:
		tx = tx.Order("price asc")
	case 4:
		tx = tx.Order("price desc")
	}

	if err := tx.Find(&res).Error; err != nil {
		return nil, err
	}
	return res, nil
}
