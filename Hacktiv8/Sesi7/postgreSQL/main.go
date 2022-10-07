package main

import (
	"database/sql"
	"fmt"

	_ "github.com/lib/pq"
)

var (
	db  *sql.DB
	err error
)

type Employee struct {
	ID        int
	Full_name string
	Email     string
	Age       int
	Division  string
}

func main() {
	db, err = sql.Open("postgres", "postgres://postgres:password@localhost/dev?sslmode=disable")
	if err != nil {
		panic(err)
	}
	defer db.Close()

	err = db.Ping()
	if err != nil {
		panic(err)
	}

	fmt.Println("Successfully connected to database")

	CreateEmployee()
}

func CreateEmployee() {
	var employee = Employee{}

	sqlStatement := `
	INSERT INTO employees (id, full_name, email, age, division) 
	VALUES ($1, $2, $3, $4, $5)
	Returning *
	`

	err = db.QueryRow(sqlStatement, 1, "Vika", "vikaputri@gmail.com", 23, "IT").
		Scan(&employee.ID, &employee.Full_name, &employee.Email, &employee.Age, &employee.Division)

	if err != nil {
		panic(err)
	}

	fmt.Printf("New Employee Data : %+v\n", employee)
}
