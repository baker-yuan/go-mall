package grpcsrv

import (
	"github.com/baker-yuan/go-mall/application/admin/internal/usecase"
	pb "github.com/baker-yuan/go-mall/proto/mall"
)

var _ pb.AdminApiServer = &AdminApiImpl{}

type AdminApiImpl struct {
	pb.UnimplementedAdminApiServer

	category                 usecase.IProductCategoryUseCase
	brand                    usecase.IBrandUseCase
	productAttributeCategory usecase.IProductAttributeCategoryUseCase
	productAttribute         usecase.IProductAttributeUseCase
	product                  usecase.IProductUseCase
}

func New(category usecase.IProductCategoryUseCase,
	brand usecase.IBrandUseCase,
	productAttributeCategory usecase.IProductAttributeCategoryUseCase,
	productAttribute usecase.IProductAttributeUseCase,
	product usecase.IProductUseCase,
) pb.AdminApiServer {
	return &AdminApiImpl{
		category:                 category,
		brand:                    brand,
		productAttributeCategory: productAttributeCategory,
		productAttribute:         productAttribute,
		product:                  product,
	}

}
