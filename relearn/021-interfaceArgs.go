package main

import "fmt"

func printInterface(arg interface{}) {
	switch arg.(type) {
	case int:
		value, _ := arg.(int)
		println(value)
	case float64:
		value, _ := arg.(float64)
		fmt.Println(value)
	case []string:
		slice, _ := arg.([]string)
		for _, value := range slice {
			print(value)
		}
	}
}

func main() {
	printInterface(22)
	printInterface(22.01)
	printInterface([]string{"hello", " ", "world", "!"})
}
