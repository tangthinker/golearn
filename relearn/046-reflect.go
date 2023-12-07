package main

import (
	"fmt"
	"reflect"
)

type AnotherInt int

type PureStruct struct {
	name string
}

func (p PureStruct) PrintAndLog(msg string) {
	fmt.Println(msg, p.name)
}

/*
	type.Kind() 总是返回变量的底层类型
	结构中只有被导出的字段才是可以被设置的
*/

func main() {
	pi := 3.14
	t := reflect.TypeOf(pi)
	v := reflect.ValueOf(pi)
	fmt.Println(t) // float64
	fmt.Println(t.Kind())
	fmt.Println(v) // 3.14
	var num AnotherInt
	num = 32
	t = reflect.TypeOf(num)
	fmt.Println(t)        // main.AnotherInt
	fmt.Println(t.Kind()) // int
	var pureStruct PureStruct
	pureStruct.name = "shanliao"
	t = reflect.TypeOf(pureStruct)
	fmt.Println(t)        // main.PureStruct
	fmt.Println(t.Kind()) // struct

	// v.SetFloat(3.42)	// panic: reflect: reflect.Value.SetFloat using unaddressable value
	v1 := reflect.ValueOf(&pi)
	fmt.Println(v1)        // 0x1400000e098
	fmt.Println(v1.Elem()) // 3.14
	v1.Elem().SetFloat(3.1415926)
	fmt.Println(v1)
	fmt.Println(v1.Elem())          // 3.1415926
	fmt.Println(v1.CanSet())        // false
	fmt.Println(v1.Elem().CanSet()) // true
	fmt.Println(v.CanSet())         // false
	// fmt.Println(v.Elem().CanSet()) // panic
	// v.Elem()函数是解引用操作

	v = reflect.ValueOf(pureStruct)
	for i := 0; i < v.NumField(); i++ {
		fmt.Println(v.Field(i)) // shanliao
	}
	for i := 0; i < t.NumField(); i++ {
		fmt.Println(t.Field(i)) // {name main string  0 [0] false}
	}

	msg := "Hello"
	v.Method(0).Call([]reflect.Value{reflect.ValueOf(msg)}) // Hello shanliao
}
