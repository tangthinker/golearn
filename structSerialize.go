package main

import (
	"encoding/json"
	"fmt"
)

func main() {

	//需要序列化的字段必须以大写字母开头
	type location struct {
		X int
		Y int
	}

	var center location
	center = location{
		222,
		222,
	}

	fmt.Println("the center is", center)

	//Marshal函数只会对结构中以大写字母开头的导出字段进行编码
	//且编码的结构时byte类型
	resultByte, err := json.Marshal(center)

	fmt.Println("the result is ", resultByte, " the err is ", err)
	fmt.Println("the string is ", string(resultByte))

	//使用结构标签改变输出字段名
	type userinfo struct {
		Username string `json:"name"`
		Password string `json:"password"`
	}

	shanliao := userinfo{
		"shanliao", "333",
	}

	resultByte, err = json.Marshal(shanliao)
	fmt.Println("the result is ", resultByte, "the err is ", err)
	fmt.Println("the string is ", string(resultByte))
}
