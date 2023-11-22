package assembler

import (
	"github.com/baker-yuan/go-mall/application/admin/internal/entity"
	pb "github.com/baker-yuan/go-mall/proto/mall"
)

// ProductAttributeCategoryEntityToModel entity转pb
func ProductAttributeCategoryEntityToModel(productAttributeCategory *entity.ProductAttributeCategory) *pb.ProductAttributeCategory {
	return &pb.ProductAttributeCategory{
		Id:             productAttributeCategory.ID,
		Name:           productAttributeCategory.Name,
		AttributeCount: uint32(productAttributeCategory.AttributeCount),
		ParamCount:     uint32(productAttributeCategory.ParamCount),
	}
}

// AddOrUpdateProductAttributeCategoryParamToEntity pb转entity
func AddOrUpdateProductAttributeCategoryParamToEntity(param *pb.AddOrUpdateProductAttributeCategoryParam) *entity.ProductAttributeCategory {
	return &entity.ProductAttributeCategory{
		Name: param.GetName(),
	}
}
