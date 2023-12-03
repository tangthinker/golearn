package main

//go支持将函数赋值给变量、闭包等操作

func print() {
	println("Hello World!")
}

func callFunction(f func()) {
	println("call function:")
	f()
}

//闭包
func swap1(a *int, b *int) func() {
	return func() {
		*a, *b = *b, *a
	}
}

func main() {
	p := print
	p()

	callFunction(print)

	var num1 = 1
	var num2 = 2
	println("the num1 is ", num1, "the num2 is ", num2)
	f := swap1(&num1, &num2)
	f()
	println("the num1 is ", num1, "the num2 is ", num2)
}
