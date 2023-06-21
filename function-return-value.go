package main

import "fmt"

func getHello(name string) string {
	if name == "" {
		return "Hallo guys!"
	} else {
		return "Hello " + name
	}
}

func main() {
	fullName := getHello("Thomas")
	fmt.Println(fullName)

	fullNames := getHello("")
	fmt.Println(fullNames)
}