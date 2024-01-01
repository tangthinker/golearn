package main

import (
	"fmt"
	"sync/atomic"
	"time"
)

type Config struct {
	NodeName   string
	NodeMemory int64
	Addr       string
}

func initConfig() Config {
	return Config{
		NodeName:   "shanliao",
		NodeMemory: 2048,
		Addr:       "192.168.0.101",
	}
}

func main() {
	var config atomic.Value
	config.Store(initConfig())

	// 修改
	go func() {
		for {
			time.Sleep(2 * time.Second)
			old := config.Load().(Config)
			old.NodeMemory = 4096
			config.Store(old)
		}
	}()

	// 拉取
	go func() {
		old := config.Load().(Config)
		fmt.Println("old:", old)
		time.Sleep(3 * time.Second)
		updated := config.Load().(Config)
		fmt.Println("new:", updated)
	}()

	time.Sleep(5 * time.Second)
}
