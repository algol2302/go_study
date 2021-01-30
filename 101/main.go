package main

import "fmt"

func main() {
	n := factorialRecursion(4)
	fmt.Println(n)

	number := 4
	m := factorialLoop(number)
	fmt.Println(m)
	fmt.Println(number)
}

func factorialRecursion(n int) int {
	if n == 0 {
		return 1
	}
	return n * factorialRecursion(n-1)
}

func factorialLoop(n int) int {
	factorial := 1
	for ; n > 0; n-- {
		factorial *= n
	}
	return factorial
}
