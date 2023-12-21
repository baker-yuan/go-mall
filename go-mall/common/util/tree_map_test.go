package util

import (
	"testing"
)

// TestTreeMapInsertAndGet 测试 TreeMap 的插入和获取功能。
func TestTreeMapInsertAndGet(t *testing.T) {
	// 创建一个新的 TreeMap 实例。
	treeMap := NewTreeMap[int, string]()
	// 定义一系列测试数据和期望结果。
	tests := []struct {
		key      int
		value    string
		expected bool
	}{
		{1, "A", true},
		{2, "B", true},
		{3, "C", true},
		{4, "D", true},
		{5, "E", true},
	}

	// 遍历测试数据，插入并尝试获取键值对。
	for _, test := range tests {
		treeMap.Insert(test.key, test.value)
		if val, ok := treeMap.Get(test.key); !ok || val != test.value {
			t.Errorf("Get(%d) = %v, %v; want %v, true", test.key, val, ok, test.value)
		}
	}

	// 测试获取不存在的键。
	if _, ok := treeMap.Get(999); ok {
		t.Error("预期 Get(999) 返回 false，但是得到了 true")
	}
}

// TestTreeMapUpdate 测试 TreeMap 的更新功能。
func TestTreeMapUpdate(t *testing.T) {
	// 创建一个新的 TreeMap 实例。
	treeMap := NewTreeMap[int, string]()
	key := 1
	initialValue := "A"
	updatedValue := "Updated"

	// 插入一个键值对，然后更新它的值。
	treeMap.Insert(key, initialValue)
	treeMap.Insert(key, updatedValue) // 更新同一个键的值。

	// 验证更新是否成功。
	if val, ok := treeMap.Get(key); !ok || val != updatedValue {
		t.Errorf("Get(%d) = %v, %v; want %v, true", key, val, ok, updatedValue)
	}
}
