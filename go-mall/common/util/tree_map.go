package util

import (
	"golang.org/x/exp/constraints"
)

// TreeMap 是一个基于泛型的二叉搜索树实现的映射表。
type TreeMap[K constraints.Ordered, V any] struct {
	root *TreeNode[K, V]
}

// TreeNode 表示 TreeMap 中的一个节点。
type TreeNode[K constraints.Ordered, V any] struct {
	key   K
	value V
	left  *TreeNode[K, V]
	right *TreeNode[K, V]
}

// NewTreeMap 创建一个新的 TreeMap 实例。
func NewTreeMap[K constraints.Ordered, V any]() *TreeMap[K, V] {
	return &TreeMap[K, V]{}
}

// Insert 向 TreeMap 中添加一个新的键值对。
func (t *TreeMap[K, V]) Insert(key K, value V) {
	t.root = insertNode(t.root, key, value)
}

func insertNode[K constraints.Ordered, V any](node *TreeNode[K, V], key K, value V) *TreeNode[K, V] {
	if node == nil {
		return &TreeNode[K, V]{key: key, value: value}
	}
	if key < node.key {
		node.left = insertNode(node.left, key, value)
	} else if key > node.key {
		node.right = insertNode(node.right, key, value)
	} else {
		node.value = value // 如果键已存在，则更新其值。
	}
	return node
}

// Get 根据给定的键检索关联的值。
func (t *TreeMap[K, V]) Get(key K) (V, bool) {
	node := t.root
	for node != nil {
		if key < node.key {
			node = node.left
		} else if key > node.key {
			node = node.right
		} else {
			return node.value, true
		}
	}
	var zeroValue V
	return zeroValue, false
}

// Range 返回一个通道，可以用来遍历 TreeMap 中的所有键值对。
func (t *TreeMap[K, V]) Range() <-chan struct {
	Key   K
	Value V
} {
	ch := make(chan struct {
		Key   K
		Value V
	})
	go func() {
		t.inOrderSend(t.root, ch)
		close(ch)
	}()
	return ch
}

// inOrderSend 是一个递归辅助函数，用于中序遍历树并发送每个节点到通道。
func (t *TreeMap[K, V]) inOrderSend(node *TreeNode[K, V], ch chan<- struct {
	Key   K
	Value V
}) {
	if node == nil {
		return
	}
	t.inOrderSend(node.left, ch) // 遍历左子树
	ch <- struct {
		Key   K
		Value V
	}{
		Key:   node.key,
		Value: node.value,
	} // 发送当前节点
	t.inOrderSend(node.right, ch) // 遍历右子树
}
