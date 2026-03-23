package main

import "fmt"

func main() {
	var array [3]string
	array[0] = "M"
	array[1] = "Hais"
	array[2] = "Batara"

	fmt.Println(array[0])
	fmt.Println(array[1])
	fmt.Println(array[2])

	array[0] = "P"
	fmt.Println(";;;;;;;;;;;;;;;;;;;;;;;;;")
	fmt.Println(array[0])
	fmt.Println(array[1])
	fmt.Println(array[2])


	var values = [3]int {
		90,
		2,
		1,
	}

	fmt.Println(values)
	fmt.Println(values[0])
	fmt.Println(values[1])
	fmt.Println(values[2])



	var values1 = [...]int {
		90,
		2,
		1,
		2,
		9,
	}
	fmt.Println("--------------------------")
	fmt.Println(len(values1))

}