package main

import (
	"crypto/md5"
	"fmt"
)

func main() {
	hash := md5.New()
	shanliao := "shanliao"
	hash.Write([]byte(shanliao))
	fmt.Println("result :", string(hash.Sum(nil)))
	fmt.Printf("result: %x\n", hash.Sum([]byte{}))
}
