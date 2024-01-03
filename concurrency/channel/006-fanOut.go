package main

import (
	"reflect"
)

func fanOut[T any](in <-chan T, out []chan T, async bool) {
	go func() {
		defer func() {
			for _, ch := range out {
				close(ch)
			}
		}()
	}()

	for recv := range in {
		recv := recv
		for _, ch := range out {
			if async {
				go func() {
					ch <- recv
				}()
			} else {
				ch <- recv
			}
		}
	}
}

func fanOutReflect[T any](in <-chan T, out []chan T, async bool) {
	go func() {
		defer func() {
			for _, ch := range out {
				close(ch)
			}
		}()

		caseIn := reflect.SelectCase{
			Dir:  reflect.SelectRecv,
			Chan: reflect.ValueOf(in),
		}
		var cases []reflect.SelectCase
		cases = append(cases, caseIn)
		for {
			_, recv, ok := reflect.Select(cases)
			if !ok {
				return
			}
			for _, ch := range out {
				if async {
					go func() {
						ch <- recv.Interface().(T)
					}()
				} else {
					ch <- recv.Interface().(T)
				}
			}
		}
	}()
}
