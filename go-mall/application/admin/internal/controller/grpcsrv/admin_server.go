package grpcsrv

import (
	"github.com/baker-yuan/go-mall/application/admin/internal/usecase"
	pb "github.com/baker-yuan/go-mall/proto/mall"
)

var _ pb.CmsAdminApiServer = &AdminApiImpl{}
var _ pb.OmsAdminApiServer = &AdminApiImpl{}

type AdminApi interface {
	pb.CmsAdminApiServer
	pb.OmsAdminApiServer
}

type AdminApiImpl struct {
	pb.UnimplementedCmsAdminApiServer
	pb.UnimplementedOmsAdminApiServer

	category                 usecase.IProductCategoryUseCase
	brand                    usecase.IBrandUseCase
	productAttributeCategory usecase.IProductAttributeCategoryUseCase
	productAttribute         usecase.IProductAttributeUseCase
	product                  usecase.IProductUseCase
	skuStock                 usecase.ISkuStockUseCase
	subject                  usecase.ISubjectUseCase
	prefrenceArea            usecase.IPrefrenceAreaUseCase
	//
	orderReturnReason usecase.IOrderReturnReasonUseCase
	order             usecase.IOrderUseCase
	orderReturnApply  usecase.IOrderReturnApplyUseCase
}

func New(category usecase.IProductCategoryUseCase,
	brand usecase.IBrandUseCase,
	productAttributeCategory usecase.IProductAttributeCategoryUseCase,
	productAttribute usecase.IProductAttributeUseCase,
	product usecase.IProductUseCase,
	skuStock usecase.ISkuStockUseCase,
	subject usecase.ISubjectUseCase,
	prefrenceArea usecase.IPrefrenceAreaUseCase,
	//
	orderReturnReason usecase.IOrderReturnReasonUseCase,
	order usecase.IOrderUseCase,
	orderReturnApply usecase.IOrderReturnApplyUseCase,
) AdminApi {
	return &AdminApiImpl{
		category:                 category,
		brand:                    brand,
		productAttributeCategory: productAttributeCategory,
		productAttribute:         productAttribute,
		product:                  product,
		skuStock:                 skuStock,
		subject:                  subject,
		prefrenceArea:            prefrenceArea,
		//
		orderReturnReason: orderReturnReason,
		order:             order,
		orderReturnApply:  orderReturnApply,
	}

}
