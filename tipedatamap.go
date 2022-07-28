package main

import "fmt"

func main() {
	person := map[string]string{
		"name":    "Vika",
		"address": "Pekalongan",
	}
	person["title"] = "Programmer"

	fmt.Println(person)
	fmt.Println(person["name"])
	fmt.Println(person["address"])
	fmt.Println(len(person))

	var book map[string]string = make(map[string]string)
	book["title"] = "Golang"
	book["author"] = "Vika"
	book["ups"] = "Salah"

	fmt.Println(book)
	fmt.Println(len(book))

	delete(book, "ups")
	fmt.Println(book)
	fmt.Println(len(book))

}
