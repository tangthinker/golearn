package main

import "fmt"

func main() {
	array := []uint64{
		33333,
		2222,
		10086,
	}
	ch := make(chan uint64)
	go sum(array, ch)
	fmt.Println(<-ch)
}

func sum(array []uint64, ch chan uint64) {
	var ret uint64
	for _, e := range array {
		ret += e
	}
	ch <- ret
}
