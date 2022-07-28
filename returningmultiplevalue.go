package main

import "fmt"

func getFullName() (string, string, string) {
	return "Vika", "Putri", "Ariyanti"
}

func main() {
	firstName, middleName, lastName := getFullName()
	fmt.Println(firstName)
	fmt.Println(middleName)
	fmt.Println(lastName)

	namaawal, _, _ := getFullName()
	fmt.Println(namaawal)
}
