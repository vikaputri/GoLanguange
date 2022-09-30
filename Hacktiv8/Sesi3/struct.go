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

//Embedded struct
type Employe struct {
	division string
	person   Person
}

func main() {
	//Giving value to struct
	var employee Employee
	employee.name = "Airell"
	employee.age = 23
	employee.division = "Curiculum Developer"

	fmt.Println(employee.name)
	fmt.Println(employee.age)
	fmt.Println(employee.division)

	//Initializing struct
	var employee1 Employee
	employee1.name = "Airell"
	employee1.age = 23
	employee1.division = "Curiculum Developer"
	var employee2 = Employee{name: "Ananda", age: 23, division: "Finance"}
	fmt.Printf("Employee1: %+v\n", employee1)
	fmt.Printf("Employee2: %+v\n", employee2)

	//Pointer to a struct
	var employee3 *Employee = &employee1
	fmt.Println("Employee1 name :", employee1.name)
	fmt.Println("Employee3 name :", employee3.name)

	fmt.Println(strings.Repeat("#", 20))

	employee3.name = "Ananda"
	fmt.Println("Employee1 name :", employee1.name)
	fmt.Println("Employee3 name :", employee3.name)

	//Embedded struct
	var employee4 = Employe{}
	employee4.person.name = "Airell"
	employee4.person.age = 23
	employee4.division = "Curiculum Developer"

	fmt.Printf("%+v\n", employee4)

	//Anonymous struct tanpa pengisian field
	var employee5 = struct {
		person   Person
		division string
	}{}
	employee5.person = Person{name: "Airell", age: 23}
	employee5.division = "Curiculum Developer"

	//Anonymous struct pengisian field
	var employee6 = struct {
		person   Person
		division string
	}{
		Person{name: "Ananda", age: 23},
		"Finance",
	}
	fmt.Printf("%+v\n", employee5)
	fmt.Printf("%+v\n", employee6)

}
