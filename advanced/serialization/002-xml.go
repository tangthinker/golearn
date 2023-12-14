package main

import (
	"encoding/xml"
	"fmt"
	"strings"
)

type PersonX struct {
	id   uint64
	Name string
	Age  int
}

func main() {
	shanliao := PersonX{
		001,
		"shanliao",
		22,
	}
	ret, err := xml.Marshal(shanliao)
	if err == nil {
		fmt.Println(string(ret))
	}
	reader := strings.NewReader(string(ret))
	decoder := xml.NewDecoder(reader)
	for token, err := decoder.Token(); err == nil; token, err = decoder.Token() {
		switch t := token.(type) {
		case xml.StartElement:
			name := t.Name.Local
			fmt.Println("Token start name:", name)
			for _, attr := range t.Attr {
				fmt.Println("attr name:", attr.Name, " attr value:", attr.Value)
			}
		case xml.EndElement:
			fmt.Println("Token end, name:", t.Name)
		case xml.CharData:
			fmt.Println("Token content:", string(t))
		}
	}
}
