package main

import "fmt"

func main() {
	var fruit1 = []string{"Apple", "Banana", "Manggo"}
	var fruit2 = []string{"Watermelon", "Cerry", "Leccy"}
	fruit1 = append(fruit1, fruit2...)

	fmt.Printf("%#v", fruit1)
}
