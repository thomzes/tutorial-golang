package main

import "fmt"

func sayHelloTo(firstName string, lastName string) {
	fmt.Println("Hello my name is ",firstName, lastName)
}

func main() {
	firstName := "Thomas"
	sayHelloTo(firstName, "Ardiansah")
	sayHelloTo("Tiyanfernan", "Madafaka")
}