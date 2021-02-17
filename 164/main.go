package main

import (
	"fmt"
)

func main() {
	// using anonymous function
	c1 := make(chan int)

	go func() {
		c1 <- 42
	}()

	fmt.Println(<-c1)

	// using buffered channel
	c2 := make(chan int, 1)

	c2 <- 43

	fmt.Println(<-c2)

}
