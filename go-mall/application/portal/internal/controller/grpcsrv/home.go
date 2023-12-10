package grpcsrv

import (
	"context"

	"github.com/baker-yuan/go-mall/common/retcode"
	pb "github.com/baker-yuan/go-mall/proto/mall"
)

// ProductCategoryList 获取首页商品分类
func (s PortalApiImpl) ProductCategoryList(ctx context.Context, req *pb.ProductCategoryListReq) (*pb.ProductCategoryListRsp, error) {
	var (
		res = &pb.ProductCategoryListRsp{}
	)
	categories, err := s.home.ProductCategoryList(ctx, req)
	if err != nil {
		return nil, err
	}
	res.Data = categories

	res.Code, res.Message = retcode.GetRetCodeMsg(retcode.RetSuccess)
	return res, nil
}
