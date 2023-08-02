package main

import (
	"fmt"
	"tutorial-golang/database"
	// pakai (_) blink identifier jika tidak ingin menggunakan package tetapi masih ada di list import
)


func main() {
	result := database.GetDatabase()
	fmt.Println(result)
}
