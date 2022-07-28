package main

import "fmt"

type Customer struct {
	Name, Address string
	Age           int
}

func (customer Customer) sayHello() {
	fmt.Println("Hello, My Name is", customer.Name)
}

func (customer Customer) sayHi(name string) {
	fmt.Println("Hello", name, "My Name is", customer.Name)
}

func main() {
	rully := Customer{Name: "Rully"}
	rully.sayHello()

	rully.sayHi("Della")
}
