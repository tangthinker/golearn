package main

import "fmt"

func main() {

	var price float32 = 0.3
	var tall float64 = 0.4 //浮点数默认float64

	fmt.Printf("price is %.2f , and tall is %.2f !\n", price, tall)

	price01 := 0.2
	price01 += 0.1

	price02 := 0.3

	//会出现浮点数精度问题
	fmt.Println("is 0.3 equal 0.3 ? ", price01 == price02)
	fmt.Printf("the first num is \t%.30f !\n", price01)
	fmt.Printf("the second num is \t%.30f !\n", price02)
	//the first num is  0.300000000000000044408920985006 !
	//the second num is 0.299999999999999988897769753748 !

}
