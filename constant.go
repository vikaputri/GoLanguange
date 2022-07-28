package main

import "fmt"

func main() {
	const firstName string = "Vika"
	const lastName = "Putri"
	const value = 1000

	fmt.Println(firstName)
	fmt.Println(lastName)
	//fmt.Println(value)

	const (
		firstName1 string = "Vika"
		lastName1         = "Putri"
	)
	fmt.Println(firstName1)
	fmt.Println(lastName1)
}
