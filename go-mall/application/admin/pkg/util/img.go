package util

import "strings"

var baseUrl string

func InitBaseUrl(base string) {
	baseUrl = base
}

// GetFullUrl 获取完整路径
func GetFullUrl(path string) string {
	if len(path) == 0 {
		return ""
	}
	return baseUrl + path
}

// GetRelativeUrl 获取相对路径
func GetRelativeUrl(url string) string {
	if strings.HasPrefix(url, baseUrl) {
		return strings.TrimPrefix(url, baseUrl)
	}
	return url
}
