package main

import (
	"fmt"
	"runtime"
	"strconv"
	"strings"
	"sync"
	"sync/atomic"
	"time"
)

/*
	重入锁的关键是：如何给线程一个唯一标识以便于同一线程无需再锁
	实现可重入锁的两种方案：
	1. 获取goroutine id（特殊方法），记录其id，实现Locker接口
	2. 调用lock/unlock方法时，验证goroutine提供的token

	方案一获取goroutine id的方法有两种：
	（1）分析当前堆栈信息，解析其ID
	（2）获取运行时g结构，得倒ID

*/

/*
分析堆栈信息以获取goroutine id
还可以使用kortschak/goroutine库获取
*/
func getGoroutineId() int64 {
	var buff [64]byte
	len := runtime.Stack(buff[:], false)
	idField := strings.Fields(strings.TrimPrefix(string(buff[:len]), "goroutine "))[0]
	id, err := strconv.Atoi(idField)
	if err != nil {
		panic(fmt.Sprintf("cannot get goroutine id: %v", err))
	}
	return int64(id)
}

type RecursiveMutex struct {
	sync.Mutex
	owner     int64
	recursion int64
}

func (m *RecursiveMutex) Lock() {
	gid := getGoroutineId()
	// 如果以获取到锁的goroutine再次获取锁
	if atomic.LoadInt64(&m.owner) == gid {
		atomic.AddInt64(&m.recursion, 1)
		return
	}

	// 其他goroutine尝试获取锁 阻塞
	m.Mutex.Lock()
	atomic.StoreInt64(&m.owner, gid)
	atomic.StoreInt64(&m.recursion, 1)
}

func (m *RecursiveMutex) Unlock() {
	gid := getGoroutineId()
	// 非持有锁的goroutine无法释放锁
	if atomic.LoadInt64(&m.owner) != gid {
		panic(fmt.Sprintf("wrong the owner(%d): %d", m.owner, gid))
	}

	// 重入次数减一
	recursive := atomic.AddInt64(&m.recursion, -1)
	if recursive != 0 {
		return
	}
	// 如果重入次数为0 则表示goroutine完成所有释放操作，释放锁
	atomic.StoreInt64(&m.owner, -1)
	m.Mutex.Unlock()
}

func main() {
	var recursiveMutex RecursiveMutex
	go func() {
		recursiveMutex.Lock()
		recursiveMutex.Lock()
		recursiveMutex.Lock()
		fmt.Println(recursiveMutex.recursion)
	}()
	time.Sleep(time.Second)
}
