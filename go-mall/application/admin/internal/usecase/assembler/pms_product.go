package assembler

import (
	"github.com/baker-yuan/go-mall/application/admin/internal/entity"
	"github.com/baker-yuan/go-mall/application/admin/pkg/util"
	pb "github.com/baker-yuan/go-mall/proto/mall"
)

// ProductEntityToModel entity转pb
func ProductEntityToModel(product *entity.Product, categoryNames map[uint64]string, brandNames map[uint64]string) *pb.Product {
	return &pb.Product{
		// 基本信息
		Id:                product.ID,
		ProductCategoryId: product.ProductCategoryID,
		Name:              product.Name,
		SubTitle:          product.SubTitle,
		BrandId:           product.BrandID,
		Description:       product.Description,
		ProductSn:         product.ProductSN,
		Price:             product.Price,
		OriginalPrice:     product.OriginalPrice,
		Stock:             product.Stock,
		Unit:              product.Unit,
		Weight:            product.Weight,
		Sort:              product.Sort,

		// 促销信息
		GiftPoint:          product.GiftPoint,
		GiftGrowth:         product.GiftGrowth,
		UsePointLimit:      product.UsePointLimit,
		PreviewStatus:      uint32(product.PreviewStatus),
		PublishStatus:      uint32(product.PublishStatus),
		NewStatus:          uint32(product.NewStatus),
		RecommandStatus:    uint32(product.RecommandStatus),
		ServiceIds:         product.ServiceIDs,
		DetailTitle:        product.DetailTitle,
		DetailDesc:         product.DetailDesc,
		Keywords:           product.Keywords,
		Note:               product.Note,
		PromotionType:      uint32(int32(product.PromotionType)),
		PromotionPrice:     product.PromotionPrice,
		PromotionStartTime: uint32(int32(product.PromotionStartTime)),
		PromotionEndTime:   uint32(int32(product.PromotionEndTime)),

		// 属性信息
		ProductAttributeCategoryId: product.ProductAttributeCategoryID,
		Pic:                        util.GetFullUrl(product.Pic),
		AlbumPics:                  product.AlbumPics,
		DetailHtml:                 product.DetailHTML,
		DetailMobileHtml:           product.DetailMobileHTML,

		// 状态
		VerifyStatus: int32(product.VerifyStatus),
		DeleteStatus: int32(product.DeleteStatus),

		// 其他
		FeightTemplateId:  product.FeightTemplateID,
		Sale:              product.Sale,
		LowStock:          product.LowStock,
		PromotionPerLimit: product.PromotionPerLimit,

		// 冗余字段
		BrandName:           brandNames[product.BrandID],
		ProductCategoryName: categoryNames[product.ProductCategoryID],
	}
}

// AddOrUpdateProductParamToEntity pb转entity
func AddOrUpdateProductParamToEntity(param *pb.AddOrUpdateProductParam) *entity.Product {
	return &entity.Product{
		// 基本信息
		ProductCategoryID: param.ProductCategoryId,
		Name:              param.Name,
		SubTitle:          param.SubTitle,
		BrandID:           param.BrandId,
		Description:       param.Description,
		ProductSN:         param.ProductSn,
		Price:             param.Price,
		OriginalPrice:     param.OriginalPrice,
		Stock:             param.Stock,
		Unit:              param.Unit,
		Weight:            param.Weight,
		Sort:              param.Sort,

		// 促销信息
		GiftPoint:          param.GiftPoint,
		GiftGrowth:         param.GiftGrowth,
		UsePointLimit:      param.UsePointLimit,
		PreviewStatus:      uint8(param.PreviewStatus),
		PublishStatus:      uint8(param.PublishStatus),
		NewStatus:          uint8(param.NewStatus),
		RecommandStatus:    uint8(param.RecommandStatus),
		ServiceIDs:         param.ServiceIds,
		DetailTitle:        param.DetailTitle,
		DetailDesc:         param.DetailDesc,
		Keywords:           param.Keywords,
		Note:               param.Note,
		PromotionType:      uint8(param.PromotionType),
		PromotionPrice:     param.PromotionPrice,
		PromotionStartTime: param.PromotionStartTime,
		PromotionEndTime:   param.PromotionEndTime,

		// 属性信息
		ProductAttributeCategoryID: param.ProductAttributeCategoryId,
		Pic:                        util.GetRelativeUrl(param.Pic),
		AlbumPics:                  param.AlbumPics,
		DetailHTML:                 param.DetailHtml,
		DetailMobileHTML:           param.DetailMobileHtml,
		VerifyStatus:               uint8(param.VerifyStatus),
		DeleteStatus:               uint8(param.DeleteStatus),

		// 其他
		FeightTemplateID:  param.FeightTemplateId,
		Sale:              param.Sale,
		LowStock:          param.LowStock,
		PromotionPerLimit: param.PromotionPerLimit,

		// 冗余字段
		BrandName:           param.BrandName,
		ProductCategoryName: param.ProductCategoryName,
	}
}
