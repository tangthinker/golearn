package main

func main() {
	// 使用复合自变量初始化数组
	var array = [5]int{
		1, 3, 5, 7, 9,
	}

	for i := range array {
		println(array[i])
	}

	// 无需在意数组元素数量
	array1 := [...]int{
		2, 4, 6, 8, 10,
	}

	for i := range array1 {
		println(array1[i])
	}

	strs := [...]string{
		"shanliao",
		"kingtang",
		"snowflake",
	}

	for i, str := range strs {
		println(i, " --> ", str)
	}

}
