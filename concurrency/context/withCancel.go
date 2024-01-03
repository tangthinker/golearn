package main

import (
	"context"
	"log"
	"time"
)

func consumptionTimeFunction(ctx context.Context, duration time.Duration, retCh chan string) {
	cur := 0
	for {
		select {
		case <-ctx.Done():
			log.Println("context is canceled")
			return
		default:
			cur++
			if cur == 10000000000 {
				retCh <- "It's done!"
				return
			}
		}
	}
}

func main() {
	ctx, cancel := context.WithCancel(context.Background())
	ch := make(chan string, 1)
	go consumptionTimeFunction(ctx, 4*time.Second, ch)

	time.Sleep(1 * time.Millisecond)
	select {
	case result := <-ch:
		log.Println(result)
		cancel()
	default:
		cancel()
		log.Println("program cancel!")
	}
	time.Sleep(time.Second)
}
