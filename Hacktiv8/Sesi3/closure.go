package main

import "fmt"

func main() {
	var evenNumber = func(numbers ...int) []int {
		var result []int
		for _, v := range numbers {
			if v%2 == 0 {
				result = append(result, v)
			}
		}
		return result
	}

	var numbers = []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}
	fmt.Println(evenNumber(numbers...))
}
