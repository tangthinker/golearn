package main

import (
	"fmt"
	"math/rand"
)

func main() {
	var count = 0
	var sum = 0
	for count < 100000000 {
		sum += rand.Intn(100) + 1
		count++
	}
	fmt.Printf("the average is %v !\n", sum/count)
	//the average is 50 !
}
