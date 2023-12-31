package main

import (
	"errors"
	"fmt"
	"sync"
	"sync/atomic"
)

type Once struct {
	done  uint32
	mutex sync.Mutex
}

func (once *Once) Do(f func() error) error {
	if atomic.LoadUint32(&once.done) == 1 {
		return nil
	}
	return once.doSlow(f)
}

func (once *Once) doSlow(f func() error) error {
	once.mutex.Lock()
	defer once.mutex.Unlock()

	var err error
	if once.done == 0 {
		err = f()
		if err == nil {
			atomic.StoreUint32(&once.done, 1)
		}
	}
	return err
}

type AutoRetryOnce struct {
	once          Once
	maxRetryCount int32
}

func (retryOnce *AutoRetryOnce) Do(f func() error) bool {
	currentCount := int32(0)
	var err error
	for {
		err = retryOnce.once.Do(f)
		// fmt.Println("try")
		if err == nil || currentCount >= retryOnce.maxRetryCount {
			break
		}
		currentCount++
	}
	return err == nil
}

func main() {
	var autoRetryOnce AutoRetryOnce
	autoRetryOnce.maxRetryCount = 3
	count := 0
	fmt.Println(autoRetryOnce.Do(func() error {
		if count <= 3 {
			count++
			return errors.New("error")
		}
		return nil
	}))

}
