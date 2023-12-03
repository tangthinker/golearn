package main

import "fmt"

type phone struct {
	color  string
	number string
}

func main() {
	//new(T)会为类型为T的新项分配已置零的内存空间，并返回它的地址。也就是一个类型为*T的值
	p := new(phone)
	p.number = "10086"
	p.color = "white"
	fmt.Println("my phone number is", p.number, " and it is", p.color)
}
