package main

import "fmt"

/*

	1. 切片是一种引用类型，是对内部数组空间的一种索引域
	2. 切片数据结构包含3个数据：数组指针ptr，长度len，容量cap
	3. 切片创建：
		1. slice1 := make([]type, len, cap)
		2. slice1 := make([]type, len)
		3. slice1 := arr[:]
	4. 在使用for-range时，（index，value）中value只是位置上值的拷贝，不能用于修改切片或数组中的值
	5. 如果想增加切片的容量，我们必须创建一个新的更大的切片并把原分片的内容都拷贝过来。 append、copy
*/

func main() {
	var s1 []int
	println(s1)                   // [0/0]0x0
	fmt.Println(len(s1), cap(s1)) // 0 0
	s1 = append(s1, 0, 0, 0, 0, 0, 0)
	println(s1)                   // [6/6]0xc00000a450 append操作之后地址改变
	fmt.Println(len(s1), cap(s1)) // 5 6
	s2 := s1[1:]
	fmt.Println(s1) // [0 0 0 0 0 0]
	s2[0] = 3
	fmt.Println(s1) // [0 3 0 0 0 0]
	modifySlice(s1)
	fmt.Println(s1) // [1 3 0 0 0 0]
}

/*
	new(T) 为每个新的类型T分配一片内存，初始化为 0 并且返回类型为*T的内存地址：这种方法 返回一个指向类型为 T，值为 0 的地址的指针，
			它适用于值类型如数组和结构体；它相当于 &T{}。
	make(T) 返回一个类型为 T 的初始值，
			它只适用于3种内建的引用类型：切片、map 和 channel（参见第 8 章，第 13 章）
*/

func modifySlice(s []int) {
	s[0] = 1
}

/*
	append方法类似于这样
	内部判断数组容量是否足够
	不够创建一个双倍的容量，并拷贝
*/
func AppendByte(slice []byte, data ...byte) []byte {
	m := len(slice)
	n := m + len(data)
	if n > cap(slice) { // if necessary, reallocate
		// allocate double what's needed, for future growth.
		newSlice := make([]byte, (n+1)*2)
		copy(newSlice, slice)
		slice = newSlice
	}
	slice = slice[0:n]
	copy(slice[m:n], data)
	return slice
}
