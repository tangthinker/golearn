package main

import (
	"errors"
	"fmt"
)

func main() {
	err := errors.New("I do not want to see you any more!")
	fmt.Println(err)
	//error 本质为一个字符串
	//func New(text string) error {
	//	return &errorString{text}
	//}
	//
	//// errorString is a trivial implementation of error.
	//type errorString struct {
	//	s string
	//}
	panic("I love this world!")
}
