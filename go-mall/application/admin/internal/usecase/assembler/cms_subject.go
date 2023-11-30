package assembler

import (
	"github.com/baker-yuan/go-mall/application/admin/internal/entity"
	pb "github.com/baker-yuan/go-mall/proto/mall"
)

// SubjectEntityToModel entity转pb
func SubjectEntityToModel(subject *entity.Subject) *pb.Subject {
	return &pb.Subject{
		Id:              subject.ID,
		CategoryId:      subject.CategoryID,
		Title:           subject.Title,
		Pic:             subject.Pic,
		ProductCount:    subject.ProductCount,
		RecommendStatus: uint32(subject.RecommendStatus),
		CreateTime:      subject.CreateTime,
		CollectCount:    subject.CollectCount,
		ReadCount:       subject.ReadCount,
		CommentCount:    subject.CommentCount,
		AlbumPics:       subject.AlbumPics,
		Description:     subject.Description,
		ShowStatus:      uint32(subject.ShowStatus),
		Content:         subject.Content,
		ForwardCount:    subject.ForwardCount,
		CategoryName:    subject.CategoryName,
	}
}

// AddOrUpdateSubjectParamToEntity pb转entity
func AddOrUpdateSubjectParamToEntity(param *pb.AddOrUpdateSubjectParam) *entity.Subject {
	return &entity.Subject{}
}
