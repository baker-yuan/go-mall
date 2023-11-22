package util

var baseUrl string

func InitBaseUrl(base string) {
	baseUrl = base
}

func GetFullUrl(path string) string {
	return baseUrl + path
}
