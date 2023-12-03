package main

import (
	"log"
	"runtime"
)

var where = func() {
	_, file, line, _ := runtime.Caller(1)
	// 打印在哪个文件的哪一行
	log.Printf("%s:%d", file, line)
}

func main() {
	where()
	println("hello world!")
	where()
}
