package main

import "fmt"

func main() {
	var currentYear = 2021
	age := currentYear - 1998
	if age < 17 {
		fmt.Println("Kamu belum boleh membuat kartu sim")
	} else {
		fmt.Println("Kamu sudah boleh membuat kartu sim")
	}
}
