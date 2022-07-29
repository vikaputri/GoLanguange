package helper

import "fmt"

var Application = "Belajar Golang"
var version = 1

func SayHello(name string) {
	fmt.Println("Hello", name)
}

func sayGoodbye(name string) {
	fmt.Println("Goodbye", name)
}
