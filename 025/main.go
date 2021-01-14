package main

import "fmt"

var x int
var y float64
var z int8

func main() {
	x = 42
	y = 42.34543
	z = 127

	fmt.Println(x)
	fmt.Println(y)
	fmt.Println(z)
	fmt.Printf("%T\n", x)
	fmt.Printf("%T\n", y)
	fmt.Printf("%T\n", z)
}
