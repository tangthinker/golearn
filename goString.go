package main

import "fmt"

func main() {

	//原始字符串字面量
	fmt.Println(`the world is beautiful!
hello world!`)

	fmt.Printf("%v is a %[1]T\n", "literal string")
	fmt.Printf("%v is a %[1]T\n", "raw string literal")
	//无论原始字符串字面量还是字符串字面量都最终将变成字符串

	//go语言提供rune（符文）类型来表示单个UTF编码的代码点
	//虽然rune类型代表的是单个字符，但实际储存的是数字值
	//rune类型与char类型十分相似
	var pi rune = 960
	var omega rune = 969 //rune本质是一种int32 byte本质为uint8
	fmt.Printf("%c\n%c\n", pi, omega)

	//go中字符串无法被修改
	//cannot assign to message[3] (strings are immutable)
	message := "shanliao"
	//message[3] = 'n'
	fmt.Println(message)

}
