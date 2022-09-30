package main

import (
	"fmt"
	"runtime"
	"time"
)

func main() {
	fmt.Println("main execution started")
	go firstProses(8)
	SecondProses(8)
	fmt.Println("No of Gorountines", runtime.NumGoroutine())
	time.Sleep(time.Second * 2)
	fmt.Println("main execution ended")

}

func firstProses(index int) {
	fmt.Println("First proses func started")
	for i := 0; i <= index; i++ {
		fmt.Println("i=", i)
	}
	fmt.Println("First process func ended")
}

func SecondProses(index int) {
	fmt.Println("Second proses func started")
	for j := 0; j <= index; j++ {
		fmt.Println("j=", j)
	}
	fmt.Println("First process func ended")
}
