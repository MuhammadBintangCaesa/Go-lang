package main

import "fmt"

func main() {
	var a = 20
	var b = 15
	var c = 2
	var d = a + b/c

	fmt.Println(d)

	var i = 20
	i += 10        // i = i + 10
	fmt.Println(i) // mencetak perhitungan nilai i di awal

	i /= 2         // i = i : 2
	fmt.Println(i) // mencetak perhitunga nilai i di awal dan di bagi di akhir sesuai perintah

	var j = 2
	j++ // j = 2 + 1   // step 1
	j++ // j = 3 + 1 // step 2
	j-- // j = 4 - 1 // step 3

	fmt.Println(j)
}
