package main

import (
	"fmt"
	"net"
	"net/http"
)

func main() {
	http.HandleFunc("/", func(writer http.ResponseWriter, request *http.Request) {
		ctx := request.Context()

		server := ctx.Value(http.ServerContextKey).(*http.Server)
		fmt.Println("server addr:", server.Addr)

		local := ctx.Value(http.LocalAddrContextKey).(net.Addr)
		fmt.Println("local addr:", local)
	})
	// 注册处理器到DefaultServeMux中

	go http.ListenAndServe(":8080", nil) // 监听8080端口，并使用DefaultServeMux处理请求

	resp, _ := http.Get("http://localhost:8080")
	_ = resp.Body.Close()
}
