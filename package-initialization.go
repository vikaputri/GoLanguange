package main

import (
	"belajargolang/database"
	"fmt"
)

func main() {
	result := database.GetDatabase()
	fmt.Println(result)
}
