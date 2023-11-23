package grpcsrv

import (
	"context"

	"github.com/baker-yuan/go-mall/application/admin/pkg/retcode"
	pb "github.com/baker-yuan/go-mall/proto/mall"
)

// CreateProductCategory 添加商品分类
func (s *AdminApiImpl) CreateProductCategory(ctx context.Context, param *pb.AddOrUpdateProductCategoryParam) (*pb.CommonRsp, error) {
	if err := s.category.CreateProductCategory(ctx, param); err != nil {
		return nil, err
	}

	res := &pb.CommonRsp{}
	res.Code, res.Message = retcode.GetRetCodeMsg(retcode.RetSuccess)
	return res, nil
}

// UpdateProductCategory 修改商品分类
func (s *AdminApiImpl) UpdateProductCategory(ctx context.Context, param *pb.AddOrUpdateProductCategoryParam) (*pb.CommonRsp, error) {
	if err := s.category.UpdateProductCategory(ctx, param); err != nil {
		return nil, err
	}

	res := &pb.CommonRsp{}
	res.Code, res.Message = retcode.GetRetCodeMsg(retcode.RetSuccess)
	return res, nil
}

// GetProductCategories 分页查询商品分类
func (s *AdminApiImpl) GetProductCategories(ctx context.Context, param *pb.GetProductCategoriesParam) (*pb.GetProductCategoriesRsp, error) {
	categories, pageTotal, err := s.category.GetProductCategories(ctx, param)
	if err != nil {
		return nil, err
	}
	res := &pb.GetProductCategoriesRsp{
		Data: &pb.ProductCategoriesData{
			Data:      categories,
			PageTotal: pageTotal,
			PageNum:   param.PageNum,
			PageSize:  param.PageSize,
		},
	}

	res.Code, res.Message = retcode.GetRetCodeMsg(retcode.RetSuccess)
	return res, nil
}

// GetProductCategory 根据id获取商品分类
func (s *AdminApiImpl) GetProductCategory(ctx context.Context, param *pb.GetProductCategoryReq) (*pb.GetProductCategoryRsp, error) {
	productCategory, err := s.category.GetProductCategory(ctx, param.GetId())
	if err != nil {
		return nil, err
	}

	res := &pb.GetProductCategoryRsp{
		Data: productCategory,
	}
	res.Code, res.Message = retcode.GetRetCodeMsg(retcode.RetSuccess)
	return res, nil
}

// DeleteProductCategory 删除商品分类
func (s *AdminApiImpl) DeleteProductCategory(ctx context.Context, param *pb.DeleteProductCategoryReq) (*pb.CommonRsp, error) {
	if err := s.category.DeleteProductCategory(ctx, param.GetId()); err != nil {
		return nil, err
	}

	res := &pb.CommonRsp{}
	res.Code, res.Message = retcode.GetRetCodeMsg(retcode.RetSuccess)
	return res, nil
}

// GetProductCategoriesWithChildren 查询所有一级分类及子分类
func (s *AdminApiImpl) GetProductCategoriesWithChildren(ctx context.Context, param *pb.GetProductCategoriesWithChildrenReq) (*pb.GetProductCategoriesWithChildrenRsp, error) {
	items, err := s.category.GetProductCategoriesWithChildren(ctx)
	if err != nil {
		return nil, err
	}

	res := &pb.GetProductCategoriesWithChildrenRsp{
		Data: items,
	}
	res.Code, res.Message = retcode.GetRetCodeMsg(retcode.RetSuccess)
	return res, nil
}
