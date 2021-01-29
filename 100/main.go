package main

import "fmt"

func main() {
	a := incrementor()
	b := incrementor()
	fmt.Println(a(), b())
	fmt.Println(a(), b())
	fmt.Println(a(), b())
	fmt.Println(a(), b())
	fmt.Println(a(), b())
}

func incrementor() func() int {
	var x int
	return func() int {
		x++
		return x
	}
}
