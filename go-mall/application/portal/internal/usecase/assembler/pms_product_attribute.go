package assembler

import (
	"github.com/baker-yuan/go-mall/common/entity"
	pb "github.com/baker-yuan/go-mall/proto/mall"
)

func ProductAttributesToDetail(productAttributes entity.ProductAttributes) []*pb.ProductAggregation_ProductAttribute {
	res := make([]*pb.ProductAggregation_ProductAttribute, 0)
	for _, productAttribute := range productAttributes {
		res = append(res, &pb.ProductAggregation_ProductAttribute{
			Id:                         productAttribute.ID,
			Type:                       productAttribute.Type,
			InputType:                  productAttribute.InputType,
			InputList:                  productAttribute.InputList,
			ProductAttributeCategoryId: productAttribute.ProductAttributeCategoryID,
			Name:                       productAttribute.Name,
			HandAddStatus:              productAttribute.HandAddStatus,
		})
	}
	return res
}
