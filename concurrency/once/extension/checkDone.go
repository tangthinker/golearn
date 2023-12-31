package main

import (
	"fmt"
	"sync"
	"sync/atomic"
	"unsafe"
)

/*
	type Once struct {
		done uint32
		m    Mutex
	}

因为done是Once结构体的第一个字段，所以无须偏移，如果需要取得m，则需先转换为uintptr再进行偏移
*/
func OnceDone(once *sync.Once) bool {
	return atomic.LoadUint32((*uint32)(unsafe.Pointer(once))) == 1
}

func getOnceMutex(once *sync.Once) *sync.Mutex {
	return (*sync.Mutex)(unsafe.Pointer(uintptr(unsafe.Pointer(once)) + unsafe.Sizeof(uint32(0))))
}

func main() {
	var once sync.Once
	fmt.Println(OnceDone(&once))
	// fmt.Println(getOnceMutex(&once).TryLock()) // 会出现死锁 lock重入死锁

	once.Do(func() {
		fmt.Println("done")
	})

	fmt.Println(OnceDone(&once))

}
