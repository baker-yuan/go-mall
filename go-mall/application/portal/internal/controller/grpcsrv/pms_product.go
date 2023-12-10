package grpcsrv

import (
	"context"

	"github.com/baker-yuan/go-mall/common/retcode"
	pb "github.com/baker-yuan/go-mall/proto/mall"
)

// SearchProduct 综合搜索商品
func (s PortalApiImpl) SearchProduct(ctx context.Context, req *pb.SearchProductReq) (*pb.SearchProductRsp, error) {
	var (
		res = &pb.SearchProductRsp{}
	)
	products, err := s.product.SearchProduct(ctx, req)
	if err != nil {
		return nil, err
	}
	res.Data = &pb.ProductRspItem{
		Data: products,
	}

	res.Code, res.Message = retcode.GetRetCodeMsg(retcode.RetSuccess)
	return res, nil
}
