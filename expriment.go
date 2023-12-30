package main

const (
	mutexLocked = 1 << iota
	mutexWoken
	mutexWaiterShift = iota
)
