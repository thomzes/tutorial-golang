package main

import "fmt"

func main() {
	var name = "Thomas Ardiansah"

	if name == "Thomas" {
		fmt.Println("Hello Thomas")
	} else if name == "Jambang" {
		fmt.Println("Masuk ke Else If")
	} else {
		fmt.Println("Masuk ke else")
	}

	if length := len(name); length > 5 {
		fmt.Println("Terlalu panjang")
	} else {
		fmt.Println("Nama sudah besar")
	}

}
