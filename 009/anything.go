package main

import "fmt"

func main() {
	// 1
	n, err := fmt.Println("Hello")
	fmt.Println(n)
	fmt.Println(err)

	// 2
	foo()

	// 3
	fmt.Println("After foo")

	// 4 loop and conditions
	for i := 0; i < 100; i++ {
		if i%2 == 0 {
			fmt.Println(i)
		}
	}

	bar()
}

func foo() {
	fmt.Println("Foo func")
}

func bar() {
	fmt.Println("and the end!")
}
