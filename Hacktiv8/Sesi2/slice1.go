package main

import "fmt"

func main() {
	var fruits = make([]string, 3)
	_ = fruits

	fruits[0] = "Apple"
	fruits[1] = "Banana"
	fruits[2] = "Manggo"

	fmt.Printf("%#v", fruits)
}
