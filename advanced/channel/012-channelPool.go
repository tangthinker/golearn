package main

import (
	"fmt"
	"time"
)

type Task struct {
	fn     func() string
	result string
}

func sendWork(ch chan *Task) {
	ch <- &Task{
		fn: func() string {
			fmt.Println("shanliao")
			return "shanliao"
		},
	}
	ch <- &Task{
		fn: func() string {
			fmt.Println("king")
			return "king"
		},
	}
}

func Worker(in, out chan *Task) {
	for task := range in {
		task.result = task.fn()
		out <- task
	}
}

func consumerResult(ch chan *Task) {
	for v := range ch {
		fmt.Println(v.result)
	}
}

func main() {
	pending, done := make(chan *Task), make(chan *Task)
	go sendWork(pending)
	handlerThread := 4
	for i := 0; i < handlerThread; i++ {
		go Worker(pending, done)
	}
	go consumerResult(done)
	time.Sleep(time.Second * 1)
}
