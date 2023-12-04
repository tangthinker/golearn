package main

import (
	"fmt"
	"math/big"
)

func main() {
	big1 := big.NewInt(1008686)
	big2 := big.NewInt(1001010)
	fmt.Println(big1.Mul(big1, big2).Mul(big1, big2)) // 1010724574680588600

	bigFloat1 := big.NewRat(1, 3)
	bigFloat2 := big.NewRat(1, 9)
	fmt.Println(bigFloat1.Mul(bigFloat1, bigFloat2)) // 1/27
}
