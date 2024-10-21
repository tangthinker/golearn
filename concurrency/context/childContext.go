package main

import (
	"context"
	"log"
	"time"
)

func main() {
	parentContext := context.Background()

	childContext, cancel := context.WithTimeout(parentContext, 10*time.Second)

	err := doSomething(childContext)

	if err != nil {
		log.Println("error:", err)
	}

	defer cancel()

	select {
	case <-parentContext.Done():
		log.Println("parent context is canceled")
	default:
		log.Println("parent context is not canceled")
	}

	// 结论： child context 超时后，parent context 未被取消
}

func doSomething(ctx context.Context) error {
	select {
	case <-ctx.Done():
		return ctx.Err()
	}
}
