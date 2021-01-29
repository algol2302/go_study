package main

import "fmt"

func main() {

	items := []int{1, 2, 3, 4, 5, 6, 7, 8, 9}
	res := sum(items...)
	fmt.Println(res)

	res2 := even(sum, items...)
	fmt.Println(res2)

	res3 := odd(sum, items...)
	fmt.Println(res3)

	// test
	res4 := sum(even(sum, items...), odd(sum, items...))
	fmt.Println(res4 == res)
}

func sum(xi ...int) int {
	total := 0

	for _, v := range xi {
		total += v
	}

	return total
}

func even(f func(xi ...int) int, vi ...int) int {
	var yi []int

	for _, v := range vi {
		if v%2 == 0 {
			yi = append(yi, v)
		}
	}

	return f(yi...)
}

func odd(f func(xi ...int) int, vi ...int) int {
	var yi []int

	for _, v := range vi {
		if v%2 != 0 {
			yi = append(yi, v)
		}
	}

	return f(yi...)
}
