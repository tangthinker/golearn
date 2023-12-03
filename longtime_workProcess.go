package main

import (
	"fmt"
	"time"
)

func main() {
	go worker()
	for {
	}
}

func worker() {
	num := 0
	next := time.After(time.Second)
	for {
		select {
		case <-next:
			num++
			fmt.Println(num)
			//为下一次事件循环创建新的计时器通道
			next = time.After(time.Second)
		}
	}
}
