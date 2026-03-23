package main

import "fmt"

func main() {
	var nilaiakhir = 90
	var absensi = 70

	var lulus = nilaiakhir > 75 && absensi > 85

	fmt.Println(lulus)
}