package main

import (
	"fmt"
	"reflect"
)

/*
	varargs 可变参数
	只能放在函数最后一个参数
	以slice类型进行接收
*/

func printArgs(arg ...string) {
	if arg == nil || len(arg) == 0 {
		return
	}
	t := reflect.TypeOf(arg)
	fmt.Println(t) // []string
	for _, value := range arg {
		fmt.Print(value)
	}
}

func main() {
	printArgs("hello", " ", "world")
	printArgs()
	strSlice := []string{"hello", " ", "world"}
	printArgs(strSlice...) // 如果参数被存储在一个 slice 类型的变量 slice 中，则可以通过 slice... 的形式来传递参数
}
