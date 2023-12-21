package grpcsrv

import (
	"github.com/baker-yuan/go-mall/application/portal/internal/usecase"
	pb "github.com/baker-yuan/go-mall/proto/mall"
)

var _ pb.PortalHomeApiServer = &PortalApiImpl{}
var _ pb.PortalProductApiServer = &PortalApiImpl{}
var _ pb.PortalMemberApiServer = &PortalApiImpl{}
var _ pb.PortalCartItemApiServer = &PortalApiImpl{}
var _ pb.PortalOrderApiServer = &PortalApiImpl{}

type PortalApi interface {
	pb.PortalHomeApiServer
	pb.PortalProductApiServer
	pb.PortalMemberApiServer
	pb.PortalCartItemApiServer
	pb.PortalOrderApiServer
}

type PortalApiImpl struct {
	pb.UnimplementedPortalHomeApiServer
	pb.UnimplementedPortalProductApiServer
	pb.UnimplementedPortalMemberApiServer
	pb.UnimplementedPortalCartItemApiServer
	pb.UnimplementedPortalOrderApiServer

	home            usecase.IHomeUseCase
	product         usecase.IProductUseCase
	memberUseCase   usecase.IMemberUseCase
	cartItemUseCase usecase.ICartItemUseCase
	orderUseCase    usecase.IOrderUseCase
}

func New(
	home usecase.IHomeUseCase,
	product usecase.IProductUseCase,
	memberUseCase usecase.IMemberUseCase,
	cartItemUseCase usecase.ICartItemUseCase,
	orderUseCase usecase.IOrderUseCase,
) PortalApi {
	return &PortalApiImpl{
		home:            home,
		product:         product,
		memberUseCase:   memberUseCase,
		cartItemUseCase: cartItemUseCase,
		orderUseCase:    orderUseCase,
	}
}
