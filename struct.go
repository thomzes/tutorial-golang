package main

import "fmt"

type Customer struct {
	Name, Address string
	Age           int
}

func (customer Customer) sayHi(name string) {
	fmt.Println("Hello", name, "My Name is", customer.Name)
}

func (a Customer) sayHuuu() {
	fmt.Println("Huuuuuuu from", a.Name)
}

func main() {
	var thomas Customer
	thomas.Name = "Thomas"
	thomas.Address = "Indonesia"
	thomas.Age = 23

	thomas.sayHi("Adam")
	thomas.sayHuuu()

	// fmt.Println(thomas)

	// fmt.Println(thomas.Name)
	// fmt.Println(thomas.Address)
	// fmt.Println(thomas.Age)

	// // menggunakan struct bisa code dibawah
	// Tiyan := Customer{
	// 	Name:    "Tiyan",
	// 	Address: "Jakarta",
	// 	Age:     25,
	// }
	// fmt.Println(Tiyan)

	// // menggunakan stuct bisa dengan code dibawah, tetapi tidak disarankan karena field2 tersebut harus berurutan dan jelas
	// Manohara := Customer{"Manohara", "Bandung", 28}
	// fmt.Println(Manohara)

}
