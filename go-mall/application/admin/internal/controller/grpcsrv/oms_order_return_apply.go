package grpcsrv

import (
	"context"

	"github.com/baker-yuan/go-mall/application/admin/pkg/retcode"
	pb "github.com/baker-yuan/go-mall/proto/mall"
)

// GetOrderReturnApplies 分页查询订单退货申请
func (s *AdminApiImpl) GetOrderReturnApplies(ctx context.Context, param *pb.GetOrderReturnAppliesParam) (*pb.GetOrderReturnAppliesRsp, error) {
	var (
		res = &pb.GetOrderReturnAppliesRsp{}
	)

	orderReturnApplies, pageTotal, err := s.orderReturnApply.GetOrderReturnApplies(ctx, param)
	if err != nil {
		return nil, err
	}

	res.Data = &pb.OrderReturnAppliesData{
		Data:      orderReturnApplies,
		PageTotal: pageTotal,
		PageSize:  param.GetPageSize(),
		PageNum:   param.GetPageNum(),
	}

	res.Code, res.Message = retcode.GetRetCodeMsg(retcode.RetSuccess)
	return res, nil
}

// GetOrderReturnApply 根据id获取订单退货申请
func (s *AdminApiImpl) GetOrderReturnApply(ctx context.Context, param *pb.GetOrderReturnApplyReq) (*pb.GetOrderReturnApplyRsp, error) {
	var (
		res = &pb.GetOrderReturnApplyRsp{}
	)

	orderReturnApply, err := s.orderReturnApply.GetOrderReturnApply(ctx, param.GetId())
	if err != nil {
		return nil, err
	}
	res.Data = orderReturnApply

	res.Code, res.Message = retcode.GetRetCodeMsg(retcode.RetSuccess)
	return res, nil
}
