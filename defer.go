package main

import "fmt"

// defer untuk menjalankan function tetap berjalan meskipun di dalam code (selain function defer) itu error

func logging() {
	fmt.Println("Selesai memanggil function")
}

func runApplication(value int) {
	defer logging()

	fmt.Println("Run Application")
	result := 10 / value
	fmt.Println("Result ", result)


}

func main() {
	runApplication(10)
}

