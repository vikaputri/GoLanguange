package main

import "fmt"

func main() {
	var name = "Ariyanti"
	switch name {
	case "Vika":
		fmt.Println("Hello Vika")
	case "Putri":
		fmt.Println("Hello Putri")
	default:
		fmt.Println("Hi, kenalan dong")
	}

	switch length := len(name); length > 5 {
	case true:
		fmt.Println("Nama terlalu panjang")
	case false:
		fmt.Println("Nama sudah benar")

	}

	length := len(name)
	switch {
	case length > 10:
		fmt.Println("Nama terlalu panjang")
	case length > 5:
		fmt.Println("Nama lumayan panjang")
	default:
		fmt.Println("Nama sudah benar")
	}

}
