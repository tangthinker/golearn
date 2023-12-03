package main

import (
	"bytes"
	"fmt"
	"strconv"
)

func main() {
	var byteBuffer bytes.Buffer
	for i := 0; i < 10; i++ {
		byteBuffer.WriteString(strconv.Itoa(i))
	}
	fmt.Println(byteBuffer.String()) // 0123456789

}
