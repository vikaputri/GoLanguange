package main

import (
	"container/list"
	"fmt"
)

func main() {
	data := list.New()
	data.PushBack("Putri")
	data.PushBack("Ariyanti")
	data.PushFront("Vika")

	fmt.Println(data.Front().Value)
	fmt.Println(data.Back().Value)
	fmt.Println(data.Front().Prev())
	fmt.Println(data.Back().Next())

	for e := data.Front(); e != nil; e = e.Next() {
		fmt.Println(e.Value)
	}

	for e := data.Back(); e != nil; e = e.Prev() {
		fmt.Println(e.Value)
	}

}
