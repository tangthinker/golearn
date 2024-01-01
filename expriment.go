package main

import (
	"fmt"
	"time"
)

const (
	mutexLocked = 1 << iota
	mutexWoken
	mutexWaiterShift = iota
)

func main() {
	ch := make(chan string, 1)

	go func() {
		for i := 0; i < 2; i++ {
			name := <-ch
			fmt.Println(name)
		}
	}()
	ch <- "shanliao"
	ch <- "shanliao"
	ch <- "shanliao"
	time.Sleep(time.Second)
}
