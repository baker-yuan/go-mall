package assembler

import (
	"github.com/baker-yuan/go-mall/application/admin/internal/entity"
	"github.com/baker-yuan/go-mall/application/admin/pkg/util"
	pb "github.com/baker-yuan/go-mall/proto/mall"
)

func ProductCategoryEntityToModel(category *entity.ProductCategory) *pb.ProductCategory {
	return &pb.ProductCategory{
		Id:           category.ID,
		ParentId:     category.ParentID,
		Name:         category.Name,
		Level:        uint32(category.Level),
		ProductCount: category.ProductCount,
		ProductUnit:  category.ProductUnit,
		NavStatus:    uint32(category.NavStatus),
		ShowStatus:   uint32(category.ShowStatus),
		Sort:         category.Sort,
		Icon:         util.GetFullUrl(category.Icon),
		Keywords:     category.Keywords,
		Description:  category.Description,
	}
}

func AddOrUpdateProductCategoryParamToEntity(param *pb.AddOrUpdateProductCategoryParam) *entity.ProductCategory {
	return &entity.ProductCategory{
		ParentID:    param.ParentId,
		Name:        param.Name,
		ProductUnit: param.ProductUnit,
		NavStatus:   uint8(param.NavStatus),
		ShowStatus:  uint8(param.ShowStatus),
		Sort:        param.Sort,
		Icon:        util.GetRelativeUrl(param.Icon),
		Keywords:    param.Keywords,
		Description: param.Description,
	}
}
