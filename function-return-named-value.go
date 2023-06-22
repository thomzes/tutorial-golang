package main

import "fmt"

func getFullName2() (firstName, middleName, lastName string) {
	firstName = "Thomas"
	middleName = "Ardiansah"
	lastName = "Terakhir"

	return
}

func main() {
	a, b, c := getFullName2()
	fmt.Println(a)
	fmt.Println(b)
	fmt.Println(c)
}
