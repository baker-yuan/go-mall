package assembler

import (
	"github.com/baker-yuan/go-mall/common/entity"
	"github.com/baker-yuan/go-mall/common/util"
	pb "github.com/baker-yuan/go-mall/proto/mall"
)

func ProductEntityToProductListItem(product *entity.Product) *pb.ProductList_Product {
	return &pb.ProductList_Product{
		Id:       product.ID,
		Pic:      util.GetFullUrl(product.Pic),
		Name:     product.Name,
		SubTitle: product.SubTitle,
		Price:    product.Price,
		Sale:     product.Sale,
	}
}

func ProductEntityToDetail(product *entity.Product) *pb.ProductAggregation_Product {
	return &pb.ProductAggregation_Product{
		AlbumPics: util.GetFullUrls(product.AlbumPics),
		Pic:       util.GetFullUrl(product.Pic),
	}
}
