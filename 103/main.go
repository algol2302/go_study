package main

import "fmt"

func main() {
	values := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 42}

	sum1 := foo(values...)
	fmt.Println(sum1)

	sum2 := bar(values)
	fmt.Println(sum2)
}

func foo(nn ...int) int {
	total := 0
	for _, v := range nn {
		total += v
	}
	return total
}

func bar(nn []int) int {
	total := 0
	for _, v := range nn {
		total += v
	}
	return total
}
