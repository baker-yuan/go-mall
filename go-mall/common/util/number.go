package util

import "strconv"

// StringToUint64 字符串转换为 uint64 的函数
func StringToUint64(s string) (uint64, error) {
	value, err := strconv.ParseUint(s, 10, 64)
	if err != nil {
		return 0, err
	}
	return value, nil
}
