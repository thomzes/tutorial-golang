package main

import "fmt"

func main() {
	counter := 1

	for counter <= 10 {
		fmt.Println("Perulangan ke", counter)
		counter++
	}

	for counter := 1; counter <= 10; counter++ {
		fmt.Println("Ini perulangan ke", counter)
	}

	slice := []string {"Thomas", "Ardiansah"}

	for i := 0; i < len(slice); i++ {
		fmt.Println(slice[i]);
	}

	for i, value := range slice {
		fmt.Println("Index ke", i, "=", value)
	}


	// jika variable tidak di butuhkann maka tambahkan "_"
	for _, value := range slice {
		fmt.Println(value)
	}

	person := make(map[string]string)
	person["name"] = "Thomas"
	person["title"] = "Ardiansah"

	// kalau berupa map, harus menggunakan key
	for key, value := range person {
		fmt.Println(key, "=", value)
	}











}