package main

import (
	"bytes"
	"encoding/gob"
	"fmt"
)

type PersonG struct {
	id   uint64
	Name string
	Age  int
}

func main() {
	shanliao := PersonG{
		001,
		"shanliao",
		22,
	}
	var buff bytes.Buffer
	encoder := gob.NewEncoder(&buff)
	_ = encoder.Encode(shanliao)
	fmt.Println(buff.Bytes())

	fmt.Println(buff.String())
	// %PersonG��Name
	//               Age�shanliao,
}
