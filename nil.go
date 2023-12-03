package main

import "fmt"

//nil意为零或者无 在go中，nil代表一个零值
//如果一个指针没有明确指向，则值为nil
//nil还是切片、映射和接口的零值

func main() {

	var zero *int
	fmt.Println(zero)
	//<nil>

	//无指向指针解引用引发panic
	//fmt.Println(*zero)
	//panic: runtime error: invalid memory address or nil pointer dereference
	//[signal 0xc0000005 code=0x0 addr=0x0 pc=0x70c555]

	//！解引用一个空指针将会引发程序崩溃

	//防范方法
	if zero != nil {
		fmt.Println(*zero)
	}

}
