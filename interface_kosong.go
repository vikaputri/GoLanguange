package main

import "fmt"

func Ups(i int) interface{} {
	if i == 1 {
		return 1
	} else if i == 2 {
		return true
	} else {
		return "Ups"
	}
	//return "Ups"
}

func main() {
	//kosong := Ups(1)
	var kosong interface{} = Ups(1)
	fmt.Println(kosong)

}
