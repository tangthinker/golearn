package main

import (
	"fmt"
	"net/http"
)

func main() {
	http.HandleFunc("/", HelloServer)
	http.HandleFunc("/echo/", EchoServer)
	http.HandleFunc("/index/", IndexServer)
	// 必须在listen支持注册处理器
	err := http.ListenAndServe("localhost:8080", nil)
	if err != nil {
		panic(err)
	}
}

func HelloServer(writer http.ResponseWriter, request *http.Request) {
	fmt.Println("into hello server")
	_, _ = writer.Write([]byte("Hello world!"))
}

func EchoServer(writer http.ResponseWriter, request *http.Request) {
	fmt.Println("into echo server")
	_, _ = writer.Write([]byte(request.URL.Path[6:]))
}

func IndexServer(writer http.ResponseWriter, request *http.Request) {
	fmt.Println("into index server")
	title := "Welcome to shanliao site!"
	body := "Hello Beautiful World!"
	_, _ = fmt.Fprintf(writer, "<h1>%s</h1>\n<div>%s</div>", title, body)
}
