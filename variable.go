package main

import "fmt";

func main() {
	var name string; //buat variable declare data

	name = "Thomas Ardiansah";
	fmt.Println(name);

	name = "Thomas Ardiansyah";
	fmt.Println(name);

	var friendName = "DoomDoom"; //buat variable tanpa declare tipe data
	fmt.Println(friendName);

	var age = 30;
	fmt.Println(age);

	country := "Indonesia"; //menggunakan := untuk declare awal variable tanpa menggunakan tipe datanya, untuk selanjutnya tidak bisa menggunakan :=
	fmt.Println(country);

	// contoh selanjutnya
	country = "America"; //ini bisa, tapi kalau pakai := tidak bisa

	// membuat lebih dari satu variable sehingga memudahkan
	var (
		firstName = "Thomas";
		lastName = "Ardiansah";
	)
	fmt.Println(firstName,lastName);



}