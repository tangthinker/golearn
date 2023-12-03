package main

import "fmt"

//定义接口
type animalAction interface {
	action() string
}

type animal struct {
	name string
}

//实现接口
func (ani animal) action() string {
	fmt.Println(ani.name, "animal action!")
	return ani.name + " animal action!"
}

//只要实现了接口中的方法的结构都可以作为参数
func printAnimal(action animalAction) string {
	return action.action()
}

func main() {
	ani := animal{"cat"}
	returnString := printAnimal(ani)
	fmt.Println("return string is", returnString)
}
