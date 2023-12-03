package main

import (
	"fmt"
	"time"
)

func main() {

	//使用make函数创建channel
	channel := make(chan string)
	go phone01(channel)
	go phone02(channel)
	time.Sleep(time.Second * 1)

}

func phone01(channel chan string) {
	//阻塞等待有另一个goroutine尝试对相同通道进行接受操作为止
	channel <- "hello world!"
	fmt.Println("send successful!")
}

func phone02(channel chan string) {
	//阻塞等待有另一个goroutine尝试对相同通道进行发送操作为止
	receive := <-channel
	fmt.Println("receive : ", receive)
}
