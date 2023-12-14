package main

import (
	"fmt"
	"time"
)

/*
	go中channel
	如果是
	ch := make(chan string)
	则为同步阻塞通道 无缓冲
	如果是
	ch := make(chan string, buff) // buff := 10
	则为异步非阻塞通道 有缓冲
*/

func main() {
	var ch chan string
	ch = make(chan string)
	//var chFuc chan func()
	//chFuc = make(chan func())
	go sendData(ch)
	go receiveData(ch)
	time.Sleep(time.Second * 1)
}

func sendData(ch chan string) {
	ch <- "shanliao"
	ch <- "king"
}

func receiveData(ch chan string) {
	for receive := range ch {
		fmt.Println(receive)
	}
}
