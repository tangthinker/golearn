package main

import (
	"fmt"
	"net/http"
	"sync"
	"time"
)

func main() {
	var pool sync.Pool
	pool.New = func() any {
		return &http.Client{
			Timeout: 5 * time.Second,
		}
	}
	//for i := 0; i < 5; i++ { 手动初始化5个成功5个失败5个
	//	pool.Put(&http.Client{
	//		Timeout: 5 * time.Second,
	//	})
	//}
	//runtime.GC()
	//runtime.GC() // 两次gc后全部失败

	var wg sync.WaitGroup
	wg.Add(10)
	go func() {
		for i := 0; i < 10; i++ {
			go func() {
				defer wg.Done()
				client, ok := pool.Get().(*http.Client)
				if !ok {
					fmt.Println("get client failed")
					return
				}
				defer pool.Put(client)

				resp, err := client.Get("https://baidu.com")
				if err != nil {
					fmt.Printf("get error")
					return
				}
				_ = resp.Body.Close()
				fmt.Println("get successful")
			}()
		}
	}()
	wg.Wait()
}
