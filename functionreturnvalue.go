package main

import "fmt"

func gethello(name string) string {
	//return "Hello " + name
	if name == "" {
		return "Hello Bro"
	} else {
		return "Hello " + name
	}
}

func main() {
	result := gethello("Vika")
	fmt.Println(result)

	fmt.Println(gethello(""))
}
