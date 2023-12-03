package main

import (
	"strconv"
	"strings"
)

func main() {
	var str = "hello world"
	println(strings.HasPrefix(str, "hello"))
	println(strings.HasSuffix(str, "world"))

	println(strings.Contains(str, "llo"))
	println(strings.Index(str, "world"))
	println(strings.Index(str, "red"))

	var zhStr = "你好 世界！"
	println(strings.IndexRune(zhStr, rune('世')))
	// rune 字符
	// 实际上是一个int32类型

	println(strings.Replace(str, "hello", "bye", -1))

	println(strings.ToUpper(str))

	var withSpace = "      hello Vue    "
	println(strings.TrimSpace(withSpace))
	printStringArray(strings.Fields(withSpace))

	var fileType = "exe,txt,py,go,vue,html"
	printStringArray(strings.Split(fileType, ","))
	fileTypeArray := strings.Split(fileType, ",")
	println(strings.Join(fileTypeArray, ","))

	var number = "123"
	num, err := strconv.Atoi(number)
	if err == nil {
		println(num)
	}
}

func printStringArray(strArray []string) {
	print("[")
	for i := range strArray {
		print(strArray[i], ",")
	}
	println("]")
}
