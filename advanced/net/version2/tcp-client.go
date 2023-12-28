package main

import (
	"bufio"
	"fmt"
	"net"
	"os"
	"strings"
)

const (
	NamePrefix1         = "name-"
	ContentPrefix1      = "content-"
	CommandPrefix1      = "command-"
	CommandOnline1      = "online"
	CommandLogout1      = "logout"
	CommandCloseServer1 = "close-server"
	CommandInputFlag    = "sh"
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

	_, _ = conn.Write([]byte(NamePrefix1 + clientName))

	buff := make([]byte, 512)
	lens, _ := conn.Read(buff)
	read := string(buff[:lens])
	if read == CommandLogout1 {
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
			_, _ = conn.Write([]byte(CommandPrefix1 + command))
			if command == CommandLogout1 || command == CommandCloseServer1 {
				return
			}
			lens, _ = conn.Read(buff)
			fmt.Println(string(buff[:lens]))
		} else {
			_, err = conn.Write([]byte(ContentPrefix1 + input))
		}
	}
}
