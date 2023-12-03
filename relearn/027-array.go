package main

import "fmt"

/*
	1. [5]int和[10]int是属于不同类型
	2. 数组的长度最大为2Gb
	3. 如果我们想让数组元素类型为任意类型的话可以使用空接口作为类型。当使用值时我们必须先做一个类型判断。
	4. 数组的长度是固定的，长度必须是一个常量表达式，不能为变量，如果长度不写就是切片
	5. 声明之后所有值将会被初始化为0值
	6. Go中数组为值类型（原生类型），不像C中为一个指针，函数传参的时候将会复制
*/

func main() {
	// n := 3
	// var arr [n]int 		// 报错
	var intArr [10]int
	fmt.Println(intArr) // [0 0 0 0 0 0 0 0 0 0]
	var strArr [10]string
	fmt.Println(strArr) // [         ]
	var complexArr [10]complex64
	fmt.Println(complexArr) // [(0+0i) (0+0i) (0+0i) (0+0i) (0+0i) (0+0i) (0+0i) (0+0i) (0+0i) (0+0i)]
	var pointerArr [10]*int
	fmt.Println(pointerArr) // [<nil> <nil> <nil> <nil> <nil> <nil> <nil> <nil> <nil> <nil>]

	modifyArr(intArr)
	fmt.Println(intArr) // [0 0 0 0 0 0 0 0 0 0]
	modifyArrByPointer(&intArr)
	fmt.Println(intArr) // [1 2 0 0 0 0 0 0 0 0]

	partInit := [5]string{2: "shanliao", 3: "king"} // [  shanliao king ]
	fmt.Println(partInit)
}

func modifyArr(arr [10]int) {
	arr[0] = 1
}

func modifyArrByPointer(arr *[10]int) {
	(*arr)[0] = 1
	arr[1] = 2 // 这样也是🆗的
}
