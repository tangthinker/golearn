package main

import (
	"context"
	"fmt"
	"sync"
)

type Pool struct {
	taskQueue  chan func()
	workerSize int
	wg         sync.WaitGroup
}

func NewPool(queueCapacity int, workerSize int) *Pool {
	pool := &Pool{
		taskQueue:  make(chan func(), queueCapacity),
		workerSize: workerSize,
	}
	pool.wg.Add(workerSize)
	return pool
}

func (pool *Pool) Start() {
	for i := 0; i < pool.workerSize; i++ {
		go func() {
			defer pool.wg.Done()
			for {
				f, ok := <-pool.taskQueue
				if !ok {
					return
				}
				f()
			}
		}()
	}
}

func (pool *Pool) Submit(f func()) {
	pool.taskQueue <- f
}

func (pool *Pool) Close() {
	close(pool.taskQueue)
	pool.wg.Wait()
}

func main() {
	pool := NewPool(10, 5)
	context.Background()
	pool.Start()
	for i := 0; i <= 100; i++ {
		i := i
		pool.Submit(func() {
			fmt.Println("Hello world!", i)
		})
	}
	pool.Close()
}
