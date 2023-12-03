package main

/*
	像int float bool string complex error基本类型为值类型
	此外 数组 结构体 类型也为值类型
	可通过&获取值类型变量的内存地址

	在Go语言中，指针、slices、maps、channel、interface、func属于引用类型
	被引用的变量会存放在堆中，以便垃圾回收，且比栈有更大的内存空间

	Go语言中不存在隐式转换

	Go中 byte 是 uint8 的别名
*/

var number int = 32
var sli = make([]int, 3)

func modifyNumber(number int) { // 值传递
	number = 33
}

func modifySli(s []int) {
	s[1] = 3
}

func printSlice(s []int) {
	for i := range s {
		print(s[i])
	}
	println()
}

func main() {
	println("==============值类型==================")
	println(number)
	modifyNumber(number)
	println(number)
	println("==============引用类型==================")
	printSlice(sli)
	modifySli(sli)
	printSlice(sli)
}
