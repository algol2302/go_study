package main

import "fmt"

var y = 42
var z string = "Shaken, not stirred"
var a string = `I'm said
"Shaken,
not stirred"`

func main() {
	fmt.Println(y)
	fmt.Printf("%T\n", y)

	fmt.Println(z)
	fmt.Printf("%T\n", z)

	fmt.Println(a)
	fmt.Printf("%T\n", a)

	// z = 43
	// fmt.Println(z)
	// fmt.Printf("%T\n", z)
}
