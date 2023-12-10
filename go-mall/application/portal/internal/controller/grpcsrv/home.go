package grpcsrv

import pb "github.com/baker-yuan/go-mall/proto/mall"

var _ pb.PortalApiServer = &PortalApiImpl{}

type PortalApi interface {
	pb.PortalApiServer
}

type PortalApiImpl struct {
	pb.UnimplementedPortalApiServer
}

func New() PortalApi {
	return nil
}
