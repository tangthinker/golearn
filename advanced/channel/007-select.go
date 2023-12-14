package main

import (
	"fmt"
	"time"
)

func generateOne(ch chan int) {
	for i := 1; ; i += 2 {
		ch <- i
	}
}

func generateTwo(ch chan int) {
	for i := 2; ; i += 2 {
		ch <- i
	}
}

func Suck(ch1, ch2 chan int) {
	for {
		select {
		case value := <-ch1:
			fmt.Println("receive ch1 value:", value)
		case value := <-ch2:
			fmt.Println("receive ch2 value:", value)
		}
	}
}

func main() {
	ch1 := make(chan int)
	ch2 := make(chan int)
	go generateOne(ch1)
	go generateTwo(ch2)
	go Suck(ch1, ch2) // end 116065
	time.Sleep(time.Second)
}
