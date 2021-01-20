package main

import "fmt"

func main() {
	x := []int{4, 5, 7, 8, 42}
	fmt.Println(x)
	x = append(x, 77, 99, 88, 101)
	fmt.Println(x)

	y := []int{123, 345, 567}
	fmt.Println(y)
	x = append(x, y...)
	fmt.Println(x)
	fmt.Println(y)
}
