package main

import (
	"fmt"
	"runtime"
)

type PersonS struct {
	id   uint64
	Name string
	Age  int
}

func main() {
	p := PersonS{
		id:   22,
		Name: "shanliao",
		Age:  22,
	}
	runtime.SetFinalizer(&p, finalizer)

	fmt.Println("I am alive!")
	runtime.GC()
	fmt.Println("I am alive again!")
}

func finalizer(p *PersonS) {
	fmt.Println("object is removed!")
}
