package main

import "fmt"

func main() {
	var names [3]string
	names[0] = "Vika"
	names[1] = "Putri"
	names[2] = "Ariyanti"

	fmt.Println(names[0])
	fmt.Println(names[1])
	fmt.Println(names[2])

	var value = [3]int{
		98,
		80,
		90,
	}
	fmt.Println(value)
	fmt.Println(value[0])
	fmt.Println(value[1])
	fmt.Println(value[2])

	fmt.Println(len(value))

	var lagi [10]string
	fmt.Println(len(lagi))

}
