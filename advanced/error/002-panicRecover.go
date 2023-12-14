package main

import (
	"fmt"
)

/*
	Go错误处理原则
	1）在包内部，总是应该从 panic 中 recover：不允许显式的超出包范围的 panic()
	2）向包的调用者返回错误值（而不是 panic）。
*/

func throwPanic() {
	panic("panic!")
}

func recoverFrom() (err error) {
	defer func() {
		if e := recover(); e != nil {
			fmt.Println("panic:", e)
			err = fmt.Errorf("find error %v", e)
		}
	}()
	throwPanic()
	fmt.Println("after call!")
	return
}

func main() {
	fmt.Println("call")
	err := recoverFrom()
	fmt.Println("end")

	fmt.Println(err)
}
