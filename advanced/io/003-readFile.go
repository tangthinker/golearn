package main

import (
	"bufio"
	"fmt"
	"os"
)

// Unix和Linux的行结束符是 \n，而Windows的行结束符是 \r\n

func main() {
	inputFile, inputError := os.Open("data.txt")
	if inputError != nil {
		fmt.Println("an error is occurred in file open!")
		return
	}
	defer func(inputFile *os.File) {
		err := inputFile.Close()
		if err != nil {
			fmt.Println("an error is occurred in file close!")
		}
	}(inputFile)

	//inputReader := bufio.NewReader(inputFile)
	//for {
	//	readString, readError := inputReader.ReadString('\n')
	//	if readError == io.EOF {
	//		break
	//	}
	//	fmt.Print(readString)
	//}

	inputReader := bufio.NewReader(inputFile)
	buff := make([]byte, 1024)
	for {
		n, err := inputReader.Read(buff)
		if n == 0 || err != nil {
			break
		}
		fmt.Println(string(buff))
	}

}
