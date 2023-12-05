package main

import "fmt"

type Person struct {
	name string
	age  int32
}

/*
	1. 在go语言中struct.field中的.叫选择器（selector）。
	无论变量是一个结构体类型，还是结构体的指针都可以直接使用选择器访问其变量或方法

	2. go语言中结构体的数据在内存中是以连续块的形式存在的
		与Java中对象是引用类型相比，一个对象和它里面包含的对象可能在不同的内存空间，这点可用go中的指针做到
*/

func printPerson(person *Person) {
	fmt.Printf("name: %s, age: %d\n", person.name, person.age)
}

func printPersons(persons ...*Person) {
	if persons == nil || len(persons) == 0 {
		return
	}
	for _, person := range persons {
		printPerson(person)
	}
}

func main() {
	persons := []*Person{
		{
			"shanliao",
			22,
		},
		{
			"king",
			21,
		},
	}

	printPersons(persons...)
}
