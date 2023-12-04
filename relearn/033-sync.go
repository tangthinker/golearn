package main

import (
	"bytes"
	"fmt"
	"sync"
)

type SyncedBuffer struct {
	lock   sync.Mutex
	buffer bytes.Buffer
}

func (receiver *SyncedBuffer) read(p []byte) (int, error) {
	receiver.lock.Lock()
	defer receiver.lock.Unlock()
	return receiver.buffer.Read(p)
}

func (receiver *SyncedBuffer) write(p []byte) (int, error) {
	receiver.lock.Lock()
	defer receiver.lock.Unlock()
	return receiver.buffer.Write(p)
}

func main() {
	syncedBuffer := new(SyncedBuffer)
	_, _ = syncedBuffer.write([]byte("hello world"))
	buff := make([]byte, 20)
	n, err := syncedBuffer.read(buff)
	if err == nil {
		fmt.Println(n)
		fmt.Println(string(buff))
	}
}
