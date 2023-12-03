package main

import (
	"fmt"
	strconv "strconv"
)

func main() {

	num := 10
	fmt.Println("Hello World" + "!")
	fmt.Println("hello " + strconv.Itoa(num) + " years !")
	//使用strconv.Itoa（Int to ASCII）方法来将数字转为string
	//同样，可以使用strconv.Atoi(ASCII to Int)将string转为数字

	strnum := "20"
	var num1, err = strconv.Atoi(strnum)
	//Atoi方法有两返回值 第二个返回值若为<nil>则转换成功
	//<nil>为go中一个特殊常量
	fmt.Println("the num is ", num1, " and err is ", err)
	fmt.Printf("the err type is %T\n", err)
	if err == nil {
		fmt.Println("the nil is the special const variable!")
	}

}
