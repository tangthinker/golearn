package main

import "fmt"

type Person struct {
	name string
	age  int32
}

func printPerson(person *Person) {
	fmt.Printf("name: %s, age: %d\n", person.name, person.age)
}

func printPersons(persons ...*Person) {
	if persons == nil || len(persons) == 0 {
		return
	}
	for _, person := range persons {
		printPerson(person)
	}
}

func main() {
	persons := []*Person{
		{
			"shanliao",
			22,
		},
		{
			"king",
			21,
		},
	}

	printPersons(persons...)
}
