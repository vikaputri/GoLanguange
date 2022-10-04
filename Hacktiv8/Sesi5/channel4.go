package main

import (
	"fmt"
	"time"
)

func main() {
	c1 := make(chan int, 3)
	go func(c chan int) {
		for i := 1; i <= 5; i++ {
			fmt.Printf("func gorountine #%d starts sending data into the channel\n", i)
			c <- i
			fmt.Printf("func gorountine #%d after sending data into the channel\n", i)
		}
		close(c)
	}(c1)

	fmt.Println("main gorountine sleeps for 2 second")
	time.Sleep(time.Second * 2)

	for v := range c1 {
		fmt.Println("main gorountine value from channel:", v)
	}
}
