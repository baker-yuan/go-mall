package retcode

import (
	"net/http"

	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

const (
	RetSuccess       codes.Code = 200
	RetInternalError codes.Code = 500
)

// customError 自定义错误结构体
type customError struct {
	HTTPStatus int    // HTTP状态码
	ErrorCode  int32  // 自定义错误码
	Message    string // 错误消息
}

// customErrorMap 是一个映射，它将 gRPC 状态码映射到自定义错误
var customErrorMap = map[codes.Code]customError{
	// grpc异常
	codes.OK: {
		HTTPStatus: http.StatusOK,
		Message:    "Success",
	},
	codes.Canceled: {
		HTTPStatus: 499,
		Message:    "Request canceled by the client",
	},
	codes.Unknown: {
		HTTPStatus: 500,
		Message:    "Unknown error occurred",
	},
	codes.InvalidArgument: {
		HTTPStatus: 400,
		Message:    "Invalid argument provided",
	},
	codes.DeadlineExceeded: {
		HTTPStatus: 504,
		Message:    "Deadline exceeded",
	},
	codes.NotFound: {
		HTTPStatus: 404,
		Message:    "Resource not found",
	},
	codes.AlreadyExists: {
		HTTPStatus: 409,
		Message:    "Resource already exists",
	},
	codes.PermissionDenied: {
		HTTPStatus: 403,
		Message:    "Permission denied",
	},
	codes.ResourceExhausted: {
		HTTPStatus: 429,
		Message:    "Resource exhausted",
	},
	codes.FailedPrecondition: {
		HTTPStatus: 412,
		Message:    "Failed precondition",
	},
	codes.Aborted: {
		HTTPStatus: 409,
		Message:    "Operation aborted",
	},
	codes.OutOfRange: {
		HTTPStatus: 400,
		Message:    "Out of range",
	},
	codes.Unimplemented: {
		HTTPStatus: 501,
		Message:    "Not implemented",
	},
	codes.Internal: {
		HTTPStatus: 500,
		Message:    "Internal server error",
	},
	codes.Unavailable: {
		HTTPStatus: 503,
		Message:    "Service unavailable",
	},
	codes.DataLoss: {
		HTTPStatus: 500,
		Message:    "Data loss",
	},
	codes.Unauthenticated: {
		HTTPStatus: 401,
		Message:    "Unauthenticated",
	},
	// 业务异常
	RetSuccess: {
		HTTPStatus: 200,
		Message:    "请求成功",
	},
	RetInternalError: {
		HTTPStatus: http.StatusInternalServerError,
		Message:    "服务器异常",
	},
}

// GetRetCodeMsg 根据错误码获取错误码及错误信息
func GetRetCodeMsg(code codes.Code) (uint32, string) {
	return uint32(code), customErrorMap[code].Message
}

// GetHttpCodeAndMsg 根据错误码获取http错误码及错误信息
func GetHttpCodeAndMsg(code codes.Code) (uint32, string) {
	if _, exist := customErrorMap[code]; exist {
		return uint32(customErrorMap[code].HTTPStatus), customErrorMap[code].Message
	} else {
		return uint32(customErrorMap[RetInternalError].HTTPStatus), customErrorMap[RetInternalError].Message
	}
}

// NewError 创建自定义异常
func NewError(code codes.Code, message string) error {
	return status.New(code, message).Err()
}
