package main

import (
	"bufio"
	"fmt"
	"net"
	"os"
	"strings"
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

	for {
		fmt.Println("please say something:")
		input, _ := inputReader.ReadString('\n')
		input = strings.Trim(input, "\n")
		if input == "q" {
			fmt.Println("good bay")
			return
		}
		_, err = conn.Write([]byte(clientName + ": " + input))
	}
}
