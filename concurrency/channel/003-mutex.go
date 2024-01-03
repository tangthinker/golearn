package main

import (
	"fmt"
	"time"
)

type ChannelMutex struct {
	ch chan struct{}
}

func NewChannelMutex() *ChannelMutex {
	m := &ChannelMutex{
		make(chan struct{}, 1),
	}
	m.ch <- struct{}{}
	return m
}

func (m *ChannelMutex) Lock() {
	<-m.ch
}

func (m *ChannelMutex) Unlock() {
	select {
	case m.ch <- struct{}{}:
	default:
		panic("unlock a unlocked mutex")
	}
}

func (m *ChannelMutex) TryLock() bool {
	select {
	case <-m.ch:
		return true
	default:
	}
	return false
}

func (m *ChannelMutex) TimeoutLock(timeout time.Duration) bool {
	timer := time.NewTimer(timeout)
	select {
	case <-m.ch:
		timer.Stop()
		return true
	case <-timer.C:
	}
	return false
}

func (m *ChannelMutex) isLocked() bool {
	return len(m.ch) == 0
}

func main() {
	m := NewChannelMutex()
	ok := m.TryLock()
	fmt.Println("isLocked:", ok)
	ok = m.TryLock()
	fmt.Println("isLocked:", ok)
	m.Unlock()
	ok = m.TryLock()
	fmt.Println("isLocked:", ok)

}
