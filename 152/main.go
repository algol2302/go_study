package main

import (
	"fmt"
	"runtime"
	"sync"
	"sync/atomic"
)

func main() {
	fmt.Println("Num CPUs:", runtime.NumCPU())
	fmt.Println("Num Goroutines:", runtime.NumGoroutine())

	var counter int64

	var wg sync.WaitGroup
	wg.Add(100)

	for i := 1; i <= 100; i++ {
		go func() {
			atomic.AddInt64(&counter, 1)

			runtime.Gosched()

			fmt.Println("Middle num Goroutines:", runtime.NumGoroutine(),
				"counter:", atomic.LoadInt64(&counter))

			wg.Done()
		}()
	}

	wg.Wait()
	fmt.Println("Final num Goroutines:", runtime.NumGoroutine())
	fmt.Println("Final counter:", counter)
}
