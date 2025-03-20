package main

import "fmt"

func main() {

	names := [...]string{"Bambang", "Weldi", "Sina", "Eja", "Kula", "Selena"} // bambang index 0, weldi, index 1, sina index 2, eja index 3 , kula index 4, selena index 5

	slice1 := names[0:6] // index 0 - 6 // low : high
	fmt.Println(slice1)

	slice2 := names[:3] // index 0 - 3 bambang weldi sina // high
	fmt.Println(slice2)

	slice3 := names[1:3] // index 1 sampai 2 // weldi sina // low : high
	fmt.Println(slice3)

	slice4 := names[1:] // index 1 di coret jadi "Weldi", "Sina", "Eja", "Kula", "Selena" // low
	fmt.Println(slice4)

	slice5 := names[:] // keseluruhan
	fmt.Println(slice5)

	days := [...]string{"senin", "selasa", "rabu", "kamis", "jumat", "sabtu", "minggu"}

	daySlice1 := days[5:] // sabtu dan minggu
	fmt.Println(daySlice1)

	daySlice1[0] = "Sabtu baru"
	daySlice1[1] = "Minggu baru"
	fmt.Println(daySlice1)
	fmt.Println(days)

	daySlice2 := append(daySlice1, "libur lama")
	//daysBaru := [...]string{"senin", "selasa", "rabu", "kamis", "jumat", "sabtu", "minggu", "libur lama"} // buat array baru data dayslice1 takkan hilang hanya di tambahkan ke data days slice 2 yg baru // data array yg baru ada di dayslice2 jadi
	// tidak bsa di akses
	daySlice2[0] = "sabtu lama"

	fmt.Println(daySlice1)
	fmt.Println(daySlice2)
	fmt.Println(days)

	var newSlice []string = make([]string, 2, 5)
	newSlice[0] = "eko"
	newSlice[1] = "eko"

	fmt.Println(newSlice)
	fmt.Println(len(newSlice))
	fmt.Println(cap(newSlice))

	newSlice2 := append(newSlice, "Eko")
	fmt.Println(newSlice2)
	fmt.Println(len(newSlice2))
	fmt.Println(newSlice2)

	newSlice2[0] = "Budi"
	fmt.Println(newSlice2)
	fmt.Println(newSlice)

	fromSlice := days[:]
	toSlice := make([]string, len(fromSlice), cap(fromSlice)) // copy hari (semua days slice yang  sudah di buat )

	copy(toSlice, fromSlice)

	fmt.Println(fromSlice)
	fmt.Println(toSlice)

	iniArray := [...]int{1, 2, 3, 4, 5} // Perbedaan array dan slice
	iniSlice := []int{1, 2, 3, 4, 5}

	fmt.Println(iniArray)
	fmt.Println(iniSlice)
}
