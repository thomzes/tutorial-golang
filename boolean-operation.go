package main

import "fmt"

func main() {
	var nilaiAkhir = 90
	var absensi = 90

	var lulusNilaiAkhir bool = nilaiAkhir > 80
	var lulusAbsensi bool = absensi > 80

	var lulus bool = lulusAbsensi && lulusNilaiAkhir
	fmt.Println(lulus)

	fmt.Println(nilaiAkhir >= 80 && absensi >= 80)

}
