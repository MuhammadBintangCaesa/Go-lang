package main

import "fmt"

func main() {
	var name [3]string //khusus string dan ga bsa tanpa ga di batasin atau langsung di deklarasikan

	name[0] = "Eka"
	name[1] = "Gustiawana"
	name[2] = "Putrag"

	name[0] = "Eka"
	name[1] = "Gustiawana"
	name[2] = "Putrag"

	fmt.Println(name[0])
	fmt.Println(name[1])
	fmt.Println(name[2])

	var values = [...]int{90, 80, 70, 60, 100, 30} // cara singkat
	fmt.Println(values[0])
	fmt.Println(values[1])
	fmt.Println(values[2])
	fmt.Println(values[3])

	fmt.Println(len(values))
	fmt.Println(values)

}
