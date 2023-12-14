package main

import (
	"encoding/json"
	"fmt"
)

type BusinessError struct {
	Code    uint64
	Message string
}

func (err *BusinessError) Error() string {
	ret, _ := json.Marshal(err)
	return string(ret)
}

func throwError() error {
	return &BusinessError{
		10086,
		"login type error",
	}
}

func main() {
	err := throwError()
	if err != nil {
		fmt.Println(err.Error())
	}
}
