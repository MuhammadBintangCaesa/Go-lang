package main

import "fmt"

func main() {

	var nilaiAkhir = 90
	var absensi = 70

	var lulusnilaiAkhir bool = nilaiAkhir > 80
	var lulusabsensi bool = absensi > 80

	var lulus bool = lulusnilaiAkhir && lulusabsensi

	fmt.Println(lulus)
}
