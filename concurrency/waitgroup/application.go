package main

import (
	"fmt"
	"io"
	"net/http"
	"sync"
	"time"
)

func main() {
	urls := []string{
		"https://baidu.com",
		"https://bing.com",
	}

	var wg sync.WaitGroup
	http.DefaultClient.Timeout = time.Second * 2
	result := make([]string, len(urls))
	for i := 0; i < len(urls); i++ {
		wg.Add(1)
		i := i
		go func(url string) {
			defer wg.Done()

			fmt.Println("fetching", url)
			resp, err := http.Get(url)
			if err != nil {
				fmt.Println("error occurred in access", url)
				fmt.Println(err)
				return
			}
			b, _ := io.ReadAll(resp.Body)
			result[i] = string(b[:512])
			_ = resp.Body.Close()
		}(urls[i])
	}

	wg.Wait()
	fmt.Println(result[0])
	fmt.Println(result[1])
}
