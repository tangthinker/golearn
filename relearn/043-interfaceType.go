package main

import "fmt"

type Crier interface {
	Cry()
}

type Baby struct {
	name string
}

type Girl struct {
	Baby
}

type Boy struct {
	Baby
}

func (g Girl) Cry() {
	fmt.Println("Girl", g.name, "is crying!")
}

func (b Boy) Cry() {
	fmt.Println("Boy", b.name, "is crying!")
}

func main() {
	var crier Crier
	crier = &Girl{ // 如果要修改其值，必须传入指针类型的实现
		Baby{"Alice"},
	}
	judgeSex(crier)
	fmt.Println(crier)

	crier = Boy{
		Baby{"Bob"},
	}
	judgeSex(crier)
	fmt.Println(crier)
	// type-switch
	switch realTypeValue := crier.(type) {
	case *Girl:
		fmt.Printf("Type is %T with value %v \n", realTypeValue, realTypeValue)
	case Boy:
		fmt.Printf("Type is %T with value %v \n", realTypeValue, realTypeValue)
	}

	// interface type
	boy := Boy{Baby{"Bob"}}
	if interfaceTypeValue, ok := interface{}(boy).(Crier); ok {
		fmt.Println("It implement interface Crier")
		interfaceTypeValue.Cry()
	}

}

func judgeSex(crier Crier) {
	// real type
	if realTypeValue, ok := crier.(*Girl); ok {
		fmt.Println("It's a girl crying! Her name is", realTypeValue.name)
		realTypeValue.name = "shanliao"
	}
	if realTypeValue, ok := crier.(Boy); ok {
		fmt.Println("It's a boy crying! His name is", realTypeValue.name)
		realTypeValue.name = "shanliao"
	}
}
