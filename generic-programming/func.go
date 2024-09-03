package main

import "fmt"

type Comparer interface {
	Compare(a interface{}) int
}

type MyType3[T Comparer] struct {
	data []T
}

func (m *MyType3[T]) Sort() {
	for i := 0; i < len(m.data); i++ {
		for j := i + 1; j < len(m.data); j++ {
			if m.data[i].Compare(m.data[j]) > 0 {
				m.data[i], m.data[j] = m.data[j], m.data[i]
			}
		}
	}
}

type MyInt struct {
	value int
}

func (m MyInt) Compare(a interface{}) int {
	if m.value > a.(MyInt).value {
		return 1
	} else if m.value < a.(MyInt).value {
		return -1
	} else {
		return 0
	}
}

func main() {
	m := MyType3[MyInt]{
		data: []MyInt{
			{value: 5},
			{value: 3},
			{value: 1},
		},
	}
	fmt.Println(m.data)
	m.Sort()
	fmt.Println(m.data)
}
