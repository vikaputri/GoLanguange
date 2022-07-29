package main

import (
	"fmt"
	"reflect"
)

type Sample struct {
	Name string `required:"true" max:"10"`
}

func IsValid(data interface{}) bool {
	t := reflect.TypeOf(data)
	for i := 0; i < t.NumField(); i++ {
		field := t.Field(i)
		if field.Tag.Get("required") == "true" {
			value := reflect.ValueOf(data).Field(i).Interface()
			if value == "" {
				return false
			}
		}
	}
	return true
}

type ContohLagi struct {
	Name        string `required:"true"`
	Description string `required:"true"`
}

func main() {
	sample := Sample{"Vika"}

	sampleType := reflect.TypeOf(sample)
	fmt.Println(sampleType)
	fmt.Println(sampleType.NumField())

	structField := sampleType.Field(0)
	fmt.Println(structField.Name)

	required := structField.Tag.Get("required")
	fmt.Println(required)
	fmt.Println(structField.Tag.Get("max"))

	fmt.Println(IsValid(sample))

	contoh := ContohLagi{"", ""}
	fmt.Println(IsValid(contoh))

}
