package main

import "fmt"

func main() {
	var fruits = make([]string, 3)
	fruits = append(fruits, "Apple", "Banana", "Manggo")

	fmt.Printf("%#v", fruits)
}
