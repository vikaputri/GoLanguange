package main

import (
	"fmt"
	"time"
)

func main() {
	now := time.Now()

	fmt.Println(now)
	fmt.Println(now.Year())
	fmt.Println(now.Month())
	fmt.Println(now.Day())
	fmt.Println(now.Hour())
	fmt.Println(now.Minute())
	fmt.Println(now.Second())

	fmt.Println(now.Local())

	utc := time.Date(2020, 10, 10, 10, 10, 10, 10, time.UTC)
	fmt.Println(utc)

	Layout := "2006-01-02"
	parse, _ := time.Parse(Layout, "1998-03-20")
	fmt.Println(parse)

}
