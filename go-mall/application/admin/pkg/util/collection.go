package util

// SliceRemove 移除切片元素
func SliceRemove[T comparable](slice []T, removes ...T) []T {
	res := make([]T, 0)
	for _, item := range slice {
		exist := false
		for _, remove := range removes {
			if item == remove {
				exist = true
				break
			}
		}
		if exist {
			continue
		}
		res = append(res, item)
	}
	return res
}
