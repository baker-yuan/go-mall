package assembler

import (
	"github.com/baker-yuan/go-mall/common/entity"
	"github.com/baker-yuan/go-mall/common/util"
	pb "github.com/baker-yuan/go-mall/proto/mall"
)

func ProductEntityToProductListItem(product *entity.Product) *pb.SearchProductRsp_Product {
	return &pb.SearchProductRsp_Product{
		Id:       product.ID,
		Pic:      util.ImgUtils.GetFullUrl(product.Pic),
		Name:     product.Name,
		SubTitle: product.SubTitle,
		Price:    product.Price,
		Sale:     product.Sale,
	}
}

func ProductEntityToDetail(product *entity.Product) *pb.ProductDetailRsp_Product {
	return &pb.ProductDetailRsp_Product{
		Id:            product.ID,
		AlbumPics:     util.ImgUtils.GetFullUrls(product.AlbumPics),
		Pic:           util.ImgUtils.GetFullUrl(product.Pic),
		Name:          product.Name,
		SubTitle:      product.SubTitle,
		Price:         product.Price,
		OriginalPrice: product.OriginalPrice,
		Sale:          product.Sale,
		Stock:         product.Stock,
		ServiceIds:    product.ServiceIDs,
	}
}
