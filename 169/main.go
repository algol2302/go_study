package main

import "fmt"

// write a program that:
// * puts 100 numbers to a channel
// * pull the numbers off the channel and print them

func main() {
	c := make(chan int)
	go func() {
		for i := 0; i < 100; i++ {
			c <- i
		}
		close(c)
	}()

	fmt.Println(c)

	for v := range c {
		fmt.Println(v)
	}
}
