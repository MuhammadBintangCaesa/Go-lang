package main

import "fmt"

func main() {

	//counter := 1

	//for counter <= 10 {
	//fmt.Println("perulangan ke ", counter)
	//counter++

	for counter := 1; counter <= 10; counter++ {
		fmt.Println("perulangan ke ", counter)
	}
	fmt.Println("selesai")

	// manual
	names := []string{"eko", "ujang", "Jono", "anton"}
	for i := 0; i < len(names); i++ {
		fmt.Println(names[i])
	}

	for index, name := range names { // mengetahui kunci
		fmt.Println("index", index, "=", name)

	}
	for _, name := range names { // menggunakan _, untuk menghilangkan kunci index
		fmt.Println(name)
	}
}
