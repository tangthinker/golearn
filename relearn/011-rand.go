package main

import (
	"fmt"
	"math/rand"
	"time"
)

func main() {
	for i := 0; i < 10; i++ {
		n := rand.Int()
		fmt.Printf("%d \t", n)
	}
	fmt.Println()
	for i := 0; i < 10; i++ {
		n := rand.Intn(10)
		fmt.Printf("%d \t", n)
	}
	fmt.Println()

	// 以当前纳秒级时间作为随机数种子
	timeNS := int64(time.Now().Nanosecond())
	rand.Seed(timeNS)
	for i := 0; i < 10; i++ {
		n := rand.Intn(100)
		fmt.Printf("%d \t", n)
	}
}
