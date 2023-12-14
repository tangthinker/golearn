package main

import (
	"fmt"
	"time"
)

/*
	timer和ticker的区别
	timer只执行一次
	ticker会不断执行

	在底层上timer没有设置period字段，而ticker设置了：
	t := &Ticker{
		C: c,
		r: runtimeTimer{
			when:   when(d),
			period: int64(d),
			f:      sendTime,
			arg:    c,
		},
	}
	t := &Timer{
		C: c,
		r: runtimeTimer{
			when: when(d),
			f:    sendTime,
			arg:  c,
		},
	}
*/

func main() {
	ticker := time.NewTicker(time.Second * 1)
	defer ticker.Stop()
	go func() {
		for t := range ticker.C {
			fmt.Println("ticker", t)
		}
	}()
	timer := time.NewTimer(time.Second * 1)
	defer timer.Stop()
	go func() {
		for t := range timer.C {
			fmt.Println("timer", t)
		}
	}()
	time.Sleep(10 * time.Second)

}
