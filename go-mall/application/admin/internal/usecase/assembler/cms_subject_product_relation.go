package assembler

import (
	"github.com/baker-yuan/go-mall/application/admin/internal/entity"
	pb "github.com/baker-yuan/go-mall/proto/mall"
)

// SubjectProductRelationEntityToModel entity转pb
func SubjectProductRelationEntityToModel(subjectProductRelation *entity.SubjectProductRelation) *pb.SubjectProductRelation {
	return &pb.SubjectProductRelation{}
}

// SubjectProductRelationsToEntity pb转entity
func SubjectProductRelationsToEntity(subjectProductRelationPbs []*pb.SubjectProductRelation) []*entity.SubjectProductRelation {
	res := make([]*entity.SubjectProductRelation, 0)
	for _, subjectProductRelationPb := range subjectProductRelationPbs {
		res = append(res, SubjectProductRelationToEntity(subjectProductRelationPb))
	}
	return res
}

// SubjectProductRelationToEntity pb转entity
func SubjectProductRelationToEntity(subjectProductRelationPb *pb.SubjectProductRelation) *entity.SubjectProductRelation {
	return &entity.SubjectProductRelation{
		ID:        subjectProductRelationPb.Id,
		SubjectID: subjectProductRelationPb.SubjectId,
		ProductID: subjectProductRelationPb.ProductId,
	}
}
