package interceptor

import (
	"context"
	"net/http"
	"strings"

	pkg_jwt "github.com/baker-yuan/go-mall/common/pkg/jwt"
	"github.com/baker-yuan/go-mall/common/util"
	"google.golang.org/grpc"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/metadata"
	"google.golang.org/grpc/status"
)

// NewJWTAuthMiddleware 是一个用于验证 JWT 令牌的 HTTP 中间件
func NewJWTAuthMiddleware(next http.Handler, jwtTokenUtil *pkg_jwt.JWT, whitelist []string, tokenHeader string, tokenHead string) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		// 检查当前请求的路径是否在白名单中
		for _, path := range whitelist {
			if strings.HasPrefix(r.URL.Path, path) {
				// 如果在白名单中，直接传递给下一个处理器
				next.ServeHTTP(w, r)
				return
			}
		}

		// 执行 JWT 认证逻辑
		authHeader := r.Header.Get(tokenHeader)
		if authHeader == "" {
			http.Error(w, "Authorization header is required", http.StatusUnauthorized)
			return
		}

		bearerToken := strings.TrimPrefix(authHeader, tokenHead)

		token, err := jwtTokenUtil.ParseToken(bearerToken)
		if err != nil {
			http.Error(w, "Invalid token", http.StatusUnauthorized)
			return
		}

		ctx := util.SetUsername(r.Context(), token.UserInfo.Username)

		// 将上下文传递到下一个处理器
		next.ServeHTTP(w, r.WithContext(ctx))
	})
}

// CustomAnnotator 将 HTTP 上下文中的用户名添加到 gRPC 上下文的元数据中
func CustomAnnotator(ctx context.Context, req *http.Request) metadata.MD {
	// 从HTTP context中获取用户名
	username, exist := util.GetUsername(ctx)
	if !exist {
		return nil
	}

	// 创建gRPC metadata并添加用户名
	md := metadata.Pairs("username", username)
	return md
}

// JWTAuthInterceptor 创建一个新的 Unary 拦截器，用于验证 JWT 令牌。
func JWTAuthInterceptor(jwtTokenUtil *pkg_jwt.JWT, whitelist []string, tokenHeader string, tokenHead string) grpc.UnaryServerInterceptor {
	return func(
		ctx context.Context,
		req interface{},
		info *grpc.UnaryServerInfo,
		handler grpc.UnaryHandler,
	) (interface{}, error) {
		// 检查当前请求的方法是否在白名单中
		//for _, path := range whitelist {
		//	if strings.HasPrefix(info.FullMethod, path) {
		//		// 如果在白名单中，直接调用 RPC 方法
		//		return handler(ctx, req)
		//	}
		//}

		// 从 gRPC 元数据中获取 Authorization 头
		md, ok := metadata.FromIncomingContext(ctx)
		if !ok {
			return nil, status.Errorf(codes.Unauthenticated, "Authorization metadata is required")
		}

		authHeader, ok := md[strings.ToLower(tokenHeader)]
		if !ok || len(authHeader) == 0 {
			return nil, status.Errorf(codes.Unauthenticated, "Authorization header is required")
		}

		bearerToken := strings.TrimPrefix(authHeader[0], tokenHead)

		token, err := jwtTokenUtil.ParseToken(bearerToken)
		if err != nil {
			return nil, status.Errorf(codes.Unauthenticated, "Invalid token: %v", err)
		}

		// 将用户名设置到 gRPC 上下文中
		newCtx := util.SetUsername(ctx, token.UserInfo.Username)

		// 调用 RPC 方法
		return handler(newCtx, req)
	}
}
