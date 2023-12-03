package main

import (
	"fmt"
	"time"
)

func main() {
	channel := make(chan int)
	for i := 0; i < 10; i++ {
		go goSleep(i, channel)
	}
	timeout := time.After(2 * time.Second)
	//time.Sleep(1 * time.Second)
	for i := 0; i < 10; i++ {
		select {
		case id := <-channel:
			fmt.Println("the id", id, " now is wake up!")
		case <-timeout:
			fmt.Println("time out!")
			return
		}
	}

	//deadlock 死锁 goroutine因为某些无法发生的事而永远阻塞
	//c := make(chan int)
	//<-c
}

func goSleep(id int, channel chan int) {
	//t := rand.Intn(5)
	//fmt.Println("id:", id, "my time is ", t)
	time.Sleep(time.Duration(1) * time.Second)
	channel <- id
}
