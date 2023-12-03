package main

import (
	"fmt"
	"unicode/utf8"
)

func printStringSize(str string) {
	fmt.Println(len(str))
}

/*
	Go语言中的string类型 实际上是byte的定长数组
*/

func main() {
	var eng = "hello world!"
	var zh = "你好 世界！"
	fmt.Printf("%c\n", eng[0])
	fmt.Printf("%c\n", zh[0])
	printStringSize(eng)
	printStringSize(zh)
	var eng1 = "hello world!"
	println(eng1 == eng)
	eng1 += "!"
	println(eng1 == eng)
	println(eng1 >= eng)
	println(utf8.RuneCountInString(eng))
	println(len([]rune(eng)))
}
