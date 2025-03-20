package main

import "fmt"

func main() {

	type NoKTP string

	var UjangKTP NoKTP = "1111111"

	var contoh string = "222222"
	var contohKTP NoKTP = NoKTP(contoh)

	fmt.Println(UjangKTP)
	fmt.Println(contohKTP)
}
