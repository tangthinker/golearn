package main

import (
	"net/http"
	_ "net/http/pprof"
	"sync"
)

func main() {
	var counter int64
	var m sync.Mutex

	for i := 0; i < 100; i++ {
		go func() {
			m.Lock()
			//defer m.Unlock() // 故意不释放锁
			// http://localhost:8080/debug/pprof/goroutine?debug=1
			// 会发现有99个goroutine阻塞在15行

			// goroutine profile: total 103
			//99 @ 0x104e6d718 0x104e7f748 0x104e7f725 0x104e9c2f8 0x104ea9034 0x10506517c 0x105065121 0x104ea08f4
			//#	0x104e9c2f7	sync.runtime_SemacquireMutex+0x27	/usr/local/go/src/runtime/sema.go:77
			//#	0x104ea9033	sync.(*Mutex).lockSlow+0x173		/usr/local/go/src/sync/mutex.go:171
			//#	0x10506517b	sync.(*Mutex).Lock+0x7b			/usr/local/go/src/sync/mutex.go:90
			//#	0x105065120	main.main.func1+0x20			/concurrency/mutex/pprof.go:15

			counter++
		}()
	}
	err := http.ListenAndServe("localhost:8080", nil)
	if err != nil {
		panic(err)
	}
}
