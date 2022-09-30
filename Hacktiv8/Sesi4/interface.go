package main

import (
	"fmt"
	"math"
)

type shape interface {
	area() float64
	perimeter() float64
}

type rectangle struct {
	width, height float64
}

type circle struct {
	radius float64
}

func (c circle) area() float64 {
	return math.Pi * math.Pow(c.radius, 2)
}

func (r rectangle) area() float64 {
	return r.height * r.width
}

func (c circle) perimeter() float64 {
	return 2 * math.Pi * c.radius
}

func (r rectangle) perimeter() float64 {
	return 2 * (r.height + r.width)
}

func main() {
	var c1 shape = circle{radius: 5}
	var c2 shape = rectangle{width: 3, height: 2}

	fmt.Printf("Type of c1: %T \n", c1)
	fmt.Printf("Type of c2: %T \n", c2)
	fmt.Println("Luas Lingkaran : ", c1.area())
	fmt.Println("Keliling Lingkaran : ", c1.perimeter())
	fmt.Println("Luas Panjang : ", c2.area())
	fmt.Println("Keliling Persegi Panjang: ", c2.perimeter())

}
