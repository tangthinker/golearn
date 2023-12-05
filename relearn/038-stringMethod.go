package main

import (
	"fmt"
	"strconv"
)

type PersonP struct {
	id   uint64
	name string
	age  int
}

func (p PersonP) String() string {
	return "[id:" + strconv.Itoa(int(p.id)) + ",name:" + p.name + ",age:" + strconv.Itoa(p.age) + "]"
}

func main() {
	p := &PersonP{
		22,
		"shanliao",
		22,
	}
	fmt.Println(p) // [id:22,name:shanliao,age:22]
}
