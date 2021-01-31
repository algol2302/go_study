package main

import "fmt"

func main() {
	res := foo()
	fmt.Println(res)

	resInt, resString := bar()
	fmt.Println(resInt)
	fmt.Println(resString)
}

func foo() int {
	return 42
}

func bar() (int, string) {
	return 1, "hello"
}
