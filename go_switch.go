package main

import "fmt"

func main() {
	var selectc = "333"

	switch selectc {
	case "111":
		fmt.Println("Are you OK !")
	case "222":
		fmt.Println("Hello Mifans !")
	case "333":
		fmt.Println("How are you ?")
		fallthrough //在go中 switch会自动break，想要下降下一分支必须使用fallthrough关键字
	case "444":
		fmt.Println("Are you OK ?")
	default:
		fmt.Println("Thank you indian Mifans !")
	}

}
