package main

import "fmt"

func main() {
	type NOktp string

	var ktpeko NOktp = "1111111"

	var contoh string = "2222222"
	var contohktp NOktp = NOktp(contoh)

	fmt.Println(ktpeko)
	fmt.Println(contohktp)

}