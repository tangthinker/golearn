package main

import (
	"context"
	"fmt"
	"net/http"
	"time"
)

func main() {
	req, err := http.NewRequest("GET", "https://google.com", nil)
	if err != nil {
		fmt.Println("err:", err)
		return
	}

	fmt.Println(req.Context()) // 默认用context.Background进行初始化request
	ctx, cancel := context.WithTimeout(req.Context(), 3*time.Second)
	defer cancel()

	req = req.WithContext(ctx)
	client := http.DefaultClient
	res, err := client.Do(req)
	if err != nil {
		fmt.Println("err:", err)
		return
	}
	fmt.Println("status:", res.StatusCode)
}
