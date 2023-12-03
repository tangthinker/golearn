package main

import "strings"

/*
	使用函数可返回函数的特性编写一个工厂函数
	实现根据参数动态函数创建
*/
func MakeAndAddSuffix(suffix string) func(string) string {
	return func(s string) string {
		if !strings.HasSuffix(s, suffix) {
			return s + suffix
		}
		return s
	}
}

func main() {
	addTxt := MakeAndAddSuffix(".txt")
	addJpg := MakeAndAddSuffix(".jpg")
	println(addTxt("file"))
	println(addJpg("file"))
}
