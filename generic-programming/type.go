package main

import "fmt"

type MyType[T any] []T

type MyType2[T any] struct {
	data []T
	size int
}

func main() {

	intArr := &MyType[int]{1, 2, 3, 4, 5}

	fmt.Println(*intArr)

	strArr := &MyType2[string]{
		[]string{"hello", "world"},
		2,
	}

	fmt.Println(strArr)
	fmt.Println(strArr.data)
	fmt.Println(strArr.size)

}
