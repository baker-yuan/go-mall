package assembler

import (
	"github.com/baker-yuan/go-mall/common/entity"
	pb "github.com/baker-yuan/go-mall/proto/mall"
)

func ProductAttributeValuesToDetail(attributeValues entity.ProductAttributeValues) []*pb.ProductAggregation_ProductAttributeValue {
	res := make([]*pb.ProductAggregation_ProductAttributeValue, 0)
	for _, attributeValue := range attributeValues {
		res = append(res, &pb.ProductAggregation_ProductAttributeValue{
			Id:                 attributeValue.ID,
			ProductId:          attributeValue.ProductID,
			ProductAttributeId: attributeValue.ProductAttributeID,
			Value:              attributeValue.Value,
		})
	}
	return res

}
