package retcode

const (
	RetSuccess       = 0
	RetInternalError = 2
)

var (
	RetMsg = map[uint32]string{
		RetSuccess:       "请求成功",
		RetInternalError: "服务器异常",
	}
)

// GetRetCodeMsg 根据错误码获取错误码及错误信息指针
func GetRetCodeMsg(code uint32) (uint32, string) {
	return code, RetMsg[code]
}
