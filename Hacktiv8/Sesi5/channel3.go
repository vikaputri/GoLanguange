package main

import (
	"fmt"
	"time"
)

func main() {
	c1 := make(chan int)
	go func(c chan int) {
		fmt.Println("func gorountine starts sending data into the channel")
		c <- 10
		fmt.Println("func gorountine after sending data into the channel")
	}(c1)

	fmt.Println("main gorountine sleeps for 2 second")
	time.Sleep(time.Second * 2)

	fmt.Println("main gorountine starts received data")
	d := <-c1
	fmt.Println("main gorountine received data:", d)

	close(c1)
	time.Sleep(time.Second)
}
