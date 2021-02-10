package main

import (
	"fmt"
	"runtime"
	"sync"
)

func main() {
	fmt.Println("Num CPUs:", runtime.NumCPU())
	fmt.Println("Num Goroutines:", runtime.NumGoroutine())

	counter := 0

	var wg sync.WaitGroup
	wg.Add(100)

	for i := 1; i <= 100; i++ {
		go func() {

			v := counter
			runtime.Gosched()
			v++
			counter = v
			fmt.Println("Middle num Goroutines:", runtime.NumGoroutine(),
				"value:", v, "counter:", counter)

			wg.Done()
		}()
	}

	wg.Wait()
	fmt.Println("Final num Goroutines:", runtime.NumGoroutine())
}
