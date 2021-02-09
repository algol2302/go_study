package main

import (
	"fmt"
	"runtime"
	"sync"
)

func main() {
	fmt.Println("Num Goroutines:", runtime.NumGoroutine())

	var wg sync.WaitGroup
	wg.Add(3)

	go func() {
		fmt.Println("this is main goroutine")
		fmt.Println("Num Goroutines:", runtime.NumGoroutine())
		go func() {
			fmt.Println("this is first additional goroutine")
			fmt.Println("Num Goroutines:", runtime.NumGoroutine())
			wg.Done()
		}()
		go func() {
			fmt.Println("this is second additional goroutine")
			fmt.Println("Num Goroutines:", runtime.NumGoroutine())
			wg.Done()
		}()
		wg.Done()
	}()

	wg.Wait()

	fmt.Println("Final num Goroutines:", runtime.NumGoroutine())
}
