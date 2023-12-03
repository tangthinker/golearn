package main

import "fmt"

func main() {
	a()
	b()
	c()
}

func a() {
	i := 0
	defer println(i) // 0 说明defer语句值传递
	i++
	return
}

func b() {
	i := 0
	p := &i
	defer println(*p) // 依然是0 说明defer会立刻拷贝当前函数传入的值
	// 调用 defer 关键字会立刻拷贝函数中引用的外部参数
	i++
	return
}

func c() {
	i := 0
	p := &i
	defer func() { fmt.Println(*p) }() // 1 此时拷贝的是函数指针
	i++
	return
}
