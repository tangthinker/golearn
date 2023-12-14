package main

import (
	"fmt"
	"log"
	"net/http"
	_ "net/http/pprof"
	"time"
)

func main() {
	go func() {
		log.Println(http.ListenAndServe(":6060", nil))
	}()
	fmt.Println("main start!")
	time.Sleep(time.Second * 200)
}
