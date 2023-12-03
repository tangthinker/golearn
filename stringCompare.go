package main

import "fmt"

func main() {
	var str = "shanliao"
	var str1 = "king"
	var str2 = "shanliao"
	fmt.Println(str == str1)
	fmt.Println(str == str2)

	//字符串可以直接比较 并且比较的是值而非Java中的地址

	fmt.Println(str >= str1)
	fmt.Println(str >= "zfdjsfldjsfjdsj")
	//字符串可以用大于小于 比较的可能是字符的大小 比如 z > a

	//可以通过%T的方式打印出变量的类型
	fmt.Printf("the type of str is %T !\n", str)

}
