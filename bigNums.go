package main

import (
	"fmt"
	"math/big"
)

func main() {

	//var num uint64 = 24000000000000000000
	//constant 24000000000000000000 overflows int

	//想要存储一个大于uint64范围的正整数 可以使用big包
	//储存大于uint64的数必须使用SetString的方式来进行 因为NewInt接受的参数就是uint64
	distance := new(big.Int)
	distance.SetString("24000000000000000000", 10) //第二个参数为进制

	fmt.Println("the distance is ", distance)

	//对于小于uint64的则可以NewInt的方式
	lightSpeed := big.NewInt(299792)
	fmt.Println("the light speed is ", lightSpeed)

	//尽管go编译器使用big包除了过大常量 但依然无法于big.int同等
	//const distance1 = 24000000000000000000
	//fmt.Println(distance1)
	//constant 24000000000000000000 overflows int

}
