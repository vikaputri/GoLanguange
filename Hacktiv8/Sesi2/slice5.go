package main

import "fmt"

func main() {
	var fruit1 = []string{"Apple", "Banana", "Manggo", "Melon", "Lecy"}
	fmt.Println("Fruit1 =>", fruit1)
	var fruit2 = fruit1[1:4]
	fmt.Println("Fruit2 hasil copy fruit1[1:4]=>", fruit2)
	fruit2[0] = "Rambutan"
	fmt.Println("Fruit1 =>", fruit1)
	fmt.Println("Fruit2 hasil copy fruit1[1:4] kemudian fruit2[0] diganti dengan Rambutan=>", fruit2)
	var fruit3 = fruit1[0:]
	fmt.Println("Fruit3 hasil copy fruit1[0:]  =>", fruit3)
	var fruit4 = fruit1[:3]
	fmt.Println("Fruit4 hasil copy fruit1[:3]  =>", fruit4)
}
