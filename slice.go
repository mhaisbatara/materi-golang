package main

import (
	"fmt"
)

func main() {
	var name = [...]string{"hais", "batara", "joko", "widodo", "agus", "kaka"}

	slice1 := name[4:6]
	fmt.Println(slice1)

	slice2 := name[:3]
	fmt.Println(slice2)

	slice3 := name[3:]
	fmt.Println(slice3)

	slice4 := name[:]
	fmt.Println(slice4)

	fmt.Println("--------------------------------------------------------------------")

	var dasy = [...]string{"senin", "selasa", "rabu", "kamis", "jumat", "sabtu", "minggu"}
	daySlice1 := dasy[4:]
	fmt.Println(daySlice1)

	daySlice1[0] = "jumat baru"
	daySlice1[1] = "sabtu baru"
	fmt.Println(daySlice1)
	


}