package main

import (
	"bytes"
	"fmt"
)

func swap(a *int, b *int) {
	t := *b
	*b = *a
	*a = t
}

func encryption(str string) string {
	var result bytes.Buffer
	for i := range str {
		result.WriteByte(str[i] + 100)
	}
	return result.String()
}

func decryption(str string) string {
	var result bytes.Buffer
	for i := range str {
		result.WriteByte(str[i] - 100)
	}
	return result.String()
}

func main() {
	a := 1
	b := -1
	println("a is ", a, " and b is ", b)

	swap(&a, &b)

	println("swaped a is ", a, " and b is ", b)

	encrypt := encryption("Hello World!")
	fmt.Println("encryption: ", encrypt)
	fmt.Println("decryption: ", decryption(encrypt))
}
