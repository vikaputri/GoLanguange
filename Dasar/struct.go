package main

import "fmt"

type Customer struct {
	Name, Address string
	Age           int
}

func main() {
	var vika Customer
	vika.Name = "Vika Putri Ariyanti"
	vika.Address = "Pekalongan"
	vika.Age = 17
	fmt.Println(vika)
	fmt.Println(vika.Name)
}
