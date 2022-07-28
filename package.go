package main

import (
	"belajargolang/helper"
	"fmt"
)

func main() {
	helper.SayHello("Vika")
	//helper.sayGoodbye("Vika") //error
	fmt.Println(helper.Application)
	//fmt.Println(helper.version) //error
}
