package main

import (
	"fmt"
	"reflect"
)

type Good struct {
	name  string "good name"
	count int    "good count"
}

func main() {
	good := Good{
		"hat",
		22,
	}
	goodType := reflect.TypeOf(good)
	fmt.Println(goodType.Size()) // 24
	for i := 0; i < 2; i++ {
		printTag(good, i)
	}
}

func printTag(good Good, i int) {
	goodType := reflect.TypeOf(good)
	field := goodType.Field(i)
	fmt.Println(field.Tag)
}
