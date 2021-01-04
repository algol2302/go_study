package main

import "fmt"

func main() {
	x := 42
	fmt.Println(x)

	x = 99
	fmt.Println(x)

	y := x + 24
	fmt.Println(y)

	someText := "Some text"
	anotherText := "Another text"

	sumText := someText + " " + anotherText
	fmt.Println(sumText)
}
