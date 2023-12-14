package util

import (
	"context"
	"fmt"
	"strings"

	"google.golang.org/grpc/metadata"
)

// 定义一个类型，用作上下文中的键
type contextKey string

// 定义上下文中使用的键
const (
	UsernameKey     contextKey = "username"     // 用户名
	UserIDKey       contextKey = "userId"       // 用户名
	RequestIDHeader contextKey = "X-Request-ID" // 请求ID
)

// SetUserID 将用户id设置到上下文中
func SetUserID(ctx context.Context, userID uint64) context.Context {
	return context.WithValue(ctx, UserIDKey, fmt.Sprintf("%d", userID))
}

// GetUserID 从上下文中获取用户ID
func GetUserID(ctx context.Context) (uint64, bool) {
	userID, ok := getFromHttpContext(ctx, UserIDKey)
	if ok {
		return userID, ok
	}
	return getFromGrpcContext(ctx, UserIDKey)
}

func getFromHttpContext(ctx context.Context, key contextKey) (uint64, bool) {
	value, ok := ctx.Value(key).(string)
	if !ok {
		return 0, false
	}
	userID, _ := StringToUint64(value)
	return userID, true
}

func getFromGrpcContext(ctx context.Context, key contextKey) (uint64, bool) {
	md, ok := metadata.FromIncomingContext(ctx)
	if !ok {
		return 0, false
	}
	slice, ok := md[strings.ToLower(string(key))]
	if !ok || len(slice) == 0 {
		return 0, false
	}
	userID, _ := StringToUint64(slice[0])
	return userID, true
}
