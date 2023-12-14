package main

import (
	"fmt"
	"time"
)

func generateChannel() <-chan int {
	ch := make(chan int)
	go func() {
		for i := 0; ; i++ {
			ch <- i
		}
	}()
	return ch
}

func suck(ch <-chan int) {
	for value := range ch {
		fmt.Println(value)
	}
}

func main() {
	ch := generateChannel()
	go suck(ch) // end 133696
	time.Sleep(time.Second * 1)
}
