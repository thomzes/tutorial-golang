package main

import "fmt"

func main() {
	var name = "Thomas"
	var names = "Ardiansah"

	var result bool
	result = name == names

	fmt.Println(result)

	var value1 = 100
	var value2 = 200

	fmt.Println(value1 > value2)
	fmt.Println(value1 < value2)
	fmt.Println(value1 == value2)
	fmt.Println(value1 != value2)

}
