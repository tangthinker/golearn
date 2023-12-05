package main

import "fmt"

/*
	1. go中的接口十分灵活，通过接口可以实现很多面向对象的特性
	2. 接口定义了一组方法集，但不包含其实现
	3. 接口的名字由方法名加上er后缀组成，例如Printer、Reader、Logger等
	4. Go中的接口都很简短，通常只包含0～3个方法

	5. 不同于面向对象中的接口，go中接口可以创建相关变量
	   该变量的值为nil，实际上是一个指针
	   包含receiver和method table ptr两部分
			具体对象   对象实现的抽象方法地址表

	6. go中类型不需要显示声明其实现了某个接口，接口被隐式实现
*/

type Eater interface {
	Eat()
}

type Dog struct {
	name string
}

type Chicken struct {
	color string
}

func (d Dog) Eat() {
	fmt.Println("Dog - ", d.name, "is eating!")
}

func (c Chicken) Eat() {
	fmt.Println("Chicken with", c.color, "is eating!")
}

func main() {
	var eater Eater
	eater = Dog{"little yellow"} // Dog -  little yellow is eating!
	eater.Eat()
	eater = Chicken{"red"} // Chicken with red is eating!
	eater.Eat()
}
