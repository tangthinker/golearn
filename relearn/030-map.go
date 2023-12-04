package main

import (
	"fmt"
	"math/rand"
	"time"
)

func main() {
	var index map[int]string         // 仅仅是声明
	index = make(map[int]string, 10) // make在堆中创建数据结构
	index[0] = "shanliao"
	index[1] = "king"
	for k, v := range index {
		fmt.Printf("id: %d name: %s\n", k, v)
	}

	pointer := new(map[int]string) // 创建一个*map[int]string的指针
	pointer = &index
	println((*pointer)[0])

	funcMap := map[int]func(){
		0: func() {
			println(0)
		},
		1: func() {
			println(1)
		},
		2: func() {
			println(2)
		},
		3: func() {
			println(3)
		},
	}
	r := rand.New(rand.NewSource(int64(time.Now().Nanosecond())))
	for i := 0; i < 4; i++ {
		funcMap[r.Intn(4)]()
	}

	value, exits := index[1]
	if exits {
		println(value)
	}
	delete(index, 1)
	fmt.Println(index)

	smap := make([]map[int]string, 5)
	fmt.Println(smap)
	for i := range smap {
		smap[i] = make(map[int]string) // 必须为切片内部各个map进行初始化
		smap[i][0] = "shanliao"
	}
	fmt.Println(smap)

}
