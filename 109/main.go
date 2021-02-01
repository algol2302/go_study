package main

import "fmt"

func double(x int) int {
	return x * 2
}

func main() {
	g := func(f func(x int) int, x int) int {
		return f(x) * 2
	}

	f := double

	fmt.Println(g(f, 2))
}
