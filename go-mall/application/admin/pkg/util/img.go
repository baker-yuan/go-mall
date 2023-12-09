package util

import "strings"

var baseUrl string

func InitBaseUrl(base string) {
	baseUrl = base
}

// GetFullUrls 获取完整路径
func GetFullUrls(paths []string) []string {
	res := make([]string, 0)
	for _, path := range paths {
		res = append(res, GetFullUrl(path))
	}
	return res
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

// GetRelativeUrls 获取相对路径
func GetRelativeUrls(urls []string) []string {
	res := make([]string, 0)
	for _, url := range urls {
		res = append(res, GetRelativeUrl(url))
	}
	return res
}
