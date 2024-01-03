package main

import (
	"fmt"
	"time"
)

type token int

func newWorker(id int, ch chan token, nextCh chan token) {
	for {
		t := <-ch
		fmt.Println(id)
		time.Sleep(100 * time.Millisecond)
		nextCh <- t
	}
}

func main() {
	chs := []chan token{
		make(chan token),
		make(chan token),
		make(chan token),
		make(chan token),
	}
	for i := 0; i < 4; i++ {
		go newWorker(i+1, chs[i], chs[(i+1)%4])
	}
	chs[0] <- token(1)

	select {}
}
