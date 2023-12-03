package main

import (
	"fmt"
	"math/rand"
	"time"
)

func main() {
	rand.Seed(int64(time.Now().Nanosecond()))
	n := rand.Intn(100)
	switch {
	case n < 25:
		fmt.Println("It's a number below 25!")
	case n >= 25 && n < 50:
		fmt.Println("It's a number great or equal 25 but below 50!")
	case n >= 50 && n < 75:
		fmt.Println("It's a number great or equal 50 but below 75!")
	case n >= 75 && n < 100:
		fmt.Println("It's a number great or equal 75 but below 100!")
	default:
		fmt.Println("It's occur a error in the program!")
	}
}
