package assembler

import (
	"github.com/baker-yuan/go-mall/application/admin/internal/entity"
	"github.com/baker-yuan/go-mall/application/admin/pkg/util"
	pb "github.com/baker-yuan/go-mall/proto/mall"
)

// BrandEntityToModel entity转pb
func BrandEntityToModel(brand *entity.Brand) *pb.Brand {
	return &pb.Brand{
		Id:                  brand.ID,
		Name:                brand.Name,
		FirstLetter:         brand.FirstLetter,
		Sort:                brand.Sort,
		FactoryStatus:       uint32(brand.FactoryStatus),
		ShowStatus:          uint32(brand.ShowStatus),
		ProductCount:        brand.ProductCount,
		ProductCommentCount: brand.ProductCommentCount,
		Logo:                util.GetFullUrl(brand.Logo),
		BigPic:              util.GetFullUrl(brand.BigPic),
		BrandStory:          brand.BrandStory,
	}
}

// AddOrUpdateBrandParamToEntity pb转entity
func AddOrUpdateBrandParamToEntity(param *pb.AddOrUpdateBrandParam) *entity.Brand {
	return &entity.Brand{
		Name:          param.Name,
		FirstLetter:   param.FirstLetter,
		Sort:          param.Sort,
		FactoryStatus: uint8(param.FactoryStatus),
		ShowStatus:    uint8(param.ShowStatus),
		Logo:          util.GetRelativeUrl(param.Logo),
		BigPic:        util.GetRelativeUrl(param.BigPic),
		BrandStory:    param.BrandStory,
	}
}
