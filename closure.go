package main

import "fmt"

func main() {
	counter := 0
	name := "Eko"
	increment := func() {
		name := "Budi"
		fmt.Println("Increment")
		fmt.Println(name)
		counter++
	}

	increment()
	increment()
	fmt.Println(counter)
	fmt.Println(name)

}
