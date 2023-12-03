package main

import (
	"fmt"
	"time"
)

func main() {

	fmt.Println(time.Now()) // 2023-12-01 15:24:34.5828245 +0800 CST m=+0.004903001
	now := time.Now()
	fmt.Println(now.Year(), now.Month(), now.Day())
	fmt.Printf("%d-%d-%d", now.Year(), now.Month(), now.Day())
	fmt.Println(now.Hour(), ":", now.Minute(), ":", now.Second(), ".", now.Nanosecond())
	fmt.Println(now.Location())
	fmt.Println(now.UTC())
	fmt.Println(now.Format("2006-01-02 15:04:05.000"))
	time.Sleep(time.Second * 5)

}
