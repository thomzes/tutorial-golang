package main

import "fmt"

func getFullName() (string,string, string) {
	return "Thomas", "Ardiansah", "Tiyan"
}

func main() {
	// tambahkan "_" untuk mengabaikann parameter func
	firstName, _, lastName := getFullName() //tidak perduli dengan middlenane
	firstName, middleName, lastName := getFullName()
	fmt.Println(firstName, middleName, lastName)
}
