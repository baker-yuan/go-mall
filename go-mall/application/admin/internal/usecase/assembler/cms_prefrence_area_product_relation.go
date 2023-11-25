package assembler

import (
	"github.com/baker-yuan/go-mall/application/admin/internal/entity"
	pb "github.com/baker-yuan/go-mall/proto/mall"
)

// PrefrenceAreaProductRelationEntityToModel entity转pb
func PrefrenceAreaProductRelationEntityToModel(prefrenceAreaProductRelation *entity.PrefrenceAreaProductRelation) *pb.PrefrenceAreaProductRelation {
	return &pb.PrefrenceAreaProductRelation{}
}

// PrefrenceAreaProductRelationsToEntity pb转entity
func PrefrenceAreaProductRelationsToEntity(prefrenceAreaProductRelationPbs []*pb.PrefrenceAreaProductRelation) []*entity.PrefrenceAreaProductRelation {
	res := make([]*entity.PrefrenceAreaProductRelation, 0)
	for _, prefrenceAreaProductRelationPb := range prefrenceAreaProductRelationPbs {
		res = append(res, PrefrenceAreaProductRelationToEntity(prefrenceAreaProductRelationPb))
	}
	return res
}

// PrefrenceAreaProductRelationToEntity pb转entity
func PrefrenceAreaProductRelationToEntity(prefrenceAreaProductRelationPb *pb.PrefrenceAreaProductRelation) *entity.PrefrenceAreaProductRelation {
	return &entity.PrefrenceAreaProductRelation{
		ID:              prefrenceAreaProductRelationPb.Id,
		PrefrenceAreaID: prefrenceAreaProductRelationPb.PrefrenceAreaId,
		ProductID:       prefrenceAreaProductRelationPb.ProductId,
	}
}
