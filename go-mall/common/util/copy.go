package util

import "encoding/json"

// CopyProperties 使用JSON序列化和反序列化来实现属性拷贝的泛型函数
func CopyProperties[Dst any](src interface{}) (Dst, error) {
	var dst Dst
	// 将源结构体序列化为JSON
	srcJSON, err := json.Marshal(src)
	if err != nil {
		return dst, err
	}
	// 将JSON反序列化为目标结构体
	err = json.Unmarshal(srcJSON, &dst)
	if err != nil {
		return dst, err
	}
	return dst, nil
}
