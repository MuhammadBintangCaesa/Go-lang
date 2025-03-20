package main

import "fmt"

func main() {
	nama := "zainal"

	switch nama {
	case "eko":
		fmt.Println("Hello eko")
	case "zo":
		fmt.Println("Hellow zo")
	default:
		fmt.Println("Hi boleh kenalan?")

		switch length := len(nama); length > 5 {
		case true:
			fmt.Println("nama terlalu panjang")
		case false:
			fmt.Println("Nama sudah benar")
		}

		length := len(nama)
		switch {
		case length > 10:
			fmt.Println("Nama teralu panjang")
		case length < 8:
			fmt.Println("Nama lumayan panjang")
		default:
			fmt.Println("Nama Sudah benar")
		}
	}
}
