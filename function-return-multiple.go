package main

import "fmt"

func getFullName() (string,string, string) {
	return "Thomas", "Ardiansah", "Tiyan"
}

func main() {
	// tambahkan "_" untuk mengabaikann parameter func
	firstName, middleName, lastName := getFullName()
	firstName, _, lastName := getFullName() //tidak perduli dengan middlenane
	fmt.Println(firstName, middleName, lastName)
}
