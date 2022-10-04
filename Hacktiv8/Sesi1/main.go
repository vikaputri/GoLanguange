package main

import "fmt"

func main() {
	fmt.Println("Hello World")

	//variabel
	var name string = "Airell"
	var age int = 23
	fmt.Println("Ini adalah namanya ==>", name)
	fmt.Println("Ini adalah namanya ==>", age)
	fmt.Printf("%T, %T", name, age)

	var name1 string
	var age1 int
	name1 = "Airell"
	age1 = 23
	fmt.Println("Ini adalah namanya ==>", name1)
	fmt.Println("Ini adalah namanya ==>", age1)

	name2 := "Airell"
	age2 := 23
	fmt.Printf("%T, %T", name2, age2)

	var student1, student2, student3 string = "satu", "dua", "tiga"
	var first, second, third int
	first, second, third = 1, 2, 3
	name3, age3, address := "Airell", 23, "Jalan Sudirman"
	fmt.Println(student1, student2, student3)
	fmt.Println(first, second, third)
	fmt.Println(name3, age3, address)

}
