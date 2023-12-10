package grpcsrv

import (
	"context"

	"github.com/baker-yuan/go-mall/application/portal/internal/usecase"
	"github.com/baker-yuan/go-mall/common/retcode"
	pb "github.com/baker-yuan/go-mall/proto/mall"
)

var _ pb.PortalApiServer = &PortalApiImpl{}

type PortalApi interface {
	pb.PortalApiServer
}

type PortalApiImpl struct {
	pb.UnimplementedPortalApiServer

	home usecase.IHomeUseCase
}

func New(home usecase.IHomeUseCase) PortalApi {
	return &PortalApiImpl{
		home: home,
	}
}

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
