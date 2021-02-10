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
	var mu sync.Mutex
	var wg sync.WaitGroup
	wg.Add(100)

	for i := 1; i <= 100; i++ {
		go func() {
			mu.Lock()

			v := counter
			v++
			counter = v
			fmt.Println("Middle num Goroutines:", runtime.NumGoroutine(),
				"value:", v, "counter:", counter)

			mu.Unlock()

			wg.Done()
		}()

	}

	wg.Wait()
	fmt.Println("Final num Goroutines:", runtime.NumGoroutine())
	fmt.Println("Final counter:", counter)
}
