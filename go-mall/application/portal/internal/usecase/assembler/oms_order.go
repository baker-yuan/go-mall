package assembler

import (
	"github.com/baker-yuan/go-mall/common/entity"
	pb "github.com/baker-yuan/go-mall/proto/mall"
)

// OrderEntityToModel entity转pb
func OrderEntityToModel(order *entity.Order) *pb.Order {
	return &pb.Order{}
}
