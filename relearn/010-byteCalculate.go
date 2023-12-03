package main

import "fmt"

/*
	位运算只能用于整形变量 且必须为等长位
	可使用 %b 格式化输出

*/
func printByte(number uint8) {
	fmt.Printf("%08b\n", number)
}

func main() {
	// 按位与
	var a1 uint8 = 1
	var b1 uint8 = 3
	printByte(a1)
	printByte(b1)
	printByte(a1 & b1)

	// 按位或
	printByte(a1 | b1)

	// 异或 相同为0不同为1
	printByte(a1 ^ b1)

	// 位清除
	printByte(b1 &^ 3)

	// 按位取反
	printByte(^a1)

	// 位左移
	printByte(a1 << 3)

	// 位右移
	printByte(b1 >> 1)

}
