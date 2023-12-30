package main

import (
	"fmt"
	"sync"
	"sync/atomic"
	"time"
)

type TokenRecursiveMutex struct {
	sync.Mutex
	genToken  int64
	token     int64
	recursive int64
}

func (m *TokenRecursiveMutex) GenerateToken() int64 {
	return atomic.AddInt64(&m.genToken, 1)
}

func (m *TokenRecursiveMutex) Lock(token int64) {
	if atomic.LoadInt64(&m.token) == token {
		m.recursive++
		return
	}

	m.Mutex.Lock()
	atomic.StoreInt64(&m.token, token)
	atomic.StoreInt64(&m.recursive, 1)
}

func (m *TokenRecursiveMutex) Unlock(token int64) {
	if atomic.LoadInt64(&m.token) != token {
		panic(fmt.Sprintf("wrong the token(%d): %d", m.token, token))
	}

	recursive := atomic.AddInt64(&m.recursive, -1)
	if recursive != 0 {
		return
	}

	atomic.StoreInt64(&m.token, 0)
	m.Mutex.Unlock()
}

func main() {
	var tokenRecursiveMutex TokenRecursiveMutex
	go func() {
		goroutineToken := tokenRecursiveMutex.GenerateToken()
		tokenRecursiveMutex.Lock(goroutineToken)
		tokenRecursiveMutex.Lock(goroutineToken)
		tokenRecursiveMutex.Lock(goroutineToken)
		tokenRecursiveMutex.Lock(goroutineToken)
		fmt.Println(tokenRecursiveMutex.recursive)
		fmt.Println(tokenRecursiveMutex.token)
	}()
	time.Sleep(time.Second)
}
