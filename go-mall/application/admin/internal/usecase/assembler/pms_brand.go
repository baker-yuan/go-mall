package assembler

import (
	"github.com/baker-yuan/go-mall/common/entity"
	"github.com/baker-yuan/go-mall/common/util"
	pb "github.com/baker-yuan/go-mall/proto/mall"
)

// BrandEntityToModel entity转pb
func BrandEntityToModel(brand *entity.Brand) *pb.Brand {
	return &pb.Brand{
		Id:          brand.ID,
		Name:        brand.Name,
		FirstLetter: brand.FirstLetter,
		Logo:        util.GetFullUrl(brand.Logo),
		BigPic:      util.GetFullUrl(brand.BigPic),
		BrandStory:  brand.BrandStory,
		Sort:        brand.Sort,
		// status
		FactoryStatus: uint32(brand.FactoryStatus),
		ShowStatus:    uint32(brand.ShowStatus),
		// 冗余字段
		ProductCount:        brand.ProductCount,
		ProductCommentCount: brand.ProductCommentCount,
	}
}

// AddOrUpdateBrandParamToEntity pb转entity
func AddOrUpdateBrandParamToEntity(param *pb.AddOrUpdateBrandParam) *entity.Brand {
	return &entity.Brand{
		Name:        param.Name,
		FirstLetter: param.FirstLetter,
		Logo:        util.GetRelativeUrl(param.Logo),
		BigPic:      util.GetRelativeUrl(param.BigPic),
		BrandStory:  param.BrandStory,
		Sort:        param.Sort,
		// status
		FactoryStatus: uint8(param.FactoryStatus),
		ShowStatus:    uint8(param.ShowStatus),
	}
}
