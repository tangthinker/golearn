package main

import (
	"sync/atomic"
	"unsafe"
)

type LockFreeQueue[T any] struct {
	head unsafe.Pointer
	tail unsafe.Pointer
}

type node[T any] struct {
	value T
	next  unsafe.Pointer
}

func NewLockFreeQueue[T any]() *LockFreeQueue[T] {
	n := unsafe.Pointer(&node[T]{})
	return &LockFreeQueue[T]{
		head: n,
		tail: n,
	}
}

func loadNode[T any](p *unsafe.Pointer) (n *node[T]) {
	return (*node[T])(atomic.LoadPointer(p))
}

func casNode[T any](p *unsafe.Pointer, old *node[T], new *node[T]) (ok bool) {
	return atomic.CompareAndSwapPointer(p, unsafe.Pointer(old), unsafe.Pointer(new))
}

func (q *LockFreeQueue[T]) Enqueue(value T) {
	newNode := &node[T]{
		value: value,
	}
	for {
		tail := loadNode[T](&q.tail)
		next := loadNode[T](&tail.next)
		if tail == loadNode[T](&q.tail) {
			if next == nil {
				if casNode(&tail.next, next, newNode) {
					casNode(&q.tail, tail, newNode)
					return
				}
			} else {
				casNode(&q.tail, tail, next)
			}
		}
	}
}

func (q *LockFreeQueue[T]) Dequeue() T {
	var blank T
	for {
		head := loadNode[T](&q.head)
		tail := loadNode[T](&q.tail)
		next := loadNode[T](&head.next)
		if head == loadNode[T](&q.head) {
			if head == tail {
				if next == nil {
					return blank
				}
				casNode(&q.tail, tail, next)
			} else {
				result := next.value
				if casNode(&q.head, head, next) {
					return result
				}
			}
		}
	}
}
