package main

var a = "G"

func n() {
	print(a)
}

func m() {
	// 这里是声明了一个函数内部的本地变量
	a := "O"
	print(a)
}

func main() {
	n()
	m()
	n()
}
