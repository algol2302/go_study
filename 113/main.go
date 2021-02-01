package main

import (
	"fmt"
	"math"
)

func pow2() func() int {
	iterator := 0
	return func() int {
		iterator++
		return int(math.Pow(2, float64(iterator)))
	}
}

func main() {
	x := pow2()
	fmt.Println(x())
	fmt.Println(x())
	fmt.Println(x())
	fmt.Println(x())
}
