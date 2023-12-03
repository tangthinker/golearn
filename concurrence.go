package main

import (
	"fmt"
	"time"
)

func main() {
	//通过go关键字启动单个独立运行的任务goroutine
	//通过go启动的任务类似于启动一个线程
	go currency01()
	go currency02()
	time.Sleep(5 * time.Second)
}

func currency01() {
	fmt.Println("currency 01 is sleep!")
	time.Sleep(2 * time.Second)
	fmt.Println("01 wake up!")
}

func currency02() {
	fmt.Println("currency 02 is sleep!")
	time.Sleep(5 * time.Second)
	fmt.Println("02 wake up!")
}
