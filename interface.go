package main

import "fmt"

type HasName interface {
	GetName() string
}

func SayHello(hasName HasName) {
	fmt.Println("Hello", hasName.GetName())
}

type Person struct {
	Name string
}

func (person Person) GetName() string {
	return person.Name
}

type Animal struct {
	Name string
}

func (animal Animal) GetName() string {
	return animal.Name
}

func main() {
	var thomas Person
	thomas.Name = "Thomas"

	SayHello(thomas)

	kambing := Animal {
		Name: "Mbeeee",
	}

	SayHello(kambing)



}
