package main

import "fmt"

// panic membuat code tidak dijalankan ketika panic ter eksekusi
// recover menampung data panic dan code akan tetap berjalan
// recover di sarankan di taruh di defer function


func endApp() {
	message := recover()
	if message != nil {
		fmt.Println("Error dengan message:", message)
	}
	fmt.Println("aplikasi selesai")
}

func runApp(error bool) {
	defer endApp()
	if error {
		panic("APLIKASI ERROR")
	}
	fmt.Println("Aplikasi berjalan")
}

func main() {
	runApp(true)
}