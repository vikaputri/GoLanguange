package main

import "fmt"

func main() {
	var fruit1 = []string{"Apple", "Banana", "Manggo"}
	var fruit2 = []string{"Watermelon", "Cerry", "Leccy"}
	nn := copy(fruit1, fruit2)

	fmt.Println("Fruit1 =>", fruit1)
	fmt.Println("Fruit2 =>", fruit2)
	fmt.Println("Copies Element =>", nn)
}
