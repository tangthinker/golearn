package main

import (
	"fmt"
	"time"
)

type Request1 struct {
	arg1      int
	arg2      int
	replyChan chan int
}

type binaryOption func(arg1, arg2 int) int

func run(option binaryOption, req *Request1) {
	fmt.Println(option(req.arg1, req.arg2))
	req.replyChan <- option(req.arg1, req.arg2)
}

func server(option binaryOption, service chan *Request1, kill chan bool) {
	for {
		select {
		case req := <-service:
			go run(option, req)
		case k := <-kill:
			if k {
				fmt.Println("kill signal!")
				return
			}
		}

	}
}

func startServer(option binaryOption) (chan *Request1, chan bool) {
	requestChan := make(chan *Request1)
	killChan := make(chan bool)
	go server(option, requestChan, killChan)
	return requestChan, killChan
}

func main() {
	address, kill := startServer(func(arg1, arg2 int) int {
		return arg1 + arg2
	})
	const reqSize = 100
	requests := &[reqSize]Request1{}
	for i := 0; i < reqSize; i++ {
		request := &requests[i]
		request.arg1 = i + 1
		request.arg2 = i + 2
		request.replyChan = make(chan int)
		address <- request
	}

	for i := 0; i < reqSize; i++ {
		fmt.Println(<-requests[i].replyChan)
	}
	kill <- true
	time.Sleep(time.Nanosecond)
	fmt.Println("done")

}
