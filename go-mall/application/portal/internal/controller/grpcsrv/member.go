package grpcsrv

import (
	"context"

	"github.com/baker-yuan/go-mall/common/util"
	pb "github.com/baker-yuan/go-mall/proto/mall"
)

// MemberRegister 会员注册
func (s PortalApiImpl) MemberRegister(ctx context.Context, req *pb.MemberRegisterReq) (*pb.EmptyRsp, error) {
	//TODO implement me
	panic("implement me")
}

// MemberLogin 会员登录
func (s PortalApiImpl) MemberLogin(ctx context.Context, req *pb.MemberLoginReq) (*pb.MemberLoginRsp, error) {
	return s.memberUseCase.MemberLogin(ctx, req)
}

// MemberInfo 获取会员信息
func (s PortalApiImpl) MemberInfo(ctx context.Context, req *pb.MemberInfoReq) (*pb.MemberInfoRsp, error) {
	memberID, err := util.CtxUtils.GetUserID(ctx)
	if err != nil {
		return nil, err
	}
	return s.memberUseCase.MemberInfo(ctx, memberID)
}

// MemberGetAuthCode 获取验证码
func (s PortalApiImpl) MemberGetAuthCode(ctx context.Context, req *pb.MemberGetAuthCodeReq) (*pb.MemberGetAuthCodeRsp, error) {
	//TODO implement me
	panic("implement me")
}

// MemberUpdatePassword 修改密码
func (s PortalApiImpl) MemberUpdatePassword(ctx context.Context, req *pb.MemberUpdatePasswordReq) (*pb.MemberUpdatePasswordRsp, error) {
	//TODO implement me
	panic("implement me")
}

// MemberRefreshToken 刷新token
func (s PortalApiImpl) MemberRefreshToken(ctx context.Context, req *pb.MemberRefreshTokenReq) (*pb.MemberRefreshTokenRsp, error) {
	//TODO implement me
	panic("implement me")
}
