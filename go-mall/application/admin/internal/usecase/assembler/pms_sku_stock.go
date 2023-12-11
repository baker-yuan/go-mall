package assembler

import (
	"github.com/baker-yuan/go-mall/common/entity"
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
		Id:        skuStockEntity.ID,
		SkuCode:   skuStockEntity.SkuCode,
		Pic:       skuStockEntity.Pic,
		Sale:      skuStockEntity.Sale,
		SpData:    skuStockEntity.SpData,
		ProductId: skuStockEntity.ProductID,
		// 价格
		Price:          skuStockEntity.Price,
		PromotionPrice: skuStockEntity.PromotionPrice,
		// 库存
		Stock:     skuStockEntity.Stock,
		LowStock:  skuStockEntity.LowStock,
		LockStock: skuStockEntity.LockStock,
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
