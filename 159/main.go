package main

import "fmt"

func main() {
	c := make(chan int)
	length := 5
	// send
	go foo(c, length)

	// recieve
	bar(c, length)

	fmt.Println("about to exit")
}

// send
func foo(c chan<- int, length int) {
	for i := 0; i < length; i++ {
		c <- i
	}
	close(c)
}

// recieve
func bar(c <-chan int, length int) {
	for i := 0; i < length; i++ {
		fmt.Println(<-c)
	}
}
