package main

import (
	"context"
	"fmt"
	"time"
)

func main() {
	parentCtx, parentCnl := context.WithTimeout(context.Background(), 5*time.Second)
	defer parentCnl()

	childCtx, childCnl := context.WithTimeout(parentCtx, 3*time.Second)
	defer childCnl()

	select {
	case <-childCtx.Done():
		fmt.Println("childCtx.Done() at:", time.Now())
		fmt.Println("childCtx.Err():", childCtx.Err())
	}

	select {
	case <-parentCtx.Done():
		fmt.Println("parentCtx.Done() at:", time.Now())
		fmt.Println("parentCtx.Err():", parentCtx.Err())
	}
}
