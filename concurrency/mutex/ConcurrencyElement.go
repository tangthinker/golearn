package main

import (
	"fmt"
	"sync"
	"time"
)

type ConcurrencyElement[T any] struct {
	sync.Mutex
	data T
}

func NewConcurrencyElement[T any](initial T) *ConcurrencyElement[T] {
	var ele ConcurrencyElement[T]
	ele.data = initial
	return &ele
}

/*
有时候读写并不是直接访问变量，而是有一些复杂的逻辑
传入函数访问和修改变量，基本上为visitor模式
*/
func (ele *ConcurrencyElement[T]) Access(f func(value *T)) {
	ele.Mutex.Lock()
	defer ele.Mutex.Unlock()

	data := ele.data
	f(&data)
	ele.data = data
	// 这种修改值的手法太棒了
}

func (ele *ConcurrencyElement[T]) Load() T {
	ele.Lock()
	defer ele.Unlock()

	return ele.data
}

func (ele *ConcurrencyElement[T]) Store(data T) {
	ele.Lock()
	defer ele.Unlock()

	ele.data = data
}

func main() {
	concurrencyInt := NewConcurrencyElement[int](22)

	go func() {
		concurrencyInt.Access(func(value *int) {
			*value = 33
		})
	}()

	go func() {
		concurrencyInt.Store(22)
	}()

	time.Sleep(time.Second)
	fmt.Println(concurrencyInt.Load())
}
