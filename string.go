package main

import "fmt";

func main() {
	fmt.Println("Thomas");
	fmt.Println("Thomas Ardiansah");
	fmt.Println("Thomas Ardiansah Ardiansah");

	fmt.Println(len("Thomas"));
	fmt.Println("Thomas Ardiansah"[0]); //outputnya berupa byte
	fmt.Println("Thomas Ardiansah Ardiansah"[1]); //outputnya berupa byte
}