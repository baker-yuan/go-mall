package grpcsrv

import (
	"github.com/baker-yuan/go-mall/application/portal/internal/usecase"
	pb "github.com/baker-yuan/go-mall/proto/mall"
)

var _ pb.PortalHomeApiServer = &PortalApiImpl{}
var _ pb.PortalProductApiServer = &PortalApiImpl{}
var _ pb.PortalMemberApiServer = &PortalApiImpl{}

type PortalApi interface {
	pb.PortalHomeApiServer
	pb.PortalProductApiServer
	pb.PortalMemberApiServer
}

type PortalApiImpl struct {
	pb.UnimplementedPortalHomeApiServer
	pb.UnimplementedPortalProductApiServer
	pb.UnimplementedPortalMemberApiServer

	home          usecase.IHomeUseCase
	product       usecase.IProductUseCase
	memberUseCase usecase.IMemberUseCase
}

func New(
	home usecase.IHomeUseCase,
	product usecase.IProductUseCase,
	memberUseCase usecase.IMemberUseCase,
) PortalApi {
	return &PortalApiImpl{
		home:          home,
		product:       product,
		memberUseCase: memberUseCase,
	}
}
