package assembler

import (
	"github.com/baker-yuan/go-mall/application/admin/internal/entity"
	pb "github.com/baker-yuan/go-mall/proto/mall"
)

// ProductAttributeEntityToModel entity转pb
func ProductAttributeEntityToModel(productAttribute *entity.ProductAttribute) *pb.ProductAttribute {
	return &pb.ProductAttribute{
		Id:                         productAttribute.ID,
		Type:                       uint32(productAttribute.Type),
		ProductAttributeCategoryId: productAttribute.ProductAttributeCategoryID,
		Name:                       productAttribute.Name,
		SelectType:                 uint32(productAttribute.SelectType),
		InputType:                  uint32(productAttribute.InputType),
		InputList:                  productAttribute.InputList,
		Sort:                       uint32(productAttribute.Sort),
		FilterType:                 uint32(productAttribute.FilterType),
		SearchType:                 uint32(productAttribute.SearchType),
		RelatedStatus:              uint32(productAttribute.RelatedStatus),
		HandAddStatus:              uint32(productAttribute.HandAddStatus),
	}

}

// AddOrUpdateProductAttributeParamToEntity pb转entity
func AddOrUpdateProductAttributeParamToEntity(param *pb.AddOrUpdateProductAttributeParam) *entity.ProductAttribute {
	return &entity.ProductAttribute{
		Type:                       uint8(param.GetType()),
		ProductAttributeCategoryID: param.GetProductAttributeCategoryId(),
		Name:                       param.GetName(),
		SelectType:                 uint8(param.GetSelectType()),
		InputType:                  uint8(param.GetInputType()),
		InputList:                  param.GetInputList(),
		Sort:                       int(param.GetSort()),
		FilterType:                 uint8(param.GetFilterType()),
		SearchType:                 uint8(param.GetSearchType()),
		RelatedStatus:              uint8(param.GetRelatedStatus()),
		HandAddStatus:              uint8(param.GetHandAddStatus()),
	}
}
