package assembler

import (
	"github.com/baker-yuan/go-mall/common/entity"
	"github.com/baker-yuan/go-mall/common/util"
	pb "github.com/baker-yuan/go-mall/proto/mall"
)

func BrandEntityToDetail(brand *entity.Brand) *pb.ProductDetailRsp_Brand {
	return &pb.ProductDetailRsp_Brand{
		Name:        brand.Name,
		FirstLetter: brand.FirstLetter,
		Logo:        util.GetFullUrl(brand.Logo),
	}
}
