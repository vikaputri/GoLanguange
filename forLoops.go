package main

import "fmt"

func main() {
	counter := 1
	for counter <= 10 {
		fmt.Println("Perulangan ke ", counter)
		counter++
	}

	for counter := 1; counter <= 10; counter++ {
		fmt.Println("Perulangan ke ", counter)
	}

	slice := []string{"Vika", "Putri", "Ariyanti"}
	for i := 0; i < len(slice); i++ {
		fmt.Println(slice[i])
	}

	for index, name := range slice {
		fmt.Println("index", index, "=", name)
	}

	person := make(map[string]string)
	person["name"] = "Vika"
	person["address"] = "Pekalongan"

	for key, value := range person {
		fmt.Println(key, "=", value)
	}

}
