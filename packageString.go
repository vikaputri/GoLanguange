package main

import (
	"fmt"
	"strings"
)

func main() {
	fmt.Println(strings.Contains("Vika Putri Ariyanti", "Vika"))
	fmt.Println(strings.Contains("Vika Putri Ariyanti", "Della"))

	fmt.Println(strings.Split("Vika Putri Ariyanti", " "))

	fmt.Println(strings.ToLower("Vika Putri Ariyanti"))
	fmt.Println(strings.ToUpper("Vika Putri Ariyanti"))
	fmt.Println(strings.ToTitle("vika putri ariyanti"))

	fmt.Println(strings.Trim("    Vika Putri Ariyanti       ", " "))
	fmt.Println(strings.Trim("1    Vika Putri Ariyanti       100", " "))

	fmt.Println(strings.ReplaceAll("Vika Putri Ariyanti", "Vika", "Fika"))
}
