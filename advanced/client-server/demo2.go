package main

import (
	"fmt"
	"math"
)

const MAX_REQUESTS = 100

var sameTime = make(chan int, MAX_REQUESTS)

type Request2 struct {
	arg       int
	replyChan chan int
}

func process(req *Request2) {
	req.replyChan <- int(math.Sqrt(float64(req.arg)))
}

/*
通过通道的方式限制服务器同时处理的请求数量
*/
func handle(req *Request2) {
	sameTime <- 1
	process(req)
	<-sameTime
}

func service(reqChan chan *Request2) {
	for {
		req := <-reqChan
		go handle(req)
	}
}

func startService() chan *Request2 {
	reqChan := make(chan *Request2)
	go service(reqChan)
	return reqChan
}

func main() {
	reqChan := startService()
	req1 := &Request2{
		arg:       4,
		replyChan: make(chan int),
	}
	reqChan <- req1
	fmt.Println(<-req1.replyChan)

}
