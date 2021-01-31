package main

import "fmt"

func main() {
	defer foo()
	bar()
}

func foo() {
	fmt.Println("print defered function")
}

func bar() {
	fmt.Println("print first function")
}
