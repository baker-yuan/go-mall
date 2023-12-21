package assembler

import (
	"github.com/baker-yuan/go-mall/common/entity"
	pb "github.com/baker-yuan/go-mall/proto/mall"
)

// ProductLadderEntityToModel entity转pb
func ProductLadderEntityToModel(productLadder *entity.ProductLadder) *pb.ProductLadder {
	return &pb.ProductLadder{}
}
