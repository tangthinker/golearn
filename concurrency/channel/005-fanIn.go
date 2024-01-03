package main

import "reflect"

func fanInReflect[T any](chs ...chan T) <-chan T {
	out := make(chan T)
	go func() {
		defer close(out)
		var cases []reflect.SelectCase
		for _, ch := range chs {
			cases = append(cases, reflect.SelectCase{
				Dir:  reflect.SelectRecv,
				Chan: reflect.ValueOf(ch),
			})
		}

		for len(cases) > 0 {
			chosen, recv, ok := reflect.Select(cases)
			if !ok { // 所选ch关闭，从切片中移除
				cases = append(cases[:chosen], cases[chosen+1:]...)
				continue
			}
			out <- recv.Interface().(T)
		}
	}()
	return out
}

func fanInRecursive[T any](chs ...chan T) <-chan T {
	switch len(chs) {
	case 0:
		ch := make(chan T)
		close(ch)
		return ch
	case 1:
		return chs[0]
	case 2:
		return mergeTwoCh(chs[0], chs[1])
	default:
		mid := len(chs) / 2
		return mergeTwoCh(
			fanInRecursive(chs[:mid]...),
			fanInRecursive(chs[mid:]...))
	}
}

func mergeTwoCh[T any](ch1, ch2 <-chan T) <-chan T {
	result := make(chan T)

	go func() {
		defer close(result)
		for ch1 != nil || ch2 != nil { // 只要其中一个没有关闭，则可以继续收集数据到result
			select {
			case recv, ok := <-ch1:
				if !ok {
					ch1 = nil
					continue
				}
				result <- recv
			case recv, ok := <-ch2:
				if !ok {
					ch2 = nil
					continue
				}
				result <- recv
			}
		}
	}()
	return result
}
