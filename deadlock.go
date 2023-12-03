package main

func main() {
	//deadlock 死锁 goroutine因为某些无法发生的事而永远阻塞
	c := make(chan int)
	<-c
}
