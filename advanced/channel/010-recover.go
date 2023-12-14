package main

import (
	"log"
	"time"
)

func safelyDo() {
	defer func() {
		if err := recover(); err != nil {
			log.Println("an error occurred in do something")
		}
	}()
	do()
}

func do() {
	panic("error")
}

func main() {
	go safelyDo()
	time.Sleep(time.Second * 1)
}
