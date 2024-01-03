package main

import (
	"context"
	"log"
	"time"
)

func main() {
	// 超时取消
	ctx, cancel := context.WithTimeout(context.Background(), 1*time.Second)
	<-ctx.Done()
	log.Println("err:", ctx.Err())
	cancel()

	// 主动取消
	ctx, cancel = context.WithTimeout(context.Background(), time.Second)
	cancel()
	<-ctx.Done()
	log.Println("err:", ctx.Err())

	// 父context取消
	pCtx, pCancel := context.WithCancel(context.Background())
	ctx, cancel = context.WithTimeout(pCtx, time.Second)
	pCancel()
	<-ctx.Done()
	log.Println("err:", pCtx.Err())
	cancel()
}
