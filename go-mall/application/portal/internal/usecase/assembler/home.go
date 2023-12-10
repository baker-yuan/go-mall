package assembler

import (
	"github.com/baker-yuan/go-mall/common/entity"
	"github.com/baker-yuan/go-mall/common/util"
	pb "github.com/baker-yuan/go-mall/proto/mall"
)

func ProductCategoryEntityToModel(category *entity.ProductCategory) *pb.ProductCategoryItem {
	return &pb.ProductCategoryItem{
		Id:       category.ID,
		ParentId: category.ParentID,
		Name:     category.Name,
		Icon:     util.GetFullUrl(category.Icon),
	}
}
