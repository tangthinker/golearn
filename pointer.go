package main

import "fmt"

func plusOne(num *int) {
	*num++
}

func modifyMap(info map[string]int) {
	for s := range info {
		if s == "snowflake" {
			delete(info, s)
		}
	}
}

func main() {

	name := "shanliao"
	address := &name
	fmt.Println("the name", name, " in memory", address)
	//the name shanliao  in memory 0xc00005a230

	fmt.Printf("the address is the Type of %T", address)
	//the address is the Type of *string

	age := 21
	age1 := &age
	*age1 = 22
	fmt.Println("the age is", age, " now!")

	temp := 37
	plusOne(&temp)
	fmt.Println("plus one :", temp)

	//映射实际上是一种隐式指针
	userinfo := map[string]int{
		"shanliao":  333,
		"snowflake": 111,
	}
	fmt.Println("userinfo:", userinfo)
	modifyMap(userinfo)
	fmt.Println("modified userinfo:", userinfo)

}
