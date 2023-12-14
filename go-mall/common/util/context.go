package util

import (
	"context"

	"google.golang.org/grpc/metadata"
)

// 定义一个类型，用作上下文中的键
type contextKey string

// 定义上下文中使用的键
const (
	UsernameKey     contextKey = "username"     // 用户名
	RequestIDHeader contextKey = "X-Request-ID" // 请求ID
)

// SetUsername 将用户名设置到上下文中
func SetUsername(ctx context.Context, username string) context.Context {
	return context.WithValue(ctx, UsernameKey, username)
}

// GetUsername 从上下文中获取用户名
func GetUsername(ctx context.Context) (string, bool) {
	username, ok := getFromGrpcContext(ctx, UsernameKey)
	if ok {
		return username, ok
	}
	return getFromHttpContext(ctx, UsernameKey)
}

func getFromHttpContext(ctx context.Context, key contextKey) (string, bool) {
	value, ok := ctx.Value(key).(string)
	return value, ok
}

func getFromGrpcContext(ctx context.Context, key contextKey) (string, bool) {
	md, ok := metadata.FromIncomingContext(ctx)
	if !ok {
		return "", false
	}
	slice, ok := md[string(key)]
	if !ok || len(slice) == 0 {
		return "", false
	}
	return slice[0], true
}
