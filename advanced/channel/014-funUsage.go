package main

import (
	"flag"
	"fmt"
)

var nGoroutine = flag.Int("size", 100, "size of goroutines")

func function(left chan int, right chan int) {
	left <- (1 + <-right)
}

/*
	中间启动了size个goroutine
	由于right没有值一直阻塞
	当right传入0时所有协程依次执行
	最终得倒最左边的值
*/

func main() {
	flag.Parse()
	lefted := make(chan int)
	var left, right chan int = nil, lefted
	for i := 0; i < *nGoroutine; i++ {
		left, right = right, make(chan int)
		go function(left, right)
	}
	right <- 0
	ret := <-lefted
	fmt.Println(ret)
}
