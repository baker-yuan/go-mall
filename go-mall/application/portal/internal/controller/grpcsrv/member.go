package grpcsrv

import (
	"context"

	pb "github.com/baker-yuan/go-mall/proto/mall"
)

func (s PortalApiImpl) MemberRegister(ctx context.Context, req *pb.MemberRegisterReq) (*pb.EmptyRsp, error) {
	//TODO implement me
	panic("implement me")
}

func (s PortalApiImpl) MemberLogin(ctx context.Context, req *pb.MemberLoginReq) (*pb.MemberLoginRsp, error) {
	//TODO implement me
	panic("implement me")
}

func (s PortalApiImpl) MemberInfo(ctx context.Context, req *pb.MemberInfoReq) (*pb.MemberInfoRsp, error) {
	//TODO implement me
	panic("implement me")
}

func (s PortalApiImpl) MemberGetAuthCode(ctx context.Context, req *pb.MemberGetAuthCodeReq) (*pb.MemberGetAuthCodeRsp, error) {
	//TODO implement me
	panic("implement me")
}

func (s PortalApiImpl) MemberUpdatePassword(ctx context.Context, req *pb.MemberUpdatePasswordReq) (*pb.MemberUpdatePasswordRsp, error) {
	//TODO implement me
	panic("implement me")
}

func (s PortalApiImpl) MemberRefreshToken(ctx context.Context, req *pb.MemberRefreshTokenReq) (*pb.MemberRefreshTokenRsp, error) {
	//TODO implement me
	panic("implement me")
}
