package main

import "fmt"

func getGoodBye(name string) string {
	return "Good Bye " + name
}

func main() {
	SayGoodBye := getGoodBye
	result := SayGoodBye("Vika")
	fmt.Println(result)
	fmt.Println(SayGoodBye("Vika"))
}
