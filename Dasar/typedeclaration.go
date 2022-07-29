package main

import "fmt"

func main() {
	type NoKTP string
	type Married bool

	var ktpVika NoKTP = "11111111111"
	fmt.Println(ktpVika)

	var marriedStatus Married = true
	fmt.Println(marriedStatus)
}
