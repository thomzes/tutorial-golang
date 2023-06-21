package main

import "fmt"

func main() {
	name := "Thomas"

	switch name {
	case "Thomas":
		fmt.Println("Hello Thomas")
	case "Ardiansah":
		fmt.Println("Hello Ardiansah")
	default:
		fmt.Println("Hi, Tidak Keduanya")
	}

	switch length := len(name); length > 5 {
	case true:
		fmt.Println("Nama terlalu panjang")
	case false:
		fmt.Println("Nama sudah benar")
	}

}
