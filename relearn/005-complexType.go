package main

func main() {
	// 默认使用complex128类型
	var c = complex(1, 2)
	println(c)
	println("real -> ", real(c))
	println("imag -> ", imag(c))

	// 内置包中复数的操作狗使用complex128类型作为参数
	// 故最好使用complex128类型

}
