package main

import (
	"fmt"
	"net"
	"os"
	"slices"
	"strings"
)

const (
	NamePrefix         = "name-"
	ContentPrefix      = "content-"
	CommandPrefix      = "command-"
	CommandOnline      = "online"
	CommandLogout      = "logout"
	CommandCloseServer = "close-server"
)

var usernames []string

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
	buff := make([]byte, 512)
	lens, err := conn.Read(buff)
	name := string(buff[:lens])
	if name[:len(NamePrefix)] != NamePrefix {
		fmt.Println("you need input your name first!")
		return
	}
	name = name[len(NamePrefix):]
	if slices.Contains(usernames, name) {
		_, _ = conn.Write([]byte(CommandLogout))
		return
	} else {
		_, _ = conn.Write([]byte("welcome"))
	}
	usernames = append(usernames, name)
	for {
		lens, err = conn.Read(buff)
		if err != nil {
			fmt.Println("error occurred in read connection", err.Error())
			return
		}
		contentOrCommand := string(buff[:lens])
		if strings.Contains(contentOrCommand, CommandPrefix) {
			command := contentOrCommand[len(CommandPrefix):]
			switch command {
			case CommandOnline:
				_, _ = conn.Write([]byte(strings.Join(usernames, ",")))
			case CommandLogout:
				fmt.Println(name, "is logout!")
				usernames = DeleteSlice(usernames, name)
			case CommandCloseServer:
				if name == "shanliao" {
					fmt.Println("system is closing")
					os.Exit(0)
				} else {
					_, _ = conn.Write([]byte("access deny!"))
				}
			default:
				_, _ = conn.Write([]byte("undefined command!"))
				fmt.Println("undefined command!")
			}
		} else if strings.Contains(contentOrCommand, ContentPrefix) {
			content := contentOrCommand[len(ContentPrefix):]
			fmt.Println("read from "+name+", data:", content)
		} else {
			fmt.Println("message format error!")
			return
		}
	}
}

func DeleteSlice(s []string, elem string) []string {
	j := 0
	for _, v := range s {
		if v != elem {
			s[j] = v
			j++
		}
	}
	return s[:j]
}
