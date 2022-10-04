package main

import "fmt"

func main() {
	callDeferFunc()
	fmt.Println("Hai everyone")
}

func callDeferFunc() {
	defer deferFunc()
}

func deferFunc() {
	fmt.Println("defer function strat to execute")
}
