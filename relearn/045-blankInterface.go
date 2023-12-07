package main

import (
	"fmt"
)

/*
	空接口
	type Any interface{}
	any或Any式空接口很好的别名或缩写
	空接口类似Java/C#中所有类的基类 Object
	可以给空接口变量赋任意类型的值

	空接口变量在内存中占据两个字长：
	1. 所包含的类型
	2. 存储包含的数据（如果赋值为值类型）或指向数据的指针（如果赋值为引用类型）
*/

type Any interface{}

func main() {
	var val Any
	val = 5
	fmt.Println(val)
	val = 3.14
	fmt.Println(val)
	val = "shanliao"
	fmt.Println(val)

	switch val.(type) {
	case int:
		fmt.Println("is int")
	case string:
		fmt.Println("is string")
	}
}
