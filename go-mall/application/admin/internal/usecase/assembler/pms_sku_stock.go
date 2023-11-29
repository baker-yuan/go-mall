package assembler

import (
	"github.com/baker-yuan/go-mall/application/admin/internal/entity"
	pb "github.com/baker-yuan/go-mall/proto/mall"
)

// SkuStocksToModel entity转pb
func SkuStocksToModel(skuStocks []*entity.SkuStock) []*pb.SkuStock {
	res := make([]*pb.SkuStock, 0)
	for _, skuStock := range skuStocks {
		res = append(res, SkuStockToModel(skuStock))
	}
	return res
}

// SkuStockToModel entity转pb
func SkuStockToModel(skuStockEntity *entity.SkuStock) *pb.SkuStock {
	return &pb.SkuStock{
		Id:             skuStockEntity.ID,
		ProductId:      skuStockEntity.ProductID,
		SkuCode:        skuStockEntity.SkuCode,
		Price:          skuStockEntity.Price,
		Stock:          skuStockEntity.Stock,
		LowStock:       skuStockEntity.LowStock,
		Pic:            skuStockEntity.Pic,
		Sale:           skuStockEntity.Sale,
		PromotionPrice: skuStockEntity.PromotionPrice,
		LockStock:      skuStockEntity.LockStock,
		SpData:         skuStockEntity.SpData,
	}
}

// SkuStocksToEntity pb转entity
func SkuStocksToEntity(skuStockPbs []*pb.SkuStock) []*entity.SkuStock {
	res := make([]*entity.SkuStock, 0)
	for _, skuStockPb := range skuStockPbs {
		res = append(res, SkuStockToEntity(skuStockPb))
	}
	return res
}

// SkuStockToEntity pb转entity
func SkuStockToEntity(skuStockPb *pb.SkuStock) *entity.SkuStock {
	return &entity.SkuStock{
		ID:             skuStockPb.Id,
		ProductID:      skuStockPb.ProductId,
		SkuCode:        skuStockPb.SkuCode,
		Price:          skuStockPb.Price,
		Stock:          skuStockPb.Stock,
		LowStock:       skuStockPb.LowStock,
		Pic:            skuStockPb.Pic,
		Sale:           skuStockPb.Sale,
		PromotionPrice: skuStockPb.PromotionPrice,
		LockStock:      skuStockPb.LockStock,
		SpData:         skuStockPb.SpData,
	}
}
