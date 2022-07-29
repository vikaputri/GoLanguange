package main

import "fmt"

func main() {
	var name = "Putri"

	if name == "Vika" {
		fmt.Println("Hello Vika")
	} else if name == "Putri" {
		fmt.Println("Hello Putri")
	} else if name == "Ariyanti" {
		fmt.Println("Hello Ariyanti")
	} else {
		fmt.Println("Hi, Kenalan dong")
	}

	if length := len(name); length > 5 {
		fmt.Println("Terlalu Panjang")
	} else {
		fmt.Println("Nama sudah benar")
	}

}
