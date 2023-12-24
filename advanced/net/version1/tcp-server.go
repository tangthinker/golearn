package main

import (
	"fmt"
	"net"
)

func main() {
	fmt.Println("start server...")
	listener, err := net.Listen("tcp", "localhost:50000")
	if err != nil {
		fmt.Println("error occurred in listening", err.Error())
		return
	}
	for {
		conn, err := listener.Accept()
		if err != nil {
			fmt.Println("error occurred in accept", err.Error())
			return
		}
		go doServer(conn)
	}
}

func doServer(conn net.Conn) {
	for {
		buff := make([]byte, 512)
		lens, err := conn.Read(buff)
		if err != nil {
			fmt.Println("error occurred in read connection", err.Error())
			return
		}
		fmt.Println("read from connection, data:", string(buff[:lens]))
	}
}
