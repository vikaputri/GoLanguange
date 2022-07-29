package main

import "fmt"

type Customer struct {
	Name, Address string
	Age           int
}

func main() {
	vika := Customer{
		Name:    "Vika Putri Ariyanti",
		Address: "Pekalongan",
		Age:     17,
	}
	fmt.Println(vika)

	della := Customer{"Della Oktavia", "Pekalongan", 17}
	fmt.Println(della)
}
