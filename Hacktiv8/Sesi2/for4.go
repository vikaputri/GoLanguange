package main

import "fmt"

func main() {
outerLoop:
	for i := 0; i < 3; i++ {
		fmt.Println("Perulangan ke -", i+1)
		for j := i; j < 3; j++ {
			if i == 2 {
				break outerLoop
			}
			fmt.Print(j, " ")

		}

		fmt.Println()
	}
}
