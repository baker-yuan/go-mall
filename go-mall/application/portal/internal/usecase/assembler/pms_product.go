package assembler

import (
	"github.com/baker-yuan/go-mall/common/entity"
	pb "github.com/baker-yuan/go-mall/proto/mall"
)

func ProductEntityToModel(product *entity.Product) *pb.ProductItem {
	return &pb.ProductItem{
		Id:                product.ID,
		ProductCategoryId: product.ProductCategoryID,
		Name:              product.Name,
	}
}
