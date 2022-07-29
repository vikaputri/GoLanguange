package main

import "fmt"

func sayhello() {
	fmt.Println("Hello")
}

func main() {
	sayhello()
	for i := 0; i < 10; i++ {
		sayhello()
	}
}
