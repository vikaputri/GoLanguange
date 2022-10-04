package main

import (
	"fmt"
	"os"
)

func main() {
	defer fmt.Println("Invoce with defer")
	fmt.Println("Before Exiting")
	os.Exit(1)
}
