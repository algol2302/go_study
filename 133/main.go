package main

import (
	"fmt"
	"time"
)

func doSomething(x int) int {
	return x * 5
}

func say(s string) {
	for i := 0; i < 5; i++ {
		time.Sleep(100 * time.Millisecond)
		fmt.Println(s)
	}
}

func main() {
	ch := make(chan int)
	go func() {
		ch <- doSomething(5)
	}()
	a := <-ch
	fmt.Println(a)

	// go say("world")
	// say("hello")
}
