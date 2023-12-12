package interceptor

import (
	"bytes"
	"encoding/json"
	"net/http"
)

// StandardResponse 统一的响应结构
type StandardResponse struct {
	Code    int32       `json:"code"`
	Message string      `json:"message"`
	Data    interface{} `json:"data,omitempty"`
}

// WrapResponseMiddleware 中间件函数，用于包装响应
func WrapResponseMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		// 拦截响应
		wrappedWriter := &responseWriterInterceptor{ResponseWriter: w, body: &bytes.Buffer{}}
		next.ServeHTTP(wrappedWriter, r)

		// 如果响应已经被写入，则不进行包装
		if wrappedWriter.written {
			return
		}

		// 检查响应是否已经是标准格式，返回的内容如果是非结构体（结构体数组、map），需要自己定义code+message+data
		var standardResp StandardResponse
		if err := json.Unmarshal(wrappedWriter.body.Bytes(), &standardResp); err == nil && standardResp.Code != 0 {
			// 响应已经是标准格式，直接写入原始的响应Writer
			w.Header().Set("Content-Type", "application/json")
			_, _ = w.Write(wrappedWriter.body.Bytes())
			return
		}

		// 解析原始响应体
		var data interface{}
		if wrappedWriter.body.Len() > 0 {
			if err := json.Unmarshal(wrappedWriter.body.Bytes(), &data); err != nil {
				http.Error(w, "Internal Server Error", http.StatusInternalServerError)
				return
			}
		}

		// 创建统一的响应结构
		response := StandardResponse{
			Code:    http.StatusOK, // 或者其他适当的成功状态码
			Message: "OK",
			Data:    data,
		}

		// 将统一的响应结构写入原始的响应Writer
		w.Header().Set("Content-Type", "application/json")
		if err := json.NewEncoder(w).Encode(response); err != nil {
			http.Error(w, "Internal Server Error", http.StatusInternalServerError)
		}
	})
}

// responseWriterInterceptor 是一个拦截器，用于捕获响应体
type responseWriterInterceptor struct {
	http.ResponseWriter
	body    *bytes.Buffer
	written bool
}

func (w *responseWriterInterceptor) WriteHeader(statusCode int) {
	if w.written {
		return
	}
	w.ResponseWriter.WriteHeader(statusCode)
	w.written = true
}

func (w *responseWriterInterceptor) Write(b []byte) (int, error) {
	if !w.written {
		// 在写入响应之前，先将响应体捕获到缓冲区中
		return w.body.Write(b)
	}
	return w.ResponseWriter.Write(b)
}
