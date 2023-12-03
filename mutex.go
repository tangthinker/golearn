package main

import (
	"fmt"
	"sync"
	"time"
)

func main() {

	mutex := sync.Mutex{}

	strs := map[string]int{
		"shanliao": 333,
	}

	fmt.Println(strs)
	go safeModify(strs, mutex)
	time.Sleep(1)
	fmt.Println(strs)

}

func safeModify(strs map[string]int, mutex sync.Mutex) {
	mutex.Lock()
	//确保函数返回时解锁
	defer mutex.Unlock()

	strs["kingtang"] = 333

}
