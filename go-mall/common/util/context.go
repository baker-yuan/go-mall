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

	username, ok := GetUsernameFromMd(ctx)
	if ok {
		return username, ok
	}
	//
	username, ok = ctx.Value(UsernameKey).(string)
	return username, ok
}

func GetUsernameFromMd(ctx context.Context) (string, bool) {
	md, ok := metadata.FromIncomingContext(ctx)
	if !ok {
		return "", false
	}
	usernameSlice, ok := md["username"]
	if !ok || len(usernameSlice) == 0 {
		return "", false
	}
	return usernameSlice[0], true
}
