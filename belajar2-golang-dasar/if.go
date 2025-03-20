package main

import "fmt"

func main() {
	nama := "zo"

	if nama == "eko" {
		fmt.Println("Hello eko")
	} else {
		fmt.Println("Hi, aku boleh kenalan?")
	}

	length := len(nama)
	if length > 5 {
		fmt.Println("Nama terlalu panjang")
	} else {
		fmt.Println("Nama anda sudah benar")
	}
}
