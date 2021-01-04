package main

import "fmt"

var y string
var z int

func main() {
	fmt.Println("printing value of y:", y, "ending")
	fmt.Printf("%T\n", y)

	y = "Bond, James Bond"
	fmt.Println("printing value of y:", y, "ending")
	fmt.Printf("%T\n", y)

	fmt.Println("printing value of z:", z, "ending")
	fmt.Printf("%T\n", z)
}
