package main

import (
	"fmt"
	"math"
)

func main() {
	fmt.Println(math.Round(67.67))
	fmt.Println(math.Round(67.33))
	fmt.Println(math.Floor(67.67))
	fmt.Println(math.Ceil(67.33))

	fmt.Println(math.Max(20, 10))
	fmt.Println(math.Min(20, 10))
}
