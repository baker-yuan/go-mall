package assembler

import (
	"github.com/baker-yuan/go-mall/common/entity"
	pb "github.com/baker-yuan/go-mall/proto/mall"
)

// ProductFullReductionsToModel entity转pb
func ProductFullReductionsToModel(productFullReductions []*entity.ProductFullReduction) []*pb.ProductFullReduction {
	res := make([]*pb.ProductFullReduction, 0)
	for _, productFullReduction := range productFullReductions {
		res = append(res, ProductFullReductionToModel(productFullReduction))
	}
	return res
}

// ProductFullReductionToModel entity转pb
func ProductFullReductionToModel(productFullReduction *entity.ProductFullReduction) *pb.ProductFullReduction {
	return &pb.ProductFullReduction{
		Id:          productFullReduction.ID,
		ProductId:   productFullReduction.ProductID,
		FullPrice:   productFullReduction.FullPrice,
		ReducePrice: productFullReduction.ReducePrice,
	}
}

// ProductFullReductionsToEntity pb转entity
func ProductFullReductionsToEntity(productFullReductionPbs []*pb.ProductFullReduction) []*entity.ProductFullReduction {
	res := make([]*entity.ProductFullReduction, 0)
	for _, productFullReductionPb := range productFullReductionPbs {
		res = append(res, ProductFullReductionToEntity(productFullReductionPb))
	}
	return res
}

// ProductFullReductionToEntity pb转entity
func ProductFullReductionToEntity(productFullReductionPb *pb.ProductFullReduction) *entity.ProductFullReduction {
	return &entity.ProductFullReduction{
		ID:          productFullReductionPb.Id,
		ProductID:   productFullReductionPb.ProductId,
		FullPrice:   productFullReductionPb.FullPrice,
		ReducePrice: productFullReductionPb.ReducePrice,
	}
}
