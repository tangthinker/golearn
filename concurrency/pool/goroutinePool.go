package main

import "sync"

type GoroutinePool[T any] struct {
	taskQueue    chan T
	taskFunction func(T)
	workerSize   int
	wg           sync.WaitGroup
}

func NewGoroutinePool[T any](workerSize int, queueCapacity int, taskFunction func(T)) *GoroutinePool[T] {
	goroutinePool := &GoroutinePool[T]{
		taskQueue:    make(chan T, queueCapacity),
		taskFunction: taskFunction,
		workerSize:   workerSize,
	}
	goroutinePool.wg.Add(workerSize)
	return goroutinePool
}

func (pool *GoroutinePool[T]) Start() {
	for i := 0; i < pool.workerSize; i++ {
		go func() {
			defer pool.wg.Done()
			for {
				taskArg, ok := <-pool.taskQueue
				if !ok {
					return
				}
				pool.taskFunction(taskArg)
			}
		}()
	}
}

func (pool *GoroutinePool[T]) Submit(taskArg T) {
	pool.taskQueue <- taskArg
}

func (pool *GoroutinePool[T]) Close() {
	close(pool.taskQueue)
	pool.wg.Wait()
}

func main() {

}
