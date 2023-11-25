package assembler

import (
	"github.com/baker-yuan/go-mall/application/admin/internal/entity"
	pb "github.com/baker-yuan/go-mall/proto/mall"
)

// ProductAttributeValueEntityToModel entity转pb
func ProductAttributeValueEntityToModel(productAttributeValue *entity.ProductAttributeValue) *pb.ProductAttributeValue {
	return &pb.ProductAttributeValue{}
}

// ProductAttributeValuesToEntity pb转entity
func ProductAttributeValuesToEntity(productAttributeValuePbs []*pb.ProductAttributeValue) []*entity.ProductAttributeValue {
	res := make([]*entity.ProductAttributeValue, 0)
	for _, productAttributeValuePb := range productAttributeValuePbs {
		res = append(res, ProductAttributeValueToEntity(productAttributeValuePb))
	}
	return res
}

// ProductAttributeValueToEntity pb转entity
func ProductAttributeValueToEntity(productAttributeValuePb *pb.ProductAttributeValue) *entity.ProductAttributeValue {
	return &entity.ProductAttributeValue{
		ID:                 productAttributeValuePb.Id,
		ProductID:          productAttributeValuePb.ProductId,
		ProductAttributeID: productAttributeValuePb.ProductAttributeId,
		Value:              productAttributeValuePb.Value,
	}
}
