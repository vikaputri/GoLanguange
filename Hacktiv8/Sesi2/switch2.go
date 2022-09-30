package main

import "fmt"

func main() {
	var score = 6
	switch {
	case score == 8:
		fmt.Println("Perfect")
	case (score < 8) && (score > 3):
		fmt.Println("Awesome")
	default:
		fmt.Println("Study Hard")
		fmt.Println("You Need to Learn More")
	}
}
