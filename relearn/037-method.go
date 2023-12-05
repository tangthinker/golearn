package main

import "fmt"

/*
	在 Go 中，类型的代码和绑定在它上面的方法的代码可以不放置在一起，它们可以存在在不同的源文件
	唯一的要求是：它们必须是同一个包的。
	如果想要方法改变接收者的数据，就在接收者的指针类型上定义该方法

	在 Go 中，类型就是类（数据和关联的方法）。
	Go 不知道类似面向对象语言的类继承的概念。继承有两个好处：代码复用和多态。

	在 Go 中，代码复用通过组合和委托实现，
	多态通过接口的使用来实现：有时这也叫 组件编程（Component Programming）。
*/

type MyInt int64

func (m *MyInt) MinusOne() {
	*m--
}

func (m *MyInt) PlusOne() {
	*m++
}

type PersonE struct {
	id   uint64 // 未导出 相当于private
	Name string
	Age  int
}

func (p *PersonE) SetId(id uint64) {
	p.id = id
}

func (p *PersonE) GetId() uint64 {
	return p.id
}

type Animal struct {
	id   uint64
	Name string
	Age  int
}

type Cat struct {
	Color  string
	Animal // 会继承Animal的方法
	Master PersonE
}

func (p *Animal) SetId(id uint64) {
	p.id = id
}

func (p *Animal) GetId() uint64 {
	return p.id
}

func main() {
	i := MyInt(22)
	fmt.Println(i) //	22
	i.PlusOne()
	fmt.Println(i) //	23
	i.MinusOne()
	fmt.Println(i) // 22

	caty := Cat{
		Color: "WHITE",
		Master: PersonE{
			Name: "shanliao",
			Age:  21,
		},
	}
	caty.Master.id = 33
	caty.Animal.SetId(1)
	caty.Animal.Name = "caty"
	caty.Animal.Age = 3

	fmt.Println(caty) // {WHITE {1 caty 3} {33 shanliao 21}}

	caty.SetId(55)
	fmt.Println(caty) // {WHITE {55 caty 3} {33 shanliao 21}}

}
