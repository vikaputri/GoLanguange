package main

import "fmt"

func NewMap(name string) map[string]string {
	if name == "" {
		return nil
	} else {
		return map[string]string{
			"name": name,
		}
	}
}

func main() {
	var person map[string]string = nil
	fmt.Println(person)

	person1 := NewMap("Vika")
	fmt.Println(person1)

	var person2 map[string]string
	if person2["name"] == "" {
		fmt.Println("Data Kosong")
	} else {
		fmt.Println(person2)
	}

	var person3 map[string]string = NewMap("")
	if person3 == nil {
		fmt.Println("data kosong")
	} else {
		fmt.Println(person3)
	}
}
