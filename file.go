package main

import (
	"fmt"
	"io/ioutil"
	"os"
)

func main() {

	files, err := ioutil.ReadDir(".")
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
	for _, file := range files {
		fmt.Println(file.Name())
	}

	//write file
	fileWriter, errs := os.Create("shanliao.dat")
	if errs != nil {
		fmt.Println(errs)
		os.Exit(1)
	}
	_, err = fmt.Fprintln(fileWriter, "The world is beautiful! enjoy your life!")
	if err != nil {
		fileWriter.Close()
		fmt.Println(err)
		os.Exit(1)
	}

	//若写入文件的操作在函数内部，可使用defer关键字来保证函数返回前触发某些延迟操作
	//defer可以延迟某些操作
	//保证在return前触发这些延迟操作
	fileWriter.Close()

}
