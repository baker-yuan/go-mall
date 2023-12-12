package retcode

import "google.golang.org/grpc/codes"

const (
	RetSuccess       codes.Code = 200
	RetInternalError codes.Code = 500
)

//
//var (
//	RetMsg = map[codes.Code]string{
//		RetSuccess:       "请求成功",
//		RetInternalError: "服务器异常",
//	}
//)

// customError 自定义错误结构体
type customError struct {
	HTTPStatus int    // HTTP 状态码
	ErrorCode  int32  // 自定义错误码
	Message    string // 错误消息
}

// customErrorMap 是一个映射，它将 gRPC 状态码映射到自定义错误
var customErrorMap = map[codes.Code]customError{
	// grpc异常
	codes.OK: {
		HTTPStatus: 200,
		ErrorCode:  1000,
		Message:    "Success",
	},
	codes.Canceled: {
		HTTPStatus: 499,
		ErrorCode:  1004,
		Message:    "Request canceled by the client",
	},
	codes.Unknown: {
		HTTPStatus: 500,
		ErrorCode:  1005,
		Message:    "Unknown error occurred",
	},
	codes.InvalidArgument: {
		HTTPStatus: 400,
		ErrorCode:  1001,
		Message:    "Invalid argument provided",
	},
	codes.DeadlineExceeded: {
		HTTPStatus: 504,
		ErrorCode:  1006,
		Message:    "Deadline exceeded",
	},
	codes.NotFound: {
		HTTPStatus: 404,
		ErrorCode:  1002,
		Message:    "Resource not found",
	},
	codes.AlreadyExists: {
		HTTPStatus: 409,
		ErrorCode:  1007,
		Message:    "Resource already exists",
	},
	codes.PermissionDenied: {
		HTTPStatus: 403,
		ErrorCode:  1008,
		Message:    "Permission denied",
	},
	codes.ResourceExhausted: {
		HTTPStatus: 429,
		ErrorCode:  1009,
		Message:    "Resource exhausted",
	},
	codes.FailedPrecondition: {
		HTTPStatus: 412,
		ErrorCode:  1010,
		Message:    "Failed precondition",
	},
	codes.Aborted: {
		HTTPStatus: 409,
		ErrorCode:  1011,
		Message:    "Operation aborted",
	},
	codes.OutOfRange: {
		HTTPStatus: 400,
		ErrorCode:  1012,
		Message:    "Out of range",
	},
	codes.Unimplemented: {
		HTTPStatus: 501,
		ErrorCode:  1013,
		Message:    "Not implemented",
	},
	codes.Internal: {
		HTTPStatus: 500,
		ErrorCode:  1003,
		Message:    "Internal server error",
	},
	codes.Unavailable: {
		HTTPStatus: 503,
		ErrorCode:  1014,
		Message:    "Service unavailable",
	},
	codes.DataLoss: {
		HTTPStatus: 500,
		ErrorCode:  1015,
		Message:    "Data loss",
	},
	codes.Unauthenticated: {
		HTTPStatus: 401,
		ErrorCode:  1016,
		Message:    "Unauthenticated",
	},
	// 业务异常

}

// GetRetCodeMsg 根据错误码获取错误码及错误信息指针
func GetRetCodeMsg(code codes.Code) (uint32, string) {
	return uint32(code), customErrorMap[code].Message
}
