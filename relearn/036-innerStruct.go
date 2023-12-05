package main

import "fmt"

/*
	结构体中内嵌未命名的结构体类似于面向对象语言中继承
	更像组合的方式
	不过会合并命名空间
	外层同名变量会覆盖内层同名变量的值,如果同一层出现同名变量则会编译错误:
		1. 外层名字会覆盖内层名字（但是两者的内存空间都保留），这提供了一种重载字段或方法的方式；
		2. 如果相同的名字在同一级别出现了两次，如果这个名字被程序使用了，将会引发一个错误（不使用没关系）。
			没有办法来解决这种问题引起的二义性，必须由程序员自己修正。
*/

type Inner struct {
	num1 int
	num2 int
	num3 int
}

type Outer struct {
	num3 int
	num4 int
	int
	Inner
}

func main() {
	outer := Outer{
		num3: 1,
		num4: 2,
	}
	outer.Inner.num3 = 4
	outer.int = 3
	outer.num1 = 4
	outer.num2 = 5
	fmt.Println(outer)      // {1 2 3 {4 5 4}}
	fmt.Println(outer.num3) // 1

}
