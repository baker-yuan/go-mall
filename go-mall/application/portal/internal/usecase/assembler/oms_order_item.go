package assembler

import (
	"github.com/baker-yuan/go-mall/common/entity"
	pb "github.com/baker-yuan/go-mall/proto/mall"
)

// OrderItemEntityToModel entity转pb
func OrderItemEntityToModel(orderItem *entity.OrderItem) *pb.OrderItem {
	return &pb.OrderItem{}
}
