package main

import (
	"fmt"
	"os"
	"runtime"
)

func main() {
	goos := runtime.GOOS
	fmt.Println("Operating System: ", goos)
	path := os.Getenv("PATH")
	fmt.Println("Path: ", path)
}
