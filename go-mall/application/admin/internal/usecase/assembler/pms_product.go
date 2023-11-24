package assembler

import (
	"github.com/baker-yuan/go-mall/application/admin/internal/entity"
	pb "github.com/baker-yuan/go-mall/proto/mall"
)

// ProductEntityToModel entity转pb
func ProductEntityToModel(product *entity.Product, categoryNames map[uint64]string, brandNames map[uint64]string) *pb.Product {
	return &pb.Product{}
}

// AddOrUpdateProductParamToEntity pb转entity
func AddOrUpdateProductParamToEntity(param *pb.AddOrUpdateProductParam) *entity.Product {
	return &entity.Product{}
}
