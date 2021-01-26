package main

import "fmt"

func main() {
	sum := foo(1, 2, 3, 4, 5, 6, 7, 8, 9)
	fmt.Println("total sum:", sum)
}

func foo(x ...int) int {
	// Variadic parameters
	fmt.Println(x)
	fmt.Printf("%T\n", x)

	sum := 0
	for i, v := range x {
		sum += v
		fmt.Println("for index: ", i, "value:", v, "sum:", sum)
	}

	return sum
}
