package main

import "fmt"

func main() {
	var names [3]string
	names[0] = "Thomas"
	names[1] = "Ardi"
	names[2] = "Ansah"

	fmt.Println(names[0])
	fmt.Println(names[1])
	fmt.Println(names[2])

	var values = [3]int{
		90,
		80,
		70,
	}
	fmt.Println(values)

	fmt.Println(len(names))
	fmt.Println(len(values))

	var lagi [10]string

	fmt.Println(len(lagi))

}
