package main

import "fmt"

func logging() {
	fmt.Println("Selesai memanggil function")
}

func runApplication(value int) {
	defer logging() //deffer : function yang dapat dieksekusi walaupun error
	fmt.Println("Run Application")
	result := 10 / value
	fmt.Println("Result", result)
}

func main() {
	runApplication(0)
}
