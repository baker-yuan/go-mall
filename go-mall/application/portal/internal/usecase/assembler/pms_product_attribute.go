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
			Type:                       uint32(productAttribute.Type),
			ProductAttributeCategoryId: productAttribute.ProductAttributeCategoryID,
			Name:                       productAttribute.Name,
			Sort:                       productAttribute.Sort,
			//
			SelectType: uint32(productAttribute.SelectType),
			InputType:  uint32(productAttribute.InputType),
			InputList:  productAttribute.InputList,
			//
			FilterType: uint32(productAttribute.FilterType),
			SearchType: uint32(productAttribute.SearchType),
			//
			RelatedStatus: uint32(productAttribute.RelatedStatus),
			HandAddStatus: uint32(productAttribute.HandAddStatus),
		})
	}
	return res
}
