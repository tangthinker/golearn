package main

import (
	"fmt"
	"sort"
)

func main() {

	color := [...]string{
		"green",
		"white",
		"red",
		"yellow",
	}

	fmt.Println(color)

	greenAndWhite := color[0:2]
	fmt.Println(greenAndWhite)

	yellowAndRed := color[2:4]
	fmt.Println(yellowAndRed)

	white_ := color[1:]
	fmt.Println(white_)

	blank := color[:]
	have := color
	fmt.Println("blank ---> \t", blank)
	fmt.Println("have ---> \t", have)

	sort.Strings(blank)
	fmt.Println("sort ---> \t", blank)

	//使用make创建切片
	//指定长度为0 容量为10的切片
	colors := make([]string, 0, 10)
	colors = append(colors, "white", "red", "yellow", "green", "gray", "black")
	fmt.Println("use make to init slicing", colors)
}
