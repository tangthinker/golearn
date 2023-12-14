package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"path/filepath"
)

type Page struct {
	Title string
	Body  []string
}

func (page *Page) save(filePath string) {
	output, err := os.OpenFile(filePath+string(filepath.Separator)+page.Title,
		os.O_CREATE|os.O_WRONLY, 0666)
	if err != nil {
		errorPrint("file open")
	}
	defer func(output *os.File) {
		err := output.Close()
		if err != nil {
			errorPrint("file close")
		}
	}(output)
	writer := bufio.NewWriter(output)
	for _, str := range page.Body {
		_, _ = writer.WriteString(str + "\n")
	}
	_ = writer.Flush()
}

func (page *Page) load(filePathWithName string) {
	page.Title = filepath.Base(filePathWithName)
	input, err := os.Open(filePathWithName)
	if err != nil {
		errorPrint("file open")
	}
	defer func(input *os.File) {
		err := input.Close()
		if err != nil {
			errorPrint("file close")
		}
	}(input)
	reader := bufio.NewReader(input)
	for {
		str, err := reader.ReadString('\n')
		if err == io.EOF {
			break
		}
		page.Body = append(page.Body, str)
	}
}

func errorPrint(msg string) {
	fmt.Println("an error occurred in", msg)
}

func main() {
	page := Page{
		Title: "shanliao.txt",
		Body: []string{
			"Hello world!",
			"我是山鸟",
			"你好世界！",
		},
	}
	page.save("./file/")
	load := new(Page)
	load.load("./file/shanliao.txt")
	fmt.Println(load)
}
