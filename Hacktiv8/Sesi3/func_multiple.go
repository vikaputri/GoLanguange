package main

import (
	"fmt"
	"math"
)

func main() {
	var diameter float64 = 15
	var area, circumference = calculate(diameter)
	fmt.Println(area)
	fmt.Println(circumference)

	var area1, circumference1 = calculate1(diameter)
	fmt.Println(area1)
	fmt.Println(circumference1)
}

func calculate(d float64) (float64, float64) {
	var area float64 = math.Pi * math.Pow(d/2, 2)
	var circumference = math.Pi * d

	return area, circumference
}

func calculate1(d float64) (area1 float64, circumference1 float64) {
	area1 = math.Pi * math.Pow(d/2, 2)
	circumference1 = math.Pi * d

	return
}
