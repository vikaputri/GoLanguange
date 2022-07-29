package main

import "fmt"

func random() interface{} {
	return "OK"
}

func main() {
	result := random()
	//resultString := result.(string)
	//fmt.Println(resultString)

	//resultInt := result.(int) //panic jika type data beda
	//fmt.Println(resultInt)

	switch value := result.(type) {
	case string:
		fmt.Println("Value", value, "is String")
	case int:
		fmt.Println("Value", value, "is int")
	default:
		fmt.Println("Unknown Type")
	}
}
