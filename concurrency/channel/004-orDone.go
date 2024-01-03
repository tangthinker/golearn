package main

import "reflect"

func orDone1(chs ...<-chan any) <-chan any {
	switch len(chs) {
	case 0:
		return nil
	case 1:
		return chs[0]
	}
	orDone := make(chan any)
	go func() {
		defer close(orDone)

		switch len(chs) {
		case 2:
			select {
			case <-chs[0]:
			case <-chs[1]:
			}
		default: // 递归调用
			mid := len(chs) / 2
			select {
			case <-orDone1(chs[:mid]...):
			case <-orDone1(chs[mid:]...):
			}

		}
	}()
	return orDone
}

func orDone2(chs ...<-chan any) <-chan any {
	switch len(chs) {
	case 0:
		return nil
	case 1:
		return chs[0]
	}
	orDone := make(chan any)
	go func() {
		defer close(orDone)

		var cases []reflect.SelectCase
		for _, ch := range chs {
			cases = append(cases, reflect.SelectCase{
				Dir:  reflect.SelectRecv,
				Chan: reflect.ValueOf(ch),
			})
		}
		// 只要其中一个有响应就结束整个orDone
		reflect.Select(cases)
	}()
	return orDone
}
