package main

import "fmt"

func main() {
	person := map[string]string{ // bebas memilih berapa banyak data tanpa batasan
		"name":   "budi",
		"addres": "jawa",
	}
	fmt.Println(person["name"])
	fmt.Println(person["address"])
	fmt.Println(person)

	fmt.Println("salah") // tidak muncul karena tidak ada data nya dan dia akan membuat data baru sesuai jenisnya tapi berbentuk kosong

	book := make(map[string]string)
	book["tittle"] = "Buku Golang"
	book["author"] = "Bintang"
	book["wrong"] = "ups"

	delete(book, "wrong")

	fmt.Println(book)
}
