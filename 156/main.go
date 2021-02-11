package main

import (
	"fmt"
)

func main() {
	c := make(chan int)    // bidirectional channel
	cr := make(<-chan int) // recieve channel
	cs := make(chan<- int) // send channel

	fmt.Println("------initial------")
	fmt.Printf("c\t%T\n", c)
	fmt.Printf("cr\t%T\n", cr)
	fmt.Printf("cs\t%T\n", cs)

	fmt.Println("--after conversion--")
	fmt.Printf("cr\t%T\n", (<-chan int)(c))
	fmt.Printf("cs\t%T\n", (chan<- int)(c))

}
