package main

import (
	"fmt"
)

func main() {
	// using channel
	c1 := make(chan int)

	go func() {
		c1 <- 42
	}()

	fmt.Println(<-c1)

	// add two value
	c2 := make(chan int, 2)

	go func() {
		c2 <- 42
		c2 <- 43
	}()

	fmt.Println(<-c2)
	fmt.Println(<-c2)
}
