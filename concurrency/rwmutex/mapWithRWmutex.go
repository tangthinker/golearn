package main

import (
	"fmt"
	"sync"
)

/*
读写锁并发map
对于读多写少的场景很实用
此外还有
分片加锁的map(orcaman/concurrent-map)
和lock-free的map(cornelk/hashmap|alphadose/haxmap)
*/
type RWMap[k comparable, v any] struct {
	sync.RWMutex
	data map[k]v
}

func NewRWMap[k comparable, v any](n int) *RWMap[k, v] {
	return &RWMap[k, v]{
		data: make(map[k]v, n),
	}
}

func (m *RWMap[k, v]) Get(key k) (v, bool) {
	m.RLock()
	defer m.RUnlock()
	value, ok := m.data[key]
	return value, ok
}

func (m *RWMap[k, v]) Set(key k, value v) {
	m.Lock()
	defer m.Unlock()
	m.data[key] = value
}

func (m *RWMap[k, v]) Delete(key k) {
	m.Lock()
	defer m.Unlock()
	delete(m.data, key)
}

func (m *RWMap[k, v]) Len() int {
	m.RLock()
	defer m.RUnlock()
	return len(m.data)
}

func (m *RWMap[k, v]) Range(f func(key k, value v) bool) {
	m.RLock()
	defer m.RUnlock()
	for key, value := range m.data {
		if !f(key, value) { // 提前结束遍历
			return
		}
	}
}

func main() {
	rwmap := NewRWMap[int, string](3)
	rwmap.Set(0, "shanliao")
	rwmap.Set(1, "king")
	rwmap.Set(2, "kkkkk")

	rwmap.Range(func(key int, value string) bool {
		fmt.Printf("no.%d: %s\n", key, value)
		return true
	})
}
