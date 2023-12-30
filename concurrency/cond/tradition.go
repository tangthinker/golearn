package main

import (
	"fmt"
	"math/rand"
	"sync"
	"time"
)

func main() {
	c := sync.NewCond(&sync.Mutex{})

	var condition int

	for i := 0; i < 10; i++ {
		go func(i int) {
			time.Sleep(time.Duration(rand.Intn(10)) * time.Second)

			c.L.Lock()
			defer c.L.Unlock()
			condition++
			fmt.Println("goroutine", i, "is ready!")
			c.Broadcast()
		}(i)
	}

	var wakeCount int

	// 每次wait唤醒之后，条件并不是一定成立，需要再次判断
	// 以下为惯例写法
	// 注意：调用cond.wait()方法之前当前goroutine必须持有cond.L锁
	c.L.Lock()
	for condition != 10 {
		c.Wait()
		wakeCount++
		fmt.Println("judge was woken", wakeCount, "times!")
	}
	c.L.Unlock()

	fmt.Println("finally wakeCount is", wakeCount)
	fmt.Println("preparation is done!")
}

/*
	扩展：

	type Cond struct {
		noCopy noCopy

		// L is held while observing or changing the condition
		L Locker

		notify  notifyList
		checker copyChecker
	}

	Cond实现中的copyChecker用于运行时检查Cond是否被复制，否则panic
	// copyChecker holds back pointer to itself to detect object copying.
	type copyChecker uintptr

	func (c *copyChecker) check() {
		if uintptr(*c) != uintptr(unsafe.Pointer(c)) &&
			!atomic.CompareAndSwapUintptr((*uintptr)(c), 0, uintptr(unsafe.Pointer(c))) &&
			uintptr(*c) != uintptr(unsafe.Pointer(c)) {
			panic("sync.Cond is copied")
		}
	}

	copyChecker用于保存其本身的内存地址
	如果Cond被复制，那么copyChecker的值与内存地址不相等

	check方法可被展开成以下代码：
	func (c *copyChecker) check() {
		if uintptr(*c) != uintptr(unsafe.Pointer(c)) {
			// 第一次进入 c为0，赋值为其地址值
			if atomic.CompareAndSwapUintptr((*uintptr)(c), 0, uintptr(unsafe.Pointer(c))) {
				return
			} else {
				if uintptr(*c) != uintptr(unsafe.Pointer(c)) {
					panic("sync.Cond is copied")
				}
			}
		}
	}
	或者：
	func (c *copyChecker) check() {
		if uintptr(*c) != uintptr(unsafe.Pointer(c)) {
			// 第一次进入 c为0，赋值为其地址值
			if atomic.CompareAndSwapUintptr((*uintptr)(c), 0, uintptr(unsafe.Pointer(c))) {
				return
			}
			panic("sync.Cond is copied")
		}
	}

*/
