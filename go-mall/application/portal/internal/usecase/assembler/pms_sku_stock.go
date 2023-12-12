package assembler

import (
	"github.com/baker-yuan/go-mall/common/entity"
	pb "github.com/baker-yuan/go-mall/proto/mall"
)

func SkuStocksToDetail(skuStocks entity.SkuStocks) []*pb.ProductDetailRsp_SkuStock {
	res := make([]*pb.ProductDetailRsp_SkuStock, 0)
	for _, skuStock := range skuStocks {
		res = append(res, &pb.ProductDetailRsp_SkuStock{
			Id:        skuStock.ID,
			SkuCode:   skuStock.SkuCode,
			Pic:       skuStock.Pic,
			Sale:      skuStock.Sale,
			SpData:    skuStock.SpData,
			ProductId: skuStock.ProductID,
			// 价格
			Price:          skuStock.Price,
			PromotionPrice: skuStock.PromotionPrice,
			// 库存
			Stock:     skuStock.Stock,
			LowStock:  skuStock.LowStock,
			LockStock: skuStock.LockStock,
		})
	}
	return res
}
