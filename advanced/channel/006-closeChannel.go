package main

import (
	"fmt"
	"time"
)

func send(ch chan string) {
	defer close(ch)
	ch <- "shanliao"
	ch <- "king"
}

func receive(ch chan string) {
	for {
		rec, ok := <-ch
		if !ok {
			break
		}
		fmt.Println(rec)
	}
	/*
		or
		for value := range ch {
			fmt.Println(value)
		}
	*/
}

func main() {
	ch := make(chan string)
	go send(ch)
	go receive(ch)
	time.Sleep(time.Second)
}
