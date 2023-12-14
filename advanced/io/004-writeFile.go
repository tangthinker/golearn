package main

import (
	"bufio"
	"fmt"
	"os"
)

/*
	os.O_RDONLY：只读
	os.O_WRONLY：只写
	os.O_CREATE：创建：如果指定文件不存在，就创建该文件。
	os.O_TRUNC：截断：如果指定文件已存在，就将该文件的长度截为0。

	文件权限：
	- rwx rwx rwx
	0 111 111 111
	0  7  7  7
*/

func main() {
	outputFile, outputError := os.OpenFile("output.txt", os.O_WRONLY|os.O_CREATE, 0666)
	if outputError != nil {
		fmt.Println("an error is occurred in file open!")
	}
	defer func(outputFile *os.File) {
		err := outputFile.Close()
		if err != nil {
			fmt.Println("an error is occurred in file close!")
		}
	}(outputFile)
	fileWriter := bufio.NewWriter(outputFile)
	_, _ = fileWriter.WriteString("Hello world!")
	_ = fileWriter.Flush()

}
