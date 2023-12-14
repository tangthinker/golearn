package main

import (
	"fmt"
	"sync"
	"time"
)

type TaskT struct {
	fn func()
}

type PoolT struct {
	mutex sync.Mutex
	tasks []*TaskT
}

func WorkerT(pool *PoolT) {
	for {
		if len(pool.tasks) <= 0 {
			break
		}
		pool.mutex.Lock()
		top := pool.tasks[0]
		pool.tasks = pool.tasks[1:]
		pool.mutex.Unlock()
		process(top)
	}
}

func process(task *TaskT) {
	go task.fn()
}

func main() {
	pool := PoolT{
		tasks: []*TaskT{
			&TaskT{
				fn: func() {
					fmt.Println("shanliao")
				},
			},
			&TaskT{
				fn: func() {
					fmt.Println("king")
				},
			},
		},
	}
	WorkerT(&pool)
	time.Sleep(time.Second * 1)
}
