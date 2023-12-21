package assembler

import (
	"github.com/baker-yuan/go-mall/common/entity"
	pb "github.com/baker-yuan/go-mall/proto/mall"
)

// ProductFullReductionEntityToModel entity转pb
func ProductFullReductionEntityToModel(productFullReduction *entity.ProductFullReduction) *pb.ProductFullReduction {
	return &pb.ProductFullReduction{}
}
