package assembler

import (
	"github.com/baker-yuan/go-mall/common/entity"
	"github.com/baker-yuan/go-mall/common/util"
	pb "github.com/baker-yuan/go-mall/proto/mall"
)

// OrderReturnApplyEntityToModel entity转pb
func OrderReturnApplyEntityToModel(orderReturnApply *entity.OrderReturnApply) *pb.OrderReturnApply {
	return &pb.OrderReturnApply{
		Id:      orderReturnApply.ID,
		OrderId: orderReturnApply.OrderID,
		// 商品信息
		ProductPic:       util.GetFullUrl(orderReturnApply.ProductPic),
		ProductName:      orderReturnApply.ProductName,
		ProductBrand:     orderReturnApply.ProductBrand,
		ProductId:        orderReturnApply.ProductID,
		ProductRealPrice: orderReturnApply.ProductRealPrice,
		ProductAttr:      orderReturnApply.ProductAttr,
		ProductCount:     orderReturnApply.ProductCount,
		ProductPrice:     orderReturnApply.ProductPrice,
		//
		Status:         uint32(orderReturnApply.Status),
		OrderSn:        orderReturnApply.OrderSN,
		CreateTime:     orderReturnApply.CreateTime,
		MemberUsername: orderReturnApply.MemberUsername,
		ReturnName:     orderReturnApply.ReturnName,
		ReturnPhone:    orderReturnApply.ReturnPhone,
		Reason:         orderReturnApply.Reason,
		Description:    orderReturnApply.Description,
		ProofPics:      util.GetFullUrls(orderReturnApply.ProofPics),
		//
		ReturnAmount:     orderReturnApply.ReturnAmount,
		CompanyAddressId: orderReturnApply.CompanyAddressID,
		// 商家-处理人
		HandleMan:  orderReturnApply.HandleMan,
		HandleTime: orderReturnApply.HandleTime,
		HandleNote: orderReturnApply.HandleNote,
		// 商家-收货人
		ReceiveMan:  orderReturnApply.ReceiveMan,
		ReceiveTime: orderReturnApply.ReceiveTime,
		ReceiveNote: orderReturnApply.ReceiveNote,
	}
}
