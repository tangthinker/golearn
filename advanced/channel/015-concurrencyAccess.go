package main

import (
	"fmt"
	"strconv"
	"time"
)

type PersonC struct {
	Name                string
	Money               int
	MoneyAccessFunction chan func()
}

func NewPersonC(name string, money int) *PersonC {
	person := &PersonC{name, money, make(chan func())}
	go person.backend()
	return person
}

func (person *PersonC) backend() {
	for action := range person.MoneyAccessFunction {
		go action()
	}
}

func (person *PersonC) SetMoney(money int) {
	person.MoneyAccessFunction <- func() {
		person.Money = money
	}
}

func (person *PersonC) GetMoney() int {
	getChan := make(chan int)
	person.MoneyAccessFunction <- func() {
		getChan <- person.Money
	}
	return <-getChan
}

func (person *PersonC) String() string {
	return "Person: name:" + person.Name + ", money:" + strconv.Itoa(person.GetMoney())
}

func main() {
	person := NewPersonC("shanliao", 100)
	fmt.Println(person)
	person.SetMoney(200)
	fmt.Println(person)
	for i := 0; i < 100; i++ {
		time.Sleep(time.Nanosecond * 50)
		go func() {
			person.SetMoney(person.GetMoney() + 1)
			fmt.Println(person)
		}()
	}
	time.Sleep(time.Nanosecond)
	fmt.Println(person) // money 一定不是300
	// 因为for循环中的协程基本上是同一时间启动的获取到的money值是一样的
	// 猜测：可能和go运行时中协程调度的机制有关
}
