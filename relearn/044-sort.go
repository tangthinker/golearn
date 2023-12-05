package main

import (
	"fmt"
	"sort"
)

type PersonK struct {
	id   uint64
	Name string
	Age  int
}

type PersonKArray []PersonK

func (p *PersonKArray) Len() int {
	return len([]PersonK(*p))
}

func (p *PersonKArray) Less(i, j int) bool {
	v := []PersonK(*p)
	return v[i].id < v[j].id
}

func (p *PersonKArray) Swap(i, j int) {
	v := []PersonK(*p)
	v[i].id, v[j].id = v[j].id, v[i].id
	v[i].Name, v[j].Name = v[j].Name, v[i].Name
	v[i].Age, v[j].Age = v[j].Age, v[i].Age
}

func main() {
	personList := &PersonKArray{
		PersonK{
			id:   9,
			Name: "shanliao",
			Age:  22,
		},
		PersonK{
			id:   3,
			Name: "king",
			Age:  33,
		},
		PersonK{
			id:   7,
			Name: "Bob",
			Age:  44,
		},
	}
	fmt.Println(personList)
	sort.Sort(personList)
	fmt.Println(personList)
	fmt.Println(sort.IsSorted(personList))
}
