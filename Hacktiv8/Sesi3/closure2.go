package main

import "fmt"

func main() {
	var numbers = []int{2, 5, 6, 7, 8, 9, 10, 99}
	var find = findOddNumbers(numbers, func(number int) bool {
		return number%2 != 0
	})
	fmt.Println((find))

}

func findOddNumbers(numbers []int, callback func(int) bool) int {
	var totalOddNumbers int
	for _, v := range numbers {
		if callback(v) {
			totalOddNumbers++
		}
	}
	return totalOddNumbers

}
