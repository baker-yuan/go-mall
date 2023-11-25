package assembler

import (
	"github.com/baker-yuan/go-mall/application/admin/internal/entity"
	pb "github.com/baker-yuan/go-mall/proto/mall"
)

// ProductFullReductionEntityToModel entity转pb
func ProductFullReductionEntityToModel(productFullReduction *entity.ProductFullReduction) *pb.ProductFullReduction {
	return &pb.ProductFullReduction{}
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
