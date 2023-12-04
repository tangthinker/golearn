package main

import (
	"container/list"
	"container/ring"
	"fmt"
	"unsafe"
)

func main() {
	linkedList := list.New()
	linkedList.PushBack(1)
	linkedList.PushBack(2)
	linkedList.PushBack(3)
	linkedList.PushBack(4)
	linkedList.PushBack("shanliao")
	for e := linkedList.Front(); e != nil; e = e.Next() {
		fmt.Println(e.Value)
	}

	ringList := ring.New(10)
	for p := ringList; ; {
		p.Value = "hell world!"
		p = p.Next()
		if p == ringList {
			break
		}
	}
	ringList.Do(func(a2 any) {
		fmt.Println(a2)
	})
	var intp int
	p := unsafe.Sizeof(intp)
	fmt.Println(p)
}
