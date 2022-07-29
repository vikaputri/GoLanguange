package main

import (
	"fmt"
	"regexp"
)

func main() {
	regex := regexp.MustCompile("e([a-z])o")
	fmt.Println(regex.MatchString("eko"))
	fmt.Println(regex.MatchString("eka"))
	fmt.Println(regex.MatchString("eDo"))

	search := regex.FindAllString("eko eka edo eto enci edi eki", -1)
	fmt.Println(search)

}
