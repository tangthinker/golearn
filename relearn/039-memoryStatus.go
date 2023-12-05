package main

import (
	"fmt"
	"runtime"
)

func main() {
	var memoryStatus runtime.MemStats
	runtime.ReadMemStats(&memoryStatus)
	fmt.Println("used ", memoryStatus.Alloc/1024, "kb")
}
