package main

import "fmt"

// 1. 程序不优雅
// 2. 引入了反射，性能会受影响
func add(a, b interface{}) interface{} {
	switch a.(type) {
	case int:
		switch b.(type) {
		case int:
			return a.(int) + b.(int)
		case float64:
			return float64(a.(int)) + b.(float64)
		default:
			return nil
		}
	case float64:
		switch b.(type) {
		case int:
			return a.(float64) + float64(b.(int))
		case float64:
			return a.(float64) + b.(float64)
		default:
			return nil
		}
	default:
		return nil
	}
}

func main() {

	a := 10
	b := 20.0
	c := add(a, b)
	fmt.Println(c)

}
