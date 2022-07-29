package main

import "fmt"

func getFullName2() (firstName, middleName, lastName string) {
	firstName = "Vika"
	middleName = "Putri"
	lastName = "Ariyanti"

	// return firstName, middleName, lastName
	return
}

func main() {
	firstName, middleName, lastName := getFullName2()
	fmt.Println(firstName)
	fmt.Println(middleName)
	fmt.Println(lastName)

	namaawal, _, _ := getFullName2()
	fmt.Println(namaawal)
}
