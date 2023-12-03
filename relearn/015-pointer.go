package main

func modify(n int) {
	n = 3
}

func modifyByPointer(n *int) {
	*n = 3
}

func modifyString(str string) {
	str = "hello world"
}

func modifyStringByPointer(str *string) {
	*str = "hello world"
}

func main() {
	n := 1
	println(n) //	1
	modify(n)
	println(n) //	1
	modifyByPointer(&n)
	println(n) // 3

	var pointer *int = &n
	println(pointer) // 0xc000055f68
	// 指针是一种特殊的数据类型，存储变量的内存地址
	// 因为指针是引用类型，所以默认零值为nil

	str := "hello go"
	println(str) //	hello go
	modifyString(str)
	println(str) //	hello go
	modifyStringByPointer(&str)
	println(str) // 	hello world

	var ppointer **int = &pointer
	println(ppointer)   // 地址为pointer指针本身的内存地址
	println(*ppointer)  // 地址为pointer指向的内存地址
	println(**ppointer) // 具体的值
}
