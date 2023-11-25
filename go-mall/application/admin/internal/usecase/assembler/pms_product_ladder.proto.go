package assembler

import (
	"github.com/baker-yuan/go-mall/application/admin/internal/entity"
	pb "github.com/baker-yuan/go-mall/proto/mall"
)

// ProductLadderEntityToModel entity转pb
func ProductLadderEntityToModel(productLadder *entity.ProductLadder) *pb.ProductLadder {
	return &pb.ProductLadder{}
}

// ProductLaddersToEntity pb转entity
func ProductLaddersToEntity(productLadderPbs []*pb.ProductLadder) []*entity.ProductLadder {
	res := make([]*entity.ProductLadder, 0)
	for _, productLadderPb := range productLadderPbs {
		res = append(res, ProductLadderToEntity(productLadderPb))
	}
	return res
}

// ProductLadderToEntity pb转entity
func ProductLadderToEntity(productLadderPb *pb.ProductLadder) *entity.ProductLadder {
	return &entity.ProductLadder{
		ID:        productLadderPb.Id,
		ProductID: productLadderPb.ProductId,
		Count:     productLadderPb.Count,
		Discount:  productLadderPb.Discount,
		Price:     productLadderPb.Price,
	}
}
