package main

import "fmt"

func main() {

	//const artinya permanen tidak bisa di rubah dengan cara memanggil nama variablenya lagi
	// const namad string = "hais"
	// const namab = "batara"

	// fmt.Println(namad)
	// fmt.Println(namab)

	// eror
	// namad = "jokowi"
	
	const(
		namad = "hais"
		namab = "batara"
	)

	fmt.Println(namad)
	fmt.Println(namab)

 
}