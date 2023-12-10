package grpcsrv

import (
	"github.com/baker-yuan/go-mall/application/portal/internal/usecase"
	pb "github.com/baker-yuan/go-mall/proto/mall"
)

var _ pb.PortalHomeApiServer = &PortalApiImpl{}
var _ pb.PortalProductApiServer = &PortalApiImpl{}

type PortalApi interface {
	pb.PortalHomeApiServer
	pb.PortalProductApiServer
}

type PortalApiImpl struct {
	pb.UnimplementedPortalHomeApiServer
	pb.UnimplementedPortalProductApiServer

	home    usecase.IHomeUseCase
	product usecase.IProductUseCase
}

func New(
	home usecase.IHomeUseCase,
	product usecase.IProductUseCase,
) PortalApi {
	return &PortalApiImpl{
		home:    home,
		product: product,
	}
}
