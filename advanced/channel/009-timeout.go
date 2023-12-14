package main

import (
	"fmt"
	"time"
)

/*
	即使是timeout的场景
	也无法在timeout的时候立即杀死协程
	注意一点：必须使用带缓冲的channel以避免已启动的协程由于阻塞迟迟无法结束

	goroutine 被设计为不可以从外部无条件地结束掉，只能通过 channel 来与它通信。
	也就是说，每一个 goroutine 都需要承担自己退出的责任。
	(A goroutine cannot be programmatically killed. It can only commit a cooperative suicide.)

*/

func timeoutCall(ch chan string) {
	defer func() {
		fmt.Println("I am done!")
	}()
	time.Sleep(time.Second * 5)
	ch <- "Hello world!"
}

func main() {
	ch := make(chan string, 1)
	go timeoutCall(ch)
	select {
	case res := <-ch:
		fmt.Println("response:", res)
	case <-time.After(time.Second * 2):
		fmt.Println("timeout exit")
		break
	}
	time.Sleep(time.Second * 4)
	fmt.Println("end")
}
