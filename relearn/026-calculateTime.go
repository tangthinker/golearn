package main

import "time"

func calculateTime(t func()) time.Duration {
	start := time.Now()
	t()
	end := time.Now()
	duration := end.Sub(start)
	return duration
}

func f() {
	for i := 0; i < 10000000; i++ {
		a := 0
		a++
	}
}

func main() {
	println(calculateTime(f))
}
