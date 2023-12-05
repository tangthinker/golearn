package main

import "bytes"

type ReadWrite interface {
	Read(buff bytes.Buffer) bool
	Write(buff bytes.Buffer) bool
}

type Lock interface {
	Lock()
	Unlock()
}

type File interface {
	ReadWrite
	Lock // 接口嵌套
	Close()
}
