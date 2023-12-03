package main

import "fmt"

func main() {
	str := "你好 世界！"
	runes := []rune(str)
	for i := 0; i < len(runes); i++ {
		fmt.Printf("%c\n", runes[i])
	}

	for index, r := range str {
		fmt.Printf("index: %d, unicode: %U, charater: %c\n", index, r, r)
	}
}
