package main

import (
	"fmt"
	"sync"
)

type semaphore struct {
	sync.Locker // 实现locker接口
	ch          chan struct{}
}

func NewSemaphore(capacity int) sync.Locker {
	if capacity <= 0 {
		capacity = 1 // 如果capacity为则为互斥锁
	}
	return &semaphore{
		ch: make(chan struct{}, capacity),
	}
}

func (sem *semaphore) Lock() {
	sem.ch <- struct{}{}
}

func (sem *semaphore) Unlock() {
	<-sem.ch
}

func main() {
	sem := NewSemaphore(10)
	sem.Lock()
	sem.Lock()
	sem.Lock()
	sem.Lock()
	sem.Lock()
	sem.Lock()
	sem.Lock()
	sem.Lock()
	sem.Lock()
	sem.Lock()
	fmt.Println("lock 10 times ok")
}
