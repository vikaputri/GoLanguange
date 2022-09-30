package main

import (
	"fmt"
	"strings"
)

type Employee struct {
	name     string
	age      int
	division string
}

type Person struct {
	name string
	age  int
}

type Employe struct {
	division string
	person   Person
}

func main() {
	var employee Employee
	employee.name = "Airell"
	employee.age = 23
	employee.division = "Curiculum Developer"

	fmt.Println(employee.name)
	fmt.Println(employee.age)
	fmt.Println(employee.division)

	var employee1 Employee
	employee1.name = "Airell"
	employee1.age = 23
	employee1.division = "Curiculum Developer"
	var employee2 = Employee{name: "Ananda", age: 23, division: "Finance"}
	fmt.Printf("Employee1: %+v\n", employee1)
	fmt.Printf("Employee2: %+v\n", employee2)

	var employee3 *Employee = &employee1
	fmt.Println("Employee1 name :", employee1.name)
	fmt.Println("Employee3 name :", employee3.name)

	fmt.Println(strings.Repeat("#", 20))

	employee3.name = "Ananda"
	fmt.Println("Employee1 name :", employee1.name)
	fmt.Println("Employee3 name :", employee3.name)

	var employee4 = Employe{}
	employee4.person.name = "Airell"
	employee4.person.age = 23
	employee4.division = "Curiculum Developer"

	fmt.Printf("%+v\n", employee4)

}
