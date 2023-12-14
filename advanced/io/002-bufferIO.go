package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	inputReader := bufio.NewReader(os.Stdin)
	fmt.Println("please enter some input:")
	in, err := inputReader.ReadString('\n') // 会包含换行符
	if err == nil {
		fmt.Println("you input:", in[0:len(in)-1])
	}

	outputWriter := bufio.NewWriter(os.Stdout)
	_, _ = outputWriter.WriteString("Hello world!\n")
	_ = outputWriter.Flush()

}
