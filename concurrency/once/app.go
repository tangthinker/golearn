package main

import (
	"fmt"
	"net"
	"sync"
	"time"
)

/*
	使用once创建单例的常规写法
*/

var connectOnce struct {
	sync.Once
	conn net.Conn
}

func getConnect() net.Conn {
	connectOnce.Do(func() {
		connectOnce.conn, _ = net.DialTimeout("tcp", "baidu.com:80", 3*time.Second)
		fmt.Println("init completed!")
	})
	return connectOnce.conn
}

func main() {
	conn := getConnect()
	fmt.Println(conn)

	another := getConnect()
	fmt.Println(another)
}
