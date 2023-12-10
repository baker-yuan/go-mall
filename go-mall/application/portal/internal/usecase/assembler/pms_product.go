package assembler

import (
	"github.com/baker-yuan/go-mall/common/entity"
	"github.com/baker-yuan/go-mall/common/util"
	pb "github.com/baker-yuan/go-mall/proto/mall"
)

func ProductEntityToModel(product *entity.Product) *pb.ProductItem {
	return &pb.ProductItem{
		// 基本信息
		Id:                product.ID,
		ProductCategoryId: product.ProductCategoryID,
		Name:              product.Name,
		SubTitle:          product.SubTitle,
		Price:             product.Price,
		// 促销信息

		// 属性信息
		Pic: util.GetFullUrl(product.Pic),
		// 状态

		// 其他
		Sale: product.Sale,

		// 冗余字段

		// 设置
	}
}
