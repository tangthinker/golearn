package main

import (
	"fmt"
	"sync"
	"time"
)

type ConcurrencyMap[k comparable, v any] struct {
	sync.Mutex
	m map[k]v
}

// 这里使用变长参数的意义是什么？

func NewConcurrencyMap[k comparable, v any](size ...int) *ConcurrencyMap[k, v] {
	if len(size) > 0 {
		return &ConcurrencyMap[k, v]{
			m: make(map[k]v, size[0]),
		}
	}
	return &ConcurrencyMap[k, v]{
		m: make(map[k]v),
	}
}

func (m *ConcurrencyMap[k, v]) Get(key k) (v, bool) {
	m.Mutex.Lock()
	defer m.Mutex.Unlock()

	value, ok := m.m[key]
	return value, ok
}

func (m *ConcurrencyMap[k, v]) Set(key k, value v) {
	m.Mutex.Lock()
	defer m.Mutex.Unlock()

	m.m[key] = value
}

func main() {
	concurrencyMap := NewConcurrencyMap[int, int](20)
	go func() {
		concurrencyMap.Set(20, 20)
	}()

	go func() {
		value, _ := concurrencyMap.Get(20)
		fmt.Println(value)
	}()
	time.Sleep(time.Second)

}
