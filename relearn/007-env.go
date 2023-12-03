package main

import "os"

var (
	HOME      = os.Getenv("HOME")
	USER      = os.Getenv("USER")
	GOROOT    = os.Getenv("GOROOT")
	JAVA_HOME = os.Getenv("JAVA_HOME")
)

func main() {
	println(HOME)
	println(USER)
	println(GOROOT)
	println(JAVA_HOME)
}
