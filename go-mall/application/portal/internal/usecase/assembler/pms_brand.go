package assembler

import (
	"github.com/baker-yuan/go-mall/common/entity"
	"github.com/baker-yuan/go-mall/common/util"
	pb "github.com/baker-yuan/go-mall/proto/mall"
)

func BrandEntityToDetail(brand *entity.Brand) *pb.ProductAggregation_Brand {
	return &pb.ProductAggregation_Brand{
		Name:        brand.Name,
		FirstLetter: brand.FirstLetter,
		Logo:        util.GetFullUrl(brand.Logo),
	}
}
