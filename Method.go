package main

import "fmt"

type username string

func (un username) print() string {
	fmt.Println(un)
	return string(un)
}

func main() {
	var user username = "shanliao"
	uname := user.print()
	fmt.Println(uname)
}
