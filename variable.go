package main

import "fmt"

func main() {
	// var name string

	// name = "M. Hais Batara"
	// fmt.Println(name)

	// name = "jokowi"
	// fmt.Println(name)

	// tidak wajib menyantumkan Tipe data karena golong sudah membacanya
	// var name = "M. Hais Batara"
	// fmt.Println(name)

	// name = "jokowi"
	// fmt.Println(name)



	// bisa dengan := tanpa ada var di awal saja
	name := "M. Hais Batara"
	fmt.Println(name)

	name = "jokowi"
	fmt.Println(name)

	var(
		namadepan = "M."
		namatenga = " Hais"
		namabelakang = " Batara"
	)

	fmt.Print(namadepan)
	fmt.Print(namatenga)
	fmt.Print(namabelakang)
}