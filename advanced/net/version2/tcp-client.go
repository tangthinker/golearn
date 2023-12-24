package main

import (
	"bufio"
	"fmt"
	"net"
	"os"
	"strings"
)

const (
	NamePrefix         = "name-"
	ContentPrefix      = "content-"
	CommandPrefix      = "command-"
	CommandOnline      = "online"
	CommandLogout      = "logout"
	CommandCloseServer = "close-server"
	CommandInputFlag   = "sh"
)

func main() {
	conn, err := net.Dial("tcp", "localhost:50000")
	if err != nil {
		fmt.Println("error occurred in dial", err.Error())
		return
	}

	inputReader := bufio.NewReader(os.Stdin)
	fmt.Println("what is your name?")
	clientName, _ := inputReader.ReadString('\n')
	clientName = strings.TrimSpace(strings.Trim(clientName, "\n"))

	_, _ = conn.Write([]byte(NamePrefix + clientName))

	buff := make([]byte, 512)
	lens, _ := conn.Read(buff)
	read := string(buff[:lens])
	if read == CommandLogout {
		fmt.Println(clientName, "is already login!")
		return
	}
	for {
		fmt.Println("please say something:")
		input, _ := inputReader.ReadString('\n')
		input = strings.Trim(input, "\n")
		if input == "q" {
			fmt.Println("good bay")
			return
		}
		if strings.Contains(input, CommandInputFlag) {
			command := strings.TrimSpace(input[len(CommandInputFlag):])
			_, _ = conn.Write([]byte(CommandPrefix + command))
			if command == CommandLogout || command == CommandCloseServer {
				return
			}
			lens, _ = conn.Read(buff)
			fmt.Println(string(buff[:lens]))
		} else {
			_, err = conn.Write([]byte(ContentPrefix + input))
		}
	}
}
