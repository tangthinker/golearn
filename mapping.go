package main

import "fmt"

func main() {

	//不同于数组在赋值给新变量时将会复制一份
	//map映射类型在赋值给新变量时只是取了一个别名 共享同一份底层数据
	//修改其中一个，两个的数据内容都会改变

	ages := map[string]int{
		"shanliao":  22,
		"kingtang":  21,
		"snowflake": 23,
	}

	fmt.Println(ages)

	var sum int = 0
	for s, i := range ages {
		fmt.Println(s, " ---> ", i)
		sum += i
	}
	fmt.Println("the average of ages is ", sum/len(ages))

	alias := ages
	delete(alias, "snowflake")
	fmt.Println("delete alias : ", ages)

	//使用make创建映射
	set := make(map[string]int)
	set["first"] = 1
	set["second"] = 2
	set["third"] = 3
	set["forth"] = 4
	fmt.Println("use make to init map --> ", set)

}
