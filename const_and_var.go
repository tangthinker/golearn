package main

import "fmt"

func main() {
	const myName = "程晓枫"
	var myAge = 18 //在go中 变量被声明而未被使用将被视为错误
	fmt.Printf("My name is %v, and My age is %v!\n", myName, myAge)
	myAge += 3
	fmt.Printf("three years later!\n")
	fmt.Printf("My age is %v now!\n", myAge)

	//声明多个变量
	var (
		username = "kkkk"
		age      = 13
	)

	fmt.Printf("\nclient is %v, and he's age is %v !\n", username, age)
}
