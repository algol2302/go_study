package main

import "fmt"

const x int = 42
const y = 42.43

func main() {
	fmt.Printf("%T\t%d\n", x, x)
	fmt.Printf("%T\t%f", y, y)
}
