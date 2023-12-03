package main

import "fmt"

//声明全局可复用结构
type location struct {
	x int
	y int
}

func main() {
	//声明临时局部结构
	var tempStruct struct {
		username string
		password string
	}
	tempStruct.username = "shanliao"
	tempStruct.password = "333"

	fmt.Println("the user info is ", tempStruct)

	var center location
	center.x = 500
	center.y = 600

	fmt.Println("the center is ", center)

	//通过符合字面量初始化结构
	left := location{
		x: 333,
		y: 222,
	}
	fmt.Println("the left is", left)

	right := location{
		222,
		111,
	}
	fmt.Println("the right is", right)

}
