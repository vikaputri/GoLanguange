package main

import (
	"fmt"
	"strings"
)

func main() {
	var fruits = [3]string{"Apple", "Banana", "Manggo"}
	for i, v := range fruits {
		fmt.Printf("Index:%d, Value %s\n", i, v)
	}
	fmt.Println(strings.Repeat("#", 25))

	for i := 0; i < len(fruits); i++ {
		fmt.Printf("Index:%d, Value %s\n", i, fruits[i])
	}
}
