package main

import (
	"fmt"
)

func main() {
	var input int
	_, _ = fmt.Scanf("%d", &input)
	fmt.Println(input)

	var num int
	var str string
	_, _ = fmt.Scanln(&num, &str)
	fmt.Println(num, str)

	var num1 int
	var str1 string
	fmt.Sscanf("33 / helloworld", "%d / %s", &num1, &str1)
	fmt.Println(num1, str1)
}
