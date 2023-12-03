package main

import "fmt"

func main() {

	//int 和 uint 会根据目标硬件选择合适的位长
	var num8 int8 = 32   //8位 1字节
	var num16 int16 = 32 //16位 2字节
	var num32 int32 = 32 //32位 4字节
	var num64 int64 = 32 //64位 8字节

	//默认的int类型取决于相关计算机架构

	fmt.Println(num8, num16, num32, num64)

	//fmt.Println(num8 == num32)	不同int类型的值无法进行比较

	//整数回绕
	var num int8 = 127
	num++
	fmt.Println("the num have plus 1 is ", num, "!")
	//the num have plus 1 is  -128!

}
