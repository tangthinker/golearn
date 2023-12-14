package main

import (
	"fmt"
	"time"
)

type Item interface{}

type Container []Item

func (c *Container) Len() int {
	return len([]Item(*c))
}

func (c *Container) IndexOf(i int) Item {
	return []Item(*c)[i]
}

func (c *Container) Add(item Item) {
	*c = append(*c, item)
}

func (c *Container) Iter() <-chan Item {
	ch := make(chan Item)
	go func() {
		for i := 0; i < c.Len(); i++ {
			ch <- c.IndexOf(i)
		}
	}()
	return ch
}

func main() {
	var container Container
	container.Add("shanliao")
	container.Add("king")
	go func() {
		for value := range container.Iter() {
			switch value.(type) {
			case string:
				fmt.Println(value.(string))
			}
		}
	}()
	time.Sleep(time.Second)
}
