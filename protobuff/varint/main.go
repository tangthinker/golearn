package main

import (
	"fmt"
	"github.com/gogo/protobuf/proto"
)

func main() {
	var num int64

	num = -5

	if num >= 0 {
		num = num * 2
	}

	if num < 0 {
		num = -num
		num = num*2 - 1
	}

	fmt.Printf("num: %b\n", num)

	unum := uint64(num)

	fmt.Printf("unum: %b\n", unum)

	data := proto.EncodeVarint(uint64(num))

	for _, v := range data {
		fmt.Printf("%08b ", v)
	}

	fmt.Println()

	x, n := proto.DecodeVarint(data)

	fmt.Printf("x: %b, n: %d\n", x, n)

	if x%2 == 0 {
		source := x / 2
		fmt.Printf("source: %d\n", source)
	}

	if x%2 == 1 {
		source := -(int64(x) + 1) / 2
		fmt.Printf("source: %d\n", source)
	}
}
