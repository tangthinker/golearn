package main

import "fmt"

func ActionFuture(arg1 string, arg2 string) chan string {
	retChan := make(chan string)
	go func() {
		retChan <- arg1 + " say:" + arg2
	}()
	return retChan
}

func Action(arg1, arg2 string) string {
	retChan := ActionFuture(arg1, arg2)
	return <-retChan
}

func main() {
	fmt.Println(Action("shanliao", "Hello world!"))
}
