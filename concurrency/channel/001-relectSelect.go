package main

import (
	"context"
	"fmt"
	"math/rand"
	"reflect"
	"time"
)

func createSendCases(f func() int, chs ...chan int) []reflect.SelectCase {
	var result []reflect.SelectCase
	for _, ch := range chs {
		result = append(result, reflect.SelectCase{
			Dir:  reflect.SelectSend,
			Chan: reflect.ValueOf(ch),
			Send: reflect.ValueOf(f()),
		})
	}
	return result
}

func createRecvCases(chs ...chan int) []reflect.SelectCase {
	var result []reflect.SelectCase
	for _, ch := range chs {
		result = append(result, reflect.SelectCase{
			Dir:  reflect.SelectRecv,
			Chan: reflect.ValueOf(ch),
		})
	}
	return result
}

func main() {

	var ch1 = make(chan int, 10)
	var ch2 = make(chan int, 10)

	sendCases := createSendCases(func() int {
		return rand.Intn(10)
	}, ch1, ch2)

	recvCases := createRecvCases(ch1, ch2)

	var allCases []reflect.SelectCase
	allCases = append(allCases, sendCases...)
	allCases = append(allCases, recvCases...)

	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	defer cancel()

	for {
		select {
		case <-ctx.Done():
			fmt.Println("process return")
			return
		default:
			chosen, recv, ok := reflect.Select(allCases)
			if recv.IsValid() {
				fmt.Println("recv:", allCases[chosen].Dir, recv, ok)
			} else {
				fmt.Println("send:", allCases[chosen].Dir, ok)
			}
		}
	}

}
